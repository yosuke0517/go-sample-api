    package main

    import (
        "net/http"
        "./routes"
        "github.com/labstack/echo"
    )

    func main() {
        e := echo.New()
        // Routes
        routes.Init(e)
        e.GET("/", func(c echo.Context) error {
            return c.String(http.StatusOK, "Hello, World!")
        })
        e.Logger.Fatal(e.Start(":8080"))
    }
