package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RssFeedXml struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel *RssFeed
}
type RssFeed struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`
	Link           string   `xml:"link"`
	Description    string   `xml:"descripton"`
	Language       string   `xml:"language,omitempty"`
	CopyRight      string   `xml:"copyright,omitempty"`
	ManagingEditor string   `xml:"managingEditor,omitempty"`
	WebMaster      string   `xml:"webMaster,omitempty"`
	PubDate        string   `xml:"pubDate,omitempty"`
	LastbuildDate  string   `xml:"lastbuildDate,omitempty"`
	Category       string   `xml:"category,omitempty"`
	Generator      string   `xml:"generator,omitempty"`
	Docs           string   `xml:"docs,omitempty"`
	Cloud          string   `xml:"cloud,omitempty"`
	Ttl            string   `xml:"ttl,omitempty"`
	Rating         string   `xml:"rating,omitempty"`
	SkipHours      string   `xml:"skipHours,omitempty"`
	skipDays       string   `xml:"skipDays,omitempty"`
	Image          *RssImage
	TextInput      *RssTextInput
	Items          []*RssItem
}
type RssImage struct {
	XMLName xml.Name `xml:"image"`
	Url     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml"link"`
	Width   int      `xml"width,omitempty"`
	Height  int      `xml:"height,omitempty"`
}
type RssTextInput struct {
	XMLName     xml.Name `xml:"textInput"`
	Title       string   `xml:"title"`
	Description string   `xml:"name"`
	Name        string   `xml:"name"`
	Link        string   `xml:"link"`
}
type RssItem struct {
	XMLName   xml.Name `xml:"item"`
	Title     string   `xml:"title"`
	Link      string   `xml:"link"`
	Author    string   `xml:"author,omitempty"`
	Category  string   `xml:"category,omitempty"`
	Comments  string   `xml:"comments,omitempty"`
	Enclosure *RssEnclosure
	Guid      string `xml:"guid,omitempty"`
	PubDate   string `xml:"pubDate,omitempty"`
	Source    string `xml:"source,omitempty"`
}
type RssEnclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	Url     string   `xml:"url,attr"`
	Length  string   `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

func Decode(data []byte) (*RssFeedXml, error) {
	v := RssFeedXml{}
	err := xml.Unmarshal(data, &v)
	return &v, err
}
func main() {
	resp, err := http.Get("http://www.geeksforgeeks.org/feed/")
	if err != nil {
		fmt.Println("Get Rss error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http body read error")
	}
	r, err := Decode(body)
	if err != nil {
		fmt.Println("decode error,", err)
	}
	str, _ := json.Marshal(r)
	fmt.Println(string(str))
}
