package main

import (
	"database/sql"
	"fmt"
	"formative-15/controllers"
	"formative-15/database"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// ENV
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file env")
	} else {
		fmt.Println("successfully read file env")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed!")
		panic(err)
	} else {
		fmt.Println("DB Connection Succeeded!")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// Router GIN
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "GET : /persons, POST : /postpersons, UPDATE : /persons/:id, DELETE : /deletepersons/:id")
	})
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/postpersons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/deletepersons/:id", controllers.DeletePerson)

	router.Run(":8080")
}
