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

// LandingPage https://developer.paypal.com/docs/api/orders/v2/#definition-application_context
type LandingPage string

const (
	// LandingPageLogin  When the customer clicks PayPal Checkout,
	// the customer is redirected to a page to log in to PayPal and approve the payment.
	LandingPageLogin LandingPage = "LOGIN"
	// LandingPageBilling When the customer clicks PayPal Checkout,
	// the customer is redirected to a page to enter credit
	// or debit card and other relevant billing information required to complete the purchase.
	LandingPageBilling LandingPage = "BILLING"
	// LandingPageNoPreference When the customer clicks PayPal Checkout,
	// the customer is redirected to either a page to log in to
	// PayPal and approve the payment or to a page to enter credit or
	// debit card and other relevant billing information required to
	// complete the purchase, depending on their previous interaction with PayPal.
	LandingPageNoPreference LandingPage = "NO_PREFERENCE"
)

type SubscriptionPlanStatus string

const (
	SubscriptionPlanStatusCreated  SubscriptionPlanStatus = "CREATED"
	SubscriptionPlanStatusInactive SubscriptionPlanStatus = "INACTIVE"
	SubscriptionPlanStatusActive   SubscriptionPlanStatus = "ACTIVE"
)

type VerificationStatus string

const (
	VerificationStatusSuccess VerificationStatus = "SUCCESS"
	VerificationStatusFailure VerificationStatus = "FAILURE"
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

// ShippingPreference The shipping preference:
// Displays the shipping address to the customer.
// Enables the customer to choose an address on the PayPal site.
// Restricts the customer from changing the address during t
// he payment-approval process.
type ShippingPreference string

const (
	// ShippingPreferenceGetFromFile Use the customer-provided shipping address
	// on the PayPal site.
	ShippingPreferenceGetFromFile ShippingPreference = "GET_FROM_FILE"
	// ShippingPreferenceNoShipping Redact the shipping address from the PayPal site.
	// Recommended for digital goods.
	ShippingPreferenceNoShipping ShippingPreference = "NO_SHIPPING"
	// ShippingPreferenceSetProvidedAddress  Use the merchant-provided address.
	// The customer cannot change this address on the PayPal site.
	ShippingPreferenceSetProvidedAddress ShippingPreference = "SET_PROVIDED_ADDRESS"
)

// ProcessingInstruction https://developer.paypal.com/docs/api/orders/v2/#orders_create
type ProcessingInstruction string

const (
	// ProcessingInstructionOrderCompleteOnPaymentApproval API Caller expects the Order to be auto completed
	// (i.e. for PayPal to authorize or capture depending on the intent)
	// on completion of payer approval.
	// This option is not relevant for payment_source that typically
	// do not require a payer approval or interaction.
	// This option is currently only available for the following payment_source:
	// Alipay, Bancontact, BLIK, boletobancario,
	// eps, giropay, iDEAL, Multibanco, MyBank,
	// OXXO, P24, PayU, PUI, SafetyPay, SatisPay, Sofort, Trustly,
	// TrustPay, Verkkopankki, WeChat Pay
	ProcessingInstructionOrderCompleteOnPaymentApproval ProcessingInstruction = "ORDER_COMPLETE_ON_PAYMENT_APPROVAL"
	// ProcessingInstructionNoInstruction The API caller intends to authorize
	// v2/checkout/orders/id/authorize or capture v2/checkout/orders/id/capture after the payer approves the order.
	ProcessingInstructionNoInstruction ProcessingInstruction = "NO_INSTRUCTION"
)

// UserAction https://developer.paypal.com/docs/api/orders/v2/#definition-application_context
type UserAction string

const (
	// UserActionContinue After you redirect the customer to the PayPal payment page,
	// a Continue button appears.
	// Use this option when the final amount is not known
	// when the checkout flow is initiated and you want to redirect
	// the customer to the merchant page without processing the payment.
	UserActionContinue UserAction = "CONTINUE"
	// UserActionPayNow After you redirect the customer to the PayPal payment page,
	// a Pay Now button appears.
	// Use this option when the final amount is known
	// when the checkout is initiated and you want to process
	// the payment immediately when the customer clicks
	UserActionPayNow UserAction = "PAY_NOW"
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
	// PayeePreferredUnrestricted  Accepts any type of payment from the customer.
	PayeePreferredUnrestricted PayeePreferred = "UNRESTRICTED"
	// PayeePreferredImmediatePaymentRequired Accepts only immediate payment
	// from the customer. For example, credit card, PayPal balance,
	// or instant ACH. Ensures that at the time of capture,
	// the payment does not have the `pending` status.
	PayeePreferredImmediatePaymentRequired PayeePreferred = "IMMEDIATE_PAYMENT_REQUIRED"
)

// StandardEntryClassCode Doc: https://developer.paypal.com/api/orders/v2/#definition-payment_method
type StandardEntryClassCode string

const (
	// StandardEntryClassCodeTel The API caller (merchant/partner) accepts
	// authorization and payment information from a consumer over the telephone.
	StandardEntryClassCodeTel StandardEntryClassCode = "TEL"
	// StandardEntryClassCodeWeb The API caller (merchant/partner) accepts Debit transactions
	// from a consumer on their website.
	StandardEntryClassCodeWeb StandardEntryClassCode = "WEB"
	// StandardEntryClassCodeCcd Cash concentration and disbursement
	// for corporate debit transaction. Used to disburse or consolidate funds.
	// Entries are usually Optional high-dollar, low-volume, and time-critical.
	// (e.g. intra-company transfers or invoice payments to suppliers).
	StandardEntryClassCodeCcd StandardEntryClassCode = "CCD"
	// StandardEntryClassCodePpd Prearranged payment and deposit entries.
	// Used for debit payments authorized by a consumer account holder,
	// and usually initiated by a company. These are usually recurring debits (such as insurance premiums).
	StandardEntryClassCodePpd StandardEntryClassCode = "PPD"
)
