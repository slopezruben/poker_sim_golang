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

    fmt.Println("Mazo Inicial")
    d.Shuffle();
    deck.ShowCards(d.GetCards());

    fmt.Println("Cartas Robadas")
    drawn := d.Draw(3)
    deck.ShowCards(drawn);

    fmt.Println("Cartas Restantes")
    deck.ShowCards(d.GetCards());
}
