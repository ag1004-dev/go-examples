
##### How to turn off the test cache

https://github.com/openshift/machine-config-operator/pull/873

The new way to disable the go testing cache is to use -count=1 in the go test command.

https://github.com/golang/go/issues/24573

For the top level Makefile for the flux project

```
export GO_TEST=env GO111MODULE=on go test -count=1 $(GO_ARGS)
```

### Testing

The first thing you need to do before reading this is to run
the command

##### go help test

If you really got what it says then you should not have to read these notes...

Golang testing is much more challenging to wrap your head around
than testing in Rust files...

In Go testing there are 2 different ways to think about tests

Go testing works on test names inside files not individual file names...
But you can test across all files in a particular directory,
but not individual files...

When you type go test -run you are talking about test names inside files

The **-run** signals you are dealing with test names

Here are my aliases for testing...

### Aliases for Testing across all filenames in a directory

```
alias gt='go test'
alias gtv='gt -v'
```

### Aliases for Testing names of tests inside files

```
alias gtr='go test -run'
alias gtvr='gtv -run'
```

### This will run all of the tests in a particular directory

```
alias gtd='go test ./...'
alias gtdv='go test -v ./...'
```

##### Bench testing

```
alias gtb='go test -v -bench .'
alias gtb1='go test -v -bench=SlowIcingManyIcers'
```

##### You must look inside the file at the code to grab an actual test name

This is not about the file name but about the name of the test in the actual code

```
go test -run TestBulkStringRequestSetString
```

#### Benchmark Testing

[Ref dave.cheney](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

[index_test.go](https://github.com/influxdata/influxdb/blob/master/tsdb/index_test.go)

```
go test -bench=BenchmarkIndex_ConcurrentWriteQuery
```

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
