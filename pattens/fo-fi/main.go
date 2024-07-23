package main

import (
	"net/http"
	"sync"
)

func checkStatus(urls []string) ([]string, []error) {
	// define channels
	var wg sync.WaitGroup
	results := make(chan string)
	errors := make(chan error)
	// step 2 fan out

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				errors <- err
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				errors <- err
			}
			results <- resp.Status
		}(url)

	}
	// fan in

	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()
	var status []string
	var errs []error
	// sync and collect result
	for {
		select {
		case resp, ok := <-results:
			if ok {
				status = append(status, resp)
			} else {
				results = nil
			}
		case err, ok := <-errors:
			if ok {
				errs = append(errs, err)
			} else {
				errs = nil
			}
		}
		if results == nil && errs == nil {
			break
		}

	}
	return status, errs
}

func main() {
	urls := []string{"http://www.google.com", "http://www.facebook.com", "http://www.yahoo.com"}
	status, errs := checkStatus(urls)
	for _, s := range status {
		println(s)
	}
	for _, e := range errs {
		println(e)
	}
}
