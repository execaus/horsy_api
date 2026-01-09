package service

var (
	ErrAccountEmailExists       = NewServiceError("Аккаунт с таким email уже существует")
	ErrSignInInvalidCredentials = NewServiceError("Неверный email или пароль")
)
