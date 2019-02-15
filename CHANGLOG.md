# CHANGELOG

## 1.1.0

### BREAK

- 将设置默认 DNS 的行为调整为一个函数, 由引入者决定是否调用. 
并且如果环境变量 "RECURSIVE_NAMESERVERS" 不提供的话是无法设置 DNS 服务器的
