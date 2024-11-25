# go-UUID
CLI to generate UUID (based on [gofrs/uuid](https://github.com/gofrs/uuid))

## Go Build

```
GOOS=darwin GOARCH=arm64 go build -o uuid-arm main.go
GOOS=darwin GOARCH=amd64 go build -o uuid-amd main.go
lipo -create uuid-arm uuid-amd -output uuid
GOOS=windows GOARCH=amd64 go build -o uuid.exe main.go
```
