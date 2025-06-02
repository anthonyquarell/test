package constant

type Err string

func (e Err) Error() string {
	return string(e)
}

// common errors
const (
	NoRows                 = Err("err_no_rows")
	ServiceNA              = Err("service_not_available")
	NotAuthorized          = Err("not_authorized")
	ObjectNotFound         = Err("object_not_found")
	IncorrectPageSize      = Err("incorrect_page_size")
	BadStatusCode          = Err("bad_status_code")
	InvalidId              = Err("invalid_id")
	SendLimitExceeded      = Err("send_limit_exceeded")
	SecurityCodeNotValid   = Err("security_code_not_valid")
	SecurityCodeHasNotSent = Err("security_code_has_not_sent")
	FrequentRequests       = Err("frequent_requests")
	OrderBlocked           = Err("order_blocked")
	AlreadyVerified        = Err("already_verified")
	VerificationDenied     = Err("verification_denied")
	RouteIdRequired        = Err("route_id_required")
	DbNameRequired         = Err("db_name_required")
	OrderIdRequired        = Err("order_id_required")
	InvalidPhone           = Err("invalid_phone")
	EmptyData              = Err("empty_data")
	CheckoutNotAllowed     = Err("checkout_not_allowed")
	CourierIinRequired     = Err("courier_iin_required")
	ImeiCodeNotValid       = Err("imei_code_not_valid")

	// user errors
	UserNameRequired  = Err("user_name_required")
	UserIdRequired    = Err("user_id_required")
	UserPhoneRequired = Err("user_phone_required")
	PhoneAlreadyExist = Err("phone_already_exist")
	BadUserPhone      = Err("bad_user_phone")
	UserDeactivated   = Err("user_deactivated")

	// role errors
	RoleNameRequired = Err("role_name_required")
	RoleIdRequired   = Err("role_id_required")
)

// ErrFull
type ErrFull struct {
	Err    error
	Desc   string
	Fields map[string]string
}

func (e ErrFull) Error() string {
	return e.Err.Error() + ", desc: " + e.Desc
}
