package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Desk888/common"
	pb "github.com/Desk888/common/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type db struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var (
	mongoURI = common.EnvString("MONGODB_URI", "mongodb://localhost:27017/orders")
)

func NewDB() *db {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")
	collection := client.Database("orders").Collection("orders")

	return &db{
		client:     client,
		collection: collection,
	}
}

func (d *db) Create(ctx context.Context) error {
	// This method is now redundant as we're initializing the DB in NewDB
	// We'll keep it for interface compatibility
	return nil
}

func (d *db) SaveOrder(order *pb.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// If order has no ID, create a new one
	if order.Id == "" {
		order.Id = primitive.NewObjectID().Hex()
	}

	// Convert to BSON document
	doc := bson.M{
		"_id":         order.Id,
		"customer_id": order.CustomerId,
		"status":      order.Status,
		"items":       order.Items,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	}

	// Use upsert to create or update
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": order.Id}
	update := bson.M{"$set": doc}

	_, err := d.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Printf("Failed to save order: %v", err)
		return err
	}

	return nil
}

func (d *db) GetOrderById(id string) (*pb.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result bson.M
	err := d.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("order not found")
		}
		log.Printf("Failed to get order: %v", err)
		return nil, err
	}

	// Convert BSON to Order
	order := &pb.Order{
		Id:         result["_id"].(string),
		CustomerId: result["customer_id"].(string),
		Status:     result["status"].(string),
	}

	// Extract items
	if items, ok := result["items"].(primitive.A); ok {
		for _, item := range items {
			if itemMap, ok := item.(bson.M); ok {
				var quantity int32
				if q, ok := itemMap["quantity"].(int32); ok {
					quantity = q
				} else if q, ok := itemMap["quantity"].(int64); ok {
					quantity = int32(q)
				} else if q, ok := itemMap["quantity"].(float64); ok {
					quantity = int32(q)
				}

				order.Items = append(order.Items, &pb.ItemsWithQuantity{
					Id:       itemMap["id"].(string),
					Quantity: quantity,
				})
			}
		}
	}

	return order, nil
}

func (d *db) DeleteOrder(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	_, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Failed to delete order: %v", err)
		return err
	}

	return nil
}
