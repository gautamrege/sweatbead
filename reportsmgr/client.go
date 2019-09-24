package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	sweatmgr "github.com/gautamrege/packt/sweatbead/proto/sweatmgr"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:33010", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

// printFeatures lists all the features within the given bounding Rectangle.
func getSweatStats(client sweatmgr.SweatMgrClient, req *sweatmgr.SweatStatsRequest) (res *sweatmgr.SweatStatsResponse) {
	log.Printf("Getting listing all sweat objects")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.GetSweatStats(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetSweatStats(_) = _, %v: ", client, err)
	}
	log.Println(res)
	return
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := sweatmgr.NewSweatMgrClient(conn)

	// Looking for a valid feature
	getSweatStats(client, &sweatmgr.SweatStatsRequest{Userid: "5d6c2bca147e62d574c73f18"})

	time.Sleep(2 * time.Second)
}
