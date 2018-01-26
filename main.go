package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Jump struct {
	number   int
	altitude int
	dropzone string
}

func main() {
	db, err := sql.Open("postgres", "dbname=jumps sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM jumps")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	jumps := make([]*Jump, 0)
	for rows.Next() {
		jump := new(Jump)
		err := rows.Scan(&jump.number, &jump.altitude, &jump.dropzone)
		if err != nil {
			log.Fatal(err)
		}
		jumps = append(jumps, jump)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, jump := range jumps {
		fmt.Printf("Logbook entry: Jump # %v, altitude %v, dropzone %v ", jump.number, jump.altitude, jump.dropzone)
	}
}
