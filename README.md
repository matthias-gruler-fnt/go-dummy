# go-dummy
Go Dummy Web App

## API
see file `test.http`

Use VS Code extension [Rest Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) to call endpoits from within editor.


## Develop

### Prerequisites
Install tools:

[Air](https://github.com/cosmtrek/air) - for live-reloading
```shell
go install github.com/cosmtrek/air@latest
```

[Task](https://taskfile.dev/installation/) - build tool
```shell
go install github.com/go-task/task/v3/cmd/task@latest
```

### Run local dev server

Starts a local server on port 8000 with hot-reloading:
```shell
task dev
```
Access server:  
[http://localhost:8000/](http://localhost:8000/)