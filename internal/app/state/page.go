package state

import (
	"encoding/json"
	"fmt"
)

type Page int

const (
	HomePage Page = iota
	RulesPage
	PlayPage
)

func (p Page) String() string {
	pages := [...]string{
		"HomePage",
		"RulesPage",
		"PlayPage",
	}
	return pages[p]
}

func ToPage(s string) (Page, error) {
	var page Page
	switch s {
	case "HomePage":
		page = HomePage
	case "RulesPage":
		page = RulesPage
	case "PlayPage":
		page = PlayPage
	default:
		return page, fmt.Errorf("invalid Page '%s'", s)
	}
	return page, nil
}

func (p Page) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *Page) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	var err error
	*p, err = ToPage(s)
	return err
}
