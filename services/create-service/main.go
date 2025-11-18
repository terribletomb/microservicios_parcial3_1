package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"create-service/controllers"
	"create-service/repositories"
	"create-service/services"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")
	dbName := os.Getenv("MONGO_DB")
	collectionName := os.Getenv("MONGO_COLLECTION")

	var cliente *mongo.Client
	var err error

	// Reintentar conexión a MongoDB hasta que funcione
	for i := 0; i < 10; i++ {
		cliente, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
		if err != nil {
			log.Printf("Error conectando a MongoDB: %v, reintentando en 2s...", err)
			time.Sleep(2 * time.Second)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := cliente.Ping(ctx, nil); err != nil {
			log.Printf("Mongo no responde, reintentando en 2s... (%v)", err)
			time.Sleep(2 * time.Second)
			continue
		}

		// Conexión exitosa
		break
	}

	if err != nil {
		log.Fatal("Mongo no responde después de varios intentos:", err)
	}

	coleccion := cliente.Database(dbName).Collection(collectionName)

	repo := repositories.NuevoLibroRepositorio(coleccion)
	servicio := services.NuevoServicioCrear(repo)
	controlador := controllers.NuevoControladorCrear(servicio)

	r := chi.NewRouter()
	r.Post("/libros", controlador.CrearLibro)

	fmt.Println("📗 Servicio CREATE corriendo en puerto 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
