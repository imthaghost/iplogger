package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

type target struct {
	status int
	meta   string
	ip     string
	client string
	data   string
}

// GetIP handler method
func GetIP(c echo.Context) error {
	timestamp := time.Now().UTC()
	// real ip
	ip := c.RealIP()
	// client ip
	client := c.Request().RemoteAddr
	// output file
	file, err := os.OpenFile("log/ip.log", os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	// logging
	// log.SetOutput(file)
	logger := log.New(file, "prefix", log.LstdFlags)
	fmt.Println(&logger)
	logger.Printf("Real ip %s", ip)
	logger.Printf("Client ip %s", client)
	logger.Printf("Time %s", timestamp)
	defer file.Close()
	// redirect to Google
	return c.Redirect(http.StatusPermanentRedirect, "https://google.com")

}
