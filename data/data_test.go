package data

import "testing"

func TestActionBuffer(t *testing.T){
	InitiateActionBuffer()
	if(cap(ActionBuffer)==0){
		t.Error("Echec Initialisation ActionBuffer")
	}
	AddNewAction(0,"test","test2","test3")
	if(ActionBuffer[0].Description["test"]["test2"]!="test3"){
		t.Error("Echec ajout ActionBuffer")
	}
	CleanActionBuffer()
	_,ok:=ActionBuffer[0].Description["test"]
	if ok{
		t.Error("Echec nettoyage ActionBuffer")
	}
}

func TestObjectID(t *testing.T){
	a:=1
	b:="test"
	c:=true
	IDMap=NewObjectID()
	ida:=(&IDMap).AddObject(&a)
	idb:=(&IDMap).AddObject(&b)
	(&IDMap).AddObject(&c)
	if((&IDMap).GetObjectFromID(ida)!=&a){
		t.Error("Echec GetObjectFromID")
	}
	if((&IDMap).GetIDFromObject(&b)!=idb){
		t.Error("Echec GetIDFromObject")
	}
}