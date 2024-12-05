package root

import (
	"fmt"

	"github.com/ButterHost69/PKr-cli/dialer"
	"github.com/ButterHost69/PKr-cli/models"
)

func Server_Setup() {
	// [X] Create a main Server Json file in tmp(root) dir
	// [X] Allow user to connect to multiple server
	// [X] Store Server IP, and your username and password (user can have multiple username and password)
	// [X] Send Create User request to server, save to Server Json file and display the username to user at terminal

	var server_ip string
	fmt.Print("Please Enter Server IP: ")
	fmt.Scan(&server_ip)

	var server_username string
	fmt.Print("Please Enter A Username for Server Connection: ")
	fmt.Scan(&server_username)

	var server_password string
	fmt.Print("Please Enter A Password for Server Connection: ")
	fmt.Scan(&server_password)

	if server_ip == "" || server_username == "" || server_password == "" {
		fmt.Println("ip or username or password cannot be empty")
		return
	}

	username, err := dialer.RegisterServer(server_ip, server_username, server_password)
	if err != nil {
		fmt.Println("error Occured in Sending Request to server")
		fmt.Println(err)
		return
	}
	fmt.Printf("Registered to server:%s as username: %s\n", server_ip, username)

	// TODO: [ ] If this function fails revert/delete the user from the server - by sending some delete request
	err = models.AddNewServerToConfig(username, server_password, server_ip)
	if err != nil {
		fmt.Println("Error Occured in Adding Server to serverConfig.json")
		fmt.Println(err)
		return
	}

	fmt.Println("Entry added to serverConfig.json file")
	return
}