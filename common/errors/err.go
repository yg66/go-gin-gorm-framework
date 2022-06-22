package errors

import "encoding/json"

type Err struct {
	Code    int
	Message string
}

const (
	ServerError                   = 1000
	NotFound                      = 1001
	UnknownError                  = 1002
	ParameterError                = 1003
	NetworkAnomaly                = 1004
	OperationTooFrequently        = 1005
	EmailRegFailed                = 1006
	PasswordRegFailed             = 1007
	EmailCodeVerifyFailed         = 1008
	EmailHasBeenUsed              = 1009
	AccountNotExist               = 1010
	RecaptchaVerifyError          = 1011
	DataNotFound                  = 1012
	ApiTokenNameUnrepeatable      = 1013
	ApiTokenUltraLimit            = 1014
	FileUploadFailed              = 1015
	VerifyCodeExpired             = 1016
	DuplicateApply                = 1017
	NotFoundApplyWithEmail        = 1018
	FileUploading                 = 1019
	UserStorageWarn               = 1020
	RequestBodyTooLarge           = 1021
	EmailVerifyExpired            = 1022
	EmailNotVerify                = 1023
	MessageUnavailable            = 1024
	RegionSubscribe               = 1025
	RegionSubscribed              = 1026
	RegionSubscribeReject         = 1027
	RegionDisabled                = 1028
	ApplyStatusAbnormal           = 1029
	ApplyPassed                   = 1030
	RegionUnsubscribe             = 1031
	Unauthorized                  = 1401
	UriNotFoundOrMethodNotSupport = 1404
)

var ErrCodeText = map[int]string{
	ServerError:                   "Server Error",
	NotFound:                      "Not Found",
	UnknownError:                  "Unknown Error",
	ParameterError:                "Parameter Error",
	NetworkAnomaly:                "Network Anomaly",
	OperationTooFrequently:        "Operation too frequently",
	EmailRegFailed:                "E-mail format is incorrect",
	PasswordRegFailed:             "The password must be 8-16 characters and contain letters and numbers",
	EmailCodeVerifyFailed:         "Email Code verify failed, Please try again",
	Unauthorized:                  "User unauthorized or disable",
	EmailHasBeenUsed:              "The mailbox has been used",
	AccountNotExist:               "The current email has not been approved",
	RecaptchaVerifyError:          "ReCAPTCHA verify failed",
	DataNotFound:                  "Data not found",
	ApiTokenNameUnrepeatable:      "The ApiToken name cannot be repeated",
	ApiTokenUltraLimit:            "Each account creation API cannot exceed 20",
	VerifyCodeExpired:             "The verification code has expired",
	DuplicateApply:                "Your application has been submitted, please be patient",
	NotFoundApplyWithEmail:        "No application submitted by this mailbox was found",
	FileUploadFailed:              "File upload failed. Please try again",
	FileUploading:                 "File uploading, please wait patiently",
	UserStorageWarn:               "Your storage capacity could not be queried or insufficient",
	RequestBodyTooLarge:           "http: request body too large",
	EmailVerifyExpired:            "The email is invalid. Please obtain the verification email again",
	EmailNotVerify:                "The mailbox was not authenticated or the application was not submitted",
	MessageUnavailable:            "The message is invalid or unavailable",
	RegionSubscribe:               "The storage region subscription, please wait patiently",
	RegionSubscribed:              "The storage region has already subscribed",
	RegionSubscribeReject:         "The storage region subscribe has been reject",
	RegionDisabled:                "The storage region is unavailable",
	ApplyStatusAbnormal:           "The status of the application is incorrect",
	ApplyPassed:                   "Your previous application has been approved. Please log in via email",
	RegionUnsubscribe:             "You do not subscribe to the storage of this area",
	UriNotFoundOrMethodNotSupport: "Uri not found or method can not support",
}

func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func New(code int) *Err {
	return &Err{
		Code:    code,
		Message: ErrCodeText[code],
	}
}
