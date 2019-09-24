module github.com/gautamrege/packt/sweatbead/reportsmgr

go 1.12

replace github.com/gautamrege/packt/sweatbead/proto => /Users/gautam/work/sweatbead/proto

require (
	github.com/gautamrege/packt/sweatbead/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.23.0
)
