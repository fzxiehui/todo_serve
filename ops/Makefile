.PHONY: run clean dev default

default: run

DEV_CONFIG_PATH = "config/dev"
RELEASE_CONFIG_PATH = "config/release"

RUN_DIR = "bin"
BUILD_DIR = "build"

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  help:     Show this help"
	@echo "  run:      Run the application"
	@echo "  clean:    Clean the project"
	@echo "  dev:      Run the application in development mode"
	@echo "  default:  Run the application"
	@echo "  release:  Run the application in release mode"
	@echo "  build:    Build the application"

run:
	@echo 'Running...'
	@docker-compose -f ${RUN_DIR}/docker-compose.yml -p ${RUN_DIR} up -d

stop:
	@echo 'Stopping...'
	@docker-compose -f ${RUN_DIR}/docker-compose.yml -p ${RUN_DIR} down

clean:
	@echo 'Cleaning...'
	@test ! -d ${RUN_DIR} || sudo rm -rf ${RUN_DIR}
	@test ! -d ${BUILD_DIR} || sudo rm -rf ${BUILD_DIR}

clean-bin:
	@echo 'Cleaning...'
	@test ! -d ${RUN_DIR} || sudo rm -rf ${RUN_DIR}

clean-build:
	@echo 'Cleaning...'
	@test ! -d ${BUILD_DIR} || sudo rm -rf ${BUILD_DIR}

dev:
	@echo 'Running in dev mode...'
	@mkdir -p ${RUN_DIR}
	@cp -r ${DEV_CONFIG_PATH}/* ${RUN_DIR}

release:
	@echo 'Running in release mode...'
	@mkdir -p ${RUN_DIR}
	@cp -r ${RELEASE_CONFIG_PATH}/* ${RUN_DIR}

build:
	@echo 'Building...'
	@echo '@TODO'

