module go.zenithar.org/spotigraph

go 1.12

replace github.com/opencensus-integrations/gomongowrapper => github.com/Zenithar/gomongowrapper v0.0.2

require (
	github.com/99designs/gqlgen v0.8.3
	github.com/DataDog/zstd v1.4.0 // indirect
	github.com/Masterminds/squirrel v1.1.0
	github.com/Shopify/sarama v1.22.0 // indirect
	github.com/StackExchange/wmi v0.0.0-20181212234831-e0a55b97c705 // indirect
	github.com/agnivade/levenshtein v1.0.2 // indirect
	github.com/asaskevich/govalidator v0.0.0-20180720115003-f9ffefc3facf // indirect
	github.com/beorn7/perks v1.0.0 // indirect
	github.com/cloudflare/tableflip v0.0.0-20190329062924-8392f1641731
	github.com/common-nighthawk/go-figure v0.0.0-20180619031829-18b2b544842c
	github.com/coreos/etcd v3.3.12+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/dchest/uniuri v0.0.0-20160212164326-8902c56451e9
	github.com/fatih/color v1.7.0
	github.com/fatih/structs v1.1.0 // indirect
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/gobuffalo/depgen v0.1.1 // indirect
	github.com/gobuffalo/genny v0.1.1 // indirect
	github.com/gobuffalo/gogen v0.1.1 // indirect
	github.com/gobuffalo/packr v1.25.0
	github.com/gogo/protobuf v1.2.1
	github.com/golang/mock v1.2.0
	github.com/golang/protobuf v1.3.1
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.2.0
	github.com/google/gops v0.3.6
	github.com/google/pprof v0.0.0-20190404155422-f8f10df84213 // indirect
	github.com/google/wire v0.2.1
	github.com/gorilla/mux v1.7.1 // indirect
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/gosimple/slug v1.5.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/hokaccha/go-prettyjson v0.0.0-20180920040306-f579f869bbfe
	github.com/jmoiron/sqlx v1.2.0
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/kisielk/errcheck v1.2.0 // indirect
	github.com/kr/pty v1.1.4 // indirect
	github.com/lyft/protoc-gen-validate v0.0.14
	github.com/magefile/mage v1.8.0
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/mcuadros/go-defaults v1.1.0
	github.com/oklog/run v1.0.0
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0
	github.com/opencensus-integrations/gomongowrapper v0.0.1
	github.com/pelletier/go-toml v1.3.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90 // indirect
	github.com/prometheus/common v0.3.0 // indirect
	github.com/prometheus/procfs v0.0.0-20190416084830-8368d24ba045 // indirect
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/rogpeppe/fastuuid v1.0.0 // indirect
	github.com/rubenv/sql-migrate v0.0.0-20190327083759-54bad0a9b051
	github.com/shirou/gopsutil v2.18.12+incompatible // indirect
	github.com/shurcooL/httpfs v0.0.0-20181222201310-74dc9339e414 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd // indirect
	github.com/sirupsen/logrus v1.4.1
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.3.2
	github.com/vektah/gqlparser v1.1.2
	github.com/xlab/treeprint v0.0.0-20181112141820-a009c3971eca // indirect
	github.com/ziutek/mymysql v1.5.4 // indirect
	go.opencensus.io v0.20.2
	go.uber.org/zap v1.9.1
	go.zenithar.org/pkg/cache v0.0.1
	go.zenithar.org/pkg/db v0.0.2
	go.zenithar.org/pkg/db/adapter/mongodb v0.0.1
	go.zenithar.org/pkg/db/adapter/postgresql v0.0.1
	go.zenithar.org/pkg/db/adapter/rethinkdb v0.0.1
	go.zenithar.org/pkg/flags v0.0.1
	go.zenithar.org/pkg/log v0.0.1
	go.zenithar.org/pkg/platform v0.0.1
	go.zenithar.org/pkg/testing v0.0.12
	go.zenithar.org/pkg/tlsconfig v0.0.1
	go.zenithar.org/pkg/web v0.0.1
	golang.org/x/crypto v0.0.0-20190422183909-d864b10871cd
	golang.org/x/exp v0.0.0-20190422150234-47ea93f3503f // indirect
	golang.org/x/image v0.0.0-20190417020941-4e30a6eb7d9a // indirect
	golang.org/x/lint v0.0.0-20190409202823-959b441ac422 // indirect
	golang.org/x/mobile v0.0.0-20190415191353-3e0bab5405d6 // indirect
	golang.org/x/net v0.0.0-20190420063019-afa5a82059c6
	golang.org/x/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 // indirect
	golang.org/x/sys v0.0.0-20190422165155-953cdadca894 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	golang.org/x/tools v0.0.0-20190422233926-fe54fb35175b // indirect
	google.golang.org/grpc v1.20.1
	gopkg.in/gorp.v1 v1.7.2 // indirect
	gopkg.in/rethinkdb/rethinkdb-go.v5 v5.0.1
	honnef.co/go/tools v0.0.0-20190418001031-e561f6794a2a // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
	rsc.io/goversion v1.2.0 // indirect
	sourcegraph.com/sourcegraph/appdash v0.0.0-20190107175209-d9ea5c54f7dc // indirect
)
