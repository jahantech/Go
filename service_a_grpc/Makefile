build:
	cd ./cmd/client_a_grpc/ && go build
	cd ./cmd/service_a_grpc/ && go build
	mv ./cmd/client_a_grpc/client_a_grpc bin/client
	mv ./cmd/service_a_grpc/service_a_grpc bin/server

image:
	docker build -t service-a-grpc:1.0 . 

deploy:
	cd k8s/ && kustomize build . | kubectl apply -f -