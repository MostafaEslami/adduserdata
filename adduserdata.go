// Package plugin_adduserdata a plugin to block a path.
package plugin_adduserdata

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
)

// Config holds the plugin configuration.
type Config struct {
	Regex []string `json:"regex,omitempty"`
}

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

type blockPath struct {
	name    string
	next    http.Handler
	regexps []*regexp.Regexp
}

// New creates and returns a plugin instance.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	regexps := make([]*regexp.Regexp, len(config.Regex))

	for i, regex := range config.Regex {
		re, err := regexp.Compile(regex)
		if err != nil {
			return nil, fmt.Errorf("error compiling regex %q: %w", regex, err)
		}

		regexps[i] = re
	}

	return &blockPath{
		name:    name,
		next:    next,
		regexps: regexps,
	}, nil
}

func (b *blockPath) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	req.Header.Set("mostafa", "mostafa")
	b.next.ServeHTTP(rw, req)
}
