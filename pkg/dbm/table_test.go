package dbm

import (
	"bytes"
	"strings"
	"testing"
)

const createTable1 = `
CREATE TABLE "users" (
	"id" INTEGER NOT NULL PRIMARY KEY,
	"name" VARCHAR(64) NOT NULL,
	"age" SMALLINT NOT NULL
);
`

func Test_model_creation(t *testing.T) {
	users := Table{
		Name: "users",
		Columns: []Column{
			{Name: "id", Type: Integer(), PrimaryKey: true},
			{Name: "name", Type: Varchar(64)},
			{Name: "age", Type: SmallInteger()},
		},
	}
	var output string
	buf := bytes.NewBufferString(output)
	users.WriteSQL(buf)
	if strings.Join(strings.Split(buf.String(), "\n"), "") != strings.Join(strings.Split(createTable1, "\n"), "") {
		t.Errorf("expected %s, got %s", createTable1, buf.String())
	}

}

func Test_model_alter(t *testing.T) {
	users1 := Table{
		Name: "users",
		Columns: []Column{
			{Name: "id", Type: Integer(), PrimaryKey: true},
			{Name: "name", Type: Varchar(64)},
			{Name: "age", Type: SmallInteger()},
		},
	}
	users2 := Table{
		Name: "users",
		Columns: []Column{
			{Name: "id", Type: Integer(), PrimaryKey: true},
			{Name: "name", Type: Varchar(64)},
			{Name: "age", Type: SmallInteger()},
		},
	}
	diffCalc := newTableDiffCalc()
	diff := diffCalc.Difference(&users1, &users2)
	if len(diff) != 0 {
		t.Errorf("expected %d, got %d", 0, len(diff))
	}
}

func Test_model_drop_column(t *testing.T) {
	users1 := Table{
		Name: "users",
		Columns: []Column{
			{Name: "id", Type: Integer(), PrimaryKey: true},
			{Name: "name", Type: Varchar(64)},
			{Name: "age", Type: SmallInteger()},
		},
	}
	users2 := Table{
		Name: "users",
		Columns: []Column{
			{Name: "id", Type: Integer(), PrimaryKey: true},
			{Name: "name", Type: Varchar(64)},
		},
	}
	diffCalc := newTableDiffCalc()
	diff := diffCalc.Difference(&users1, &users2)
	if len(diff) != 1 {
		t.Errorf("expected %d, got %d", 1, len(diff))
	}
	if diff[0].Type != TDiffTypeAlterColumn {
		t.Errorf("expected %d, got %d", TDiffTypeAlterColumn, diff[0].Type)
	}
	if len(diff[0].ColumnDifferences) != 1 {
		t.Errorf("expected %d, got %d", 1, len(diff[0].ColumnDifferences))
	}
	if diff[0].ColumnDifferences[0].ColumnName != "age" {
		t.Errorf("expected %s, got %s", "age", diff[0].ColumnDifferences[0].ColumnName)
	}
	if diff[0].ColumnDifferences[0].Type != CDiffTypeDropColumn {
		t.Errorf("expected %d, got %d", CDiffTypeDropColumn, diff[0].ColumnDifferences[0].Type)
	}
}
func Test_model_add_column(t *testing.T) {
	users1 := Table{
		Name: "users",
		Columns: []Column{
			{Name: "id", Type: Integer(), PrimaryKey: true},
			{Name: "name", Type: Varchar(64)},
		},
	}
	users2 := Table{
		Name: "users",
		Columns: []Column{
			{Name: "id", Type: Integer(), PrimaryKey: true},
			{Name: "name", Type: Varchar(64)},
			{Name: "age", Type: SmallInteger()},
		},
	}
	diffCalc := newTableDiffCalc()
	diff := diffCalc.Difference(&users1, &users2)
	if len(diff) != 1 {
		t.Errorf("expected %d, got %d", 1, len(diff))
	}
	if diff[0].Type != TDiffTypeAlterColumn {
		t.Errorf("expected %d, got %d", TDiffTypeAlterColumn, diff[0].Type)
	}
	if len(diff[0].ColumnDifferences) != 1 {
		t.Errorf("expected %d, got %d", 1, len(diff[0].ColumnDifferences))
	}
	if diff[0].ColumnDifferences[0].ColumnName != "age" {
		t.Errorf("expected %s, got %s", "age", diff[0].ColumnDifferences[0].ColumnName)
	}
	if diff[0].ColumnDifferences[0].Type != CDiffTypeAddColumn {
		t.Errorf("expected %d, got %d", CDiffTypeAddColumn, diff[0].ColumnDifferences[0].Type)
	}
	if diff[0].ColumnDifferences[0].DataType.SQLType() != SmallInteger().SQLType() {
		t.Errorf("expected %s, got %s", SmallInteger().SQLType(), diff[0].ColumnDifferences[0].DataType.SQLType())
	}

}
