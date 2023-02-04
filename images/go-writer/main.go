package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	path := os.Getenv("HTML_LOCATION")
	if path == "" {
		path = "html"
	}

	for {
		select {
		case <-time.After(5 * time.Second):
			//file, err := create("app/html/index.html") di dockerfile sudah dalam app
			file, err := create(path + "/index.html")
			fmt.Println(err)
			fmt.Fprintf(file, "<html><body>"+time.Now().Format("2006-01-02 15:04:05.999999999")+"</body></html>")
			file.Close()
		}
	}
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

//docker build -t vincen320/go-writer:tag .

//docker build -t (hubuser/namaimages:tag) (lokasiFolder Dockerfile) >> pakai titik karena kebetulan di "cd" sampe kesini
//terus untuk dipakai di kubernetes jangan lupa di push ke dockerhub terlebih dahulu
