package healthcheck

import (
	"strings"
    "os"
    "encoding/csv"
    "encoding/json"
    "io/ioutil"
)

func loadConfig(fileName string) (*Config,error) {
    jsonFile, err := os.Open(fileName)
    if err != nil {
        return nil,err
    }
    byteValue, _ := ioutil.ReadAll(jsonFile)
    var config Config
    json.Unmarshal(byteValue, &config)
    return &config, nil
}

// Parsing URLs with CSV format
func parseURLs(filePath string) ([]string, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    // Create a new reader.
    csvr := csv.NewReader(f)
    csvr.FieldsPerRecord = -1

    lines,err := csvr.ReadAll()
    if err != nil {
        return nil, err
    }
    urls := []string{}
    for _, line := range lines {
        for _, column := range line {
            url := strings.TrimSpace(column)
            if url == "" {
                continue
            }
            urls = append(urls, url)
        }
    }
    return urls,nil
}