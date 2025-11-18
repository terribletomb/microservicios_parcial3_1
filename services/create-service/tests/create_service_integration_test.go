package tests

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoConnection(t *testing.T) {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	if user == "" || pass == "" || host == "" || port == "" {
		t.Fatalf("Variables de entorno faltantes: user='%s', pass='%s', host='%s', port='%s'", user, pass, host, port)
	}

	uri := "mongodb://" + user + ":" + pass + "@" + host + ":" + port
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatalf("No se pudo conectar a MongoDB: %v", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		t.Fatalf("Ping falló: %v", err)
	}
}
