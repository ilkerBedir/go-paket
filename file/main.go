package main

import	(
	"log"
	"os"
)
func main() {
	tempDir,err:=os.MkdirTemp("","cars-")
	if err!=nil{
        log.Fatal(err)
    }
	defer os.RemoveAll(tempDir)

	file, err:= os.CreateTemp(tempDir,"car-*.png")
	if err!=nil{
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	log.Println(file.Name())
	if _, err := file.Write([]byte("hello world\n")); err != nil {
        log.Println(err)
    }
	data, err := os.ReadFile(file.Name())
	if err!= nil {
		log.Println(err)
	}
	log.Println(string(data))
}