package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("terminating steamcmd...")
	cmd := exec.Command("pkill", "-9", "steamcmd")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("steamcmd killing error!\nplease check is it running?:\n%v\n", err)
		return
	}
	fmt.Println("killed steamcmd!")
	os.Exit(0)
}
