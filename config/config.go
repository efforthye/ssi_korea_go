// config/config.go
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	SystemConfig Configuration
)

// 패키지가 로드될 때 가장 먼저 호출되는 함수
// init()은 누가 호출해주지 않아도 자동으로 실행된다.
func init() {
	// config 파일의 환경변수 값을 가져온다.
	configFileName := os.Getenv("CONFIG_FILE")
	if configFileName == "" {
		configFileName = "./config/config.json"
	}
	SystemConfig = loadConfigration(configFileName)
}

// JSON struct
type Configuration struct {
	IssuerAddr           string `json:"issuer_addr"`
	VerifierAddr         string `json:"verifier_addr"`
	RegistrarAddr        string `json:"registrar_addr"`
	RegistrarGatewayAddr string `json:"registrar_gateway_addr"`
	ResolverAddr         string `json:"resolver_addr"`
	ResolverGatewayAddr  string `json:"resolver_gateway_addr"`
}

// 읽어와서 집어넣어 주는 역할을 한다.
func loadConfigration(path string) Configuration {
	var configuration Configuration
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&configuration)
	return configuration
}
