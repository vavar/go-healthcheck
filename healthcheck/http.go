package healthcheck
import (
	"net/http"
	"sync"
	"time"
)

func fetch(url string, wg *sync.WaitGroup, ch chan<- HTTPResult) {
    start := time.Now()
    resp, err := http.Get(url)
    defer wg.Done()
    if err != nil {
        ch <- HTTPResult{ false, time.Since(start).Nanoseconds() }
        return
    }
    success := resp.StatusCode >= 200 && resp.StatusCode < 300
    ch <- HTTPResult{ success, time.Since(start).Nanoseconds() }
}

func asyncFetch(urls []string) (chan HTTPResult) {
    var wg sync.WaitGroup

    fetchSize := len(urls)
    queue := make(chan HTTPResult, fetchSize)
    wg.Add(fetchSize)
    for _, url := range urls {
        go fetch(url, &wg, queue)
    }

    wg.Wait()
    close(queue)
    return queue
}

// Return summary healthcheck result for providing URLs 
func FetchURLs(urls []string) HealthCheckResult {
    queue := asyncFetch(urls)
    result := HealthCheckResult{ 0, 0, 0, 0}
    for t := range queue {
        if t.Success {
            result.Success = result.Success + 1
        } else {
            result.Failure = result.Failure + 1
        }
        result.TotalWebsites = result.TotalWebsites + 1
        result.TotalTime = result.TotalTime + t.Time
    }
    return result
}