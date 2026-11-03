package main

import (
	"log"
	"net/http"
	"time"
)

// TODO: lihat SOAL.md untuk kontrak lengkap. Ganti setiap "panic" dengan
// implementasi yang benar; tambahkan import (log, time, runtime/debug, mime,
// ...) sendiri kalau dibutuhkan.

// Middleware membungkus sebuah http.Handler dengan perilaku tambahan.
type Middleware func(http.Handler) http.Handler

// Chain membungkus h dengan mws. Urutan: middleware PERTAMA di daftar adalah
// yang PALING LUAR (dijalankan paling dulu saat request masuk). (Level 3)
func Chain(h http.Handler, mws ...Middleware) http.Handler {
	panic("belum diimplementasikan")
}

// statusRecorder membungkus ResponseWriter untuk mengingat status code yang
// ditulis handler. Status awal (kalau handler tak pernah memanggil
// WriteHeader) adalah 200. (Level 1-2)
type statusRecorder struct {
	http.ResponseWriter
	status  int
	written bool
}

// WriteHeader mencatat status lalu meneruskannya ke ResponseWriter asli.
// Hanya panggilan PERTAMA yang dicatat. (Level 2)
func (sr *statusRecorder) WriteHeader(code int) {
	if !sr.written {
		sr.status = code
		sr.written = true
	}
	sr.ResponseWriter.WriteHeader(code)
}

func (sr *statusRecorder) Write(b []byte) (int, error) {
	sr.written = true
	return sr.ResponseWriter.Write(b)
}

// Logging mencatat SATU baris per request lewat log.Printf dengan format
// "%s %s %d %v" -> method, path, status, durasi. (Level 1-2)
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mulai := time.Now()
		rec := &statusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		next.ServeHTTP(rec, r)
		durasi := time.Since(mulai)
		log.Printf("%s %s %d %v", r.Method, r.URL.Path, rec.status, durasi)
	})
}

// Recover menangkap panic dari handler di bawahnya: catat ke log (pesan +
// stack trace), balas 500 JSON generik. (Level 4-6)
func Recover(next http.Handler) http.Handler {
	panic("belum diimplementasikan")
}

// RequireJSON menolak POST/PUT/PATCH yang Content-Type-nya bukan
// application/json dengan 415. (Level 9)
func RequireJSON(next http.Handler) http.Handler {
	panic("belum diimplementasikan")
}

// NewHandler menyusun API + seluruh middleware menjadi satu http.Handler yang
// siap dipakai server. (Level 10)
func NewHandler(store TaskStore) http.Handler {
	panic("belum diimplementasikan")
}
