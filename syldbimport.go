package main

import (
    "flag"
)

var opt_dbname = flag.String("d", "", "database name")
var opt_hostname = flag.String("h", "", "hostname of database server")
var opt_port = flag.Int("p", 0, "port number of database server")
var opt_username = flag.String("U", "", "username of database")
var opt_password = flag.String("P", "", "password of database")
var opt_mecab_encoding = flag.String("-mecab-encoding", "", "encoding of MeCab dictionary (default: Unix: EUC-JP / Win32: Shift_JIS)")
var opt_dont_remove_messeges = flag.Bool("n", false, "don't remove nonexist messages")
var opt_recursive_import = flag.Bool("r", false, "recursive import")
var opt_verbose = flag.Bool("v", false, "verbose output")
var opt_exclude = flag.String("-exclude", "", "exclude foldername from import targets (can be specified multiple times)")
var opt_delete = flag.Bool("-delete", false, "recursively delete folders from DB")
var opt_debug = flag.Bool("-debug", false, "show debug console window")


func main() {
    flag.Parse()
    // if command has no args, show help and exit
    if flag.NFlag() == 0 {
        flag.Usage()
        return
    }

}
