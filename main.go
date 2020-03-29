package main

import "log"

func main() {
	err, _ := scrape()
	if err != nil {
		log.Fatal("An error has occured.")
	}

}
