package holder

import (
	"context"
	"fmt"
	"log"
	"ssikr/config"
	"ssikr/core"
	"ssikr/protos"
	"time"

	"google.golang.org/grpc"
)

type Holder struct {
	Kms         *core.ECDSAManager
	Did         *core.DID
	DidDocument *core.DIDDocument
}

func (holder *Holder) GenerateDID() {
	// 키생성(ECDSA) - 향후 KMS로 대체.
	holder.Kms = core.NewEcdsa()

	// DID 생성.
	did, _ := core.NewDID("comnic", holder.Kms.PublicKeyBase58())

	holder.Did = did

	// DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", did)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               core.VERIFICATION_KEY_TYPE_SECP256K1,
			Controller:         did.String(),
			PublicKeyMultibase: holder.Kms.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(did.String(), verificationMethod)
	holder.DidDocument = didDocument
}

func RequestVC(holderDid string) string {
	conn, err := grpc.Dial(config.SystemConfig.IssuerAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Issuer not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewSimpleIssuerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Issuer로 부터 VC를 요청한다.
	res, err := c.IssueSimpleVC(ctx, &protos.MsgIssueVC{
		Did:   holderDid,
		Nonce: "100001",
		// 이슈어(발급자)도 검증자가 되어야 한다. (?)
		ReqVP: "{'Korea eID Credential': '12345'}",
	})
	if err != nil {
		log.Fatalf("could not request: %v", err)
	}

	return res.GetVc()
}

func RegisterDid(did string, document *core.DIDDocument) error {
	err := core.RegisterDid(did, document.String())
	if err != nil {
		return err
	}
	return nil
}

// 검증자에게 VP를 보낸다.
func SubmitVP(vpToken string) error {
	conn, err := grpc.Dial(config.SystemConfig.VerifierAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Verifier not connect: %v", err)
		return err
	}
	defer conn.Close()
	c := protos.NewVerifierClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Issuer로 부터 VC를 요청한다.
	res, err := c.SubmitVP(ctx, &protos.SubmitVPRequest{
		Vp: vpToken,
	})
	if err != nil {
		log.Printf("could not request: %v", err)
		return err
	}

	fmt.Printf("Verifier's response: %s\n", res.Result)

	return nil
}
