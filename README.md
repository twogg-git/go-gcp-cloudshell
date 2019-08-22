# go-light
Simple go endpoints to test with ELK

```sh
docker run -v $pwd:/go/src/app -p 8181:8080 -d --name go-light golang:1.8 go run src/app/main.go

docker run -v $(pwd):/go/src/app --rm --name testgo golang:1.8 go run main.go
```


```sh
localhost:8181
localhost:8181/time
localhost:8181/version
```

```sh
docker run -v $pwd:/go/src/app -p 8181:8080 -d --name go-light golang:1.8 /bin/bash -c "go run src/app/main.go"
```

```sh
docker run -v $pwd:/go/src/app -p 8181:8080 -d --name go-light golang:1.8 /bin/bash -c "cd src/app;go get -v ./...;go run /src/app/main.go"
```

```sh
docker rm -f go-light
```

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://console.cloud.google.com/cloudshell/editor?shellonly=true&cloudshell_image=gcr.io/cloudrun/button&cloudshell_git_repo=https://github.com/twogg-git/go-light.git)
