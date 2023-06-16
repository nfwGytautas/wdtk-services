package auth_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"testing"

	"github.com/nfwGytautas/gdev/jwt"
)

func TestLogin(t *testing.T) {
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

	resp, err := http.Post("http://localhost:8080/Login/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Error(err)
		return
	}

	b, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(b))

	if resp.StatusCode != http.StatusOK {
		t.Fail()
		return
	}
}

func TestRegister(t *testing.T) {
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

	resp, err := http.Post("http://localhost:8080/Register/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Error(err)
		return
	}

	b, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(b))

	if resp.StatusCode != http.StatusNoContent {
		t.Fail()
		return
	}
}

func TestMe(t *testing.T) {
	jwt.APISecret = "TEST_KEY"

	request, err := http.NewRequest("GET", "http://localhost:8080/Me/", nil)
	if err != nil {
		t.Error(err)
		return
	}

	tokenString, err := jwt.GenerateToken(123, "Role")
	if err != nil {
		t.Error(err)
		return
	}

	request.Header.Add("Authorization", "bearer "+tokenString)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
		return
	}

	b, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(b))

	if resp.StatusCode != http.StatusOK {
		t.Fail()
		return
	}
}
