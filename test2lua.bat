go build -o excel2lua.exe ./bak/main2lua.go 
excel2lua.exe -i "./excel" -o "luadata" -d true
rem go run main2lua.go -i "./excel" -o "luadata" -d true
pause