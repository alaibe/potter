package potter

import "fmt"

type Book struct {
  name, count int
}

type Basket []*Book

var prices = [5]float64 {8.0, 15.2, 21.6, 25.6, 30}

func Price(books []int) float64 {
  basket   := build(books)

  return basket.Price()
}

func build(books []int) Basket {
  basket := Basket{}

  for _, value := range books {
    if book, present := basket.Get(value); present == true {
      book.count += 1
    } else {
      book := new(Book)
      book.name = value
      book.count = 1
      basket = append(basket, book)
    }
  }
  return basket
}

func (basket Basket) Get(name int) (*Book, bool) {
  for _, book := range basket {
    if book.name == name {
      return book, true
    }
  }
  return nil, false
}

func (basket Basket) countBooks() int {
  count := 0
  for _, book := range basket {
    if book.count > 0 {
      count += 1
    }
  }

  return count
}

func (basket Basket) countUniqueBooks() int {
  count := 0
  for _, book := range basket {
    if book.count == 0 {
      count += 1
    }
  }

  return count
}

func (basket Basket) incrementCounterWithOne() {
  for _, book := range basket {
    if book.count == 1 {
      book.count += 1
      break
    }
  }
}

func (basket Basket) Price() float64 {
  price := 0.0

  for basket.countBooks() > 0 {
    len := basket.countBooks()
    if basket.countUniqueBooks() == 2 {
      price += prices[len-2]
      basket.incrementCounterWithOne()
    } else {
      price += prices[len-1]
    }
    basket.decrementCounters()
  }

  return price
}

func (basket Basket) decrementCounters() {
  for _, book := range basket {
    book.count -= 1
  }
}

func bestPriceFor(sortedBooks map[int]int) {
  fmt.Println("test")
}
