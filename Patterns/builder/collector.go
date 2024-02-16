package main

type Film struct {
	Studio   string
	Year     int
	IsSeries bool
	Actors   []string
}

type FilmBuilderI interface {
	SetStudio(studio string) FilmBuilderI
	SetYear(year int) FilmBuilderI
	SetIsSeries(isSeries bool) FilmBuilderI
	SetActors(actors []string) FilmBuilderI
	GetFilm() Film
}

type FilmBuilder struct {
	Studio   string
	Year     int
	IsSeries bool
	Actors   []string
}

func NewFilmBuilder() FilmBuilderI {
	return FilmBuilder{}
}

func (f FilmBuilder) SetStudio(studio string) FilmBuilderI {
	f.Studio = studio
	return f
}

func (f FilmBuilder) SetYear(year int) FilmBuilderI {
	f.Year = year
	return f
}

func (f FilmBuilder) SetIsSeries(isSeries bool) FilmBuilderI {
	f.IsSeries = isSeries
	return f
}

func (f FilmBuilder) SetActors(actors []string) FilmBuilderI {
	f.Actors = actors
	return f
}

func (f FilmBuilder) GetFilm() Film {
	return Film{
		Studio:   f.Studio,
		Year:     f.Year,
		IsSeries: f.IsSeries,
		Actors:   f.Actors,
	}
}
