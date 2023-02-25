package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	p := fmt.Println

	s := "Mysql://admin:password@serverhost.com:8081/server/page1?key=value&key2=value2#X"

	u, err := url.Parse(s)

	if err != nil {
		panic(err)
	}

	p(u.Scheme)
	p(u.User)
	p(u.User.Username())
	pass, passErr := u.User.Password()
	if !passErr {
		p("no any password available")
	}
	p(pass)
	p(u.Host)
	host, port, splitErr := net.SplitHostPort(u.Host)
	if splitErr != nil {
		p("error while extracting host and port", splitErr)
	}
	p(host)
	p(port)
	p(u.Path)
	p(u.Fragment)
	p(u.RawQuery)
	m, queryErr := url.ParseQuery(u.RawQuery)
	if queryErr != nil {
		p("error while extracting raw query", queryErr)
	}
	p(m)
}
