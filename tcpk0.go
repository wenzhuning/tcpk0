package main

import (
    "fmt"
    "net"
    "time"
)

func doTask() {
time.Sleep(1 * time.Second)
main()
}

func main() {
for {
    conn, err := net.DialTimeout("tcp", "223.5.5.5:80",time.Second*4)
    if err != nil {
      fmt.Println("请求失败", err)
      doTask()
    }else{
      fmt.Println("请求成功", conn)
    }

    time.Sleep(1*time.Second)

    conn0, err := net.DialTimeout("tcp", "223.5.5.5:80",time.Second*4)
    if err != nil {
      fmt.Println("请求失败", err)
      doTask()
    }else{ 
      fmt.Println("请求成功", conn0)
    }

    time.Sleep(2*time.Second)
}
}
