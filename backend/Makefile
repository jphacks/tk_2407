TOOLS = \
	github.com/deepmap/oapi-codegen/cmd/oapi-codegen

gen: $(TOOLS)
	@sh -c "oapi-codegen --package=openapi -generate gin --old-config-style ../openapi/openapi.yaml > pkg/openapi/server.gen.go"
	@sh -c "oapi-codegen --package=openapi -generate types --old-config-style ../openapi/openapi.yaml > pkg/openapi/types.gen.go"
	@sh -c "oapi-codegen --package=openapi -generate client --old-config-style ../openapi/openapi.yaml > pkg/openapi/client.gen.go"

$(TOOLS):
	go install $@@latest
