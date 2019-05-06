package healthcheck

import (
    "net/http"
    "encoding/json"
    "bytes"
    "fmt"
)

//Report via Healcheck Report API
func (h *Healthcheck) Report(report Report) (error) {
    if (h.config.ReportURL == "" || h.config.Token == "") {
        return fmt.Errorf("missing config for ReportAPI")
    }

    b, _ := json.Marshal(report)
    req, _ := http.NewRequest("POST", h.config.ReportURL, bytes.NewBuffer(b))
    req.Header.Set("Authorization", "Bearer " + h.config.Token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("Sending report ... failed(%d)\n", resp.StatusCode)
    }

    return nil
}