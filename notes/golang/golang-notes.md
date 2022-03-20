
#### Example Go builds

```
go help get
```

The -t flag instructs get to also download the packages required to build
the tests for the specified packages.

The -u flag instructs get to use the network to update the named packages
and their dependencies. By default, get uses the network to check out
missing packages but does not use it to look for updates to existing packages.

Assuming you do NOT have a **GOPATH** set; this will build and put
the working binaries in ~/go/bin

We will use
[vegeta](https://github.com/tsenart/vegeta)
```
go get -u github.com/tsenart/vegeta
```

We will use
[k6](https://github.com/loadimpact/k6)
```
go get github.com/loadimpact/k6
```

If you do have a **GOPATH** set; this will put everything
in the exact same spot as not having a **GOPATH** set but
in the spot of where GOPATH is set...

We will use
[badger](https://github.com/dgraph-io/badger#installing)

```
cd /tmpXX
gop
gopc
go get -t github.com/dgraph-io/badger/...
cd src/github.com/dgraph-io/badger/
go test -v
```

We will use
[tormenta](https://github.com/jpincas/tormenta)

```
cd /tmpXX
gop
gopc
go get -t github.com/jpincas/tormenta
cd src/github.com/jpincas/tormenta/
go test -v
```

We will use
[istio](https://github.com/istio/istio) as an example
of how to build a Golang app that does not have **go.mod**
at the top level in its repository.

```
mkdir src
cd src
mkdir istio.io
cd istio.io
git clone git@github.com:istio/istio.git
make depend
make build
```

We will use
[hugo](https://github.com/gohugoio/hugo) as an example
of how to build a Golang app that does have it...

```
mkdir $HOME/src
cd $HOME/src
git clone https://github.com/gohugoio/hugo.git
cd hugo
go install
```

[Here are the instructions for Hugo](https://github.com/gohugoio/hugo#fetch-from-github)

#### Setup from Scratch

Assuming you do not have a go.mod in your repository  
At the top level of any repo type this command

```
go mod init github.com/stormasm/redishacker
```

### Testing

```
go test -run TestBulkStringRequestSetString
```

#### Benchmark Testing

[count_test.go](https://github.com/influxdata/flux/blob/master/stdlib/universe/count_test.go)

```
gtvr BenchmarkCount -bench=BenchmarkCount
```

#### How to clear the test cache and do verbose logging

* Shows how to clear the testing cache every time
* Shows how to do verbose test logging

```
go test -count=1 -v ./galley/pkg/source/kube -run Builtin
```

#### Running

```
cd fetch
go run fetch.go
```

#### Building

So this took me a while to figure out...
Assumes we are using go.mod and no GOPATH

From a top level tmp directory

```
mkdir src   
git clone git@github.com:pelletier/go-toml.git
```

**And here is the key point**

You must be in a directory with a package **main** for go build to work...

```
cd cmd/tomljson
go build
```

And if you want the binary called **rick** instead of tomljson

```
go build -o rick
```

##### Yet another example

In the [flux repo](https://github.com/influxdata/flux)

If you...

```
cd cmd/flux
go build
```

It will put the binary right there in that directory.

If you...

```
go install
```

It will put the binary in **~/go/bin** or whatever GOBIN is defined to be...   
**PLUS** it will remove the binary from the directory where go build was previously run...
