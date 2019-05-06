package healthcheck

// Summary Healthcheck for overall websites
type Report struct {
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

type Config struct {
    Token string `json:"api_token"`
    ReportURL string `json:"report_url"`
}

type Healthcheck struct {
    config *Config
    urls []string
}

func NewChecker(filePath string) (*Healthcheck,error) {
    config, err := loadConfig("config.json")
    if (err != nil) {
        return nil, err
    }

    urls, err := parseURLs(filePath)
    if (err != nil) {
        return nil, err
    }

    return &Healthcheck{ config, urls }, nil
}