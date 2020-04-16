build :
	cd ./cmd/client_a_grpc/ && go build
	cd ./cmd/service_a_grpc/ && go build
	cp ./cmd/client_a_grpc/client bin/.
	cp ./cmd/service_a_grpc/service_a_grpc bin/.

image:
	docker build -t service_a_grpc:1.0 . 