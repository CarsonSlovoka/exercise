//go:build generate

//go:generate go test -v -coverpkg=./... -coverprofile=coverage.txt ./...
// //go:generate go test -v -coverpkg=./... -coverprofile=coverage.txt ./algorithms/search/binary_search/exercise1/ //<-- 指測試某一個目錄
//go:generate go tool cover -func=coverage.txt
//go:generate go tool cover -html=coverage.txt
//go:generate powershell -Command Remove-Item coverage.txt

package main

func main() {

}
