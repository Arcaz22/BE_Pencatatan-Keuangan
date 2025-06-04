package repository

import (
	"fmt"
	"pencatatan_keuangan/internal/service/dashboard/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DashboardRepository interface {
    GetTotalIncome(userID uuid.UUID, year int, month *int) (float64, error)
    GetTotalExpense(userID uuid.UUID, year int, month *int) (float64, error)
    GetTotalBudget(userID uuid.UUID, year int, month *int) (float64, error)
    // GetMonthlyIncomes(userID uuid.UUID, year int) ([]dto.MonthlyTotal, error)
    // GetMonthlyExpenses(userID uuid.UUID, year int) ([]dto.MonthlyTotal, error)
    GetExpenseDistribution(userID uuid.UUID, year int, month *int) ([]dto.CategoryDistribution, error)
    GetBudgetComparison(userID uuid.UUID, year int, month *int) ([]dto.BudgetComparison, error)
}

type dashboardRepository struct {
    db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
    return &dashboardRepository{db: db}
}

func (r *dashboardRepository) GetTotalIncome(userID uuid.UUID, year int, month *int) (float64, error) {
    var total float64
    query := r.db.Table("incomes").
        Select("COALESCE(SUM(amount), 0) as total").
        Where("user_id = ? AND EXTRACT(YEAR FROM date) = ?", userID, year)

    if month != nil {
        query = query.Where("EXTRACT(MONTH FROM date) = ?", *month)
    }

    err := query.Row().Scan(&total)
    return total, err
}

func (r *dashboardRepository) GetTotalExpense(userID uuid.UUID, year int, month *int) (float64, error) {
    var total float64
    query := r.db.Table("expenses").
        Select("COALESCE(SUM(amount), 0) as total").
        Where("user_id = ? AND EXTRACT(YEAR FROM date) = ?", userID, year)

    if month != nil {
        query = query.Where("EXTRACT(MONTH FROM date) = ?", *month)
    }

    err := query.Row().Scan(&total)
    return total, err
}

func (r *dashboardRepository) GetTotalBudget(userID uuid.UUID, year int, month *int) (float64, error) {
    var total float64
    query := r.db.Table("budgets").
        Select("COALESCE(SUM(amount), 0) as total").
        Where("user_id = ? AND is_active = true", userID)

    query = query.Where("(EXTRACT(YEAR FROM effective_from) <= ? AND EXTRACT(YEAR FROM effective_to) >= ?)",
        year, year)

    if month != nil {
        query = query.Where("(EXTRACT(YEAR FROM effective_from) < ? OR (EXTRACT(YEAR FROM effective_from) = ? AND EXTRACT(MONTH FROM effective_from) <= ?)) AND (EXTRACT(YEAR FROM effective_to) > ? OR (EXTRACT(YEAR FROM effective_to) = ? AND EXTRACT(MONTH FROM effective_to) >= ?))",
            year, year, *month, year, year, *month)
    }

    err := query.Row().Scan(&total)
    return total, err
}

// func (r *dashboardRepository) GetMonthlyIncomes(userID uuid.UUID, year int) ([]dto.MonthlyTotal, error) {
//     var results []dto.MonthlyTotal

//     err := r.db.Raw(`
//         SELECT
//             EXTRACT(MONTH FROM date)::int as month,
//             EXTRACT(YEAR FROM date)::int as year,
//             COALESCE(SUM(amount), 0) as total,
//             'IDR' as currency
//         FROM incomes
//         WHERE user_id = ? AND EXTRACT(YEAR FROM date) = ?
//         GROUP BY EXTRACT(MONTH FROM date), EXTRACT(YEAR FROM date)
//         ORDER BY year, month
//     `, userID, year).Scan(&results).Error

//     return results, err
// }

// func (r *dashboardRepository) GetMonthlyExpenses(userID uuid.UUID, year int) ([]dto.MonthlyTotal, error) {
//     var results []dto.MonthlyTotal

