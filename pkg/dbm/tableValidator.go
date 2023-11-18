package dbm

type tableValidator struct{}

func NewTableValidator() *tableValidator {
	return &tableValidator{}
}
func (t *tableValidator) Validate(table *Table) error {
	if table.Name == "" {
		return ErrNoTableName
	}
	if len(table.Columns) == 0 {
		return ErrIncompleteTable
	}
	return nil
}
