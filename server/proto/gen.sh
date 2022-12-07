protoc -I=. --go_out=paths=source_relative:gen/go trace.proto
protoc -I=. --go-grpc_out=paths=source_relative:gen/go trace.proto
protoc -I=. --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:gen/go trace.proto