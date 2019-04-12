package config

import (
	"go.zenithar.org/pkg/platform/jaeger"
	"go.zenithar.org/pkg/platform/prometheus"
)

// Configuration contains spotigraph settings
type Configuration struct {
	Debug struct {
		Enable    bool   `toml:"enable" default:"false" comment:"allow debugging with gops"`
		RemoteURL string `toml:"remoteDebugURL" comment:"start a gops agent on specified URL. Ex: localhost:9999"`
	} `toml:"Debug" comment:"###############################\n Debug with gops \n##############################"`

	Instrumentation struct {
		Network string `toml:"network" default:"tcp" comment:"Network class used for listen (tcp, tcp4, tcp6, unixsocket)"`
		Listen  string `toml:"listen" default:":5556" comment:"Listen address for instrumentation server"`
		Logs    struct {
			Level     string `toml:"level" default:"warn" comment:"Log level: debug, info, warn, error, dpanic, panic, and fatal"`
			SentryDSN string `toml:"sentryDSN" comment:"Sentry DSN"`
		} `toml:"Logs" comment:"###############################\n Logs Settings \n##############################"`
		Prometheus struct {
			Enabled bool              `toml:"enabled" default:"true" comment:"Enable metric instrumentation"`
			Config  prometheus.Config `toml:"Config" comment:"Prometheus settings"`
		} `toml:"Prometheus" comment:"###############################\n Prometheus exporter \n##############################"`
		Jaeger struct {
			Enabled bool          `toml:"enabled" default:"true" comment:"Enable trace instrumentation"`
			Config  jaeger.Config `toml:"Config" comment:"Jaeger settings"`
		} `toml:"Jaeger" comment:"###############################\n Jaeger exporter \n##############################"`
	} `toml:"Instrumentation" comment:"###############################\n Instrumentation \n##############################"`

	Core struct {
		Mode  string `toml:"mode" default:"local" comment:"Use remote or local as backend"`
		Local struct {
			AutoMigrate bool   `toml:"-" default:"false"`
			Type        string `toml:"type" default:"rethinkdb" comment:"Database connector to use: rethinkdb."`
			Hosts       string `toml:"hosts" default:"127.0.0.1:28015" comment:"Database hosts (comma separated)"`
			Database    string `toml:"database" default:"spotigraph" comment:"Database namespace"`
			Username    string `toml:"username" default:"" comment:"Database connection username"`
			Password    string `toml:"password" default:"" comment:"Database connection password"`
		} `toml:"Local" comment:"###############################\n Local Settings \n##############################"`
		Remote struct {
			Address string `toml:"address" default:"" comment:"Remote gRPC service address"`
			UseTLS  bool   `toml:"useTLS" default:"false" comment:"Enable TLS listener"`
			TLS     struct {
				CertificatePath              string `toml:"certificatePath" default:"" comment:"Certificate path"`
				PrivateKeyPath               string `toml:"privateKeyPath" default:"" comment:"Private Key path"`
				CACertificatePath            string `toml:"caCertificatePath" default:"" comment:"CA Certificate Path"`
				ClientAuthenticationRequired bool   `toml:"clientAuthenticationRequired" default:"false" comment:"Force client authentication"`
			} `toml:"TLS" comment:"TLS Socket settings"`
		} `toml:"Remote" comment:"###############################\n Remote Settings \n##############################"`
	} `toml:"Core" comment:"###############################\n Core \n##############################"`

	Server struct {
		GRPC struct {
			Network string `toml:"network" default:"tcp" comment:"Network class used for listen (tcp, tcp4, tcp6, unixsocket)"`
			Listen  string `toml:"listen" default:":5555" comment:"Listen address for gRPC server"`
			UseTLS  bool   `toml:"useTLS" default:"false" comment:"Enable TLS listener"`
			TLS     struct {
				CertificatePath              string `toml:"certificatePath" default:"" comment:"Certificate path"`
				PrivateKeyPath               string `toml:"privateKeyPath" default:"" comment:"Private Key path"`
				CACertificatePath            string `toml:"caCertificatePath" default:"" comment:"CA Certificate Path"`
				ClientAuthenticationRequired bool   `toml:"clientAuthenticationRequired" default:"false" comment:"Force client authentication"`
			} `toml:"TLS" comment:"TLS Socket settings"`
		} `toml:"GRPC" comment:"###############################\n gRPC Settings \n##############################"`
		HTTP struct {
			Network string `toml:"network" default:"tcp" comment:"Network class used for listen (tcp, tcp4, tcp6, unixsocket)"`
			Listen  string `toml:"listen" default:":8080" comment:"Listen address for HTTP server"`
			UseTLS  bool   `toml:"useTLS" default:"false" comment:"Enable TLS listener"`
			TLS     struct {
				CertificatePath              string `toml:"certificatePath" default:"" comment:"Certificate path"`
				PrivateKeyPath               string `toml:"privateKeyPath" default:"" comment:"Private Key path"`
				CACertificatePath            string `toml:"caCertificatePath" default:"" comment:"CA Certificate Path"`
				ClientAuthenticationRequired bool   `toml:"clientAuthenticationRequired" default:"false" comment:"Force client authentication"`
			} `toml:"TLS" comment:"TLS Socket settings"`
		} `toml:"HTTP" comment:"###############################\n HTTP Settings \n##############################"`
	}
}
