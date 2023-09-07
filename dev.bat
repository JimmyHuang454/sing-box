go build -tags "with_gvisor,with_dhcp,with_wireguard,with_utls,with_reality_server,with_clash_api,with_grpc,with_quic" ./cmd/sing-box
:: sing-box.exe run --config ./jls_config.json
