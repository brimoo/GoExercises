package main

import(
    "fmt"
    "bufio"
    "os"
)


func hamming(first, second string) int {

    dist := 0
    for i := 0; i < len(first); i++ {
        if first[i] != second[i] {
            dist++
        }
    }

    return dist
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter first strand: ")
    firstStrand, _ := reader.ReadString('\n')
    fmt.Print("Enter second strand: ")
    secondStrand, _ := reader.ReadString('\n')

    if len(firstStrand) != len(secondStrand) {
        fmt.Println("Strands must be of equal length!")
        os.Exit(0)
    }

    fmt.Println("Hamming Distance: ", hamming(firstStrand, secondStrand))
}
