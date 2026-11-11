package domain

import (
	"time"
)

// User Role constants
type Role string

const (
	RoleSuperadmin Role = "SUPERADMIN_DISKOMINFO" // Developer / IT Diskominfo
	RoleAdmin      Role = "ADMIN_SATPOL"          // Kasat / Pimpinan
	RoleStaff      Role = "STAFF_SATPOL"          // Staff Pengelola Informasi / Regu Patroli
	RolePublic     Role = "PUBLIC"                // Warga Masyarakat
)

// User represents an administrative user
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	Email     string    `json:"email" gorm:"size:100;uniqueIndex;not null"`
	Username  string    `json:"username" gorm:"size:50;uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"` // Hashed with bcrypt
	Role      Role      `json:"role" gorm:"size:30;not null;default:'STAFF_SATPOL'"`
	Bidang    string    `json:"bidang" gorm:"size:100"` // Sekretariat, Trantibum, Perda, Linmas, Binmas
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Article represents news & announcements on the portal
type Article struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"size:255;not null"`
	Slug        string    `json:"slug" gorm:"size:255;uniqueIndex;not null"`
	Summary     string    `json:"summary" gorm:"type:text"`
	Content     string    `json:"content" gorm:"type:longtext;not null"`
	Thumbnail   string    `json:"thumbnail" gorm:"size:255"`
	Category    string    `json:"category" gorm:"size:50;index"` // Berita, Pengumuman, Agenda
	AuthorID    uint      `json:"author_id"`
	AuthorName  string    `json:"author_name" gorm:"size:100"`
	ViewCount   int       `json:"view_count" gorm:"default:0"`
	IsPublished bool      `json:"is_published" gorm:"default:true;index"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PublicDocument represents official government documents (Renja, DPA, LAKIP, etc.)
type PublicDocument struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Title         string    `json:"title" gorm:"size:255;not null"`
	DocNumber     string    `json:"doc_number" gorm:"size:100"`
	Category      string    `json:"category" gorm:"size:50;index"` // renja, renstra, dpa, lakip, keuangan, rpjmd, perda, sop
	Year          int       `json:"year" gorm:"index"`
	FileURL       string    `json:"file_url" gorm:"size:255;not null"`
	FileSize      string    `json:"file_size" gorm:"size:20"`
	DownloadCount int       `json:"download_count" gorm:"default:0"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Gallery represents photo/event documentation
type Gallery struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"type:text"`
	ImageURL    string    `json:"image_url" gorm:"size:255;not null"`
	EventDate   time.Time `json:"event_date"`
	CreatedAt   time.Time `json:"created_at"`
}

// ComplaintStatus for SIAP TERTIB
type ComplaintStatus string

const (
	StatusPending    ComplaintStatus = "MENUNGGU_VERIFIKASI"
	StatusDisposed   ComplaintStatus = "DISPOSISI_BIDANG"
	StatusInProgress ComplaintStatus = "PENINDAKAN_LAPANGAN"
	StatusCompleted  ComplaintStatus = "SELESAI"
	StatusRejected   ComplaintStatus = "DITOLAK"
)

// Complaint represents public reports from SIAP TERTIB
type Complaint struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	TicketCode    string `json:"ticket_code" gorm:"size:30;uniqueIndex;not null"` // e.g. TRB-202611-0012
	IsAnonymous   bool   `json:"is_anonymous" gorm:"default:false"`
	ReporterNIK   string `json:"reporter_nik" gorm:"size:16"`   // Encrypted or hidden if anonymous
	ReporterName  string `json:"reporter_name" gorm:"size:100"` // Masked on public endpoints
	ReporterPhone string `json:"reporter_phone" gorm:"size:25"`
	ReporterEmail string `json:"reporter_email" gorm:"size:100"`

	// Violation Category
	ViolationType string `json:"violation_type" gorm:"size:100;not null;index"`

	// Geographic Location (Sulawesi Selatan)
	RegencyID     string  `json:"regency_id" gorm:"size:50;not null"`
	RegencyName   string  `json:"regency_name" gorm:"size:100"`
	DistrictName  string  `json:"district_name" gorm:"size:100"`
	VillageName   string  `json:"village_name" gorm:"size:100"`
	AddressDetail string  `json:"address_detail" gorm:"type:text;not null"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`

	// Description & Photos
	Description      string    `json:"description" gorm:"type:text;not null"`
	IncidentDate     time.Time `json:"incident_date"`
	EvidenceImageURL string    `json:"evidence_image_url" gorm:"type:text"`
	EvidencePhotos   string    `json:"evidence_photos" gorm:"type:text"`

	// Status & Follow Up
	Status      ComplaintStatus `json:"status" gorm:"size:30;default:'MENUNGGU_VERIFIKASI';index"`
	DisposedTo  string          `json:"disposed_to" gorm:"size:100"`
	ActionNotes string          `json:"action_notes" gorm:"type:text"`
	ActionProof string          `json:"action_proof" gorm:"type:text"`
	OfficerID   *uint           `json:"officer_id"`
	OfficerName string          `json:"officer_name" gorm:"size:100"`
	CompletedAt *time.Time      `json:"completed_at"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// SiteSetting represents anti-hardcoded dynamic configuration
type SiteSetting struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Key       string    `json:"key" gorm:"size:100;uniqueIndex;not null"`
	Value     string    `json:"value" gorm:"type:text"`
	Group     string    `json:"group" gorm:"size:50;default:'general'"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuditLog tracks sensitive operations by staff and superadmins
type AuditLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"index"`
	UserName  string    `json:"user_name" gorm:"size:100"`
	Action    string    `json:"action" gorm:"size:100;not null"`
	Resource  string    `json:"resource" gorm:"size:100"`
	Details   string    `json:"details" gorm:"type:text"`
	IPAddress string    `json:"ip_address" gorm:"size:45"`
	CreatedAt time.Time `json:"created_at"`
}

// Regency master data for Sulawesi Selatan (24 Kab/Kota)
type Regency struct {
	ID   string `json:"id" gorm:"primaryKey;size:10"`
	Name string `json:"name" gorm:"size:100;not null"`
}

// OrganizationNode represents a dynamic SOTK structure unit editable via admin panel
type OrganizationNode struct {
	ID            uint     `json:"id" gorm:"primaryKey"`
	OrgCode       string   `json:"org_code" gorm:"size:30;uniqueIndex;not null"`
	ParentCode    *string  `json:"parent_code" gorm:"size:30"`
	PositionTitle string   `json:"position_title" gorm:"size:150;not null"`
	OfficialName  string   `json:"official_name" gorm:"size:150;not null"`
	NIP           string   `json:"nip,omitempty" gorm:"size:40"`
	Rank          string   `json:"rank,omitempty" gorm:"size:50"`
	Description   string   `json:"description,omitempty" gorm:"type:text"`
	SubUnits      []string `json:"sub_units" gorm:"type:text"`
	PhotoURL      string   `json:"photo_url,omitempty" gorm:"size:255"`
	Level         int      `json:"level" gorm:"default:2"`
	SortOrder     int      `json:"sort_order" gorm:"default:1"`
}
