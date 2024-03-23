package main

import (
	"log"
	"net"
	"ssikr/actors/verifier"
	"ssikr/config"
	"ssikr/protos"

	"google.golang.org/grpc"
)

func main() {
	// New Verifier
	vrfr := new(verifier.Verifier)
	//issuer.generateDID()

	lis, err := net.Listen("tcp", config.SystemConfig.VerifierAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	verifierServer := verifier.Server{}
	verifierServer.Verifier = vrfr

	s := grpc.NewServer()
	protos.RegisterVerifierServer(s, &verifierServer)

	log.Printf("Issuer Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// export CONFIG_FILE="./config/config.json"
// go run ./actors/verifier/cmd/verifierd.go
