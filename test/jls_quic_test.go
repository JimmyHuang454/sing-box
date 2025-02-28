package main

import (
	"net/netip"
	"testing"

	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/option"
)

func TestJLSQUIC(t *testing.T) {
	_, certPem, keyPem := createSelfSignedCertificate(t, "example.org")

	startInstance(t, option.Options{
		Inbounds: []option.Inbound{
			{
				Type: C.TypeMixed,
				Tag:  "mixed-in",
				MixedOptions: option.HTTPMixedInboundOptions{
					ListenOptions: option.ListenOptions{
						Listen:     option.NewListenAddress(netip.IPv4Unspecified()),
						ListenPort: clientPort,
					},
				},
			},
			{
				Type: C.TypeTrojan,
				TrojanOptions: option.TrojanInboundOptions{
					ListenOptions: option.ListenOptions{
						Listen:     option.NewListenAddress(netip.IPv4Unspecified()),
						ListenPort: serverPort,
					},
					Users: []option.TrojanUser{
						{
							Name:     "sekai",
							Password: "password",
						},
					},
					TLS: &option.InboundTLSOptions{
						Enabled:         true,
						ServerName:      "www.jsdelivr.com",
						JLS:             &option.JLSOptions{Enabled: true, Password: "123", IV: "456"},
						CertificatePath: certPem,
						KeyPath:         keyPem,
					},
					Transport: &option.V2RayTransportOptions{
						Type: C.V2RayTransportTypeQUIC,
					},
				},
			},
		},
		Outbounds: []option.Outbound{
			{
				Type: C.TypeDirect,
			},
			{
				Type: C.TypeTrojan,
				Tag:  "trojan-out",
				TrojanOptions: option.TrojanOutboundOptions{
					ServerOptions: option.ServerOptions{
						Server:     "127.0.0.1",
						ServerPort: serverPort,
					},
					Password: "password",
					TLS: &option.OutboundTLSOptions{
						Enabled:         true,
						ServerName:      "www.jsdelivr.com",
						JLS:             &option.JLSOptions{Enabled: true, Password: "123", IV: "456"},
						CertificatePath: certPem,
					},
					Transport: &option.V2RayTransportOptions{
						Type: C.V2RayTransportTypeQUIC,
					},
				},
			},
		},
		Route: &option.RouteOptions{
			Rules: []option.Rule{
				{
					DefaultOptions: option.DefaultRule{
						Inbound:  []string{"mixed-in"},
						Outbound: "trojan-out",
					},
				},
			},
		},
	})
	testSuit(t, clientPort, testPort)
}
