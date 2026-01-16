package types

import "os"

type File struct {
	Path string
	Info os.FileInfo
}

type SignatureType string

const (
	ExpressRoute   SignatureType = "express_route"
	ExpressSource  SignatureType = "express_source"
	HttpFetch      SignatureType = "http_fetch"
	DatabaseAccess SignatureType = "database_access"
)

type SignatureHit struct {
	Path          string
	LineNumber    int
	SignatureType SignatureType
}

type SignatureSpan struct {
	Path      string
	StartLine int
	EndLine   int
	SHA256    string
}

type ExpressRouteType string

const (
	ExpressEndpointGet    ExpressRouteType = "get"
	ExpressEndpointPost   ExpressRouteType = "post"
	ExpressEndpointPut    ExpressRouteType = "put"
	ExpressEndpointDelete ExpressRouteType = "delete"
)
