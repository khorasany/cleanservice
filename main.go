package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kianooshaz/cleanservice/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
