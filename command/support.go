package command

import (
	"lsf"
)

// REVU: ALL can be env methods TODO

// assert resource does not exist
func _assertNotExists(env *lsf.Environment, docId string) {
	doc, e := env.LoadDocument(docId)
	if e == nil && doc != nil {
		panic(lsf.ERR.ResourceExists("docId:", docId))
	}
}

// assert resource does not exist
func _assertExists(env *lsf.Environment, docId string) {
	doc, e := env.LoadDocument(docId)
	if e != nil || doc == nil {
		panic(lsf.ERR.ResourceDoesNotExist("docId:", docId))
	}
}
