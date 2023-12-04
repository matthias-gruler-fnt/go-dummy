# go-dummy
Go Dummy Web App

## API
see file `test.http`

Use VS Code extension [Rest Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) to call endpoits from within editor.

## Build

[goreleaser](https://goreleaser.com/) is used for building and releasing. It's a preqrequisite which needs to be installed before:
```shell
go install github.com/goreleaser/goreleaser@latest
```

### Local binaries only

Build snapshot binaries for Linux and Windows in subdirectories of `dist`.
```shell
task build
```

### Local archives and container image

Build snapshot binaries and archives for Linux and Windows in subdirectories of `dist`.  
Builds a container image and stores it in local registry. Image name is: `matthias-gruler-fnt/go-dummy`
```shell
task release-local
```

## Release

A release builds, packages and publishes artifacts and images to Github/container registry.  
Releases depend on Git tags.

⚠️ __Prerequisite__:  
Make sure you have a clean Git working directory without changes.

Steps:
1. Increment version number and create a new Git tag:
    ```shell
    # assumed current version is 0.1.0
    git tag v0.2.0
    ```
    Alternatively, use [SVU](https://github.com/caarlos0/svu) to generate next version number:
    ```shell
    git tag "$(svu next)"
    ```
1. Make sure to have variable GITHUB_TOKEN set in your environment:
    ```powershell
   $env:GITHUB_TOKEN="___REPLACE_ME___"
    ```
1. Create the release: 
    ```shell
    task release
    ```

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

### Local build

Builds `go-dummy.exe` in project directory:
```shell
task build
```