package main

import (
	"fmt"
	"os/exec"
	"runtime"
)
func main() {
    if(runtime.GOOS == "windows"){
		fmt.Println("windows machine")
	}else{
		fmt.Println(runtime.GOOS)
		out,err:=exec.Command("ls","-ltr").Output()
		if err!= nil {
			fmt.Println(err)
		}
		output:=string(out[:])
		fmt.Println(output)
	}
}