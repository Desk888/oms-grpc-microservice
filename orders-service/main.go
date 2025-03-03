package main 

import (
	"context"
)

func main() {
	db := NewDB()
	svc := NewService(db)
	svc.CreateOrder(context.Background())
}

