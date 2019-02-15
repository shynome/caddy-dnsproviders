### 警告

使用该插件会更改默认dns服务器为 `dns21.hichina.com,dns22.hichina.com`.		
更改的原因是使用谷歌dns服务器经常无法申请阿里云域名下的证书

### usage

```go
package caddyhttp
import (
  "github.com/shynome/caddy-dnsproviders/alidns"
	acme "github.com/xenolf/lego/acmev2"
)
var _alidnsInit = alidns.SetRecursiveNameservers(&acme.RecursiveNameservers)
```

### caddy config

```conf
mydomain.com {
  tls {
    dns alidns
  }
}
mydomain.com {
  tls {
    dns alicloud
  }
}
```
