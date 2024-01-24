package main

import (
	"encoding/xml"

	"net/http"
	"sync"

	"github.com/google/uuid"
)
type returnVal struct {
	val * Channel
	id   uuid.UUID
}
//Channel struct for RSS
type Channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Language      string `xml:"language"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Item          []Item `xml:"item"`
}

//ItemEnclosure struct for each Item Enclosure
type ItemEnclosure struct {
	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

//Item struct for each Item in the Channel
type Item struct {
	Title       string          `xml:"title"`
	Link        string          `xml:"link"`
	Comments    string          `xml:"comments"`
	PubDate     string            `xml:"pubDate"`
	GUID        string          `xml:"guid"`
	Category    []string        `xml:"category"`
	Enclosure   []ItemEnclosure `xml:"enclosure"`
	Description string          `xml:"description"`
	Author      string          `xml:"author"`
	Content     string          `xml:"content"`
	FullText    string          `xml:"full-text"`
}

func (cfg *apiConfig)GetRssData(url string, id uuid.UUID, c chan * returnVal, wg *sync.WaitGroup){
	defer wg.Done()
    resp, err:= http.Get(url)
	if err != nil {
		c <- nil
		return
	}
	data :=  struct {
		Channel Channel `xml:"channel"`
	}{}
	decoder := xml.NewDecoder(resp.Body)
	err2 := decoder.Decode(&data)
	if err2 != nil {
		c <- nil
		return
	}
	c <- &returnVal{
		val: &data.Channel,
		id: id,
	}
}