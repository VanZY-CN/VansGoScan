package myflag

import (
	"fmt"
)

func Banner() {
	banner := `
	______  _____  _______ _______ _______ __   _
	|  ____ |     | |______ |       |_____| | \  |
	|_____| |_____| ______| |_____  |     | |  \_|   
   
	   GoScan version:1.0
	   Author:Myth3me
		   `
	fmt.Println(banner)
}