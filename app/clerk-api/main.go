package main

import (
	"github.com/csixteen/clerk/app/clerk-api/broker"
)

func main() {
	b := broker.New()
	b.Start()
}
