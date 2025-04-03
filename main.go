package main
import "fmt"
func main (){


slice := []string {"shadow","sonic","superman","taylor swift","peter parker"}
doisprimeiros := slice[0:2]
fmt.Println(doisprimeiros)
doisultimos := slice[3:]
fmt.Println(doisultimos)
nomedomeio := slice[2]
fmt.Println(nomedomeio)

}
