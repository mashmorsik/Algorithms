package main

import "fmt"

func main() {
	filmBuilder := NewFilmBuilder()

	GOT := filmBuilder.SetStudio("HBO").SetYear(2011).SetIsSeries(true).
		SetActors([]string{"Peter Dinklage", "Emilia Clarke", "Kit Harington", "Sophie Turner"}).BuildFilm()
	fmt.Println(GOT)

	HomeAlone := filmBuilder.SetStudio("20th Century").SetYear(1990).SetIsSeries(false).
		SetActors([]string{"Macaulay Culkin", "Joe Pesci", "Daniel Stern", "Catherine O'Hara"}).BuildFilm()
	fmt.Println(HomeAlone)
}
