package dialer

import (
	"fmt"
	"net/rpc"

	"github.com/ButterHost69/kcp-go"
)

const (
	HANDLER_NAME = "Handler"
)

func call(rpcname string, args interface{}, reply interface{}, ripaddr, lipaddr string) error {

	conn, err := kcp.Dial(ripaddr, lipaddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	c := rpc.NewClient(conn)
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err != nil {
		return err
	}

	return nil
}

func (h *CallHandler) CallRegisterWorkspace(server_ip, username, password, workspace_name string) error {

	var req RegisterWorkspaceRequest
	var res RegisterUserResponse

	req.Username = username
	req.Password = password
	req.WorkspaceName = workspace_name

	if err := call(HANDLER_NAME + ".RegisterWorkspace", req, &res, server_ip, ""); err != nil{

		return fmt.Errorf("Error in Calling RPC...\nError: %v", err)
	}

	if res.Response != 200 {
		return fmt.Errorf("Calling Ping Method was not Successful.\nReturn Code - %d", res.Response)
	}

	return nil
}

func (h *CallHandler) CallRegisterUser(server_ip, username, password string) (string, error) {
	var req RegisterUserRequest
	var res RegisterUserResponse

	req.Username = username
	req.Password = password
	
	if err := call(HANDLER_NAME + ".RegisterUser", req, &res, server_ip, ""); err != nil{
		return "", fmt.Errorf("Error in Calling RPC...\nError: %v", err)
	}

	if res.Response != 200 {
		return "", fmt.Errorf("Calling CallRegisterUser Method was not Successful.\nReturn Code - %d", res.Response)
	}

	return res.UniqueUsername, nil
}