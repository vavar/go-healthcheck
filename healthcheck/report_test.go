package healthcheck

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "gotest.tools/assert"
)

func TestReport(t *testing.T) {
    testCases := map[string]struct {
        status int
        err string
    }{
        "send report api failed":{ http.StatusBadRequest, "Sending report ... failed(400)\n" },
        "send report api success":{ http.StatusOK, "" },
    }

    for name, it := range testCases {
		t.Run(name, func(t *testing.T) {
            ts := httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
                assert.Equal(t, r.Header.Get("Content-Type"), "application/json")
                assert.Equal(t, r.Header.Get("Authorization"), "Bearer aaa")
                w.WriteHeader(it.status)
            }))
            defer ts.Close()
            err := (&Healthcheck{ &Config{ "aaa", ts.URL}, []string{} }).Report(*&Report{})
            if it.err != "" {
                assert.Error(t, err, it.err)
            }
        })
    }
}