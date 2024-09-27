package main

import (
	"golang.org/x/net/context"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		ctx context.Context
		in  []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Check basic div parse",
			args: args{
				ctx: context.Background(),
				in:  []byte("<div></div>"),
			},
			want: "package example\n" +
				"\n" +
				"import (\n" +
				"\tg \"github.com/maragudk/gomponents\"\n" +
				"\tc \"github.com/maragudk/gomponents/components\"\n" +
				"\t. \"github.com/maragudk/gomponents/html\"\n" +
				")\n" +
				"\n" +
				"func example() g.Node {\n" +
				"\treturn Body(\n" +
				"\t\tDiv(),\n" +
				"\t)\n" +
				"}\n",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
