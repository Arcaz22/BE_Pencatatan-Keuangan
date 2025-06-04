package dto

type FilterRequest struct {
    Year  int `form:"year" json:"year"`
    Month int `form:"month" json:"month,omitempty"`
}

// Summary struktur untuk total
type Summary struct {
    TotalIncome  float64 `json:"total_income"`
    TotalExpense float64 `json:"total_expense"`
    TotalBudget  float64 `json:"total_budget"`
    Balance      float64 `json:"balance"`
}

type MonthlyTotal struct {
    Month    int     `json:"month"`
    Year     int     `json:"year"`
    Total    float64 `json:"total"`
    Currency string  `json:"currency"`
}

type CategoryDistribution struct {
    CategoryID   string  `json:"category_id"`
    CategoryName string  `json:"category_name"`
    Amount       float64 `json:"amount"`
    Percentage   float64 `json:"percentage"`
}

type BudgetComparison struct {
    Month         int     `json:"month"`
    Year          int     `json:"year"`
    BudgetAmount  float64 `json:"budget_amount"`
    ActualExpense float64 `json:"actual_expense"`
    Difference    float64 `json:"difference"`
    IsOverBudget  bool    `json:"is_over_budget"`
}

type DashboardResponse struct {
    Summary             Summary                `json:"summary"`
    // MonthlyIncomes      []MonthlyTotal         `json:"monthly_incomes"`
    // MonthlyExpenses     []MonthlyTotal         `json:"monthly_expenses"`
    ExpenseDistribution []CategoryDistribution `json:"expense_distribution"`
    BudgetComparison    []BudgetComparison     `json:"budget_comparison"`
}
