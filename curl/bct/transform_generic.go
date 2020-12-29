// +build !amd64 appengine gccgo

package bct

import (
	"github.com/loveandpeople-DAG/goClient/curl"
)

func transform(lto, hto, lfrom, hfrom *[curl.StateSize]uint, rounds uint) {
	transformGeneric(lto, hto, lfrom, hfrom, rounds)
}
