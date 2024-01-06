package global

//go:generate stringer -type=Source
type Source int

const (
	SourceNil Source = iota
	SourceBangumi
	SourceTMDB
	SourceDouban
)

//go:generate stringer -type=Type
type Type int

const (
	TypeNil Type = iota
	TypeAnime
	TypeMovie
	TypeGame
	TypeBook
	TypeMusic
)

//go:generate stringer -type=Status
type Status int

const (
	StatusNil Status = iota
	StatusWish
	StatusWatched
	StatusWatching
	StatusOnHold
	StatusDropped
)
