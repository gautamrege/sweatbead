module github.com/gautamrege/packt/sweatbead/samplemgr

require (
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/mux v1.7.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/urfave/negroni v1.0.0
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.5.1
)

replace (
	github.com/gautamrege/packt/sweatbead/samplemgr/db => /Users/gautam/work/sweatbead/samplemgr/db
	github.com/gautamrege/packt/sweatbead/samplemgr/service => /Users/gautam/work/sweatbead/samplemgr/service
)
