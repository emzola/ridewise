# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## protoc: generate protoc
.PHONY: protoc
protoc:
	protoc -I=api --go_out=. --go-grpc_out=. ridewise.proto
	
## run/rider: run the rider service
.PHONY: run/rider
run/rider:
	@go run riderservice/cmd/*.go