package constant

const (
    ErrTypeValidation = "VALIDATION_ERROR"
    ErrTypeBusiness   = "BUSINESS_ERROR"
    ErrTypeSystem     = "SYSTEM_ERROR"
    ErrTypeNotFound   = "NOT_FOUND_ERROR"
    ErrTypeAuth       = "AUTH_ERROR"
)

const (
    CategoryTypeIncome  = "income"
    CategoryTypeExpense = "exepense"
)

var ValidCategoryTypes = []string{
    CategoryTypeIncome,
    CategoryTypeExpense,
}
