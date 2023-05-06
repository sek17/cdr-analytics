package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sek17/go-activedirectory/handlers"
)

var (
	db *sql.DB
)

func main() {
	var err error
	db, err = sql.Open("mysql", "root:admin@tcp(localhost:3306)/cdr")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	// serve static files
	r.Static("/css", "./css")
	r.Static("/js", "./js")
	r.Static("/assets", "./assets")
	r.Static("/base", "./base")
	r.Static("/icons", "./icons")
	r.Static("/vendors", "./vendors")

	r.GET("/", handlers.GetDashboardHandler(db))
	r.GET("/create-email", handlers.GetAliasesHandler(db))
	r.POST("/add-user", handlers.CreateUserHandler(db))
	r.POST("/add-positional", handlers.CreatePositionalHandler(db))
	r.DELETE("/delete-email/:id", handlers.DeleteEmailHandler(db))
	r.POST("/edit-email/:id", handlers.EditEmailHandler(db))
	/*	r.GET("/add-alias", handlers.GetUsers(db))
		r.GET("/reset-password", handlers.GetPassword(db))
		r.GET("/enable-disable-user", handlers.GetEnableDisable(db))
		r.GET("/migrate-user", handlers.GetMigrate(db))
		r.GET("/add-mac-address", handlers.GetMacAddress(db))
		r.GET("/manage-user", handlers.GetManageUsers(db))

	*/
	r.Run("0.0.0.0:8888")
}
