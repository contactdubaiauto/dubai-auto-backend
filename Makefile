dev:
	@go run main.go

build:
	@go build main.go

deploy:
	@echo "Started building..."
	@env GOOS=linux GOARCH=amd64 go build -o bin/
	@echo "Building done"

	@echo "Deploying..."
	@scp ./bin/payment payment@95.85.108.126:/var/www/
	@echo "Done"