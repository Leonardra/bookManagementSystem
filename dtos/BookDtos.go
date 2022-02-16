package dtos

type BookRequestDto struct{
	Author string
	Title string
	Genre string
	BookDescription string
	Tags []string
	Quantity int
	ReleaseDate string
}