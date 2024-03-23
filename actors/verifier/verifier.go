package verifier

import (
	"context"
	"log"
	"ssikr/core"
	"ssikr/protos"
)

type Server struct {
	protos.UnimplementedVerifierServer

	Verifier *Verifier
}

type Verifier struct {
	kms         *core.ECDSAManager
	did         *core.DID
	didDocument *core.DIDDocument
}

func (server *Server) SubmitVP(ctx context.Context, req *protos.SubmitVPRequest) (*protos.SubmitVPResponse, error) {
	log.Printf("VP: %s\n", req.Vp)

	// VC들의 검증자들의 도큐먼트들을 가져와서 검증하는게 다들어있다고 함(?-제대로 들은 건진 잘 모르겠음)
	verify, _, err := core.ParseAndVerifyJwtForVP(req.Vp)

	res := &protos.SubmitVPResponse{Result: "fail"}
	if verify && err == nil {
		res.Result = "ok"
	}

	return res, nil
}
