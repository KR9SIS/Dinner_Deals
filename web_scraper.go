package main

import (
	"fmt"
	"io"
	"log"
	"net/html"
	"net/http"
)

func scrape_site(site string) {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}

	body, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// body, err := io.ReadAll(resp.Body)
	// resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with stats code: %d and \nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
}
