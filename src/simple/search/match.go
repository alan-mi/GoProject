package search

import "fmt"

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//实例化一个接口并把实现接口的结构体赋值给他 然后接口就可以工作
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, result := range searchResults {
		results <- result
	}

}

func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)

	}
}
