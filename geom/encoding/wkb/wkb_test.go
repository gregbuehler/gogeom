package wkb

import (
	"github.com/twpayne/gogeom/geom"
	"reflect"
	"testing"
)

func TestWKB(t *testing.T) {

	var testCases = []struct {
		g   geom.T
		xdr []byte
		ndr []byte
	}{
		{
			geom.Point{1, 2},
			[]byte("\x00\x00\x00\x00\x01?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@"),
		},
		{
			geom.PointZ{1, 2, 3},
			[]byte("\x00\x00\x00\x03\xe9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\xe9\x03\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			geom.PointM{1, 2, 3},
			[]byte("\x00\x00\x00\x07\xd1?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\xd1\x07\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			geom.PointZM{1, 2, 3, 4},
			[]byte("\x00\x00\x00\x0b\xb9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\xb9\x0b\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@"),
		},
		{
			geom.LineString{[]geom.Point{{1, 2}, {3, 4}}},
			[]byte("\x00\x00\x00\x00\x02\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\x02\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@"),
		},
		{
			geom.LineStringZ{[]geom.PointZ{{1, 2, 3}, {4, 5, 6}}},
			[]byte("\x00\x00\x00\x03\xea\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\xea\x03\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		},
		{
			geom.LineStringM{[]geom.PointM{{1, 2, 3}, {4, 5, 6}}},
			[]byte("\x00\x00\x00\x07\xd2\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\xd2\x07\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		},
		{
			geom.LineStringZM{[]geom.PointZM{{1, 2, 3, 4}, {5, 6, 7, 8}}},
			[]byte("\x00\x00\x00\x0b\xba\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\xba\x0b\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @"),
		},
		{
			geom.Polygon{[][]geom.Point{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}},
			[]byte("\x00\x00\x00\x00\x03\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\x03\x00\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@"),
		},
		{
			geom.PolygonZ{[][]geom.PointZ{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}},
			[]byte("\x00\x00\x00\x03\xeb\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\xeb\x03\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			geom.PolygonM{[][]geom.PointM{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}},
			[]byte("\x00\x00\x00\x07\xd3\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\xd3\x07\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			geom.PolygonZM{[][]geom.PointZM{{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {1, 2, 3, 4}}}},
			[]byte("\x00\x00\x00\x0b\xbb\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00@$\x00\x00\x00\x00\x00\x00@&\x00\x00\x00\x00\x00\x00@(\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00"),
			[]byte("\x01\xbb\x0b\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00$@\x00\x00\x00\x00\x00\x00&@\x00\x00\x00\x00\x00\x00(@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@"),
		},
	}

	for _, tc := range testCases {

		// test XDR reading
		if got, err := Unmarshal(tc.xdr); err != nil || !reflect.DeepEqual(got, tc.g) {
			t.Errorf("Unmarshal(%q) == %q, %s, want %q, nil", tc.xdr, got, err, tc.g)
		}

		// test XDR writing
		if got, err := Marshal(tc.g, XDR); err != nil || !reflect.DeepEqual(got, tc.xdr) {
			t.Errorf("Marshal(%q, %q) == %q, %q, want %q, nil", tc.g, XDR, got, err, tc.xdr)
		}

		// test NDR reading
		if got, err := Unmarshal(tc.ndr); err != nil || !reflect.DeepEqual(got, tc.g) {
			t.Errorf("Unmarshal(%q) == %q, %s, want %q, nil", tc.ndr, got, err, tc.g)
		}

		// test NDR writing
		if got, err := Marshal(tc.g, NDR); err != nil || !reflect.DeepEqual(got, tc.ndr) {
			t.Errorf("Marshal(%q, %q) == %q, %q, want %q, nil", tc.g, NDR, got, err, tc.ndr)
		}

	}

}