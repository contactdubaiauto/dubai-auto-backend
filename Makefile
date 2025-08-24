
deploy:
	@echo "Started building..."
	@GOOS=linux GOARCH=amd64 go build -o ./bin/da2 ./cmd/http/main.go
	@echo "Building done"

	@echo "Deploying..."
	@scp ./bin/da2 root@84.200.87.48:/var/www/da/da2
	@ssh root@84.200.87.48 "rm -f /var/www/da/da && mv /var/www/da/da2 /var/www/da/da"
	@echo "Restarting remote service..."
	@ssh root@84.200.87.48 "sudo -S systemctl restart da.service"
	@echo "Done"
	
	@scp -r ./docs root@84.200.87.48:/var/www/da
	# @scp -r ./images/logo root@84.200.87.48:/var/www/da/images
	# @scp -r ./images/body root@84.200.87.48:/var/www/da/images
	# @scp ./.env root@84.200.87.48:/var/www/da
	@echo "Done"
folder:
	@echo "deploying images..."
	@scp -r ./images root@84.200.87.48:/var/www/da/images
	@echo "done"
swag:
	@swag init -g cmd/http/main.go  

# swag init -g ./cmd/http/main.go
# todo: write swagger init command in this file

# -- for i in $(seq 1 30); do
# --   mkdir -p ./images/cars/$i
# --   cp ./images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1_l.jpg ./images/cars/$i/
# --   cp ./images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f1_m.jpg ./images/cars/$i/
# --   cp ./images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2_l.jpg ./images/cars/$i/
# --   cp ./images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f2_m.jpg ./images/cars/$i/
# --   cp ./images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3_l.jpg ./images/cars/$i/
# --   cp ./images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f3_m.jpg ./images/cars/$i/
# --   cp ./images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4_l.jpg ./images/cars/$i/
# --   cp ./images/cars/1/c3fba494-ca35-4be0-8345-3e3b1eb6f7f4_m.jpg ./images/cars/$i/
# -- done
