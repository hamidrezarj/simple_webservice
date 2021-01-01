package main

import (
	"fmt"
	"time"

	"github.com/hamidrezaRanjbarpour/simple_webservice/handler"
	"github.com/labstack/echo/v4"
	// "github.com/stretchr/testify/http"
)

func main() {

	// fmt.Println(time.Now().Year + "-" + time.Now().Month + "-" + time.Now().Day)
	fmt.Println(time.Now().Format("2006-1-2"))

	e := echo.New()

	e.GET("/hello", handler.Hello)
	e.POST("/customers", handler.Customer{}.Create)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		fmt.Println(err)
	}

}
