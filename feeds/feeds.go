package feeds

import (
	"log"
	"sort"
	"time"
	"ziglu_tech_test/cache"
	"ziglu_tech_test/data"
)

// here we need a list of the current feeds
// the ability to create a new cache to get feeds
// stored in a map of caches
// need feed categories as well

type newsfeed struct {
	Provider string
	Category string
	data     *cache.Cache
}

var newsfeeds map[string]newsfeed

func GetAllFeeds() []*data.Item {
	var allItems = []*data.Item{}

	for _, feeds := range newsfeeds {
		allItems = append(allItems, feeds.data.GetFeed()...)
	}

	return sortItems(allItems)
}

func GetNewsFeedByName(name string) []*data.Item {
	if _, ok := newsfeeds[name]; ok {
		return sortItems(newsfeeds[name].data.GetFeed())
	}
	return nil
}
func GetNewsFeedByProviderAndCategory(provider string, category string) []*data.Item {

	var allItems = []*data.Item{}

	// find appropriate feeds

	for _, feeds := range newsfeeds {
		if feeds.Provider == provider && feeds.Category == category {
			allItems = append(allItems, feeds.data.GetFeed()...)
		}
	}

	return sortItems(allItems)
}

func GetNewsFeedByProvider(provider string) []*data.Item {

	var allItems = []*data.Item{}

	// find appropriate feeds

	for _, feeds := range newsfeeds {
		if feeds.Provider == provider {
			allItems = append(allItems, feeds.data.GetFeed()...)
		}
	}

	return sortItems(allItems)
}

func GetNewsFeedByCategory(category string) []*data.Item {

	var allItems = []*data.Item{}

	// find appropriate feeds

	for _, feeds := range newsfeeds {
		if feeds.Category == category {
			allItems = append(allItems, feeds.data.GetFeed()...)
		}
	}

	return sortItems(allItems)
}

func AddNewsFeed(name string, provider string, category string, url string) {

	// check fields
	if len(name) < 1 || len(provider) < 1 || len(category) < 1 || len(url) < 1 {
		return
	}

	newsFeed := newsfeed{
		Provider: provider,
		Category: category,
		data: &cache.Cache{
			Ttl: 15,
			Feeds: data.Feed{
				Url: url,
			},
		},
	}
	if newsfeeds == nil {
		// first run init
		newsfeeds = map[string]newsfeed{}
	}
	newsfeeds[name] = newsFeed
}

func sortItems(items []*data.Item) []*data.Item {

	sort.Slice(items, func(i, j int) bool {
		iTime, err := time.Parse(time.RFC1123, items[i].PubDate)
		if err != nil {
			log.Fatalln(err)
		}
		jTime, err := time.Parse(time.RFC1123, items[j].PubDate)
		if err != nil {
			log.Fatalln(err)
		}
		return iTime.After(jTime)
	})

	return items
}
