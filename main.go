package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/alecthomas/template"
)

type songs struct {
	CurrentSong int    `json:"currentSong"`
	Songs       []song `json:"songs"`
}

type song struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Sounds []sound `json:"sounds"`
}

type sound struct {
	Note  note `json:"note"`
	Delay int  `json:"delay"`
}

type note struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"`
}

func main() {
	http.HandleFunc("/songs", songsHandler)
	http.ListenAndServe(":8080", nil)

}
func songsHandler(rw http.ResponseWriter, r *http.Request) {
	url := "http://localhost:7878/songs"

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	songs1 := songs{}
	jsonErr := json.Unmarshal(body, &songs1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(songs1.Songs)

	t, err := template.ParseFiles("songs.html")
	if err != nil {
		panic(err)
	}
	t.Execute(rw, songs1)

}
