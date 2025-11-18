package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"update-service/controller"
	"update-service/repository"
	"update-service/service"

	"github.com/gorilla/mux"
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
	repo := &repository.LibroRepository{Collection: db.Collection(collectionName)}
	svc := &service.LibroService{Repo: repo}
	ctrl := &controller.LibroController{Service: svc}

	router := mux.NewRouter()
	router.HandleFunc("/libros/{id}", ctrl.ActualizarLibro).Methods("PUT")

	fmt.Println("✏️ Servicio UPDATE corriendo en puerto 8084")
	log.Fatal(http.ListenAndServe(":8084", router))
}
