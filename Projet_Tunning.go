package main
import (
	"fmt"
)

func main () {
	x := 67
	y := 45
	z := 650
	w := 0
	for x<z {
		x+=y
		w++
	}
	fmt.Println("Jhon va devoir attendre", w/4, "mois et", w%4, "une semaine avant de finir sa 205 tunning")
}