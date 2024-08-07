build:
	echo "Generating docs"
	swag init --parseDependency
	echo "Compiling for your OS"
	go build -o ./bin/http-scheduler

compile:
	echo "Generating docs"
	swag init --parseDependency
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o ./bin/http-scheduler-linux-arm 
	GOOS=linux GOARCH=arm64 go build -o ./bin/http-scheduler-linux-arm64 
	GOOS=windows GOARCH=386 go build -o ./bin/http-scheduler-windows-386 
	GOOS=windows GOARCH=arm64 go build -o ./bin/http-scheduler-windows-arm64

run:
	clear
	echo "Generating docs"
	swag init --parseDependency
	echo "Running your program"
	go run main.go

clean:
	echo "Cleaning outputs"
	@if [ -d "./bin" ]; then rm -r ./bin; fi
	swag fmt

remove-setup: # Use this command to remove all Docker containers and network
	docker compose down --volumes --remove-orphans

init-setup: # Use this command to init the setup
	docker compose up -d
