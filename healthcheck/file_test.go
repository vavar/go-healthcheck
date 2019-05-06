package healthcheck

import (
    "testing"
    "gotest.tools/assert"
)

func TestLoadConfig(t *testing.T) {
    testCases := map[string]struct {
        filePath string
        config *Config
        err string
    }{
        "config load failed": {"a.json", nil, "open a.json: no such file or directory"},
        "config load success": {"fixtures/config_test.json", &Config{"aaaa", "https://example.com"}, ""},
    }

    for name, it := range testCases {
		t.Run(name, func(t *testing.T) {
            config, err := loadConfig(it.filePath)
            if it.err != "" {
                assert.Error(t, err, it.err)
            }

            if it.config != nil {
                assert.DeepEqual(t, config, it.config)
            }
        })
    }
}

func TestParseURLs(t *testing.T) {
    testCases := map[string]struct {
        filePath string
        urls []string
        err string
    }{
        "file load failed": {"a.json", nil, "open a.json: no such file or directory"},
        "file load success": {"fixtures/test.csv", []string{"http://example.com", "http://unittest.com"}, ""},
    }

    for name, it := range testCases {
		t.Run(name, func(t *testing.T) {
            urls, err := parseURLs(it.filePath)
            if it.err != "" {
                assert.Error(t, err, it.err)
            }

            if it.urls != nil {
                assert.DeepEqual(t, urls, it.urls)
            }
        })
    }
}