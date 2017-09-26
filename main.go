package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	JustHeaders bool `short:"I" long:"justheaders" description:"Return just the response headers"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		panic(err)
	}

	response, err := http.Get(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if opts.JustHeaders {
		for k, v := range response.Header {
			fmt.Println(k + ": " + strings.Join(v, ","))
		}
		os.Exit(0)
	}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
