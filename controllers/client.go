package controllers

import (
	"fmt"
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

// check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetIP handler method
func GetIP(c echo.Context) error {
	// timestamp
	timestamp := time.Now().String()
	// real ip
	ip := string(c.RealIP())
	// client ip
	client := string(c.Request().RemoteAddr)
	// each visitor
	visitor := fmt.Sprintf("ip: %s", ip) + " | " + fmt.Sprintf("client: %s ", client) + " - " + timestamp + "\n"
	// output file
	file, err := os.Create("/tmp/ip.log")
	// uh oh
	check(err)
	// write to file
	file.WriteString(visitor)
	// redirect to Google
	return c.Redirect(http.StatusPermanentRedirect, "https://google.com")

}
