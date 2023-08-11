DOTENV_FILE=./secrets.env
DOCKER_IMAGE=game_server_image
DOCKER_CONTAINER=game_server
DOCKER_ROOMPORT=8080


all: build run 

test:
	go run cmd/main.go

build: 
	docker build -t ${DOCKER_IMAGE} . 

run: 
	docker run --name ${DOCKER_CONTAINER} --env-file ${DOTENV_FILE}  -p ${DOCKER_ROOMPORT}:${DOCKER_ROOMPORT} ${DOCKER_IMAGE}

clean:
	docker stop ${DOCKER_CONTAINER} &&  docker rm ${DOCKER_CONTAINER} && docker rmi ${DOCKER_IMAGE}