//     err := r.db.Raw(`
//         SELECT
//             EXTRACT(MONTH FROM date)::int as month,
//             EXTRACT(YEAR FROM date)::int as year,
//             COALESCE(SUM(amount), 0) as total,
//             'IDR' as currency
//         FROM expenses
//         WHERE user_id = ? AND EXTRACT(YEAR FROM date) = ?
//         GROUP BY EXTRACT(MONTH FROM date), EXTRACT(YEAR FROM date)
//         ORDER BY year, month
//     `, userID, year).Scan(&results).Error

//     return results, err
// }

func (r *dashboardRepository) GetExpenseDistribution(userID uuid.UUID, year int, month *int) ([]dto.CategoryDistribution, error) {
    var results []dto.CategoryDistribution

    query := `
        WITH total_expenses AS (
            SELECT COALESCE(SUM(e.amount), 0) as total
            FROM expenses e
            WHERE e.user_id = ? AND EXTRACT(YEAR FROM e.date) = ?
    `

    queryParams := []interface{}{userID, year}

    if month != nil {
        query += ` AND EXTRACT(MONTH FROM e.date) = ?`
        queryParams = append(queryParams, *month)
    }

    query += `)
        SELECT
            c.id as category_id,
            c.name as category_name,
            COALESCE(SUM(e.amount), 0) as amount,
            CASE
                WHEN (SELECT total FROM total_expenses) > 0
                THEN ROUND((COALESCE(SUM(e.amount), 0) / (SELECT total FROM total_expenses) * 100), 2)
                ELSE 0
            END as percentage
        FROM categories c
        LEFT JOIN expenses e ON c.id = e.category_id
            AND e.user_id = ?
            AND EXTRACT(YEAR FROM e.date) = ?
    `

    queryParams = append(queryParams, userID, year)

    if month != nil {
        query += ` AND EXTRACT(MONTH FROM e.date) = ?`
        queryParams = append(queryParams, *month)
    }

    query += `
        WHERE c.type = 'expense'
        GROUP BY c.id, c.name
        ORDER BY amount DESC
    `

    err := r.db.Raw(query, queryParams...).Scan(&results).Error
    return results, err
}

func (r *dashboardRepository) GetBudgetComparison(userID uuid.UUID, year int, month *int) ([]dto.BudgetComparison, error) {
    var results []dto.BudgetComparison

    query := `WITH months AS (`

    if month != nil {
        query += fmt.Sprintf(`SELECT %d AS month`, *month)
    } else {
        query += `SELECT generate_series(1, 12) AS month`
    }

    query += `),
        budget_monthly AS (
            SELECT
                m.month,
                COALESCE(SUM(b.amount), 0) as budget_amount
            FROM months m
            LEFT JOIN budgets b ON
                b.user_id = ? AND
                b.is_active = true AND
                ((EXTRACT(YEAR FROM b.effective_from) < ? OR (EXTRACT(YEAR FROM b.effective_from) = ? AND EXTRACT(MONTH FROM b.effective_from) <= m.month)) AND
                (EXTRACT(YEAR FROM b.effective_to) > ? OR (EXTRACT(YEAR FROM b.effective_to) = ? AND EXTRACT(MONTH FROM b.effective_to) >= m.month)))
            GROUP BY m.month
        ),
        expense_monthly AS (
            SELECT
                EXTRACT(MONTH FROM date)::int as month,
                COALESCE(SUM(amount), 0) as actual_expense
            FROM expenses
            WHERE user_id = ? AND EXTRACT(YEAR FROM date) = ?`

    if month != nil {
        query += ` AND EXTRACT(MONTH FROM date) = ?`
    }

    query += `
            GROUP BY EXTRACT(MONTH FROM date)
        )
        SELECT
            m.month,
            ?::int as year,
            COALESCE(b.budget_amount, 0) as budget_amount,
            COALESCE(e.actual_expense, 0) as actual_expense,
            COALESCE(b.budget_amount, 0) - COALESCE(e.actual_expense, 0) as difference
        FROM months m
        LEFT JOIN budget_monthly b ON m.month = b.month
        LEFT JOIN expense_monthly e ON m.month = e.month
        ORDER BY m.month`

    var err error
    if month != nil {
        err = r.db.Raw(query, userID, year, year, year, year, userID, year, *month, year).Scan(&results).Error
    } else {
        err = r.db.Raw(query, userID, year, year, year, year, userID, year, year).Scan(&results).Error
    }

    return results, err
}
