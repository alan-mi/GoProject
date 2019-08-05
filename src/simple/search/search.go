package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds() //解析json文件中的字段
	if err != nil {
		log.Fatal(err)
	}
	results := make(chan *Result)
	var waitGrop sync.WaitGroup
	waitGrop.Add(len(feeds))
	for _, feed := range feeds {
		matcher, exusts := matchers[feed.Type]
		if !exusts {
			matcher = matchers["default"]
		}
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGrop.Done()
		}(matcher, feed)
	}
	go func() {
		waitGrop.Wait()
		close(results)
	}()
	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "matcher already registerd")
	}
	log.Println("register", feedType, "matcher")
	matchers[feedType] = matcher
}
