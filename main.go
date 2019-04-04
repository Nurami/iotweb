package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var ip string = "169.254.121.188"

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
	tmp := os.Args[1:]
	ip = tmp[0]

	http.HandleFunc("/songs", songsHandler)
	http.ListenAndServe(":8080", nil)

}
func songsHandler(rw http.ResponseWriter, r *http.Request) {
	url := "http://" + ip + ":7878/songs"

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
	t, err := template.ParseFiles("songs.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(rw, songs1)

}
