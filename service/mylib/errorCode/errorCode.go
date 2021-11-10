package errorCode

import "server/service/mylib/errorhandler"

const (
	// 一般常用訊息
	Success                    = iota
	Error                      = 1
	ParamsError                = 2
	DecodeJsonError            = 3
	EncodeJsonError            = 4
	FormatError                = 5
	HashError                  = 6
	QrError                    = 7
	DataValidError             = 8
	TimeOutError               = 9
	EncryptError               = 10
	NoRouteError               = 11
	AccountAndPwdError         = 12
	IllegalError               = 13
	RequestCreateError         = 14
	RequestSendError           = 15
	ResponseDecodeError        = 16
	FileOpenError              = 17
	FileRemoveError            = 18
	DirectoryMakeError         = 19
	DirectoryRemoveError       = 20
	UploadFileFormatError      = 21
	UploadFileSizeError        = 22
	PermissionDenyError        = 23
	RetryError                 = 24
	AccountFormatError         = 25
	PasswordFormatError        = 26
	EmailFormatError           = 27
	CellPhoneFormatError       = 28
	TimeFormatError            = 29
	TimeSettingError           = 30
	MaxLimitError              = 31
	MinLimitError              = 32
	Relogin                    = 33
	AlreadyLoginFromOtherPlace = 34
	TimeCantBeEarlierThanNow   = 35
	TimeCantBeLaterThanNow     = 36
	TimeRangeIsTooLog          = 37

	// DB錯誤
	DBConnectError      = 1000
	DBInsertError       = 1001
	DBUpdateError       = 1002
	DBSelectError       = 1003
	DBDeleteError       = 1004
	DBNoData            = 1005
	DBDuplicateError    = 1006
	DBDataNoChangeError = 1007

	DBAccountDuplicateError = 1500

	// Redis錯誤
	RedisConnectError = 2000
	RedisSetError     = 2001
	RedisDelError     = 2002
	RedisGetError     = 2003

	// Websocket錯誤
	WebsocketDecodeError                   = 3000
	WebsocketActionError                   = 3001
	WebsocketTypeError                     = 3002
	WebsocketRespError                     = 3003
	WebsocketReceiveError                  = 3004
	WebsocketNotFoundClientError           = 3005
	WebsocketChannelCloseError             = 3006
	WebsocketFinishCallButClientCloseError = 3007

	// RabbitMq錯誤
	RabbitMqCreateConnError       = 4000
	RabbitMqCreateChannelError    = 4001
	RabbitMqCreateQueueError      = 4002
	RabbitMqNotYetCreateConnError = 4003
	RabbitMqActionNotDefineError  = 4004

	// ocpp message 錯誤
	OCPPJSonDecodeError         = 5000
	OCPPJWSError                = 5001
	OCPPWebSocketHeaderNotFound = 5002
	OCPPAccountAndPasswordErr   = 5003
	OCPPVersionNotSupport       = 5004
	OCPPValidError              = 5005
	OCPPMessageIdNotFound       = 5006

	// ocpp message的訊息(reason code)
	CSNotAccepted        = 6000 //BootNotification of Charging Station has not (yet) been accepted by CSMS. RequestStartTransaction, RequestStopTransaction
	DuplicateProfile     = 6001 //A charging profile with same stackLevel - chargingProfilePurpose combination already exists on the Charging Station and has an overlapping validity period. SetChargingProfile
	DuplicateRequestId   = 6002 //A requestId is provided, that has already been used for this type of request. UpdateFirmware, PublishFirmware and requests for reports.
	FixedCable           = 6003 //The connector has its own fixed cable that cannot be unlocked. UnlockConnector
	FwUpdateInProgress   = 6004 //Operation is not possible, because a firmware update is in progress. Reset
	InternalError        = 6005 //Operation cannot be completed due to an internal error. (generic)
	InvalidCertificate   = 6006 //Provided certificate is invalid. CertificateSigned, InstallCertificate
	InvalidCSR           = 6007 //Provided CSR is invalid SignCertificate
	InvalidIdToken       = 6008 //Provided idToken is not valid. RequestStartTransaction
	InvalidProfile       = 6009 //Provided chargingProfile contains invalid elements. SetChargingProfile, RequestStartTransaction
	InvalidSchedule      = 6010 //Provided chargingSchedule contains invalid elements. SetChargingProfile, RequestStartTransaction
	InvalidStackLevel    = 6011 //Provided value for stackLevel is invalid. SetChargingProfile
	InvalidURL           = 6012 //Provided URL is invalid. UpdateFirmware,PublishFirmware
	InvalidValue         = 6013 //An invalid value has been provided. (generic)
	MissingParam         = 6014 //A parameter that is required for the request is missing. (generic)
	NoCable              = 6015 //No cable is connected at this time. UnlockConnector
	NoError              = 6016 //No error has occurred, but some extra information is in additionalInfo .(generic)
	NotEnabled           = 6017 //Feature is not enabled. ClearCache
	NotFound             = 6018 //No object(s) found that match a provided ID or criteria. ClearVariableMonitoring, CustomerInformation,GetChargingProfiles,GetDisplayMessages,GetInstalledCertificateIds,GetReport
	OutOfMemory          = 6019 //Operation not possible, because system does not have enough memory. (generic)
	OutOfStorage         = 6020 //Operation not possible, because system does not have enough storage. (generic)
	ReadOnly             = 6021 //Targeted variable is read-only and cannot be set. SetVariables
	TooLargeElement      = 6022 //Provided element is too large to handle. CertificateSigned, InstallCertificate
	TooManyElements      = 6023 //Too many elements have been provided. SetChargingProfile, SetVariables, SendLocalList
	TxInProgress         = 6024 //A transaction is in progress. ChangeAvailability, Reset, RequestStartTransaction
	TxNotFound           = 6025 //There is no such transaction. RequestStopTransaction, SetChargingProfile
	TxStarted            = 6026 //A transaction had already started (e.g. due to cable being plugged in). RequestStartTransaction
	UnknownConnectorId   = 6027 //Connector Id is not known on EVSE ChangeAvailability, UnlockConnector
	UnknownConnectorType = 6028 //Connector type is not known on EVSE ReserveNow
	UnknownEvse          = 6029 //EVSE is not known on Charging Stations ChangeAvailability, ReserveNow, RequestStartTransaction
	UnknownTxId          = 6030 //Provided transactionId is not known. RequestStopTransaction
	Unspecified          = 6031 //No reason is specified, but some extra information is in additionalInfo (generic)
	UnsupportedParam     = 6032 //A parameter was provided that is not supported. (generic)
	UnsupportedRateUnit  = 6033 //A chargingRateUnit is provided that is not supported. SetChargingProfile
	UnsupportedRequest   = 6034 //This request is not supported. (generic)
	ValueOutOfRange      = 6035 //Provided value is out of range. SetVariables, SetVariableMonitoring
	ValuePositiveOnly    = 6036 //Provided value is not greater than zero. (generic)
	ValueTooHigh         = 6037 //Provided value is too high. (generic)
	ValueTooLow          = 6038 //Provided value is too low. (generic)
	ValueZeroNotAllowed  = 6039 //Provided value cannot be zero. (generic)
	WriteOnly            = 6040 //Targeted variable is write-only and cannot be read. GetVariables

	// ocpp message的訊息(error code)
	FormatViolation = 6500
	GenericError    = 6501
	// InternalError                 = 6502	//(跟reason code有重複)
	MessageTypeNotSupported       = 6503
	NotImplemented                = 6504
	NotSupported                  = 6505
	OccurrenceConstraintViolation = 6506
	PropertyConstraintViolation   = 6507
	ProtocolError                 = 6508
	RpcFrameworkError             = 6509
	SecurityError                 = 6510
	TypeConstraintViolation       = 6511

	// GRPC Error
	GrpcConnectError     = 7000
	GrpcNotFoundCMSError = 7001
	GrpcAlreadyExist     = 7002
	GrpcNotExist         = 7003

	// TBox Error
	TBoxParamError         = 10000 //Invalid parameter
	TBoxNotFound           = 10001 //T-Box not online
	TBoxFwUpdateInProgress = 10002 //During firmware update
	TBoxAddNFCInProgress   = 10003 //During NFC Card Modification locally
	TBoxNFCAboveMax        = 10004 //Above maximum NFC Card IDs
	TBoxDuplicateNFC       = 10005 //Duplicate NFC Card ID
	TBoxNFCNotFound        = 10006 //NFC Card ID not exist

	// Image hosting
	ImageHostingConnectError = 20000
	PutObjectError           = 20001
	GenerateUrlError         = 20002
	ImageHostingObjectError  = 20003
	Base64DecoderError       = 20004
	PhotoEncodeError         = 20005

	// telegram錯誤
	TgNotFoundUser    = 40001
	TgSetCommandError = 40002

	// 尚未定義到的錯誤訊息
	Undefined = -999
)

