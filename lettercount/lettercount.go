/* Note: we could read the file names and call the routines in the same loop, avoiding
   an extra loop and storing the file names in a slce. However, the purpose here
   is to demonstrate concurrency and the time it takes for a user to enter a file
   name would ruin this since each routine would finish before the next file name
   could be entered
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var fileCount int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the number of files to be read: ")
	fmt.Scanf("%d", &fileCount)

	if fileCount < 1 {
		fmt.Println("Please enter a value of at least 1")
		os.Exit(0)
	}

	files := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fmt.Print("Enter name of file ", i+1, " (make sure it's in the package directory): ")
		file, _ := reader.ReadString('\n')
		file = file[:len(file)-1]
		files[i] = file
	}

	fmt.Print("Enter letter to search for: ")
	letter, _ := reader.ReadString('\n')
	letter = letter[:len(letter)-1]

	wg.Add(fileCount)
	for i := 0; i < len(files); i++ {
		go letterCount(files[i], letter)
	}

	wg.Wait()
}

func letterCount(file, letter string) {
	defer wg.Done()
	pwd, _ := os.Getwd()
	fileBytes, err := ioutil.ReadFile(pwd + "/" + file)

	if err != nil {
		fmt.Println(err)
		return
	}

	fileString := string(fileBytes)

	count := 0
	for i := 0; i < len(fileString); i++ {
		if fileString[i] == letter[0] {
			count++
		}
	}

	fmt.Println("Occurrences of ", letter, " in ", file, ": ", count)
}
