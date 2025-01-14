package rfc3164_test

import (
	"fmt"

	"github.com/mu-box/golang-syslogparser/rfc3164"
)

func ExampleNewParser() {
	b := "<34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8"
	buff := []byte(b)

	p := rfc3164.NewParser(buff)
	err := p.Parse()
	if err != nil {
		panic(err)
	}

	for k, v := range p.Dump() {
		fmt.Println(k, ":", v)
	}
}
