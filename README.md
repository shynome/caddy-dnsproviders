### usage

```go
package caddyhttp
import (
  "github.com/shynome/caddy-dnsproviders/alidns"
  acme "github.com/xenolf/lego/acmev2"
)
// 如果存在环境变量 RECURSIVE_NAMESERVERS 的话会将 DNS 服务器更改为 RECURSIVE_NAMESERVERS 的值
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
