// 열린 서버들을 가지고 실제 통신을 하는 부분(중요)
package core

import (
	"context"
	"errors"
	"fmt"
	"log"
	"ssikr/config"
	"ssikr/protos"
	"time"

	"google.golang.org/grpc"
)

// 등록할 did와 didDocument 를 매개변수로 받는다.
func RegisterDid(did string, didDocument string) error {
	// 컨피그에서 연결 정보를 가져온다. (커넥션 정보: conn)
	conn, err := grpc.Dial(config.SystemConfig.RegistrarAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Registrar not connect: %v\n", err)
		return errors.New(fmt.Sprintf("Registrar not connect: %v", err))
	}
	defer conn.Close()

	// GRPC 레지스트라 클라이언트 생성
	client := protos.NewRegistrarClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 중요한 건 RegisterDid 함수 명이다.
	// 블룸RPC에서 실제 메서드 서버 호출했던 것과 똑같은 부분이다. (중요)
	res, err := client.RegisterDid(ctx, &protos.RegistrarRequest{Did: did, DidDocument: didDocument})
	if err != nil {
		log.Println("Failed to register DID.")
		return errors.New("Failed to register DID.")
	}

	// 응답을 콘솔로그에 찍어 정상적인지 확인한다.
	fmt.Printf("Registrar Response: %s\n", res)

	return nil
}

// did를 가지고 didDocument를 찾는 함수이다.
func ResolveDid(did string) (string, error) {
	conn, err := grpc.Dial(config.SystemConfig.ResolverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Resolver not connect: %v\n", err)
		return "", errors.New(fmt.Sprintf("Resolver not connect: %v", err))
	}
	defer conn.Close()

	client := protos.NewResolverClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ResolveDid(ctx, &protos.ResolverRequest{Did: did})
	if err != nil {
		log.Fatalf("Failed to resolve DID.")
	}

	fmt.Printf("Result: %s\n", res)

	// 가져온 didDocument 값을 리턴한다.
	return res.DidDocument, nil
}
