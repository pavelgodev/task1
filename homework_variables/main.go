// Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
// Ми маємо 23 грн.
// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
// 2. Скільки груш ми можемо купити?
// 3. Скільки яблук ми можемо купити?
// 4. Чи ми можемо купити 2 груші та 2 яблука?

package main

import "fmt"

func main(){
    const budget float32 = 23
    var applePrice float32 = 5.99
    var pearPrice float32 = 7
    var check float32
    var quantityPear, quantityApple int
    var isPossibleCheck bool

    // 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
    check = 9 * applePrice + 8 * pearPrice

    // 2. Скільки груш ми можемо купити?
    quantityPear =  int( budget / pearPrice )

    // 3. Скільки яблук ми можемо купити?
    quantityApple =  int( budget / applePrice )

    // 4. Чи ми можемо купити 2 груші та 2 яблука?
    isPossibleCheck = ( 2 * pearPrice + 2 * applePrice ) <= budget

    fmt.Println( "Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?", check )
    fmt.Println( "Скільки груш ми можемо купити?", quantityPear )
    fmt.Println( "Скільки яблук ми можемо купити?", quantityApple )
    fmt.Println( "Чи ми можемо купити 2 груші та 2 яблука?", isPossibleCheck )

}