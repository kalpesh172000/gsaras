package main

func main() {
	httpServer := NewHttpServer(":9001")
	go httpServer.Run()
	
	grpcServer := NewGRPCServer(":9000")
	grpcServer.Run()


}
