module example/hn-clone

go 1.18

replace example/internal => ./internal

require (
	example/internal v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)
