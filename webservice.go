package main

import (
	"fmt"

	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"

	"github.com/hamidrezaRanjbarpour/simple_webservice/handler"
	"github.com/labstack/echo/v4"
	// "github.com/stretchr/testify/http"
)

func main() {

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect to database")
	// }

	e := echo.New()

	e.GET("/hello", handler.Hello)
	e.POST("/customers", handler.Customer{}.Create)
	e.PUT("/customers/:id", handler.Customer{}.Update)
	e.DELETE("/customers/:id", handler.Customer{}.Delete)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		fmt.Println(err)
	}

}
