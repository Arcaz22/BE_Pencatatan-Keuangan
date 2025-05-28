package constant

const (
    HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
    HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
    HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
    HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
    HeaderAccessControlMaxAge           = "Access-Control-Max-Age"

    AllowOriginAll     = "*"
    AllowMethods       = "POST, GET, PUT, DELETE, OPTIONS"
    AllowHeadersBasic  = "Content-Type, Authorization"
    AllowCredentials   = "true"
    MaxAge            = "86400"
)
