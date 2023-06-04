package examples

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Example3() {
	start := time.Now()
	websites := []string{
		"https://stackoverflow.com/",
		"https://github.com/",
		"https://www.linkedin.com/",
		"http://medium.com/",
		"https://golang.org/",
		"https://www.udemy.com/",
		"https://www.coursera.org/",
		"https://wesionary.team/",
	}

	for _, website := range websites {
		go getWebsiteByWaitGroup(website)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}
func getWebsiteByWaitGroup(website string) {
	defer wg.Done()
	if res, err := http.Get(website); err != nil {
		fmt.Println(website, "is down")

	} else {
		fmt.Printf("[%d] %s is up\n", res.StatusCode, website)
	}

}
