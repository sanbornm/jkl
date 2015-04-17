package main

import (
	"testing"
)

func TestGetShortDescription(t *testing.T) {

	p := Page{"content": "foo<!--more-->blah foobar"}
	resp := p.GetShortDescription()
	if resp != "foo" {
		t.Errorf("Expected foo got [%s]", resp)
	}

	p = Page{"content": "fooblah foobar"}
	resp = p.GetShortDescription()
	if resp != "fooblah foobar" {
		t.Errorf("Expected fooblah foobar got [%s]", resp)
	}
}
