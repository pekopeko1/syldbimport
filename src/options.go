package main

import (
	"fmt"
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
