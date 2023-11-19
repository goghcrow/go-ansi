package ansi

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestColor(t *testing.T) {
	{
		fmt.Println(Green.Text("?"))
		fmt.Println(Red.Text("!"))
		fmt.Println(Cyan.Light().Bg().Text("background"))
		fmt.Println(Yellow.Text("Hello").Underline().Bold().Bg(Blue))
	}

	{
		rainbow := Red.Text("H").Reset().
			Fg(Green).Text("E").Reset().
			Fg(Yellow).Text("L").Reset().
			Fg(Blue).Text("L").Reset().
			Fg(Purple).Text("O").Reset().
			Fg(Cyan).Text("!")
		fmt.Println(rainbow)
	}

	{
		rainbow := Red.Text("H").
			Append(Green.Text("E").Bold()).
			Append(Yellow.Text("L")).
			Append(Blue.Text("L").Bold()).
			Append(Purple.Text("O")).
			Append(Cyan.Text("!").Bold())
		fmt.Println(rainbow)
	}

	{
		colors := []Color{Red, Green, Yellow, Blue, Purple, Cyan}
		randColor := func() Color { return colors[rand.Intn(len(colors))] }

		xs := []string{
			"欲买桂花同载酒,终不似,少年游!",
			"醉后不知天在水，满船清梦压星河!",
		}

		for _, s := range xs {
			rb := &Ansi{}
			for i, c := range s {
				word := randColor().Text(string(c))
				if i%2 == 0 {
					rb = rb.Append(word.Bold())
				} else {
					rb = rb.Append(word)
				}
			}
			fmt.Println(rb)

			fmt.Println(Strip(rb.String()))
		}
	}
}
