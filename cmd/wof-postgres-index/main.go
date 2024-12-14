package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-whosonfirst-database-postgres"

	"github.com/whosonfirst/go-whosonfirst-database/app/sql/index"
)

func main() {

	ctx := context.Background()
	err := index.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to index database, %v", err)
	}
}
