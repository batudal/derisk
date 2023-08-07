start: build
	docker-compose down
	@echo "Docker images closed!"
	docker-compose up --build -d 
	@echo "Docker images built and started!"

stop:
	docker-compose down
	@echo "Docker images closed!"

build:
	@echo "Building app binary..."
	cd ./app && bash /opt/tailwindcss -i ./public/app.css -o ./public/tw.css --minify
	cd ./app && env GOOS=linux CGO_ENABLED=0 go build -o de-risk .
	@echo "Done!"

dev: hmr 
	@echo "Initialized dev mode."
hmr:
	@echo "Engaging hmr mode..."
	cd ./app && air -- -dev
	@echo "Done!"

