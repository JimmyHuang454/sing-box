cd test
go mod tidy
go test -v -tags "with_gvisor,with_quic,with_wireguard,with_grpc,with_ech,with_utls,with_shadowsocksr" . -run TestTrojanPlainSelf
:: go test -v -tags "with_gvisor,with_quic,with_wireguard,with_grpc,with_ech,with_utls,with_shadowsocksr" .
