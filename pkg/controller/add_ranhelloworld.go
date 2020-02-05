package controller

import (
	"github.com/bsuro10/ran-hw-operator/pkg/controller/ranhelloworld"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, ranhelloworld.Add)
}
