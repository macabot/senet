module github.com/macabot/senet

go 1.18

require (
	github.com/macabot/fairytale v0.0.0
	github.com/macabot/hypp v0.0.0
	github.com/stretchr/testify v1.8.4
	golang.org/x/crypto v0.15.0
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gosimple/slug v1.12.0 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/macabot/fairytale => ../fairytale
	github.com/macabot/hypp => ../hypp
)
