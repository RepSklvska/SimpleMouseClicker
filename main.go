// This tool can be very useful when you have too many tokens or oranges in 100% Orange Juice!

package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	chPressedCtrl := make(chan bool)
	go func() {
		pressed := robotgo.AddEvents("ctrl")
		if pressed {
			chPressedCtrl <- true
		}
	}()

	fmt.Println("Move the mouse, press CTRL to set a point.")
	var x, y int
	func() {
		for {
			x, y = robotgo.GetMousePos()
			fmt.Print("\rPos: (", x, ",", y, ")")

			select {
			case _ = <-chPressedCtrl:
				fmt.Print("\rPos: (", x, ",", y, ") selected\n")
				return
			default:
				continue
			}
		}
	}()
	fmt.Print("Set a interval time between clicks (ms): ")
	var inter int
	_, _ = fmt.Scanln(&inter)
	fmt.Print("Clicking (", x, ",", y, ") with interval:", inter, "milliseconds.\n")

	chPressedCtrlAgain := make(chan bool)
	go func() {
		pressed := robotgo.AddEvents("ctrl")
		if pressed {
			chPressedCtrlAgain <- true
		}
	}()
	i := 1
	func() {
		for {
			i++
			fmt.Print("\rPress CTRL to quit (clicked ", i, " times)")
			robotgo.Move(x, y)
			robotgo.Click("left")
			time.Sleep(time.Duration(inter) * time.Millisecond)
			select {
			case _ = <-chPressedCtrlAgain:
				fmt.Println("\nQuited by pressing CTRL.")
				return
			default:
				continue
			}
		}
	}()
}
