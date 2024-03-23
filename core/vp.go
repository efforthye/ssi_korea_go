package core

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/getlantern/deepcopy"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type VP struct {
	Context      []string `json:"@context"`
	Id           string   `json:"id,omitempty"`
	Type         []string `json:"type,omitempty"`
	Issuer       string   `json:"issuer,omitempty"`
	IssuanceDate string   `json:"issuanceDate,omitempty"`

	// jwt의 token형식으로 저장한다.
	VerifiableCredential []string `json:"verifiableCredential"`
	// VP의 Proof 사용.
	Proof *Proof `json:"proof,omitempty"`
}

// JWT를 위한 claim
type JwtClaimsForVP struct {
	jwt.StandardClaims

	Nonce string
	Vp    VP `json:"vp,omitempty"`
}

// VP 만들기 (중요)
func NewVP(id string, typ []string, issuer string, vcTokens []string) (*VP, error) {
	newVP := &VP{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/v2",
		},
		Id:                   id,
		Type:                 typ,
		Issuer:               issuer,
		VerifiableCredential: vcTokens, // VC를 JWT로 바꾼 것이다.
	}
	return newVP, nil
}

// JWT 생성
func (vp *VP) GenerateJWT(verificationId string, pvKey *ecdsa.PrivateKey) string {
	aud := ""
	exp := time.Now().Add(time.Minute * 5).Unix()       //만료 시간. 현재 + 5분
	jti := uuid.NewString()                             // JWT ID
	t, err := time.Parse(time.RFC3339, vp.IssuanceDate) //unixtime으로 바꾸기 위해.
	iat := t.Unix()
	nbf := iat
	iss := vp.Issuer // VP에서는 홀더의 아이디가(????)
	sub := "Verifiable Presentation"

	// Proof를 제거하고 JWT를 만들기 위해 복제한다.
	vpTmp := new(VP)
	deepcopy.Copy(vpTmp, vp)
	vpTmp.Proof = nil

	jwtClaims := JwtClaimsForVP{
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
		*vpTmp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwtClaims)
	token.Header["kid"] = verificationId

	tokenString, err := token.SignedString(pvKey)

	if err != nil {

	}

	return tokenString
}

// JWT 검증 (전달받은 토큰, 검증할 퍼블릭 키)
func (vp *VP) VerifyJwt(tokenString string, pbKey *ecdsa.PublicKey) (bool, error) {

	parts := strings.Split(tokenString, ".")
	err := jwt.SigningMethodES256.Verify(strings.Join(parts[0:2], "."), parts[2], pbKey)
	if err != nil {
		return false, nil
	}

	// 개별적으로 내부 VC들을 다시 다 검증해야 한다.
	// 클레임을 먼저 파싱한다.
	parseToken, err := jwt.ParseWithClaims(tokenString, &JwtClaimsForVP{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			log.Fatalln("unexpected signing method.")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		//did := token.Header["kid"].(string)
		//pbKeyBase58 := did // getPbKey(did, "") //DID를 통해 DID-Document의 pbKey를 구한다.
		//pbKey, _ := x509.ParsePKIXPublicKey(base58.Decode(pbKeyBase58))

		// 함수로 받아온 퍼블릭 키를 바로 리턴하였기 때문에 위 코드는 주석처리 하였음.
		return pbKey, nil
	})

	fmt.Println("parseToken: ", parseToken)
	claims, ok := parseToken.Claims.(*JwtClaimsForVP)
	fmt.Println("claims: ", claims)

	if ok && parseToken.Valid {

		if &claims.Vp != nil {
			// VP가 있으면 VP를 가져와서 VP 내부의 VC를 모두 가져온다. (다수의 VC List)
			vpMapClaims := claims.Vp
			vcList := vpMapClaims.VerifiableCredential

			// VP 내의 모든 VC를 하나하나 돌면서 하나씩 전부 검증한다.
			for _, vcToken := range vcList {
				fmt.Println("VC: ", vcToken)
				//verify jwt
				//VC.VerifyJwt(vcToken)

			}
		} else {
			return false, fmt.Errorf("VC is not exist.")
		}

	} else {
		return false, fmt.Errorf("VP is not valid.")
	}

	return true, nil
}
