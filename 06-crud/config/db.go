package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Conexion() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB_CONECTION"))

	if err != nil {
		log.Fatal("DB Error...")
	}

	fmt.Println("Conectado a la base")

	return db
}
