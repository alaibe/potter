package potter

type Book struct {
  name, count int
}

func (book Book) unique() bool {
  return book.count == 1
}

func (book Book) active() bool {
  return book.count >= 1
}
