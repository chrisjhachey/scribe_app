# OpenAPI

You can generate the models for scribe using the openAPI spec scribe.yml in the docs folder with the following command:

```
 go run /Users/christopherhachey/go/pkg/mod/github.com/deepmap/oapi-codegen@v1.9.1/cmd/oapi-codegen/oapi-codegen.go -generate types -o ./app/domain/scribe/model/scribe.gen.go ./docs/openapi/scribe.yaml
```

Additionally you can generate server code with this:

```
go run /Users/christopherhachey/go/pkg/mod/github.com/deepmap/oapi-codegen@v1.9.1/cmd/oapi-codegen/oapi-codegen.go -generate gin -o ./app/infrastructure/server/server.gen.go ./docs/openapi/scribe.yaml
```
