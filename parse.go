package s3syncparse

import (
	"strings"
)

// An Upload represents a file that was synced.
type Upload struct {
	Src  string
	Dest string
	Path string
}

// Parse a sync string and return the uploads
func Parse(cont string) *Upload {
	s := strings.Split(cont, " ")
	length := len(s)

	if length != 4 {
		return nil
	}

	src := s[1]
	dest := s[3]

	// Parse the dest to get the path
	dp := strings.Split(dest, "/")
	pp := dp[3:]
	path := "/" + strings.Join(pp, "/")

	upload := Upload{
		Src:  src,
		Dest: dest,
		Path: path,
	}

	return &upload
}

func ParseAll(cont string) []Upload {
	var uploads []Upload

	p := strings.Split(cont, "\n")

	for _, item := range p {
		upload := Parse(item)

		if upload != nil {
			uploads = append(uploads, *upload)
		}

	}

	return uploads
}
