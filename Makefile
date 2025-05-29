dev:
	@go run main.go

build:
	@go build main.go

deploy:
	@echo "Started building..."
	@env GOOS=linux GOARCH=amd64 go build -o bin/
	@echo "Building done"

	@echo "Deploying..."
	@scp ./bin/dubai-auto user@0.0.0.0:/var/www/
	@echo "Done"