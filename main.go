package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数から接続情報を取得
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	run_port := os.Getenv("RUN_PORT")
	
	// PostgreSQLの接続文字列を作成
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname,
	)
	// データベース接続を初期化
	db,err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to Supabase: %v", err)
	}
	defer db.Close()

	log.Println("Successfully connected to Supabase!")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// サーバーを起動
	e.Logger.Fatal(e.Start(run_port))
}
