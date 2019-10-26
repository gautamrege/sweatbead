package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/urfave/negroni"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/gautamrege/packt/sweatbead/usermgr/config"
	"github.com/gautamrege/packt/sweatbead/usermgr/db"
	"github.com/gautamrege/packt/sweatbead/usermgr/logger"
	"github.com/gautamrege/packt/sweatbead/usermgr/service"
	// @grpc - Enable
	pb "github.com/gautamrege/packt/sweatbead/proto/usermgr"
)

func main() {
	config.Load()

	logger.Init()
	db.Init()

	deps := service.Dependencies{
		DB: db.GetStorer(db.Get()),
	}

	// mux router
	router := service.InitRouter(deps)

	// init web server
	server := negroni.Classic()
	server.UseHandler(router)

	port := config.AppPort() // This should be changed to the service port number via argument or environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	// @grpc
	go GRPCServe()
	server.Run(addr)
}

func GRPCServe() {
	host := config.ReadEnvString("GRPC_HOST")
	port := config.ReadEnvInt("GRPC_PORT")
	tls := config.ReadEnvBool("TLS")
	certFile := config.ReadEnvString("CERT_FILE")
	keyFile := config.ReadEnvString("KEY_FILE")

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logger.Get().Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if tls {
		if certFile == "" {
			logger.Get().Fatalf("No certificate file specified")
		}
		if keyFile == "" {
			logger.Get().Fatalf("No key file specified")
		}
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			logger.Get().Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	s := service.GrpcServer{
		DB: db.GetStorer(db.Get()),
	}
	logger.Get().Infof("Grpc config: %v", s)

	// @grpc - Enable after proto file is in place
	pb.RegisterUserMgrServer(grpcServer, &s)

	logger.Get().Infof("Listening for gRPC on %s:%d", host, port)

	grpcServer.Serve(lis)
}
