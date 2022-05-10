package main

import (
	"finance-management-web/routes"
	"net/http"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8000", nil)
}
