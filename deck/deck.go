package deck

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/judedaryl/go-arrayutils"
) 

type card struct {
	nominal_value int
	value         string
	suit          string
    ToString func() string
}

var suits []string = []string{"Diamonds", "Clubs", "Hearts", "Spades"}

type deck struct {
	cards [52]card
    GetCards func() []card
    Shuffle func();
}

func NewCard(value int, suit string) (card, error) {
	c := card{nominal_value: value, suit: suit}

    if(!arrayutils.Some(suits, func(s string) bool { return s == suit})) {
        return c, errors.New(fmt.Sprintf("%v is not a suit", suit))
    }

    
    c.ToString = func() string {return fmt.Sprintf("%v of %v", c.value, c.suit)};
    

    switch value {
        case 1:
            c.value = "Ace"
            break;
        case 11:
            c.value = "Jack"
            break;
        case 12:
            c.value = "Queen"
            break;
        case 13:
            c.value = "King"
            break;
        default:
            c.value = strconv.Itoa(value);
            break;
    }

    return c, nil;
}


func CreateDeck() *deck {
    d := deck{}
    d.GetCards = func() []card { return d.cards[:]}
    d.Shuffle = func() {
        for i := 0; i < 52; i++ {
            j := i + (rand.Int() % (52 -i));   
            d.cards[i], d.cards[j] = d.cards[j], d.cards[i];
        }
    }

    return &d 
}

func StartDeck(d *deck) error {
    var index int = 0;
    for _, s := range suits {
        for i:=1 ; i < 14; i++ {
          var card, err = NewCard(i, s);

          if(err != nil) {
            return err
          }

          d.cards[index] = card;
          index++;
        }
    } 
    return nil
}
