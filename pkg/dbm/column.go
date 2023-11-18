package dbm

type SQLTyper interface {
	SQLType() string
}

type Column struct {
	Name          string
	Type          SQLTyper
	Nullable      bool
	AutoIncrement bool
	PrimaryKey    bool
}
