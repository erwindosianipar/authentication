dep:
	go mod tidy

run:
	go run app/main.go

test:
	go test -short -cover ./...

build:
	go build -o bin/authentication app/main.go

docker-image:
	docker build -t authentication:v1 .

docker-run:
	docker run -it -d -p 3000:3000 --name authentication
