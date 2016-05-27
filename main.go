package main

import (
  "os"
  "bufio"
  "log"
  "sync"
  "countwords/sorter"
  "fmt"
)

var wg sync.WaitGroup

func main() {

  // Take the args
  filenames := os.Args[1:]

  // Create a wait group with the amount of arguments
  wg.Add(len(filenames))

  // Loop over the filenames
  for _, file := range filenames {
    go readFile(file)
  }

  // Wait for execution to stop
  wg.Wait()
}


// Read the contents of a file
func readFile(filename string) {
  log.Println("Opening file: " + filename)
  var words map[string]int
  words = make(map[string]int)

  f, err := os.Open(filename)

  if err != nil {
    log.Println(err)
    return
  }

  scanner := bufio.NewScanner(f)

  // Set the Split method to ScanWords.
  scanner.Split(bufio.ScanWords)

  // Scan all words from the file.
  for scanner.Scan() {
    word := scanner.Text()
    words[word]++
  }

  list := sorter.Sort(words)

  log.Print("Top 3 (or less) common words in " + filename + " are: ")

  for i, item := range list {
    if i > 2 {
      break;
    } else {
      fmt.Printf("%s:%d\n", item.Word, item.Occurances )
    }
  }

  wg.Done()
}

