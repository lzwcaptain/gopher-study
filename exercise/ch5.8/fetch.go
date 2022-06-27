package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
)

func main() {
	fmt.Println(num(1))
}

func num(a int64) string {
	defer func() {
		a += 1
		fmt.Println("1")
	}()
	defer func() {
		a += 2
		fmt.Println("2")
	}()
	a += 3
	return strconv.FormatInt(a, 10)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local != "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	defer f.Close()
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
		os.Exit(1)
	}

	return local, n, err
}
