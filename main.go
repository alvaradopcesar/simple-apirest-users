package main

import (
	"fmt"
	"net/http"

	"github.com/wilian1808/simple-apirest-users/routes"
)

func main() {
	mux := routes.AllRoute()
	fmt.Println("Serving in port http://localhost:6060")
	_ = http.ListenAndServe(":6060", mux)
}
