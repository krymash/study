package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQeury := `
	DELETE FROM users
	WHERE id=1
	`
	_, err := conn.Exec(ctx, sqlQeury)

	return err
}
