//go:build !release
// +build !release

package consts

const (
	DbConnInfo = "host=localhost port=5432 user=postgres password= dbname=storiesservice sslmode=disable"
	IsRelease  = false
)
