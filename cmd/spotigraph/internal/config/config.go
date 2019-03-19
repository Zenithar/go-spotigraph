package config

// Configuration contains spotigraph settings
type Configuration struct {
	Debug struct {
		Enable    bool   `toml:"enable" default:"false" comment:"allow debugging with gops"`
		RemoteURL string `toml:"remoteDebugURL" comment:"start a gops agent on specified URL. Ex: localhost:9999"`
	} `toml:"Debug" comment:"###############################\n Debug with gops \n##############################"`

	Observability struct {
		Logs struct {
			Level     string `toml:"level" default:"warn" comment:"Log level: debug, info, warn, error, dpanic, panic, and fatal"`
			SentryDSN string `toml:"sentryDSN" comment:"Sentry DSN"`
		} `toml:"Logs" comment:"###############################\n Logs Settings \n##############################"`
		Metrics struct {
			Enable bool `toml:"enable" default:"true" comment:"Enable metric scrapper endpoint for Prometheus"`
		} `toml:"Metrics" comment:"###############################\n Metric exporter \n##############################"`
		Traces struct {
			TracerURL string `toml:"tracerURL" comment:"OpenTracing compatible tracer"`
		} `toml:"Traces" comment:"###############################\n Trace exporter \n##############################"`
	} `toml:"Observability" comment:"###############################\n Observability \n##############################"`

	Server struct {
		Database struct {
			AutoMigrate bool   `toml:"-" default:"false"`
			Type        string `toml:"type" default:"rethinkdb" comment:"Database connector to use: rethinkdb."`
			Hosts       string `toml:"hosts" default:"127.0.0.1:28015" comment:"Database hosts (comma separated)"`
			Database    string `toml:"database" default:"spotigraph" comment:"Database namespace"`
			Username    string `toml:"username" default:"" comment:"Database connection username"`
			Password    string `toml:"password" default:"" comment:"Database connection password"`
		} `toml:"Database" comment:"###############################\n Database Settings \n##############################"`
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
	}
}
