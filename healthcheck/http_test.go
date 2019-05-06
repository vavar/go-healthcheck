package healthcheck

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "gotest.tools/assert"
)

func mockEndpoint(status int) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(status)
    }))
}

func getURLs(endpoints []*httptest.Server) []string {
    result := []string{}
    for _, endpoint := range endpoints {
        result = append(result, endpoint.URL)
    }
    return result
}

func TestHttpHealthcheck(t *testing.T) {
    testCases := map[string]struct {
        endpoints []*httptest.Server
        report *Report
    }{
        "single StatusOK site":{
            []*httptest.Server{ mockEndpoint(http.StatusOK) }, 
            &Report{1,1,0,0},
        },
        "single StatusNotFound site":{
            []*httptest.Server{ mockEndpoint(http.StatusNotFound) },
            &Report{1,0,1,0},
        },
        "multiple StatusOK sites":{
            []*httptest.Server{ mockEndpoint(http.StatusOK), mockEndpoint(http.StatusOK) }, 
            &Report{2,2,0,0},
        },
        "multiple StatusNotFound sites":{
            []*httptest.Server{ mockEndpoint(http.StatusNotFound), mockEndpoint(http.StatusNotFound) }, 
            &Report{2,0,2,0},
        },
        "multiple mix sites":{
            []*httptest.Server{ mockEndpoint(http.StatusOK), mockEndpoint(http.StatusNotFound) }, 
            &Report{2,1,1,0},
        },
    }

    for name, it := range testCases {
        t.Run(name, func(t *testing.T) {
            defer func() { for _, e := range it.endpoints { e.Close() } }()
            urls := getURLs(it.endpoints)
            report := (&Healthcheck{ urls: urls }).Check()
            assert.Equal(t, report.TotalWebsites, it.report.TotalWebsites)
            assert.Equal(t, report.Success, it.report.Success)
            assert.Equal(t, report.Failure, it.report.Failure)
        })
    }
}