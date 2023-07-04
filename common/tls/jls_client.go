package tls

import (
	"crypto/x509"
	"net"
	"net/netip"
	"os"

	JLS "github.com/JimmyHuang454/JLS-go/tls"
	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

type JLSClientConfig struct {
	config *JLS.Config
}

func (s *JLSClientConfig) ServerName() string {
	return s.config.ServerName
}

func (s *JLSClientConfig) SetServerName(serverName string) {
	s.config.ServerName = serverName
}

func (s *JLSClientConfig) NextProtos() []string {
	return s.config.NextProtos
}

func (s *JLSClientConfig) SetNextProtos(nextProto []string) {
	s.config.NextProtos = nextProto
}

func (s *JLSClientConfig) Config() (*STDConfig, error) {
	return nil, E.New("unsupported usage for JLS")
}

func (s *JLSClientConfig) Client(conn net.Conn) (Conn, error) {
	return &JLSConnWrapper{JLS.Client(conn, s.config)}, nil
}

func (s *JLSClientConfig) Clone() Config {
	return &JLSClientConfig{s.config.Clone()}
}

func NewJLSlient(router adapter.Router, serverAddress string, options option.OutboundTLSOptions) (Config, error) {
	var serverName string
	if options.ServerName != "" {
		serverName = options.ServerName
	} else if serverAddress != "" {
		if _, err := netip.ParseAddr(serverName); err != nil {
			serverName = serverAddress
		}
	}
	if serverName == "" && !options.Insecure {
		return nil, E.New("missing server_name or insecure=true")
	}

	var tlsConfig JLS.Config
	tlsConfig.Time = router.TimeFunc()
	if options.DisableSNI {
		tlsConfig.ServerName = "127.0.0.1"
	} else {
		tlsConfig.ServerName = serverName
	}
	if options.Insecure {
		tlsConfig.InsecureSkipVerify = options.Insecure
	} else if options.DisableSNI {
		tlsConfig.InsecureSkipVerify = true
		tlsConfig.VerifyConnection = func(state JLS.ConnectionState) error {
			verifyOptions := x509.VerifyOptions{
				DNSName:       serverName,
				Intermediates: x509.NewCertPool(),
			}
			for _, cert := range state.PeerCertificates[1:] {
				verifyOptions.Intermediates.AddCert(cert)
			}
			_, err := state.PeerCertificates[0].Verify(verifyOptions)
			return err
		}
	}
	if len(options.ALPN) > 0 {
		tlsConfig.NextProtos = options.ALPN
	}
	if options.MinVersion != "" {
		minVersion, err := ParseTLSVersion(options.MinVersion)
		if err != nil {
			return nil, E.Cause(err, "parse min_version")
		}
		tlsConfig.MinVersion = minVersion
	}
	if options.MaxVersion != "" {
		maxVersion, err := ParseTLSVersion(options.MaxVersion)
		if err != nil {
			return nil, E.Cause(err, "parse max_version")
		}
		tlsConfig.MaxVersion = maxVersion
	}
	if options.CipherSuites != nil {
	find:
		for _, cipherSuite := range options.CipherSuites {
			for _, tlsCipherSuite := range JLS.CipherSuites() {
				if cipherSuite == tlsCipherSuite.Name {
					tlsConfig.CipherSuites = append(tlsConfig.CipherSuites, tlsCipherSuite.ID)
					continue find
				}
			}
			return nil, E.New("unknown cipher_suite: ", cipherSuite)
		}
	}

	tlsConfig.JLSPWD = []byte(options.JLS.Password)
	tlsConfig.JLSIV = []byte(options.JLS.IV)

	return &JLSClientConfig{&tlsConfig}, nil
}
