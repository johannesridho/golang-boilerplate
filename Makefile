gen-proto:
	for dir in $(wildcard protobuf-files/*/); do \
		cd $$dir && protoc -I. \
			--go_out=plugins=grpc:. \
			--swagger_out=logtostderr=true,allow_delete_body=true:. \
			./*.proto \
		&& cd ../; \
	done

run:
	go run cmd/api/main.go
