package dbm

type TDiffType int

const (
	TDiffTypeCreateTable = iota + 1
	TDiffTypeAlterColumn
)

const (
	CDiffTypeAddColumn = iota + 1
	CDiffTypeDropColumn
)

type tableDiffCalc struct{}

func newTableDiffCalc() *tableDiffCalc {
	return &tableDiffCalc{}
}

func (t *tableDiffCalc) Difference(before, after *Table) []*tableDifference {
	diff := make([]*tableDifference, 0)
	if cDiff := t.diffColumns(before, after); cDiff != nil {
		diff = append(diff, cDiff)
	}
	diff = append(diff)
	return diff
}

func (t *tableDiffCalc) diffColumns(before, after *Table) *tableDifference {
	beforeColumns, afterColumns := t.mapColumns(before, after)
	collDiff := make([]columnDifference, 0)
	collDiff = append(collDiff, t.diffDrop(beforeColumns, afterColumns)...)
	collDiff = append(collDiff, t.diffAddColumn(beforeColumns, afterColumns)...)

	if len(collDiff) == 0 {
		return nil
	}
	return &tableDifference{
		Type:              TDiffTypeAlterColumn,
		ColumnDifferences: collDiff,
	}
}

func (t *tableDiffCalc) mapColumns(before, after *Table) (map[string]Column, map[string]Column) {
	beforeColumns := make(map[string]Column)
	for _, column := range before.Columns {
		beforeColumns[column.Name] = column
	}
	afterColumns := make(map[string]Column)
	for _, column := range after.Columns {
		afterColumns[column.Name] = column
	}
	return beforeColumns, afterColumns
}

func (t *tableDiffCalc) diffDrop(beforeColumns, afterColumns map[string]Column) []columnDifference {
	collDiff := make([]columnDifference, 0)
	for name := range beforeColumns {
		if _, ok := afterColumns[name]; !ok {
			collDiff = append(collDiff, columnDifference{
				Type:       CDiffTypeDropColumn,
				ColumnName: name,
			})
		}
	}
	return collDiff
}

func (t *tableDiffCalc) diffAddColumn(beforeColumns, afterColumns map[string]Column) []columnDifference {
	collDiff := make([]columnDifference, 0)
	for name, col := range afterColumns {
		if _, ok := beforeColumns[name]; !ok {
			collDiff = append(collDiff, columnDifference{
				Type:       CDiffTypeAddColumn,
				ColumnName: name,
				DataType:   col.Type,
			})
		}
	}
	return collDiff
}
