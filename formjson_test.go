package formjson

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/rs/xhandler"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestPOST(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, url.Values{"name": []string{"foo"}}, r.Form)
		assert.Equal(t, url.Values{"name": []string{"foo"}}, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestPUT(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, url.Values{"name": []string{"foo"}}, r.Form)
		assert.Equal(t, url.Values{"name": []string{"foo"}}, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestPATCH(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, url.Values{"name": []string{"foo"}}, r.Form)
		assert.Equal(t, url.Values{"name": []string{"foo"}}, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestGET(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Nil(t, r.Form)
		assert.Nil(t, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestXhandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := xhandler.HandlerFuncC(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, url.Values{"name": []string{"foo"}}, r.Form)
		assert.Equal(t, url.Values{"name": []string{"foo"}}, r.PostForm)
	})

	HandlerC(h).ServeHTTPC(context.Background(), res, req)
}

func TestStringJSONValue(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, url.Values{"name": []string{"foo"}}, r.Form)
		assert.Equal(t, url.Values{"name": []string{"foo"}}, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestStringArrayJSONValue(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":[\"foo\",\"bar\"]}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, url.Values{"name": []string{"foo", "bar"}}, r.Form)
		assert.Equal(t, url.Values{"name": []string{"foo", "bar"}}, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestMixArrayJSONValue(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":[\"foo\",1,true,false]}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, url.Values{"name": []string{"foo", "1", "1", "0"}}, r.Form)
		assert.Equal(t, url.Values{"name": []string{"foo", "1", "1", "0"}}, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestNumberJSONValue(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":1.2}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, url.Values{"name": []string{"1.2"}}, r.Form)
		assert.Equal(t, url.Values{"name": []string{"1.2"}}, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestBoolJSONValue(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"a\":true,\"b\":false}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, url.Values{"a": []string{"1"}, "b": []string{"0"}}, r.Form)
		assert.Equal(t, url.Values{"a": []string{"1"}, "b": []string{"0"}}, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestUnsupportedJSONValue(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":{}}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Nil(t, r.Form)
		assert.Equal(t, url.Values{}, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestUnsuportedArrayJSONValue(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{\"name\":[\"foo\",{}]}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Nil(t, r.Form)
		assert.Equal(t, url.Values{}, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestMixQueryString(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/?foo=bar", bytes.NewBufferString("{\"name\":\"baz\"}"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, url.Values{"name": []string{"baz"}, "foo": []string{"bar"}}, r.Form)
		assert.Equal(t, url.Values{"name": []string{"baz"}}, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestInvalidJSON(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString("{"))
	req.Header.Add("Content-Type", "application/json")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Nil(t, r.Form)
		assert.Nil(t, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}

func TestInvalidContentType(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/", bytes.NewBufferString("{\"name\":\"foo\"}"))
	req.Header.Add("Content-Type", "application/bson")

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Nil(t, r.Form)
		assert.Nil(t, r.PostForm)
	})

	Handler(h).ServeHTTP(res, req)
}
