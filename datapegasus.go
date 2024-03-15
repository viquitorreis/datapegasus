package datapegasus

import (
	"github.com/valyala/fasthttp"
)

// container de configuração para o Pegasus
type Options struct {
	UserAgent string
	Headers   *fasthttp.RequestHeader
}

// Handler do Pegasus
type Pegasus struct {
	UserAgent string
	Headers   *fasthttp.RequestHeader
}

func NewPegasus(options Options) *Pegasus {
	p := &Pegasus{
		UserAgent: options.UserAgent,
		Headers:   options.Headers,
	}

	if p.UserAgent == "" {
		p.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/"
	}

	p.Headers = &fasthttp.RequestHeader{}
	p.Headers.Set("User-Agent", p.UserAgent)

	// ARRUMAR CONDIÇÕES CORRETAMENTE
	if p.Headers == nil {
		p.Headers = &fasthttp.RequestHeader{}
		p.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		p.Headers.Set("Accept-Encoding", "gzip, deflate, br, sdch")
		p.Headers.Set("Accept-Language", "en-US,en;q=0.8")
		p.Headers.Set("Referer", "https://www.google.com/")
	}

	return p
}

func Default() *Pegasus {
	return NewPegasus(Options{})
}
