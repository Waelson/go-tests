stack_build:
	@echo "Running docker-compose..."
	docker-compose up --build -d

stack_up:
	@echo "Running docker-compose..."
	docker-compose up -d

stack_down:
	@echo "Stopping docker-compose..."
	docker-compose down
