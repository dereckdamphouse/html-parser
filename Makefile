.PHONY: init test build deploy build-HTMLParserFunction

REGION=eu-west-2
AWS_PROFILE=dereck-damphouse

ifeq ($(env), prod)
CONFIG_ENV=prod
else
CONFIG_ENV=dev
endif

init:
	go mod download
	go mod tidy

test:
	@echo "running tests with race flag..."
	@go test ./... -race -coverprofile cover.out
	@go tool cover -func cover.out | grep total:
	@rm cover.out

build: test
	@echo "building ${CONFIG_ENV} handler for AWS Lambda"
	sam validate \
		--profile ${AWS_PROFILE} \
		--region ${REGION}
	sam build \
		--profile ${AWS_PROFILE} \
		--region ${REGION} \

deploy: build
	@echo "deploying ${CONFIG_ENV} infrastructure and code"
	sam deploy \
		--config-env ${CONFIG_ENV} \
		--profile ${AWS_PROFILE} \
		--region ${REGION} \
		--parameter-overrides ParameterKey=Environment,ParameterValue=${CONFIG_ENV}

build-HTMLParserFunction:
	$(MAKE) -C cmd build