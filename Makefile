DEFAULT_GOAL := build

run:
	@echo "🦦🏃‍♂️ Starting Project..."
	@cd src && go run main.go

build:
	@echo "🛠 Building Project..."
	@cd src && go build main.go

test:
	@echo "👁‍🗨 Testing Project..."
	@cd src && go test ./... -cover

coverage:
	@echo "👁‍🗨 Testing Project with coverage..."
	@cd src && go test ./... -cover -coverprofile=c.out && go tool cover -html=c.out -o coverage.html

up:
	@echo "🍃🔓 Starting mongo database..."
	@docker-compose up

up-d:
	@echo "🍃🔓 Starting mongo database with -d option..."
	@docker-compose up -d

down:
	@echo "🍃🔒 Closing mongo database..."
	@docker-compose down

down-rm:
	@echo "🍃🔒 Closing mongo database and removing mongodb directory..."
	@docker-compose down && sudo rm -rf mongodb

docker-backend:
	@echo "🐳 Running docker backend container..."
	@docker build . -t nutriguide-backend-image:test && docker run --rm --name nutriguide-backend -p 8080:8080 nutriguide-backend-image:test
