package global

//go:generate stringer -type=Source
type Source int

const (
	SourceBangumi Source = iota
	SourceTMDB
	SourceDouban
	SourceOther
)

//go:generate stringer -type=Type
type Type int

const (
	TypeAnime Type = iota
	TypeMovie
	TypeGame
	TypeBook
	TypeMusic
)
