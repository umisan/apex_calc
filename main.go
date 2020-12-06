package main

import "log"

func main() {
	log.Println("This is a Apex Reverse Calculator")
	server := CreateHttpServer()
	server.Start()
}
