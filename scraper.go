package datapegasus

import "github.com/valyala/fasthttp"

func (p *Pegasus) Get(dst []byte, url string) (statusCode int, body []byte, err error) {
	req := fasthttp.AcquireRequest()
	// req.Header.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	// req.Header.SetUserAgent(p.UserAgent)
	// req.Header = *p.Headers

	req.Header = fasthttp.RequestHeader{}

	p.Headers.CopyTo(&req.Header)

	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return 0, nil, err
	}

	statusCode = resp.StatusCode()
	body = resp.Body()

	return statusCode, body, nil
}
