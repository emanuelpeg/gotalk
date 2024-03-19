# gotalk
Chat in golang

## Getting started

### Global Dedendencies

- Goctl is a code generation tool based on *.api* files. Run `go install github.com/zeromicro/go-zero/tools/goctl@latest`.
- Run `nano ~/.profile` and add a new line with `export PATH="$HOME/go/bin:$PATH"` to have GoCTL avaiable globally.
- Run `. ~/.profile`.

### Create a Go service

- Create a folder with the name of the service you want to create `mkdir apiname`.
- Create an *api* subfolder inside the newly created folder `mkdir apiname/api`.
- Go to the new service folder `cd apiname`.
- To generate the API, run `goctl api -o api/apiname.api`, open the file and assign a title and a description to it.
- To generate Go code with a service based off an existing API, run `goctl api go -api api/apiname.api -dir .`, the service will have the same name as the API.
- Run `go mod tidy` to install all of the service dependencies.
- To install goctl missing dependencies if necessary, run `goctl env check --install --verbose --force`


## Run a Service

- To start the **gotalk** service run `go run gotalk/gotalk.go -f gotalk/etc/gotalk-api.yaml`.
- To start the **gotalkRouter** service run `go run gotalkRouter/gotalkrouter.go -f gotalkRouter/etc/gotalkRouter-api.yaml`.
- Alternatively, if the first option gives you an error about the import paths:
  - Navigate to the gotalk directory first using `cd gotalk`
  - And then `go run gotalk.go -f etc/gotalk-api.yaml`.


### Endpoints
- Gotalk example
  - `curl -i "http://127.0.0.1:8888/from/me"`
  - `curl -i "http://127.0.0.1:8888/from/you"`
- Gotalk router
  - `curl -i "http://127.0.0.1:8889/from/me"`
- Health check
  - `curl -i "http://127.0.0.1:8888/health"`


## Good Reads

- [Rest vs gRPC](https://blog.postman.com/how-to-choose-http-or-grpc-for-your-next-api/)