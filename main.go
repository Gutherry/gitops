package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello Argo, estamos Online!!!</h1>"))
	})
	http.ListenAndServe(":8090", nil)
}
