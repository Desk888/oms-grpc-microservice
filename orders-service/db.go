package main

import (
	"context"
)

type db struct {
	// add mongodb here instance here
}

func NewDB() *db {
	return &db{}
}

func (d *db) Create(context.Context) error {
	return nil
}