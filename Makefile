openapi-download:
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
openapi-http:
	oapi-codegen -generate types -o internal/orders/delivery/http/types.openapi.gen.go -package http api/openapi/main.yml
	oapi-codegen -generate server -o internal/orders/delivery/http/api.openapi.gen.go -package http api/openapi/main.yml
generate:
	go generate ./...
	