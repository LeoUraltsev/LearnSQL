package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Открытие соединения с базой данных, первый параметр это имя используемой БД, второй параметр строка с параметрами подключения к БД.
	db, err := sql.Open("postgres", "user=leo password=1234 host=localhost dbname=learn_sql sslmode=disable")
	if err != nil {
		log.Fatalf("connection to db failed: %v", err)
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("connection to db failed: %v", err)
	}
	fmt.Printf("Success conection to db \n")

	q := `SELECT * FROM spendings WHERE price > $1 AND user_id = $2`
	price := 7500
	id := 7
	var sp Spending

	rows, err := db.Query(q, price, id)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&sp.Id, &sp.Price, &sp.CreatedAt, &sp.UserID, &sp.CategoryID)
		if err != nil {
			log.Fatalf("scan error: %v", err)
		}
		fmt.Println(sp)
	}

}
