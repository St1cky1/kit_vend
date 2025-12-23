# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [pb/v1/vending_machine.proto](#pb_v1_vending_machine-proto)
    - [Action](#api-v1-Action)
    - [Command](#api-v1-Command)
    - [Event](#api-v1-Event)
    - [GetActionsRequest](#api-v1-GetActionsRequest)
    - [GetActionsResponse](#api-v1-GetActionsResponse)
    - [GetEventsRequest](#api-v1-GetEventsRequest)
    - [GetEventsResponse](#api-v1-GetEventsResponse)
    - [GetSalesRequest](#api-v1-GetSalesRequest)
    - [GetSalesResponse](#api-v1-GetSalesResponse)
    - [GetVMStatesRequest](#api-v1-GetVMStatesRequest)
    - [GetVMStatesResponse](#api-v1-GetVMStatesResponse)
    - [GetVendingMachineByIDRequest](#api-v1-GetVendingMachineByIDRequest)
    - [GetVendingMachineByIDResponse](#api-v1-GetVendingMachineByIDResponse)
    - [GetVendingMachineRemainsRequest](#api-v1-GetVendingMachineRemainsRequest)
    - [GetVendingMachineRemainsResponse](#api-v1-GetVendingMachineRemainsResponse)
    - [Sale](#api-v1-Sale)
    - [SendCommandRequest](#api-v1-SendCommandRequest)
    - [SendCommandResponse](#api-v1-SendCommandResponse)
    - [VMState](#api-v1-VMState)
    - [VendingMachine](#api-v1-VendingMachine)
    - [VendingMachineRemains](#api-v1-VendingMachineRemains)
  
    - [VendingMachineService](#api-v1-VendingMachineService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="pb_v1_vending_machine-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pb/v1/vending_machine.proto



<a name="api-v1-Action"></a>

### Action
Action содержит информацию о действии (обслуживании, загрузке) с автоматом


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Уникальный идентификатор действия |
| vending_machine_id | [int32](#int32) |  | ID торгового автомата |
| action_type | [int32](#int32) |  | Тип действия (1=обслуживание, 2=загрузка товара, 3=инкассация и т.д.) |
| date_time | [string](#string) |  | Дата и время действия (ISO 8601) |






<a name="api-v1-Command"></a>

### Command
Command содержит команду для выполнения на торговом автомате


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vending_machine_id | [int32](#int32) |  | ID торгового автомата, на который отправляется команда |
| command_code | [int32](#int32) |  | Код команды (см. таблицу кодов команд: 1=обновление, 2=отчёт, 3=загрузка матрицы и т.д.) |






<a name="api-v1-Event"></a>

### Event
Event содержит информацию о событии (ошибке, предупреждении) в автомате


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Уникальный идентификатор события |
| vending_machine_id | [int32](#int32) |  | ID торгового автомата, где произошло событие |
| event_code | [int32](#int32) |  | Код события/ошибки (определяется спецификацией оборудования) |
| date_time | [string](#string) |  | Дата и время события (ISO 8601) |
| description | [string](#string) |  | Описание события на русском языке |






<a name="api-v1-GetActionsRequest"></a>

### GetActionsRequest
GetActionsRequest - запрос списка действий за период


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| from_date | [string](#string) |  | Начальная дата периода (ISO 8601) |
| to_date | [string](#string) |  | Конечная дата периода (ISO 8601) |
| vending_machine_id | [int32](#int32) |  | ID торгового автомата для фильтрации (опционально) |






<a name="api-v1-GetActionsResponse"></a>

### GetActionsResponse
GetActionsResponse - ответ со списком действий


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actions | [Action](#api-v1-Action) | repeated | Список всех действий за указанный период |






<a name="api-v1-GetEventsRequest"></a>

### GetEventsRequest
GetEventsRequest - запрос списка событий за период


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| from_date | [string](#string) |  | Начальная дата периода (ISO 8601) |
| to_date | [string](#string) |  | Конечная дата периода (ISO 8601) |
| vending_machine_id | [int32](#int32) |  | ID торгового автомата для фильтрации (опционально) |






<a name="api-v1-GetEventsResponse"></a>

### GetEventsResponse
GetEventsResponse - ответ со списком событий


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| events | [Event](#api-v1-Event) | repeated | Список всех событий за указанный период |






<a name="api-v1-GetSalesRequest"></a>

### GetSalesRequest
GetSalesRequest - запрос списка продаж за период


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| from_date | [string](#string) |  | Начальная дата периода (ISO 8601) |
| to_date | [string](#string) |  | Конечная дата периода (ISO 8601) |
| vending_machine_id | [int32](#int32) |  | ID торгового автомата для фильтрации (опционально) |






<a name="api-v1-GetSalesResponse"></a>

### GetSalesResponse
GetSalesResponse - ответ со списком продаж


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sales | [Sale](#api-v1-Sale) | repeated | Список всех продаж за указанный период |






<a name="api-v1-GetVMStatesRequest"></a>

### GetVMStatesRequest
GetVMStatesRequest - запрос состояния всех автоматов






<a name="api-v1-GetVMStatesResponse"></a>

### GetVMStatesResponse
GetVMStatesResponse - ответ с состоянием всех автоматов


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vm_states | [VMState](#api-v1-VMState) | repeated | Список состояний всех торговых автоматов |






<a name="api-v1-GetVendingMachineByIDRequest"></a>

### GetVendingMachineByIDRequest
GetVendingMachineByIDRequest - запрос информации об автомате по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | ID торгового автомата |






<a name="api-v1-GetVendingMachineByIDResponse"></a>

### GetVendingMachineByIDResponse
GetVendingMachineByIDResponse - ответ с информацией об автомате


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vending_machine | [VendingMachine](#api-v1-VendingMachine) |  | Информация о торговом автомате |






<a name="api-v1-GetVendingMachineRemainsRequest"></a>

### GetVendingMachineRemainsRequest
GetVendingMachineRemainsRequest - запрос остатков товара в автомате


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | ID торгового автомата |






<a name="api-v1-GetVendingMachineRemainsResponse"></a>

### GetVendingMachineRemainsResponse
GetVendingMachineRemainsResponse - ответ с остатками товара


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| remains | [VendingMachineRemains](#api-v1-VendingMachineRemains) | repeated | Список остатков всех товаров в автомате |






<a name="api-v1-Sale"></a>

### Sale
Sale содержит информацию об одной продаже товара из автомата


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Уникальный идентификатор продажи |
| vending_machine_id | [int32](#int32) |  | ID торгового автомата, из которого была сделана продажа |
| goods_id | [int32](#int32) |  | ID товара |
| goods_name | [string](#string) |  | Название товара |
| count | [int32](#int32) |  | Количество единиц товара |
| sum | [double](#double) |  | Сумма продажи |
| date_time | [string](#string) |  | Дата и время продажи (ISO 8601) |
| payment_method | [int32](#int32) |  | Метод оплаты (1=наличные, 2=карта, 3=мобильный платёж и т.д.) |






<a name="api-v1-SendCommandRequest"></a>

### SendCommandRequest
SendCommandRequest - запрос на отправку команды автомату


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| command | [Command](#api-v1-Command) |  | Команда для отправки |






<a name="api-v1-SendCommandResponse"></a>

### SendCommandResponse
SendCommandResponse - ответ с результатом отправки команды


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result_code | [int32](#int32) |  | Код результата (0=успех, &gt;0=ошибка) |
| command_id | [int32](#int32) |  | ID команды для отслеживания её выполнения |






<a name="api-v1-VMState"></a>

### VMState
VMState содержит текущее состояние торгового автомата и его компонентов


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vending_machine_id | [int32](#int32) |  | ID торгового автомата |
| machine_name | [string](#string) |  | Название автомата |
| is_online | [bool](#bool) |  | Статус подключения (true=онлайн, false=оффлайн) |
| last_activity_time | [string](#string) |  | Время последней активности (ISO 8601) |
| power_supply | [int32](#int32) |  | Статус источника питания (0=нормально, 1=низкое напряжение, 2=критическое) |
| bill_acceptor | [int32](#int32) |  | Статус приёмника банкнот (0=рабочий, 1=ошибка, 2=отсутствует) |
| coin_acceptor | [int32](#int32) |  | Статус приёмника монет (0=рабочий, 1=ошибка, 2=отсутствует) |
| non_cash_payment | [int32](#int32) |  | Статус системы бесконтактной оплаты (0=рабочая, 1=ошибка, 2=отсутствует) |
| cash_register | [int32](#int32) |  | Статус кассы (0=рабочая, 1=ошибка, 2=полная) |
| qr_display | [int32](#int32) |  | Статус QR дисплея (0=рабочий, 1=ошибка, 2=отсутствует) |






<a name="api-v1-VendingMachine"></a>

### VendingMachine
VendingMachine содержит базовую информацию о торговом автомате


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Уникальный идентификатор автомата |
| name | [string](#string) |  | Название/описание автомата |
| company_id | [int32](#int32) |  | ID компании, которой принадлежит автомат |






<a name="api-v1-VendingMachineRemains"></a>

### VendingMachineRemains
VendingMachineRemains содержит информацию об остатках товара в автомате


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vending_machine_id | [int32](#int32) |  | ID торгового автомата |
| goods_id | [int32](#int32) |  | ID товара |
| goods_name | [string](#string) |  | Название товара |
| count | [int32](#int32) |  | Количество единиц товара в автомате |





 

 

 


<a name="api-v1-VendingMachineService"></a>

### VendingMachineService
VendingMachineService предоставляет методы для работы с торговыми автоматами
Включает получение информации о состоянии, продажах, действиях и событиях автоматов

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetVendingMachineByID | [GetVendingMachineByIDRequest](#api-v1-GetVendingMachineByIDRequest) | [GetVendingMachineByIDResponse](#api-v1-GetVendingMachineByIDResponse) | GetVendingMachineByID получает информацию о конкретном торговом автомате по его идентификатору Возвращает базовую информацию об автомате: ID, название и ID компании |
| GetSales | [GetSalesRequest](#api-v1-GetSalesRequest) | [GetSalesResponse](#api-v1-GetSalesResponse) | GetSales получает список всех продаж для указанного периода Может быть отфильтрована по ID автомата Возвращает детальную информацию о каждой продаже: товар, количество, сумма, метод оплаты |
| GetActions | [GetActionsRequest](#api-v1-GetActionsRequest) | [GetActionsResponse](#api-v1-GetActionsResponse) | GetActions получает список действий (обслуживание, загрузка товара) для указанного периода Может быть отфильтрована по ID автомата |
| GetVMStates | [GetVMStatesRequest](#api-v1-GetVMStatesRequest) | [GetVMStatesResponse](#api-v1-GetVMStatesResponse) | GetVMStates получает текущее состояние всех торговых автоматов Включает информацию о подключении, состоянии компонентов и времени последней активности |
| GetEvents | [GetEventsRequest](#api-v1-GetEventsRequest) | [GetEventsResponse](#api-v1-GetEventsResponse) | GetEvents получает список событий (ошибки, предупреждения) для указанного периода Может быть отфильтрована по ID автомата События включают код ошибки и описание |
| SendCommand | [SendCommandRequest](#api-v1-SendCommandRequest) | [SendCommandResponse](#api-v1-SendCommandResponse) | SendCommand отправляет команду на выполнение конкретному торговому автомату Команда включает тип команды (перезагрузка, обновление, синхронизация и т.д.) Возвращает ID команды для отслеживания её выполнения |
| GetVendingMachineRemains | [GetVendingMachineRemainsRequest](#api-v1-GetVendingMachineRemainsRequest) | [GetVendingMachineRemainsResponse](#api-v1-GetVendingMachineRemainsResponse) | GetVendingMachineRemains получает информацию об остатках товара в автомате Включает ID товара, название и количество единиц |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

