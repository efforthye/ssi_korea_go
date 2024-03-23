// core/did.go
// 나중에 이거 사용할 때, core.NewDID 이러한 형식으로 사용 가능하다.
package core

import (
	"errors"
	"fmt"
	"ssikr/util"
)

type DID struct {
	did    string
	method string
}

func NewDID(method string, pbKey string) (did *DID, err error) {
	if method == "" || pbKey == "" {
		return nil, errors.New("parameter is not valid")
	}

	var newDid = new(DID)
	newDid.method = method
	specificIdentifier := util.MakeHashBase58(pbKey)
	// DID:Method:specific
	newDid.did = fmt.Sprintf("did:%s:%s", method, specificIdentifier)

	return newDid, nil
}

func (d *DID) String() string {
	return d.did
}
