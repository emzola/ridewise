# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## protoc: generate protoc
.PHONY: protoc
protoc:
	protoc -I=api --go_out=. --go-grpc_out=. ridewise.proto