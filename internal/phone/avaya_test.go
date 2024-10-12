package phone

import (
	"testing"
)

func TestJSeries(t *testing.T) {
	cases := []struct {
		name string
		ua   string
		want Identity
	}{
		{
			name: "j139",
			ua:   "Mozilla/4.0 (compatible; MSIE 6.0) AVAYA/J139-4.1.5.0.6 (MAC:c81f4aef9689)",
			want: Identity{
				Model: "J139",
				MAC:   "c81f4aef9689",
			},
		},
		{
			name: "j159",
			ua:   "Mozilla/4.0 (compatible; MSIE 6.0) AVAYA/J159-4.1.0 (MAC:c41f3aef9689)",
			want: Identity{
				Model: "J159",
				MAC:   "c41f3aef9689",
			},
		},
		{
			name: "j179",
			ua:   "Mozilla/4.0 (compatible; MSIE 6.0) AVAYA/J179-4.2.0 (MAC:c41f3ae55689)",
			want: Identity{
				Model: "J179",
				MAC:   "c41f3ae55689",
			},
		},
		{
			name: "wrongModel",
			ua:   "Mozilla/4.0 (compatible; MSIE 6.0) AVAYA/R134-4.2.0 (MAC:c41f3ae55689)",
			want: Identity{},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			identity := avayaJSeries(tt.ua)
			if identity != tt.want {
				t.Errorf("got %q, want %q", identity, tt.want)
			}
		})
	}
}
