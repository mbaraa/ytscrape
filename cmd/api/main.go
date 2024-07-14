package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mbaraa/ytscrape"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Refer to https://github.com/mbaraa/ytscrape for more info!"))
}

func handleSearchYt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Header", "*")
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query().Get("q")
	if q == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Missing `q` in the query list!"))
		return
	}
	results, err := ytscrape.Search(q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		_, _ = w.Write([]byte("Something went wrong..."))
		return
	}

	_ = json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/search", handleSearchYt)
	log.Println("Starting ytscrape server at port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
