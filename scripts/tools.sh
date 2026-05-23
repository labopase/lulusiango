#!/bin/bash

set -e

echo "Installing tools..."

# sqlc
# https://docs.sqlc.dev/en/stable/overview/install.html
if command -v sqlc &> /dev/null; then
    echo "sqlc already installed"
else
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.30.0
fi

# mockery
# https://docs.mockery.dev/en/stable/installation.html
if command -v mockery &> /dev/null; then
    echo "mockery already installed"
else
    go install github.com/vektra/mockery/v2@latest
fi

# migrate
# https://github.com/golang-migrate/migrate/releases
if command -v migrate &> /dev/null; then
    echo "migrate already installed"
else
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.3
fi

echo "All tools installed successfully!"