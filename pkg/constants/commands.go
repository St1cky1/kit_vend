// коды команд
package constants

type CommandCode uint

const (
	CommandSoftwareUpdate CommandCode = 1

	CommandGetReport CommandCode = 2

	CommandLoadMatrix CommandCode = 3

	CommandApplyMatrix CommandCode = 4

	CommandEnableFiscalization CommandCode = 5

	CommandDisableFiscalization CommandCode = 6

	CommandApplyDefaultVATRate CommandCode = 7

	CommandDisableServiceMode CommandCode = 8

	CommandCompleteUpdate CommandCode = 9

	CommandRestartKitBox CommandCode = 10

	CommandStopTerminalWork CommandCode = 11

	CommandResumeTerminalWork CommandCode = 12

	CommandSyncTimeWithServer CommandCode = 14

	CommandSendLogToServer CommandCode = 15

	CommandAllowFreeOfChargeSales CommandCode = 16

	CommandRestartTerminal CommandCode = 18

	CommandRestartPaymentSystem CommandCode = 19

	CommandEnableDelayedFiscalization CommandCode = 21

	CommandDisableDelayedFiscalization CommandCode = 22

	CommandCollect CommandCode = 24

	CommandMaintenance CommandCode = 25

	CommandLoad CommandCode = 26

	CommandEnableMultisale CommandCode = 27

	CommandDisableMultisale CommandCode = 28

	CommandUpdateReaderSoftware CommandCode = 30

	CommandCompleteReaderUpdate CommandCode = 31

	CommandEnableReaderOfflineMode CommandCode = 32

	CommandDisableReaderOfflineMode CommandCode = 33

	CommandLoadAndApplyMatrix CommandCode = 34
)

var CommandNames = map[CommandCode]string{
	CommandSoftwareUpdate:              "Обновление ПО",
	CommandGetReport:                   "Снятие отчета",
	CommandLoadMatrix:                  "Загрузка матрицы",
	CommandApplyMatrix:                 "Применение матрицы",
	CommandEnableFiscalization:         "Включение обл. фискализации",
	CommandDisableFiscalization:        "Выключение обл. фискализации",
	CommandApplyDefaultVATRate:         "Применить деф. ставку НДС",
	CommandDisableServiceMode:          "Выкл. сервисный режим",
	CommandCompleteUpdate:              "Завершить обновление",
	CommandRestartKitBox:               "Перезагрузить Kit Box",
	CommandStopTerminalWork:            "Остановить работу ТА",
	CommandResumeTerminalWork:          "Возобновить работу ТА",
	CommandSyncTimeWithServer:          "Синхронизировать время с сервером",
	CommandSendLogToServer:             "Выслать лог на сервер",
	CommandAllowFreeOfChargeSales:      "Разрешить бесплатную продажу",
	CommandRestartTerminal:             "Перезагрузить ТА",
	CommandRestartPaymentSystem:        "Перезагрузить ПС",
	CommandEnableDelayedFiscalization:  "Вкл. отложенную фискализацию",
	CommandDisableDelayedFiscalization: "Выкл. отложенную фискализацию",
	CommandCollect:                     "Сделать инкассацию",
	CommandMaintenance:                 "Сделать обслуживание",
	CommandLoad:                        "Сделать загрузку",
	CommandEnableMultisale:             "Включить мультипродажу",
	CommandDisableMultisale:            "Выключить мультипродажу",
	CommandUpdateReaderSoftware:        "Обновление ПО ридера",
	CommandCompleteReaderUpdate:        "Завершить обновление ПО ридера",
	CommandEnableReaderOfflineMode:     "Включить оффлайн режим ридера",
	CommandDisableReaderOfflineMode:    "Выключить оффлайн режим ридера",
	CommandLoadAndApplyMatrix:          "Загрузить и применить матрицу",
}

func (cc CommandCode) String() string {
	if name, ok := CommandNames[cc]; ok {
		return name
	}
	return "Неизвестная команда"
}
