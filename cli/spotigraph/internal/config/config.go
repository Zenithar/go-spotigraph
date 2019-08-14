package config

import "go.zenithar.org/pkg/platform"

// Configuration contains spotigraph settings
type Configuration struct {
	Debug struct {
		Enable bool `toml:"enable" default:"false" comment:"allow debug mode"`
	} `toml:"Debug" comment:"###############################\n Debug with gops \n##############################"`

	Instrumentation platform.InstrumentationConfig `toml:"Instrumentation" comment:"###############################\n Instrumentation \n##############################"`

	Core struct {
		Mode  string `toml:"mode" default:"local" comment:"Use remote or local as backend"`
		Local struct {
			AutoMigrate bool   `toml:"-" default:"false"`
			Type        string `toml:"type" default:"postgresql" comment:"Database connector to use: rethinkdb."`
			Hosts       string `toml:"hosts" default:"postgresql://spotigraph:changeme@localhost:5432/spotigraph?driver=pgx" comment:"Database hosts (comma separated)"`
			Database    string `toml:"database" default:"spotigraph" comment:"Database namespace"`
			Username    string `toml:"username" default:"spotigraph" comment:"Database connection username"`
			Password    string `toml:"password" default:"changeme" comment:"Database connection password"`
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
		GraphQL struct {
			Network string `toml:"network" default:"tcp" comment:"Network class used for listen (tcp, tcp4, tcp6, unixsocket)"`
			Listen  string `toml:"listen" default:":8080" comment:"Listen address for HTTP server"`
			UseTLS  bool   `toml:"useTLS" default:"false" comment:"Enable TLS listener"`
			TLS     struct {
				CertificatePath              string `toml:"certificatePath" default:"" comment:"Certificate path"`
				PrivateKeyPath               string `toml:"privateKeyPath" default:"" comment:"Private Key path"`
				CACertificatePath            string `toml:"caCertificatePath" default:"" comment:"CA Certificate Path"`
				ClientAuthenticationRequired bool   `toml:"clientAuthenticationRequired" default:"false" comment:"Force client authentication"`
			} `toml:"TLS" comment:"TLS Socket settings"`
		} `toml:"GraphQL" comment:"###############################\n GraphQL Settings \n##############################"`
	}
}
