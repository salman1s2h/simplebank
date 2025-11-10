name: ci-test

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:
  build:
    runs-on: ubuntu-latest

    services:
      postgres:
        image: postgres:13
        env:
          POSTGRES_USER: admin
          POSTGRES_PASSWORD: admin
          POSTGRES_DB: go_db
        ports:
          - 5433:5432
        options: >-
          --health-cmd pg_isready -U admin
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5

    steps:
    - uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.25.1'

    - name: Install migrate CLI
      run: |
        curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz
        sudo mv migrate /usr/local/bin/
        migrate -version  # sanity check

    - name: Build
      run: go build -v ./...

    - name: Migrate up
      env:
        DB_URL: postgresql://admin:admin@localhost:5433/go_db?sslmode=disable
      run: make migrateup

    - name: Test
      run: go test -v ./...