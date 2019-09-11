module go.zenithar.org/spotigraph

go 1.13

replace (
	github.com/envoyproxy/protoc-gen-validate => github.com/Zenithar/protoc-gen-validate v0.2.0-java.0.20190808132234-82cadd0ebcfe
	github.com/go-critic/go-critic => github.com/go-critic/go-critic v0.3.5-0.20190904082202-d79a9f0c64db
	github.com/golangci/errcheck => github.com/golangci/errcheck v0.0.0-20181223084120-ef45e06d44b6
	github.com/golangci/go-tools => github.com/golangci/go-tools v0.0.0-20190124090046-35a9f45a5db0
	github.com/golangci/gofmt => github.com/golangci/gofmt v0.0.0-20181222123516-0b8337e80d98
	github.com/golangci/gosec => github.com/golangci/gosec v0.0.0-20180901114220-8afd9cbb6cfb
	github.com/golangci/ineffassign => github.com/golangci/ineffassign v0.0.0-20180808204949-2ee8f2867dde
	github.com/golangci/lint-1 => github.com/golangci/lint-1 v0.0.0-20181222135242-d2cdd8c08219
	github.com/opencensus-integrations/gomongowrapper => github.com/Zenithar/gomongowrapper v0.0.2
	github.com/timakin/bodyclose => github.com/timakin/bodyclose v0.0.0-20190721030226-87058b9bfcec
	go.mongodb.org/mongo-driver => go.mongodb.org/mongo-driver v1.0.1-0.20190812160042-74cffef35f2e
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190905072037-92dd089d5514
	mvdan.cc/unparam => mvdan.cc/unparam v0.0.0-20190720180237-d51796306d8f
)

require (
	github.com/Masterminds/squirrel v1.1.1-0.20190801214710-0f6e36219a8f
	github.com/cloudflare/tableflip v1.0.0
	github.com/common-nighthawk/go-figure v0.0.0-20190529165535-67e0ed34491a
	github.com/dchest/uniuri v0.0.0-20160212164326-8902c56451e9
	github.com/envoyproxy/protoc-gen-validate v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.7.0
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/gobuffalo/packr v1.30.1
	github.com/gogo/protobuf v1.3.0
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/google/go-cmp v0.3.1
	github.com/google/wire v0.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/hokaccha/go-prettyjson v0.0.0-20190818114111-108c894c2c0e
	github.com/jackc/pgx v3.5.0+incompatible
	github.com/jmoiron/sqlx v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.2.0
	github.com/magefile/mage v1.8.0
	github.com/oklog/run v1.0.0
	github.com/onsi/gomega v1.7.0
	github.com/pkg/errors v0.8.1
	github.com/rubenv/sql-migrate v0.0.0-20190902133344-8926f37f0bc1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	go.opencensus.io v0.22.1
	go.uber.org/zap v1.10.0
	go.zenithar.org/pkg/config v0.0.6
	go.zenithar.org/pkg/db v0.0.3
	go.zenithar.org/pkg/db/adapter/postgresql v0.0.7
	go.zenithar.org/pkg/errors v0.0.1
	go.zenithar.org/pkg/flags v0.0.2
	go.zenithar.org/pkg/log v0.2.0
	go.zenithar.org/pkg/platform v0.1.4
	go.zenithar.org/pkg/reactor v0.0.4
	go.zenithar.org/pkg/testing v0.0.13
	go.zenithar.org/pkg/tlsconfig v0.0.1
	go.zenithar.org/pkg/types v0.0.1
	golang.org/x/crypto v0.0.0-20190907121410-71b5226ff739
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/grpc v1.23.1
	gopkg.in/gorp.v1 v1.7.2 // indirect
)
