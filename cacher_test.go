package main

import (
	"testing"
)

func TestStdoutCacherGenerateCacheFilename(t *testing.T) {
	type fields struct {
		ttl         int
		command     string
		args        []string
		environment []string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "foo",
			fields: fields{
				ttl:         2,
				command:     "thiscommmanddoesnotexist",
				args:        nil,
				environment: nil,
			},
			want: "thiscommmanddoesnotexist_eac6bfe995dfb0973514d4e6cf8d31f2ba159ce3301cb8a454e3eeb08982557f.cache",
		},
		{
			name: "foo",
			fields: fields{
				ttl:         2,
				command:     "thiscommmanddoesnotexist",
				args:        []string{"foo", "bar"},
				environment: nil,
			},
			want: "thiscommmanddoesnotexist_2636829e518b6df820cfa8f277c31718873c3276eebe8fa80f6f2453d9399366.cache",
		},
		{
			name: "foo",
			fields: fields{
				ttl:         2,
				command:     "thiscommmanddoesnotexist",
				args:        []string{"foo", "bar"},
				environment: []string{"duu=daa"},
			},
			want: "thiscommmanddoesnotexist_f422bb509bf5605fff0e89c44969bce4c7f6142d24c465921aaca244ddf0704c.cache",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacher := &StdoutCacher{
				ttl:         tt.fields.ttl,
				command:     tt.fields.command,
				args:        tt.fields.args,
				environment: tt.fields.environment,
			}
			if got := cacher.generateCacheFilename(); got != tt.want {
				t.Errorf("StdoutCacher.generateCacheFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStdoutCacherRunCommand(t *testing.T) {
	type fields struct {
		ttl         int
		command     string
		args        []string
		environment []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacher := &StdoutCacher{
				ttl:         tt.fields.ttl,
				command:     tt.fields.command,
				args:        tt.fields.args,
				environment: tt.fields.environment,
			}
			got, err := cacher.RunCommand()
			if (err != nil) != tt.wantErr {
				t.Errorf("StdoutCacher.RunCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StdoutCacher.RunCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
