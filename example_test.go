package formjson

import (
	"fmt"
	"net/http"
)

func ExampleHello() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := fmt.Sprintf("Hello %s!", r.FormValue("name"))
		w.Write([]byte(body))
	})

	handler := Handler(h)
	http.ListenAndServe(":8080", handler)
}
