package alidns

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/alidns"
)

func init() {
	caddytls.RegisterDNSProvider("alidns", NewDNSProvider)
}

// NewDNSProvider returns a new Aliyun DNS challenge provider.
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return alidns.NewDNSProvider()
	case 2:
		config := alidns.NewDefaultConfig()
		config.APIKey = credentials[0]
		config.SecretKey = credentials[1]
		return alidns.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
