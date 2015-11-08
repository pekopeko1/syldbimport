package main

import (
	"fmt"
	"strconv"
)

type Options struct {
	dbname           string
	hostname         string
	port             int
	username         string
	password         string
	mecab_encoding   string
	del_noexist      bool
	recursive_import bool
	verbose          bool
	exclude          string
	del              bool
	debug            bool
}

func (options *Options) Check() (err error) {
	switch {
	case options.dbname == "":
		err = fmt.Errorf("Invalid Arugment")
		return err
	case options.port < 0 || options.port > 655335:
		err = fmt.Errorf("Invalid Arugment")
		return err
	}
	err = nil
	return err
}

func (options *Options) GetConnInfo() (conninfo string) {
	conninfo += "dbname=" + options.dbname
	if options.hostname != "" {
		conninfo += " host=" + options.hostname
	}
	conninfo += " port=" + strconv.Itoa(options.port)
	if options.username != "" {
		conninfo += " user=" + options.username
	}
	if options.password != "" {
		conninfo += " password=" + options.password
	}
	fmt.Println(conninfo)
	return conninfo
}
