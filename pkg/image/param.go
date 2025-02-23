package image

import "github.com/gaia-charge/go-datastore/pkg/db"

func NewUpdateImageParams(image db.Image) db.UpdateImageParams {
	return db.UpdateImageParams(image)
}
