# generate the proto files
all:
	protoc -I. --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative auth/*.proto

generate-auth:
	protoc -I. --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative auth/*.proto

generate-finances:
	protoc -I. --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative --grpc-gateway_out=./ --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out ./ finances/*.proto