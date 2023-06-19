package params

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	r := &gin.Context{
		Request: &http.Request{
			URL: &url.URL{
				RawQuery: "src=src&hl=hl&r=r&c=c&f=f&ssml=ssml",
			},
		},
	}

	key := "key"

	p, err := Parse(r, &key)

	assert.Nil(t, err)
	assert.Equal(t, "key", p.key)
	assert.Equal(t, "src", p.src)
	assert.Equal(t, "hl", p.hl)
	assert.Equal(t, "r", p.r)
	assert.Equal(t, "c", p.c)
	assert.Equal(t, "f", p.f)
	assert.Equal(t, "ssml", p.ssml)
	assert.Equal(t, "true", p.b64)
}

func TestParseFail(t *testing.T) {
	r := &gin.Context{
		Request: &http.Request{
			URL: &url.URL{
				RawQuery: "src=src",
			},
		},
	}

	key := "key"

	_, err := Parse(r, &key)

	assert.NotNil(t, err)
}

func TestGenUrl(t *testing.T) {
	p := Params{
		key:  "key",
		src:  "src",
		hl:   "hl",
		r:    "r",
		c:    "c",
		f:    "f",
		ssml: "ssml",
		b64:  "b64",
	}

	assert.Equal(t, "http://api.voicerss.org?b64=b64&c=c&f=ssml&hl=hl&key=key&r=r&src=src", p.GenUrl())
}
