build:
	cd ./cmd/ && go build server.go
	mv ./cmd/server bin/server

image:
	docker build -t service-b-rest:1.0 . 

deploy:
	cd k8s/ && kustomize build . | kubectl apply -f -
