package main

import (
	"fmt"
	"reflect"
)

type controllerInfo struct {
	url            string
	controllerType reflect.Type
}

type ControllerRegistor struct {
	routers []*controllerInfo
}

type ControllerInterface interface {
	Do()
}

type DefaultController struct {
}

type AddController struct {
}

type DeleteController struct {
}

type UpdateController struct {
}

type SearchController struct {
}

func (d *DefaultController) Do() {
	fmt.Println("in default controller")
}

func (u *AddController) Do() {
	fmt.Println("mongodb add something")
}

func (u *DeleteController) Do() {
	fmt.Println("mongodb delete something")
}

func (d *SearchController) Do() {
	fmt.Println("mongodb search something")
}

func (d *UpdateController) Do() {
	fmt.Println("mongodb update something")
}