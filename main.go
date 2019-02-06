package main

import "fmt"
import npc "server/npc"
import carte "server/carte"

func main() {
	mat:=carte.New(4)
	carte.Debug(mat)
	fmt.Println("test")
	npc.Test()
}