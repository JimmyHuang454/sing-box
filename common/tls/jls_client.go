package tls

import (
	"net"

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
	tlsConfig := &JLS.Config{}
	tlsConfig.Time = router.TimeFunc()

	if options.ServerName == "" {
		return nil, E.New("fallback website is needed.")
	}
	tlsConfig.ServerName = options.ServerName

	if len(options.ALPN) > 0 {
		tlsConfig.NextProtos = options.ALPN
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
	tlsConfig.UseJLS = true

	return &JLSClientConfig{tlsConfig}, nil
}
