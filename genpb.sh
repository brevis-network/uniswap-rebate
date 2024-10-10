export MODU=github.com/brevis-network/uniswap-rebate
protoc -I./proto --go_out . --go_opt module=$MODU \
        --go-grpc_out . --go-grpc_opt require_unimplemented_servers=false,module=$MODU \
        --grpc-gateway_out . --grpc-gateway_opt module=$MODU ./proto/webapi.proto