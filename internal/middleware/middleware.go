package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
)

func MiddleLog(next http.Handler) http.Handler {
	fmt.Println("Start program1")
	defer fmt.Println("Stop program1")

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request -> [%s] %q", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
