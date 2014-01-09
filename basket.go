package potter

type Basket []*Book

var prices = [5]float64 {8.0, 15.2, 21.6, 25.6, 30}

func (basket *Basket) get(name int) (*Book, bool) {
  for _, book := range *basket {
    if book.name == name {
      return book, true
    }
  }
  return nil, false
}

func (basket *Basket) countActiveBooks() int {
  count := 0
  for _, book := range *basket {
    if book.active() == true {
      count++
    }
  }
  return count
}

func (basket *Basket) countUniqueBooks() int {
  count := 0
  for _, book := range *basket {
    if book.unique() == true {
      count++
    }
  }
  return count
}

func (basket *Basket) incFirstUniqueBook() {
  for _, book := range *basket {
    if book.unique() == true {
      book.count++
      return
    }
  }
}

func (basket *Basket) decBooks() {
  for _, book := range *basket {
    book.count--
  }
}

func (basket *Basket) price() float64 {
  price := 0.0
  for basket.countActiveBooks() > 0 {
    len := basket.countActiveBooks()
    if len == 5 && basket.countUniqueBooks() == 2 {
      price += prices[len-2]
      basket.incFirstUniqueBook()
    } else {
      price += prices[len-1]
    }
    basket.decBooks()
  }

  return price
}

