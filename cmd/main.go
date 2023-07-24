package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"
	"os"

	"github.com/yuiuae/Debugging/internal/handlers"
	"github.com/yuiuae/Debugging/internal/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.Handle("/", middleware.MiddleLog(http.HandlerFunc(handlers.Index)))
	http.Handle("/user", middleware.MiddleLog(http.HandlerFunc(handlers.UserCreate)))
	http.Handle("/user/login", middleware.MiddleLog(http.HandlerFunc(handlers.UserLogin)))
	http.Handle("/admin", middleware.MiddleLog(http.HandlerFunc(handlers.GetUserAll)))
	http.Handle("/actusers", middleware.MiddleLog(http.HandlerFunc(handlers.ActiveUsers)))

	http.Handle("/chat", middleware.MiddleLog(http.HandlerFunc(handlers.RequestWithToken)))

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
