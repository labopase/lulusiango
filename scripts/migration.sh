#!/usr/bin/env bash
set -e

# Path to config file
CONFIG_FILE="configs/config.json"
MIGRATION_DIR="database/migrations"

mkdir -p "$MIGRATION_DIR"

if [ ! -f "$CONFIG_FILE" ]; then
    echo "Error: Configuration file $CONFIG_FILE not found."
    exit 1
fi

# Extract database configuration using Python (standard on macOS/Linux)
DB_HOST=$(python3 -c "import json; print(json.load(open('$CONFIG_FILE'))['postgres']['host'])")
DB_PORT=$(python3 -c "import json; print(json.load(open('$CONFIG_FILE'))['postgres']['port'])")
DB_USER=$(python3 -c "import json; print(json.load(open('$CONFIG_FILE'))['postgres']['user'])")
DB_PASS=$(python3 -c "import json; print(json.load(open('$CONFIG_FILE'))['postgres']['password'])")
DB_NAME=$(python3 -c "import json; print(json.load(open('$CONFIG_FILE'))['postgres']['dbname'])")
DB_SSL=$(python3 -c "import json; print(json.load(open('$CONFIG_FILE'))['postgres']['ssl_mode'])")

if [ -z "$DB_HOST" ] || [ -z "$DB_USER" ] || [ -z "$DB_NAME" ]; then
    echo "Error: Failed to parse database configuration from $CONFIG_FILE"
    exit 1
fi

# Construct DSN
DSN="postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL}"

# Check if migrate is installed
if ! command -v migrate &> /dev/null; then
    echo "Error: 'migrate' command not found. Please run 'make install' or install go-migrate."
    exit 1
fi

COMMAND=$1
shift

case "$COMMAND" in
    up)
        echo "Running migrations up..."
        migrate -path "$MIGRATION_DIR" -database "$DSN" up "$@"
        ;;
    down)
        echo "Running migrations down..."
        migrate -path "$MIGRATION_DIR" -database "$DSN" down "$@"
        ;;
    drop)
        echo "Dropping all migrations..."
        migrate -path "$MIGRATION_DIR" -database "$DSN" drop -f
        ;;
    force)
        if [ -z "$1" ]; then
            echo "Error: 'force' requires a version number."
            exit 1
        fi
        echo "Forcing migration version $1..."
        migrate -path "$MIGRATION_DIR" -database "$DSN" force "$1"
        ;;
    version)
        migrate -path "$MIGRATION_DIR" -database "$DSN" version
        ;;
    create)
        if [ -z "$1" ]; then
            echo "Error: 'create' requires a migration name."
            exit 1
        fi
        echo "Creating migration $1..."
        migrate create -ext sql -dir "$MIGRATION_DIR" -seq "$1"
        ;;
    *)
        echo "Usage: $0 {up|down|drop|force|version|create} [args...]"
        exit 1
        ;;
esac
