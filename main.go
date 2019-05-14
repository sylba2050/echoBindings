package main

import (
    "os"
    "fmt"
    "net/http"

    "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    e.GET("/", test)
    e.POST("/", test)

    e.Start(":8080")
}

type myStruct struct {
    UserId string `json:"userid" form:"userid" query:"userid"`
    PW string `json:"pw" form:"pw" query:"pw"`
}

func test(c echo.Context) error {
    s := new(myStruct)
    if err := c.Bind(s); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return err
    }

    fmt.Fprintln(os.Stderr, s.UserId)

    return c.NoContent(http.StatusOK)
}
