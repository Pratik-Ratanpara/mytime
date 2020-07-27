package mytime 
import "fmt"
import "time"

func Date() {
t := time.Now()
fmt.Println(t.Day(),"/",t.Month(),"/",t.Year())
}

func Time() {
t := time.Now()
fmt.Println(t.Hour(),"hr",t.Minute(),"min",t.Second(),"sec")
}