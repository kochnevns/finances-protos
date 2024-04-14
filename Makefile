# generate the proto files
all:
	protoc -I. --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative auth/*.proto

generate-auth:
	protoc -I. --go_out=./auth --go_opt=paths=source_relative --go-grpc_out=./auth --go-grpc_opt=paths=source_relative auth/*.proto

generate-finances:
	protoc -I. --go_out=./finances --go_opt=paths=source_relative --go-grpc_out=./finances --go-grpc_opt=paths=source_relative finances/*.proto