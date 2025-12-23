// коды ошибок
package constants

type ErrorCode uint

const (
	ResultCodeSuccess ErrorCode = 0

	ResultCodeUnknownError ErrorCode = 1

	ResultCodeInvalidJSON ErrorCode = 2

	ResultCodeReadingError ErrorCode = 3

	ResultCodeEmptyAuth ErrorCode = 4

	ResultCodeInvalidCompanyId ErrorCode = 5

	ResultCodeInvalidRequestId ErrorCode = 6

	ResultCodeEmptyUserLogin ErrorCode = 7

	ResultCodeEmptySign ErrorCode = 8

	ResultCodeCompanyNotFound ErrorCode = 9

	ResultCodeCompanyBlocked ErrorCode = 10

	ResultCodeUserNotFound ErrorCode = 11

	ResultCodeUserBlocked ErrorCode = 12

	ResultCodeNoPermission ErrorCode = 13

	ResultCodeInvalidRequestIDValue ErrorCode = 14

	ResultCodeInvalidSignature ErrorCode = 15

	ResultCodeEmptyFilter ErrorCode = 16

	ResultCodeEmptyUpDate ErrorCode = 17

	ResultCodeEmptyToDate ErrorCode = 18

	ResultCodeInvalidUpDate ErrorCode = 19

	ResultCodeInvalidToDate ErrorCode = 20

	ResultCodeToDateLessUpDate ErrorCode = 21

	ResultCodeInvalidFilterCompanyId ErrorCode = 22

	ResultCodeInvalidFilterVendingMachineId ErrorCode = 23
)

var ErrorMessages = map[ErrorCode]string{
	ResultCodeSuccess:                       "Успешно",
	ResultCodeUnknownError:                  "Неизвестная ошибка сервиса",
	ResultCodeInvalidJSON:                   "Неверная структура JSON",
	ResultCodeReadingError:                  "Ошибка чтения данных запроса",
	ResultCodeEmptyAuth:                     "Неверный формат запроса. Объект Auth не может быть пустым",
	ResultCodeInvalidCompanyId:              "Неверный формат запроса. Значение CompanyId должно быть числом",
	ResultCodeInvalidRequestId:              "Неверный формат запроса. Значение RequestId должно быть числом",
	ResultCodeEmptyUserLogin:                "Неверный формат запроса. Значение UserLogin не может быть пустым",
	ResultCodeEmptySign:                     "Неверный формат запроса. Значение Sign не может быть пустым",
	ResultCodeCompanyNotFound:               "Компания с таким Id не найдена",
	ResultCodeCompanyBlocked:                "Компания заблокирована",
	ResultCodeUserNotFound:                  "Пользователь с таким логином не найден",
	ResultCodeUserBlocked:                   "Пользователь заблокирован",
	ResultCodeNoPermission:                  "У пользователя нет прав на данную операцию",
	ResultCodeInvalidRequestIDValue:         "Неверное значение RequestId",
	ResultCodeInvalidSignature:              "Неверная подпись",
	ResultCodeEmptyFilter:                   "Неверный формат запроса. Объект Filter не может быть пустым",
	ResultCodeEmptyUpDate:                   "Неверный формат запроса. Значение UpDate не может быть пустым",
	ResultCodeEmptyToDate:                   "Неверный формат запроса. Значение ToDate не может быть пустым",
	ResultCodeInvalidUpDate:                 "Неверный формат запроса. Значение UpDate должно быть датой",
	ResultCodeInvalidToDate:                 "Неверный формат запроса. Значение ToDate должно быть датой",
	ResultCodeToDateLessUpDate:              "Неверный формат запроса. Значение ToDate должно быть больше UpDate",
	ResultCodeInvalidFilterCompanyId:        "Неверный формат запроса. Значение CompanyId в фильтре должно быть числом",
	ResultCodeInvalidFilterVendingMachineId: "Неверный формат запроса. Значение VendingMachineId в фильтре должно быть числом",
}

func (ec ErrorCode) String() string {
	if msg, ok := ErrorMessages[ec]; ok {
		return msg
	}
	return "Неизвестная ошибка"
}
