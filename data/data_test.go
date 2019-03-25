package data

import (
	"testing"

	"git.unistra.fr/AOEINT/server/constants"
)

func TestActionBuffer(t *testing.T) {
	constants.PlayerUID1 = "A912HF18H129HF"
	constants.PlayerUID2 = "9H91HF91FHFH1J"
	InitiateActionBuffer()
	if cap(ActionBuffer["A912HF18H129HF"]) == 0 {
		t.Error("Echec Initialisation ActionBuffer")
	}
	AddNewAction(constants.PlayerUID1, 0, "test", "test2", "test3")
	if ActionBuffer[constants.PlayerUID1][0].Description["test"]["test2"] != "test3" {
		t.Error("Echec ajout ActionBuffer")
	}
	CleanActionBuffer()
	_, ok := ActionBuffer[constants.PlayerUID1][0].Description["test"]
	if ok {
		t.Error("Echec nettoyage ActionBuffer")
	}
}

func TestObjectID(t *testing.T) {
	a := 1
	b := "test"
	c := true
	IDMap = NewObjectID()
	ida := (&IDMap).AddObject(&a)
	idb := (&IDMap).AddObject(&b)
	(&IDMap).AddObject(&c)
	if (&IDMap).GetObjectFromID(ida) != &a {
		t.Error("Echec GetObjectFromID")
	}
	if (&IDMap).GetIDFromObject(&b) != idb {
		t.Error("Echec GetIDFromObject")
	}
}
