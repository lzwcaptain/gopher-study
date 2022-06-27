package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func fillUrl(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return url
}

func main() {
	for _, url := range os.Args[1:] {
		url := fillUrl(url)
		resp, err := http.Get(url)
		fmt.Println(resp.StatusCode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		// ex 1.7
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}
}
