
deploy:
	@echo "Started building..."
	@GOOS=linux GOARCH=amd64 go build -o ./bin/da2 ./cmd/http/main.go
	@echo "Building done"

	@echo "Stopping remote service..."
	# @ssh ubuntu@95.85.126.220 "sudo -S systemctl stop da.service"

	@echo "Deploying..."
	@scp ./bin/da2 ubuntu@95.85.126.220:/var/www/
	# @scp -r ./docs ubuntu@95.85.126.220:/var/www/
	
	# @scp -r ./images ubuntu@95.85.126.220:/var/www
	# @scp ./.env ubuntu@95.85.126.220:/var/www/
	
	@echo "Starting remote service..."
	# @ssh ubuntu@95.85.126.220 "sudo -S systemctl start da.service"
	@echo "Done"
folder:
	@echo "deploying images..."
	@scp -r ./images ubuntu@95.85.126.220:/var/www/images
	@echo "done"

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
