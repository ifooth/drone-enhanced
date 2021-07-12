.PHONY: build
build:
	@echo ">> installing promu"
	GO111MODULE=on go install github.com/prometheus/promu
	@echo ">> rebuilding binaries using promu"
	GO111MODULE=on promu build --prefix $(PREFIX)

.PHONY: test
test:
	@echo ">> test not support"
