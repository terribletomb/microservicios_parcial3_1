package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"read-service/controller"
	"read-service/repository"
	"read-service/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")
	dbName := os.Getenv("MONGO_DB")
	collectionName := os.Getenv("MONGO_COLLECTION")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error conectando a MongoDB:", err)
	}

	db := client.Database(dbName)
	repo := &repository.LibroRepositoryMongo{Collection: db.Collection(collectionName)}
	svc := &service.LibroService{Repo: repo}
	ctrl := &controller.LibroController{Service: svc}

	http.HandleFunc("/libros", ctrl.ObtenerLibros)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "OK")
	})

	fmt.Println("📘 Servicio READ corriendo en puerto 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
