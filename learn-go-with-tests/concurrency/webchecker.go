package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	ch := make(chan result)

	for _, url := range urls {
		//pass to the anonymous func a copy of the url variable
		//otherwise all goroutines will process only the last url
		go func(url string) {
			ch <- result{string: url, bool: wc(url)}
		}(url)
	}

	results := make(map[string]bool)
	for r := range ch {
		results[r.string] = r.bool
	}
	return results
}
