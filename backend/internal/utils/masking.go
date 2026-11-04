package utils

import (
	"strings"
)

// MaskNIK masks confidential citizen NIK (16 digits), showing first 4 and last 2 digits
// e.g. 7371101234567890 -> 7371**********90
func MaskNIK(nik string) string {
	clean := strings.TrimSpace(nik)
	if len(clean) >= 16 {
		return clean[:4] + "**********" + clean[len(clean)-2:]
	}
	if len(clean) > 4 {
		return clean[:2] + strings.Repeat("*", len(clean)-4) + clean[len(clean)-2:]
	}
	return "7371**********12"
}

// MaskName masks confidential citizen Name, showing initial and last char with protected label
// e.g. "Ahmad Pratama" -> "Ah*** P*** (Anonim Terlindungi)"
func MaskName(name string) string {
	clean := strings.TrimSpace(name)
	if clean == "" {
		return "Warga (Anonim Terlindungi)"
	}

	parts := strings.Fields(clean)
	maskedParts := make([]string, 0, len(parts))
	for _, p := range parts {
		runes := []rune(p)
		if len(runes) <= 2 {
			maskedParts = append(maskedParts, string(runes[0])+"*")
		} else {
			maskedParts = append(maskedParts, string(runes[:2])+"***")
		}
	}

	return strings.Join(maskedParts, " ") + " (Anonim Terlindungi)"
}

// MaskPhone masks phone numbers, keeping prefix and last 3 digits
// e.g. 081234567890 -> 0812****890
func MaskPhone(phone string) string {
	clean := strings.TrimSpace(phone)
	if len(clean) >= 10 {
		return clean[:4] + "****" + clean[len(clean)-3:]
	}
	if len(clean) >= 6 {
		return clean[:3] + strings.Repeat("*", len(clean)-5) + clean[len(clean)-2:]
	}
	return "0812****8901"
}
