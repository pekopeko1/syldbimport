package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/lib/pq"
	"io"
	"path/filepath"
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

func (cli *CLI) Run(args []string) int {
	options, err := NewOptions(args, cli.errStream)
	if err != nil {
		return ExitCodeFlagParseError
	}

	if err := options.Check(); err != nil {
		options.Usage(cli.errStream)
		return ExitCodeError
	}

	var folder_id string
	folder_id = options.folder_id

	fmt.Println(folder_id)
	var db_conninfo string
	db_conninfo = options.GetConnInfo()

	db, err := sql.Open("postgres", db_conninfo)
	if err != nil {
		fmt.Println("connection to database failed")
	}
	defer db.Close()

	// * if opt_delete {
	// * db_delete_folder()
	// * if fail "remove folder failed: %s", folder_id

	db_import_folder(db, folder_id, options)

	return ExitCodeOK
}

func db_import_mh_folder(db *sql.DB, pathname string, options *Options) int {
	fmt.Println("run db_import_mh_folder")
	fmt.Printf("\nimporting %s ...\n", pathname)
	return ExitCodeOK
}

func db_import_folder_item(db *sql.DB, item string, options *Options) int {
	fmt.Println("run db_import_folder_item")
	return ExitCodeOK
}

func db_import_folder(db *sql.DB, folder_id string, options *Options) int {
	if filepath.IsAbs(folder_id) {
		db_import_mh_folder(db, folder_id, options)
	} else {
		var new_folder_id string
		new_folder_id, _ = filepath.Abs(folder_id)
		db_import_mh_folder(db, new_folder_id, options)
	}
	return ExitCodeOK
}
