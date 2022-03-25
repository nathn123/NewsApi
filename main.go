package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"ziglu_tech_test/data"
	"ziglu_tech_test/feeds"
)

type newNewsfeed struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
	Category string `json:"category"`
	Url      string `json:"url"`
}

func main() {
	// http.HandleFunc("/feed", getFeed)

	feeds.AddNewsFeed("sky_news_uk", "sky", "uk", "https://feeds.skynews.com/feeds/rss/uk.xml")
	feeds.AddNewsFeed("sky_news_tech", "sky", "tech", "https://feeds.skynews.com/feeds/rss/technology.xml")
	feeds.AddNewsFeed("bbc_news_uk", "bbc", "uk", "http://feeds.bbci.co.uk/news/uk/rss.xml")
	feeds.AddNewsFeed("bbc_news_tech", "bbc", "tech", "http://feeds.bbci.co.uk/news/technology/rss.xml")

	http.Handle("/all", http.HandlerFunc(getAllNewsFeeds))
	http.Handle("/name", http.HandlerFunc(getNewsFeedByName))
	http.Handle("/filter", http.HandlerFunc(getNewsFeedWithFilter))
	http.Handle("/add", http.HandlerFunc(addNewsFeed))

	http.ListenAndServe(":8080", nil)

}

func getNewsFeedByName(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query()["name"]) < 1 {
		// error
		http.Error(w, "no query string feed by name", 500)
		return
	}
	feed := feeds.GetNewsFeedByName(r.URL.Query()["name"][0])
	response, err := json.Marshal(feed)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getAllNewsFeeds(w http.ResponseWriter, r *http.Request) {
	// call to get newsfeed via cache

	feed := feeds.GetAllFeeds()

	response, err := json.Marshal(feed)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getNewsFeedWithFilter(w http.ResponseWriter, r *http.Request) {

	var feed []*data.Item
	provider := r.URL.Query()["provider"]
	category := r.URL.Query()["category"]

	switch {
	// both query strings
	case len(provider) > 0 && len(category) > 0:
		feed = feeds.GetNewsFeedByProviderAndCategory(provider[0], category[0])
	case len(provider) > 0:
		feed = feeds.GetNewsFeedByProvider(provider[0])
	case len(category) > 0:
		feed = feeds.GetNewsFeedByCategory(category[0])
	}

	response, err := json.Marshal(feed)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func addNewsFeed(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	var newfeed newNewsfeed
	err = json.Unmarshal(body, &newfeed)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	feeds.AddNewsFeed(newfeed.Name, newfeed.Provider, newfeed.Category, newfeed.Url)
	if err != nil {
		w.WriteHeader(400)
	}
}
