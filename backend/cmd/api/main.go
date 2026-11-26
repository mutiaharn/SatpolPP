package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpadapter "satpolpp-sulsel-api/internal/adapters/http"
	"satpolpp-sulsel-api/internal/adapters/http/handlers"
	"satpolpp-sulsel-api/internal/adapters/repository/postgres"
	"satpolpp-sulsel-api/internal/adapters/storage"
	"satpolpp-sulsel-api/internal/config"
	"satpolpp-sulsel-api/internal/worker"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	cfg := config.LoadConfig()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 1. Connect to PostgreSQL via PgBouncer (SimpleProtocol)
	db, err := postgres.NewDB(ctx, cfg)
	if err != nil {
		slog.Error("Gagal terhubung ke PostgreSQL/PgBouncer", "error", err)
		os.Exit(1)
	}
	defer db.Pool.Close()

	// 2. Run Idempotent Seeder
	if err := db.SeedIfEmpty(ctx); err != nil {
		slog.Error("Gagal menjalankan seeder awal database", "error", err)
		os.Exit(1)
	}

	// 3. Initialize Unified Storage (RustFS S3 or Dev Local Fallback)
	store, err := storage.NewUnifiedStorage(cfg)
	if err != nil {
		slog.Error("Gagal menginisialisasi storage RustFS", "error", err)
		os.Exit(1)
	}

	// 4. Start Background Email Outbox Worker
	outbox := worker.NewOutboxWorker(db, cfg)
	go outbox.Start(ctx)

	// 5. Wire Modular Handlers & Router
	authH := handlers.NewAuthHandler(db, cfg)
	compH := handlers.NewComplaintHandler(db, store, cfg)
	portalH := handlers.NewPortalHandler(db)
	uploadH := handlers.NewUploadHandler(store)

	router := httpadapter.NewRouter(cfg, db, authH, compH, portalH, uploadH)

	srv := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		slog.Info("Satpol PP Sulsel & SIAP TERTIB Backend API aktif",
			"port", cfg.AppPort,
			"env", cfg.AppEnv,
			"db_host", cfg.DBHost,
			"s3_endpoint", cfg.S3Endpoint,
		)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("HTTP server error", "error", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	slog.Info("Menerima sinyal shutdown, menghentikan server dengan aman...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownCtx)
}
