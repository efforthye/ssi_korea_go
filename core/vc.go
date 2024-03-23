package core

import (
	"crypto/ecdsa"
	"strings"
	"time"

	"github.com/getlantern/deepcopy"
	"github.com/golang-jwt/jwt"
)

// https://www.w3.org/TR/vc-data-model
// https://www.w3.org/TR/vc-data-model/#json-web-token
type VC struct { // VC 객체
	// Mendatory
	Context []string `json:"@context"`

	// Id: vc의 유일한 아이디
	Id   string   `json:"id,omitempty"`
	Type []string `json:"type,omitempty"`
	// Issuer: VC를 발행한 사람이 누군지 적는다.
	Issuer            string                 `json:"issuer,omitempty"`
	IssuanceDate      string                 `json:"issuanceDate,omitempty"`
	CredentialSubject map[string]interface{} `json:"credentialSubject,omitempty"`
	Proof             *Proof                 `json:"proof,omitempty"`
}

// 형식, 언제, 어떤키서명, 서명값뭐고, 자바스크립트웹서명 등 프로퍼티 설명은 위 문서에 있다.
type Proof struct {
	Type               string `json:"type,omitempty"`
	Created            string `json:"created,omitempty"`
	ProofPurpose       string `json:"proofPurpose,omitempty"`
	VerificationMethod string `json:"verificationMethod,omitempty"`
	ProofValue         string `json:"proofValue,omitempty"`
	Jws                string `json:"jws,omitempty"`
}

// JWT를 위한 claim (중요)
type JwtClaims struct {
	// jwt에서 기본적으로 제공해주는 프로퍼티를 자동으로 상속받아온다.
	jwt.StandardClaims

	// 추가적으로 넣는 것은 여기에, Nonce: 구분값
	Nonce string
	// Vc: 실제 VC (중요)
	Vc VC `json:"vc,omitempty"`
}

// 새로운 VC 생성 함수 (중요)
func NewVC(id string, typ []string, issuer string, credentialSubject map[string]interface{}) (*VC, error) {
	newVC := &VC{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/v2",
		},
		Id:                id,
		Type:              typ,
		Issuer:            issuer,
		IssuanceDate:      time.Now().Format(time.RFC3339), //"2010-01-01T19:23:24Z",
		CredentialSubject: credentialSubject,
	}
	return newVC, nil
}

type VCInterface interface {
	GenerateJWT() string
	VerifyJwt() (bool, error)
}

// 이 아래부터는 JWT 관련 함수들임.
// VC를 JTW로 생성하고 string으로 반환한다. (VC객체에 해당하는 JWT를 생성하는 함수이다.)
// JTW의 경우 JWS로 증명되기에 Proofs를 빼고, JWT와 중복되는 properties를 제거한다.
func (vc *VC) GenerateJWT(verificationId string, pvKey *ecdsa.PrivateKey) (string, error) {
	// JWT 속성 정의
	aud := ""
	exp := time.Now().Add(time.Minute * 5).Unix()       //만료 시간. 현재 + 5분
	jti := "1112342"                                    // JWT ID (일단 아무거나 넣었음)
	t, err := time.Parse(time.RFC3339, vc.IssuanceDate) // unixtime으로 날짜 형식을 바꾸기 위해.
	iat := t.Unix()                                     // 이슈어 엣
	nbf := iat                                          // 지금 발행했지만 아직 유요하지않다(?) 이런의미
	iss := vc.Issuer                                    // 발행자의 DID
	sub := "Verifiable Credential"                      // 제목

	// Proof를 제거하고 JWT를 만들기 위해 복제한다.
	vcTmp := new(VC)
	deepcopy.Copy(vcTmp, vc)
	vcTmp.Proof = nil

	// JWT 형식으로 값을 채워준다. (상속 받아서 JWT에 필요한 값들을 넣어준다.)
	jwtClaims := JwtClaims{
		jwt.StandardClaims{
			Audience:  aud,
			ExpiresAt: exp,
			Id:        jti,
			IssuedAt:  iat,
			Issuer:    iss,
			NotBefore: nbf,
			Subject:   sub,
		},
		"qwasd!234",
		*vcTmp, // VC가 실제로 여기 담긴다. (중요)
	}

	// 구조 생성이 다 되었으니, JWT . . 구분하는 string 토큰으로 만들어 준다.
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwtClaims)
	// kid: key-id 이다. VC를 검증할때 어떤 키를 쓸 것인가를 여기에서 가져온다. (중요)
	token.Header["kid"] = verificationId // 중요한 값이다.

	// 토큰에 대해서 프라이빗 키로 서명을 해준다.
	tokenString, err := token.SignedString(pvKey)

	if err != nil {

	}

	return tokenString, nil
}

// 토큰을 받아서 pbKey로 유효한지 검증을 한다.
func (vc *VC) VerifyJwt(token string, pbKey *ecdsa.PublicKey) (bool, error) {
	// .으로 구분해서 맨 뒤에있는게 서명이다.
	parts := strings.Split(token, ".")
	// 검증해서 err가 있으면 에러를 넘기고 없으면 true를 넘겨 유효함을 리턴한다.
	err := jwt.SigningMethodES256.Verify(strings.Join(parts[0:2], "."), parts[2], pbKey)
	if err != nil {
		return false, err
	}

	return true, nil
}
