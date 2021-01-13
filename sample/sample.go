// +build run

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spiegel-im-spiegel/fetch"
)

func main() {
	u, err := fetch.URL("https://github.com/spiegel-im-spiegel.gpg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	resp, err := fetch.New(fetch.WithHTTPClient(&http.Client{})).
		Get(u, fetch.WithContext(context.Background()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Close()
	if _, err := io.Copy(os.Stdout, resp.Body()); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
