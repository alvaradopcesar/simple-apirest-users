package routes

import "net/http"

// RouterRoute func
func RouterRoute() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/delete/", DeleteRoute)
	mux.HandleFunc("/update/", UpdateRoute)
	mux.HandleFunc("/create", CreateRoute)
	mux.HandleFunc("/", ReadRoute)
	return mux
}
