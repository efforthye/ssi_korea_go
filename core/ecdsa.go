package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/btcsuite/btcd/btcutil/base58" // 비트코인
	"github.com/multiformats/go-multibase"
)

// const로 만들어 놓는 것이 편하기 때문에 밖으로 빼 두었다.
const (
	ELLIPTIC_CURVE = "p256"
)

// 비밀키와 공개키를 담을 구조체
type ECDSAManager struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s Signature) String() string {
	return s.R.String() + s.S.String()
}

// Not use now.
type ECDSAInterface interface {
	Sign(msg string) ([]byte, error)
	SignToString(msg string) (string, error)
	Verify() bool
	Encode() string
	//Decode() string
}

// 프라이빗 키와 퍼블릭 키를 한 번에 넘겨준다. (중요)
// go 언어 특징: 앞을 대문자로 쓰면 퍼블릭 함수, 소문자로 쓰면 프라이빗 함수가 된다(참고!)
// Generate ECDSAManager
func NewEcdsa() (ecdsa *ECDSAManager) {
	ecdsa = new(ECDSAManager)
	err := ecdsa.Generate()
	if err != nil {
		log.Printf("Fail to ECDSA Generate.")
		return nil
	}

	return
}

// Genetate ecdsa keys(P256)
func (e *ECDSAManager) Generate() error {
	// 비밀 키 생성
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // elliptic.p224, elliptic.P384(), elliptic.P521()

	if err != nil {
		return errors.New("ECDSA Keypair generation was Fail!")
	}

	// 퍼블릭 키와 프라이빗 키
	e.PrivateKey = pvKey
	e.PublicKey = &pvKey.PublicKey

	return nil
}

// sign
func (e *ECDSAManager) Sign(digest []byte) (*Signature, error) {
	r := big.NewInt(0)
	s := big.NewInt(0)

	r, s, err := ecdsa.Sign(rand.Reader, e.PrivateKey, digest)
	if err != nil {
		return nil, err //errors.New("failed to sign to msg.")
	}

	// prepare a signature structure to marshal into json
	signature := &Signature{
		R: r,
		S: s,
	}
	/*
		signature := r.Bytes()
		signature = append(signature, s.Bytes()...)
	*/
	return signature, nil
}

// signASN1
func (e *ECDSAManager) SignASN1(digest []byte) ([]byte, error) {
	signature, err := ecdsa.SignASN1(rand.Reader, e.PrivateKey, digest)
	if err != nil {
		return nil, err //errors.New("failed to sign to msg.")
	}

	return signature, nil
}

// Verify
func (e *ECDSAManager) Verify(signature *Signature, digest []byte) bool {
	return ecdsa.Verify(e.PublicKey, digest, signature.R, signature.S)
}

// VerifyASN1
func (e *ECDSAManager) VerifyASN1(signature []byte, digest []byte) bool {
	return ecdsa.VerifyASN1(e.PublicKey, digest, signature)
}

func (e *ECDSAManager) SignToString(digest []byte) (string, error) {
	signature, err := e.Sign(digest)
	if err != nil {
		return "", err
	}

	return signature.String(), nil
}

// 퍼블릭 키 자료형을 문자열으로 변경하는 함수
func (e *ECDSAManager) PublicKeyToString() (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&e.PublicKey)
	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return "", err
	}

	publicKeyHash := sha256.Sum256(publicKeyBytes)

	return hex.EncodeToString(publicKeyHash[:]), nil
}

func (e *ECDSAManager) PublicKeyBase58() string {
	if e.PublicKey == nil {
		return ""
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(e.PublicKey)
	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return ""
	}

	return base58.Encode(publicKeyBytes)
}

func (e *ECDSAManager) PublicKeyMultibase() string {
	if e.PublicKey == nil {
		return ""
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(e.PublicKey)

	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return ""
	}

	// The current multibase table: https://github.com/multiformats/multibase/blob/master/multibase.csv
	str, err := multibase.Encode(multibase.Base58BTC, publicKeyBytes)
	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return ""
	}
	return str
}

func (e *ECDSAManager) PrintPublicKey() {
	str, err := e.PublicKeyToString()
	if err != nil {

	}
	fmt.Printf("Public Key: %s\n", str)
}
