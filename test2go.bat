go build -o excel2go.exe ./bak/main2go.go
excel2go -i "./excel" -o "godata"
rem run main2go.go -i "./excel" -o "godata" -n "godata" -d true
pause