package constant

const (
	HeaderContentType    = "Content-Type"
	HeaderAuthorization  = "Authorization"
)

const (
	ContentTypeJSON      = "application/json"
	ContentTypeForm      = "application/x-www-form-urlencoded"
	ContentTypeMultipart = "multipart/form-data"
	ContentTypeText      = "text/plain"
	ContentTypeHTML      = "text/html"
)

const (
    ErrTypeValidation = "VALIDATION_ERROR"
    ErrTypeBusiness   = "BUSINESS_ERROR"
    ErrTypeSystem     = "SYSTEM_ERROR"
    ErrTypeNotFound   = "NOT_FOUND_ERROR"
    ErrTypeAuth       = "AUTH_ERROR"
)

const (
    ErrCodeUnauthorized       = "UNAUTHORIZED"
    ErrCodeForbidden         = "FORBIDDEN"
    ErrCodeBadRequest        = "BAD_REQUEST"
    ErrCodeNotFound          = "NOT_FOUND"
    ErrCodeInternalError     = "INTERNAL_ERROR"
    ErrCodeValidation        = "VALIDATION_ERROR"
    ErrCodeDatabaseError     = "DATABASE_ERROR"
    ErrCodeDuplicateEntry    = "DUPLICATE_ENTRY"
    ErrCodeInvalidCredential = "INVALID_CREDENTIAL"
    ErrCodeExpiredToken      = "EXPIRED_TOKEN"
    ErrCodeInvalidToken      = "INVALID_TOKEN"
)

const (
    MsgInternalError     = "Terjadi kesalahan sistem"
    MsgNotFound         = "Data tidak ditemukan"
    MsgUnauthorized     = "Tidak memiliki akses"
    MsgInvalidInput     = "Input tidak valid"
    MsgDuplicateEntry   = "Data sudah ada"
    MsgInvalidToken     = "Token tidak valid"
    MsgExpiredToken     = "Token sudah kadaluarsa"
	MsgInvalidCredential = "Email atau password salah"
)

const (
	MsgRegisterSuccess = "Registrasi berhasil"
	MsgLoginSuccess    = "Login berhasil"
	MsgRetrivedUserSuccess = "Berhasil mendapatkan profile"
)
