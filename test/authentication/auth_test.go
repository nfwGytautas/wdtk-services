package auth_test

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"testing"
)

func TestLogin(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/Login/", "application/json", nil)
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
