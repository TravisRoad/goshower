package global

//go:generate stringer -type=Source
type Source int

const (
	Bangumi Source = iota
	TMDB
	Douban
)

//go:generate stringer -type=Type
type Type int

const (
	Anime Type = iota
	Movie
	Game
)
