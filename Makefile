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
# RUN MICROSERVICES
# ==================================================================================== #
	
# Run the rider service
.PHONY: run_rider
run_rider:
	@go run riderservice/cmd/*.go

# Run the authentication service
.PHONY: run_authentication
run_authentication:
	@go run authenticationservice/cmd/*.go

# Run the sms notification service
.PHONY: run_sms
run_sms:
	@go run smsnotificationservice/cmd/*.go

# Run the gateway service
.PHONY: run_gateway
run_gateway:
	@go run gatewayservice/cmd/*.go