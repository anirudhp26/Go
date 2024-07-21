package main

import (
	"encoding/json"
	"fmt"
	"hello/config"
	"hello/controllers"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	client, err := config.SetupClient()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connected to database")
	}
	fmt.Println("Starting server on port: 3000")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		post, err := controllers.CreatePost(client)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
	defer client.Prisma.Disconnect()
}
