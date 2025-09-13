package netscape

import (
	"path"
	"strings"
)

func joinURLPath(a, b string) string {
	ap := strings.TrimSuffix(a, "/")
	bp := strings.TrimPrefix(b, "/")
	return path.Clean(ap + "/" + bp)
}
