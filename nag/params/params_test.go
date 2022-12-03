package params_test

import (
	"testing"
	"time"

	"github.com/dylantientcheu/nbacli/nag/params"
)

// probably PBT here
func TestSeason(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "2021-10-01",
			want: "2021-22",
		},
		{
			in:   "2021-09-30",
			want: "2020-21",
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			tim, _ := time.Parse("2006-01-02", tt.in)
			got := params.Season(tim)
			if got != tt.want {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
