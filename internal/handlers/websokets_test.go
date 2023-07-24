package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func indexRequest(b *testing.B) {
	url := "/"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)
	// fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		b.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func adminRequest(b *testing.B) {
	url := "/admin"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserAll)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		b.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusOK)
	}
}
func BenchmarkDebugging(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// adminRequest(b)
		// indexRequest(b)
		// benchmarkAll(b)
		requestWithToken(b)
	}
}

func requestWithToken(b *testing.B) {
	url := "/chat?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM1MjM5ODAsInVzZXJuYW1lIjoibG9naW4xIn0.n4eEaQB3rPienS8d5NzDHke2PxAUWhEdp8cOpOYm9xA"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		b.Fatal(err)
	}

	// req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Connection", "Debugging")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RequestWithToken)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		b.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func benchmarkAll(b *testing.B) {
	url := "/"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)
	// fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		b.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		b.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)
	// fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		b.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusOK)
	}
	// fmt.Println("/admin")
	url = "/admin"
	req, err = http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		b.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetUserAll)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		b.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		b.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetUserAll)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		b.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	url = "/user/login"
	payload := `
			{
				"username":"login1",
				"password":"password1"
			}
		`
	req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		b.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(UserLogin)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		b.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	fmt.Println("TEST----")

	url = "/chat?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM1MjM5ODAsInVzZXJuYW1lIjoibG9naW4xIn0.n4eEaQB3rPienS8d5NzDHke2PxAUWhEdp8cOpOYm9xA"

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		b.Fatal(err)
	}

	// req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Connection", "Debugging")

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(RequestWithToken)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		b.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TTestDebugging(t *testing.T) {
	url := "/"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)
	// fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)
	// fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	url = "/admin"
	req, err = http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetUserAll)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetUserAll)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	url = "/user/login"
	payload := `
			{
				"username":"login1",
				"password":"password1"
			}
		`
	req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(UserLogin)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	fmt.Println("TEST----")

	url = "/chat?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM1MjM5ODAsInVzZXJuYW1lIjoibG9naW4xIn0.n4eEaQB3rPienS8d5NzDHke2PxAUWhEdp8cOpOYm9xA"
	// req, err = http.NewRequest(http.MethodGet, url, nil)
	// reqbody := "HTTP/1.1 1 1 map[Connection:[Upgrade] Sec-Websocket-Extensions:[permessage-deflate; client_max_window_bits] Sec-Websocket-Key:[4w09D5xKQEyVuSu/dcB59w==] Sec-Websocket-Version:[13] Upgrade:[websocket]] {} <nil> 0 [] false localhost:8080 map[] map[] <nil> map[] 127.0.0.1:63002 /chat?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMyODg4MTEsInVzZXJuYW1lIjoibG9naW4xIn0.CusR1TJ1rLSgEjJnpJFpHPjE7hZurKTK0ZpaDDNMavs <nil> <nil> <nil> 0xc0000602d0"
	// req, err = http.NewRequest(http.MethodGet, url, strings.NewReader(reqbody))

	// jsonBody := []byte(`{"Connection": "Upgrade"}`)
	// bodyReader := bytes.NewReader(jsonBody)
	// req, err = http.NewRequest(http.MethodGet, url, bodyReader)

	// fmt.Println(bodyReader)

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Connection", "Debugging")

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(RequestWithToken)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

}
