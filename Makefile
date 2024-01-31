# Define the protoc command as a variable for easy modification and readability
PROTOC_CMD = protoc --proto_path=. \
       --proto_path=/home/levniko/go-proto-validators \
       --go_out=paths=source_relative:. \
       --go-grpc_out=paths=source_relative:. \
       --govalidators_out=paths=source_relative:. \
       --go-grpc_opt=require_unimplemented_servers=false \
       protobuf/user/user.proto
       


# Define a target for generating protobuf files
generate-proto:
	$(PROTOC_CMD) protobuf/user/user.proto
