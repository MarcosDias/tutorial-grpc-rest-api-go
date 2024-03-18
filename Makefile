.PHONY: proto
proto:
	buf generate --template proto/buf.gen.yaml proto

.PHONY: runapi
runapi:
	go run ./cmd/api

.PHONY: rungateway
rungateway:
	go run ./cmd/gateway

.PHONY: runcomplete
runcomplete:
	go run ./cmd/complete