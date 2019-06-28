package routes

import "net/http"

// AllRoute func
func AllRoute() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/delete/", DeleteRoute)
	mux.HandleFunc("/update/", UpdateRoute)
	mux.HandleFunc("/create", CreateRoute)
	mux.HandleFunc("/", ReadRoute)
	return mux
}
