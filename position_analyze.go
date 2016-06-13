package main

import (
    "encoding/xml"
    "fmt"
	"io/ioutil"
    "os"
)

type position struct {
    //code      string    `xml:"id,attr"`
    pos      string    `xml:"boundedBy>EnvelopeWithTimePeriod>lowerCorner"`
}

func main() {
	file, err := os.Open("/Users/matsui/Desktop/L01-10_GML/L01-10-g.xml") // For read access.        
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    defer file.Close()
    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
	v := new(position)
    err = xml.Unmarshal(data, &v)
    if err != nil {
        fmt.Println("XML Unmarshal error:", err)
        return
    }
	fmt.Printf("%+v", v)
    //fmt.Println(v[0].code)
    //fmt.Println(v[0].pos)
    //fmt.Println(v[1].code)
    //fmt.Println(v[2].pos)
}
