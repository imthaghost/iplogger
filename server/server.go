package server

import (
	"github.com/imthaghost/iplogger/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Start the logging server
func Start(port string) {
	e := echo.New()

	// Log
	e.Use(middleware.Logger())

	// Recover
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Root
	e.GET("/", controllers.GetIP)

	// user specified port
	if port != "" {
		e.Logger.Fatal(e.Start(":" + port))
	}
	// default
	e.Logger.Fatal(e.Start(":8000"))
}
