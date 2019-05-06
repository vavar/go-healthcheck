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
    ch <- HTTPResult{ (resp.StatusCode == http.StatusOK), time.Since(start).Nanoseconds() }
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
func (h *Healthcheck) Check() Report {
    queue := asyncFetch(h.urls)
    report := Report{ 0, 0, 0, 0}
    for t := range queue {
        if t.Success {
            report.Success++
        } else {
            report.Failure++
        }
        report.TotalWebsites++
        report.TotalTime = report.TotalTime + t.Time
    }
    return report
}