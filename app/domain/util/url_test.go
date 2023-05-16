package util

import "testing"

func TestJoinPaths(t *testing.T) {
	type args struct {
		url   string
		paths []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "using path with slash",
			args: args{
				url:   "http://url",
				paths: []string{"/path1", "/path2"},
			},
			want: "http://url/path1/path2",
		},
		{
			name: "using path without slash",
			args: args{
				url:   "http://url",
				paths: []string{"path1", "path2"},
			},
			want: "http://url/path1/path2",
		},
		{
			name: "https scheme",
			args: args{
				url:   "https://url",
				paths: []string{"/path1", "/path2"},
			},
			want: "https://url/path1/path2",
		},
		{
			name: "url without scheme",
			args: args{
				url:   "://url",
				paths: []string{"/path1", "/path2"},
			},
			want: "http://url/path1/path2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinPaths(tt.args.url, tt.args.paths...); got != tt.want {
				t.Errorf("JoinPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
