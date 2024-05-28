include .envrc

# ==================================================================================== #
# PROTOBUF
# ==================================================================================== #

# Define the paths
PROTO_DIR := proto
AUTHENTICATION_PROTO := $(PROTO_DIR)/authenticationservice/authentication.proto
RIDER_PROTO := $(PROTO_DIR)/riderservice/rider.proto
SMS_NOTIFICATION_PROTO := $(PROTO_DIR)/smsnotificationservice/smsnotification.proto
GATEWAY_PROTO := $(PROTO_DIR)/gatewayservice/gateway.proto

AUTHENTICATION_OUT_DIR := authenticationservice/genproto
RIDER_OUT_DIR := riderservice/genproto
SMS_NOTIFICATION_OUT_DIR := smsnotificationservice/genproto
GATEWAY_OUT_DIR := gatewayservice/genproto

# Check if output directories exist, if not create them
.PHONY: create_dirs
create_dirs:
	@mkdir -p $(AUTHENTICATION_OUT_DIR)
	@mkdir -p $(RIDER_OUT_DIR)
	@mkdir -p $(SMS_NOTIFICATION_OUT_DIR)
	@mkdir -p $(GATEWAY_OUT_DIR)

# Define the protoc command
PROTOC := protoc
PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GRPC_GO := protoc-gen-go-grpc

# Generate the protobuf files
.PHONY: proto
proto: create_dirs authentication_proto rider_proto sms_notification_proto gateway_proto

authentication_proto: $(AUTHENTICATION_PROTO)
	$(PROTOC) --go_out=. --go-grpc_out=. $(AUTHENTICATION_PROTO)

rider_proto: $(RIDER_PROTO)
	$(PROTOC) --go_out=. --go-grpc_out=. $(RIDER_PROTO)

sms_notification_proto: $(SMS_NOTIFICATION_PROTO)
	$(PROTOC) --go_out=. --go-grpc_out=. $(SMS_NOTIFICATION_PROTO)
	
gateway_proto: $(GATEWAY_PROTO)
	$(PROTOC) --go_out=. --go-grpc_out=. $(GATEWAY_PROTO)
	
# Clean the generated files
.PHONY: clean
clean:
	rm -f $(AUTHENTICATION_OUT_DIR)/*.pb.go
	rm -f $(RIDER_OUT_DIR)/*.pb.go
	rm -f $(SMS_NOTIFICATION_OUT_DIR)/*.pb.go
	rm -f $(GATEWAY_OUT_DIR)/*.pb.go

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
	
# Run the rider service
.PHONY: rider
rider:
	@go run riderservice/cmd/*.go

# Run the authentication service
.PHONY: authentication
authentication:
	@go run authenticationservice/cmd/*.go

# Run the sms notification service
.PHONY: sms
sms:
	@go run smsnotificationservice/cmd/*.go

# Run the gateway service
.PHONY: gateway
gateway:
	@go run gatewayservice/cmd/*.go

# Run Hashicorp Consul locally
.PHONY: consul
consul:
	docker run -d -p 8500:8500 -p 8600:8600/udp --name=ridewise-dev-consul hashicorp/consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0