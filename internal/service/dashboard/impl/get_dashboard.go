package impl

import (
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/dashboard/dto"
    "pencatatan_keuangan/internal/service/dashboard/mapper"
    "pencatatan_keuangan/pkg/constant"
    "pencatatan_keuangan/pkg/utils"
    "time"

    "github.com/google/uuid"
)

func GetDashboard(
    repo repository.DashboardRepository,
    dashboardMapper *mapper.DashboardMapper,
    userID uuid.UUID,
    filter dto.FilterRequest,
) (*dto.DashboardResponse, error) {
    if filter.Year == 0 {
        filter.Year = time.Now().Year()
    }

    var month *int
    if filter.Month > 0 {
        month = &filter.Month
    }

    totalIncome, err := repo.GetTotalIncome(userID, filter.Year, month)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    totalExpense, err := repo.GetTotalExpense(userID, filter.Year, month)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    balance := totalIncome - totalExpense

    totalBudget, err := repo.GetTotalBudget(userID, filter.Year, month)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    expenseDistribution, err := repo.GetExpenseDistribution(userID, filter.Year, month)
    if err != nil {
        return nil, utils.NewSystemError(constant.MsgInternalError, err)
    }

    // monthlyIncomes, err := repo.GetMonthlyIncomes(userID, filter.Year)
    // if err != nil {
    //     return nil, utils.NewSystemError(constant.MsgInternalError, err)
    // }

    // monthlyExpenses, err := repo.GetMonthlyExpenses(userID, filter.Year)
    // if err != nil {
    //     return nil, utils.NewSystemError(constant.MsgInternalError, err)
    // }

    budgetComparison, err := repo.GetBudgetComparison(userID, filter.Year, month)
	if err != nil {
		return nil, utils.NewSystemError(constant.MsgInternalError, err)
	}

    for i := range budgetComparison {
        budgetComparison[i].IsOverBudget = budgetComparison[i].ActualExpense > budgetComparison[i].BudgetAmount
    }

    summary := dto.Summary{
        TotalIncome:  totalIncome,
        TotalExpense: totalExpense,
        TotalBudget:  totalBudget,
        Balance:      balance,
    }

    dashboardResponse := &dto.DashboardResponse{
        Summary:             summary,
        // MonthlyIncomes:      monthlyIncomes,
        // MonthlyExpenses:     monthlyExpenses,
        ExpenseDistribution: expenseDistribution,
        BudgetComparison:    budgetComparison,
    }

    return dashboardResponse, nil
}
