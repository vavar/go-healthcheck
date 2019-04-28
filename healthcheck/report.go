package healthcheck

import (
    "net/http"
    "encoding/json"
    "bytes"
    "fmt"
)

const (
    reportURL = "https://hiring-challenge.appspot.com/healthcheck/report"
)

//Report via Healcheck Report API
func Report(apiToken string, result HealthCheckResult) {
    b, _ := json.Marshal(result)
    req, _ := http.NewRequest("POST", reportURL, bytes.NewBuffer(b))
    req.Header.Set("Authorization", "Bearer " + apiToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}

    if resp, err := client.Do(req); err == nil {
        if resp.StatusCode != http.StatusOK {
            fmt.Printf("Sending report failed")
        }
    }
}