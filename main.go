package main

import (
	"pagerduty/src"
)

func main() {
	src.InitDatabase()
	src.StartServer()
}
