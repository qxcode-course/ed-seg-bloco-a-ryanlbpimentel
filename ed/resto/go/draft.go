package main
import "fmt"

func Div_Resto(n int){
    if n/2 != 0 {
        Div_Resto(n/2)
    }

    fmt.Printf("%d %d\n", n/2, n%2)
}

func main() {
    var valor int
    fmt.Scan(&valor)
    
    Div_Resto(valor)
}
