package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main5() {

	for _, url := range os.Args[1:] {
		bool := strings.HasPrefix(url, "http://")
		if bool == false {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stderr, resp.Body)
		fmt.Println(resp.Status)
		resp.Body.Close()
	}

}
