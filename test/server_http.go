package test

import "net/http"

func SeverHttp() {
	//Listen server port 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	//server running
	http.ListenAndServe(":8080", nil)
}
