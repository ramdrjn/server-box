all: *.proto
	protoc --go_out=. --go_opt=paths=source_relative $^
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative $^
