package formjson

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
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
	var d map[string]interface{}
	if err := decoder.Decode(&d); err != nil {
		return
	}
	// Inject parsed data into PostForm
	r.PostForm = url.Values{}
	for k, v := range d {
		switch t := v.(type) {
		case string:
			r.PostForm.Set(k, t)
		case float64:
			r.PostForm.Set(k, strconv.FormatFloat(t, 'f', -1, 64))
		case bool:
			if t {
				r.PostForm.Set(k, "1")
			} else {
				r.PostForm.Set(k, "0")
			}
		case []interface{}:
			r.PostForm[k] = []string{}
			for _, sv := range t {
				switch st := sv.(type) {
				case string:
					r.PostForm.Add(k, st)
				case float64:
					r.PostForm.Add(k, strconv.FormatFloat(st, 'f', -1, 64))
				case bool:
					if st {
						r.PostForm.Add(k, "1")
					} else {
						r.PostForm.Add(k, "0")
					}
				default:
					// Do not translate array partially
					r.PostForm.Del(k)
					break
				}
			}
		}
	}
	// Build the Form property
	if len(r.PostForm) > 0 {
		r.Form = url.Values{}
		for k, vs := range r.PostForm {
			r.Form[k] = []string{}
			for _, v := range vs {
				r.Form.Add(k, v)
			}
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
