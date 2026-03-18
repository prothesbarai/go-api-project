package database

import (
	"context"
	"database/sql"
	"fmt"
	"go-api-project/logger"
	"os"
	"time"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectionDB(){
	// >>> First Load .env File
	err := godotenv.Load()
	if(err != nil){
		logger.AppLogger.Error.Println("Error loading .env file",err)
	}
	/// >>> Then Get env file link
	dsn := os.Getenv("DB_DSN")
	fmt.Println(dsn)
	
	/// >>> Open DB connection
	db,err := sql.Open("mysql",dsn)
	if(err != nil){
		logger.AppLogger.Error.Println("Database connection failed:",err)
	}

	/// >>> For Production Purpose stablility
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second) // >> If DB response takes more than 5 seconds → operation will be canceled
	defer cancel()

	/// >>> It won't work just by opening the connection, test the real connection. DB alive কিনা check
	if err := db.PingContext(ctx); err != nil{
		logger.AppLogger.Error.Println("Database unreachable:",err)
	}
 
	db.SetMaxOpenConns(25) // >> max active connection limit. মানে: একসাথে maximum 25টা DB connection open হতে পারবে
	db.SetMaxIdleConns(25) // >> idle connection pool. মানে: idle (unused) connection pool এ রাখা হবে
	db.SetConnMaxLifetime(5*time.Minute) // >> connection lifetime control. মানে: একটা connection 5 মিনিট পর expire হবে
	DB = db
	fmt.Print("Database Connected")
}