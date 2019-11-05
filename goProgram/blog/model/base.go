package model

type BaseObject interface {
	GetClass() interface{}
	GetSliceClass() interface{}
}
