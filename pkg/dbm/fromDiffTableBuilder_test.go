package dbm

import "testing"

func Test_from_diff_table_builder_build_failed_no_name(t *testing.T) {
	fromDiffTableBuilder := NewFromDiffTableBuilder()
	_, err := fromDiffTableBuilder.build()
	if nil == err {
		t.Fatal("expected error with missing name, got nil")
	}
}

func Test_from_diff_table_builder_build_failed_incomplete(t *testing.T) {
	fromDiffTableBuilder := NewFromDiffTableBuilder()
	fromDiffTableBuilder.SetName("Actor")
	_, err := fromDiffTableBuilder.build()
	if nil == err {
		t.Fatal("expected error with incomplete table, got nil")
	}
}

func Test_from_diff_table_builder_build_create_success(t *testing.T) {
	fromDiffTableBuilder := NewFromDiffTableBuilder()
	fromDiffTableBuilder.AddDiff(&tableDifference{
		Type:      TDiffTypeCreateTable,
		TableName: "Actor",
		ColumnDifferences: []columnDifference{
			{Type: CDiffTypeAddColumn, ColumnName: "id", DataType: Integer(), AutoIncrement: DFTrue, PrimaryKey: DFTrue},
			{Type: CDiffTypeAddColumn, ColumnName: "name", DataType: VarcharDefault()},
			{Type: CDiffTypeAddColumn, ColumnName: "comment", DataType: Varchar(1024), Nullable: DFTrue},
		},
	})
	table, err := fromDiffTableBuilder.build()
	if err != nil {
		t.Fatal(err)
	}
	expectedTable := &Table{
		Name: "Actor",
		Columns: []Column{
			{Name: "id", Type: Integer(), PrimaryKey: true, AutoIncrement: true},
			{Name: "name", Type: VarcharDefault()},
			{Name: "comment", Type: Varchar(1024), Nullable: true},
		},
	}
	if table.Name != expectedTable.Name {
		t.Errorf("expected %s, got %s", expectedTable.Name, table.Name)
	}
	for i := 0; i < len(table.Columns); i++ {
		if table.Columns[i].Name != expectedTable.Columns[i].Name {
			t.Errorf("expected %s, got %s", expectedTable.Columns[i].Name, table.Columns[i].Name)
		}
	}
}
func Test_from_diff_table_builder_build_add_column(t *testing.T) {
	fromDiffTableBuilder := NewFromDiffTableBuilder()
	fromDiffTableBuilder.AddDiff(&tableDifference{
		Type:      TDiffTypeCreateTable,
		TableName: "Actor",
		ColumnDifferences: []columnDifference{
			{Type: CDiffTypeAddColumn, ColumnName: "id", DataType: Integer(), AutoIncrement: DFTrue, PrimaryKey: DFTrue},
			{Type: CDiffTypeAddColumn, ColumnName: "name", DataType: VarcharDefault()},
			{Type: CDiffTypeAddColumn, ColumnName: "comment", DataType: Varchar(1024), Nullable: DFTrue},
		},
	})
	fromDiffTableBuilder.AddDiff(&tableDifference{
		Type: CDiffTypeAddColumn,
		ColumnDifferences: []columnDifference{
			{Type: CDiffTypeAddColumn, ColumnName: "name2", DataType: VarcharDefault()},
			{Type: CDiffTypeAddColumn, ColumnName: "comment2", DataType: Varchar(1024), Nullable: DFTrue},
		},
	})
	table, err := fromDiffTableBuilder.build()
	if err != nil {
		t.Fatal(err)
	}
	expectedTable := &Table{
		Name: "Actor",
		Columns: []Column{
			{Name: "id", Type: Integer(), PrimaryKey: true, AutoIncrement: true},
			{Name: "name", Type: VarcharDefault()},
			{Name: "comment", Type: Varchar(1024), Nullable: true},
		},
	}
	if table.Name != expectedTable.Name {
		t.Errorf("expected %s, got %s", expectedTable.Name, table.Name)
	}
	if len(table.Columns) != len(expectedTable.Columns) {
		t.Fatalf("expected %d columns, got %d", len(expectedTable.Columns), len(table.Columns))
	}
	for i := 0; i < len(table.Columns); i++ {
		if table.Columns[i].Name != expectedTable.Columns[i].Name {
			t.Errorf("expected %s, got %s", expectedTable.Columns[i].Name, table.Columns[i].Name)
		}
	}
}
