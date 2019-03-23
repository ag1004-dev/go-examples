
[This is a fork of this repo](https://github.com/bechurch/reverse-proxy-demo)

[And here is the blog post on the repo](https://hackernoon.com/writing-a-reverse-proxy-in-just-one-line-with-go-c1edfa78c84b)

Make sure you have a
[caddy binary](https://github.com/mholt/caddy/releases)
somewhere on your system.

Do the same thing for all of the caddy directories.

```
cd caddy10
caddy
```

```
source env1

go build bechurch.go
./bechurch

go run bechurch.go
```
