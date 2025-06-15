package utils

import (
    "gorm.io/gorm"
    "pencatatan_keuangan/pkg/response"
    "strings"
)

// ApplyCaseInsensitiveFilters menerapkan filter pagination, sorting, dan filtering yang tidak sensitif terhadap huruf besar/kecil
func ApplyCaseInsensitiveFilters(db *gorm.DB, params response.PaginationParams, searchableFields []string) *gorm.DB {
    query := db

    if params.Search != "" && len(searchableFields) > 0 {
        search := "%" + strings.ToLower(params.Search) + "%"

        var conditions []string
        var args []interface{}

        for _, field := range searchableFields {
            conditions = append(conditions, "LOWER("+field+") LIKE ?")
            args = append(args, search)
        }

        query = query.Where(strings.Join(conditions, " OR "), args...)
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
            if strValue, ok := value.(string); ok {
                query = query.Where("LOWER("+fieldName+") LIKE ?", "%"+strings.ToLower(strValue)+"%")
            } else {
                query = query.Where(fieldName+" LIKE ?", "%"+value.(string)+"%")
            }
        } else {
            if strValue, ok := value.(string); ok {
                query = query.Where("LOWER("+field+") = ?", strings.ToLower(strValue))
            } else {
                query = query.Where(field+" = ?", value)
            }
        }
    }

    if params.SortBy != "" {
        query = query.Order(params.SortBy + " " + params.SortDir)
    }

    return query
}
