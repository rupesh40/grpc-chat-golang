package main

import (
	//"grpcChatServer/chatserver"
	"grpcbidir/chatserver"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main(){
	Port:= os.Getenv("PORT");
	if Port ==""  {
		Port = "5000"
	}
	listner,err := net.Listen("tcp", ":" + Port);
	if err!=nil {
		log.Fatalf(" could notnlisten on port @ %v :: " + Port,err)
	}
	log.Printf("listenig on port %v" , Port);

	//new grpc server instance created 
	grpcserver := grpc.NewServer()
	
	//register chat service
	cs := chatserver.ChatServer{}
	chatserver.RegisterServicesServer(grpcserver, &cs)
	//grpc serve and listen and serve
	err = grpcserver.Serve(listner)
	if err!=nil{
		log.Fatalf("failed to start grpc server ::%v " , err);
		} 

}