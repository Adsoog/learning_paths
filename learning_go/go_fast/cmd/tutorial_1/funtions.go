package main


import (
    "fmt"
    "errors"
  )

func main()  {
  var dividendo int = 15
  var divisor int = 4

  var resultado, modulo, err = division(dividendo,divisor)

  if err != nil {
    fmt.Printf(err.Error())
  }else if modulo == 0{
    fmt.Printf("El resultado es: %v", resultado)  
  }else{
  fmt.Printf("El resultado es: %v y el residuo es: %v", resultado,modulo)
  }
}


func division(dividendo int, divisor int) (int,int,error) {

  var err error
  
  if divisor == 0 {
    err = errors.New("Ningun numero puede dividirse por 0")
    return 0,0, err
  }


  var cociente int= dividendo / divisor
  var residuo int = dividendo % divisor
  return cociente, residuo, err

}
