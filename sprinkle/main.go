package main

import "bufio"
import "fmt"
import "math/rand"
import "os"
import "strings"
import "time"

const (
	add   = "*"
	input = "a"
)

var transforms = []string{
	input + add,
	add + input}

func main() {
	//	var (
	//		readfile string
	//	)
	//
	//	flag.StringVar(&readfile, "f", "hoge.txt", "read file")
	//
	//	flag.Parse()
	//
	//	//Read File
	//	inputFile, err := os.Open(readfile)
	//	if err != nil {
	//		fmt.Fprint(os.Stderr, err)
	//		os.Exit(1)
	//	}

	//	file := bufio.NewReader(inputFile)

	var addition = []string{"app", "lets", "go", "get"}

	//	for {
	//		s, err := file.ReadString('\n')
	//		if err == io.EOF {
	//			break
	//		}
	//		addition = append(addition, strings.Replace(s, "\n", "", -1))
	//	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		str := fmt.Sprintf(strings.Replace(t, input, s.Text(), -1))
		fmt.Println(strings.Replace(str, add, addition[rand.Intn(len(addition))], -1))
	}
}
