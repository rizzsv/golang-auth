package initializers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectToDB() {
	// Ambil URL dari .env
	uri := os.Getenv("mongoDB")
	if uri == "" {
		log.Fatal("❌ MongoDB URL not found in environment variables")
	}

	// Client options
	clientOpts := options.Client().ApplyURI(uri)

	// Timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB untuk memastikan koneksi berhasil
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ MongoDB ping error: %v", err)
	}

	// Set database - sesuaikan dengan nama database kamu (bukan Cluster0 kalau beda)
	DB = client.Database("Cluster0") // GANTI jika nama databasenya beda

	fmt.Println("✅ Connected to MongoDB!")
}
