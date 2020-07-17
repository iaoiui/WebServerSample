package server

import "net/http"

// Serveはwebサーバを起動して待ち受けます
func Serve() {
	http.HandleFunc("/", rootAction)
	http.ListenAndServe(":8080", nil)
}

func rootAction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Webserver Sample\n"))
}
