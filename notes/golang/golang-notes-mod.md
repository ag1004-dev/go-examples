
Last updated March 14, 2020

[How to create a go.mod from scratch](https://blog.golang.org/using-go-modules)

[Using this repo as an example](https://github.com/stormasm/yhoo-gq1)

```
To test how to run a repo without go.mod
Remove the go.mod file from this repo and run this command...
go mod init github.com/stormasm/yhoo-gq1

Otherwise since go.mod is present...
go get -t
go run quote/main.go ui

Or to install a binary

cd quote
go install

quote will now be installed in $HOME/go/bin

Then once quote is installed go to any directory and type

quote ibm
quote aapl
quote msft
```
