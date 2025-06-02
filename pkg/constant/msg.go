package constant

const (
    ErrCodeUnauthorized       = "Tidak Terotorisasi"
    ErrCodeForbidden         = "Akses Dilarang"
    ErrCodeBadRequest        = "Permintaan Tidak Valid"
    ErrCodeNotFound          = "Tidak Ditemukan"
    ErrCodeInternalError     = "Kesalahan Internal"
    ErrCodeValidation        = "Kesalahan Validasi"
    ErrCodeDatabaseError     = "Database Error"
    ErrCodeDuplicateEntry    = "Duaplikat Entri"
    ErrCodeInvalidCredential = "Kredensial Tidak Valid"
    ErrCodeExpiredToken      = "Token Kadaluarsa"
    ErrCodeInvalidToken      = "Token Tidak Valid"
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
	Msg
)

const (
	MsgRegisterSuccess = "Registrasi berhasil"
	MsgLoginSuccess    = "Login berhasil"
	MsgRetrivedUserSuccess = "Berhasil mendapatkan profile"
	MsgLogoutSuccess = "Logout berhasil"
	MsgCreateCategorySuccess = "Kategori berhasil dibuat"
	MsgRetrievedCategoriesSuccess = "Berhasil mendapatkan daftar kategori"
	MsgUpdateCategorySuccess = "Category updated successfully"
    MsgCategoryNotFound     = "Category not found"
	MsgCreateIncomeSuccess = "Pendapatan berhasil dibuat"
	MsgGetIncomesSuccess = "Berhasil mendapatkan daftar pendapatan"
	MsgIncomeNotFound     = "Pemasukan not found"
	MsgUpdateIncomeSuccess = "Pendapatan berhasil diperbarui"
	MsgDeleteIncomeSuccess = "Pendapatan berhasil dihapus"
	MsgExpenseNotFound     = "Pengeluran not found"
	MsgCreateExpenseSuccess = "Pengeluaran berhasil dibuat"
	MsgGetExpensesSuccess = "Berhasil mendapatkan daftar pengeluaran"
	MsgUpdateExpenseSuccess = "Pengeluaran berhasil diperbarui"
	MsgDeleteExpenseSuccess = "Pengeluaran berhasil dihapus"
	MsgBudgetNotFound = "Anggaran tidak ditemukan"
	MsgCreateBudgetSuccess = "Anggaran berhasil dibuat"
	MsgGetBudgetsSuccess = "Berhasil mendapatkan daftar anggaran"
	MsgUpdateBudgetSuccess = "Anggaran berhasil diperbarui"
	MsgDeleteBudgetSuccess = "Anggaran berhasil dihapus"
)
