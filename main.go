package main

import (
	"context"
	"study/postgres"
	"study/simple_connection"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := postgres.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	if err := postgres.InsertRow(ctx, conn); err != nil {
		panic(err)
	}
	if err := postgres.DeleteRow(ctx, conn); err != nil {
		panic(err)
	}
	if err := postgres.SmthSelect(ctx, conn); err != nil {
		panic(err)
	}
}
