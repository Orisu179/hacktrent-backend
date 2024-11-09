package main

import "net/http"

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Server", "Go")
	_, err := w.Write([]byte("Hello from rare animal!"))
	if err != nil {
		return
	}
}
