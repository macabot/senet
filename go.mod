module github.com/macabot/senet

go 1.18

require (
	github.com/macabot/fairytale v0.0.0
	github.com/macabot/hypp v0.0.0
)

require (
	github.com/gosimple/slug v1.12.0 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	golang.org/x/exp v0.0.0-20220613132600-b0d781184e0d // indirect
)

replace (
	github.com/macabot/fairytale => ../fairytale
	github.com/macabot/hypp => ../hypp
)
