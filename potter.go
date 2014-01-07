package potter

func Price(books []int) float64 {
  basket   := build(books)

  return basket.price()
}

func build(books []int) Basket {
  basket := Basket{}

  for _, value := range books {
    if book, present := basket.get(value); present == true {
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

