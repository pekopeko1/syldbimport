package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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
	folder_id        string
}

func (options *Options) Check() (err error) {
	switch {
	case options.dbname == "":
		err = fmt.Errorf("Invalid Arugment")
		return err
	case options.port < 0 || options.port > 655335:
		err = fmt.Errorf("Invalid Arugment")
		return err
	case options.folder_id == "":
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

func (options *Options) Usage() {
	fmt.Fprint(os.Stderr, helpText)
}

func NewOptions(args []string, output io.Writer) (p_options *Options, err error) {
	var options Options
	// オプション引数のパース
	flags := flag.NewFlagSet("syldbimport", flag.ContinueOnError)
	flags.SetOutput(output)
	flags.Usage = func() {
		fmt.Fprint(os.Stderr, helpText)
	}
	flags.StringVar(&options.dbname, "d", "sylph", "database name")
	flags.StringVar(&options.hostname, "h", "", "hostname of database server")
	flags.IntVar(&options.port, "p", 25432, "port number of database server")
	flags.StringVar(&options.username, "U", "", "username of database")
	flags.StringVar(&options.password, "P", "", "password of database")
	flags.StringVar(&options.mecab_encoding, "mecab-encoding", "Shift_JIS", "encoding of MeCab dictionary (default: Unix: EUC-JP / Win32: Shift_JIS)")
	flags.BoolVar(&options.del_noexist, "n", false, "don't remove nonexist messages")
	flags.BoolVar(&options.recursive_import, "r", false, "recursive import")
	flags.BoolVar(&options.verbose, "v", false, "verbose output")
	flags.StringVar(&options.exclude, "exclude", "", "exclude foldername from import targets (can be specified multiple times)")
	flags.BoolVar(&options.del, "delete", false, "recursively delete folders from DB")
	flags.BoolVar(&options.debug, "debug", false, "show debug console window")

	if err := flags.Parse(args[1:]); err != nil {
		return nil, err
	}
	if len(flags.Args()) == 0 {
		return nil, fmt.Errorf("specify target folder.")
	}
	options.folder_id = flags.Args()[0]
	p_options = &options
	return p_options, nil
}

var helpText = `Usage: syldbimport.exe [-d dbname] [-h hostname] [-p port] [-U username] [-P password] [--mecab-encoding encoding] ...
Common options:
  -d dbname                     database name
  -h hostname                   hostname of database server
  -p port                       port number of database server
  -U username                   username for database
  -P password                   password for database
  --mecab-encoding encoding     encoding of MeCab dictionary
                                (default: Unix: EUC-JP / Win32: Shift_JIS)
  --help                        show this message
Options for syldbimport:
  -n                            don't remove nonexist messages
  -r                            recursive import
  -v                            verbose output
  --exclude foldername          exclude foldername from import targets
                                (can be specified multiple times)
  --delete                      recursively delete folders from DB
  --debug                       show debug console window
`
