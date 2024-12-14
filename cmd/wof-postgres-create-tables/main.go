package main

// go run cmd/wof-postgres-create-tables/main.go -database-uri 'sql://postgres?dsn=user=asc host=localhost dbname=whosonfirst sslmode=disable' -properties

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-whosonfirst-database-postgres"

	"github.com/whosonfirst/go-whosonfirst-database/app/sql/tables/create"
)

func main() {

	ctx := context.Background()
	err := create.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to create database tables, %v", err)
	}
}
