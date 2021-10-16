package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/MathewKostiuk/sorting"
)

var tracks = []*sorting.Track{
	{Title: "Go", Artist: "Delilah", Album: "From the Roots up", Year: 2012, Length: length("3m38s")},
	{Title: "Go", Artist: "Moby", Album: "Moby", Year: 1992, Length: length("3m37s")},
	{Title: "Go Ahead", Artist: "Alicia Keys", Album: "As I Am", Year: 2007, Length: length("4m36s")},
	{Title: "Ready 2 Go", Artist: "Martin Solveig", Album: "Smash", Year: 2011, Length: length("4m24s")},
}

var tt sorting.TrackTable

type Payload struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/sort", sortHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	tt.Tracks = tracks
	sorting.WriteTracks(tt, w, tmpl)
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var p Payload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	sorting.UpdateSortKeys(strings.Title(p.Name))
	sorting.SortTracks(tt.Tracks)
	sorting.WriteTracks(tt, w, tmpl)
}
