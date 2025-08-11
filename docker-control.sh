#!/bin/bash

# ------------------------------------------------------
# INSTRUCTIONS
# make executable with: chmod +x docker-control.sh
# Then use:
#   ./docker-control.sh build/start/stop...
# ------------------------------------------------------

set -e

COMMAND=$1

case $COMMAND in
    build)
        echo "Building Docker images..."
        docker compose build --no-cache
        ;;

    start)
        echo "Starting containers (API + Dev DB)..."
        docker compose up -d db_dev db_test backend
        echo "Containers started!"
        echo "API:     http://localhost:8080/api/hello"
        ;;

    stop)
        echo "Stopping and removing containers..."
        docker compose down
        ;;

    stop-all)
        echo "Stopping containers and deleting volumes (Postgres data will be lost)..."
        docker compose down -v
        ;;

    test)
        echo "Exporting .env variables"
        export TEST_DATABASE_URL="$(awk -F= '/^TEST_DATABASE_URL=/{print substr($0,index($0,$2))}' .env)"
        echo "Running tests"
        go test -tags=postgres ./test -v -race -count=1
        ;;

    logs)
        echo "Showing live logs..."
        docker compose logs -f
        ;;

    cleanup)
        echo "Stopping and removing all containers, volumes, and images for this project..."
        docker compose down -v --rmi all --remove-orphans
        echo "Pruning dangling volumes (if any)..."
        docker volume prune -f
        ;;

    help|*)
        echo ""
        echo "Usage: $0 {build|start|stop|stop-all|init|logs|cleanup}"
        echo ""
        echo "  build      Build database and web services"
        echo "  start      Start services (for dev use)"
        echo "  stop       Stop services (for dev use)"
        echo "  stop-all   Stop services and delete volumes (db data)"
        echo "  test       Runs test suite"
        echo "  init       Run initial setup for database (run after stop-all)"
        echo "  logs       Show container logs"
        echo "  cleanup    Remove containers, volumes, and images (complete reset)"
        echo ""
        exit 1
        ;;
esac