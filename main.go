package main

import "log"

func main() {
	_, err := scrape()
	if err != nil {
		log.Fatal("An error has occured.")
	}
}
