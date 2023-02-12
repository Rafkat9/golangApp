package main

import (
	"fmt"
	"log"

	// "myappcorn/internal/adapters"
	"myappcorn/loadEnv"
	"myappcorn/pkg/db/postgres"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

//	func GetGames() {
//		fmt.Println("Hello")
//	}

func init() {
	//.env
	loadEnv.LoadEnv()
	//postgresdb
	postgres.Connect()

}
func main() {
	///////////////////////////////////// const env

	loadEnv.LoadEnv()
	log.Println("create const env")
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env read function")
	}
	///////////////////////////////////////db connect
	log.Println("db connect")

	////////////////////////////////////////gin start
	log.Println("gin start")
	r := gin.Default()
	// r := adapters.Handler
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	///////////////////////////////////////////serve run
	log.Println("serve run")
	port := myEnv["PORT"]
	srv := &http.Server{
		Addr:           port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("ssssssssssStarting server at port %s \n", port)
	// if err := srv.ListenAndServe(); err != nil {
	// 	log.Fatal(err)
	// }
	srv.ListenAndServe()
	log.Println("connect serve")
}
