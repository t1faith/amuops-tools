package main

import (
	"fmt"
)

func main() {
	transfer := "\t ________   _______  _______  __  ___ .______        ______   ______        _______.___________.\n	|       /  |   ____||   ____||  |/  / |   _  \\      /      | /  __  \\      /       |           |\n	`---/  /   |  |__   |  |__   |  '  /  |  |_)  |    |  ,----'|  |  |  |    |   (----`---|  |----`\n	   /  /    |   __|  |   __|  |    <   |      /     |  |     |  |  |  |     \\   \\       |  |     \n	  /  /----.|  |____ |  |____ |  .  \\  |  |\\  \\----.|  `----.|  `--'  | .----)   |      |  |     \n	 /________||_______||_______||__|\\__\\ | _| `._____| \\______| \\______/  |_______/       |__|     \n																									"
	fmt.Println(transfer)
	var t = []byte(transfer) //强制类型转换

	fmt.Println(t)
}
