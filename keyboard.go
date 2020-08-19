package main

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/eiannone/keyboard"
)

func main() {
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
		} else if key == keyboard.KeyEnter {

			resp, err := http.Get("http://192.168.86.181:1880/hello")
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
	}
}
