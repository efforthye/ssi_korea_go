// TODO. 주석 채우기(카카오톡 나에게 보기 참고)
// core/did_document.go

package core

import "encoding/json"

const (
	VERIFICATION_KEY_TYPE_SECP256K1 = "EcdsaSecp256k1VerificationKey2019"
	VERIFICATION_KEY_TYPE_JWS       = "JsonWebKey2020"
	VERIFICATION_KEY_TYPE_X25519    = "X25519KeyAgreementKey2019"
	VERIFICATION_KEY_TYPE_ED25519   = "Ed25519VerificationKey2018"

	PROOF_TYPE_ED25519 = "Ed25519Signature2018"
	PROOF_TYPE_JWS     = "JsonWebSignature2020"
)

type didDocumentInterface interface {
	produce(doc DIDDocument) string
	consume(str string) DIDDocument
}

// DID Document
// Decentralized Identifiers (DIDs) v1.0 https://www.w3.org/TR/did-core/
// Simple Example: https://www.w3.org/TR/did-core/#a-simple-example
type DIDDocument struct {
	// @context
	// Mendatory
	Context []string `json:"@context"`

	Id                   string               `json:"id"`
	AlsoKnownAs          []string             `json:"alsoKnownAs,omitempty"`
	Controller           string               `json:"controller,omitempty"`
	VerificationMethod   []VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication       []Authentication     `json:"authentication,omitempty"`
	AssertionMethod      string               `json:"assertionMethod,omitempty"`
	KeyAgreement         string               `json:"keyAgreement,omitempty"`
	CapabilityInvocation string               `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation string               `json:"capabilityDelegation,omitempty"`
	Service              []Service            `json:"service,omitempty"`
}

type VerificationMethod struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	Controller         string `json:"controller"`
	PublicKeyMultibase string `json:"PublicKeyMultibase,omitempty"`
}

type Authentication struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	Controller         string `json:"controller"`
	PublicKeyBase58    string `json:"publicKeyBase58,omitempty"`
	PublicKeyMultibase string `json:"PublicKeyMultibase,omitempty"`
}

type Service struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

func NewDIDDocument(did string, verificationMethod []VerificationMethod) (doc *DIDDocument) {
	var docTmp = new(DIDDocument)
	docTmp.Context = []string{"https://www.w3.org/ns/did/v1"}
	docTmp.Id = did
	docTmp.VerificationMethod = verificationMethod

	return docTmp
}

// 구조체로 만들어진 것을 string 형식으로 리턴하는 함수
func NewDIDDocumentForString(docStr string) (didDoc *DIDDocument, err error) {
	didDoc = new(DIDDocument)
	e := json.Unmarshal([]byte(docStr), didDoc)
	if e != nil {
		return nil, e
	}

	return didDoc, nil
}

////////////////////////////////////////////////
// methods define
////////////////////////////////////////////////

// https://www.w3.org/TR/did-core/#production-and-consumption
func (doc *DIDDocument) Produce() string {
	// object -> string
	str, err := json.Marshal(doc)
	if err != nil {
		// fmt.Println("Fail to marshal: ", err)
		return ""
	}
	return string(str)
}

// Not use
func (doc *DIDDocument) Consume(str string) {
	// string -> object
	err := json.Unmarshal([]byte(str), doc)
	if err != nil {

	}
}

func (doc *DIDDocument) GetVerificationMethod() []VerificationMethod {
	return doc.VerificationMethod
}

// 구조체 그대로 데이터를 집어넣고 set(여러개 들어갈 수 있음)이므로 기존 것 뒤에 붙여넣고 리턴
func (doc *DIDDocument) AddVerificationMethod(id string, typ string, controller string, publicKeyMultibase string) {
	newVm := VerificationMethod{
		Id:                 id,
		Type:               typ,
		Controller:         controller,
		PublicKeyMultibase: publicKeyMultibase,
	}

	doc.VerificationMethod = append(doc.VerificationMethod, newVm)
}

func (doc *DIDDocument) String() string {
	return doc.Produce()
}
