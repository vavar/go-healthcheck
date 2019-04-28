package healthcheck

// Summary Healthcheck for overall websites
type HealthCheckResult struct {
    TotalWebsites int `json:"total_websites"`
    Success int `json:"success"`
    Failure int `json:"failure"`
    TotalTime int64 `json:"total_time"`
}

// Healthcheck result
type HTTPResult struct {
    Success bool
    Time int64
}