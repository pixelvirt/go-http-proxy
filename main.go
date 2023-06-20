package main

import (
	"encoding/json"
	"fmt"
	"github.com/elazarl/goproxy"
	"io"
	"log"
	"net/http"
	"os"
        "bytes"
)

var port = os.Getenv("PORT")

func main() {
	if port == "" {
		port = "8080"
	}
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			r.Header.Set("X-GoProxy", "yxorPoG-X")
			decoder := json.NewDecoder(r.Body)
			var data root
			err := decoder.Decode(&data)
			if err != nil {
				fmt.Println(err)
				return r, nil
			}

			data.Country = "Nepal"
			// Encode the modified data
			encodedData, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
				return r, nil
			}

			// Update the request body with the modified data
			r.Body = io.NopCloser(bytes.NewReader(encodedData))
			r.ContentLength = int64(len(encodedData))

			return r, nil
		})

	fmt.Println("Startring server ... ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), proxy))
}

type root struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}
