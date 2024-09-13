package main

import "github.com/f7ed0/golog/lg"

func main() {
	lg.Init(lg.ALL, true)
	lg.Info.Println("Starting authentication microservice")
}
