##gunk get
<hr/>
https://github.com/gunk/gunk

##generate
<hr/>
./generate.sh

##goose add
<hr/>
go get github.com/pressly/goose


##migrations
<hr/>
go run migrations/migrate.go create create_blog_tabls sql
<hr/>
go run migrations/migrate.go up
<hr/>
Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version


##test
<hr/>

cd blog/storage/postgres

DATABASE_CONNECTION="user=postgres password=dipto host=localhost port=5432 sslmode=disable" go test 