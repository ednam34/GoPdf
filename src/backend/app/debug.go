package Application

import "fmt"

func (a *App) PrintAny(an any) {
	fmt.Println(an)
}

func (a *App) PrintString(s string) {
	fmt.Println(s)
}
func (a *App) PrintArr(s []any) {
	fmt.Println(s)
}
