ifeq ($(OS), Windows_NT)
	RM = del
	TARGET = main.exe

else
	RM = rm
	TARGET = main
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