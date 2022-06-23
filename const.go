package paypal

const (
	// HTTPHeaderUserAgent the user agent name in request header
	HTTPHeaderUserAgent = "User-Agent"
	// HTTPHeaderToken the token name in request header
	HTTPHeaderToken = "Authorization"
	// HTTPHeaderContentType the content-type name in request header
	HTTPHeaderContentType = "content-type"
)

const (
	// SDK_VERSION sdk version
	SDK_VERSION = "v0.0.1"
	// PROJECT_NAME project name
	PROJECT_NAME = "paypal"
	// SANDBOX_ENDPOINT sandbox api endpoint
	SANDBOX_ENDPOINT = "https://api-m.sandbox.paypal.com"
	// LIVE_ENDPOINT live endpoint
	LIVE_ENDPOINT = "https://api-m.paypal.com"
)

type SubscriptionPlanStatus string

const (
	SubscriptionPlanStatusCreated  SubscriptionPlanStatus = "CREATED"
	SubscriptionPlanStatusInactive SubscriptionPlanStatus = "INACTIVE"
	SubscriptionPlanStatusActive   SubscriptionPlanStatus = "ACTIVE"
)

type BillingPlanStatus string

const (
	BillingPlanStatusActive BillingPlanStatus = "ACTIVE"
)

type IntervalUnit string

const (
	IntervalUnitDay   IntervalUnit = "DAY"
	IntervalUnitWeek  IntervalUnit = "WEEK"
	IntervalUnitMonth IntervalUnit = "MONTH"
	IntervalUnitYear  IntervalUnit = "YEAR"
)

// PatchOp patch operate
type PatchOp string

const (
	PatchOpReplace PatchOp = "replace"
	PatchOpAdd     PatchOp = "add"
	PatchOpRemove  PatchOp = "remove"
)

type Intent string

const (
	IntentCapture   Intent = "CAPTURE"
	IntentAuthorize Intent = "AUTHORIZE"
)

type TenureType string

const (
	TenureTypeRegular TenureType = "REGULAR"
	TenureTypeTrial   TenureType = "TRIAL"
)

type SetupFeeFailureAction string

const (
	SetupFeeFailureActionContinue SetupFeeFailureAction = "CONTINUE"
	SetupFeeFailureActionCancel   SetupFeeFailureAction = "CANCEL"
)

type ShippingPreference string

const (
	ShippingPreferenceGetFromFile        ShippingPreference = "GET_FROM_FILE"
	ShippingPreferenceNoShipping         ShippingPreference = "NO_SHIPPING"
	ShippingPreferenceSetProvidedAddress ShippingPreference = "SET_PROVIDED_ADDRESS"
)

type UserAction string

const (
	UserActionContinue     UserAction = "CONTINUE"
	UserActionPayNow       UserAction = "PAY_NOW"
	UserActionSubscribeNow UserAction = "SUBSCRIBE_NOW"
)

type SubscriptionStatus string

const (
	SubscriptionStatusApprovalPending SubscriptionStatus = "APPROVAL_PENDING"
	SubscriptionStatusApproved        SubscriptionStatus = "APPROVED"
	SubscriptionStatusActive          SubscriptionStatus = "ACTIVE"
	SubscriptionStatusSuspended       SubscriptionStatus = "SUSPENDED"
	SubscriptionStatusCancelled       SubscriptionStatus = "CANCELLED"
	SubscriptionStatusExpired         SubscriptionStatus = "EXPIRED"
)

//Doc: https://developer.paypal.com/docs/api/subscriptions/v1/#definition-transaction
type SubscriptionTransactionStatus string

const (
	SubscriptionCaptureStatusCompleted         SubscriptionTransactionStatus = "COMPLETED"
	SubscriptionCaptureStatusDeclined          SubscriptionTransactionStatus = "DECLINED"
	SubscriptionCaptureStatusPartiallyRefunded SubscriptionTransactionStatus = "PARTIALLY_REFUNDED"
	SubscriptionCaptureStatusPending           SubscriptionTransactionStatus = "PENDING"
	SubscriptionCaptureStatusRefunded          SubscriptionTransactionStatus = "REFUNDED"
)

type CaptureType string

const (
	CaptureTypeOutstandingBalance CaptureType = "OUTSTANDING_BALANCE"
)

type ProductType string

// ProductCategory Doc: https://developer.paypal.com/docs/api/catalog-products/v1/#definition-product_category
type ProductCategory string

const (
	ProductTypePhysical ProductType = "PHYSICAL"
	ProductTypeDigital  ProductType = "DIGITAL"
	ProductTypeService  ProductType = "SERVICE"

	ProductCategorySoftware                                  ProductCategory = "SOFTWARE"
	ProductCategorySoftwareComputerAndDataProcessingServices ProductCategory = "COMPUTER_AND_DATA_PROCESSING_SERVICES"
	ProductCategorySoftwareDigitalGames                      ProductCategory = "DIGITAL_GAMES"
	ProductCategorySoftwareGameSoftware                      ProductCategory = "GAME_SOFTWARE"
	ProductCategorySoftwareGames                             ProductCategory = "GAMES"
	ProductCategorySoftwareGeneral                           ProductCategory = "GENERAL"
	ProductCategorySoftwareGraphicAndCommercialDesign        ProductCategory = "GRAPHIC_AND_COMMERCIAL_DESIGN"
	ProductCategorySoftwareOemSoftware                       ProductCategory = "OEM_SOFTWARE"
	ProductCategorySoftwareOnlineGaming                      ProductCategory = "ONLINE_GAMING"
	ProductCategorySoftwareOnlineGamingCurrency              ProductCategory = "ONLINE_GAMING_CURRENCY"
	ProductCategorySoftwareOnlineServices                    ProductCategory = "ONLINE_SERVICES"
	ProductCategorySoftwareOther                             ProductCategory = "OTHER"
	ProductCategorySoftwareServices                          ProductCategory = "SERVICES"
)

// PayeePreferred Doc: https://developer.paypal.com/api/orders/v2/#definition-payment_method
type PayeePreferred string

const (
	PayeePreferredUnrestricted             PayeePreferred = "UNRESTRICTED"
	PayeePreferredImmediatePaymentRequired PayeePreferred = "IMMEDIATE_PAYMENT_REQUIRED"
)

// StandardEntryClassCode Doc: https://developer.paypal.com/api/orders/v2/#definition-payment_method
type StandardEntryClassCode string

const (
	StandardEntryClassCodeTel StandardEntryClassCode = "TEL"
	StandardEntryClassCodeWeb StandardEntryClassCode = "WEB"
	StandardEntryClassCodeCcd StandardEntryClassCode = "CCD"
	StandardEntryClassCodePpd StandardEntryClassCode = "PPD"
)
