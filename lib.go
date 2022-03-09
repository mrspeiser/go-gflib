package gfl

import (
  "fmt"
  "math/rand"
  "time"
)

const alnums = "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWYXZ1234567890"
const alphabet = "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWYXZ"

func Greet(name string){
  fmt.Printf("Hello %s!",name)
}

func Getalphabet() string {
  return alphabet;
}

func Getalnums() string {
  return alnums;
}

func SetSeed(){
  /*
    Seed uses the provided seed value to initialize the default Source to a deterministic state. 
    If Seed is not called, the generator behaves as if seeded by Seed(1). 
    Seed values that have the same remainder when divided by 2³¹-1 
    generate the same pseudo-random sequence. 
    Seed, unlike the Rand.Seed method, is safe for concurrent use.
  */
  rand.Seed(time.Now().UTC().UnixNano())
}

func UnsignedInt(max int) int {
  m := max
  if m < 1 {
     m = 1
  }
  Int := rand.Intn(m)
  return Int
}

func PositiveUnsignedInt(max int) int {
  Int := UnsignedInt(max)
  if Int == 0 {
    Int = Int+1
  }
  return Int
}

func RandomAlnum(length int) string {
  s := make([]byte,length)
  for i := range s {
    s[i] = alnums[UnsignedInt(len(alnums))]
  }
  return string(s)
}

func ArrayofStringMaps(length int) ([]map[string]string) {
  arr := make([]map[string]string,length)
  return arr
}

func ArrayofIntMaps(length int) ([]map[string]int) {
  arr := make([]map[string]int,length)
  return arr
}

func IndexOf(vs []string, s string) int {
  for i, v := range vs {
    if v == s {
      return i
    }
  }
  return -1
}

func Include(vs []string, s string) bool {
  return IndexOf(vs, s) >= 0
}

func Any(vs []string, f func(string) bool) bool {
  for _,v := range vs {
    if f(v) {
      return true
    }
  }
  return false
}

func All(vs []string, f func(string) bool) bool {
  for _,v := range vs {
    if !f(v) {
      return false
    }
  }
  return true
}

func Filter(vs []string, f func(string) bool) []string {
  nvs := make([]string,0)
  for _,v := range nvs {
    if f(v) {
      nvs = append(nvs,v)
    }
  }
  return nvs
}

func Map(vs []string, f func(string) string) []string {
  nvs := make([]string, len(vs))
  for i,v := range nvs {
    nvs[i] = f(v)
  }
  return nvs
}
