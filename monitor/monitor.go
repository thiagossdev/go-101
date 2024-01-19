package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const monitoringTimes = 10
const monitoring = 5

func main() {
	displayIntroduction()

	for {
		displayMenu()
		command := requestCommandInput()

		switch command {
		case 1:
			startMonitor()
		case 2:
			fmt.Println("Displaying logs...")
			fmt.Println()
		case 0:
			fmt.Println("Exiting the program...")
			os.Exit(0)
		default:
			fmt.Println("Undefined command!")
			os.Exit(-1)
		}
	}
}

func displayIntroduction() {
	name := "Thiago Sousa Santos"
	version := 0.1
	fmt.Println("Hello,", name)
	fmt.Println("This program is in version", version)
	fmt.Println("The type of name var is", reflect.TypeOf(name))
}

func displayMenu() {
	fmt.Println("1- Start monitor")
	fmt.Println("2- Display logs")
	fmt.Println("0- Exit the program")
}

func requestCommandInput() int {
	var command int
	fmt.Scanf("%d", &command)
	fmt.Println("\nThe chose command was", command)
	fmt.Println()

	return command
}

func startMonitor() {
	fmt.Println("Monitoring...")
	sites := []string{"https://app.revgas.com/api/health-check", "https://random-status-code.herokuapp.com",
		"https://httpbin.org/status/500", "https://httpbin.org/status/201"}

	for i := 0; i < monitoringTimes; i++ {
		fmt.Println("Run ", i+1)
		for _, site := range sites {
			resp, _ := http.Get(site)

			if resp.StatusCode == 200 {
				fmt.Println("Site:", site, "has been loaded successfully!")
			} else {
				fmt.Println("Site:", site, "has problems. Status Code:", resp.StatusCode)
			}
		}
		time.Sleep(monitoring * time.Second)
		fmt.Println()
	}

	fmt.Println()
}
