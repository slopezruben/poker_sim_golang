package main;

import ( 
    "fmt"
    "evilbalatro/deck"
)

func main() {
    fmt.Println("Hello Evil Balatro")

    d := deck.CreateDeck();
    err := deck.StartDeck(d);

    if err != nil {
       fmt.Println(err);
    }

    for i := 0; i < 52; i++ {
        c := d.GetCards();
        fmt.Println(c[i].ToString())
    }

    d.Shuffle();

    for i := 0; i < 52; i++ {
        c := d.GetCards();
        fmt.Println(c[i].ToString())
    }
}
