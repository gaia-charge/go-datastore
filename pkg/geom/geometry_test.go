package geom_test

import (
	"testing"

	"github.com/gaia-charge/go-datastore/pkg/geom"
	"github.com/paulmach/orb/geojson"
)

func TestNullGeometry(t *testing.T) {
	t.Run("Null", func(t *testing.T) {
		g := geom.NullGeometry4326{}
		err := g.Scan(nil)

		if err != nil {
			t.Errorf("Scan: %v", err)
		}

		if err != nil {
			t.Errorf("UnmarshalGeometry: %v", err)
		}

		if g.Valid != false {
			t.Errorf("Valid: %v", g.Valid)
		}
	})

	t.Run("Point", func(t *testing.T) {
		g := geom.NullGeometry4326{}
		err := g.Scan([]byte(`\x0101000020e6100000e0d57267266e4840b22ac24d46b50240`))

		if err != nil {
			t.Errorf("Scan: %v", err)
		}

		p, err := geojson.UnmarshalGeometry([]byte(`{
			"type": "Point",
			"coordinates": [48.860547, 2.338513]
		}`))

		if err != nil {
			t.Errorf("UnmarshalGeometry: %v", err)
		}

		if g.Valid != true {
			t.Errorf("Valid: %v", g.Valid)
		}

		if p.Coordinates != g.Geometry4326.Coordinates {
			t.Errorf("Coordinates: %v %v", p.Coordinates, g.Geometry4326.Coordinates)
		}
	})
}

func TestGeometry(t *testing.T) {
	t.Run("Null", func(t *testing.T) {
		g := geom.Geometry4326{}
		err := g.Scan(nil)

		if err == nil {
			t.Errorf("Scan: %v", err)
		}
	})

	t.Run("Point", func(t *testing.T) {
		g := geom.Geometry4326{}
		err := g.Scan([]byte(`\x0101000020e6100000e0d57267266e4840b22ac24d46b50240`))

		if err != nil {
			t.Errorf("Scan: %v", err)
		}

		p, err := geojson.UnmarshalGeometry([]byte(`{
			"type": "Point",
			"coordinates": [48.860547, 2.338513]
		}`))

		if err != nil {
			t.Errorf("UnmarshalGeometry: %v", err)
		}

		if p.Coordinates != g.Coordinates {
			t.Errorf("Coordinates: %v %v", p.Coordinates, g.Coordinates)
		}
	})
}
