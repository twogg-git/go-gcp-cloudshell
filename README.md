# GO + GCP 
Simple GO endpoints deployed in Google Cloud Platform directy from GitHub clicking just one button!

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://console.cloud.google.com/cloudshell/editor?shellonly=true&cloudshell_image=gcr.io/cloudrun/button&cloudshell_git_repo=https://github.com/twogg-git/go-gcp-cloudshell.git)

### Testing locally
```sh
docker run -v $pwd:/go/src/app -p 8181:8080 -d --name go-endpoints golang:1.8 go run src/app/main.go

docker run -v $(pwd):/go/src/app --rm --name go-endpoints golang:1.8 go run main.go
```

### Endpoints available
```sh
localhost:8181
localhost:8181/time
localhost:8181/version
```

### Docker + Go commands
```sh
docker run -v $pwd:/go/src/app -p 8181:8080 -d --name go-endpoints golang:1.8 /bin/bash -c "go run src/app/main.go"

docker run -v $pwd:/go/src/app -p 8181:8080 -d --name go-endpoints golang:1.8 /bin/bash -c "cd src/app;go get -v ./...;go run /src/app/main.go"
```
### Cleaning up
```sh
docker rm -f go-endpoints
```

[GCP Cloud run button repo](https://github.com/GoogleCloudPlatform/cloud-run-button#add-the-cloud-run-button-to-your-repos-readme)
