protoc -I=. --go_out=paths=source_relative:gen/go trace.proto
#protoc -I=. --go-grpc_out=paths=source_relative:gen/go trace.proto
protoc -I=. --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:gen/go trace.proto
# generate RESTful JSON APIs
protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=trace.yaml:gen/go trace.proto