package cmd

import (
	"fmt"

	"github.com/imthaghost/iplogger/server"
)

// LogIP main function for starting the server
func LogIP(args []string) {
	if Tunnel == true {
		// todo create a reverse ssh tunnel
		fmt.Println("tunnel is used")

	} else if Port != "" {
		// user specified port
		server.Start(Port)
	} else {
		// default server
		server.Start("")
	}

}