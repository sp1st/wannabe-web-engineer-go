package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// type jsonData struct {
// 	Number int    `json:"number,omitempty"`
// 	String string `json:"string,omitempty"`
// 	Bool   bool   `json:"bool,omitempty"`
// }

type jsonData struct {
	Number int
	String string
	Bool   bool
}

var num int = 0

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/add", addHandler)

	e.GET("/fizzbuzz", fizzbuzzHandler)

	e.GET("/incremental", incrementalHandler)

	e.GET("/ping", pingHandler)

	e.POST("/hello/:name", helloHandler)

	e.POST("/post", postHandler)

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "ペンギンです\n")
	})

	e.GET("/json", jsonHandler)

	e.Start(":8080")
}

type addData struct {
	Right int `json:"right,omitempty"`
	Left  int `json:"left,omitempty"`
}

type answerData struct {
	Answer int `json:"answer,omitempty"`
}

type errorData struct {
	Error int `json:"error,omitempty"`
}

func addHandler(c echo.Context) error {
	var data addData
	var errData errorData
	var ansData answerData

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, errData)
	}
	ansData.Answer = data.Right + data.Left
	return c.JSON(http.StatusOK, ansData)
}

func fizzbuzzHandler(c echo.Context) error {
	count, err := strconv.Atoi(c.QueryParam("count"))
	if err != nil || count < 0 {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	return c.String(http.StatusOK, fizzbuzz(count))
}

func fizzbuzz(count int) string {
	var str string
	for i := 1; i <= count; i++ {
		if i%15 == 0 {
			str += "FizzBuzz"
		} else if i%3 == 0 {
			str += "Fizz"
		} else if i%5 == 0 {
			str += "Buzz"
		} else {
			str += strconv.Itoa(i)
		}
		str += "\n"
	}
	return str
}

func incrementalHandler(c echo.Context) error {
	num++
	return c.String(http.StatusOK, strconv.Itoa(num))
}

func pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong\n")
}

func helloHandler(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello, "+name+".\n")
}

func postHandler(c echo.Context) error {
	var data jsonData

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}
	return c.JSON(http.StatusOK, data)
}

func jsonHandler(c echo.Context) error {
	res := jsonData{
		Number: 10,
		String: "hoge",
		Bool:   false,
	}

	return c.JSON(http.StatusOK, &res)
}
