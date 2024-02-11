.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make install     Install dependencies"
	@echo "  make dev         Start development server"
	@echo "  make build       Build for production"
	@echo "  make preview     Start preview server"
	@echo "  make lint        Lint code"

.PHONY: install
install:
	@echo "Installing..."
	npm install

.PHONY: dev
dev:
	@echo "Starting development server..."
	npm run dev

.PHONY: build
build:
	@echo "Building..."
	npm run build

.PHONY: preview
preview:
	@echo "Starting preview server..."
	npm run start

.PHONY: lint
lint:
	@echo "Linting..."
	npm run lint
