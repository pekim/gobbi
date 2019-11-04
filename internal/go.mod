module github.com/pekim/gobbi/internal

replace github.com/dave/jennifer v1.3.0 => github.com/pekim/jennifer v1.3.1-0.20190504080055-d88ee566d149

replace github.com/pekim/gobbi v0.0.0 => ../

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/boyter/scc v2.4.0+incompatible
	github.com/buger/goterm v0.0.0-20181115115552-c206103e1f37
	github.com/dave/jennifer v1.3.0
	github.com/gomarkdown/markdown v0.0.0-20191104174740-4d42851d4d5a
	github.com/pekim/gobbi v0.0.0
	github.com/stretchr/testify v1.3.0
	golang.org/x/sys v0.0.0-20191029155521-f43be2a4598c // indirect
)

go 1.13
