## Postgres playground

### Setup

1. Install psql (postgres client) (if you don't have it already)
2. Run `make setup` to setup golang-migrate
3. Run `docker compose up -d` to start the database
4. Run `make m_up` to apply the migrations

### Connect to the database

```bash
make start
```
Type `pass` for the password

## Migrations

### Create a new migration

```bash
make m_create NAME="your_migration_name"
```

### Apply migrations

```bash
make m_up
```

```bash
make m_down
```
