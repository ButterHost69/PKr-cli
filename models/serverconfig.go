package models

import (
	"encoding/json"
	"fmt"
	// "log"
	"os"
)

type ServerConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ServerIP string `json:"server_ip"`

	SendWorkspaces []BetterSendWorkspace `json:"send_workspace"`
	GetWorkspaces  []BetterGetWorkspace  `json:"get_workspace"`
}

type ServerFileConfig struct {
	ServerLists []ServerConfig `json:"server_lists"`
}

func CreateServerConfigFiles() error {
	serverConfig := ServerFileConfig{}
	jsonbytes, err := json.Marshal(serverConfig)
	if err != nil {
		return fmt.Errorf("~ Unable to Parse JSON to server config file.\nError:%v", err)
	}
	err = os.WriteFile(ROOT_DIR+"/serverConfig.json", jsonbytes, 0777)
	if err != nil {
		return fmt.Errorf("error could not write server config file.\nError:%v", err)
	}
	return nil
}

func writeToServerConfigFile(newUserConfig ServerFileConfig) error {
	jsonData, err := json.MarshalIndent(newUserConfig, "", "	")
	// fmt.Println(jsonData)
	if err != nil {
		fmt.Println("error occured in Marshalling the data to JSON")
		fmt.Println(err)
		return err
	}

	// fmt.Println(string(jsonData))
	err = os.WriteFile(SERVER_CONFIG_FILE, jsonData, 0777)
	if err != nil {
		fmt.Println("error occured in storing data in server file")
		fmt.Println(err)
		return err
	}

	return nil
}

func readFromServerConfigFile() (ServerFileConfig, error) {
	file, err := os.Open(SERVER_CONFIG_FILE)
	if err != nil {
		fmt.Println("error in opening config file.... pls check if tmp/userConfig.json available ")
		return ServerFileConfig{}, err
	}
	defer file.Close()

	var serverConfig ServerFileConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&serverConfig)
	if err != nil {
		fmt.Println("error in decoding json data")
		return ServerFileConfig{}, err
	}

	// fmt.Println(userConfig)
	return serverConfig, nil
}

func AddNewServerToConfig(username, password, serverip string) error {
	serverConfig, err := readFromServerConfigFile()
	if err != nil {
		fmt.Println("Error in reading From the ServerConfig File...")
		return err
	}

	sconf := ServerConfig{
		Username: username,
		Password: password,
		ServerIP: serverip,
	}

	serverConfig.ServerLists = append(serverConfig.ServerLists, sconf)
	if err := writeToServerConfigFile(serverConfig); err != nil {
		fmt.Println("Error Occured in Writing To the Server File")
		return err
	}

	return nil
}

func GetServerUsernamePassword(server_ip string) (string, string, error) {
	serverConfig, err := readFromServerConfigFile()
	if err != nil {
		fmt.Println("Error in reading From the ServerConfig File...")
		return "", "", err
	}

	for _, server := range serverConfig.ServerLists {
		if server.ServerIP == server_ip {
			return server.Username, server.Password, nil
		}
	}
	return "", "", fmt.Errorf("error could not find server with ip: %s", server_ip)
}
