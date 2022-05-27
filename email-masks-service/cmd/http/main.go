package main

import "email-masks-service/src/infrastructure/bootstrap/http"

func main() {
	app := http.NewServer()
	app.Run()
}
