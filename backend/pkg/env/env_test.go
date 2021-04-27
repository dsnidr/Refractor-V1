/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package env

import (
	"os"
	"reflect"
	"testing"
)

func TestEnv_GetError(t *testing.T) {
	type fields struct {
		missingVars []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "TestEnv_GetError-01",
			fields: fields{
				missingVars: []string{
					"TEST",
					"TEST2",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Env{
				missingVars: tt.fields.missingVars,
			}
			if err := e.GetError(); (err != nil) != tt.wantErr {
				t.Errorf("GetError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEnv_RequireEnv(t *testing.T) {
	if err := os.Setenv("TESTENV-01", "1"); err != nil {
		t.Fatal("Could not set TESTENV-01 env variable.")
	}

	type fields struct {
		missingVars []string
	}
	type args struct {
		varName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Env
	}{
		{
			name: "TestEnv_RequireEnv-01",
			fields: fields{
				missingVars: []string{},
			},
			args: args{
				varName: "TESTENV-01",
			},
			want: &Env{missingVars: []string{}},
		},
		{
			name: "TestEnv_RequireEnv-02",
			fields: fields{
				missingVars: []string{},
			},
			args: args{
				varName: "TESTENV-UNDEFINED",
			},
			want: &Env{missingVars: []string{"TESTENV-UNDEFINED"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Env{
				missingVars: tt.fields.missingVars,
			}
			if got := e.RequireEnv(tt.args.varName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequireEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequireEnv(t *testing.T) {
	if err := os.Setenv("TESTENV-01", "1"); err != nil {
		t.Fatal("Could not set TESTENV-01 env variable.")
	}

	type args struct {
		varName string
	}
	tests := []struct {
		name string
		args args
		want *Env
	}{
		{
			name: "TestRequireEnv-01",
			args: args{
				varName: "TESTENV-01",
			},
			want: &Env{missingVars: []string{}},
		},
		{
			name: "TestRequireEnv-02",
			args: args{
				varName: "TESTENV-UNDEFINED",
			},
			want: &Env{missingVars: []string{"TESTENV-UNDEFINED"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequireEnv(tt.args.varName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequireEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
