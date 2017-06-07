package main

import "net/http"

const (
	port = ":19022"
)

const (
	rpFeeds = "/feeds"
)

func main() {
	http.HandleFunc(rpFeeds, handleFeedReq)
}

func handleFeedReq(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodPut:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodDelete:
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
	return
}
