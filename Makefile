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

# Start the docker-compose services
docker-run:
	docker-compose up --build

# Stop the docker-compose services
docker-stop:
	docker-compose down

# Build docker-compose services without starting them
docker-build:
	docker-compose build

# Stop and remove containers, networks, images, and volumes
docker-clean-all:
	docker-compose rm -f
	docker system prune -af --volumes

# Remove all stopped containers
docker-clean-containers:
	docker container prune -f

# Remove all unused images
docker-clean-images:
	docker image prune -af

# Remove all unused volumes
docker-clean-volumes:
	docker volume prune -f

# Remove all unused networks
docker-clean-networks:
	docker network prune -f

# Restart all services
docker-restart: stop run