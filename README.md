# Go JSON handler [![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/rs/formjson) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/rs/formjson/master/LICENSE) [![build](https://img.shields.io/travis/rs/formjson.svg?style=flat)](https://travis-ci.org/rs/formjson)

FormJSON is a `net/http` handler implementing content negotiation for posted data in order to transparently expose posted JSON as if it was `application/x-www-form-urlencoded`. The posted data is then available using built-in `r.FormValue("key")`'s `http.Request` method.

To match capabilities of `application/x-www-form-urlencoded`, only single depth JSON object with `string` as keys and values is supported.

## Usage

```go
package main

import (
    "net/http"
    "fmt"

    "github.com/rs/formjson"
)

func main() {
    h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        body := fmt.Sprintf("Hello %s!", r.FormValue("name"))
        w.Write([]byte(body))
    })

    handler := formjson.Handler(h)
    http.ListenAndServe(":8080", handler)
}
```

Then this request:

    curl -H "Content-Type:application/json" -d '{"name":"World"}' :8080

is now equivalent to this one:

    curl -d name=World :8080

## Licenses

All source code is licensed under the [MIT License](https://raw.github.com/rs/formjson/master/LICENSE).
