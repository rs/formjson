package formjson

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/context"

	"github.com/rs/xhandler"
)

// Handler detects when a POST/PUT/PATCH content type is JSON, and transparently convert
// the JSON content into a standard PostForm. This does only support posting of a JSON dictionary
// containing string => string key value pairs.
func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleFormJSONRequest(r)
		h.ServeHTTP(w, r)
	})
}

// HandlerC detects when a POST/PUT/PATCH content type is JSON, and transparently convert
// the JSON content into a standard PostForm. This does only support posting of a JSON dictionary
// containing string => string key value pairs.
func HandlerC(h xhandler.HandlerC) xhandler.HandlerC {
	return xhandler.HandlerFuncC(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		handleFormJSONRequest(r)
		h.ServeHTTPC(ctx, w, r)
	})
}

func handleFormJSONRequest(r *http.Request) {
	if strings.Index(r.Header.Get("Content-Type"), "application/json") == -1 {
		return
	}
	switch r.Method {
	case "POST":
	case "PUT":
	case "PATCH":
		// whitelisted methods
	default:
		return
	}

	// Try to decode body using a restrictive type
	decoder := json.NewDecoder(r.Body)
	var d map[string]string
	if err := decoder.Decode(&d); err != nil {
		return
	}
	// Inject parsed data into PostForm
	r.PostForm = url.Values{}
	for k, v := range d {
		r.PostForm.Set(k, v)
	}
	// Build the Form property
	if len(r.PostForm) > 0 {
		r.Form = url.Values{}
		for k, v := range r.PostForm {
			r.Form.Set(k, v[0])
		}
		if r.URL != nil {
			if newValues, err := url.ParseQuery(r.URL.RawQuery); err == nil {
				for k, v := range newValues {
					r.Form.Set(k, v[0])
				}
			}
		}
	}
}
