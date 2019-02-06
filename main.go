package main

import "fmt"
import npc "server/npc"
<<<<<<< HEAD
import carte "server/carte"
=======
>>>>>>> 9ce59af5765c970d6d6a9b586473c8361207d84c

func main() {
	mat:=carte.New(4)
	carte.Debug(mat)
	fmt.Println("test")
	npc.Test()
}
