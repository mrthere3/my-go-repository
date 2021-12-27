package main

import (
    "fmt"
    "os"
)

//func main() {
   // 新建文件
//   file, err := os.OpenFile("./慢慢.txt",os.O_APPEND,0666)
//   if err != nil {
//       fmt.Println(err)
//       return
//   }
//   defer file.Close()
//   for i := 0; i < 5; i++ {
//       file.WriteString("hello\n")
//       file.Write([]byte("cd\n"))
//   }
//}
func main() {
   file,err := os.OpenFile("./慢慢.txt",os.O_APPEND|os.O_CREATE,0666)
   if err != nil{
       fmt.Println(err)
       return

   }
   defer file.Close()
      for i := 0; i < 5; i++ {
      file.WriteString("hello\n")
      file.Write([]byte("cd\n"))
  }
}


