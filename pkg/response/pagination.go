package response

import (
    "math"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type PaginationMeta struct {
    CurrentPage  int `json:"current_page"`
    PerPage      int `json:"per_page"`
    TotalPages   int `json:"total_pages"`
    TotalRecords int `json:"total_records"`
}

type PaginatedResponse struct {
    Message    string         `json:"message,omitempty"`
    Data       interface{}    `json:"data,omitempty"`
    Pagination PaginationMeta `json:"pagination"`
}

type PaginationParams struct {
    Page    int    `form:"page" binding:"min=1"`
    Limit   int    `form:"limit" binding:"min=1,max=100"`
    Offset  int

    SortBy  string `form:"sort_by"`
    SortDir string `form:"sort_dir"`

    Search  string                 `form:"search"`
    Filters map[string]interface{}
}

func SendPaginated(c *gin.Context, statusCode int, message string, data interface{}, params PaginationParams, totalRecords int) {
    totalPages := int(math.Ceil(float64(totalRecords) / float64(params.Limit)))

    c.JSON(statusCode, PaginatedResponse{
        Message: message,
        Data:    data,
        Pagination: PaginationMeta{
            CurrentPage:  params.Page,
            PerPage:      params.Limit,
            TotalPages:   totalPages,
            TotalRecords: totalRecords,
        },
    })
}

func SuccessPaginated(c *gin.Context, message string, data interface{}, params PaginationParams, totalRecords int) {
    SendPaginated(c, http.StatusOK, message, data, params, totalRecords)
}

func NewPaginationParams(c *gin.Context) PaginationParams {
    params := PaginationParams{
        Page:    1,
        Limit:   10,
        SortDir: "desc",
        Filters: make(map[string]interface{}),
    }

    c.ShouldBindQuery(&params)

    if params.Page < 1 {
        params.Page = 1
    }
    if params.Limit < 1 {
        params.Limit = 10
    } else if params.Limit > 100 {
        params.Limit = 100
    }

    params.SortDir = strings.ToLower(params.SortDir)
    if params.SortDir != "asc" && params.SortDir != "desc" {
        params.SortDir = "desc"
    }

    params.Offset = (params.Page - 1) * params.Limit

    queryParams := c.Request.URL.Query()
    standardParams := map[string]bool{
        "page":     true,
        "limit":    true,
        "sort_by":  true,
        "sort_dir": true,
        "search":   true,
    }

    for key, values := range queryParams {
        if !standardParams[key] && len(values) > 0 {
            params.Filters[key] = values[0]
        }
    }

    return params
}

func ApplyFilters(db *gorm.DB, params PaginationParams, searchableFields []string) *gorm.DB {
    query := db

    if params.Search != "" && len(searchableFields) > 0 {
        search := "%" + params.Search + "%"

        query = query.Where(buildSearchQuery(searchableFields), generateSearchArgs(search, len(searchableFields))...)
    }

    for field, value := range params.Filters {
        if strings.HasSuffix(field, "_gt") {
            fieldName := strings.TrimSuffix(field, "_gt")
            query = query.Where(fieldName+" > ?", value)
        } else if strings.HasSuffix(field, "_lt") {
            fieldName := strings.TrimSuffix(field, "_lt")
            query = query.Where(fieldName+" < ?", value)
        } else if strings.HasSuffix(field, "_gte") {
            fieldName := strings.TrimSuffix(field, "_gte")
            query = query.Where(fieldName+" >= ?", value)
        } else if strings.HasSuffix(field, "_lte") {
            fieldName := strings.TrimSuffix(field, "_lte")
            query = query.Where(fieldName+" <= ?", value)
        } else if strings.HasSuffix(field, "_like") {
            fieldName := strings.TrimSuffix(field, "_like")
            query = query.Where(fieldName+" LIKE ?", "%"+value.(string)+"%")
        } else {
            query = query.Where(field+" = ?", value)
        }
    }

    if params.SortBy != "" {
        query = query.Order(params.SortBy + " " + params.SortDir)
    }

    return query
}

func buildSearchQuery(fields []string) string {
    var conditions []string

    for _, field := range fields {
        conditions = append(conditions, field+" LIKE ?")
    }

    return strings.Join(conditions, " OR ")
}

func generateSearchArgs(searchTerm string, count int) []interface{} {
    args := make([]interface{}, count)
    for i := range args {
        args[i] = searchTerm
    }
    return args
}

func ApplyPaginationToQuery(db *gorm.DB, params PaginationParams) *gorm.DB {
    return db.Offset(params.Offset).Limit(params.Limit)
}

func CountTotalRecords(db *gorm.DB, model interface{}) (int, error) {
    var count int64
    err := db.Model(model).Count(&count).Error
    return int(count), err
}
