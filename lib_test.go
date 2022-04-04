package gfl_test

import (
  "testing"
  "fmt"
  "github.com/mrspeiser/go-gflib"
)

func TestGreet(t *testing.T){
  gfl.Greet("Matthew")
}

func TestUnsignedInt(t *testing.T){
  gfl.SetSeed()
  a := gfl.UnsignedInt(100)
  b := gfl.UnsignedInt(100)
  fmt.Printf("\nrandom integer: %d",a)
  fmt.Printf("\nrandom integer: %d\n",b)
}

func TestRandomAlnum(t *testing.T){
  gfl.SetSeed()
  uia := gfl.UnsignedInt(16)
  uib := gfl.UnsignedInt(16)
  uic := gfl.UnsignedInt(16)

  fmt.Printf("\nrandom integer: %d",uia)
  fmt.Printf("\nrandom integer: %d",uib)
  fmt.Printf("\nrandom integer: %d\n",uic)

  a := gfl.RandomAlnum(uia)
  b := gfl.RandomAlnum(uib)
  c := gfl.RandomAlnum(uic)

  fmt.Printf("\nrandom alnum: %s",a)
  fmt.Printf("\nrandom alnum: %s",b)
  fmt.Printf("\nrandom alnum: %s\n",c)

  x := gfl.RandomAlnum(24)
  y := gfl.RandomAlnum(24)
  z := gfl.RandomAlnum(24)

  fmt.Printf("\nrandom alnum 24 chars: %s",x)
  fmt.Printf("\nrandom alnum 24 chars: %s",y)
  fmt.Printf("\nrandom alnum 24 chars: %s\n",z)
}

func TestArrayMaps(t *testing.T){
  a := gfl.ArrayofStringMaps(10)
  b := gfl.ArrayofIntMaps(10)
  fmt.Printf("\n")
  fmt.Println(a)
  fmt.Println(b)
}

func TestWriteFile(t *testing.T){
  gfl.WriteFile("./testfile.txt","hello world 123!")
}

func TestChmodFile(t *testing.T){
  gfl.ChmodFile("./testfile.txt", 0667)
}

func TestOpenFileReadOnly(t *testing.T){
  f := gfl.OpenFileReadOnly("./testfile.txt")
  defer f.Close()

  stat,staterr := f.Stat()
  if staterr != nil {
    fmt.Println(staterr)
  }
  fmt.Println("")
  fmt.Println(stat)

  filesize := stat.Size()
  buffer := make([]byte, filesize)
  rd, readerr := f.Read(buffer)
  if readerr != nil {
    fmt.Println(readerr)
  }
  fmt.Println("\nbytes read: ", rd)
  fmt.Println("bytestream to string: ", string(buffer))

  d := []byte("Should not write")
  r, writeerr := f.Write(d)
  if writeerr != nil {
    fmt.Printf("Error! %s",writeerr)
  }
  fmt.Println(r)
}

func TestAddIntegers(t *testing.T){
    STest := gfl.AddIntegers(10,2)
    fmt.Printf("\nthe sum is: %d \n", STest)
}

func TestSubtractIntegers(t *testing.T){
    DTest := gfl.SubtractIntegers(10,2)
    fmt.Printf("the difference is: %d \n", DTest)
}

func TestMultiplyIntegers(t *testing.T){
  int1 := 411
  int2 := 122
  result := gfl.MultiplyIntegers(int1,int2)
  fmt.Printf("the product is: %d \n",result)
}

func TestDivideIntegers(t *testing.T){
  int1 := 400
  int2 := 101
  result := gfl.DivideIntegers(int1,int2)
  fmt.Printf("the quotient is: %d \n",result)
}

func TestBinaryExponent(t *testing.T){
  result := gfl.BinaryExponent(30)
  fmt.Printf("the result of 2^30 is: %d \n",result)
}

func TestMax32BitInteger(t *testing.T){
  result := gfl.BinaryExponent(32)
  max32int := gfl.DivideIntegers(result,2)
  fmt.Printf("the max positive 32 bit integer is: %d \n",max32int)
}

func TestFactorialRecursive (t *testing.T){
    result := gfl.FactorialRecursive(5)
    fmt.Printf("the factorial of 5 is: %d \n", result)
}

func TestPermutations(t *testing.T){
    result := gfl.Permutations(5,3)
    fmt.Printf("the total permutations of 5 options and 3 choices is: %d \n", result)
}

func TestCombinations(t *testing.T){
    result := gfl.Combinations(5,3)
    fmt.Printf("the total Combinations of 5 options and 3 choices is: %d \n", result)
}

func TestPermutationsWithRep(t *testing.T){
    result := gfl.PermutationsWithRep(5,3)
    fmt.Printf("the total permutations with repititions of 5 options and 3 choices is: %d \n", result)
}

func TestModifyStringMap(t *testing.T){
  mymap := make(map[string]string)
  key := "foo"
  value := "bar"
  fmt.Printf("\n")
  gfl.PrintMap(mymap)
  gfl.ModifyStringMap(&mymap,key,value)
  gfl.PrintMap(mymap)

}
