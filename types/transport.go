package types

import "net/http"

type DefaultHeaderTransport struct {
	N    int64
	next http.RoundTripper
}

func (tr *DefaultHeaderTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("service", "search-microservice")
	if tr.next != nil {
		return tr.next.RoundTrip(r)
	}
	return http.DefaultTransport.RoundTrip(r)
}
