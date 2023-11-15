//go:build generate

//go:generate go test -v -coverpkg=./... -coverprofile=coverage.txt ./...
//go:generate go tool cover -func=coverage.txt
//go:generate go tool cover -html=coverage.txt
//go:generate powershell -Command Remove-Item coverage.txt

package main

func main() {

}
