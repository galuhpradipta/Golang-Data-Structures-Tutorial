package main

import (
  "fmt"
  "os"
)

type node struct {
  next *node
  val string
  timesSeen int
}

type table struct {
  list []*node
  length int
}

func createTable(newLen int) *table {
  return &table{
    list: make([]*node, newLen),
    length : newLen,
  }
}

func findSpot(length int, myVal string) int {
    var insertLoc int
    for _, i := range myVal {
      insertLoc += int(i)
    }
    return insertLoc % length
}

func insertNode(myNode *node, myVal string) *node {
  newNode := &node{
    next : nil,
    val : myVal,
  }
  if myNode != nil{
    for newNode.next = myNode; newNode.next != nil; newNode.next = newNode.next.next {
      if newNode.next.val == myVal{
        newNode.next.timesSeen++
        return nil
      }
    }
  }
  return newNode
}

func getLength(vals []*node) int {
  var count int
  for _, pointer := range vals {
    if pointer != nil{
      count++
    }
  }
  return count
}


func reSize(myTable *table) *table {
  newTable := createTable(myTable.length * 2)
  for i := range myTable.list{
    for item := myTable.list[i]; item != nil; item = item.next{
      newTable = insertTable(newTable, item.val)
    }
  }
  return newTable
}

func insertTable(myTable *table, myVal string) *table {
  if getLength(myTable.list) == (myTable.length - 1) {
    myTable = reSize(myTable)
  }else{
    location := findSpot(myTable.length, myVal)
    newNode := insertNode(myTable.list[location], myVal)
    if newNode != nil{
      myTable.list[location] = newNode
    }
  }
  return myTable
}

func printTable(myTable *table){
  for i := 0; i < len(myTable.list); i++{
    fmt.Println(i)
    for temp := myTable.list[i]; temp != nil; temp = temp.next{
      fmt.Printf("%s %d", temp.val, temp.timesSeen)

    }
    fmt.Printf("\n")
  }
}

func main() {
  myTal := createTable(5)

  for _, word := range os.Args[1:] {
    myTal = insertTable(myTal, word)
  }
  printTable(myTal)
  fmt.Println("Hello World")

  fmt.Println(myTal)
}
