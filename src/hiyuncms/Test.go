package main

import (
  "github.com/go-vgo/robotgo"
   "fmt"
)


func main() {
	keve := robotgo.AddEvent("k")
	if keve == 0 {
		fmt.Println("you press...", "k")
	}

	mleft := robotgo.AddEvent("mleft")
	if mleft == 0 {
		fmt.Println("you press...", "mouse left button")
	}
}
/*
func main() {
	robotgo.TypeString("Hello World")
	robotgo.KeyTap("enter")
	robotgo.TypeString("en")
	robotgo.KeyTap("i", "alt", "command")
	arr := []string{"alt", "command"}
	robotgo.KeyTap("i", arr)

	robotgo.WriteAll("测试")
	text, err := robotgo.ReadAll()
	if err == nil {
		fmt.Println(text)
	}
}
*/
/*
func main() {
	fpid, err := robotgo.FindIds("chrome")
	if err == nil {
		fmt.Println("pids...", fpid)
	}

	isExist, err := robotgo.PidExists(100)
	if err == nil {
		fmt.Println("pid exists is", isExist)
	}

	abool := robotgo.ShowAlert("test", "robotgo")
	if abool == 0 {
		fmt.Println("ok@@@", "ok")
	}

	title := robotgo.GetTitle()
	fmt.Println("title@@@", title)
}
*/
/*
func main() {
	robotgo.ScrollMouse(10, "up")
	robotgo.MouseClick("left", true)
	robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
}
*/
