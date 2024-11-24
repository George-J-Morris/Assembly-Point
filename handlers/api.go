package handlers

import (
	"blindsig/internal"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Types
type sPublicKey struct {
	ServerPublicKey string `json:"serverpublickey" xml:"serverpublickey"`
}

type errorJson struct {
	Error string `json:"error" xml:"error"`
}

// JSON Errors
var jsonError501 = errorJson{
	Error: "Not implemented",
}

var jsonError500 = errorJson{
	Error: "Internal Server Error",
}

// Super User HTMX Handlers

// JSON Handlers
func apiJsonPubkey(c echo.Context) error {

	// Server: generate an RSA keypair.
	sk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {

		err2 := &errorJson{
			Error: fmt.Sprintf("failed to generate RSA key: %v", err),
		}
		fmt.Printf("failed to generate RSA key: %v", err)
		return c.JSON(http.StatusInternalServerError, err2)
	}

	publicKey, _ := x509.MarshalPKIXPublicKey(&sk.PublicKey)

	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKey,
	}

	encodedPublicKey := string(pem.EncodeToMemory(block))

	jsonPublicKey := sPublicKey{ServerPublicKey: encodedPublicKey}

	return c.JSON(http.StatusOK, jsonPublicKey)
}

func apiJsonReqBlindSignature(c echo.Context) error {

	reqJson := internal.GetJSONRawBody(c)
	base64Str := reqJson["blindedMsg"].(string)
	_, err := internal.Base64StringToBytes(base64Str)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, jsonError500)
	}

	//fmt.Println(blindedMsgArr)

	return c.JSON(http.StatusNotImplemented, jsonError501)
}
