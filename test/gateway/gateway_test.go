package gateway_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"testing"
)

func TestForwarding(t *testing.T) {
	requestData := struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}{
		Identifier: "identifier",
		Password:   "password",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := http.Post("http://localhost:8080/Authentication/Login/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Error(err)
		return
	}

	b, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(b))

	// We are not running authentication service here so for testing it will return a 502 not a 501
	if resp.StatusCode != http.StatusBadGateway {
		t.Fail()
		return
	}
}

func TestError(t *testing.T) {
	requestData := struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}{
		Identifier: "identifier",
		Password:   "password",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := http.Post("http://localhost:8080/BadService/Login/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Error(err)
		return
	}

	b, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(b))

	if resp.StatusCode != http.StatusBadRequest {
		t.Fail()
		return
	}
}
