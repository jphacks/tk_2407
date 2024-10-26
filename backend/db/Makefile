.PHONY: start
start:
	psql -h localhost --port 5432 -U user -d jphacks

.PHONY: setup
setup:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: m_up
migrate_up:
	migrate -path migrations -database "postgresql://user:password@localhost:5432/jphacks?sslmode=disable" up 1

.PHONY: m_down
migrate_down:
	migrate -path migrations -database "postgresql://user:password@localhost:5432/jphacks?sslmode=disable" down 1

m_create:
ifndef NAME
	@echo "Usage: make NAME=migration_name m_create"
else
	migrate create -ext sql -dir migrations -seq $(NAME)
endif
