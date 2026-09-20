package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func SmthSelect(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		SELECT id, title, description, completed, created_at, completed_at 
		FROM users	
	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		var description string
		var completed bool
		var created_at time.Time
		var completed_at *time.Time

		err := rows.Scan(
			&id,
			&title,
			&description,
			&completed,
			&created_at,
			&completed_at,
		)
		if err != nil {
			return err
		}
		fmt.Println(id, title, description, completed, created_at, completed_at)
	}
	return err
}
