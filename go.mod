module gohash_miner_simulator

go 1.18

require (
	github.com/spf13/cobra v1.4.0
	loader v0.0.0-00010101000000-000000000000
	server v0.0.0-00010101000000-000000000000
	validation v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069 // indirect
	golang.org/x/text v0.3.7 // indirect
	miner_sim v0.0.0-00010101000000-000000000000 // indirect
)

// replace validation => ../validation

replace validation => ./validation

replace server => ./server

replace loader => ./loader

replace miner_sim => ./miner_sim
