package main

import "fmt"
import "unicode/utf8"

func main()  {

  var intNum int //var/db/  
  fmt.Println(intNum)

  var floatNum float64 // or float 32
  fmt.Println(floatNum) // no podemos mutiplcar direntes tipos 

  var numero float32 = 32.0
  var numero2 int32 = 32
  var result float32 = numero + float32(numero2)
  fmt.Println(result)
  

  // esto es para sacar el modulo

  var letras string = "Hello Atutui"
  fmt.Println(letras)

  fmt.Println(len(letras))
  fmt.Println(utf8.RuneCountInString("Y"))

  var runas rune = 'a'
  fmt.Println(runas)

  // DEfault values for al data types is "0"
  var1 , var2 := 1,2
  fmt.Println(var1,var2)

  const myConstante string = "const values"
  fmt.Println(myConstante)

  const pi = 3.1415
}
