package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

type errorResponse struct {
	Message string `json:"msg"`
}

type getResponse struct {
	Size      int           `json:"size"`
	Customers []interface{} `json:"customers"`
	Message   string        `json:"msg"`
}

type report struct {
	Total   int    `json:"totalCustomers"`
	Period  int    `json:"period"`
	Message string `json:"msg"`
}

var cnt uint64 = 0
var customers []interface{}

// returns customer's index in slice if found. else -1
func findCustomer(Name string, Tel uint64, Address string) (int, model.Customer) {

	for i := 0; i < len(customers); i++ {
		if customers[i] == nil {
			continue
		}

		c, ok := customers[i].(model.Customer)

		if ok {

			if c.Name == Name && c.Tel == Tel && c.Address == Address {
				return i, c
			}

		}
	}
	return -1, model.Customer{}

}

func findByID(ID uint64) (int, model.Customer) {

	for i := 0; i < len(customers); i++ {
		if customers[i] == nil {
			continue
		}

		c, ok := customers[i].(model.Customer)

		if ok {

			if c.ID == ID {
				return i, c
			}

		}
	}
	return -1, model.Customer{}

}

// Create ... creates new customers based on http req body. It raises http forbidden (403) if customer with given info was created before.
func (customer Customer) Create(c echo.Context) error {
	var req request

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	regDate := time.Now().Format("2006-1-2")

	if index, _ := findCustomer(req.Name, req.Tel, req.Address); index == -1 {
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

		if cap(customers) == 0 {
			customers = make([]interface{}, 0, 5)
		}

		if int(cnt-1) == len(customers) {
			customers = append(customers, m)
		} else {
			customers[cnt-1] = m
		}

		fmt.Println(m)
		fmt.Println(customers, cap(customers), len(customers))
		return c.JSON(http.StatusCreated, res)
	}

	res := errorResponse{
		Message: "Customer already exists.",
	}

	return echo.NewHTTPError(http.StatusForbidden, res)

}

func (customer Customer) Update(c echo.Context) error {
	var req request

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errorResponse{
			Message: "Please enter valid url.",
		})
	}

	if index, cus := findByID(uint64(intID)); index != -1 {

		cus.Name = req.Name
		cus.Tel = req.Tel
		cus.Address = req.Address

		res := createResponse{
			Name:         req.Name,
			Tel:          req.Tel,
			Address:      req.Address,
			ID:           cus.ID,
			RegisterDate: cus.RegisterDate,
			Message:      "success",
		}

		customers[index] = cus
		fmt.Println(customers)

		return c.JSON(http.StatusOK, res)
	}

	res := errorResponse{
		// Name:    req.Name,
		// Tel:     req.Tel,
		// Address: req.Address,
		Message: "cID is not available. (Customer with this information doesn't exist)",
	}

	return c.JSON(http.StatusBadRequest, res)
}

func getCustomerArray() ([]interface{}, bool) {

	var newArr []interface{}
	newArr = make([]interface{}, 0, cap(customers))

	// index := 0
	notNull := false
	for i := 0; i < len(customers); i++ {
		if customers[i] != nil {
			notNull = true
			// newArr[index] = customers[i]
			// index++
			newArr = append(newArr, customers[i])
		}
	}

	return newArr, notNull
}

func getCustomersByName(name string) []interface{} {

	var newArr []interface{}
	newArr = make([]interface{}, 0, cap(customers))

	// index := 0
	// notNull := false
	for i := 0; i < len(customers); i++ {
		if customers[i] != nil {

			c := customers[i].(model.Customer)
			if strings.HasPrefix(c.Name, name) {
				// notNull = true
				newArr = append(newArr, customers[i])
			}

		}
	}

	return newArr
}

func (customer Customer) Delete(c echo.Context) error {

	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errorResponse{
			Message: "Please enter valid url.",
		})
	}

	if index, _ := findByID(uint64(intID)); index != -1 {
		customers[index] = nil

		fmt.Println(customers)

		return c.JSON(http.StatusOK, errorResponse{
			Message: "success",
		})
	}
	return c.JSON(http.StatusNotFound, errorResponse{
		Message: "cID is not available. (Customer with this information doesn't exist)",
	})

}

// Get ... handles GET method with or without query parameters. It's case-sensitive in case of query params provided.
func (customer Customer) Get(c echo.Context) error {

	q := c.QueryParam("cName")

	// fmt.Println(q == "")
	if q == "" {
		if newCustomers, ok := getCustomerArray(); ok {
			fmt.Println(newCustomers)
			return c.JSON(http.StatusOK, getResponse{
				Size:      len(newCustomers),
				Customers: newCustomers,
				Message:   "success",
			})
		}

		return c.JSON(http.StatusNotFound, errorResponse{
			Message: "(error) No customers are available!",
		})
	} else {

		matchedCustomers := getCustomersByName(q)
		if len(matchedCustomers) > 0 {
			return c.JSON(http.StatusOK, getResponse{
				Size:      len(matchedCustomers),
				Customers: matchedCustomers,
				Message:   "success",
			})
		}

		return c.JSON(http.StatusNotFound, errorResponse{
			Message: "No customers found with given (prefix) name.",
		})
	}
}

func countNumOfCustomers(month int) int {
	cnt := 0
	if validCustomers, ok := getCustomerArray(); ok {
		for i := 0; i < len(validCustomers); i++ {

			c := validCustomers[i].(model.Customer)
			if intMonth, err := strconv.Atoi(strings.Split(c.RegisterDate, "-")[1]); err == nil && intMonth == month {
				cnt++
			}
		}
	}
	return cnt
}

func (customer Customer) MakeReport(c echo.Context) error {

	month := c.Param("month")
	intMOnth, err := strconv.Atoi(month)
	if err != nil || intMOnth < 0 || intMOnth > 11 {
		return echo.NewHTTPError(http.StatusNotFound, errorResponse{
			Message: "Please enter valid month.",
		})
	}

	if cnt := countNumOfCustomers(intMOnth); cnt > 0 {

		r := report{
			Total:   cnt,
			Period:  1,
			Message: "success",
		}

		return c.JSON(http.StatusOK, r)
	}

	return echo.NewHTTPError(http.StatusNotFound, errorResponse{
		Message: "(error) No customers registered for this month!",
	})
}
