package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Method: GET, POST, PUT, DELETE, PATCH dll
// Status Code: 2XX, 3XX, 4XX, 5XX

type UserAddressInfo struct {
	City string `json:"city"`
}

type UserInfo struct {
	ID      int              `json:"id"`
	Name    string           `json:"name"`
	Age     int              `json:"age"`
	Address *UserAddressInfo `json:"address"`
}

type UserCreateRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserUpdateRequest struct {
	Age int `json:"age"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {

	rateLimiterConfig := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 500, ExpiresIn: 1 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	echoServer := echo.New()
	echoServer.Use(middleware.SecureWithConfig(middleware.DefaultSecureConfig))
	echoServer.Use(middleware.RateLimiterWithConfig(rateLimiterConfig))

	// Path / route / endpoint

	// Query Param
	// http://localhost:8080/user?id=1&name=dayat&age=24
	echoServer.GET("/user", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		name := c.QueryParam("name")
		age, _ := strconv.Atoi(c.QueryParam("age"))

		user1 := &UserInfo{
			ID:   id,
			Name: name,
			Age:  age,
		}

		return c.JSON(http.StatusOK, user1)
	})

	// Path Param
	// http://localhost:8080/user/:id
	echoServer.GET("/user/:id", func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		user1 := &UserInfo{
			ID:   id,
			Name: "Alan",
			Age:  24,
			Address: &UserAddressInfo{
				City: "Jakarta Barat",
			},
		}

		return c.JSON(http.StatusOK, user1)
	})

	echoServer.POST("/user", func(c echo.Context) error {
		body := &UserCreateRequest{}
		c.Bind(body)

		if body.Name == "" {
			return c.JSON(http.StatusBadRequest, &Response{
				Message: "Nama tidak boleh kosong",
			})
		}

		return c.JSON(http.StatusCreated, body)
	})

	echoServer.PUT("/user/:id", func(c echo.Context) error {
		body := &UserUpdateRequest{}
		c.Bind(body)

		return c.JSON(http.StatusOK, body)
	})

	echoServer.DELETE("/user/:id", func(c echo.Context) error {
		log.Println("Deleted user", c.Param("id"))
		return c.JSON(http.StatusOK, &Response{
			Message: "User successfully deleted",
		})
	})

	err := echoServer.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
