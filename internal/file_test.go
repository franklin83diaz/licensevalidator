package internal

import "testing"

func TestReadInternalfile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test 1", args{"testing_file"}, "test_ok", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadInternalfile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInternalfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadInternalfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
