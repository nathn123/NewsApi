package data

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type Feed struct {
	Url  string
	Feed []*Item
}

type Channel struct {
}
type Rss struct {
	XMLName      xml.Name `xml:"rss"`
	Items        []*Item  `xml:"channel>item"`
	DefaultImage string   `xml:"channel>image>url"`
}

type Item struct {
	Title       string `xml:"title" json:"title"`
	Link        string `xml:"link" json:"link"`
	Description string `xml:"description" json:"description"`
	PubDate     string `xml:"pubDate" json:"pubDate"`
	Guid        string `xml:"guid" json:"guid"`
	Image       string `xml:"media:content" json:"media:content"`
}

func (f *Feed) RefreshFeed() {
	f.Feed = getFeed(f.Url)

}

func getFeed(url string) []*Item {

	var f = Rss{}
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = xml.Unmarshal(body, &f)
	if err != nil {
		log.Fatalln(err)
	}

	setDefaultImage(f.DefaultImage, f.Items)
	return f.Items
}

func setDefaultImage(imageUrl string, items []*Item) {
	for _, item := range items {
		if len(item.Image) < 1 {
			item.Image = imageUrl
		}
	}
}
