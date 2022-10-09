package main

import (
	"fmt"
	"log"

	"order/ent"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s",
		"test",
		"yrv1+LtyRjLUb7QbVSlxXlxjcJ8=",
		"172.16.20.232",
		"3311",
		"nt_order",
		"utf8mb4",
		"12s",
	)
	fmt.Println(dbDSN)
	client, err := ent.Open("mysql", dbDSN)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	n, err := client.OrderMain.
		Query().
		Where().
		Count(ctx)
	fmt.Println(n)
}
