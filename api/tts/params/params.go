package params

import (
	"errors"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Params struct {
	key  string
	src  string
	hl   string
	r    string
	c    string
	f    string
	ssml string
	b64  string
}

func (p Params) GenUrl() string {
	url := url.URL{
		Scheme: "http",
		Host:   "api.voicerss.org",
	}

	q := url.Query()

	q.Set("key", p.key)
	q.Set("src", p.src)
	q.Set("hl", p.hl)
	q.Set("b64", p.b64)

	if p.r != "" {
		q.Set("r", p.r)
	}

	if p.c != "" {
		q.Set("c", p.c)
	}

	if p.f != "" {
		q.Set("f", p.f)
	}

	if p.ssml != "" {
		q.Set("f", p.ssml)
	}

	return url.String() + "?" + q.Encode()
}

func Parse(c *gin.Context, key *string) (Params, error) {
	src := c.Query("src")
	hl := c.Query("hl")

	if src == "" || hl == "" {
		return Params{}, errors.New("Missing required parameters")
	}

	p := Params{
		key: *key,
		src: src,
		hl:  hl,
		r:   c.Query("r"),
		c:   c.Query("c"),
		f:   c.Query("f"),
		ssml:   c.Query("ssml"),
		b64: "true",
	}

	return p, nil
}
