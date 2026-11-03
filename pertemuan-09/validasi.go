package main

import (
	"strings"
	"unicode/utf8"
)

// ValidationError membawa alasan penolakan per field, mis.
// {"judul": "tidak boleh kosong"}. (Level 7)
type ValidationError struct {
	Fields map[string]string
}

// Error mengembalikan teks "validasi gagal" (supaya *ValidationError memenuhi
// interface error). (Level 7)
func (e *ValidationError) Error() string {
	return "validasi gagal"
}

// ValidasiTugas memeriksa input tugas. Kembalikan nil kalau valid, atau
// *ValidationError berisi SEMUA field yang bermasalah. (Level 7)
func ValidasiTugas(in taskInput) error {
	judul := strings.TrimSpace(in.Judul)
	if judul == "" {
		return &ValidationError{
			Fields: map[string]string{
				"judul": "tidak boleh kosong",
			},
		}
	}

	if utf8.RuneCountInString(judul) > 100 {
		return &ValidationError{
			Fields: map[string]string{
				"judul": "maksimal 100 karakter",
			},
		}
	}

	return nil
}
