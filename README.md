# Golang Migration Tool

Go Migration Tool is a command-line utility for managing database migrations. It supports running SQL migration scripts to apply or rollback changes to databases.

## Features

- Connect to a database and run migrations
- Apply (`up`) or rollback (`down`) SQL migration scripts
- Handle single SQL files or directories containing multiple SQL files
- Supports **PostgreSQL, MySQL, and SQLite3** (currently)

## Prerequisites

- Go **1.23.6** or higher
- PostgreSQL, MySQL, or SQLite3 (depending on your target database)

## Installation

### Using "go install"

`go install github.com/ayayaakasvin/gomig@latest`
