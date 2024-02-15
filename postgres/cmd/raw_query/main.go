package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v5"
)

const dbDSN = "host=localhost port=54321 dbname=auth user=auth-user password=Su23&cGR sslmode=disable"

func main() {
	ctx := context.Background()

	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func() {
		if err = con.Close(ctx); err != nil {
			log.Printf("Error when closing connection: %v", err)
		}

	}()

	res, err := con.Exec(ctx, "INSERT INTO tuser (uname, email, urole) VALUES ($1, $2, $3)", gofakeit.BeerName(), gofakeit.Email(), gofakeit.Number(1, 2))
	if err != nil {
		log.Fatalf("failed to insert new row to table tuser: %v", err)
	}
	resoutput := res.RowsAffected()
	log.Printf("inserted %v rows.", resoutput)

	rows, err := con.Query(ctx, "SELECT id, uname, email, urole, created_at, updated_at FROM tuser")
	if err != nil {
		log.Fatalf("failed to select users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var uname, email string
		var urole int
		var createdAt time.Time
		var updatedAt sql.NullTime
		err := rows.Scan(&id, &uname, &email, &urole, &createdAt, &updatedAt)
		if err != nil {
			log.Fatalf("failed to scan user: %v", err)
		}
		log.Printf("id: %v, uname: %v, email: %v, urole: %v, created_at: %v, updated_at: %v", id, uname, email, urole, createdAt, updatedAt)

	}

}