func init() {
	Errorcodes := map[int]string{
		// 尚未定義到的錯誤訊息
		Undefined: "InternalError", // 發生錯誤

		// 一般常用訊息
		Success:                    "",                           // 成功
		Error:                      "InternalError",              // 發生錯誤
		ParamsError:                "ParamsError",                // 參數錯誤
		DecodeJsonError:            "InternalError",              // 解析JSON資料錯誤
		EncodeJsonError:            "InternalError",              // 加密JSON資料錯誤
		FormatError:                "FormatError",                // 格式錯誤
		HashError:                  "InternalError",              // 系統內部錯誤
		QrError:                    "InternalError",              // 系統內部錯誤
		DataValidError:             "DataInvalidError",           // 資料驗證錯誤
		TimeOutError:               "TimeOutError",               // 已超時錯誤
		EncryptError:               "InternalError",              // 加密錯誤
		NoRouteError:               "UndefinedRouteError",        // 尚未定義的Router
		AccountAndPwdError:         "AccountOrPasswordError",     // 帳號或密碼錯誤
		IllegalError:               "IllegalError",               // 非法操作
		RequestCreateError:         "InternalError",              // Request建置錯誤
		RequestSendError:           "InternalError",              // Request發送錯誤
		ResponseDecodeError:        "InternalError",              // Request解析錯誤
		FileOpenError:              "FileOpenError",              // 檔案開啟錯誤
		FileRemoveError:            "FileRemoveError",            // 檔案刪除錯誤
		DirectoryMakeError:         "InternalError",              // 目錄建立錯誤
		DirectoryRemoveError:       "InternalError",              // 目錄刪除錯誤
		UploadFileFormatError:      "UploadFileFormatError",      // 上傳檔案格式錯誤
		UploadFileSizeError:        "UploadFileSizeError",        // 上傳檔案大小錯誤
		PermissionDenyError:        "PermissionDenyError",        // 權限不足
		RetryError:                 "RetryError",                 // 請稍後再試
		AccountFormatError:         "AccountFormatError",         // 帳號格式錯誤
		PasswordFormatError:        "PasswordFormatError",        // 密碼格式錯誤
		EmailFormatError:           "EmailFormatError",           // mail格式錯誤
		CellPhoneFormatError:       "CellPhoneFormatError",       // 電話格式錯誤
		TimeFormatError:            "TimeFormatError",            // 時間格式錯誤
		TimeSettingError:           "TimeSettingError",           // 時間設置有誤
		MaxLimitError:              "MaxLimitError",              // 已達上限值
		MinLimitError:              "MinLimitError",              // 已達下限值
		Relogin:                    "Relogin",                    // 請重新登入
		AlreadyLoginFromOtherPlace: "AlreadyLoginFromOtherPlace", // 已從其他地方登入
		TimeCantBeEarlierThanNow:   "TimeCantBeEarlyThanNow",     // 時間不得早於現在
		TimeCantBeLaterThanNow:     "TimeCantBeLaterThanNow",     // 時間不得晚於現在
		TimeRangeIsTooLog:          "TimeRangeIsTooLong",         // 時間區間設置過長

		// DB錯誤
		DBConnectError:      "InternalError",  // DB連線建置錯誤
		DBInsertError:       "InternalError",  // DB新增資料錯誤
		DBUpdateError:       "InternalError",  // DB更新資料錯誤
		DBSelectError:       "InternalError",  // DB查詢資料錯誤
		DBDeleteError:       "InternalError",  // DB刪除資料錯誤
		DBNoData:            "NoData",         // 查無資料
		DBDuplicateError:    "DuplicatedData", // 資料重複
		DBDataNoChangeError: "DataNoChange",   // 資料未異動

		DBAccountDuplicateError: "DuplicatedAccount", // 帳號重複

		// Redis錯誤
		RedisConnectError: "InternalError", // Redis連線建置錯誤
		RedisSetError:     "InternalError", // Redis Set資料錯誤
		RedisDelError:     "InternalError", // Redis Del資料錯誤
		RedisGetError:     "InternalError", // Redis Get資料錯誤

		// Websocket錯誤
		WebsocketDecodeError:                   "InternalError", // 解析錯誤
		WebsocketActionError:                   "InternalError", // 尚未定義該動作
		WebsocketTypeError:                     "InternalError", // 尚未定義type
		WebsocketRespError:                     "InternalError", // 尚未定義該回覆動作
		WebsocketReceiveError:                  "InternalError", // 收到client端回傳錯誤
		WebsocketNotFoundClientError:           "InternalError", // 查無此client
		WebsocketChannelCloseError:             "InternalError", // channel已經關閉
		WebsocketFinishCallButClientCloseError: "InternalError", // 已完成Call的動作 但client已經關閉連線

		// RabbitMq錯誤
		RabbitMqCreateConnError:       "InternalError", // Rabbitmq 建立連線失敗
		RabbitMqCreateChannelError:    "InternalError", // Rabbitmq 建立channel失敗
		RabbitMqCreateQueueError:      "InternalError", // Rabbitmq 建立queue失敗
		RabbitMqNotYetCreateConnError: "InternalError", // Rabbitmq 尚未建立連線
		RabbitMqActionNotDefineError:  "InternalError", // Rabbitmq 尚未定義該action

		// ocpp message 錯誤(自定義)
		OCPPJSonDecodeError:         "InternalError", // ocpp json decode error
		OCPPJWSError:                "InternalError", // jws encode error
		OCPPWebSocketHeaderNotFound: "InternalError", // ocpp header not found
		OCPPAccountAndPasswordErr:   "InternalError", // account and password is err
		OCPPVersionNotSupport:       "InternalError", // ocpp version is not support
		OCPPValidError:              "InternalError", // ocpp token is not valid
		OCPPMessageIdNotFound:       "InternalError", // message id can't not found

		// ocpp message的訊息(reason code)
		CSNotAccepted:        "InternalError", // CSNotAccepted
		DuplicateProfile:     "InternalError", // DuplicateProfile
		DuplicateRequestId:   "InternalError", // DuplicateRequestId
		FixedCable:           "InternalError", // FixedCable
		FwUpdateInProgress:   "InternalError", // FwUpdateInProgress
		InternalError:        "InternalError", // InternalError
		InvalidCertificate:   "InternalError", // InvalidCertificate
		InvalidCSR:           "InternalError", // InvalidCSR
		InvalidIdToken:       "InternalError", // InvalidIdToken
		InvalidProfile:       "InternalError", // InvalidProfile
		InvalidSchedule:      "InternalError", // InvalidSchedule
		InvalidStackLevel:    "InternalError", // InvalidStackLevel
		InvalidURL:           "InternalError", // InvalidURL
		InvalidValue:         "InternalError", // InvalidValue
		MissingParam:         "InternalError", // MissingParam
		NoCable:              "InternalError", // NoCable
		NoError:              "InternalError", // NoError
		NotEnabled:           "InternalError", // NotEnabled
		NotFound:             "InternalError", // NotFound
		OutOfMemory:          "InternalError", // OutOfMemory
		OutOfStorage:         "InternalError", // OutOfStorage
		ReadOnly:             "InternalError", // ReadOnly
		TooLargeElement:      "InternalError", // TooLargeElement
		TooManyElements:      "InternalError", // TooManyElements
		TxInProgress:         "InternalError", // TxInProgress
		TxNotFound:           "InternalError", // TxNotFound
		TxStarted:            "InternalError", // TxStarted
		UnknownConnectorId:   "InternalError", // UnknownConnectorId
		UnknownConnectorType: "InternalError", // UnknownConnectorType
		UnknownEvse:          "InternalError", // UnknownEvse
		UnknownTxId:          "InternalError", // UnknownTxId
		Unspecified:          "InternalError", // Unspecified
		UnsupportedParam:     "InternalError", // UnsupportedParam
		UnsupportedRateUnit:  "InternalError", // UnsupportedRateUnit
		UnsupportedRequest:   "InternalError", // UnsupportedRequest
		ValueOutOfRange:      "InternalError", // ValueOutOfRange
		ValuePositiveOnly:    "InternalError", // ValuePositiveOnly
		ValueTooHigh:         "InternalError", // ValueTooHigh
		ValueTooLow:          "InternalError", // ValueTooLow
		ValueZeroNotAllowed:  "InternalError", // ValueZeroNotAllowed
		WriteOnly:            "InternalError", // WriteOnly

		// ocpp message的訊息(error code)

		FormatViolation: "InternalError", // FormatViolation
		GenericError:    "InternalError", // GenericError
		// InternalError:                 "InternalError", //(跟reason code有重複)
		MessageTypeNotSupported:       "InternalError", // MessageTypeNotSupported
		NotImplemented:                "InternalError", // NotImplemented
		NotSupported:                  "InternalError", // NotSupported
		OccurrenceConstraintViolation: "InternalError", // OccurrenceConstraintViolation
		PropertyConstraintViolation:   "InternalError", // PropertyConstraintViolation
		ProtocolError:                 "InternalError", // ProtocolError
		RpcFrameworkError:             "InternalError", // RpcFrameworkError
		SecurityError:                 "InternalError", // SecurityError
		TypeConstraintViolation:       "InternalError", // TypeConstraintViolation

		// GRPC Error
		GrpcConnectError:     "InternalError", // GRPC連線錯誤
		GrpcNotFoundCMSError: "InternalError", // GRPC無法連線到FMS錯誤
		GrpcAlreadyExist:     "InternalError", // GRPC連線已存在
		GrpcNotExist:         "InternalError", // GRPC連線不存在

		// T Box error
		TBoxParamError:         "ParamsError",                      // tbox參數錯誤
		TBoxNotFound:           "NotOnLineError",                   // tbox查無資料
		TBoxFwUpdateInProgress: "InternalError",                    // Firmware 更新中
		TBoxAddNFCInProgress:   "DuringNFCCardModificationLocally", //
		TBoxNFCAboveMax:        "AboveMaximumNFCCardIDs",           // 已達NFC卡上限
		TBoxDuplicateNFC:       "DuplicateNFCCardID",               // 重複的 nfc card
		TBoxNFCNotFound:        "NFCCardIDNotExist",                // card id 不存在

		// Image hosting Error
		ImageHostingConnectError: "InternalError", // Image Hosting 連線錯誤
		PutObjectError:           "InternalError", // Image Hosting 上傳物件錯誤
		GenerateUrlError:         "InternalError", // Image Hosting 製作產生錯誤
		ImageHostingObjectError:  "InternalError", // Image Hosting 物件異常
		Base64DecoderError:       "InternalError", // Base64 解碼錯誤
		PhotoEncodeError:         "InternalError", // 照片編碼錯誤

		// Telegram錯誤
		TgNotFoundUser:    "InternalError", //  該使用者不在可接受的範圍內
		TgSetCommandError: "InternalError", // 無法設定初始化tg command
	}
	for code, message := range Errorcodes {
		errorhandler.Errorcodes[code] = message
	}
}
