package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":3000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Application 3.0")
}

//docker build -t vincen320/go-web images/go-web
//docker build -t vincen320/go-web:tag images/go-web
//docker build -t vincen320/go-web:latest -t vincen320/go-web:3 images/go-web

//docker build -t (hubuser/namaimages:tag) (lokasiFolder Dockerfile)
//terus untuk dipakai di kubernetes jangan lupa di push ke dockerhub terlebih dahulu

//docker push vincen320/go-web [-a]
//-a : push all tags
