
deploy:
	@echo "Started building..."
	@GOOS=linux GOARCH=amd64 go build -o ./bin/da2 ./cmd/http/main.go
	@echo "Building done"

	@echo "Deploying..."
	@scp ./bin/da2 root@84.200.87.48:/var/www/da/da2
	@ssh root@84.200.87.48 "rm -f /var/www/da/da && mv /var/www/da/da2 /var/www/da/da"
	
	@scp -r ./docs root@84.200.87.48:/var/www/da
	# @scp -r ./images/logo root@84.200.87.48:/var/www/da/images
	# @scp -r ./images/body root@84.200.87.48:/var/www/da/images
	# @scp ./.env root@84.200.87.48:/var/www/da
	@echo "Restarting remote service..."
	@ssh root@84.200.87.48 "sudo -S systemctl restart da.service"
	@echo "Done"
	
swag:
	@swag init -g cmd/http/main.go  


