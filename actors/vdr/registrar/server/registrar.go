// vdr 데이터를 저장할 데이터 서버를 띄워 놓는다.
// 서버 만드는 첫번째 파일(이후 비슷하다)
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"ssikr/config"
	"ssikr/protos"

	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc"
)

type registrarServer struct {
	// 이미 만들어져 있으므로 아래 한 줄을 똑같이 앞으로도 사용
	protos.UnimplementedRegistrarServer
}

// 실제 레지스트라 함수의 기본 - RegisterDid
func (server *registrarServer) RegisterDid(ctx context.Context, req *protos.RegistrarRequest) (*protos.RegistrarResponse, error) {
	// did와 did document 리턴
	log.Printf("Register DID: %s\n", req.Did)
	log.Printf("Register DID Document: %s\n", req.DidDocument)

	// 레벨디비를 사용하기 위해 특정 경로의 레벨디비를 오픈
	db, err := leveldb.OpenFile("did_db/dids", nil)
	if err != nil {
		panic(err)
	}
	// 함수가 종료될 때 열었던 데이터베이스를 닫는드ㅏ.
	defer db.Close()

	// (키, 값)
	err = db.Put([]byte(req.Did), []byte(req.DidDocument), nil)

	return &protos.RegistrarResponse{Result: "OK"}, nil
}

// 특정 포트를 얼어서 대기(listen) - 실습이 아니므로 제거
func main() {
	fmt.Println("### Start Registrar ###")
	lis, err := net.Listen("tcp", config.SystemConfig.RegistrarAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := registrarServer{}
	s := grpc.NewServer()
	// 등록 했던 것을 띄워주고
	protos.RegisterRegistrarServer(s, &server)

	log.Printf("Registrar Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// go mod tidy
// export CONFIG_FILE="./config/config.json"
// echo $CONFIG_FILE
// go run ./actors/vdr/registrar/server/registrar.go
// lsof -i TCP:9000
