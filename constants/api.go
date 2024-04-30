package constants

import "golang.org/x/exp/maps"

var (
	ArtAPI = map[string]string{
		"cn": "",
		"tw": "",
		"vn": "",
	}

	API = maps.Keys(ArtAPI)
)
