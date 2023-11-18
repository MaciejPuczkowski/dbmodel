package dbm

type CDiffType int

type DiffFlag int

const (
	DFNone  = 0
	DFTrue  = 1
	DFFalse = -1
)

type columnDifference struct {
	Type          CDiffType
	ColumnName    string
	ColumnNameNew string
	DataType      SQLTyper
	Default       any
	Nullable      DiffFlag
	AutoIncrement DiffFlag
	PrimaryKey    DiffFlag
	Unique        DiffFlag
}
