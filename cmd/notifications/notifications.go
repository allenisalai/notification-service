package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
	"log"
	"github.com/allenisalai/notification-service/internal"
	"github.com/allenisalai/notification-service/internal/controller"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	enableGracefulShutdown()

	db := database.GetDb()
	defer db.Close()

	r := mux.NewRouter()
	r.Use(jsonReturnContentTypeMiddleware)
	r.Use(checkIncomingContentType)

	r.HandleFunc("/notification-type", controller.NotificationTypeCGet()).Methods("GET")
	r.HandleFunc("/notification-type/{id}", controller.NotificationTypeGet()).Methods("GET")
	r.HandleFunc("/notification-type", controller.NotificationTypePost()).Methods("POST")
	r.HandleFunc("/notification-type/{id}", controller.NotificationTypePatch()).Methods("PATCH")
	r.HandleFunc("/notification-type/{id}", controller.NotificationTypeDelete()).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func enableGracefulShutdown() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		if sig == syscall.SIGTERM || sig == syscall.SIGINT {
			os.Exit(0)
		}
	}()
}

func jsonReturnContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func checkIncomingContentType(next http.Handler) http.Handler {
	return handlers.ContentTypeHandler(next, "application/json")
	/*
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(415)
			return
		}

		next.ServeHTTP(w, r)
	})*/
}
