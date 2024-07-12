package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int) []int {
    // Write your code here
    multipl05 := []int {40,45,50,55,60,65,70,75,80,85,90,95,100}
    max := 38 ;
   // var res []int ;
    
    for i := 0; i < len(grades); i++ {
        //fmt.Printf("\n%d",arr[i]);
        if grades[i] >= max {
            curent := 0;
            for j := 0;  j < len(multipl05); j++ {
                if grades[i]+2 == multipl05[j] {
                    fmt.Printf("\n%d",multipl05[j]); 
                    curent = multipl05[j] ;
                    grades[i] = multipl05[j];
                }
            }
            
            if curent == 0 {
            fmt.Printf("\n%d",grades[i]);
            grades[i] = grades[i] ;  
            }
                
            //fmt.Printf("\n%d",arr[i]);
        }else {
             fmt.Printf("\n%d",grades[i]);
             grades[i] =grades[i] ;
        }
        
    }
    
    return grades ;

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    gradesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var grades []int

    for i := 0; i < int(gradesCount); i++ {
        gradesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        gradesItem := int(gradesItemTemp)
        grades = append(grades, gradesItem)
    }

    result := gradingStudents(grades)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
