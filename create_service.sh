#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <directory_name>"
    exit 1
fi

BASE_DIR=$1

# Create base directory
mkdir -p "apps/$BASE_DIR"
cd "apps/$BASE_DIR"

# domain
mkdir -p domain/model
mkdir -p domain/repository
mkdir -p domain/service
touch domain/.gitkeep
touch domain/model/.gitkeep
touch domain/repository/.gitkeep
touch domain/service/.gitkeep

# infra
mkdir -p infra/dao/query
mkdir -p infra/db/ddl
mkdir -p infra/db/migrations
mkdir -p infra/db/seed
mkdir -p infra/dto
touch infra/.gitkeep
touch infra/dao/.gitkeep
touch infra/dao/query/.gitkeep
touch infra/db/ddl/.gitkeep
touch infra/db/migrations/.gitkeep
touch infra/db/seed/.gitkeep
touch infra/dto/.gitkeep

# ui
mkdir -p ui/grpc
touch ui/.gitkeep
touch ui/grpc/.gitkeep

# usecase
mkdir -p usecase/query
touch usecase/.gitkeep
touch usecase/query/.gitkeep
