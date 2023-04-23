package gots_sdk

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinRouterGenerator struct {
	R   *gin.Engine
	Gen *SdkGenerator
}

func (gen *GinRouterGenerator) Group(relativePath string, handlers ...gin.HandlerFunc) *Group {
	g := gen.R.Group(relativePath, handlers...)

	return &Group{
		Path: relativePath,
		G:    g,
		Gen:  gen.Gen,
	}
}

func (gen *GinRouterGenerator) Generate(fname string) {
	gen.Gen.Generate(fname)
}

type Group struct {
	Path string
	G    *gin.RouterGroup
	Gen  *SdkGenerator
}

func (group *Group) POST(relativePath string, query interface{}, body interface{}, res interface{}, handlers ...gin.HandlerFunc) gin.IRoutes {
	path := group.Path + relativePath
	group.Gen.CreateTsFunc(path, http.MethodPost, query, body, res)
	return group.G.POST(relativePath, handlers...)
}

func (group *Group) GET(relativePath string, query interface{}, body interface{}, res interface{}, handlers ...gin.HandlerFunc) gin.IRoutes {
	path := group.Path + relativePath
	group.Gen.CreateTsFunc(path, http.MethodGet, query, body, res)
	return group.G.GET(relativePath, handlers...)
}

func (group *Group) PUT(relativePath string, query interface{}, body interface{}, res interface{}, handlers ...gin.HandlerFunc) gin.IRoutes {
	path := group.Path + relativePath
	group.Gen.CreateTsFunc(path, http.MethodPut, query, body, res)
	return group.G.PUT(relativePath, handlers...)
}

func (group *Group) DELETE(relativePath string, query interface{}, body interface{}, res interface{}, handlers ...gin.HandlerFunc) gin.IRoutes {
	path := group.Path + relativePath
	group.Gen.CreateTsFunc(path, http.MethodDelete, query, body, res)
	return group.G.DELETE(relativePath, handlers...)
}

func NewGinRouterGenerator(router *gin.Engine) *GinRouterGenerator {

	return &GinRouterGenerator{
		R:   router,
		Gen: NewSdkGenerator(),
	}
}
