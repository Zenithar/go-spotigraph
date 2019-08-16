module go.zenithar.org/spotigraph

go 1.12

replace github.com/opencensus-integrations/gomongowrapper => github.com/Zenithar/gomongowrapper v0.0.2

replace go.mongodb.org/mongo-driver => go.mongodb.org/mongo-driver v1.0.1-0.20190812160042-74cffef35f2e

require (
	github.com/99designs/gqlgen v0.9.3
	github.com/Masterminds/squirrel v1.1.0
	github.com/bxcodec/faker v2.0.1+incompatible
	github.com/cloudflare/tableflip v1.0.0
	github.com/common-nighthawk/go-figure v0.0.0-20190529165535-67e0ed34491a
	github.com/dchest/uniuri v0.0.0-20160212164326-8902c56451e9
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/fatih/color v1.7.0
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/gobuffalo/packr v1.30.1
	github.com/gogo/protobuf v1.2.1
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.3.1
	github.com/google/wire v0.3.0
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.6.2 // indirect
	github.com/gosimple/slug v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/hokaccha/go-prettyjson v0.0.0-20180920040306-f579f869bbfe
	github.com/jackc/pgx v3.5.0+incompatible
	github.com/jmoiron/sqlx v1.2.0
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.2.0
	github.com/magefile/mage v1.8.0
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/oklog/run v1.0.0
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0
	github.com/pkg/errors v0.8.1
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/rubenv/sql-migrate v0.0.0-20190717103323-87ce952f7079
	github.com/sony/gobreaker v0.4.1
	github.com/spf13/afero v1.2.1 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/vektah/gqlparser v1.1.2
	github.com/ziutek/mymysql v1.5.4 // indirect
	go.opencensus.io v0.22.0
	go.uber.org/zap v1.10.0
	go.zenithar.org/pkg/cache v0.0.5
	go.zenithar.org/pkg/config v0.0.6
	go.zenithar.org/pkg/db v0.0.3
	go.zenithar.org/pkg/db/adapter/mongodb v0.0.12 // indirect
	go.zenithar.org/pkg/db/adapter/postgresql v0.0.6
	go.zenithar.org/pkg/db/adapter/rethinkdb v0.0.3 // indirect
	go.zenithar.org/pkg/flags v0.0.2
	go.zenithar.org/pkg/log v0.2.0
	go.zenithar.org/pkg/platform v0.1.4
	go.zenithar.org/pkg/testing v0.0.13
	go.zenithar.org/pkg/tlsconfig v0.0.1
	go.zenithar.org/pkg/web v0.0.2
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/grpc v1.22.1
	gopkg.in/gorp.v1 v1.7.2 // indirect
)
