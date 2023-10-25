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



docker-up:
	@docker-compose up -d

migrate-up:
	@go build -o $(TARGET_M) ./cmd/migrator/main.go
	@$(TARGET_M)
	@$(RM) $(TARGET_M)

db-connect:
	@docker exec -it person-storage psql -U supervisor -d person_dp