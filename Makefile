build:
	echo "Generating docs"
	swag init
	echo "Compiling for your OS"
	go build -o ./bin/http-scheduler

compile:
	echo "Generating docs"
	swag init
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o ./bin/http-scheduler-linux-arm 
	GOOS=linux GOARCH=arm64 go build -o ./bin/http-scheduler-linux-arm64 
	GOOS=windows GOARCH=386 go build -o ./bin/http-scheduler-windows-386 
	GOOS=windows GOARCH=arm64 go build -o ./bin/http-scheduler-windows-arm64

run:
	clear
	echo "Generating docs"
	swag init
	echo "Running your program"
	go run main.go

clean:
	echo "Cleaning outputs"
	@if [ -d "./bin" ]; then rm -r ./bin; fi
	swag fmt

setup: # Use this command to setup the environment
	echo "Setting docker container"
	sudo docker run -d --name http-scheduler-redis -p 6379:6379 -e REDIS_PASSWORD=redis redis

init-setup: # Use this command to init the setup
	echo "Initializing setup"
	echo "Initializing Redis server"
	sudo docker start http-scheduler-redis

