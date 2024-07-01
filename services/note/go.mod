module github.com/yahn1ukov/scribble/services/note

go 1.22.4

replace (
	github.com/yahn1ukov/scribble/libs/grpc => ../../libs/grpc
	github.com/yahn1ukov/scribble/libs/respond => ../../libs/respond
)

require (
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/lib/pq v1.10.9
	github.com/yahn1ukov/scribble/libs/grpc v0.0.0-00010101000000-000000000000
	github.com/yahn1ukov/scribble/libs/respond v0.0.0-00010101000000-000000000000
	go.uber.org/fx v1.22.1
	google.golang.org/grpc v1.64.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
