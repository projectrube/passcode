package main

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/eiannone/keyboard"
)

func main() {
	nodeRedUrl := "http://192.168.86.181:1880/"

	secretOne := "917"
	secretTwo := "512"
	secretThree := "888"
	stepOneStatus := false
	stepTwoStatus := false
	stepThreeStatus := false

	attempt := ""

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		if key == keyboard.KeyEsc {
			break
		} else {
			if key == keyboard.KeyEnter {

				switch attempt {
					case secretOne:
						displayMessage(nodeRedUrl + "step-1")
						stepOneStatus = true
						attempt = ""
					case secretTwo:
						displayMessage(nodeRedUrl + "step-2")
						stepTwoStatus = true
						attempt = ""
					case secretThree:
						displayMessage(nodeRedUrl + "step-3")
						stepThreeStatus = true
						attempt = ""
					default:
						fmt.Println("Nope, try again!")
						displayMessage(nodeRedUrl + "nope")
						attempt = ""
				}

				if (stepOneStatus && stepTwoStatus && stepThreeStatus) {
					displayMessage(nodeRedUrl + "success")
				}

			} else {
				attempt += string(char)
			}
		}
	}
}

func displayMessage(step string) {
	resp, err := http.Get(step)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
