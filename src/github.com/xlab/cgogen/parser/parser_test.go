package parser

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/cznic/c/internal/cc"
)

func TestParse(t *testing.T) {
	p, err := New(NewConfig("test/parser_test.h"))
	if err != nil {
		t.Fatal(err)
	}
	unit, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	testUnit(t, unit)
}

func testUnit(t *testing.T, u *cc.TranslationUnit) {
	buf, err := ioutil.ReadFile("test/parser_test.out")
	if err != nil {
		t.Fatal(err)
	} else if len(buf) == 0 {
		t.Fatal("no reference output provided")
	}
	if u == nil {
		t.Fatal("no translation unit returned")
	}
	if u.String() != string(buf) {
		log.Println(u)
		t.Error("output doesn't match reference")
	}
}
