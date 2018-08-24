/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package xkcd

// APICurrentComic is the URL of latest comic (use to get comic id)
const APICurrentComic = "https://xkcd.com/info.0.json"

// APIFormat is a string template for URL of comid with Id.
const APIFormat = "http://xkcd.com/%d/info.0.json"

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

// Comic struct is the serialization container for XKCD JSON files.
type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}
