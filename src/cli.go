package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// 終了コード
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
	ExitCodeFlagParseError
	ExitCodeRemoveError
	ExitCodeRestoreError
	ExitCodeLoggingError
	ExitCodeBadArgs
)

type CLI struct {
	outStream io.Writer
	errStream io.Writer
}

// 引数処理を含めた具体的な処理
func (cli *CLI) Run(args []string) int {
	// オプション引数のパース
	var options Options
	flags := flag.NewFlagSet("syldbimport", flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
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
		return ExitCodeFlagParseError
	}

	if err := options.Check(); err != nil {
		flags.Usage()
		return ExitCodeError
	}

	if len(flags.Args()) == 0 {
		fmt.Println("specify target folder.")
		return ExitCodeError
	}
	var folder_id string
	folder_id = flags.Args()[0]

	fmt.Println(folder_id)

	return ExitCodeOK
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
