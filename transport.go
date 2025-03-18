package unsend

import (
	"fmt"
	"net/http"
)

type UnsendTransport struct {
	ApiKey string
}

func (t *UnsendTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("user-agent", fmt.Sprintf("%s/%s", PACKAGE_NAME, VERSION))
	req.Header.Add("version", VERSION)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.ApiKey))
	return http.DefaultTransport.RoundTrip(req)
}
