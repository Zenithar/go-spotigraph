module go.zenithar.org/spotigraph

go 1.12

require (
	github.com/CircleCI-Public/circleci-cli v0.1.5490
	github.com/Masterminds/squirrel v1.1.0
	github.com/asaskevich/govalidator v0.0.0-20180720115003-f9ffefc3facf // indirect
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/cloudfoundry-incubator/candiedyaml v0.0.0-20170901234223-a41693b7b7af // indirect
	github.com/dchest/uniuri v0.0.0-20160212164326-8902c56451e9
	github.com/fatih/structs v1.1.0
	github.com/frapposelli/wwhrd v0.2.1
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/gogo/protobuf v1.2.1
	github.com/golang/mock v1.2.0
	github.com/golang/protobuf v1.3.1
	github.com/golangci/golangci-lint v1.15.0
	github.com/google/gops v0.3.6
	github.com/google/wire v0.2.1
	github.com/gosimple/slug v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hexdigest/gowrap v1.1.2
	github.com/hokaccha/go-prettyjson v0.0.0-20180920040306-f579f869bbfe
	github.com/iancoleman/strcase v0.0.0-20180726023541-3605ed457bf7 // indirect
	github.com/jessevdk/go-flags v1.4.0 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/lyft/protoc-gen-validate v0.0.14
	github.com/magefile/mage v1.8.0
	github.com/mcuadros/go-defaults v1.1.0
	github.com/mongodb/mongo-go-driver v1.0.0
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/ryanuber/go-license v0.0.0-20180405065157-c69f41c2c8d6 // indirect
	github.com/sirupsen/logrus v1.4.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.3.2
	go.mongodb.org/mongo-driver v1.0.0
	go.opencensus.io v0.19.2
	go.uber.org/zap v1.9.1
	go.zenithar.org/pkg v0.0.6
	go.zenithar.org/protoc-gen-cobra v0.0.3
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/oauth2 v0.0.0-20190319182350-c85d3e98c914
	google.golang.org/grpc v1.19.1
	gopkg.in/rethinkdb/rethinkdb-go.v5 v5.0.1
	gotest.tools/gotestsum v0.3.4
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)

replace go.zenithar.org/pkg => github.com/Zenithar/go-pkg v0.0.6

replace go.zenithar.org/protoc-gen-cobra => github.com/Zenithar/go-protoc-gen-cobra v0.0.3
