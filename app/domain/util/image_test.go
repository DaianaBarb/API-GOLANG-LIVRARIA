package util

import "testing"

func TestNormalizeURL(t *testing.T) {
	type args struct {
		url  string
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "successfully complete url",
			args: args{
				url:  "http://www.americanas.com.br",
				path: "/produto/12345678",
			},
			want: "http://www.americanas.com.br/produto/12345678",
		},
		{
			name: "should return empty string when path is empty",
			args: args{
				url:  "http://www.americanas.com.br",
				path: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeURL(tt.args.url, tt.args.path); got != tt.want {
				t.Errorf("NormalizeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
