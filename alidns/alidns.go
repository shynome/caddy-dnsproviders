package alidns

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/challenge/dns01"
	"github.com/xenolf/lego/providers/dns/alidns"
)

// Nameservers of ali
var Nameservers = []string{"dns21.hichina.com:53", "dns22.hichina.com:53"}

func init() {
	caddytls.RegisterDNSProvider("alidns", NewDNSProvider)
	caddytls.RegisterDNSProvider("alicloud", NewDNSProvider)
	dns01.AddRecursiveNameservers(Nameservers)(&dns01.Challenge{})
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
		dnsProvider, err := alidns.NewDNSProviderConfig(config)
		return dnsProvider, err
	default:
		return nil, errors.New("invalid credentials length")
	}
}
