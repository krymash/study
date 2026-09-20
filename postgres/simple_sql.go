package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQeury := `
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			title VARCHAR(200) NOt NULL,
			description VARCHAR(1000) NOT NULL,
			completed BOOLEAN NOT NULL,
			created_at TIMESTAMP NOT NULL,
			completed_at TIMESTAMP

		);
	`
	_, err := conn.Exec(ctx, sqlQeury)
	if err != nil {
		return err
	}
	return nil
}
