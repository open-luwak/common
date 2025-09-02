package common

import "testing"

func TestNewAPIName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{
			name:    "basic case",
			input:   "db.schema.table.action",
			want:    []string{"", "db.schema.table.action", "db.schema.table", "action"},
			wantErr: false,
		},
		{
			name:    "single prefix",
			input:   "prefix.myDb.mySchema.myTable.insert",
			want:    []string{"prefix", "myDb.mySchema.myTable.insert", "myDb.mySchema.myTable", "insert"},
			wantErr: false,
		},
		{
			name:    "multiple prefixes",
			input:   "a.b.c.myDb.mySchema.myTable.delete",
			want:    []string{"a.b.c", "myDb.mySchema.myTable.delete", "myDb.mySchema.myTable", "delete"},
			wantErr: false,
		},
		{
			name:    "invalid input - too few parts",
			input:   "db.schema.action",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMethodName(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMethodName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if got.OrgName != tt.want[0] {
				t.Errorf("ParseMethodName() OrgName = %v, want %v",
					got.OrgName, tt.want[0])
			}
			if got.ApiName != tt.want[1] {
				t.Errorf("ParseMethodName() ApiName = %v, want %v",
					got.ApiName, tt.want[1])
			}
			if got.EntityName != tt.want[2] {
				t.Errorf("ParseMethodName() EntityName = %v, want %v",
					got.EntityName, tt.want[3])
			}
			if got.EntityMethod != tt.want[3] {
				t.Errorf("ParseMethodName() EntityMethod = %v, want %v",
					got.EntityMethod, tt.want[3])
			}
		})
	}
}
