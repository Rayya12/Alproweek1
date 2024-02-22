package main
import "fmt"

func main(){
  var i,j,ukuran,n int
  fmt.Scanln(&ukuran)
  n = 1
  
  for i = 1 ; i <= ukuran ; i ++ {
    for j = 1 ; j <= n ; j++ {
      fmt.Print("* ")
    }
    
    if n < ukuran {
        n ++
    }
    
    fmt.Print("\n")
  }
}
