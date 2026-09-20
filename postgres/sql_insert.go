package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		INSERT INTO users (title,description,completed,created_at)
		VALUES('Домашка','Сделать домашку по матеше',FALSE,'2026-04-11 18:00:05');
	`
	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
