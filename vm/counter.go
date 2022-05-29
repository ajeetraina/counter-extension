package main 
  
import "fmt"
import "io/ioutil"
import "os"
import "path/filepath"
import "strconv"
import "time"
  
func main() { 
	timestamp := strconv.FormatInt(time.Now().UnixMicro(), 10)
	os.Mkdir("button_pressed", os.FileMode(0522))

	// We create a new file everytime this program is invoked
	file, _ := os.Create(filepath.Join("button_pressed", timestamp))
	defer file.Close()

	files,_ := ioutil.ReadDir("button_pressed")
    numFiles := len(files)
	fmt.Println(numFiles)
}
