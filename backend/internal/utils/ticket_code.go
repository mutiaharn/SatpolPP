package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// GenerateTicketCode generates unique formatted ticket like TRB-202611-4829
func GenerateTicketCode() string {
	now := time.Now()
	datePart := now.Format("200601") // YYYYMM
	
	// generate 4 random alphanumeric chars
	const charset = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	b := make([]byte, 4)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}
	
	return fmt.Sprintf("TRB-%s-%s", datePart, string(b))
}
