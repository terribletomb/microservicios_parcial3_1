package tests

import (
    "context"
    "net/url"
    "os"
    "testing"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoURI() string {
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        return "mongodb://localhost:27017"
    }
    // sanitize: if userinfo exists but username is empty, remove userinfo
    u, err := url.Parse(uri)
    if err != nil {
        return "mongodb://localhost:27017"
    }
    if u.User != nil {
        if u.User.Username() == "" {
            u.User = nil
            return u.String()
        }
    }
    return uri
}

func TestMongoConnection(t *testing.T) {
    uri := getMongoURI()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOpts := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(ctx, clientOpts)
    if err != nil {
        t.Fatalf("No se pudo conectar a MongoDB: %v", err)
    }
    if err := client.Ping(ctx, nil); err != nil {
        t.Fatalf("No se pudo hacer ping a MongoDB: %v", err)
    }
}

