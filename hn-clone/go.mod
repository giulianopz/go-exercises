module example/hn-clone

go 1.18

replace example/internal => ./internal

require (
	example/internal v0.0.0-00010101000000-000000000000
	golang.org/x/exp v0.0.0-20220613132600-b0d781184e0d // indirect
)
