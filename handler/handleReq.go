package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hamidrezaRanjbarpour/simple_webservice/model"

	"github.com/labstack/echo/v4"
)

type Customer struct {
}

type request struct {
	Name    string `json:"cName"`
	Tel     uint64 `json:"cTel"`
	Address string `json:"cAddress"`
}

type createResponse struct {
	Name         string `json:"cName"`
	Tel          uint64 `json:"cTel"`
	Address      string `json:"cAddress"`
	ID           uint64 `json:"cID"`
	RegisterDate string `json:"cRegisterDate"`
	Message      string `json:"msg"`
}

func Hello(c echo.Context) error {
	var jsonStr = "{\"name\" : \"Hamid\"}"

	return c.String(http.StatusOK, jsonStr)
}

var cnt uint64 = 0

func (customer Customer) Create(c echo.Context) error {
	var req request

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	regDate := time.Now().Format("2006-1-2")

	cnt++

	m := model.Customer{
		Name:         req.Name,
		Tel:          req.Tel,
		Address:      req.Address,
		ID:           cnt,
		RegisterDate: regDate,
	}

	res := createResponse{
		Name:         req.Name,
		Tel:          req.Tel,
		Address:      req.Address,
		ID:           cnt,
		RegisterDate: regDate,
		Message:      "success",
	}

	fmt.Println(m)

	return c.JSON(http.StatusCreated, res)
}
