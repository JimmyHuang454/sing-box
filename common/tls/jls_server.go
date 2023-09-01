package tls

import (
	"context"
	"crypto/tls"
	"net"

	JLS "github.com/JimmyHuang454/JLS-go/tls"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
	"github.com/sagernet/sing/common/ntp"
)

var certPem = []byte(`-----BEGIN CERTIFICATE-----
MIIBhTCCASugAwIBAgIQIRi6zePL6mKjOipn+dNuaTAKBggqhkjOPQQDAjASMRAw
DgYDVQQKEwdBY21lIENvMB4XDTE3MTAyMDE5NDMwNloXDTE4MTAyMDE5NDMwNlow
EjEQMA4GA1UEChMHQWNtZSBDbzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD0d
7VNhbWvZLWPuj/RtHFjvtJBEwOkhbN/BnnE8rnZR8+sbwnc/KhCk3FhnpHZnQz7B
5aETbbIgmuvewdjvSBSjYzBhMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggr
BgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdEQQiMCCCDmxvY2FsaG9zdDo1
NDUzgg4xMjcuMC4wLjE6NTQ1MzAKBggqhkjOPQQDAgNIADBFAiEA2zpJEPQyz6/l
Wf86aX6PepsntZv2GYlA5UpabfT2EZICICpJ5h/iI+i341gBmLiAFQOyTDT+/wQc
6MF9+Yw1Yy0t
-----END CERTIFICATE-----`)

var keyPem = []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIrYSSNQFaA2Hwf1duRSxKtLYX5CB04fSeQ6tF1aY/PuoAoGCCqGSM49
AwEHoUQDQgAEPR3tU2Fta9ktY+6P9G0cWO+0kETA6SFs38GecTyudlHz6xvCdz8q
EKTcWGekdmdDPsHloRNtsiCa697B2O9IFA==
-----END EC PRIVATE KEY-----`)

type JLSServerConfig struct {
	config   *JLS.Config
	isCompat bool
}

// NextProtos implements tls.ServerConfig.
func (c *JLSServerConfig) NextProtos() []string {
	return c.config.NextProtos
}

// SetNextProtos implements tls.ServerConfig.
func (c *JLSServerConfig) SetNextProtos(nextProto []string) {
	c.config.NextProtos = nextProto
}

func (c *JLSServerConfig) ServerName() string {
	return c.config.ServerName
}

func (c *JLSServerConfig) SetServerName(serverName string) {
	c.config.ServerName = serverName
}

func (c *JLSServerConfig) Config() (*STDConfig, error) {
	return nil, E.New("unsupported usage for JLS")
}

func (c *JLSServerConfig) Client(conn net.Conn) (Conn, error) {
	return &JLSConnWrapper{JLS.Client(conn, c.config)}, nil
}

func (c *JLSServerConfig) Server(conn net.Conn) (Conn, error) {
	return &JLSConnWrapper{JLS.Server(conn, c.config)}, nil
}

func (c *JLSServerConfig) Clone() Config {
	return &JLSServerConfig{
		config: c.config.Clone(),
	}
}

func (c *JLSServerConfig) Start() error {
	return nil
}

func (c *JLSServerConfig) Close() error {
	return nil
}

func NewJLSServer(ctx context.Context, logger log.Logger, options option.InboundTLSOptions) (ServerConfig, error) {
	tlsConfig := &JLS.Config{}
	tlsConfig.Time = ntp.TimeFuncFromContext(ctx)

	if options.ServerName == "" {
		return nil, E.New("fallback website is needed.")
	}
	tlsConfig.ServerName = options.ServerName

	if len(options.ALPN) > 0 {
		tlsConfig.NextProtos = append(options.ALPN, tlsConfig.NextProtos...)
	}

	if options.CipherSuites != nil {
	find:
		for _, cipherSuite := range options.CipherSuites {
			for _, tlsCipherSuite := range tls.CipherSuites() {
				if cipherSuite == tlsCipherSuite.Name {
					tlsConfig.CipherSuites = append(tlsConfig.CipherSuites, tlsCipherSuite.ID)
					continue find
				}
			}
			return nil, E.New("unknown cipher_suite: ", cipherSuite)
		}
	}

	cert, _ := JLS.X509KeyPair(certPem, keyPem)
	tlsConfig.Certificates = []JLS.Certificate{cert}
	tlsConfig.JLSPWD = []byte(options.JLS.Password)
	tlsConfig.JLSIV = []byte(options.JLS.IV)
	tlsConfig.UseJLS = true

	return &JLSServerConfig{
		config:   tlsConfig,
		isCompat: false,
	}, nil
}

type JLSConnWrapper struct {
	*JLS.Conn
}

func (c *JLSConnWrapper) ConnectionState() tls.ConnectionState {
	state := c.Conn.ConnectionState()
	return tls.ConnectionState{
		Version:                     state.Version,
		HandshakeComplete:           state.HandshakeComplete,
		DidResume:                   state.DidResume,
		CipherSuite:                 state.CipherSuite,
		NegotiatedProtocol:          state.NegotiatedProtocol,
		NegotiatedProtocolIsMutual:  state.NegotiatedProtocolIsMutual,
		ServerName:                  state.ServerName,
		PeerCertificates:            state.PeerCertificates,
		VerifiedChains:              state.VerifiedChains,
		SignedCertificateTimestamps: state.SignedCertificateTimestamps,
		OCSPResponse:                state.OCSPResponse,
		TLSUnique:                   state.TLSUnique,
	}
}
