package library

import (
	"fmt"
	"sort"
)

type User struct {
  Name string 
  Age int
}

type UserSlice []User

func (s UserSlice) Len() int {
  return len(s)
}

func (s UserSlice) Less(i, j int) bool {
  return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}

func Sort() {
  users := []User{
    {"Eko", 30},
    {"Budi", 45},
    {"Bedu", 55},
    {"Janur", 23},
  } 

  sort.Sort(UserSlice(users))
  fmt.Println(users)
}