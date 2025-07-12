
deploy:
	@echo "Started building..."
	@GOOS=linux GOARCH=amd64 go build -o ./bin/da ./cmd/http/main.go
	@echo "Building done"

	@echo "Stopping remote service..."
	@ssh ubuntu@95.85.126.220 "sudo -S systemctl stop da.service"

	@echo "Deploying..."
	@scp ./bin/da ubuntu@95.85.126.220:/var/www/
	@scp ./docs/docs.go ubuntu@95.85.126.220:/var/www/docs/
	@scp ./docs/swagger.json ubuntu@95.85.126.220:/var/www/docs/
	@scp ./docs/swagger.yaml ubuntu@95.85.126.220:/var/www/docs/
	
	# @scp -r ./images ubuntu@95.85.126.220:/var/www
	# @scp ./.env ubuntu@95.85.126.220:/var/www/
	
	@echo "Starting remote service..."
	@ssh ubuntu@95.85.126.220 "sudo -S systemctl start da.service"
	@echo "Done"
folder:
	@echo "deploying images..."
	@scp -r ./images ubuntu@95.85.126.220:/var/www/images
	@echo "done"

# swag init -g ./cmd/http/main.go
# todo: write swagger init command in this file