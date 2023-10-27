ifeq ($(OS), Windows_NT)
	RM = del
	TARGET = main.exe
	TARGET_M = migrator.exe

else
	RM = rm
	TARGET = main
	TARGET_M = migrator
endif

run: build
	@$(TARGET)

test:
	@go test -cover ./...

cover:
	@go test -coverprofile=coverage.out ./...

cover_html: cover
	@go tool cover -html=coverage.out

build:
	@go build -o $(TARGET) ./cmd/app/main.go

clear:
	@$(RM) $(TARGET)

mock_gen:
	@mockgen -source=internal/domain/ports/storage/person.go -destination=pkg/mocks/storage/person_mock.go
	@mockgen -source=internal/domain/ports/enricher/enricher.go -destination=pkg/mocks/enricher/enricher_mock.go
	@mockgen -source=internal/adapters/api/service/person.go -destination=pkg/mocks/api/service/person_mock.go

swagger-gen:
	@swag init -g ./internal/adapters/api/router/person.go

docker-up:
	@docker-compose up -d

migrate-up:
	@go build -o $(TARGET_M) ./cmd/migrator/main.go
	@$(TARGET_M) -action="up"
	@$(RM) $(TARGET_M)

migrate-init:
	@go build -o $(TARGET_M) ./cmd/migrator/main.go
	@$(TARGET_M) -action="init"
	@$(RM) $(TARGET_M)

migrate-down:
	@go build -o $(TARGET_M) ./cmd/migrator/main.go
	@$(TARGET_M) -action="down"
	@$(RM) $(TARGET_M)

db-connect:
	@docker exec -it person-storage psql -U supervisor -d person_db