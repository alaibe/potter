package potter

import "fmt"

func price(books []int) float64 {
  sortedBooks := organize(books)
  discount    := calc(sortedBooks)
  
  return discount
}

func organize(books []int) map[int]int {
  sortedBooks := make(map[int]int)
  
  for index, value := range books {
    if _, present := sortedBooks[value]; present == false {
      sortedBooks[value] = 1
    }
    
    if index + 1 == len(books) {
      break
    }
    
    if next := books[index+1]; value == next {
      sortedBooks[value] += 1
    }
  }
  
  return sortedBooks
}

func calc(sortedBooks map[int]int) float64 {
  sum := 0.0
  fmt.Println(sortedBooks)
  for len(sortedBooks) > 0 {
        
    switch len(sortedBooks) {
      case 1:
        sum += 8.0
      case 2:
        sum += 15.2
      case 3:
        sum += 21.6
      case 4:
        sum += 25.6
      case 5:
        sum += 30.0
    }
    
    decrementBooks(sortedBooks)
  }
  
  return sum
}

func decrementBooks(sortedBooks map[int]int) map[int]int {
  for key, value := range sortedBooks {
    sortedBooks[key] -= 1
    if value == 1 {
      delete(sortedBooks, key)
    }
  }
  
  return sortedBooks
}

func bestPriceFor(sortedBooks map[int]int) float64 {
  countOne := 0
  for _, value := range sortedBooks {
    if value == 1 {
      countOne += 1
    }
  }
  if countOne >= 2 {
    return 25.6
  } else {
    return 30.0
  }
}
