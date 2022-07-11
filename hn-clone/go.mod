module example/hn-clone

go 1.18

replace example/internal => ./internal

require (
	example/internal v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

require (
	github.com/go-sql-driver/mysql v1.6.0
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
)
