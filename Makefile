build:
	docker build -t backend:1 .
	docker-compose build
run: build
	docker-compose up -d