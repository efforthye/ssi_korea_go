// core/did_test.go
// 테스트: 형식(앞에 테스트라고 앞글자 대문자로 쓰면 테스트 코드로 인식한다.)
// 테스트 코드에서는 맞는지 검증도 중요하지만 틀린지 검증도 중요하다.

package core

import (
	"fmt"
	"strings"
	"testing"
)

// 이 아래 왼쪽에 있는 녹색 삼각형 run 버튼을 누르면 실행된다.
func TestGenerateDID(t *testing.T) {
	method := "comnic"
	// 키생성(ECDSA) - 향후 KMS로 대체.
	kms := NewEcdsa()

	// DID 생성.
	did, err := NewDID(method, kms.PublicKeyBase58())

	// 에러를 뱉지 않으면 정상적으로 판단.
	if err != nil {
		t.Error("Failed to generate DID.")
	}

	// 형식에 맞는지 확인
	if did == nil || !strings.HasPrefix(did.String(), fmt.Sprintf("did:%s:", method)) {
		t.Error("Failed to generate DID.")
	}
}
