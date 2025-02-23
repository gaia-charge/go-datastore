package util

import (
	"database/sql"
	"time"

	"github.com/gaia-charge/go-datastore/pkg/geom"
)

func SqlNullBool(i interface{}) sql.NullBool {
	n := sql.NullBool{}

	switch t := i.(type) {
	case *bool:
		if t == nil {
			n.Scan(nil)
		} else {
			n.Scan(*t)
		}
	default:
		n.Scan(t)
	}

	return n
}

func SqlNullFloat64(i interface{}) sql.NullFloat64 {
	n := sql.NullFloat64{}

	switch t := i.(type) {
	case *float64:
		if t == nil {
			n.Scan(nil)
		} else {
			n.Scan(*t)
		}
	default:
		n.Scan(t)
	}

	return n
}

func SqlNullGeometry4326(i interface{}) geom.NullGeometry4326 {
	n := geom.NullGeometry4326{}

	switch t := i.(type) {
	case *geom.Geometry4326:
		if t == nil {
			n.Scan(nil)
		} else {
			n.Scan(*t)
		}
	default:
		n.Scan(t)
	}

	return n
}

func SqlNullInt32(i interface{}) sql.NullInt32 {
	n := sql.NullInt32{}

	switch t := i.(type) {
	case *int32:
		if t == nil {
			n.Scan(nil)
		} else {
			n.Scan(*t)
		}
	default:
		n.Scan(t)
	}

	return n
}

func SqlNullZeroInt64(i interface{}) sql.NullInt64 {
	n := SqlNullInt64(i)

	if n.Int64 == 0 {
		n.Valid = false
	}

	return n
}

func SqlNullInt64(i interface{}) sql.NullInt64 {
	n := sql.NullInt64{}

	switch t := i.(type) {
	case *int64:
		if t == nil {
			n.Scan(nil)
		} else {
			n.Scan(*t)
		}
	default:
		n.Scan(t)
	}

	return n
}

func SqlNullString(i interface{}) sql.NullString {
	n := sql.NullString{}

	switch t := i.(type) {
	case *string:
		if t == nil {
			n.Scan(nil)
		} else {
			n.Scan(*t)
		}
	default:
		n.Scan(t)
	}

	return n
}

func SqlNullTime(i interface{}) sql.NullTime {
	n := sql.NullTime{}

	switch t := i.(type) {
	case *time.Time:
		if t == nil {
			n.Scan(nil)
		} else {
			n.Scan(*t)
		}
	default:
		n.Scan(t)
	}

	return n
}
