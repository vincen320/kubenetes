package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Job executed at " + time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST"))
}

//docker build -t vincen320/go-job:tag .

//docker build -t (hubuser/namaimages:tag) (lokasiFolder Dockerfile) >> pakai titik karena kebetulan di "cd" sampe kesini
//terus untuk dipakai di kubernetes jangan lupa di push ke dockerhub terlebih dahulu
