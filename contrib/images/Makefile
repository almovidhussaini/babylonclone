RELAYER_TAG := $(shell grep '^ENV RELAYER_TAG' cosmos-relayer/Dockerfile | cut -f3 -d\ )
ybtc_FULL_PATH := $(shell git rev-parse --show-toplevel)
ybtc_VERSION_BEFORE_UPGRADE ?= v0.9.4

all: ybtcd cosmos-relayer

ybtcd: ybtcd-rmi
	docker build --tag ybtclabs-io/ybtcd -f ybtcd/Dockerfile ${ybtc_FULL_PATH}

ybtcd-e2e:
	docker build --tag ybtclabs-io/ybtcd -f ybtcd/Dockerfile ${ybtc_FULL_PATH} \
		--build-arg ybtc_BUILD_OPTIONS="testnet"

ybtcd-rmi:
	docker rmi ybtclabs-io/ybtcd --force 2>/dev/null; true

e2e-init-chain-rmi:
	docker rmi ybtclabs-io/ybtcd-e2e-init-chain --force 2>/dev/null; true

e2e-init-chain: e2e-init-chain-rmi
	@DOCKER_BUILDKIT=1 docker build -t ybtclabs-io/ybtcd-e2e-init-chain --build-arg E2E_SCRIPT_NAME=chain --platform=linux/x86_64 \
		-f e2e-initialization/init.Dockerfile --build-arg VERSION="${ybtc_VERSION_BEFORE_UPGRADE}" ${ybtc_FULL_PATH}

cosmos-relayer: cosmos-relayer-rmi
	docker build --tag ybtclabs-io/cosmos-relayer:${RELAYER_TAG} -f cosmos-relayer/Dockerfile \
				${ybtc_FULL_PATH}/contrib/images/cosmos-relayer
	docker tag ybtclabs-io/cosmos-relayer:${RELAYER_TAG} ybtclabs-io/cosmos-relayer:latest

cosmos-relayer-rmi:
	docker rmi ybtclabs-io/cosmos-relayer 2>/dev/null; true

.PHONY: all ybtcd cosmos-relayer e2e-init-chain ybtcd-rmi cosmos-relayer-rmi 
