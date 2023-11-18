package dbm

type tableDifference struct {
	Type              TDiffType
	TableName         string
	TableNameNew      string
	ColumnName        string
	ColumnNameNew     string
	ConstraintName    string
	ConstraintNameNew string
	ColumnDifferences []columnDifference
}
