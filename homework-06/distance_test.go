package distance

import "testing"

func TestCalculateDistance(t *testing.T) {
	type args struct {
		x1 float64
		x2 float64
		y1 float64
		y2 float64
	}
	tests := []struct {
		name         string
		x1           float64
		x2           float64
		y1           float64
		y2           float64
		wantDistance float64
	}{
		{
			name:         "When coordinates are positive",
			x1:           2.2,
			x2:           4.2,
			y1:           3.3,
			y2:           6.3,
			wantDistance: 3.605551275463989,
		},
		{
			name:         "When some coordinates are negative",
			x1:           -2.2,
			x2:           4.2,
			y1:           -3.3,
			y2:           6.3,
			wantDistance: 3.605551275463989,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := CalculateDistance(tt.x1, tt.x2, tt.y1, tt.y2); gotDistance != tt.wantDistance {
				t.Errorf("CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
