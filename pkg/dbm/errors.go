package dbm

import "errors"

var ErrNoTableName = errors.New("no table name specified")
var ErrIncompleteTable = errors.New("incomplete table")
