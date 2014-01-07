package potter

import "testing"

func TestBasics(t *testing.T){
  expected := float64(0)
  books := []int{}
  result := Price(books)
  assert_equal(expected, result, t)

  expected = float64(8)
  books = []int{0}
  result = Price(books)
  assert_equal(expected, result, t)

  expected = float64(16)
  books = []int{0, 0}
  result = Price(books)
  assert_equal(expected, result, t)
}

func TestSimpleDiscounts(t *testing.T){
  expected := 8 * 2 * 0.95
  books := []int{0, 1}
  result := Price(books)
  assert_equal(expected, result, t)

  expected = 8 * 3 * 0.9
  books = []int{0, 1, 2}
  result = Price(books)
  assert_equal(expected, result, t)

  expected = 8 * 4 * 0.8
  books = []int{0, 1, 2, 3}
  result = Price(books)
  assert_equal(expected, result, t)

  expected = 8 * 5 * 0.75
  books = []int{0, 1, 2, 3, 4}
  result = Price(books)

  assert_equal(expected, result, t)
}

func TestSeveralDiscounts(t *testing.T){
  expected := 8 + (8 * 2 * 0.95)
  books := []int{0, 0, 1}
  result := Price(books)
  assert_equal(expected, result, t)

  expected = 2 * (8 * 2 * 0.95)
  books = []int{0, 0, 1, 1}
  result = Price(books)
  assert_equal(expected, result, t)

  expected = (8 * 4 * 0.8) + (8 * 2 * 0.95)
  books = []int{0, 0, 1, 2, 2, 3}
  result = Price(books)
  assert_equal(expected, result, t)

  expected = 8 + (8 * 5 * 0.75)
  books = []int{0, 1, 1, 2, 3, 4}
  result = Price(books)
  assert_equal(expected, result, t)
}

func TestEdgeCases(t *testing.T){
  expected := 2 * (8 * 4 * 0.8)
  books := []int{0, 0, 1, 1, 2, 2, 3, 4}
  result := Price(books)
  assert_equal(expected, result, t)

  expected = 3 * (8 * 5 * 0.75) + 2 * (8 * 4 * 0.8)
  books = []int{0, 0, 0, 0, 0,
                 1, 1, 1, 1, 1,
                 2, 2, 2, 2,
                 3, 3, 3, 3, 3,
                 4, 4, 4, 4}
  result = Price(books)
  assert_equal(expected, result, t)
}

func assert_equal(expected float64, result float64, t *testing.T){
  if expected == result {
    t.Log(".")
  } else {
    t.Errorf("%v is not equal to %v", expected, result)
  }
}

