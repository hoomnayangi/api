# api

## Tech Stack

- Go
- Docker
- Postgresql
- Nginx

## API documentation

- [Insomnia](https://github.com/hoomnayangi/api/blob/main/hoomnayangi_2020-11-15.json)

## Setup

### Pre-install

- Install [golang](https://golang.org/doc/install)
- `go get -v github.com/rubenv/sql-migrate/...`

### Usage

- database setup:

  - `make init`

- app run on `local` environment
  - `make dev`

## App Architect

- `cmd`: app execute interface
- `data`: manage and migrating app's data
  - `seed`: app seed data for developing phase
  - `migration`: app data migration
- `src`: contains app main logic
  - `handler`: app routing, parsing
  - `model`: app's data model
  - `server`: portal between service, storage and app's handler
  - `service`: app's third-party implementation
  - `store`: database connection bridge
