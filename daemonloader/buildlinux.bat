SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go clean
go build -o ..\..\build\bin\ldaemonloader daemonloader.go

