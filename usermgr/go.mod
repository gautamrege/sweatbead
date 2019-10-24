module github.com/gautamrege/packt/sweatbead/usermgr

go 1.12

require (
	github.com/gautamrege/packt/sweatbead/proto v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.3.1
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.2.2
	github.com/urfave/negroni v1.0.0
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.1.1
	go.uber.org/zap v1.10.0
	google.golang.org/grpc v1.23.1
)

replace (
	github.com/gautamrege/packt/sweatbead/proto => /Users/gautam/work/sweatbead/proto
	github.com/gautamrege/packt/sweatbead/usermgr/config => /Users/gautam/work/sweatbead/usermgr/config
	github.com/gautamrege/packt/sweatbead/usermgr/db => /Users/gautam/work/sweatbead/usermgr/db
	github.com/gautamrege/packt/sweatbead/usermgr/logger => /Users/gautam/work/sweatbead/usermgr/logger
	github.com/gautamrege/packt/sweatbead/usermgr/service => /Users/gautam/work/sweatbead/usermgr/service
)
