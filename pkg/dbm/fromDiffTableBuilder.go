package dbm

type TableValidatorer interface {
	Validate(table *Table) error
}

type FromDiffTableBuilder struct {
	diffs     []*tableDifference
	name      string
	validator TableValidatorer
}

func NewFromDiffTableBuilder() *FromDiffTableBuilder {
	return &FromDiffTableBuilder{
		diffs:     make([]*tableDifference, 0),
		validator: NewTableValidator(),
	}
}

func (b *FromDiffTableBuilder) SetName(name string) {
	b.name = name
}

func (b *FromDiffTableBuilder) AddDiff(diff *tableDifference) {
	b.diffs = append(b.diffs, diff)
}

func (b *FromDiffTableBuilder) build() (*Table, error) {
	table := &Table{
		Name:    b.name,
		Columns: make([]Column, 0),
	}
	if err := b.applyDiffs(table); err != nil {
		return nil, err
	}
	if err := b.validator.Validate(table); err != nil {
		return nil, err
	}
	return table, nil
}

func (b *FromDiffTableBuilder) applyDiffs(table *Table) error {
	for _, diff := range b.diffs {
		if err := b.applyDiff(table, diff); err != nil {
			return err
		}
	}
	return nil
}

func (b *FromDiffTableBuilder) applyDiff(table *Table, diff *tableDifference) (err error) {
	switch diff.Type {
	case TDiffTypeCreateTable:
		err = b.applyCreateTableDiff(table, diff)
	case TDiffTypeAlterColumn:
		err = b.applyAlterColumnDiff(table, diff)
	}
	return err
}

func (b *FromDiffTableBuilder) applyCreateTableDiff(table *Table, diff *tableDifference) error {
	table.Name = diff.TableName
	b.applyColumnDiffs(table, diff.ColumnDifferences)
	return nil
}

func (b *FromDiffTableBuilder) applyAlterColumnDiff(table *Table, diff *tableDifference) error {
	return nil
}

func (b *FromDiffTableBuilder) applyColumnDiffs(table *Table, diff []columnDifference) error {
	for _, columnDiff := range diff {
		if err := b.applyColumnDiff(table, columnDiff); err != nil {
			return err
		}
	}
	return nil
}

func (b *FromDiffTableBuilder) applyColumnDiff(table *Table, columnDiff columnDifference) error {
	switch columnDiff.Type {
	case CDiffTypeAddColumn:
		b.applyAddColumnDiff(table, columnDiff)
	}
	return nil
}

func (b *FromDiffTableBuilder) applyAddColumnDiff(table *Table, columnDiff columnDifference) error {
	table.Columns = append(table.Columns, Column{
		Name:       columnDiff.ColumnName,
		Type:       columnDiff.DataType,
		Nullable:   columnDiff.Nullable == DFTrue,
		PrimaryKey: columnDiff.PrimaryKey == DFTrue,
	})
	return nil
}
