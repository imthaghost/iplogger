package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// GetIP handler method
func GetIP(c echo.Context) error {
	// real ip
	ip := c.RealIP()
	// client ip
	client := c.Request().RemoteAddr
	// output file
	file, err := os.OpenFile("log/ip.log", os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// logging
	log.SetOutput(file)
	log.Println(ip)
	log.Println(client)
	log.Printf("Real ip %s", ip)
	log.Printf("Client ip %s", client)

	// redirect to Google
	return c.Redirect(http.StatusPermanentRedirect, "https://google.com")

}
