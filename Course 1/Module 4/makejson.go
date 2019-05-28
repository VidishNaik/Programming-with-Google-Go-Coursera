package main

import (
"fmt"
"bufio"
"os"
"strings"
"encoding/json"
)

func main() {
	m := make(map[string]string)	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	name, _ := reader.ReadString('\n')
	fmt.Println("Enter your address:")
	address, _ := reader.ReadString('\n')
	m["name"] = strings.TrimSpace(name)
	m["address"] = strings.TrimSpace(address)
	fmt.Println(m)
	json_map, _ := json.Marshal(m)
	fmt.Println(string(json_map))
}
