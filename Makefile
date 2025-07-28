
deploy:
	@echo "Started building..."
	@GOOS=linux GOARCH=amd64 go build -o ./bin/da2 ./cmd/http/main.go
	@echo "Building done"

	@echo "Stopping remote service..."
	# @ssh root@95.85.126.220 "sudo -S systemctl stop da.service"

	@echo "Deploying..."
	@scp ./bin/da2 root@95.85.126.220:/var/www/
	@ssh root@95.85.126.220 "rm -f /var/www/da && mv /var/www/da2 /var/www/da"
	@echo "Restarting remote service..."
	@ssh root@95.85.126.220 "sudo -S systemctl restart da.service"
	@echo "Done"
	
	# @scp ./images/logo/audi.png root@95.85.126.220:/var/www/images/logo
	# @scp -r ./docs root@95.85.126.220:/var/www/
	# @scp -r ./images/body root@95.85.126.220:/var/www/images
	# @scp ./.env root@95.85.126.220:/var/www/
	
	@echo "Starting remote service..."
	# @ssh root@95.85.126.220 "sudo -S systemctl start da.service"
	@echo "Done"
folder:
	@echo "deploying images..."
	@scp -r ./images root@95.85.126.220:/var/www/images
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
