package gfl

import (
  "fmt"
  "math/rand"
  "time"
  "os"
)

const alnums = "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWYXZ1234567890"
const alphabet = "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWYXZ"

func Greet(name string){
  fmt.Printf("Hello %s!",name)
}

func GetAlphabet() string {
  return alphabet;
}

func GetAlnums() string {
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
  /*
    rand.Intn() returns value that includes 0
  */
  return Int
}

func PositiveUnsignedInt(max int) int {
  Int := UnsignedInt(max)
  if Int == 0 {
    /*
      if UnsignedInt returns value that is 0, add by 1
    */
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


/*
  File Open-time Flags

  // Only One of these MUST be specified

  O_RDONLY - open the file read-only
  O_WRONLY - open the file write-only
  O_RDWR   - open the file read-write

  // These other flags can be or'ed in to control behavior

  O_APPEND - append new data to the file when writing
  O_CREATE - create a new file if none exists
  O_EXCL   - used with O_CREATE, file must not exist
  O_SYNC   - open for synchronous I/O
  O_TRUNC  - truncate regular writable file when opened
     e.g. os.OpenFile(filename, os.O_RDWR | O_APPEND | O_CREATE, 0744)

*/

func WriteFile(filename string, s string) int {
  fn := filename
  s1 := []byte(s)
  err := os.WriteFile(fn,s1,0777)
  if err != nil {
    fmt.Printf("Error writing file: %s",err)
    return 1
  }
  return 0
}

func ChmodFile(filename string, mod uint32) int {
  fn := filename
  var mode = os.FileMode(mod)
  err := os.Chmod(fn, mode)
  if err != nil {
    fmt.Printf("Error Chmod on file file: %s; err: %s",fn,err)
    return 1
  }
  return 0
}

func OpenFileReadOnly(filename string) (file *os.File) {
  fn := filename
  /*
    os.OpenFile(filename, safety_flag, permissions)
  */
  f,err := os.OpenFile(fn,os.O_RDONLY,755)
  if err != nil {
    fmt.Printf("Error opening file: %s",err)
    panic(err)
  }
  return f
}

func OpenFileReadWrite(filename string) (file *os.File) {
  fn := filename
  /*
    os.OpenFile(filename, safety_flag, permissions)
  */
  f,err := os.OpenFile(fn,os.O_RDONLY,755)
  if err != nil {
    fmt.Printf("Error opening file: %s",err)
    panic(err)
  }
  return f
}

func WriteToFile(file *os.File, content []byte) int {
  f := file
  data := content
  //out, err := f.Write(data)
  if _, err := f.Write(data); err != nil {
    fmt.Printf("WriteToFile Error! %s",err)
  }
  return 0
}

func AddIntegers(num1 int, num2 int) int{
    sum := num1 + num2
    return sum
}

func SubtractIntegers(num1 int, num2 int) int{
    difference := num1 - num2
    return difference
}
