package services

import (
	// "ButterHost69/PKr-base/encrypt"
	"ButterHost69/PKr-base/models"
	"ButterHost69/PKr-base/pb"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/peer"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type BackgroundServer struct {
	pb.UnimplementedBackgroundServiceServer
}



func (s *BackgroundServer) GetPublicKey(ctx context.Context, request *emptypb.Empty)(*pb.PublicKey, error) {
	keyData, err := ReadPublicKey(); 
	p, _ := peer.FromContext(ctx)
  	ip := p.Addr.String()
	if err != nil {
		logentry := "Could Not Provide Public Key To IP: " + ip
		models.AddUsersLogEntry("[X]",logentry)
		models.AddUsersLogEntry("[X]",err)

		return &pb.PublicKey{
			Key: "",
		}, err
	}
	logentry := "Successfully Provided Public Key To IP: " + ip
	models.AddUsersLogEntry("[X]", logentry)

	return &pb.PublicKey{
		Key: keyData,
	}, nil

}

func (s *BackgroundServer) InitNewWorkSpaceConnection (ctx context.Context, request *pb.InitRequest)(*pb.InitResponse, error){
	// 1. Decrypt password [X]
	// 2. Authenticate Request [X]
	// 3. Add the New Connection to the .PKr Config File [X]
	// 4. Store the Public Key [X]
	// 5. Send the Response with port [X]
	// 6. Open a Data Transfer Port and shit [Will be a separate Function not here] [X]
	
	p, _ := peer.FromContext(ctx)
  	ip := p.Addr.String()

	encrypted_password := request.Password


	// This is Comment temporary
	// password, err := encrypt.DecryptData(encrypted_password)
	password := encrypted_password

	// UNCOMMENT THIS SHIT --------
	// if err != nil {
	// 	AddUserLogEntry("Failed to Init Workspace Connection for User IP: " + ip)
	// 	AddUserLogEntry(err)	
	// 	return &pb.InitResponse{
	// 		Response: 4000,
	// 		Port: 0000,
	// 	}, nil
	// }

	// Authenticates Workspace Name and Password and Get the Workspace File Path
	file_path := models.AuthenticateWorkspaceInfo(request.WorkspaceName, password)
	if file_path == "" {
		models.AddLogEntry(request.WorkspaceName, "Failed to Init Workspace Connection for User IP: " + ip)
		return &pb.InitResponse{
			Response: 4000,
			Port: 0000,
		}, nil
	}

	var connection models.Connection
	connection.CurrentIP = ip
	connection.CurrentPort = request.Port

	// Save Public Key
	keysPath, err := models.StorePublicKeys(file_path + "\\.PKr\\keys\\", request.PublicKey)
	if err != nil {
		models.AddLogEntry(request.WorkspaceName, "Failed to Init Workspace Connection for User IP: " + ip)
		models.AddLogEntry(request.WorkspaceName, err)
		return &pb.InitResponse{
			Response: 4000,
			Port: 0000,
		}, err
	}	
	
	// Store the New Connection in the .PKr Config file
	connection.PublicKeyPath = keysPath
	if err := models.AddConnectionToPKRConfigFile(file_path + "\\.PKr\\workspaceConfig.json", connection); err != nil {
		models.AddLogEntry(request.WorkspaceName, "Failed to Init Workspace Connection for User IP: " + ip)
		models.AddLogEntry(request.WorkspaceName, err)
		return &pb.InitResponse{
			Response: 4000,
			Port: 0000,
		}, err
	}
	models.AddLogEntry(request.WorkspaceName, fmt.Sprintf("Added User with IP: %v to the Connection List", ip))

	// Start New Data grpc Server and Transfer Data
	var portchan chan int
	var errorchan chan error
	
	go StartDataServer(1024 * time.Second, request.WorkspaceName,file_path, portchan, errorchan)
	select {
	case port_num := <- portchan:
		return &pb.InitResponse{
			Response: 200,
			Port: int32(port_num),
		}, nil	
	case err := <- errorchan:
		return &pb.InitResponse{
			Response: 4000,
			Port: 0000,
		}, err	
	}
	// 
}
