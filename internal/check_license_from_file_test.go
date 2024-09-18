package internal

import "testing"

func TestCheckLicenseFromFile(t *testing.T) {
	type args struct {
		protectedId     string
		pathLicenseFile string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Test 1", args{"X123", "../cmd/example_app.lic"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckLicenseFromFile(tt.args.protectedId, tt.args.pathLicenseFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckLicenseFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckLicenseFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
