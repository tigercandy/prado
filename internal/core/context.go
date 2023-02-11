package core

import "github.com/gin-gonic/gin"

type Handler func(c Context)

var _ Context = (*context)(nil)

type Context interface {
}

type context struct {
	ctx *gin.Context
}
