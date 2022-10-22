package controller

import "fmt"

func Run(power <-chan bool) {
	for {
		select {
		case isPowerOn := <-power:
			{
				var powerOnString string
				if isPowerOn {
					powerOnString = "ON"
				} else {
					powerOnString = "OFF"
				}
				fmt.Printf("recieved mesage: %s", powerOnString)
				if !isPowerOn {
					break
				}
			}
		default:
			{
				// continue
			}
		}
		fmt.Printf("doing work\n")
	}

}
