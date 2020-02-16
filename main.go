    package main

    import (
        "net/http"
        "app/routes"
        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
        "github.com/sirupsen/logrus"
    )

    func init() {
         logrus.SetLevel(logrus.DebugLevel)
         logrus.SetFormatter(&logrus.JSONFormatter{})
    }

    func main() {
        e := echo.New()

        //Middlewares
        e.Use(middleware.Logger())

        // Routes
        routes.Init(e)
        e.GET("/", func(c echo.Context) error {
            return c.String(http.StatusOK, "Hello, World!")
        })
        e.Logger.Fatal(e.Start(":8080"))
    }
