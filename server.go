package main
 
import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "sort"
    "strings"
    
    

)

func wordCount(str string) map[string]int {
    wordList := strings.Fields(str)
    counts := make(map[string]int)
    for _, word := range wordList {
        _, ok := counts[word]
        if ok {
            counts[word] += 1
        } else {
            counts[word] = 1
        }
    }
    return counts
}


 
func main() {
 	
    fileName := "book.txt"

    bs, err := ioutil.ReadFile(fileName)

    if err != nil {

        log.Fatal(err)
    }

    text := string(bs)

    fields := strings.FieldsFunc(text, func(r rune) bool {

        return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '\'')
    })

    wordsCount := make(map[string]int)

    for _, field := range fields {

        wordsCount[field]++
    }

    keys := make([]string, 0, len(wordsCount))

    for key := range wordsCount {

        keys = append(keys, key)
    }

    sort.Slice(keys, func(i, j int) bool {

        return wordsCount[keys[i]] > wordsCount[keys[j]]
    })
 
	
    var theArray [11]string
    var keyArray [11]int
 
    for idx, key := range keys {

        fmt.Printf("%s %d\n", key, wordsCount[key])
        theArray[idx] = key
        keyArray[idx] = wordsCount[key]
 
        if idx == 10 {
            break
        }
    }   

    

    // API routes
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "top ten used of word in a book \n %s %d \n %s %d \n %s %d \n %s %d \n %s %d \n %s %d \n %s %d \n %s %d \n %s %d", theArray[0], keyArray[0], 
			theArray[1], keyArray[1], theArray[2], keyArray[2], theArray[3], keyArray[3], 
			theArray[4], keyArray[4], theArray[5], keyArray[5], theArray[6], keyArray[6],
			theArray[7], keyArray[7], theArray[8], keyArray[8], theArray[9], keyArray[9])
			
	
    
    })
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })
 
    port := ":5000"
    fmt.Println("Server is running on port" + port)
 
    // Start server on port specified above
    log.Fatal(http.ListenAndServe(port, nil))

    

}
 
