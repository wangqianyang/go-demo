package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Goods struct {
	ID    string
	Name  string
	Stock int
}

func connect(cName string) (*mgo.Session, *mgo.Collection) {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/") //Mongodb's connection
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	//return a instantiated collect
	return session, session.DB("local").C(cName)
}

/**
添加数据
*/
func (a *Goods) save() error {
	s, c := connect("goods")
	defer s.Close()
	a.ID = bson.NewObjectId().Hex()
	return c.Insert(&a)
}

/**
查询所有数据
*/
func (a Goods) findAll() ([]Goods, error) {
	s, c := connect("goods")
	defer s.Close()
	var group []Goods
	err := c.Find(nil).All(&group)
	return group, err
}

/**
根据id查找数据
*/
func (a *Goods) find(id string) error {
	s, c := connect("goods")
	defer s.Close()
	return c.Find(bson.M{"id": id}).One(&a)
}

/**
删除
*/
func (a Goods) delete() error {
	s, c := connect("goods")
	defer s.Close()
	return c.Remove(bson.M{"id": a.ID})
}

/**
更新
*/
func (a *Goods) update() error {
	s, c := connect("goods")
	defer s.Close()
	c.Update(bson.M{"id": a.ID}, a)
	return a.find(a.ID)
}


