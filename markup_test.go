package markup

import (
	"testing"
)

func TestMarkup(t *testing.T) {
	metas, err := Parse(`hello="world"
contains_quotes = "contains\" quotes"
key =  value text 
goodnight = "Moon"
`)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", metas)
}
