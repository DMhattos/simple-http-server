package server

import (
	"fmt"
	"net/http"
)

func Listen(port int) {
	home()
	about()
	serverPort := fmt.Sprintf(`:%v`, port)
	http.ListenAndServe(serverPort, nil)
}
func home() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home page")
	})
}
func about() {
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About page")
	})
}
