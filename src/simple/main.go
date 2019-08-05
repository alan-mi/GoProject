package main

import (
	"log"
	"os"
	_ "simple/matchers"
	"simple/search"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("president")
}
