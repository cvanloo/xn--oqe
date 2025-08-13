package markup

import (
	"testing"
	"reflect"
)

func TestMarkupMeta(t *testing.T) {
	const case1 = `hello="world"
contains_quotes = "here is a \" quote"
key =  value text 
goodnight= Moon
`
	testCases := []struct{
		src string
		correct bool
		exp map[string]string
	} {
		{case1, true, map[string]string{"hello": "world", "contains_quotes": "here is a \" quote", "key": "value text ", "goodnight": "Moon"}},
		{"hello=\"world\"\ngoodnight=\"Moon\"", true, map[string]string{"hello": "world", "goodnight": "Moon"}},
		{"", true, map[string]string{}},
		{"\n\n  \n \n", true, map[string]string{}},
		{"\n\nhello=\"world\"  \n \n", true, map[string]string{"hello": "world"}},
		{"\n\n  hello=\"world\"\n  goodnight=\"moon\" \n", true, map[string]string{"hello": "world", "goodnight": "moon"}},
		{"\n\n  hello=\"world\"  \n  goodnight=\"moon\" \n", true, map[string]string{"hello": "world", "goodnight": "moon"}},
	}
	for _, testCase := range testCases {
		metas, err := Parse(testCase.src)
		if !testCase.correct && err == nil {
			t.Error("expected to fail, but actually returned err=nil")
		}
		if testCase.correct && err != nil {
			t.Error(err)
		}
		if testCase.correct && !reflect.DeepEqual(metas, testCase.exp) {
			t.Errorf("expected: `%#v`, got: `%#v`", testCase.exp, metas)
		}
	}
}
