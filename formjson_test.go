package formjson

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/xhandler"
	"golang.org/x/net/context"
)

func TestPOST(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("name")))
	})

	Handler(h).ServeHTTP(res, req)

	if res.Body.String() != "foo" {
		t.Fail()
	}
}

func TestPUT(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("name")))
	})

	Handler(h).ServeHTTP(res, req)

	if res.Body.String() != "foo" {
		t.Fail()
	}
}

func TestPATCH(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("name")))
	})

	Handler(h).ServeHTTP(res, req)

	if res.Body.String() != "foo" {
		t.Fail()
	}
}

func TestGET(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("name")))
	})

	Handler(h).ServeHTTP(res, req)

	if res.Body.String() != "" {
		t.Fail()
	}
}

func TestXhandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := xhandler.HandlerFuncC(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("name")))
	})

	HandlerC(h).ServeHTTPC(context.Background(), res, req)

	if res.Body.String() != "foo" {
		t.Fail()
	}
}

func TestNonStringJSONValue(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":1}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("name")))
	})

	Handler(h).ServeHTTP(res, req)

	if res.Body.String() != "" {
		t.Fail()
	}
}

func TestMixQueryString(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/?foo=bar", bytes.NewBufferString("{\"name\":\"baz\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("foo")))
		w.Write([]byte(r.FormValue("name")))
	})

	Handler(h).ServeHTTP(res, req)

	if res.Body.String() != "barbaz" {
		t.Fail()
	}
}

func TestInvalidJSON(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("name")))
	})

	Handler(h).ServeHTTP(res, req)

	if res.Body.String() != "" {
		t.Fail()
	}
}

func TestInvalidContentType(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/bson")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("name")))
	})

	Handler(h).ServeHTTP(res, req)

	if res.Body.String() != "" {
		t.Fail()
	}
}
