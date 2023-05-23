package http_requests

import (
    "log"
	"net/http"
	"io"
)

func CallGET(APIURL string){
    resp, err := http.Get(APIURL)
    if err!= nil {
            log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err!= nil {
            log.Fatal(err)
    }
    log.Println(string(body))
}

