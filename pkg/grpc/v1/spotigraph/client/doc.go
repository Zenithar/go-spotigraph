package client

//go:generate gowrap gen -g -p go.zenithar.org/spotigraph/internal/services -i Chapter -t ../../../../../tools/templates/services/grpc/client.txt -o chapter.gen.go
//go:generate gowrap gen -g -p go.zenithar.org/spotigraph/internal/services -i Guild -t ../../../../../tools/templates/services/grpc/client.txt -o guild.gen.go
//go:generate gowrap gen -g -p go.zenithar.org/spotigraph/internal/services -i Squad -t ../../../../../tools/templates/services/grpc/client.txt -o squad.gen.go
//go:generate gowrap gen -g -p go.zenithar.org/spotigraph/internal/services -i Tribe -t ../../../../../tools/templates/services/grpc/client.txt -o tribe.gen.go
//go:generate gowrap gen -g -p go.zenithar.org/spotigraph/internal/services -i User -t ../../../../../tools/templates/services/grpc/client.txt -o user.gen.go
