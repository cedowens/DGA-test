
package main

import (
    "bufio"
    "fmt"
    "os"
		"net"
)

func main() {

    readFile, err := os.Open("dg-test-domains.com")

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
			iprecords, _ := net.LookupIP(fileScanner.Text())
			if err !=nil {
				fmt.Println(os.Stderr,"Could not get IPs: %v\n",err)
			}
			for _, ip:= range iprecords{
				fmt.Println(fileScanner.Text(),ip)
			}
        //fmt.Println(fileScanner.Text())
    }

    readFile.Close()
}







// func main() {
//
// f, err := os.Open("dg-test-domains.com")
// if err != nil {
//     fmt.Printf("error opening file: %v\n",err)
//     os.Exit(1)
// }
// r := bufio.NewReader(f)
// s, e := Readln(r)
// for e == nil {
//     fmt.Println(s)
//     s,e = Readln(r)
// }
//
// 	//iprecords, _ := net.LookupIP("facebook.com")
// 	//for _, ip := range iprecords {
// 	//	fmt.Println(ip)
// 	//}
// }
