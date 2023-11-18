package dbm

import (
	"fmt"
	"io"
)

type Table struct {
	Name    string
	Columns []Column
}

func (t *Table) WriteSQL(w io.Writer) error {
	_, err := w.Write([]byte(fmt.Sprintf(`CREATE TABLE "%s" (`, t.Name)))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\n"))
	if err != nil {
		return err
	}
	for i, c := range t.Columns {
		_, err = w.Write([]byte("\t"))
		if err != nil {
			return err
		}
		_, err := w.Write([]byte(fmt.Sprintf(`"%s" %s`, c.Name, c.Type.SQLType())))
		if err != nil {
			return err
		}
		if !c.Nullable {
			_, err = w.Write([]byte(" NOT NULL"))
			if err != nil {
				return err
			}
		}
		if c.PrimaryKey {
			_, err := w.Write([]byte(" PRIMARY KEY"))
			if err != nil {
				return err
			}
		}
		if i < len(t.Columns)-1 {
			_, err := w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte(");\n"))
	return nil
}
