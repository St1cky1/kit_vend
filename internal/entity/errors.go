// Описание таблицы кодов ответов (Приложение 1)
package entity

const (
	// 0 - Успешно
	ResultCodeSuccess = 0

	// 1 - Неизвестная ошибка сервиса
	ResultCodeUnknownError = 1

	// 2 - Неверная структура JSON
	ResultCodeInvalidJSON = 2

	// 3 - Ошибка чтения данных запроса
	ResultCodeReadingError = 3

	// 4 - Неверный формат запроса. Объект Auth не может быть пустым
	ResultCodeEmptyAuth = 4

	// 5 - Неверный формат запроса. Значение CompanyId должно быть числом
	ResultCodeInvalidCompanyId = 5

	// 6 - Неверный формат запроса. Значение RequestId должно быть числом
	ResultCodeInvalidRequestId = 6

	// 7 - Неверный формат запроса. Значение UserLogin не может быть пустым
	ResultCodeEmptyUserLogin = 7

	// 8 - Неверный формат запроса. Значение Sign не может быть пустым
	ResultCodeEmptySign = 8

	// 9 - Компания с таким Id не найдена
	ResultCodeCompanyNotFound = 9

	// 10 - Компания заблокирована
	ResultCodeCompanyBlocked = 10

	// 11 - Пользователь с таким логином не найден
	ResultCodeUserNotFound = 11

	// 12 - Пользователь заблокирован
	ResultCodeUserBlocked = 12

	// 13 - У пользователя нет прав на данную операцию
	ResultCodeNoPermission = 13

	// 14 - Неверное значение RequestId
	ResultCodeInvalidRequestIDValue = 14

	// 15 - Неверная подпись
	ResultCodeInvalidSignature = 15

	// 16 - Неверный формат запроса. Объект Filter не может быть пустым
	ResultCodeEmptyFilter = 16

	// 17 - Неверный формат запроса. Значение UpDate не может быть пустым
	ResultCodeEmptyUpDate = 17

	// 18 - Неверный формат запроса. Значение ToDate не может быть пустым
	ResultCodeEmptyToDate = 18

	// 19 - Неверный формат запроса. Значение UpDate должно быть датой
	ResultCodeInvalidUpDate = 19

	// 20 - Неверный формат запроса. Значение ToDate должно быть датой
	ResultCodeInvalidToDate = 20

	// 21 - Неверный формат запроса. Значение ToDate должно быть больше UpDate
	ResultCodeToDateLessUpDate = 21

	// 22 - Неверный формат запроса. Значение CompanyId в фильтре должно быть числом
	ResultCodeInvalidFilterCompanyId = 22

	// 23 - Неверный формат запроса. Значение VendingMachineId в фильтре должно быть числом
	ResultCodeInvalidFilterVendingMachineId = 23
)
