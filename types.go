package gocin7corerestapi

type AttributeSet struct {
	ID                string                  `json:"ID,omitempty"`
	Name              string                  `json:"Name,omitempty"`
	Attribute1Name    string                  `json:"Attribute1Name,omitempty"`
	Attribute1Type    string                  `json:"Attribute1Type,omitempty"`
	Attribute1Values  string                  `json:"Attribute1Values,omitempty"`
	Attribute2Name    string                  `json:"Attribute2Name,omitempty"`
	Attribute2Type    string                  `json:"Attribute2Type,omitempty"`
	Attribute2Values  string                  `json:"Attribute2Values,omitempty"`
	Attribute3Name    string                  `json:"Attribute3Name,omitempty"`
	Attribute3Type    string                  `json:"Attribute3Type,omitempty"`
	Attribute3Values  string                  `json:"Attribute3Values,omitempty"`
	Attribute4Name    string                  `json:"Attribute4Name,omitempty"`
	Attribute4Type    string                  `json:"Attribute4Type,omitempty"`
	Attribute4Values  string                  `json:"Attribute4Values,omitempty"`
	Attribute5Name    string                  `json:"Attribute5Name,omitempty"`
	Attribute5Type    string                  `json:"Attribute5Type,omitempty"`
	Attribute5Values  string                  `json:"Attribute5Values,omitempty"`
	Attribute6Name    string                  `json:"Attribute6Name,omitempty"`
	Attribute6Type    string                  `json:"Attribute6Type,omitempty"`
	Attribute6Values  string                  `json:"Attribute6Values,omitempty"`
	Attribute7Name    string                  `json:"Attribute7Name,omitempty"`
	Attribute7Type    string                  `json:"Attribute7Type,omitempty"`
	Attribute7Values  string                  `json:"Attribute7Values,omitempty"`
	Attribute8Name    string                  `json:"Attribute8Name,omitempty"`
	Attribute8Type    string                  `json:"Attribute8Type,omitempty"`
	Attribute8Values  string                  `json:"Attribute8Values,omitempty"`
	Attribute9Name    string                  `json:"Attribute9Name,omitempty"`
	Attribute9Type    string                  `json:"Attribute9Type,omitempty"`
	Attribute9Values  string                  `json:"Attribute9Values,omitempty"`
	Attribute10Name   string                  `json:"Attribute10Name,omitempty"`
	Attribute10Type   string                  `json:"Attribute10Type,omitempty"`
	Attribute10Values string                  `json:"Attribute10Values,omitempty"`
	Attributes        []AttributeSetLineModel `json:"Attributes,omitempty"`
}

type BankAccounts struct {
	AccountID        string  `json:"AccountID,omitempty"`
	Bank             string  `json:"Bank,omitempty"`
	AccountName      string  `json:"AccountName,omitempty"`
	AccountNumber    string  `json:"AccountNumber,omitempty"`
	AccountCode      string  `json:"AccountCode,omitempty"`
	Currency         string  `json:"Currency,omitempty"`
	StatementBalance float64 `json:"StatementBalance,omitempty"`
	BalanceInDear    float64 `json:"BalanceInDear,omitempty"`
	InitialBalance   string  `json:"InitialBalance,omitempty"`
}

type Brand struct {
	ID   string `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
}

type Carrier struct {
	CarrierID   string `json:"CarrierID,omitempty"`
	Description string `json:"Description,omitempty"`
}

type ChartofAccounts struct {
	Code              string `json:"Code,omitempty"`
	Name              string `json:"Name,omitempty"`
	Type              string `json:"Type,omitempty"`
	Status            string `json:"Status,omitempty"`
	Description       string `json:"Description,omitempty"`
	Class             string `json:"Class,omitempty"`
	SystemAccount     string `json:"SystemAccount,omitempty"`
	ForPayments       string `json:"ForPayments,omitempty"`
	DisplayName       string `json:"DisplayName,omitempty"`
	OldCode           string `json:"OldCode,omitempty"`
	Bank              string `json:"Bank,omitempty"`
	BankAccountNumber string `json:"BankAccountNumber,omitempty"`
	BankAccountId     string `json:"BankAccountId,omitempty"`
	Currency          string `json:"Currency,omitempty"`
}

type Customer struct {
	ID                    string                 `json:"ID,omitempty"`
	Name                  string                 `json:"Name,omitempty"`
	Status                string                 `json:"Status,omitempty"`
	Currency              string                 `json:"Currency,omitempty"`
	PaymentTerm           string                 `json:"PaymentTerm,omitempty"`
	AccountReceivable     string                 `json:"AccountReceivable,omitempty"`
	RevenueAccount        string                 `json:"RevenueAccount,omitempty"`
	TaxRule               string                 `json:"TaxRule,omitempty"`
	PriceTier             string                 `json:"PriceTier,omitempty"`
	Carrier               string                 `json:"Carrier,omitempty"`
	SalesRepresentative   string                 `json:"SalesRepresentative,omitempty"`
	Location              string                 `json:"Location,omitempty"`
	Discount              int                    `json:"Discount,omitempty"`
	Comments              string                 `json:"Comments,omitempty"`
	TaxNumber             int                    `json:"TaxNumber,omitempty"`
	CreditLimit           int                    `json:"CreditLimit,omitempty"`
	Tags                  string                 `json:"Tags,omitempty"`
	AttributeSet          string                 `json:"AttributeSet,omitempty"`
	AdditionalAttribute1  string                 `json:"AdditionalAttribute1,omitempty"`
	AdditionalAttribute2  string                 `json:"AdditionalAttribute2,omitempty"`
	AdditionalAttribute3  string                 `json:"AdditionalAttribute3,omitempty"`
	AdditionalAttribute4  string                 `json:"AdditionalAttribute4,omitempty"`
	AdditionalAttribute5  string                 `json:"AdditionalAttribute5,omitempty"`
	AdditionalAttribute6  string                 `json:"AdditionalAttribute6,omitempty"`
	AdditionalAttribute7  string                 `json:"AdditionalAttribute7,omitempty"`
	AdditionalAttribute8  string                 `json:"AdditionalAttribute8,omitempty"`
	AdditionalAttribute9  string                 `json:"AdditionalAttribute9,omitempty"`
	AdditionalAttribute10 string                 `json:"AdditionalAttribute10,omitempty"`
	LastModifiedOn        string                 `json:"LastModifiedOn,omitempty"`
	IsOnCreditHold        bool                   `json:"IsOnCreditHold,omitempty"`
	ProductPrices         []ProductPriceModel    `json:"ProductPrices,omitempty"`
	Addresses             []SupplierAddressModel `json:"Addresses,omitempty"`
	Contacts              []SupplierContactModel `json:"Contacts,omitempty"`
}

type DisassemblyList struct {
	TaskID            string  `json:"TaskID,omitempty"`
	DisassemblyNumber string  `json:"DisassemblyNumber,omitempty"`
	ProductID         string  `json:"ProductID,omitempty"`
	ProductCode       string  `json:"ProductCode,omitempty"`
	ProductName       string  `json:"ProductName,omitempty"`
	Quantity          float64 `json:"Quantity,omitempty"`
	LocationID        string  `json:"LocationID,omitempty"`
	Location          float64 `json:"Location,omitempty"`
	Date              string  `json:"Date,omitempty"`
	Status            string  `json:"Status,omitempty"`
}

type Disassembly struct {
	TaskID                 string                             `json:"TaskID,omitempty"`
	DisassemblyNumber      string                             `json:"DisassemblyNumber,omitempty"`
	Status                 string                             `json:"Status,omitempty"`
	ProductID              string                             `json:"ProductID,omitempty"`
	ProductCode            string                             `json:"ProductCode,omitempty"`
	ProductName            string                             `json:"ProductName,omitempty"`
	LocationID             string                             `json:"LocationID,omitempty"`
	Location               float64                            `json:"Location,omitempty"`
	WIPAccount             string                             `json:"WIPAccount,omitempty"`
	Quantity               float64                            `json:"Quantity,omitempty"`
	AssemblyInstructionURL string                             `json:"AssemblyInstructionURL,omitempty"`
	PickLines              []DisassemblyPickLineModel         `json:"PickLines,omitempty"`
	OrderLines             []DisassemblyOrderLineModel        `json:"OrderLines,omitempty"`
	OrderServiceLines      []DisassemblyOrderServiceLineModel `json:"OrderServiceLines,omitempty"`
	Transactions           []TransactionStockLineModel        `json:"Transactions,omitempty"`
	Errors                 []ErrorModel                       `json:"Errors,omitempty"`
}

type FinishedGoodsOrder struct {
	TaskID     string                        `json:"TaskID,omitempty"`
	Status     string                        `json:"Status,omitempty"`
	OrderLines []FinishedGoodsOrderLineModel `json:"OrderLines,omitempty"`
}

type FinishedGoodsPick struct {
	TaskID         string                       `json:"TaskID,omitempty"`
	Status         string                       `json:"Status,omitempty"`
	WIPAccount     string                       `json:"WIPAccount,omitempty"`
	WIPDate        string                       `json:"WIPDate,omitempty"`
	Account        string                       `json:"Account,omitempty"`
	CompletionDate string                       `json:"CompletionDate,omitempty"`
	PickLines      []FinishedGoodsPickLineModel `json:"PickLines,omitempty"`
}

type FixedAssetTypes struct {
	FixedAssetTypeID                   string  `json:"FixedAssetTypeID,omitempty"`
	Name                               string  `json:"Name,omitempty"`
	DepreciationMethod                 string  `json:"DepreciationMethod,omitempty"`
	AveragingMethod                    string  `json:"AveragingMethod,omitempty"`
	Rate                               float64 `json:"Rate,omitempty"`
	EffectiveLife                      float64 `json:"EffectiveLife,omitempty"`
	AssetAccountCode                   float64 `json:"AssetAccountCode,omitempty"`
	AccumulatedDepreciationAccountCode string  `json:"AccumulatedDepreciationAccountCode,omitempty"`
	AssetAccountName                   string  `json:"AssetAccountName,omitempty"`
	AccumulatedDepreciationAccountName string  `json:"AccumulatedDepreciationAccountName,omitempty"`
	DepreciationExpenseAccountName     string  `json:"DepreciationExpenseAccountName,omitempty"`
}

type Location struct {
	ID                   string `json:"ID,omitempty"`
	Name                 string `json:"Name,omitempty"`
	IsDefault            bool   `json:"IsDefault,omitempty"`
	Deprecated           bool   `json:"Deprecated,omitempty"`
	Bins                 []Bin  `json:"Bins,omitempty"`
	FixedAssetsLocation  bool   `json:"FixedAssetsLocation,omitempty"`
	ParentID             string `json:"ParentID,omitempty"`
	ReferenceCount       int    `json:"ReferenceCount,omitempty"`
	AddressLine1         string `json:"AddressLine1,omitempty"`
	AddressLine2         string `json:"AddressLine2,omitempty"`
	AddressCitySuburb    string `json:"AddressCitySuburb,omitempty"`
	AddressStateProvince string `json:"AddressStateProvince,omitempty"`
	AddressZipPostCode   string `json:"AddressZipPostCode,omitempty"`
	AddressCountry       string `json:"AddressCountry,omitempty"`
	PickZones            string `json:"PickZones,omitempty"`
	IsShopfloor          bool   `json:"IsShopfloor,omitempty"`
	IsCoMan              bool   `json:"IsCoMan,omitempty"`
	IsStaging            bool   `json:"IsStaging,omitempty"`
}

type Bin struct {
	ID           string `json:"ID,omitempty"`
	Name         string `json:"Name,omitempty"`
	IsDeprecated bool   `json:"IsDeprecated,omitempty"`
	IsStaging    bool   `json:"IsStaging,omitempty"`
}

type ME struct {
	Company                                   string               `json:"Company,omitempty"`
	Currency                                  string               `json:"Currency,omitempty"`
	TimeZone                                  string               `json:"TimeZone,omitempty"`
	DefaultWeightUnits                        string               `json:"DefaultWeightUnits,omitempty"`
	DefaultDimensionsUnits                    string               `json:"DefaultDimensionsUnits,omitempty"`
	LockDate                                  string               `json:"LockDate,omitempty"`
	OpeningBalanceDate                        string               `json:"OpeningBalanceDate,omitempty"`
	TaxCalculationMethod                      string               `json:"TaxCalculationMethod,omitempty"`
	DefaultSaleTaxRuleId                      string               `json:"DefaultSaleTaxRuleId,omitempty"`
	DefaultSaleTaxRuleName                    string               `json:"DefaultSaleTaxRuleName,omitempty"`
	Maximumfloat64PlacesInQuantity            string               `json:"Maximumfloat32PlacesInQuantity,omitempty"`
	ApplyCustomerDiscountsAfterOtherDiscounts bool                 `json:"ApplyCustomerDiscountsAfterOtherDiscounts,omitempty"`
	DiscountRule                              string               `json:"DiscountRule,omitempty"`
	AutomaticallyApplyDiscounts               bool                 `json:"AutomaticallyApplyDiscounts,omitempty"`
	RoundingTable                             []RoundingTableModel `json:"RoundingTable,omitempty"`
}

type RoundingTableModel struct {
	RangeTo         float64 `json:"RangeTo,omitempty"`
	RoundToNearest  float64 `json:"RoundToNearest,omitempty"`
	AdjustmentRule  string  `json:"AdjustmentRule,omitempty"`
	AdjustmentValue string  `json:"AdjustmentValue,omitempty"`
}

type MeAddress struct {
	AddressID      string `json:"AddressID,omitempty"`
	Line1          string `json:"Line1,omitempty"`
	Line2          string `json:"Line2,omitempty"`
	CitySuburb     string `json:"CitySuburb,omitempty"`
	StateProvince  string `json:"StateProvince,omitempty"`
	ZipPostCode    string `json:"ZipPostCode,omitempty"`
	Country        string `json:"Country,omitempty"`
	Type           string `json:"Type,omitempty"`
	DefaultForType bool   `json:"DefaultForType,omitempty"`
}

type MeContact struct {
	ContactID      string `json:"ContactID,omitempty"`
	Name           string `json:"Name,omitempty"`
	Phone          string `json:"Phone,omitempty"`
	Fax            string `json:"Fax,omitempty"`
	Email          string `json:"Email,omitempty"`
	Website        string `json:"Website,omitempty"`
	Comment        string `json:"Comment,omitempty"`
	Type           string `json:"Type,omitempty"`
	DefaultForType bool   `json:"DefaultForType,omitempty"`
}

type MoneyOperation struct {
	TaskID                 string                      `json:"TaskID,omitempty"`
	TaskType               string                      `json:"TaskType,omitempty"`
	Status                 string                      `json:"Status,omitempty"`
	BankAccount            string                      `json:"BankAccount,omitempty"`
	CurrencyConversionRate float64                     `json:"CurrencyConversionRate,omitempty"`
	SupplierCustomerName   string                      `json:"SupplierCustomerName,omitempty"`
	SupplierID             string                      `json:"SupplierID,omitempty"`
	CustomerID             string                      `json:"CustomerID,omitempty"`
	Reference              string                      `json:"Reference,omitempty"`
	Date                   string                      `json:"Date,omitempty"`
	TaxInclusive           bool                        `json:"TaxInclusive,omitempty"`
	Note                   string                      `json:"Note,omitempty"`
	Lines                  []MoneyTaskLineModel        `json:"Lines,omitempty"`
	Transactions           []TransactionStockLineModel `json:"Transactions,omitempty"`
	Attachments            []AttachmentLineModel       `json:"Attachments,omitempty"`
}

type BankTransfer struct {
	TaskID                 string                      `json:"TaskID,omitempty"`
	Status                 string                      `json:"Status,omitempty"`
	FromAccount            string                      `json:"FromAccount,omitempty"`
	ToAccount              string                      `json:"ToAccount,omitempty"`
	FromAmount             float64                     `json:"FromAmount,omitempty"`
	ToAmount               float64                     `json:"ToAmount,omitempty"`
	CurrencyConversionRate float64                     `json:"CurrencyConversionRate,omitempty"`
	Reference              string                      `json:"Reference,omitempty"`
	Date                   string                      `json:"Date,omitempty"`
	Note                   string                      `json:"Note,omitempty"`
	Transactions           []TransactionStockLineModel `json:"Transactions,omitempty"`
	Attachments            []AttachmentLineModel       `json:"Attachments,omitempty"`
}

type PaymentTerm struct {
	ID        string `json:"ID,omitempty"`
	Name      string `json:"Name,omitempty"`
	Duration  int    `json:"Duration,omitempty"`
	Method    string `json:"Method,omitempty"`
	IsActive  bool   `json:"IsActive,omitempty"`
	IsDefault bool   `json:"IsDefault,omitempty"`
}

type PriceTier struct {
	Code int    `json:"Code,omitempty"`
	Name string `json:"Name,omitempty"`
}

type Product struct {
	ID                           string                       `json:"ID,omitempty"`
	SKU                          string                       `json:"SKU,omitempty"`
	Name                         string                       `json:"Name,omitempty"`
	Category                     string                       `json:"Category,omitempty"`
	Brand                        string                       `json:"Brand,omitempty"`
	Type                         string                       `json:"Type,omitempty"`
	CostingMethod                string                       `json:"CostingMethod,omitempty"`
	DropShipMode                 string                       `json:"DropShipMode,omitempty"`
	DefaultLocation              string                       `json:"DefaultLocation,omitempty"`
	Length                       float64                      `json:"Length,omitempty"`
	Width                        float64                      `json:"Width,omitempty"`
	Height                       float64                      `json:"Height,omitempty"`
	Weight                       float64                      `json:"Weight,omitempty"`
	CartonLength                 float64                      `json:"CartonLength,omitempty"`
	CartonWidth                  float64                      `json:"CartonWidth,omitempty"`
	CartonHeight                 float64                      `json:"CartonHeight,omitempty"`
	CartonQuantity               float64                      `json:"CartonQuantity,omitempty"`
	CartonInnerQuantity          float64                      `json:"CartonInnerQuantity,omitempty"`
	UOM                          string                       `json:"UOM,omitempty"`
	WeightUnits                  string                       `json:"WeightUnits,omitempty"`
	DimensionsUnits              string                       `json:"DimensionsUnits,omitempty"`
	Barcode                      string                       `json:"Barcode,omitempty"`
	MinimumBeforeReorder         float64                      `json:"MinimumBeforeReorder,omitempty"`
	ReorderQuantity              float64                      `json:"ReorderQuantity,omitempty"`
	PriceTier1                   float64                      `json:"PriceTier1,omitempty"`
	PriceTier2                   float64                      `json:"PriceTier2,omitempty"`
	PriceTier3                   float64                      `json:"PriceTier3,omitempty"`
	PriceTier4                   float64                      `json:"PriceTier4,omitempty"`
	PriceTier5                   float64                      `json:"PriceTier5,omitempty"`
	PriceTier6                   float64                      `json:"PriceTier6,omitempty"`
	PriceTier7                   float64                      `json:"PriceTier7,omitempty"`
	PriceTier8                   float64                      `json:"PriceTier8,omitempty"`
	PriceTier9                   float64                      `json:"PriceTier9,omitempty"`
	PriceTier10                  float64                      `json:"PriceTier10,omitempty"`
	PriceTiers                   []PriceTierModel             `json:"PriceTiers,omitempty"`
	AverageCost                  float64                      `json:"AverageCost,omitempty"`
	ShortDescription             string                       `json:"ShortDescription,omitempty"`
	Description                  string                       `json:"Description,omitempty"`
	InternalNote                 string                       `json:"InternalNote,omitempty"`
	AdditionalAttribute1         string                       `json:"AdditionalAttribute1,omitempty"`
	AdditionalAttribute2         string                       `json:"AdditionalAttribute2,omitempty"`
	AdditionalAttribute3         string                       `json:"AdditionalAttribute3,omitempty"`
	AdditionalAttribute4         string                       `json:"AdditionalAttribute4,omitempty"`
	AdditionalAttribute5         string                       `json:"AdditionalAttribute5,omitempty"`
	AdditionalAttribute6         string                       `json:"AdditionalAttribute6,omitempty"`
	AdditionalAttribute7         string                       `json:"AdditionalAttribute7,omitempty"`
	AdditionalAttribute8         string                       `json:"AdditionalAttribute8,omitempty"`
	AdditionalAttribute9         string                       `json:"AdditionalAttribute9,omitempty"`
	AdditionalAttribute10        string                       `json:"AdditionalAttribute10,omitempty"`
	AttributeSet                 string                       `json:"AttributeSet,omitempty"`
	DiscountRule                 string                       `json:"DiscountRule,omitempty"`
	Tags                         string                       `json:"Tags,omitempty"`
	Status                       string                       `json:"Status,omitempty"`
	StockLocator                 string                       `json:"StockLocator,omitempty"`
	COGSAccount                  string                       `json:"COGSAccount,omitempty"`
	RevenueAccount               string                       `json:"RevenueAccount,omitempty"`
	ExpenseAccount               string                       `json:"ExpenseAccount,omitempty"`
	InventoryAccount             string                       `json:"InventoryAccount,omitempty"`
	PurchaseTaxRule              string                       `json:"PurchaseTaxRule,omitempty"`
	SaleTaxRule                  string                       `json:"SaleTaxRule,omitempty"`
	LastModifiedOn               string                       `json:"LastModifiedOn,omitempty"`
	Sellable                     bool                         `json:"Sellable,omitempty"`
	PickZones                    string                       `json:"PickZones,omitempty"`
	BillOfMaterial               bool                         `json:"BillOfMaterial,omitempty"`
	AutoAssembly                 bool                         `json:"AutoAssembly,omitempty"`
	AutoDisassembly              bool                         `json:"AutoDisassembly,omitempty"`
	QuantityToProduce            float64                      `json:"QuantityToProduce,omitempty"`
	AssemblyInstructionURL       string                       `json:"AssemblyInstructionURL,omitempty"`
	AssemblyCostEstimationMethod string                       `json:"AssemblyCostEstimationMethod,omitempty"`
	BOMType                      string                       `json:"BOMType,omitempty"`
	HSCode                       string                       `json:"HSCode,omitempty"`
	CountryOfOrigin              string                       `json:"CountryOfOrigin,omitempty"`
	CountryOfOriginCode          string                       `json:"CountryOfOriginCode,omitempty"`
	Suppliers                    []ProductSupplierModel       `json:"Suppliers,omitempty"`
	ReorderLevels                []ReorderLevelModel          `json:"ReorderLevels,omitempty"`
	BillOfMaterialsProducts      []BillOfMaterialProductModel `json:"BillOfMaterialsProducts,omitempty"`
	BillOfMaterialsServices      []BillOfMaterialServiceModel `json:"BillOfMaterialsServices,omitempty"`
	Movements                    []ProductMovementModel       `json:"Movements,omitempty"`
	Attachments                  []AttachmentLineModel        `json:"Attachments,omitempty"`
	CustomPrices                 []ProductPriceModel          `json:"CustomPrices,omitempty"`
}

type ProductAvailability struct {
	ID          string  `json:"ID,omitempty"`
	SKU         string  `json:"SKU,omitempty"`
	Name        string  `json:"Name,omitempty"`
	Barcode     string  `json:"Barcode,omitempty"`
	Location    string  `json:"Location,omitempty"`
	Bin         string  `json:"Bin,omitempty"`
	Batch       string  `json:"Batch,omitempty"`
	ExpiryDate  string  `json:"ExpiryDate,omitempty"`
	OnHand      float64 `json:"OnHand,omitempty"`
	Allocated   float64 `json:"Allocated,omitempty"`
	Available   float64 `json:"Available,omitempty"`
	OnOrder     float64 `json:"OnOrder,omitempty"`
	StockOnHand float64 `json:"StockOnHand,omitempty"`
	InTransit   float64 `json:"InTransit,omitempty"`
}

type ProductAvailablityPaginatedResponse struct {
	Total                   int                   `json:"Total,omitempty"`
	Page                    int                   `json:"Page,omitempty"`
	ProductAvailabilityList []ProductAvailability `json:"ProductAvailabilityList,omitempty"`
}

type ProductCategory struct {
	ID   string `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
}

type ProductFamily struct {
	Products             []ProductFamilyProductLineModel `json:"Products,omitempty"`
	ID                   string                          `json:"ID,omitempty"`
	SKU                  string                          `json:"SKU,omitempty"`
	Name                 string                          `json:"Name,omitempty"`
	Category             string                          `json:"Category,omitempty"`
	Brand                string                          `json:"Brand,omitempty"`
	CostingMethod        string                          `json:"CostingMethod,omitempty"`
	DefaultLocation      string                          `json:"DefaultLocation,omitempty"`
	UOM                  string                          `json:"UOM,omitempty"`
	MinimumBeforeReorder *float64                        `json:"MinimumBeforeReorder,omitempty"`
	ReorderQuantity      *float64                        `json:"ReorderQuantity,omitempty"`
	PriceTier1           float64                         `json:"PriceTier1,omitempty"`
	PriceTier2           float64                         `json:"PriceTier2,omitempty"`
	PriceTier3           float64                         `json:"PriceTier3,omitempty"`
	PriceTier4           float64                         `json:"PriceTier4,omitempty"`
	PriceTier5           float64                         `json:"PriceTier5,omitempty"`
	PriceTier6           float64                         `json:"PriceTier6,omitempty"`
	PriceTier7           float64                         `json:"PriceTier7,omitempty"`
	PriceTier8           float64                         `json:"PriceTier8,omitempty"`
	PriceTier9           float64                         `json:"PriceTier9,omitempty"`
	PriceTier10          float64                         `json:"PriceTier10,omitempty"`
	ShortDescription     string                          `json:"ShortDescription,omitempty"`
	Description          string                          `json:"Description,omitempty"`
	AttributeSet         string                          `json:"AttributeSet,omitempty"`
	DiscountRule         string                          `json:"DiscountRule,omitempty"`
	Tags                 string                          `json:"Tags,omitempty"`
	COGSAccount          string                          `json:"COGSAccount,omitempty"`
	RevenueAccount       string                          `json:"RevenueAccount,omitempty"`
	InventoryAccount     string                          `json:"InventoryAccount,omitempty"`
	PurchaseTaxRule      string                          `json:"PurchaseTaxRule,omitempty"`
	SaleTaxRule          string                          `json:"SaleTaxRule,omitempty"`
	DropShipMode         string                          `json:"DropShipMode,omitempty"`
	Option1Name          string                          `json:"Option1Name,omitempty"`
	Option2Name          string                          `json:"Option2Name,omitempty"`
	Option3Name          string                          `json:"Option3Name,omitempty"`
	Option1Values        string                          `json:"Option1Values,omitempty"`
	Option2Values        string                          `json:"Option2Values,omitempty"`
	Option3Values        string                          `json:"Option3Values,omitempty"`
	LastModifiedOn       *string                         `json:"LastModifiedOn,omitempty"`
	Attachments          []AttachmentLineModel           `json:"Attachments,omitempty"`
	HSCode               string                          `json:"HSCode,omitempty"`
	CountryOfOrigin      string                          `json:"CountryOfOrigin,omitempty"`
	CountryOfOriginCode  string                          `json:"CountryOfOriginCode,omitempty"`
}

type MarkupPricesModel struct {
	ProductID    string                 `json:"ProductID,omitempty"`
	MarkupPrices []MarkupPriceLineModel `json:"MarkupPrices,omitempty"`
}

type MarkupPriceLineModel struct {
	TierNumber   int     `json:"TierNumber,omitempty"`
	MarkupType   string  `json:"MarkupType,omitempty"`
	UsePriceType string  `json:"UsePriceType,omitempty"`
	MarkupValue  float64 `json:"MarkupValue,omitempty"`
}

type FactoryCalendar struct {
	Year                       string                      `json:"Year,omitempty"`
	WeekStart                  string                      `json:"WeekStart,omitempty"`
	FactoryCalendarDays        []FactoryCalendarDay        `json:"FactoryCalendarDays,omitempty"`
	FactoryCalendarSpecialDays []FactoryCalendarSpecialDay `json:"FactoryCalendarSpecialDays,omitempty"`
}

type FactoryCalendarDay struct {
	IsWeekend      bool `json:"IsWeekend,omitempty"`
	StartTime      int  `json:"StartTime,omitempty"`
	EndTime        *int `json:"EndTime,omitempty"`
	BreakStartTime *int `json:"BreakStartTime,omitempty"`
	BreakEndTime   *int `json:"BreakEndTime,omitempty"`
	Capacity       int  `json:"Capacity,omitempty"`
}

type FactoryCalendarSpecialDay struct {
	Date           string `json:"Date,omitempty"`
	IsHoliday      bool   `json:"IsHoliday,omitempty"`
	IsWeekend      bool   `json:"IsWeekend,omitempty"`
	StartTime      int    `json:"StartTime,omitempty"`
	EndTime        int    `json:"EndTime,omitempty"`
	BreakStartTime *int   `json:"BreakStartTime,omitempty"`
	BreakEndTime   *int   `json:"BreakEndTime,omitempty"`
	Comment        string `json:"Comment,omitempty"`
	Capacity       int    `json:"Capacity,omitempty"`
}

type ProductionBOM struct {
	BOMID                       string                   `json:"BOMID,omitempty"`
	OutputQuantity              float64                  `json:"OutputQuantity,omitempty"`
	BufferPercent               float64                  `json:"BufferPercent,omitempty"`
	InstructionUrl              string                   `json:"InstructionUrl,omitempty"`
	IgnoreCumulativeLeadTime    bool                     `json:"IgnoreCumulativeLeadTime,omitempty"`
	ComponentProductionLeadTime int                      `json:"ComponentProductionLeadTime,omitempty"`
	Version                     int                      `json:"Version,omitempty"`
	Name                        string                   `json:"Name,omitempty"`
	IsDefault                   bool                     `json:"IsDefault,omitempty"`
	IsTracing                   bool                     `json:"IsTracing,omitempty"`
	DeliveryToID                string                   `json:"DeliveryToID,omitempty"`
	DeliveryToName              string                   `json:"DeliveryToName,omitempty"`
	IssueMethod                 string                   `json:"IssueMethod,omitempty"`
	Operations                  []ProductionBOMOperation `json:"Operations,omitempty"`
	IssueMethodComponent        string                   `json:"IssueMethodComponent,omitempty"`
	IssueMethodParameter        string                   `json:"IssueMethodParameter,omitempty"`
	MinQuantity                 float64                  `json:"MinQuantity,omitempty"`
	MaxQuantity                 float64                  `json:"MaxQuantity,omitempty"`
	DeviationPercent            float64                  `json:"DeviationPercent,omitempty"`
	RunSize                     float64                  `json:"RunSize,omitempty"`
}

type ProductionBOMOperation struct {
	OperationID                    string                          `json:"OperationID,omitempty"`
	Order                          int                             `json:"Order,omitempty"`
	Name                           string                          `json:"Name,omitempty"`
	CycleTime                      int                             `json:"CycleTime,omitempty"`
	UnitsPerCycle                  float64                         `json:"UnitsPerCycle,omitempty"`
	WorkCenterID                   string                          `json:"WorkCenterID,omitempty"`
	WorkCenterName                 string                          `json:"WorkCenterName,omitempty"`
	WorkCenterCoManProcurementType string                          `json:"WorkCenterCoManProcurementType,omitempty"`
	OperationType                  string                          `json:"OperationType,omitempty"`
	DeliveryToID                   string                          `json:"DeliveryToID,omitempty"`
	DeliveryToName                 string                          `json:"DeliveryToName,omitempty"`
	IssueMethod                    string                          `json:"IssueMethod,omitempty"`
	IsDropShip                     bool                            `json:"IsDropShip,omitempty"`
	Resources                      []ProductionBOMResource         `json:"Resources,omitempty"`
	Components                     []ProductionBOMComponent        `json:"Components,omitempty"`
	Attachments                    []ProductionBOMAttachment       `json:"Attachments,omitempty"`
	Notes                          []ProductionBOMNote             `json:"Notes,omitempty"`
	OperationLinks                 []ProductionBOMOperationLink    `json:"OperationLinks,omitempty"`
	InputProducts                  []ProductionBOMOperationProduct `json:"InputProducts,omitempty"`
	OutputProducts                 []ProductionBOMOperationProduct `json:"OutputProducts,omitempty"`
	FinishedProducts               []ProductionBOMOperationProduct `json:"FinishedProducts,omitempty"`
	IsBackflush                    bool                            `json:"IsBackflush,omitempty"`
}

type ProductionBOMResource struct {
	BOMResourceID       string  `json:"BOMResourceID,omitempty"`
	ResourceID          string  `json:"ResourceID,omitempty"`
	ResourceName        string  `json:"ResourceName,omitempty"`
	ResourceCode        string  `json:"ResourceCode,omitempty"`
	Quantity            float64 `json:"Quantity,omitempty"`
	Cost                float64 `json:"Cost,omitempty"`
	CycleTime           int     `json:"CycleTime,omitempty"`
	Position            int     `json:"Position,omitempty"`
	CostCalculationType string  `json:"CostCalculationType,omitempty"`
}

type ProductionBOMComponent struct {
	BOMComponentID string  `json:"BOMComponentID,omitempty"`
	ProductID      string  `json:"ProductID,omitempty"`
	ProductSKU     string  `json:"ProductSKU,omitempty"`
	ProductName    string  `json:"ProductName,omitempty"`
	Quantity       float64 `json:"Quantity,omitempty"`
	WastageQty     float64 `json:"WastageQty,omitempty"`
	WastagePercent float64 `json:"WastagePercent,omitempty"`
	Cost           float64 `json:"Cost,omitempty"`
	CostingMethod  string  `json:"CostingMethod,omitempty"`
	Unit           string  `json:"Unit,omitempty"`
	Position       int     `json:"Position,omitempty"`
	SalePriceTier  int     `json:"SalePriceTier,omitempty"`
	IsBackflush    bool    `json:"IsBackflush,omitempty"`
}

type ProductionBOMAttachment struct {
	AttachmentID string `json:"AttachmentID,omitempty"`
	FileName     string `json:"FileName,omitempty"`
	Date         string `json:"Date,omitempty"`
	ContentType  string `json:"ContentType,omitempty"`
	Content      string `json:"Content,omitempty"`
	Position     int    `json:"Position,omitempty"`
}

type ProductionBOMNote struct {
	NoteID   string `json:"NoteID,omitempty"`
	Note     string `json:"Note,omitempty"`
	Position int    `json:"Position,omitempty"`
}

type ProductionBOMOperationLink struct {
	ID                   string `json:"ID,omitempty"`
	RelatedOperationID   string `json:"RelatedOperationID,omitempty"`
	RelatedOperationName string `json:"RelatedOperationName,omitempty"`
	RelationType         int    `json:"RelationType,omitempty"`
}

type ProductionBOMOperationProduct struct {
	ID                  string  `json:"ID,omitempty"`
	ProductID           string  `json:"ProductID,omitempty"`
	ProductSKU          string  `json:"ProductSKU,omitempty"`
	ProductName         string  `json:"ProductName,omitempty"`
	CostCalculationType string  `json:"CostCalculationType,omitempty"`
	PriceTier           int     `json:"PriceTier,omitempty"`
	Ratio               float64 `json:"Ratio,omitempty"`
	OutputQuantity      float64 `json:"OutputQuantity,omitempty"`
	Position            int     `json:"Position,omitempty"`
	CostingMethod       string  `json:"CostingMethod,omitempty"`
	Unit                string  `json:"Unit,omitempty"`
	AverageCost         float64 `json:"AverageCost,omitempty"`
	WastageCost         int     `json:"WastageCost,omitempty"`
	DeliveryTo          string  `json:"DeliveryTo,omitempty"`
	DeliveryToName      string  `json:"DeliveryToName,omitempty"`
}

type ProductionBOMVariationComponent struct {
	BOMVariationComponentID string `json:"BOMVariationComponentID,omitempty"`
	ProductFamilyID         string `json:"ProductFamilyID,omitempty"`
	ProductFamilySKU        string `json:"ProductFamilySKU,omitempty"`
	ProductFamilyName       string `json:"ProductFamilyName,omitempty"`
	QuantitySettingsJSON    string `json:"QuantitySettingsJSON,omitempty"`
	MapVariationsJSON       string `json:"MapVariationsJSON,omitempty"`
	QuantityOptionName      string `json:"QuantityOptionName,omitempty"`
	CostingMethod           string `json:"CostingMethod,omitempty"`
	Unit                    string `json:"Unit,omitempty"`
	Position                int    `json:"Position,omitempty"`
	SalePriceTier           int    `json:"SalePriceTier,omitempty"`
}

// To be fixed -- Fields are not well defined in API documentation
// type ProductionOrder struct {
// 	ProductionOrderID         string                      `json:"ProductionOrderID,omitempty"`
// 	ProductID                 string                      `json:"ProductID,omitempty"`
// 	ProductSKU                string                      `json:"ProductSKU,omitempty"`
// 	ProductName               string                      `json:"ProductName,omitempty"`
// 	OrderNumber               string                      `json:"OrderNumber,omitempty"`
// 	LocationID                string                      `json:"LocationID,omitempty"`
// 	LocationName              string                      `json:"LocationName,omitempty"`
// 	CostingMethod             string                      `json:"CostingMethod,omitempty"`
// 	WarehouseName             string                      `json:"WarehouseName,omitempty"`
// 	Unit                      string                      `json:"Unit,omitempty"`
// 	OrderStatus               string                      `json:"OrderStatus,omitempty"`
// 	Status                    string                      `json:"Status,omitempty"`
// 	InstructionURL            string                      `json:"InstructionURL,omitempty"`
// 	SourceName                string                      `json:"SourceName,omitempty"`
// 	SourceTaskID              string                      `json:"SourceTaskID,omitempty"`
// 	SourceTaskNumber          string                      `json:"SourceTaskNumber,omitempty"`
// 	WIPAccountCode            string                      `json:"WIPAccountCode,omitempty"`
// 	FinishedGoodsAccountCode  string                      `json:"FinishedGoodsAccountCode,omitempty"`
// 	Quantity                  float64                     `json:"Quantity,omitempty"`
// 	BOMQuantity               float64                     `json:"BOMQuantity,omitempty"`
// 	CapacityCalculationType   string                      `json:"CapacityCalculationType,omitempty"`
// 	StartDate                 string                      `json:"StartDate,omitempty"`
// 	ReleaseDate               string                      `json:"ReleaseDate,omitempty"`
// 	RequiredByDate            string                      `json:"RequiredByDate,omitempty"`
// 	CompletionDate            string                      `json:"CompletionDate,omitempty"`
// 	IsIgnoreLeadTime          bool                        `json:"IsIgnoreLeadTime,omitempty"`
// 	WarehouseID               string                      `json:"WarehouseID,omitempty"`
// 	RetailID                  string                      `json:"RetailID,omitempty"`
// 	Comments                  string                      `json:"Comments,omitempty"`
// 	ScheduleStart             string                      `json:"ScheduleStart,omitempty"`
// 	PlannedBy                 string                      `json:"PlannedBy,omitempty"`
// 	ReleasedBy                string                      `json:"ReleasedBy,omitempty"`
// 	OrderCycleTime            int                       `json:"OrderCycleTime,omitempty"`
// 	BOMVersion                int                       `json:"BOMVersion,omitempty"`
// 	BOMName                   string                      `json:"BOMName,omitempty"`
// 	Tags                      string                      `json:"Tags,omitempty"`
// 	CustomField1              string                      `json:"CustomField1,omitempty"`
// 	CustomField2              string                      `json:"CustomField2,omitempty"`
// 	CustomField3              string                      `json:"CustomField3,omitempty"`
// 	CustomField4              string                      `json:"CustomField4,omitempty"`
// 	CustomField5              string                      `json:"CustomField5,omitempty"`
// 	CustomField6              string                      `json:"CustomField6,omitempty"`
// 	CustomField7              string                      `json:"CustomField7,omitempty"`
// 	CustomField8              string                      `json:"CustomField8,omitempty"`
// 	CustomField9              string                      `json:"CustomField9,omitempty"`
// 	CustomField10             string                      `json:"CustomField10,omitempty"`
// 	ProductionOrderOperations []ProductionOrderOperation  `json:"ProductionOrderOperations,omitempty"`
// 	ProductionRuns            []ProductionRun             `json:"ProductionRuns,omitempty"`
// 	ProductionOrderDeliveryTo []ProductionOrderDeliveryTo `json:"ProductionOrderDeliveryTo,omitempty"`
// 	IssueMethodComponent      string                      `json:"IssueMethodComponent,omitempty"`
// 	IssueMethodParameter      string                      `json:"IssueMethodParameter,omitempty"`
// 	SourceTasks               []ProductionOrderSourceTask `json:"SourceTasks,omitempty"`
// 	RunSize                   float64                     `json:"RunSize,omitempty"`
// }

type ProductionOrderOperation struct {
	OperationID                    string                               `json:"OperationID,omitempty"`
	Order                          int                                  `json:"Order,omitempty"`
	Name                           string                               `json:"Name,omitempty"`
	CycleTime                      int                                  `json:"CycleTime,omitempty"`
	UnitsPerCycle                  float64                              `json:"UnitsPerCycle,omitempty"`
	TotalCycleTime                 int                                  `json:"TotalCycleTime,omitempty"`
	TotalUnitsPerCycle             float64                              `json:"TotalUnitsPerCycle,omitempty"`
	OperationType                  string                               `json:"OperationType,omitempty"`
	WorkCenterID                   string                               `json:"WorkCenterID,omitempty"`
	WorkCenterName                 string                               `json:"WorkCenterName,omitempty"`
	WorkCenterCode                 string                               `json:"WorkCenterCode,omitempty"`
	WorkCenterCoManProcurementType string                               `json:"WorkCenterCoManProcurementType,omitempty"`
	ComponentLocationID            string                               `json:"ComponentLocationID,omitempty"`
	IsDropShip                     bool                                 `json:"IsDropShip,omitempty"`
	Attachments                    []ProductionOrderOperationAttachment `json:"Attachments,omitempty"`
	Components                     []ProductionOrderComponent           `json:"Components,omitempty"`
	Notes                          []ProductionOrderOperationNote       `json:"Notes,omitempty"`
	Resources                      []ProductionOrderResource            `json:"Resources,omitempty"`
	OperationLinks                 []ProductionOrderOperationLink       `json:"OperationLinks,omitempty"`
	InputProducts                  []ProductionOrderOperationProduct    `json:"InputProducts,omitempty"`
	OutputProducts                 []ProductionOrderOperationProduct    `json:"OutputProducts,omitempty"`
	FinishedProducts               []ProductionOrderOperationProduct    `json:"FinishedProducts,omitempty"`
	IsBackflush                    bool                                 `json:"IsBackflush,omitempty"`
}

type ProductionOrderOperationAttachment struct {
	AttachmentID string `json:"AttachmentID,omitempty"`
	FileName     string `json:"FileName,omitempty"`
	Date         string `json:"Date,omitempty"`
	ContentType  string `json:"ContentType,omitempty"`
	Position     int    `json:"Position,omitempty"`
	Content      string `json:"Content,omitempty"`
}

type ProductionOrderComponent struct {
	OrderComponentID string  `json:"OrderComponentID,omitempty"`
	ProductID        string  `json:"ProductID,omitempty"`
	ProductSKU       string  `json:"ProductSKU,omitempty"`
	ProductName      string  `json:"ProductName,omitempty"`
	Position         int     `json:"Position,omitempty"`
	Available        float64 `json:"Available,omitempty"`
	Quantity         float64 `json:"Quantity,omitempty"`
	TotalQuantity    float64 `json:"TotalQuantity,omitempty"`
	WastageQty       float64 `json:"WastageQty,omitempty"`
	WastagePercent   float64 `json:"WastagePercent,omitempty"`
	Cost             float64 `json:"Cost,omitempty"`
	TotalCost        float64 `json:"TotalCost,omitempty"`
	ProductCost      float64 `json:"ProductCost,omitempty"`
	CostingMethod    string  `json:"CostingMethod,omitempty"`
	Unit             string  `json:"Unit,omitempty"`
	ProductType      string  `json:"ProductType,omitempty"`
	SalePriceTier    int     `json:"SalePriceTier,omitempty"`
	IsBackflush      bool    `json:"IsBackflush,omitempty"`
}

type ProductionOrderOperationNote struct {
	NoteID   string `json:"NoteID,omitempty"`
	Note     string `json:"Note,omitempty"`
	Position int    `json:"Position,omitempty"`
}

type ProductionOrderResource struct {
	OrderResourceID     string  `json:"OrderResourceID,omitempty"`
	ResourceID          string  `json:"ResourceID,omitempty"`
	ResourceCode        string  `json:"ResourceCode,omitempty"`
	ResourceName        string  `json:"ResourceName,omitempty"`
	Position            int     `json:"Position,omitempty"`
	Quantity            float64 `json:"Quantity,omitempty"`
	Cost                float64 `json:"Cost,omitempty"`
	TotalCost           float64 `json:"TotalCost,omitempty"`
	ResourceCost        float64 `json:"ResourceCost,omitempty"`
	CycleTime           int     `json:"CycleTime,omitempty"`
	CostCalculationType string  `json:"CostCalculationType,omitempty"`
}

type ProductionOrderOperationLink struct {
	ID                 string `json:"ID,omitempty"`
	RelatedOperationID string `json:"RelatedOperationID,omitempty"`
	Position           int    `json:"Position,omitempty"`
}

type ProductionOrderOperationProduct struct {
	ID                  string  `json:"ID,omitempty"`
	ProductID           string  `json:"ProductID,omitempty"`
	ProductSKU          string  `json:"ProductSKU,omitempty"`
	ProductName         string  `json:"ProductName,omitempty"`
	Position            int     `json:"Position,omitempty"`
	CostCalculationType string  `json:"CostCalculationType,omitempty"`
	PriceTier           int     `json:"PriceTier,omitempty"`
	Ratio               float64 `json:"Ratio,omitempty"`
	OutputQuantity      float64 `json:"OutputQuantity,omitempty"`
	TotalOutputQuantity float64 `json:"TotalOutputQuantity,omitempty"`
	Cost                float64 `json:"Cost,omitempty"`
	TotalCost           float64 `json:"TotalCost,omitempty"`
	AverageCost         float64 `json:"AverageCost,omitempty"`
	FixedCost           float64 `json:"FixedCost,omitempty"`
	CostingMethod       string  `json:"CostingMethod,omitempty"`
	Unit                string  `json:"Unit,omitempty"`
	WastageCost         float64 `json:"WastageCost,omitempty"`
}

type ProductionOrderSourceTask struct {
	SourceTaskID     string `json:"SourceTaskID,omitempty"`
	SourceTaskNumber string `json:"SourceTaskNumber,omitempty"`
}

type ProductionOrderListItem struct {
	TaskID                   string                          `json:"TaskID,omitempty"`
	ProductionOrderID        string                          `json:"ProductionOrderID,omitempty"`
	Type                     string                          `json:"Type,omitempty"`
	ProductID                string                          `json:"ProductID,omitempty"`
	ProductSKU               string                          `json:"ProductSKU,omitempty"`
	ProductName              string                          `json:"ProductName,omitempty"`
	OrderNumber              string                          `json:"OrderNumber,omitempty"`
	LocationID               string                          `json:"LocationID,omitempty"`
	LocationName             string                          `json:"LocationName,omitempty"`
	Status                   string                          `json:"Status,omitempty"`
	OrderStatus              string                          `json:"OrderStatus,omitempty"`
	Quantity                 float64                         `json:"Quantity,omitempty"`
	StartDate                string                          `json:"StartDate,omitempty"`
	ReleaseDate              string                          `json:"ReleaseDate,omitempty"`
	RequiredByDate           string                          `json:"RequiredByDate,omitempty"`
	CompletionDate           string                          `json:"CompletionDate,omitempty"`
	Comments                 string                          `json:"Comments,omitempty"`
	CapacityCalculationType  string                          `json:"CapacityCalculationType,omitempty"`
	WIPAccountCode           string                          `json:"WIPAccountCode,omitempty"`
	Tags                     string                          `json:"Tags,omitempty"`
	FinishedGoodsAccountCode string                          `json:"FinishedGoodsAccountCode,omitempty"`
	SourceTaskID             string                          `json:"SourceTaskID,omitempty"`
	SourceTaskNumber         string                          `json:"SourceTaskNumber,omitempty"`
	SourceTaskType           int                             `json:"SourceTaskType,omitempty"`
	IsSourceTaskVoided       bool                            `json:"IsSourceTaskVoided,omitempty"`
	CustomField1             string                          `json:"CustomField1,omitempty"`
	CustomField2             string                          `json:"CustomField2,omitempty"`
	CustomField3             string                          `json:"CustomField3,omitempty"`
	CustomField4             string                          `json:"CustomField4,omitempty"`
	CustomField5             string                          `json:"CustomField5,omitempty"`
	CustomField6             string                          `json:"CustomField6,omitempty"`
	CustomField7             string                          `json:"CustomField7,omitempty"`
	CustomField8             string                          `json:"CustomField8,omitempty"`
	CustomField9             string                          `json:"CustomField9,omitempty"`
	CustomField10            string                          `json:"CustomField10,omitempty"`
	SourceTasks              []ProductionOrderListSourceTask `json:"SourceTasks,omitempty"`
}

type ProductionOrderListSourceTask struct {
	SourceTaskID       string `json:"SourceTaskID,omitempty"`
	SourceTaskNumber   string `json:"SourceTaskNumber,omitempty"`
	SourceTaskType     int    `json:"SourceTaskType,omitempty"`
	IsSourceTaskVoided bool   `json:"IsSourceTaskVoided,omitempty"`
}

type ProductionRun struct {
	RunID                string                       `json:"RunID,omitempty"`
	Number               int                          `json:"Number,omitempty"`
	Status               string                       `json:"Status,omitempty"`
	Quantity             float64                      `json:"Quantity,omitempty"`
	WIPAccount           string                       `json:"WIPAccount,omitempty"`
	StartDate            string                       `json:"StartDate,omitempty"`
	EndDate              string                       `json:"EndDate,omitempty"`
	DueDate              string                       `json:"DueDate,omitempty"`
	ReceivedDate         string                       `json:"ReceivedDate,omitempty"`
	ScheduleStart        string                       `json:"ScheduleStart,omitempty"`
	ScheduleDue          string                       `json:"ScheduleDue,omitempty"`
	Operations           []ProductionRunOperation     `json:"Operations,omitempty"`
	Output               []ProductionRunOutput        `json:"Output,omitempty"`
	PendingOutput        []ProductionRunPendingOutput `json:"PendingOutput,omitempty"`
	ManualJournals       []ProductionRunManualJournal `json:"ManualJournals,omitempty"`
	Traceability         []ProductionRunTraceability  `json:"Traceability,omitempty"`
	IssueMethodParameter string                       `json:"IssueMethodParameter,omitempty"`
}

type ProductionRunOperation struct {
	OperationID      string                               `json:"OperationID,omitempty"`
	Status           string                               `json:"Status,omitempty"`
	Order            int                                  `json:"Order,omitempty"`
	Name             string                               `json:"Name,omitempty"`
	PlannedTime      int                                  `json:"PlannedTime,omitempty"`
	ActualTime       int                                  `json:"ActualTime,omitempty"`
	UnitsPerCycle    float64                              `json:"UnitsPerCycle,omitempty"`
	StartDate        string                               `json:"StartDate,omitempty"`
	EndDate          string                               `json:"EndDate,omitempty"`
	DueDate          string                               `json:"DueDate,omitempty"`
	OperationType    string                               `json:"OperationType,omitempty"`
	WorkCenterID     string                               `json:"WorkCenterID,omitempty"`
	WorkCenterCode   string                               `json:"WorkCenterCode,omitempty"`
	WorkCenterName   string                               `json:"WorkCenterName,omitempty"`
	SuspendReason    string                               `json:"SuspendReason,omitempty"`
	CoManStatus      string                               `json:"CoManStatus,omitempty"`
	Components       []ProductionRunOperationComponent    `json:"Components,omitempty"`
	Resources        []ProductionRunOperationResource     `json:"Resources,omitempty"`
	ResourceCosts    []ProductionRunOperationResourceCost `json:"ResourceCosts,omitempty"`
	FinishedProducts []ProductionRunOperationProduct      `json:"FinishedProducts,omitempty"`
	OutputProducts   []ProductionRunOperationProduct      `json:"OutputProducts,omitempty"`
	InputProducts    []ProductionRunOperationProduct      `json:"InputProducts,omitempty"`
	CoManTasks       []ProductionRunOperationCoManTask    `json:"CoManTasks,omitempty"`
	CoManTaskLines   []ProductionRunOperationCoManLine    `json:"CoManTaskLines,omitempty"`
	Attachments      []ProductionRunOperationAttachment   `json:"Attachments,omitempty"`
	Notes            []ProductionRunOperationNote         `json:"Notes,omitempty"`
	IsBackflush      bool                                 `json:"IsBackflush,omitempty"`
	ManualStartDate  string                               `json:"ManualStartDate,omitempty"`
	ManualEndDate    string                               `json:"ManualEndDate,omitempty"`
}

type ProductionRunOperationComponent struct {
	RunComponentID   string  `json:"RunComponentID,omitempty"`
	ProductID        string  `json:"ProductID,omitempty"`
	ProductCode      string  `json:"ProductCode,omitempty"`
	ProductName      string  `json:"ProductName,omitempty"`
	Quantity         float64 `json:"Quantity,omitempty"`
	ExpectedQuantity float64 `json:"ExpectedQuantity,omitempty"`
	WastageQty       float64 `json:"WastageQty,omitempty"`
	WastagePercent   float64 `json:"WastagePercent,omitempty"`
	BatchSN          string  `json:"BatchSN,omitempty"`
	ExpiryDate       string  `json:"ExpiryDate,omitempty"`
	UnitCost         float64 `json:"UnitCost,omitempty"`
	LocationID       string  `json:"LocationID,omitempty"`
	LocationName     string  `json:"LocationName,omitempty"`
	ProductCost      string  `json:"ProductCost,omitempty"`
	Available        float64 `json:"Available,omitempty"`
	CostingMethod    string  `json:"CostingMethod,omitempty"`
	Unit             string  `json:"Unit,omitempty"`
	ReservedQuantity float64 `json:"ReservedQuantity,omitempty"`
	IsReserved       bool    `json:"IsReserved,omitempty"`
	IsBackflush      bool    `json:"IsBackflush,omitempty"`
}

type ProductionRunOperationResource struct {
	RunResourceID string  `json:"RunResourceID,omitempty"`
	ResourceID    string  `json:"ResourceID,omitempty"`
	ResourceCode  string  `json:"ResourceCode,omitempty"`
	ResourceName  string  `json:"ResourceName,omitempty"`
	UnitCost      float64 `json:"UnitCost,omitempty"`
	Quantity      float64 `json:"Quantity,omitempty"`
	Cost          float64 `json:"Cost,omitempty"`
	CycleTime     float64 `json:"CycleTime,omitempty"`
}

type ProductionRunOperationResourceCost struct {
	RunCostID      string  `json:"RunCostID,omitempty"`
	ProductID      string  `json:"ProductID,omitempty"`
	ProductName    string  `json:"ProductName,omitempty"`
	AccountName    string  `json:"AccountName,omitempty"`
	ExpenseAccount string  `json:"ExpenseAccount,omitempty"`
	Comments       string  `json:"Comments,omitempty"`
	Cost           float64 `json:"Cost,omitempty"`
}

type ProductionRunOperationProduct struct {
	ProductID        string  `json:"ProductID,omitempty"`
	ProductSKU       string  `json:"ProductSKU,omitempty"`
	ProductName      string  `json:"ProductName,omitempty"`
	Unit             string  `json:"Unit,omitempty"`
	OutputQuantity   float64 `json:"OutputQuantity,omitempty"`
	ExpectedQuantity float64 `json:"ExpectedQuantity,omitempty"`
	WastageQuantity  float64 `json:"WastageQuantity,omitempty"`
	BatchSN          string  `json:"BatchSN,omitempty"`
	ExpiryDate       string  `json:"ExpiryDate,omitempty"`
	LocationID       string  `json:"LocationID,omitempty"`
	LocationName     string  `json:"LocationName,omitempty"`
}

type ProductionRunOperationCoManTask struct {
	TaskID     string `json:"TaskID,omitempty"`
	Type       string `json:"Type,omitempty"`
	TaskStatus string `json:"TaskStatus,omitempty"`
	TaskNumber string `json:"TaskNumber,omitempty"`
}

type ProductionRunOperationCoManLine struct {
	TaskID       string  `json:"TaskID,omitempty"`
	ProductID    string  `json:"ProductID,omitempty"`
	ProductCode  string  `json:"ProductCode,omitempty"`
	ProductName  string  `json:"ProductName,omitempty"`
	Unit         string  `json:"Unit,omitempty"`
	Quantity     float64 `json:"Quantity,omitempty"`
	Cost         float64 `json:"Cost,omitempty"`
	BatchSN      string  `json:"BatchSN,omitempty"`
	ExpiryDate   string  `json:"ExpiryDate,omitempty"`
	LocationID   string  `json:"LocationID,omitempty"`
	LocationName string  `json:"LocationName,omitempty"`
	ReceivedDate string  `json:"ReceivedDate,omitempty"`
}

type ProductionRunOperationAttachment struct {
	AttachmentID string `json:"AttachmentID,omitempty"`
	FileName     string `json:"FileName,omitempty"`
	Date         string `json:"Date,omitempty"`
	ContentType  string `json:"ContentType,omitempty"`
	Position     int    `json:"Position,omitempty"`
	Content      string `json:"Content,omitempty"`
}

type ProductionRunOperationNote struct {
	NoteID   string `json:"NoteID,omitempty"`
	Note     string `json:"Note,omitempty"`
	Position int    `json:"Position,omitempty"`
}

type ProductionRunPendingOutput struct {
	ProductID     string `json:"ProductID,omitempty"`
	ProductCode   string `json:"ProductCode,omitempty"`
	ProductName   string `json:"ProductName,omitempty"`
	Unit          string `json:"Unit,omitempty"`
	CostingMethod string `json:"CostingMethod,omitempty"`
	BatchSN       string `json:"BatchSN,omitempty"`
	ExpiryDate    string `json:"ExpiryDate,omitempty"`
}

type ProductionRunOutput struct {
	ProductID       string  `json:"ProductID,omitempty"`
	ProductCode     string  `json:"ProductCode,omitempty"`
	ProductName     string  `json:"ProductName,omitempty"`
	Unit            string  `json:"Unit,omitempty"`
	CostingMethod   string  `json:"CostingMethod,omitempty"`
	BatchSN         string  `json:"BatchSN,omitempty"`
	ExpiryDate      string  `json:"ExpiryDate,omitempty"`
	UnitCost        float64 `json:"UnitCost,omitempty"`
	Quantity        float64 `json:"Quantity,omitempty"`
	WastageQuantity float64 `json:"WastageQuantity,omitempty"`
	ReceivedDate    string  `json:"ReceivedDate,omitempty"`
	Received        bool    `json:"Received,omitempty"`
	LocationID      string  `json:"LocationID,omitempty"`
	LocationName    string  `json:"LocationName,omitempty"`
}

type ProductionRunManualJournal struct {
	JournalID     string  `json:"JournalID,omitempty"`
	Reference     string  `json:"Reference,omitempty"`
	Amount        float64 `json:"Amount,omitempty"`
	Date          string  `json:"Date,omitempty"`
	Debit         string  `json:"Debit,omitempty"`
	Credit        string  `json:"Credit,omitempty"`
	OperationID   string  `json:"OperationID,omitempty"`
	OperationName string  `json:"OperationName,omitempty"`
}

type ProductionRunTraceability struct {
	TraceabilityID            string `json:"TraceabilityID,omitempty"`
	ProducedProductID         string `json:"ProducedProductID,omitempty"`
	ProducedProductBatchSN    string `json:"ProducedProductBatchSN,omitempty"`
	ProducedProductExpiryDate string `json:"ProducedProductExpiryDate,omitempty"`
	UsedProductID             string `json:"UsedProductID,omitempty"`
	UsedProductBatchSN        string `json:"UsedProductBatchSN,omitempty"`
	UsedProductExpiryDate     string `json:"UsedProductExpiryDate,omitempty"`
}

type Resource struct {
	ResourceID              string               `json:"ResourceID,omitempty"`
	Code                    string               `json:"Code,omitempty"`
	Name                    string               `json:"Name,omitempty"`
	ResourceType            string               `json:"ResourceType,omitempty"`
	CycleDuration           int                  `json:"CycleDuration,omitempty"`
	IsActive                bool                 `json:"IsActive,omitempty"`
	IsInfinite              bool                 `json:"IsInfinite,omitempty"`
	IsAllowAllCapacityUsage bool                 `json:"IsAllowAllCapacityUsage,omitempty"`
	IsAvailableOnHolidays   bool                 `json:"IsAvailableOnHolidays,omitempty"`
	IsAvailableOnWeekends   bool                 `json:"IsAvailableOnWeekends,omitempty"`
	Tags                    string               `json:"Tags,omitempty"`
	ResourceCapacities      []ResourceCapacity   `json:"ResourceCapacities,omitempty"`
	ResourceCosts           []ResourceCost       `json:"ResourceCosts,omitempty"`
	ResourceRemarks         []ResourceRemark     `json:"ResourceRemarks,omitempty"`
	ResourceAttachments     []ResourceAttachment `json:"ResourceAttachments,omitempty"`
}

type ResourceCapacity struct {
	LocationID         string             `json:"LocationID,omitempty"`
	LocationName       string             `json:"LocationName,omitempty"`
	ResourceQuantity   int                `json:"ResourceQuantity,omitempty"`
	NonOperationalFrom string             `json:"NonOperationalFrom,omitempty"`
	NonOperationalTo   string             `json:"NonOperationalTo,omitempty"`
	CustomWorkingDays  []CustomWorkingDay `json:"CustomWorkingDays,omitempty"`
	ResourceUnits      []ResourceUnit     `json:"ResourceUnits,omitempty"`
}

type CustomWorkingDay struct {
	DayOfWeek      string `json:"DayOfWeek,omitempty"`
	StartTime      int    `json:"StartTime,omitempty"`
	EndTime        int    `json:"EndTime,omitempty"`
	BreakStartTime int    `json:"BreakStartTime,omitempty"`
	BreakEndTime   int    `json:"BreakEndTime,omitempty"`
	Capacity       int    `json:"Capacity,omitempty"`
}

type ResourceUnit struct {
	Name               string `json:"Name,omitempty"`
	NonOperationalFrom string `json:"NonOperationalFrom,omitempty"`
	NonOperationalTo   string `json:"NonOperationalTo,omitempty"`
	OperationalFrom    string `json:"OperationalFrom,omitempty"`
	OperationalTo      string `json:"OperationalTo,omitempty"`
}

type ResourceCost struct {
	ResourceCostID              string  `json:"ResourceCostID,omitempty"`
	ProductID                   string  `json:"ProductID,omitempty"`
	AccountCode                 string  `json:"AccountCode,omitempty"`
	Cost                        float64 `json:"Cost,omitempty"`
	PriceTier                   int     `json:"PriceTier,omitempty"`
	PlannedDowntimeCost         float64 `json:"PlannedDowntimeCost,omitempty"`
	PlannedDowntimePriceTier    int     `json:"PlannedDowntimePriceTier,omitempty"`
	NotPlannedDowntimeCost      float64 `json:"NotPlannedDowntimeCost,omitempty"`
	NotPlannedDowntimePriceTier int     `json:"NotPlannedDowntimePriceTier,omitempty"`
}

type ResourceRemark struct {
	RemarkID    string `json:"RemarkID,omitempty"`
	Remark      string `json:"Remark,omitempty"`
	CreatedDate string `json:"CreatedDate,omitempty"`
}

type ResourceAttachment struct {
	AttachmentID string `json:"AttachmentID,omitempty"`
	FileName     string `json:"FileName,omitempty"`
	Date         string `json:"Date,omitempty"`
	ContentType  string `json:"ContentType,omitempty"`
	Content      string `json:"Content,omitempty"`
}

type SuspendReason struct {
	SuspendReasonID string   `json:"SuspendReasonID,omitempty"`
	Reason          string   `json:"Reason,omitempty"`
	Workcenters     []string `json:"Workcenters,omitempty"`
}

type WorkCenter struct {
	WorkCenterID         string               `json:"WorkCenterID,omitempty"`
	Code                 string               `json:"Code,omitempty"`
	Name                 string               `json:"Name,omitempty"`
	IsActive             bool                 `json:"IsActive,omitempty"`
	IsCoMan              bool                 `json:"IsCoMan,omitempty"`
	IsCoManPurchase      bool                 `json:"IsCoManPurchase,omitempty"`
	SupplierID           string               `json:"SupplierID,omitempty"`
	SupplierName         string               `json:"SupplierName,omitempty"`
	CoManProcurementType string               `json:"CoManProcurementType,omitempty"`
	WorkCenterLocations  []WorkCenterLocation `json:"WorkCenterLocations,omitempty"`
	WorkCenterSuppliers  []WorkCenterSupplier `json:"WorkCenterSuppliers,omitempty"`
}

type WorkCenterLocation struct {
	ID               string `json:"ID,omitempty"`
	LocationID       string `json:"LocationID,omitempty"`
	Type             string `json:"Type,omitempty"`
	ParentLocationID string `json:"ParentLocationID,omitempty"`
	LocationName     string `json:"LocationName,omitempty"`
}

type WorkCenterSupplier struct {
	SupplierID   string `json:"SupplierID,omitempty"`
	SupplierName string `json:"SupplierName,omitempty"`
}

type PurchaseList struct {
	ID                      string  `json:"ID,omitempty"`
	BlindReceipt            bool    `json:"BlindReceipt,omitempty"`
	OrderNumber             string  `json:"OrderNumber,omitempty"`
	Status                  string  `json:"Status,omitempty"`
	OrderDate               string  `json:"OrderDate,omitempty"`
	InvoiceDate             string  `json:"InvoiceDate,omitempty"`
	Supplier                string  `json:"Supplier,omitempty"`
	SupplierID              string  `json:"SupplierID,omitempty"`
	InvoiceNumber           string  `json:"InvoiceNumber,omitempty"`
	InvoiceAmount           float64 `json:"InvoiceAmount,omitempty"`
	PaidAmount              float64 `json:"PaidAmount,omitempty"`
	InvoiceDueDate          string  `json:"InvoiceDueDate,omitempty"`
	RequiredBy              string  `json:"RequiredBy,omitempty"`
	BaseCurrency            string  `json:"BaseCurrency,omitempty"`
	SupplierCurrency        string  `json:"SupplierCurrency,omitempty"`
	CreditNoteNumber        string  `json:"CreditNoteNumber,omitempty"`
	OrderStatus             string  `json:"OrderStatus,omitempty"`
	StockReceivedStatus     string  `json:"StockReceivedStatus,omitempty"`
	UnstockStatus           string  `json:"UnstockStatus,omitempty"`
	InvoiceStatus           string  `json:"InvoiceStatus,omitempty"`
	CreditNoteStatus        string  `json:"CreditNoteStatus,omitempty"`
	LastUpdatedDate         string  `json:"LastUpdatedDate,omitempty"`
	CombinedReceivingStatus string  `json:"CombinedReceivingStatus,omitempty"`
	CombinedInvoiceStatus   string  `json:"CombinedInvoiceStatus,omitempty"`
	CombinedPaymentStatus   string  `json:"CombinedPaymentStatus,omitempty"`
	Type                    string  `json:"Type,omitempty"`
	IsServiceOnly           bool    `json:"IsServiceOnly,omitempty"`
	DropShipTaskID          string  `json:"DropShipTaskID,omitempty"`
}

type PurchaseCreditNoteList struct {
	ID                      string  `json:"ID,omitempty"`
	BlindReceipt            bool    `json:"BlindReceipt,omitempty"`
	OrderNumber             string  `json:"OrderNumber,omitempty"`
	Status                  string  `json:"Status,omitempty"`
	OrderDate               string  `json:"OrderDate,omitempty"`
	InvoiceDate             string  `json:"InvoiceDate,omitempty"`
	Supplier                string  `json:"Supplier,omitempty"`
	SupplierID              string  `json:"SupplierID,omitempty"`
	InvoiceNumber           string  `json:"InvoiceNumber,omitempty"`
	InvoiceAmount           float64 `json:"InvoiceAmount,omitempty"`
	PaidAmount              float64 `json:"PaidAmount,omitempty"`
	InvoiceDueDate          string  `json:"InvoiceDueDate,omitempty"`
	RequiredBy              string  `json:"RequiredBy,omitempty"`
	BaseCurrency            string  `json:"BaseCurrency,omitempty"`
	SupplierCurrency        string  `json:"SupplierCurrency,omitempty"`
	CreditNoteNumber        string  `json:"CreditNoteNumber,omitempty"`
	OrderStatus             string  `json:"OrderStatus,omitempty"`
	StockReceivedStatus     string  `json:"StockReceivedStatus,omitempty"`
	UnstockStatus           string  `json:"UnstockStatus,omitempty"`
	InvoiceStatus           string  `json:"InvoiceStatus,omitempty"`
	CreditNoteStatus        string  `json:"CreditNoteStatus,omitempty"`
	LastUpdatedDate         string  `json:"LastUpdatedDate,omitempty"`
	CombinedReceivingStatus string  `json:"CombinedReceivingStatus,omitempty"`
	CombinedInvoiceStatus   string  `json:"CombinedInvoiceStatus,omitempty"`
	CombinedPaymentStatus   string  `json:"CombinedPaymentStatus,omitempty"`
	Type                    string  `json:"Type,omitempty"`
	IsServiceOnly           bool    `json:"IsServiceOnly,omitempty"`
	DropShipTaskID          string  `json:"DropShipTaskID,omitempty"`
}

type PurchaseOrder struct {
	TaskID                   string                          `json:"TaskID,omitempty"`
	CombineAdditionalCharges bool                            `json:"CombineAdditionalCharges,omitempty"`
	Memo                     string                          `json:"Memo,omitempty"`
	Status                   string                          `json:"Status,omitempty"`
	Lines                    []PurchaseOrderLineModel        `json:"Lines,omitempty"`
	AdditionalCharges        []PurchaseAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	TotalBeforeTax           float64                         `json:"TotalBeforeTax,omitempty"`
	Tax                      float64                         `json:"Tax,omitempty"`
	Total                    float64                         `json:"Total,omitempty"`
}

type PurchaseCreditNote struct {
	TaskID                   string                                 `json:"TaskID,omitempty"`
	CombineAdditionalCharges bool                                   `json:"CombineAdditionalCharges,omitempty"`
	CreditNoteNumber         string                                 `json:"CreditNoteNumber,omitempty"`
	CreditNoteDate           string                                 `json:"CreditNoteDate,omitempty"`
	Status                   string                                 `json:"Status,omitempty"`
	Lines                    []PurchaseInvoiceLineModel             `json:"Lines,omitempty"`
	AdditionalCharges        []PurchaseInvoiceAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	Unstock                  []PurchaseUnStockLineModel             `json:"Unstock,omitempty"`
	TotalBeforeTax           float64                                `json:"TotalBeforeTax,omitempty"`
	Tax                      float64                                `json:"Tax,omitempty"`
	Total                    float64                                `json:"Total,omitempty"`
}

type Purchase struct {
	ID                      string                               `json:"ID,omitempty"`
	SupplierID              string                               `json:"SupplierID,omitempty"`
	Supplier                string                               `json:"Supplier,omitempty"`
	Contact                 string                               `json:"Contact,omitempty"`
	Phone                   string                               `json:"Phone,omitempty"`
	InventoryAccount        string                               `json:"InventoryAccount,omitempty"`
	BlindReceipt            bool                                 `json:"BlindReceipt,omitempty"`
	Approach                string                               `json:"Approach,omitempty"`
	BillingAddress          []AddressModel                       `json:"BillingAddress,omitempty"`
	ShippingAddress         []PurchaseShippingAddressModel       `json:"ShippingAddress,omitempty"`
	BaseCurrency            string                               `json:"BaseCurrency,omitempty"`
	SupplierCurrency        string                               `json:"SupplierCurrency,omitempty"`
	TaxRule                 string                               `json:"TaxRule,omitempty"`
	TaxCalculation          string                               `json:"TaxCalculation,omitempty"`
	Terms                   string                               `json:"Terms,omitempty"`
	RequiredBy              string                               `json:"RequiredBy,omitempty"`
	Location                string                               `json:"Location,omitempty"`
	Note                    string                               `json:"Note,omitempty"`
	OrderNumber             string                               `json:"OrderNumber,omitempty"`
	OrderDate               string                               `json:"OrderDate,omitempty"`
	CombinedReceivingStatus string                               `json:"CombinedReceivingStatus,omitempty"`
	CombinedInvoiceStatus   string                               `json:"CombinedInvoiceStatus,omitempty"`
	CombinedPaymentStatus   string                               `json:"CombinedPaymentStatus,omitempty"`
	Type                    string                               `json:"Type,omitempty"`
	IsServiceOnly           bool                                 `json:"IsServiceOnly,omitempty"`
	Status                  string                               `json:"Status,omitempty"`
	RelatedDropShipSaleTask string                               `json:"RelatedDropShipSaleTask,omitempty"`
	CurrencyRate            float64                              `json:"CurrencyRate,omitempty"`
	LastUpdatedDate         string                               `json:"LastUpdatedDate,omitempty"`
	Order                   []PurchaseOrderModel                 `json:"Order,omitempty"`
	StockReceived           []AdvancedPurchaseStockModel         `json:"StockReceived,omitempty"`
	PutAway                 []AdvancedPurchasePutAwayModel       `json:"PutAway,omitempty"`
	Invoice                 []AdvancedPurchaseInvoiceModel       `json:"Invoice,omitempty"`
	CreditNote              []AdvancedPurchaseCreditNoteModel    `json:"CreditNote,omitempty"`
	ManualJournals          []AdvancedPurchaseManualJournalModel `json:"ManualJournals,omitempty"`
	AdditionalAttributes    []AdditionalAttributeModel           `json:"AdditionalAttributes,omitempty"`
	Attachments             []AttachmentLineModel                `json:"Attachments,omitempty"`
	InventoryMovements      []InventoryMovementLineModel         `json:"InventoryMovements,omitempty"`
}

type PurchaseStockReceived struct {
	PurchaseID string                           `json:"PurchaseID,omitempty"`
	TaskID     string                           `json:"TaskID,omitempty"`
	Status     string                           `json:"Status,omitempty"`
	Lines      []AdvancedPurchaseStockLineModel `json:"Lines,omitempty"`
}

type PurchasePutAway struct {
	PurchaseID string                             `json:"PurchaseID,omitempty"`
	TaskID     string                             `json:"TaskID,omitempty"`
	Status     string                             `json:"Status,omitempty"`
	Lines      []AdvancedPurchasePutAwayLineModel `json:"Lines,omitempty"`
}

type PurchasePayments struct {
	TaskID       string  `json:"TaskID,omitempty"`
	ID           string  `json:"ID,omitempty"`
	Type         string  `json:"Type,omitempty"`
	Reference    string  `json:"Reference,omitempty"`
	Amount       float64 `json:"Amount,omitempty"`
	DatePaid     string  `json:"DatePaid,omitempty"`
	Account      string  `json:"Account,omitempty"`
	CurrencyRate float64 `json:"CurrencyRate,omitempty"`
	DateCreated  string  `json:"DateCreated,omitempty"`
}

// TODO: Fields are not well definted in API documentation
// type ProductDeal struct {
// 	ID                    string                      `json:"ID,omitempty"`
// 	Name                  string                      `json:"Name,omitempty"`
// 	DateFrom              string                      `json:"DateFrom,omitempty"`
// 	DateTo                string                      `json:"DateTo,omitempty"`
// 	IsActive              bool                        `json:"IsActive,omitempty"`
// 	AllowCoupons          bool                        `json:"AllowCoupons,omitempty"`
// 	SingleCouponCodeUsage bool                        `json:"SingleCouponCodeUsage,omitempty"`
// 	CustomersGroup        string                      `json:"CustomersGroup,omitempty"`
// 	CouponCodes           string                      `json:"CouponCodes,omitempty"`
// 	DealCustomers         []ProductDealCustomerModel  `json:"DealCustomers,omitempty"`
// 	DealCustomerTags      []ProductDealTagModel       `json:"DealCustomerTags,omitempty"`
// 	DealDiscounts         []ProductDealDiscountsModel `json:"DealDiscounts,omitempty"`
// }

type ProductDealCustomerModel struct {
	ID           string `json:"ID,omitempty"`
	CustomerName string `json:"CustomerName,omitempty"`
	CustomerID   string `json:"CustomerID,omitempty"`
}

type ProductDealTagModel struct {
	ID      string `json:"ID,omitempty"`
	TagName string `json:"TagName,omitempty"`
}

type ProductDealDiscountModel struct {
	ID                     string                             `json:"ID,omitempty"`
	DiscountID             string                             `json:"DiscountID,omitempty"`
	DiscountName           string                             `json:"DiscountName,omitempty"`
	DiscountType           string                             `json:"DiscountType,omitempty"`
	Sequence               int                                `json:"Sequence,omitempty"`
	IsOrderLevel           bool                               `json:"IsOrderLevel,omitempty"`
	BuyType                string                             `json:"BuyType,omitempty"`
	BuyValue               float64                            `json:"BuyValue,omitempty"`
	BuyValueType           string                             `json:"BuyValueType,omitempty"`
	BuyMore                bool                               `json:"BuyMore,omitempty"`
	GetType                string                             `json:"GetType,omitempty"`
	GetValue               float64                            `json:"GetValue,omitempty"`
	GetValueType           string                             `json:"GetValueType,omitempty"`
	DealDiscountBrands     []ProductDealDiscountBrandModel    `json:"DealDiscountBrands,omitempty"`
	DealDiscountCategories []ProductDealDiscountCategoryModel `json:"DealDiscountCategories,omitempty"`
	DealDiscountTags       []ProductDealDiscountTagModel      `json:"DealDiscountTags,omitempty"`
	DealDiscountProducts   []ProductDealDiscountProductModel  `json:"DealDiscountProducts,omitempty"`
}

type ProductDealDiscountBrandModel struct {
	ID        string `json:"ID,omitempty"`
	BrandName string `json:"BrandName,omitempty"`
	BrandID   string `json:"BrandID,omitempty"`
	Type      string `json:"Type,omitempty"`
}

type ProductDealDiscountProductModel struct {
	ID          string `json:"ID,omitempty"`
	ProductSKU  string `json:"ProductSKU,omitempty"`
	ProductID   string `json:"ProductID,omitempty"`
	ProductName string `json:"ProductName,omitempty"`
	IsFamily    bool   `json:"IsFamily,omitempty"`
	Type        string `json:"Type,omitempty"`
}

type ProductDealDiscountTagModel struct {
	ID      string `json:"ID,omitempty"`
	TagName string `json:"TagName,omitempty"`
	Type    string `json:"Type,omitempty"`
}

type ProductDealDiscountCategoryModel struct {
	ID           string `json:"ID,omitempty"`
	CategoryName string `json:"CategoryName,omitempty"`
	Type         string `json:"Type,omitempty"`
}

type ProductDiscountRuleModel struct {
	ID            string              `json:"ID,omitempty"`
	Name          string              `json:"Name,omitempty"`
	IsActive      bool                `json:"IsActive,omitempty"`
	Type          string              `json:"Type,omitempty"`
	DiscountLines []DiscountLineModel `json:"DiscountLines,omitempty"`
}

type DiscountLineModel struct {
	ID           string  `json:"ID,omitempty"`
	MinValue     float64 `json:"MinValue,omitempty"`
	MaxValue     float64 `json:"MaxValue,omitempty"`
	DiscountType string  `json:"DiscountType,omitempty"`
	Amount       float64 `json:"Amount,omitempty"`
	OrderExceeds float64 `json:"OrderExceeds,omitempty"`
}

// TODO: Fields are not well defined in API Docs
// ISSUE: ShipXoneConditionsModel
// type ShippingZoneModel struct {
// 	ZoneID              string                    `json:"ZoneID,omitempty"`
// 	Name                string                    `json:"Name,omitempty"`
// 	IsRestZone          bool                      `json:"IsRestZone,omitempty"`
// 	PricesInclTax       bool                      `json:"PricesInclTax,omitempty"`
// 	Negative            bool                      `json:"Negative,omitempty"`
// 	DefaultShippingCost float64                   `json:"DefaultShippingCost,omitempty"`
// 	ShortDesc           bool                      `json:"ShortDesc,omitempty"`
// 	TaxRuleID           string                    `json:"TaxRuleID,omitempty"`
// 	TaxRuleName         string                    `json:"TaxRuleName,omitempty"`
// 	AppliesTo           []ShipZoneAppliesToModel  `json:"AppliesTo,omitempty"`
// 	Conditions          []ShipZoneConditionsModel `json:"Conditions,omitempty"`
// }

type ShipZoneAppliesToModel struct {
	ID            string  `json:"ID,omitempty"`
	Country2      string  `json:"Country2,omitempty"`
	StateID       int     `json:"StateID,omitempty"`
	StateName     string  `json:"StateName,omitempty"`
	PostCodeFrom  string  `json:"PostCodeFrom,omitempty"`
	PostCodeTo    string  `json:"PostCodeTo,omitempty"`
	PostCodesList string  `json:"PostCodesList,omitempty"`
	ShippingRate  float64 `json:"ShippingRate,omitempty"`
}

type ShipZoneConditionModel struct {
	ZoneConditionID      string  `json:"ZoneConditionID,omitempty"`
	MinValue             float64 `json:"MinValue,omitempty"`
	MaxValue             float64 `json:"MaxValue,omitempty"`
	ShippingCost         float64 `json:"ShippingCost,omitempty"`
	ConditionType        string  `json:"ConditionType,omitempty"`
	ConditionDescription string  `json:"ConditionDescription,omitempty"`
}

type SaleList struct {
	SaleID                  string  `json:"SaleID,omitempty"`
	OrderNumber             string  `json:"OrderNumber,omitempty"`
	Status                  string  `json:"Status,omitempty"`
	OrderDate               string  `json:"OrderDate,omitempty"`
	InvoiceDate             string  `json:"InvoiceDate,omitempty"`
	Customer                string  `json:"Customer,omitempty"`
	CustomerID              string  `json:"CustomerID,omitempty"`
	InvoiceNumber           string  `json:"InvoiceNumber,omitempty"`
	CustomerReference       string  `json:"CustomerReference,omitempty"`
	InvoiceAmount           float64 `json:"InvoiceAmount,omitempty"`
	PaidAmount              float64 `json:"PaidAmount,omitempty"`
	InvoiceDueDate          string  `json:"InvoiceDueDate,omitempty"`
	ShipBy                  string  `json:"ShipBy,omitempty"`
	BaseCurrency            string  `json:"BaseCurrency,omitempty"`
	CustomerCurrency        string  `json:"CustomerCurrency,omitempty"`
	CreditNoteNumber        string  `json:"CreditNoteNumber,omitempty"`
	Updated                 string  `json:"Updated,omitempty"`
	QuoteStatus             string  `json:"QuoteStatus,omitempty"`
	OrderStatus             string  `json:"OrderStatus,omitempty"`
	CombinedPickingStatus   string  `json:"CombinedPickingStatus,omitempty"`
	CombinedPackingStatus   string  `json:"CombinedPackingStatus,omitempty"`
	CombinedShippingStatus  string  `json:"CombinedShippingStatus,omitempty"`
	FulFilmentStatus        string  `json:"FulFilmentStatus,omitempty"`
	CombinedInvoiceStatus   string  `json:"CombinedInvoiceStatus,omitempty"`
	CreditNoteStatus        string  `json:"CreditNoteStatus,omitempty"`
	CombinedPaymentStatus   string  `json:"CombinedPaymentStatus,omitempty"`
	Type                    string  `json:"Type,omitempty"`
	CombinedTrackingNumbers string  `json:"CombinedTrackingNumbers,omitempty"`
	SourceChannel           string  `json:"SourceChannel,omitempty"`
	ExternalID              string  `json:"ExternalID,omitempty"`
	OrderLocationID         string  `json:"OrderLocationID,omitempty"`
}

type SaleCreditNoteList struct {
	SaleID                  string  `json:"SaleID,omitempty"`
	OrderNumber             string  `json:"OrderNumber,omitempty"`
	Status                  string  `json:"Status,omitempty"`
	OrderDate               string  `json:"OrderDate,omitempty"`
	InvoiceDate             string  `json:"InvoiceDate,omitempty"`
	Customer                string  `json:"Customer,omitempty"`
	CustomerID              string  `json:"CustomerID,omitempty"`
	InvoiceNumber           string  `json:"InvoiceNumber,omitempty"`
	CustomerReference       string  `json:"CustomerReference,omitempty"`
	InvoiceAmount           float64 `json:"InvoiceAmount,omitempty"`
	PaidAmount              float64 `json:"PaidAmount,omitempty"`
	InvoiceDueDate          string  `json:"InvoiceDueDate,omitempty"`
	ShipBy                  string  `json:"ShipBy,omitempty"`
	BaseCurrency            string  `json:"BaseCurrency,omitempty"`
	CustomerCurrency        string  `json:"CustomerCurrency,omitempty"`
	CreditNoteNumber        string  `json:"CreditNoteNumber,omitempty"`
	Updated                 string  `json:"Updated,omitempty"`
	QuoteStatus             string  `json:"QuoteStatus,omitempty"`
	OrderStatus             string  `json:"OrderStatus,omitempty"`
	CombinedPickingStatus   string  `json:"CombinedPickingStatus,omitempty"`
	CombinedPackingStatus   string  `json:"CombinedPackingStatus,omitempty"`
	CombinedShippingStatus  string  `json:"CombinedShippingStatus,omitempty"`
	FulFilmentStatus        string  `json:"FulFilmentStatus,omitempty"`
	CombinedInvoiceStatus   string  `json:"CombinedInvoiceStatus,omitempty"`
	CreditNoteStatus        string  `json:"CreditNoteStatus,omitempty"`
	CombinedPaymentStatus   string  `json:"CombinedPaymentStatus,omitempty"`
	Type                    string  `json:"Type,omitempty"`
	CombinedTrackingNumbers string  `json:"CombinedTrackingNumbers,omitempty"`
	SourceChannel           string  `json:"SourceChannel,omitempty"`
	ExternalID              string  `json:"ExternalID,omitempty"`
	OrderLocationID         string  `json:"OrderLocationID,omitempty"`
}

type Sale struct {
	ID                      string                       `json:"ID,omitempty"`
	Customer                string                       `json:"Customer,omitempty"`
	CustomerID              string                       `json:"CustomerID,omitempty"`
	Contact                 string                       `json:"Contact,omitempty"`
	Phone                   string                       `json:"Phone,omitempty"`
	Email                   string                       `json:"Email,omitempty"`
	DefaultAccount          string                       `json:"DefaultAccount,omitempty"`
	SkipQuote               bool                         `json:"SkipQuote,omitempty"`
	BillingAddress          []AddressModel               `json:"BillingAddress,omitempty"`
	ShippingAddress         []SaleShippingAddressModel   `json:"ShippingAddress,omitempty"`
	ShippingNotes           string                       `json:"ShippingNotes,omitempty"`
	BaseCurrency            string                       `json:"BaseCurrency,omitempty"`
	CustomerCurrency        string                       `json:"CustomerCurrency,omitempty"`
	TaxRule                 string                       `json:"TaxRule,omitempty"`
	TaxCalculation          string                       `json:"TaxCalculation,omitempty"`
	Terms                   string                       `json:"Terms,omitempty"`
	PriceTier               string                       `json:"PriceTier,omitempty"`
	ShipBy                  string                       `json:"ShipBy,omitempty"`
	Location                string                       `json:"Location,omitempty"`
	SaleOrderDate           string                       `json:"SaleOrderDate,omitempty"`
	LastModifiedOn          string                       `json:"LastModifiedOn,omitempty"`
	Note                    string                       `json:"Note,omitempty"`
	CustomerReference       string                       `json:"CustomerReference,omitempty"`
	COGSAmount              float64                      `json:"COGSAmount,omitempty"`
	Status                  string                       `json:"Status,omitempty"`
	CombinedPickingStatus   string                       `json:"CombinedPickingStatus,omitempty"`
	CombinedPackingStatus   string                       `json:"CombinedPackingStatus,omitempty"`
	CombinedShippingStatus  string                       `json:"CombinedShippingStatus,omitempty"`
	FulFilmentStatus        string                       `json:"FulFilmentStatus,omitempty"`
	CombinedInvoiceStatus   string                       `json:"CombinedInvoiceStatus,omitempty"`
	CombinedPaymentStatus   string                       `json:"CombinedPaymentStatus,omitempty"`
	CombinedTrackingNumbers string                       `json:"CombinedTrackingNumbers,omitempty"`
	Carrier                 string                       `json:"Carrier,omitempty"`
	CurrencyRate            float64                      `json:"CurrencyRate,omitempty"`
	SalesRepresentative     string                       `json:"SalesRepresentative,omitempty"`
	Type                    string                       `json:"Type,omitempty"`
	SourceChannel           string                       `json:"SourceChannel,omitempty"`
	ExternalID              string                       `json:"ExternalID,omitempty"`
	ServiceOnly             bool                         `json:"ServiceOnly,omitempty"`
	Quote                   []SaleQuoteModel             `json:"Quote,omitempty"`
	Order                   []SaleOrderModel             `json:"Order,omitempty"`
	Fulfilments             []SaleFulfilmentModel        `json:"Fulfilments,omitempty"`
	Invoices                []SaleInvoiceModel           `json:"Invoices,omitempty"`
	CreditNotes             []SaleCreditNoteModel        `json:"CreditNotes,omitempty"`
	ManualJournals          []SaleManualJournalModel     `json:"ManualJournals,omitempty"`
	AdditionalAttributes    []AdditionalAttributeModel   `json:"AdditionalAttributes,omitempty"`
	Attachments             []AttachmentLineModel        `json:"Attachments,omitempty"`
	InventoryMovements      []InventoryMovementLineModel `json:"InventoryMovements,omitempty"`
	Transactions            []SaleTransactionLineModel   `json:"Transactions,omitempty"`
}

type SaleQuote struct {
	SaleID                   string                      `json:"SaleID,omitempty"`
	CombineAdditionalCharges bool                        `json:"CombineAdditionalCharges,omitempty"`
	Memo                     string                      `json:"Memo,omitempty"`
	Status                   string                      `json:"Status,omitempty"`
	Lines                    []SaleQuoteLineModel        `json:"Lines,omitempty"`
	AdditionalCharges        []SaleAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	TotalBeforeTax           float64                     `json:"TotalBeforeTax,omitempty"`
	Tax                      float64                     `json:"Tax,omitempty"`
	Total                    float64                     `json:"Total,omitempty"`
}

type SaleOrder struct {
	SaleID                   string                      `json:"SaleID,omitempty"`
	SaleOrderNumber          string                      `json:"SaleOrderNumber,omitempty"`
	CombineAdditionalCharges bool                        `json:"CombineAdditionalCharges,omitempty"`
	Memo                     string                      `json:"Memo,omitempty"`
	Status                   string                      `json:"Status,omitempty"`
	Lines                    []SaleOrderLineModel        `json:"Lines,omitempty"`
	AdditionalCharges        []SaleAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	TotalBeforeTax           float64                     `json:"TotalBeforeTax,omitempty"`
	Tax                      float64                     `json:"Tax,omitempty"`
	Total                    float64                     `json:"Total,omitempty"`
}

type SaleFulfilment struct {
	SaleID      string                `json:"SaleID,omitempty"`
	Fulfilments []SaleFulfilmentModel `json:"Fulfilments,omitempty"`
}

type SaleFulfilmentPick struct {
	TaskID string                            `json:"TaskID,omitempty"`
	Status string                            `json:"Status,omitempty"`
	Lines  []SaleFulfilmentPickPackLineModel `json:"Lines,omitempty"`
}

type SaleFulfilmentPack struct {
	TaskID string                            `json:"TaskID,omitempty"`
	Status string                            `json:"Status,omitempty"`
	Lines  []SaleFulfilmentPickPackLineModel `json:"Lines,omitempty"`
}

type SaleFulfilmentShip struct {
	TaskID          string                        `json:"TaskID,omitempty"`
	Status          string                        `json:"Status,omitempty"`
	RequireBy       string                        `json:"RequireBy,omitempty"`
	ShippingAddress []SaleShippingAddressModel    `json:"ShippingAddress,omitempty"`
	ShippingNotes   string                        `json:"ShippingNotes,omitempty"`
	Lines           []SaleFulfilmentShipLineModel `json:"Lines,omitempty"`
}

type StockAdjustment struct {
	TaskID             string                      `json:"TaskID,omitempty"`
	EffectiveDate      string                      `json:"EffectiveDate,omitempty"`
	StocktakeNumber    string                      `json:"StocktakeNumber,omitempty"`
	Status             string                      `json:"Status,omitempty"`
	Account            string                      `json:"Account,omitempty"`
	Reference          string                      `json:"Reference,omitempty"`
	ExistingStockLines []ExistingStockLineModel    `json:"ExistingStockLines,omitempty"`
	NewStockLines      []NewStockLineModel         `json:"NewStockLines,omitempty"`
	Transactions       []TransactionStockLineModel `json:"Transactions,omitempty"`
}

type StockTake struct {
	TaskID                     string                      `json:"TaskID,omitempty"`
	EffectiveDate              string                      `json:"EffectiveDate,omitempty"`
	StocktakeNumber            string                      `json:"StocktakeNumber,omitempty"`
	Status                     string                      `json:"Status,omitempty"`
	Account                    string                      `json:"Account,omitempty"`
	LocationID                 string                      `json:"LocationID,omitempty"`
	Location                   string                      `json:"Location,omitempty"`
	Tags                       []string                    `json:"Tags,omitempty"`
	PickZones                  []string                    `json:"PickZones,omitempty"`
	StockLocators              []string                    `json:"StockLocators,omitempty"`
	Categories                 []IDNameModel               `json:"Categories,omitempty"`
	Brands                     []IDNameModel               `json:"Brands,omitempty"`
	Bins                       []IDNameModel               `json:"Bins,omitempty"`
	Reference                  string                      `json:"Reference,omitempty"`
	NonZeroStockOnHandProducts []ExistingStockLineModel    `json:"NonZeroStockOnHandProducts,omitempty"`
	ZeroStockOnHandProducts    []NewStockLineModel         `json:"ZeroStockOnHandProducts,omitempty"`
	Transactions               []TransactionStockLineModel `json:"Transactions,omitempty"`
}

type StockTransfer struct {
	TaskID               string                    `json:"TaskID,omitempty"`
	From                 string                    `json:"From,omitempty"`
	FromLocation         string                    `json:"FromLocation,omitempty"`
	To                   string                    `json:"To,omitempty"`
	ToLocation           string                    `json:"ToLocation,omitempty"`
	Status               string                    `json:"Status,omitempty"`
	Number               string                    `json:"Number,omitempty"`
	CostDistributionType string                    `json:"CostDistributionType,omitempty"`
	InTransitAccount     string                    `json:"InTransitAccount,omitempty"`
	DepartureDate        string                    `json:"DepartureDate,omitempty"`
	CompletionDate       string                    `json:"CompletionDate,omitempty"`
	RequiredByDate       string                    `json:"RequiredByDate,omitempty"`
	Reference            string                    `json:"Reference,omitempty"`
	SkipOrder            bool                      `json:"SkipOrder,omitempty"`
	Lines                []StockTransferLineModel  `json:"Lines,omitempty"`
	Order                []StockTransferOrderModel `json:"Order,omitempty"`
}

type StockTransferOrder struct {
	TaskID string                        `json:"TaskID,omitempty"`
	Status string                        `json:"Status,omitempty"`
	Lines  []StockTransferOrderLineModel `json:"Lines,omitempty"`
}

type Supplier struct {
	ID                    string                 `json:"ID,omitempty"`
	Name                  string                 `json:"Name,omitempty"`
	Currency              string                 `json:"Currency,omitempty"`
	PaymentTerm           string                 `json:"PaymentTerm,omitempty"`
	AccountPayable        string                 `json:"AccountPayable,omitempty"`
	TaxRule               string                 `json:"TaxRule,omitempty"`
	Status                string                 `json:"Status,omitempty"`
	Discount              int                    `json:"Discount,omitempty"`
	Comments              string                 `json:"Comments,omitempty"`
	TaxNumber             int                    `json:"TaxNumber,omitempty"`
	AttributeSet          string                 `json:"AttributeSet,omitempty"`
	AdditionalAttribute1  string                 `json:"AdditionalAttribute1,omitempty"`
	AdditionalAttribute2  string                 `json:"AdditionalAttribute2,omitempty"`
	AdditionalAttribute3  string                 `json:"AdditionalAttribute3,omitempty"`
	AdditionalAttribute4  string                 `json:"AdditionalAttribute4,omitempty"`
	AdditionalAttribute5  string                 `json:"AdditionalAttribute5,omitempty"`
	AdditionalAttribute6  string                 `json:"AdditionalAttribute6,omitempty"`
	AdditionalAttribute7  string                 `json:"AdditionalAttribute7,omitempty"`
	AdditionalAttribute8  string                 `json:"AdditionalAttribute8,omitempty"`
	AdditionalAttribute9  string                 `json:"AdditionalAttribute9,omitempty"`
	AdditionalAttribute10 string                 `json:"AdditionalAttribute10,omitempty"`
	LastModifiedOn        string                 `json:"LastModifiedOn,omitempty"`
	Addresses             []SupplierAddressModel `json:"Addresses,omitempty"`
	Contacts              []SupplierContactModel `json:"Contacts,omitempty"`
}

type Tax struct {
	ID               string              `json:"ID,omitempty"`
	Name             string              `json:"Name,omitempty"`
	Account          string              `json:"Account,omitempty"`
	IsActive         bool                `json:"IsActive,omitempty"`
	TaxInclusive     bool                `json:"TaxInclusive,omitempty"`
	TaxPercent       float64             `json:"TaxPercent,omitempty"`
	IsTaxForSale     bool                `json:"IsTaxForSale,omitempty"`
	IsTaxForPurchase bool                `json:"IsTaxForPurchase,omitempty"`
	Components       []TaxComponentModel `json:"Components,omitempty"`
}

type UnitofMeasure struct {
	ID   string `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
}

type Lead struct {
	ID                  string                 `json:"ID,omitempty"`
	LeadStatus          string                 `json:"LeadStatus,omitempty"`
	Name                string                 `json:"Name,omitempty"`
	Currency            string                 `json:"Currency,omitempty"`
	PaymentTerm         int                    `json:"PaymentTerm,omitempty"`
	PriceTier           string                 `json:"PriceTier,omitempty"`
	SalesRepresentative string                 `json:"SalesRepresentative,omitempty"`
	TaxRule             string                 `json:"TaxRule,omitempty"`
	Amount              float64                `json:"Amount,omitempty"`
	CloseChance         int                    `json:"CloseChance,omitempty"`
	CloseDate           string                 `json:"CloseDate,omitempty"`
	Comments            string                 `json:"Comments,omitempty"`
	Contacts            []SupplierContactModel `json:"Contacts,omitempty"`
	Addresses           []SupplierAddressModel `json:"Addresses,omitempty"`
}

type OpportunityAdditionalCharge struct {
	ID          string  `json:"ID,omitempty"`
	Description string  `json:"Description,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	Amount      float64 `json:"Amount,omitempty"`
	Discount    float64 `json:"Discount,omitempty"`
	Tax         float64 `json:"Tax,omitempty"`
	Total       float64 `json:"Total,omitempty"`
	TaxRule     string  `json:"TaxRule,omitempty"`
	TaxPercent  float64 `json:"TaxPercent,omitempty"`
	Comment     string  `json:"Comment,omitempty"`
}

type TaskModel struct {
	ID            string `json:"ID,omitempty"`
	Name          string `json:"Name,omitempty"`
	Description   string `json:"Description,omitempty"`
	StartDate     string `json:"StartDate,omitempty"`
	EndDate       string `json:"EndDate,omitempty"`
	IsAllDay      bool   `json:"IsAllDay,omitempty"`
	CompletedDate string `json:"CompletedDate,omitempty"`
	Category      string `json:"Category,omitempty"`
	WorkflowName  string `json:"WorkflowName,omitempty"`
	AssignedTo    string `json:"AssignedTo,omitempty"`
	EntityType    string `json:"EntityType,omitempty"`
	EntityID      string `json:"EntityID,omitempty"`
	IsImportant   bool   `json:"IsImportant,omitempty"`
	TaskStatus    string `json:"TaskStatus,omitempty"`
	AssignedBy    string `json:"AssignedBy,omitempty"`
	VoidDate      string `json:"VoidDate,omitempty"`
}

type TaskCategoryModel struct {
	ID              string `json:"ID,omitempty"`
	Name            string `json:"Name,omitempty"`
	Color           string `json:"Color,omitempty"`
	BackgroundColor string `json:"BackgroundColor,omitempty"`
}

type Workflow struct {
	ID          string         `json:"ID,omitempty"`
	Name        string         `json:"Name,omitempty"`
	IsActive    bool           `json:"IsActive,omitempty"`
	EntityType  string         `json:"EntityType,omitempty"`
	DueDaysType string         `json:"DueDaysType,omitempty"`
	Steps       []WorkflowStep `json:"Steps,omitempty"`
}

type WorkflowStep struct {
	ID          string `json:"ID,omitempty"`
	Name        string `json:"Name,omitempty"`
	Category    string `json:"Category,omitempty"`
	Days        int    `json:"Days,omitempty"`
	StartTime   string `json:"StartTime,omitempty"`
	EndTime     string `json:"EndTime,omitempty"`
	UserName    string `json:"UserName,omitempty"`
	SkipHoliday string `json:"SkipHoliday,omitempty"`
}

type Opportunity struct {
	ID                     string                        `json:"ID,omitempty"`
	CustomerID             string                        `json:"CustomerID,omitempty"`
	LeadID                 string                        `json:"LeadID,omitempty"`
	CustomerName           string                        `json:"CustomerName,omitempty"`
	Contact                string                        `json:"Contact,omitempty"`
	Phone                  string                        `json:"Phone,omitempty"`
	DefaultAccount         string                        `json:"DefaultAccount,omitempty"`
	BillingAddressLine1    string                        `json:"BillingAddressLine1,omitempty"`
	BillingAddressLine2    string                        `json:"BillingAddressLine2,omitempty"`
	Currency               string                        `json:"Currency,omitempty"`
	TaxPercent             float64                       `json:"TaxPercent,omitempty"`
	TaxInclusive           bool                          `json:"TaxInclusive,omitempty"`
	TaxRule                string                        `json:"TaxRule,omitempty"`
	TermDays               int                           `json:"TermDays,omitempty"`
	Terms                  string                        `json:"Terms,omitempty"`
	PriceTier              string                        `json:"PriceTier,omitempty"`
	ShippingAddressLine1   string                        `json:"ShippingAddressLine1,omitempty"`
	ShippingAddressLine2   string                        `json:"ShippingAddressLine2,omitempty"`
	OpportunityLocation    string                        `json:"OpportunityLocation,omitempty"`
	OpportunityNumber      string                        `json:"OpportunityNumber,omitempty"`
	OpportunityDate        string                        `json:"OpportunityDate,omitempty"`
	OpportunityComment     string                        `json:"OpportunityComment,omitempty"`
	OpportunityMemo        string                        `json:"OpportunityMemo,omitempty"`
	OpportunityStatus      string                        `json:"OpportunityStatus,omitempty"`
	TaxTotal               float64                       `json:"TaxTotal,omitempty"`
	Total                  float64                       `json:"Total,omitempty"`
	CustomerReference      string                        `json:"CustomerReference,omitempty"`
	CurrencyConversionRate float64                       `json:"CurrencyConversionRate,omitempty"`
	CustomerCurrency       string                        `json:"CustomerCurrency,omitempty"`
	TermMethod             int                           `json:"TermMethod,omitempty"`
	SalesRepresentative    string                        `json:"SalesRepresentative,omitempty"`
	TermDueNextMonth       int                           `json:"TermDueNextMonth,omitempty"`
	ShipToOther            bool                          `json:"ShipToOther,omitempty"`
	ShipToCompany          string                        `json:"ShipToCompany,omitempty"`
	ShipToContact          string                        `json:"ShipToContact,omitempty"`
	ShipToCountry          string                        `json:"ShipToCountry,omitempty"`
	ShipToPostCode         string                        `json:"ShipToPostCode,omitempty"`
	ShipToState            string                        `json:"ShipToState,omitempty"`
	ShipToCity             string                        `json:"ShipToCity,omitempty"`
	ShipToAddress1         string                        `json:"ShipToAddress1,omitempty"`
	ShipToAddress2         string                        `json:"ShipToAddress2,omitempty"`
	CustomField1           string                        `json:"CustomField1,omitempty"`
	CustomField2           string                        `json:"CustomField2,omitempty"`
	CustomField3           string                        `json:"CustomField3,omitempty"`
	CustomField4           string                        `json:"CustomField4,omitempty"`
	CustomField5           string                        `json:"CustomField5,omitempty"`
	CustomField6           string                        `json:"CustomField6,omitempty"`
	CustomField7           string                        `json:"CustomField7,omitempty"`
	CustomField8           string                        `json:"CustomField8,omitempty"`
	CustomField9           string                        `json:"CustomField9,omitempty"`
	CustomField10          string                        `json:"CustomField10,omitempty"`
	Lines                  []OpportunityLine             `json:"Lines,omitempty"`
	AdditionalCharges      []OpportunityAdditionalCharge `json:"AdditionalCharges,omitempty"`
}

type OpportunityLine struct {
	ID          string  `json:"ID,omitempty"`
	ProductID   string  `json:"ProductID,omitempty"`
	ProductName string  `json:"ProductName,omitempty"`
	ProductSKU  string  `json:"ProductSKU,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	Price       float64 `json:"Price,omitempty"`
	Discount    float64 `json:"Discount,omitempty"`
	Tax         float64 `json:"Tax,omitempty"`
	Total       float64 `json:"Total,omitempty"`
	TaxRule     string  `json:"TaxRule,omitempty"`
	TaxPercent  float64 `json:"TaxPercent,omitempty"`
	Comment     string  `json:"Comment,omitempty"`
}

type AddressModel struct {
	ID                  string `json:"ID,omitempty"`
	DisplayAddressLine1 string `json:"DisplayAddressLine1,omitempty"`
	DisplayAddressLine2 string `json:"DisplayAddressLine2,omitempty"`
	Line1               string `json:"Line1,omitempty"`
	Line2               string `json:"Line2,omitempty"`
	City                string `json:"City,omitempty"`
	State               string `json:"State,omitempty"`
	Postcode            string `json:"Postcode,omitempty"`
	Country             string `json:"Country,omitempty"`
}

type SaleShippingAddressModel struct {
	ID                  string `json:"ID,omitempty"`
	DisplayAddressLine1 string `json:"DisplayAddressLine1,omitempty"`
	DisplayAddressLine2 string `json:"DisplayAddressLine2,omitempty"`
	Line1               string `json:"Line1,omitempty"`
	Line2               string `json:"Line2,omitempty"`
	City                string `json:"City,omitempty"`
	State               string `json:"State,omitempty"`
	Postcode            string `json:"Postcode,omitempty"`
	Country             string `json:"Country,omitempty"`
	Company             string `json:"Company,omitempty"`
	Contact             string `json:"Contact,omitempty"`
	ShipToOther         bool   `json:"ShipToOther,omitempty"`
}

type PurchaseShippingAddressModel struct {
	DisplayAddressLine1 string `json:"DisplayAddressLine1,omitempty"`
	DisplayAddressLine2 string `json:"DisplayAddressLine2,omitempty"`
	Line1               string `json:"Line1,omitempty"`
	Line2               string `json:"Line2,omitempty"`
	City                string `json:"City,omitempty"`
	State               string `json:"State,omitempty"`
	Postcode            string `json:"Postcode,omitempty"`
	Country             string `json:"Country,omitempty"`
	ShipToOther         bool   `json:"ShipToOther,omitempty"`
	Company             string `json:"Company,omitempty"`
}

type AdditionalAttributeModel struct {
	AdditionalAttribute1  string `json:"AdditionalAttribute1,omitempty"`
	AdditionalAttribute2  string `json:"AdditionalAttribute2,omitempty"`
	AdditionalAttribute3  string `json:"AdditionalAttribute3,omitempty"`
	AdditionalAttribute4  string `json:"AdditionalAttribute4,omitempty"`
	AdditionalAttribute5  string `json:"AdditionalAttribute5,omitempty"`
	AdditionalAttribute6  string `json:"AdditionalAttribute6,omitempty"`
	AdditionalAttribute7  string `json:"AdditionalAttribute7,omitempty"`
	AdditionalAttribute8  string `json:"AdditionalAttribute8,omitempty"`
	AdditionalAttribute9  string `json:"AdditionalAttribute9,omitempty"`
	AdditionalAttribute10 string `json:"AdditionalAttribute10,omitempty"`
}

type SaleQuoteModel struct {
	Memo              string                      `json:"Memo,omitempty"`
	Status            string                      `json:"Status,omitempty"`
	Prepayments       []SalePaymentLineModel      `json:"Prepayments,omitempty"`
	Lines             []SaleQuoteLineModel        `json:"Lines,omitempty"`
	AdditionalCharges []SaleAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	TotalBeforeTax    float64                     `json:"TotalBeforeTax,omitempty"`
	Tax               float64                     `json:"Tax,omitempty"`
	Total             float64                     `json:"Total,omitempty"`
}

type SaleQuoteLineModel struct {
	ProductID   string  `json:"ProductID,omitempty"`
	SKU         string  `json:"SKU,omitempty"`
	Name        string  `json:"Name,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	Price       float64 `json:"Price,omitempty"`
	Discount    float64 `json:"Discount,omitempty"`
	Tax         float64 `json:"Tax,omitempty"`
	AverageCost float64 `json:"AverageCost,omitempty"`
	TaxRule     string  `json:"TaxRule,omitempty"`
	Comment     string  `json:"Comment,omitempty"`
	Total       float64 `json:"Total,omitempty"`
}

type SaleOrderModel struct {
	SaleOrderNumber   string                      `json:"SaleOrderNumber,omitempty"`
	Memo              string                      `json:"Memo,omitempty"`
	Status            string                      `json:"Status,omitempty"`
	Lines             []SaleOrderLineModel        `json:"Lines,omitempty"`
	AdditionalCharges []SaleAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	TotalBeforeTax    float64                     `json:"TotalBeforeTax,omitempty"`
	Tax               float64                     `json:"Tax,omitempty"`
	Total             float64                     `json:"Total,omitempty"`
}

type SaleOrderLineModel struct {
	ProductID         string  `json:"ProductID,omitempty"`
	SKU               string  `json:"SKU,omitempty"`
	Name              string  `json:"Name,omitempty"`
	Quantity          float64 `json:"Quantity,omitempty"`
	Price             float64 `json:"Price,omitempty"`
	Discount          float64 `json:"Discount,omitempty"`
	Tax               float64 `json:"Tax,omitempty"`
	AverageCost       float64 `json:"AverageCost,omitempty"`
	TaxRule           string  `json:"TaxRule,omitempty"`
	Comment           string  `json:"Comment,omitempty"`
	DropShip          bool    `json:"DropShip,omitempty"`
	BackorderQuantity float64 `json:"BackorderQuantity,omitempty"`
	Total             float64 `json:"Total,omitempty"`
}

type SaleFulfilmentModel struct {
	TaskID              string                        `json:"TaskID,omitempty"`
	FulfillmentNumber   int                           `json:"FulfillmentNumber,omitempty"`
	LinkedInvoiceNumber string                        `json:"LinkedInvoiceNumber,omitempty"`
	FulFilmentStatus    string                        `json:"FulFilmentStatus,omitempty"`
	Pick                []SaleFulfilmentPickPackModel `json:"Pick,omitempty"`
	Pack                []SaleFulfilmentPickPackModel `json:"Pack,omitempty"`
	Ship                []SaleFulfilmentShipModel     `json:"Ship,omitempty"`
}

type SaleFulfilmentPickPackModel struct {
	Status string                            `json:"Status,omitempty"`
	Lines  []SaleFulfilmentPickPackLineModel `json:"Lines,omitempty"`
}

type SaleFulfilmentPickPackLineModel struct {
	ProductID                  string  `json:"ProductID,omitempty"`
	SKU                        string  `json:"SKU,omitempty"`
	Name                       string  `json:"Name,omitempty"`
	Location                   string  `json:"Location,omitempty"`
	LocationID                 string  `json:"LocationID,omitempty"`
	Quantity                   float64 `json:"Quantity,omitempty"`
	BatchSN                    string  `json:"BatchSN,omitempty"`
	ExpiryDate                 string  `json:"ExpiryDate,omitempty"`
	Box                        string  `json:"Box,omitempty"`
	NonInventory               bool    `json:"NonInventory,omitempty"`
	WarrantyRegistrationNumber string  `json:"WarrantyRegistrationNumber,omitempty"`
	RestockLocation            string  `json:"RestockLocation,omitempty"`
	RestockLocationID          string  `json:"RestockLocationID,omitempty"`
}

type SaleFulfilmentShipModel struct {
	Status          string                        `json:"Status,omitempty"`
	RequireBy       string                        `json:"RequireBy,omitempty"`
	ShippingAddress []SaleShippingAddressModel    `json:"ShippingAddress,omitempty"`
	ShippingNotes   string                        `json:"ShippingNotes,omitempty"`
	Lines           []SaleFulfilmentShipLineModel `json:"Lines,omitempty"`
}

type SaleFulfilmentShipLineModel struct {
	ID             string `json:"ID,omitempty"`
	ShipmentDate   string `json:"ShipmentDate,omitempty"`
	Carrier        string `json:"Carrier,omitempty"`
	Boxes          string `json:"Boxes,omitempty"`
	TrackingNumber string `json:"TrackingNumber,omitempty"`
	TrackingURL    string `json:"TrackingURL,omitempty"`
	IsShipped      bool   `json:"IsShipped,omitempty"`
}

type SaleInvoiceModel struct {
	TaskID                  string                             `json:"TaskID,omitempty"`
	InvoiceNumber           string                             `json:"InvoiceNumber,omitempty"`
	Memo                    string                             `json:"Memo,omitempty"`
	Status                  string                             `json:"Status,omitempty"`
	InvoiceDate             string                             `json:"InvoiceDate,omitempty"`
	InvoiceDueDate          string                             `json:"InvoiceDueDate,omitempty"`
	CurrencyConversionRate  float64                            `json:"CurrencyConversionRate,omitempty"`
	BillingAddressLine1     string                             `json:"BillingAddressLine1,omitempty"`
	BillingAddressLine2     string                             `json:"BillingAddressLine2,omitempty"`
	LinkedFulfillmentNumber int                                `json:"LinkedFulfillmentNumber,omitempty"`
	Lines                   []SaleInvoiceLineModel             `json:"Lines,omitempty"`
	AdditionalCharges       []SaleInvoiceAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	Payments                []SalePaymentLineModel             `json:"Payments,omitempty"`
	TotalBeforeTax          float64                            `json:"TotalBeforeTax,omitempty"`
	Tax                     float64                            `json:"Tax,omitempty"`
	Total                   float64                            `json:"Total,omitempty"`
	Paid                    float64                            `json:"Paid,omitempty"`
}

type SaleAdditionalChargeModel struct {
	Description string  `json:"Description,omitempty"`
	Price       float64 `json:"Price,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	Discount    float64 `json:"Discount,omitempty"`
	Tax         float64 `json:"Tax,omitempty"`
	Total       float64 `json:"Total,omitempty"`
	TaxRule     string  `json:"TaxRule,omitempty"`
	Comment     string  `json:"Comment,omitempty"`
}

type SaleInvoiceAdditionalChargeModel struct {
	Description string  `json:"Description,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	Price       float64 `json:"Price,omitempty"`
	Discount    float64 `json:"Discount,omitempty"`
	Tax         float64 `json:"Tax,omitempty"`
	Total       float64 `json:"Total,omitempty"`
	TaxRule     string  `json:"TaxRule,omitempty"`
	Account     string  `json:"Account,omitempty"`
	Comment     string  `json:"Comment,omitempty"`
}

type SaleInvoiceLineModel struct {
	ProductID   string  `json:"ProductID,omitempty"`
	SKU         string  `json:"SKU,omitempty"`
	Name        string  `json:"Name,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	Price       float64 `json:"Price,omitempty"`
	Discount    float64 `json:"Discount,omitempty"`
	Tax         float64 `json:"Tax,omitempty"`
	Total       float64 `json:"Total,omitempty"`
	AverageCost float64 `json:"AverageCost,omitempty"`
	TaxRule     string  `json:"TaxRule,omitempty"`
	Account     string  `json:"Account,omitempty"`
	Comment     string  `json:"Comment,omitempty"`
}

type SaleCreditNoteModel struct {
	TaskID                   string                             `json:"TaskID,omitempty"`
	CreditNoteInvoiceNumber  string                             `json:"CreditNoteInvoiceNumber,omitempty"`
	Memo                     string                             `json:"Memo,omitempty"`
	Status                   string                             `json:"Status,omitempty"`
	CreditNoteDate           string                             `json:"CreditNoteDate,omitempty"`
	CreditNoteNumber         string                             `json:"CreditNoteNumber,omitempty"`
	CreditNoteConversionRate float64                            `json:"CreditNoteConversionRate,omitempty"`
	Lines                    []SaleInvoiceLineModel             `json:"Lines,omitempty"`
	AdditionalCharges        []SaleInvoiceAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	Refunds                  []SalePaymentLineModel             `json:"Refunds,omitempty"`
	Restock                  []SaleFulfilmentPickPackLineModel  `json:"Restock,omitempty"`
	TotalBeforeTax           float64                            `json:"TotalBeforeTax,omitempty"`
	Tax                      float64                            `json:"Tax,omitempty"`
	Total                    float64                            `json:"Total,omitempty"`
}

type SalePaymentLineModel struct {
	ID           string  `json:"ID,omitempty"`
	Reference    string  `json:"Reference,omitempty"`
	Amount       float64 `json:"Amount,omitempty"`
	DatePaid     string  `json:"DatePaid,omitempty"`
	Account      string  `json:"Account,omitempty"`
	CurrencyRate float64 `json:"CurrencyRate,omitempty"`
	DateCreated  string  `json:"DateCreated,omitempty"`
}

type SaleManualJournalModel struct {
	Status string                       `json:"Status,omitempty"`
	Lines  []SaleManualJournalLineModel `json:"Lines,omitempty"`
}

type SaleManualJournalLineModel struct {
	Reference string  `json:"Reference,omitempty"`
	Amount    float64 `json:"Amount,omitempty"`
	Date      string  `json:"Date,omitempty"`
	Debit     string  `json:"Debit,omitempty"`
	Credit    string  `json:"Credit,omitempty"`
}

type AttachmentLineModel struct {
	ID          string `json:"ID,omitempty"`
	ContentType string `json:"ContentType,omitempty"`
	IsDefault   bool   `json:"IsDefault,omitempty"`
	FileName    string `json:"FileName,omitempty"`
	DownloadUrl string `json:"DownloadUrl,omitempty"`
}

type ProductFamilyProductLineModel struct {
	ID      string `json:"ID,omitempty"`
	SKU     string `json:"SKU,omitempty"`
	Name    string `json:"Name,omitempty"`
	Option1 string `json:"Option1,omitempty"`
	Option2 string `json:"Option2,omitempty"`
	Option3 string `json:"Option3,omitempty"`
}

type InventoryMovementLineModel struct {
	TaskID    string  `json:"TaskID,omitempty"`
	ProductID string  `json:"ProductID,omitempty"`
	Date      string  `json:"Date,omitempty"`
	COGS      float64 `json:"COGS,omitempty"`
}

type SaleTransactionLineModel struct {
	TaskID        string  `json:"TaskID,omitempty"`
	TransactionID string  `json:"TransactionID,omitempty"`
	Debit         string  `json:"Debit,omitempty"`
	Credit        string  `json:"Credit,omitempty"`
	Description   string  `json:"Description,omitempty"`
	Amount        float64 `json:"Amount,omitempty"`
	EffectiveDate string  `json:"EffectiveDate,omitempty"`
}

type PurchaseManualJournalModel struct {
	Status string                           `json:"Status,omitempty"`
	Lines  []PurchaseManualJournalLineModel `json:"Lines,omitempty"`
}

type PurchaseManualJournalLineModel struct {
	Reference string  `json:"Reference,omitempty"`
	Amount    float64 `json:"Amount,omitempty"`
	Date      string  `json:"Date,omitempty"`
	Debit     string  `json:"Debit,omitempty"`
	Credit    string  `json:"Credit,omitempty"`
	IsSystem  bool    `json:"IsSystem,omitempty"`
}

type PurchaseOrderModel struct {
	Memo              string                          `json:"Memo,omitempty"`
	Status            string                          `json:"Status,omitempty"`
	Lines             []PurchaseOrderLineModel        `json:"Lines,omitempty"`
	AdditionalCharges []PurchaseAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	Prepayments       []SalePaymentLineModel          `json:"Prepayments,omitempty"`
	TotalBeforeTax    float64                         `json:"TotalBeforeTax,omitempty"`
	Tax               float64                         `json:"Tax,omitempty"`
	Total             float64                         `json:"Total,omitempty"`
}

type PurchaseOrderLineModel struct {
	ProductID   string  `json:"ProductID,omitempty"`
	SKU         string  `json:"SKU,omitempty"`
	Name        string  `json:"Name,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	Price       float64 `json:"Price,omitempty"`
	Discount    float64 `json:"Discount,omitempty"`
	Tax         float64 `json:"Tax,omitempty"`
	TaxRule     string  `json:"TaxRule,omitempty"`
	SupplierSKU string  `json:"SupplierSKU,omitempty"`
	Comment     string  `json:"Comment,omitempty"`
	Total       float64 `json:"Total,omitempty"`
}

type PurchaseAdditionalChargeModel struct {
	Description string  `json:"Description,omitempty"`
	Reference   string  `json:"Reference,omitempty"`
	Price       float64 `json:"Price,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	Discount    float64 `json:"Discount,omitempty"`
	Tax         float64 `json:"Tax,omitempty"`
	Total       float64 `json:"Total,omitempty"`
	TaxRule     string  `json:"TaxRule,omitempty"`
}

type PurchaseStockModel struct {
	Status string                   `json:"Status,omitempty"`
	Lines  []PurchaseStockLineModel `json:"Lines,omitempty"`
}

type PurchaseStockLineModel struct {
	Date        string  `json:"Date,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	ProductID   string  `json:"ProductID,omitempty"`
	SKU         string  `json:"SKU,omitempty"`
	Name        string  `json:"Name,omitempty"`
	Location    string  `json:"Location,omitempty"`
	LocationID  string  `json:"LocationID,omitempty"`
	Received    bool    `json:"Received,omitempty"`
	BatchSN     string  `json:"BatchSN,omitempty"`
	SupplierSKU string  `json:"SupplierSKU,omitempty"`
	ExpiryDate  string  `json:"ExpiryDate,omitempty"`
	CardID      string  `json:"CardID,omitempty"`
}

type PurchaseUnStockLineModel struct {
	CardID     string  `json:"CardID,omitempty"`
	Date       string  `json:"Date,omitempty"`
	Quantity   float64 `json:"Quantity,omitempty"`
	ProductID  string  `json:"ProductID,omitempty"`
	SKU        string  `json:"SKU,omitempty"`
	Name       string  `json:"Name,omitempty"`
	Location   string  `json:"Location,omitempty"`
	BatchSN    string  `json:"BatchSN,omitempty"`
	ExpiryDate string  `json:"ExpiryDate,omitempty"`
}

type PurchaseInvoiceModel struct {
	InvoiceDate       string                                 `json:"InvoiceDate,omitempty"`
	InvoiceDueDate    string                                 `json:"InvoiceDueDate,omitempty"`
	InvoiceNumber     string                                 `json:"InvoiceNumber,omitempty"`
	Status            string                                 `json:"Status,omitempty"`
	Lines             []PurchaseInvoiceLineModel             `json:"Lines,omitempty"`
	AdditionalCharges []PurchaseInvoiceAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	Payments          []SalePaymentLineModel                 `json:"Payments,omitempty"`
	TotalBeforeTax    float64                                `json:"TotalBeforeTax,omitempty"`
	Tax               float64                                `json:"Tax,omitempty"`
	Total             float64                                `json:"Total,omitempty"`
	Paid              float64                                `json:"Paid,omitempty"`
}

type PurchaseCreditNoteModel struct {
	CreditNoteNumber  string                                 `json:"CreditNoteNumber,omitempty"`
	CreditNoteDate    string                                 `json:"CreditNoteDate,omitempty"`
	Status            string                                 `json:"Status,omitempty"`
	Lines             []PurchaseInvoiceLineModel             `json:"Lines,omitempty"`
	AdditionalCharges []PurchaseInvoiceAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	Refunds           []SalePaymentLineModel                 `json:"Refunds,omitempty"`
	Unstock           []PurchaseUnStockLineModel             `json:"Unstock,omitempty"`
	TotalBeforeTax    float64                                `json:"TotalBeforeTax,omitempty"`
	Tax               float64                                `json:"Tax,omitempty"`
	Total             float64                                `json:"Total,omitempty"`
}

type PurchaseInvoiceLineModel struct {
	ProductID string  `json:"ProductID,omitempty"`
	SKU       string  `json:"SKU,omitempty"`
	Name      string  `json:"Name,omitempty"`
	Quantity  float64 `json:"Quantity,omitempty"`
	Price     float64 `json:"Price,omitempty"`
	Discount  float64 `json:"Discount,omitempty"`
	Tax       float64 `json:"Tax,omitempty"`
	TaxRule   string  `json:"TaxRule,omitempty"`
	Account   string  `json:"Account,omitempty"`
	Comment   string  `json:"Comment,omitempty"`
	Total     float64 `json:"Total,omitempty"`
}

type PurchaseInvoiceAdditionalChargeModel struct {
	Description string  `json:"Description,omitempty"`
	Reference   string  `json:"Reference,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	Price       float64 `json:"Price,omitempty"`
	Discount    float64 `json:"Discount,omitempty"`
	Tax         float64 `json:"Tax,omitempty"`
	TaxRule     string  `json:"TaxRule,omitempty"`
	Account     string  `json:"Account,omitempty"`
	Total       float64 `json:"Total,omitempty"`
}

type ExistingStockLineModel struct {
	ProductID            string                `json:"ProductID,omitempty"`
	SKU                  string                `json:"SKU,omitempty"`
	ProductName          string                `json:"ProductName,omitempty"`
	QuantityOnHand       float64               `json:"QuantityOnHand,omitempty"`
	Available            float64               `json:"Available,omitempty"`
	Adjustment           float64               `json:"Adjustment,omitempty"`
	LocationID           string                `json:"LocationID,omitempty"`
	Location             string                `json:"Location,omitempty"`
	BatchSN              string                `json:"BatchSN,omitempty"`
	ExpiryDate           string                `json:"ExpiryDate,omitempty"`
	Comments             string                `json:"Comments,omitempty"`
	ProductLength        float64               `json:"ProductLength,omitempty"`
	ProductWidth         float64               `json:"ProductWidth,omitempty"`
	ProductHeight        float64               `json:"ProductHeight,omitempty"`
	ProductWeight        float64               `json:"ProductWeight,omitempty"`
	WeightUnits          string                `json:"WeightUnits,omitempty"`
	DimensionsUnits      string                `json:"DimensionsUnits,omitempty"`
	ProductCustomField1  string                `json:"ProductCustomField1,omitempty"`
	ProductCustomField2  string                `json:"ProductCustomField2,omitempty"`
	ProductCustomField3  string                `json:"ProductCustomField3,omitempty"`
	ProductCustomField4  string                `json:"ProductCustomField4,omitempty"`
	ProductCustomField5  string                `json:"ProductCustomField5,omitempty"`
	ProductCustomField6  string                `json:"ProductCustomField6,omitempty"`
	ProductCustomField7  string                `json:"ProductCustomField7,omitempty"`
	ProductCustomField8  string                `json:"ProductCustomField8,omitempty"`
	ProductCustomField9  string                `json:"ProductCustomField9,omitempty"`
	ProductCustomField10 string                `json:"ProductCustomField10,omitempty"`
	Image                []AttachmentLineModel `json:"Image,omitempty"`
	Barcode              string                `json:"Barcode,omitempty"`
	StockLocator         string                `json:"StockLocator,omitempty"`
	Unit                 string                `json:"Unit,omitempty"`
	CostingMethod        string                `json:"CostingMethod,omitempty"`
}

type NewStockLineModel struct {
	ProductID            string                `json:"ProductID,omitempty"`
	SKU                  string                `json:"SKU,omitempty"`
	ProductName          string                `json:"ProductName,omitempty"`
	Quantity             float64               `json:"Quantity,omitempty"`
	UnitCost             float64               `json:"UnitCost,omitempty"`
	LocationID           string                `json:"LocationID,omitempty"`
	Location             string                `json:"Location,omitempty"`
	BatchSN              string                `json:"BatchSN,omitempty"`
	ExpiryDate           string                `json:"ExpiryDate,omitempty"`
	ReceivedDate         string                `json:"ReceivedDate,omitempty"`
	Comments             string                `json:"Comments,omitempty"`
	ProductLength        float64               `json:"ProductLength,omitempty"`
	ProductWidth         float64               `json:"ProductWidth,omitempty"`
	ProductHeight        float64               `json:"ProductHeight,omitempty"`
	ProductWeight        float64               `json:"ProductWeight,omitempty"`
	WeightUnits          string                `json:"WeightUnits,omitempty"`
	DimensionsUnits      string                `json:"DimensionsUnits,omitempty"`
	ProductCustomField1  string                `json:"ProductCustomField1,omitempty"`
	ProductCustomField2  string                `json:"ProductCustomField2,omitempty"`
	ProductCustomField3  string                `json:"ProductCustomField3,omitempty"`
	ProductCustomField4  string                `json:"ProductCustomField4,omitempty"`
	ProductCustomField5  string                `json:"ProductCustomField5,omitempty"`
	ProductCustomField6  string                `json:"ProductCustomField6,omitempty"`
	ProductCustomField7  string                `json:"ProductCustomField7,omitempty"`
	ProductCustomField8  string                `json:"ProductCustomField8,omitempty"`
	ProductCustomField9  string                `json:"ProductCustomField9,omitempty"`
	ProductCustomField10 string                `json:"ProductCustomField10,omitempty"`
	Image                []AttachmentLineModel `json:"Image,omitempty"`
	Barcode              string                `json:"Barcode,omitempty"`
	StockLocator         string                `json:"StockLocator,omitempty"`
	Unit                 string                `json:"Unit,omitempty"`
	CostingMethod        string                `json:"CostingMethod,omitempty"`
}

type TransactionStockLineModel struct {
	ID            string  `json:"ID,omitempty"`
	Debit         string  `json:"Debit,omitempty"`
	Credit        string  `json:"Credit,omitempty"`
	Amount        float64 `json:"Amount,omitempty"`
	EffectiveDate string  `json:"EffectiveDate,omitempty"`
}

type StockTransferLineModel struct {
	ProductID            string  `json:"ProductID,omitempty"`
	SKU                  string  `json:"SKU,omitempty"`
	ProductName          string  `json:"ProductName,omitempty"`
	QuantityOnHand       float64 `json:"QuantityOnHand,omitempty"`
	QuantityAvailable    float64 `json:"QuantityAvailable,omitempty"`
	TransferQuantity     float64 `json:"TransferQuantity,omitempty"`
	BatchSN              string  `json:"BatchSN,omitempty"`
	ExpiryDate           string  `json:"ExpiryDate,omitempty"`
	Comments             string  `json:"Comments,omitempty"`
	ProductLength        float64 `json:"ProductLength,omitempty"`
	ProductWidth         float64 `json:"ProductWidth,omitempty"`
	ProductHeight        float64 `json:"ProductHeight,omitempty"`
	ProductWeight        float64 `json:"ProductWeight,omitempty"`
	WeightUnits          string  `json:"WeightUnits,omitempty"`
	DimensionsUnits      string  `json:"DimensionsUnits,omitempty"`
	ProductCustomField1  string  `json:"ProductCustomField1,omitempty"`
	ProductCustomField2  string  `json:"ProductCustomField2,omitempty"`
	ProductCustomField3  string  `json:"ProductCustomField3,omitempty"`
	ProductCustomField4  string  `json:"ProductCustomField4,omitempty"`
	ProductCustomField5  string  `json:"ProductCustomField5,omitempty"`
	ProductCustomField6  string  `json:"ProductCustomField6,omitempty"`
	ProductCustomField7  string  `json:"ProductCustomField7,omitempty"`
	ProductCustomField8  string  `json:"ProductCustomField8,omitempty"`
	ProductCustomField9  string  `json:"ProductCustomField9,omitempty"`
	ProductCustomField10 string  `json:"ProductCustomField10,omitempty"`
}

type PriceTierModel struct {
	Price float64 `json:"Price,omitempty"`
}

type ProductSupplierModel struct {
	SupplierID             string                        `json:"SupplierID,omitempty"`
	SupplierName           string                        `json:"SupplierName,omitempty"`
	ProductID              string                        `json:"ProductID,omitempty"`
	ProductSKU             string                        `json:"ProductSKU,omitempty"`
	ProductSupplierID      string                        `json:"ProductSupplierID,omitempty"`
	SupplierInventoryCode  string                        `json:"SupplierInventoryCode,omitempty"`
	SupplierProductName    string                        `json:"SupplierProductName,omitempty"`
	Cost                   float64                       `json:"Cost,omitempty"`
	FixedCost              float64                       `json:"FixedCost,omitempty"`
	Currency               string                        `json:"Currency,omitempty"`
	DropShip               bool                          `json:"DropShip,omitempty"`
	SupplierProductURL     string                        `json:"SupplierProductURL,omitempty"`
	LastSupplied           string                        `json:"LastSupplied,omitempty"`
	ProductSupplierOptions []ProductSupplierOptionsModel `json:"ProductSupplierOptions,omitempty"`
}

type ProductSupplierOptionsModel struct {
	ID               string                                `json:"ID,omitempty"`
	LocationID       string                                `json:"LocationID,omitempty"`
	LocationName     string                                `json:"LocationName,omitempty"`
	ReorderQuantity  float64                               `json:"ReorderQuantity,omitempty"`
	Lead             int                                   `json:"Lead,omitempty"`
	Safety           int                                   `json:"Safety,omitempty"`
	MinimumToReorder float64                               `json:"MinimumToReorder,omitempty"`
	SupplyIntervals  []ProductSupplierOptionsIntervalModel `json:"SupplyIntervals,omitempty"`
}

type ProductSupplierOptionsIntervalModel struct {
	ID                string `json:"ID,omitempty"`
	DeliveryMethod    string `json:"DeliveryMethod,omitempty"`
	IntervalDays      int    `json:"IntervalDays,omitempty"`
	IntervalStartDate string `json:"IntervalStartDate,omitempty"`
	IsMonday          bool   `json:"IsMonday,omitempty"`
	IsTuesday         bool   `json:"IsTuesday,omitempty"`
	IsWednesday       bool   `json:"IsWednesday,omitempty"`
	IsThursday        bool   `json:"IsThursday,omitempty"`
	IsFriday          bool   `json:"IsFriday,omitempty"`
	IsSaturday        bool   `json:"IsSaturday,omitempty"`
	IsSunday          bool   `json:"IsSunday,omitempty"`
}

type ReorderLevelModel struct {
	LocationID           string  `json:"LocationID,omitempty"`
	LocationName         string  `json:"LocationName,omitempty"`
	MinimumBeforeReorder float64 `json:"MinimumBeforeReorder,omitempty"`
	ReorderQuantity      float64 `json:"ReorderQuantity,omitempty"`
	StockLocator         string  `json:"StockLocator,omitempty"`
	PickZones            string  `json:"PickZones,omitempty"`
}

type BillOfMaterialProductModel struct {
	ComponentProductID string  `json:"ComponentProductID,omitempty"`
	ProductCode        string  `json:"ProductCode,omitempty"`
	Name               string  `json:"Name,omitempty"`
	Quantity           float64 `json:"Quantity,omitempty"`
	WastagePercent     float64 `json:"WastagePercent,omitempty"`
	WastageQuantity    float64 `json:"WastageQuantity,omitempty"`
	CostPercentage     float64 `json:"CostPercentage,omitempty"`
}

type BillOfMaterialServiceModel struct {
	ComponentProductID string  `json:"ComponentProductID,omitempty"`
	Name               string  `json:"Name,omitempty"`
	Quantity           float64 `json:"Quantity,omitempty"`
	ExpenseAccount     string  `json:"ExpenseAccount,omitempty"`
	PriceTier          int     `json:"PriceTier,omitempty"`
}

type ProductMovementModel struct {
	TaskID     string  `json:"TaskID,omitempty"`
	Type       string  `json:"Type,omitempty"`
	Date       string  `json:"Date,omitempty"`
	Number     string  `json:"Number,omitempty"`
	Status     int     `json:"Status,omitempty"`
	Quantity   float64 `json:"Quantity,omitempty"`
	Amount     float64 `json:"Amount,omitempty"`
	Location   float64 `json:"Location,omitempty"`
	BatchSN    float64 `json:"BatchSN,omitempty"`
	ExpiryDate string  `json:"ExpiryDate,omitempty"`
	FromTo     string  `json:"FromTo,omitempty"`
}

type AttributeSetLineModel struct {
	Name   string `json:"Name,omitempty"`
	Type   string `json:"Type,omitempty"`
	Values string `json:"Values,omitempty"`
}

type ErrorModel struct {
	ErrorCode int    `json:"ErrorCode,omitempty"`
	Exception string `json:"Exception,omitempty"`
}

type DisassemblyPickLineModel struct {
	BinID      string  `json:"BinID,omitempty"`
	Bin        string  `json:"Bin,omitempty"`
	Date       string  `json:"Date,omitempty"`
	BatchSN    string  `json:"BatchSN,omitempty"`
	ExpiryDate string  `json:"ExpiryDate,omitempty"`
	Available  float64 `json:"Available,omitempty"`
	UnitCost   float64 `json:"UnitCost,omitempty"`
	Quantity   float64 `json:"Quantity,omitempty"`
	TotalCost  float64 `json:"TotalCost,omitempty"`
}

type FinishedGoodsPickLineModel struct {
	ProductID   string  `json:"ProductID,omitempty"`
	ProductCode string  `json:"ProductCode,omitempty"`
	Name        string  `json:"Name,omitempty"`
	BinID       string  `json:"BinID,omitempty"`
	Bin         string  `json:"Bin,omitempty"`
	BatchSN     string  `json:"BatchSN,omitempty"`
	ExpiryDate  string  `json:"ExpiryDate,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	Unit        string  `json:"Unit,omitempty"`
	Cost        float64 `json:"Cost,omitempty"`
}

type FinishedGoodsOrderLineModel struct {
	ProductID       string  `json:"ProductID,omitempty"`
	ProductCode     string  `json:"ProductCode,omitempty"`
	Name            string  `json:"Name,omitempty"`
	ExpenseAccount  string  `json:"ExpenseAccount,omitempty"`
	Quantity        float64 `json:"Quantity,omitempty"`
	Unit            string  `json:"Unit,omitempty"`
	WastagePercent  float64 `json:"WastagePercent,omitempty"`
	WastageQuantity float64 `json:"WastageQuantity,omitempty"`
	TotalQuantity   float64 `json:"TotalQuantity,omitempty"`
	TotalCost       float64 `json:"TotalCost,omitempty"`
}

type DisassemblyOrderLineModel struct {
	ProductID   string  `json:"ProductID,omitempty"`
	ProductCode string  `json:"ProductCode,omitempty"`
	Name        string  `json:"Name,omitempty"`
	BinID       string  `json:"BinID,omitempty"`
	Bin         string  `json:"Bin,omitempty"`
	Unit        string  `json:"Unit,omitempty"`
	BatchSN     string  `json:"BatchSN,omitempty"`
	ExpiryDate  string  `json:"ExpiryDate,omitempty"`
	Account     string  `json:"Account,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	Cost        float64 `json:"Cost,omitempty"`
}

type DisassemblyOrderServiceLineModel struct {
	ProductID string  `json:"ProductID,omitempty"`
	Name      string  `json:"Name,omitempty"`
	Account   string  `json:"Account,omitempty"`
	Comments  string  `json:"Comments,omitempty"`
	Amount    float64 `json:"Amount,omitempty"`
}

type InventoryWriteOffLineModel struct {
	ProductID      string  `json:"ProductID,omitempty"`
	ProductCode    string  `json:"ProductCode,omitempty"`
	Name           string  `json:"Name,omitempty"`
	BinID          string  `json:"BinID,omitempty"`
	Bin            string  `json:"Bin,omitempty"`
	BatchSN        string  `json:"BatchSN,omitempty"`
	ExpiryDate     string  `json:"ExpiryDate,omitempty"`
	ExpenseAccount string  `json:"ExpenseAccount,omitempty"`
	Quantity       float64 `json:"Quantity,omitempty"`
	Cost           float64 `json:"Cost,omitempty"`
	TotalCost      float64 `json:"TotalCost,omitempty"`
}

type TaxComponentModel struct {
	Name           string  `json:"Name,omitempty"`
	Percent        float64 `json:"Percent,omitempty"`
	AccountCode    string  `json:"AccountCode,omitempty"`
	ComponentOrder int     `json:"ComponentOrder,omitempty"`
}

type MoneyTaskLineModel struct {
	Name     string  `json:"Name,omitempty"`
	Comment  string  `json:"Comment,omitempty"`
	Quantity float64 `json:"Quantity,omitempty"`
	Price    float64 `json:"Price,omitempty"`
	Discount float64 `json:"Discount,omitempty"`
	Tax      float64 `json:"Tax,omitempty"`
	TaxRule  string  `json:"TaxRule,omitempty"`
	Account  string  `json:"Account,omitempty"`
	Total    float64 `json:"Total,omitempty"`
}

type SupplierAddressModel struct {
	ID             string `json:"ID,omitempty"`
	Line1          string `json:"Line1,omitempty"`
	Line2          string `json:"Line2,omitempty"`
	City           string `json:"City,omitempty"`
	State          string `json:"State,omitempty"`
	Postcode       string `json:"Postcode,omitempty"`
	Country        string `json:"Country,omitempty"`
	Type           string `json:"Type,omitempty"`
	DefaultForType bool   `json:"DefaultForType,omitempty"`
}

type SupplierContactModel struct {
	ID             string `json:"ID,omitempty"`
	Name           string `json:"Name,omitempty"`
	Phone          string `json:"Phone,omitempty"`
	MobilePhone    string `json:"MobilePhone,omitempty"`
	Fax            string `json:"Fax,omitempty"`
	Email          string `json:"Email,omitempty"`
	Website        string `json:"Website,omitempty"`
	Comment        string `json:"Comment,omitempty"`
	Default        bool   `json:"Default,omitempty"`
	IncludeInEmail bool   `json:"IncludeInEmail,omitempty"`
}

type ProductPriceModel struct {
	ProductID    string  `json:"ProductID,omitempty"`
	CustomerID   string  `json:"CustomerID,omitempty"`
	CustomerName string  `json:"CustomerName,omitempty"`
	ProductSKU   string  `json:"ProductSKU,omitempty"`
	ProductName  string  `json:"ProductName,omitempty"`
	Price        float64 `json:"Price,omitempty"`
}

type JournalLineModel struct {
	Debit      string  `json:"Debit,omitempty"`
	Credit     string  `json:"Credit,omitempty"`
	Reference  string  `json:"Reference,omitempty"`
	Amount     float64 `json:"Amount,omitempty"`
	BaseAmount float64 `json:"BaseAmount,omitempty"`
}

type PurchasePaymentLineModel struct {
	PurchaseID   string  `json:"PurchaseID,omitempty"`
	TaskID       string  `json:"TaskID,omitempty"`
	ID           string  `json:"ID,omitempty"`
	Reference    string  `json:"Reference,omitempty"`
	Amount       float64 `json:"Amount,omitempty"`
	DatePaid     string  `json:"DatePaid,omitempty"`
	Account      string  `json:"Account,omitempty"`
	CurrencyRate float64 `json:"CurrencyRate,omitempty"`
	DateCreated  string  `json:"DateCreated,omitempty"`
}

type AdvancedPurchaseStockModel struct {
	TaskID string                           `json:"TaskID,omitempty"`
	Status string                           `json:"Status,omitempty"`
	Lines  []AdvancedPurchaseStockLineModel `json:"Lines,omitempty"`
}

type AdvancedPurchaseStockLineModel struct {
	Date        string  `json:"Date,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	ProductID   string  `json:"ProductID,omitempty"`
	SKU         string  `json:"SKU,omitempty"`
	Name        string  `json:"Name,omitempty"`
	Location    string  `json:"Location,omitempty"`
	LocationID  string  `json:"LocationID,omitempty"`
	Received    bool    `json:"Received,omitempty"`
	BatchSN     string  `json:"BatchSN,omitempty"`
	SupplierSKU string  `json:"SupplierSKU,omitempty"`
	ExpiryDate  string  `json:"ExpiryDate,omitempty"`
}

type AdvancedPurchasePutAwayModel struct {
	TaskID string                             `json:"TaskID,omitempty"`
	Status string                             `json:"Status,omitempty"`
	Lines  []AdvancedPurchasePutAwayLineModel `json:"Lines,omitempty"`
}

type AdvancedPurchasePutAwayLineModel struct {
	Date        string  `json:"Date,omitempty"`
	Quantity    float64 `json:"Quantity,omitempty"`
	ProductID   string  `json:"ProductID,omitempty"`
	SKU         string  `json:"SKU,omitempty"`
	Name        string  `json:"Name,omitempty"`
	Location    string  `json:"Location,omitempty"`
	LocationID  string  `json:"LocationID,omitempty"`
	Received    bool    `json:"Received,omitempty"`
	BatchSN     string  `json:"BatchSN,omitempty"`
	SupplierSKU string  `json:"SupplierSKU,omitempty"`
	ExpiryDate  string  `json:"ExpiryDate,omitempty"`
	CardID      string  `json:"CardID,omitempty"`
}

type AdvancedPurchaseInvoiceModel struct {
	TaskID            string                                 `json:"TaskID,omitempty"`
	InvoiceDate       string                                 `json:"InvoiceDate,omitempty"`
	InvoiceDueDate    string                                 `json:"InvoiceDueDate,omitempty"`
	InvoiceNumber     string                                 `json:"InvoiceNumber,omitempty"`
	Status            string                                 `json:"Status,omitempty"`
	Lines             []PurchaseInvoiceLineModel             `json:"Lines,omitempty"`
	AdditionalCharges []PurchaseInvoiceAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	Payments          []PurchasePaymentLineModel             `json:"Payments,omitempty"`
	TotalBeforeTax    float64                                `json:"TotalBeforeTax,omitempty"`
	Tax               float64                                `json:"Tax,omitempty"`
	Total             float64                                `json:"Total,omitempty"`
	Paid              float64                                `json:"Paid,omitempty"`
}

type AdvancedPurchaseCreditNoteModel struct {
	TaskID            string                                 `json:"TaskID,omitempty"`
	CreditNoteNumber  string                                 `json:"CreditNoteNumber,omitempty"`
	CreditNoteDate    string                                 `json:"CreditNoteDate,omitempty"`
	Status            string                                 `json:"Status,omitempty"`
	Lines             []PurchaseInvoiceLineModel             `json:"Lines,omitempty"`
	AdditionalCharges []PurchaseInvoiceAdditionalChargeModel `json:"AdditionalCharges,omitempty"`
	Refunds           []PurchasePaymentLineModel             `json:"Refunds,omitempty"`
	Unstock           []PurchaseUnStockLineModel             `json:"Unstock,omitempty"`
	TotalBeforeTax    float64                                `json:"TotalBeforeTax,omitempty"`
	Tax               float64                                `json:"Tax,omitempty"`
	Total             float64                                `json:"Total,omitempty"`
}

type AdvancedPurchaseManualJournalModel struct {
	TaskID string                           `json:"TaskID,omitempty"`
	Status string                           `json:"Status,omitempty"`
	Lines  []PurchaseManualJournalLineModel `json:"Lines,omitempty"`
}

type IDNameModel struct {
	ID   string `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
}

type StockTransferOrderModel struct {
	Status string                        `json:"Status,omitempty"`
	Lines  []StockTransferOrderLineModel `json:"Lines,omitempty"`
}

type StockTransferOrderLineModel struct {
	ProductID            string  `json:"ProductID,omitempty"`
	SKU                  string  `json:"SKU,omitempty"`
	ProductName          string  `json:"ProductName,omitempty"`
	QuantityOnHand       float64 `json:"QuantityOnHand,omitempty"`
	QuantityAvailable    float64 `json:"QuantityAvailable,omitempty"`
	TransferQuantity     float64 `json:"TransferQuantity,omitempty"`
	Comments             string  `json:"Comments,omitempty"`
	ProductLength        float64 `json:"ProductLength,omitempty"`
	ProductWidth         float64 `json:"ProductWidth,omitempty"`
	ProductHeight        float64 `json:"ProductHeight,omitempty"`
	ProductWeight        float64 `json:"ProductWeight,omitempty"`
	WeightUnits          string  `json:"WeightUnits,omitempty"`
	DimensionsUnits      string  `json:"DimensionsUnits,omitempty"`
	ProductCustomField1  string  `json:"ProductCustomField1,omitempty"`
	ProductCustomField2  string  `json:"ProductCustomField2,omitempty"`
	ProductCustomField3  string  `json:"ProductCustomField3,omitempty"`
	ProductCustomField4  string  `json:"ProductCustomField4,omitempty"`
	ProductCustomField5  string  `json:"ProductCustomField5,omitempty"`
	ProductCustomField6  string  `json:"ProductCustomField6,omitempty"`
	ProductCustomField7  string  `json:"ProductCustomField7,omitempty"`
	ProductCustomField8  string  `json:"ProductCustomField8,omitempty"`
	ProductCustomField9  string  `json:"ProductCustomField9,omitempty"`
	ProductCustomField10 string  `json:"ProductCustomField10,omitempty"`
}

type FinishedGoods struct {
	TaskID                 string                        `json:"TaskID,omitempty"`
	AssemblyNumber         string                        `json:"AssemblyNumber,omitempty"`
	Status                 string                        `json:"Status,omitempty"`
	ProductID              string                        `json:"ProductID,omitempty"`
	ProductCode            string                        `json:"ProductCode,omitempty"`
	ProductName            string                        `json:"ProductName,omitempty"`
	LocationID             string                        `json:"LocationID,omitempty"`
	Location               float64                       `json:"Location,omitempty"`
	BinID                  string                        `json:"BinID,omitempty"`
	Bin                    float64                       `json:"Bin,omitempty"`
	WIPAccount             string                        `json:"WIPAccount,omitempty"`
	WIPDate                string                        `json:"WIPDate,omitempty"`
	Account                string                        `json:"Account,omitempty"`
	Quantity               float64                       `json:"Quantity,omitempty"`
	AssemblyInstructionURL string                        `json:"AssemblyInstructionURL,omitempty"`
	CompletionDate         string                        `json:"CompletionDate,omitempty"`
	BatchSN                string                        `json:"BatchSN,omitempty"`
	ExpiryDate             string                        `json:"ExpiryDate,omitempty"`
	Notes                  string                        `json:"Notes,omitempty"`
	OrderLines             []FinishedGoodsOrderLineModel `json:"OrderLines,omitempty"`
	PickLines              []FinishedGoodsPickLineModel  `json:"PickLines,omitempty"`
	Transactions           []TransactionStockLineModel   `json:"Transactions,omitempty"`
	Errors                 []ErrorModel                  `json:"Errors,omitempty"`
	CustomField1           string                        `json:"CustomField1,omitempty"`
	CustomField2           string                        `json:"CustomField2,omitempty"`
	CustomField3           string                        `json:"CustomField3,omitempty"`
	CustomField4           string                        `json:"CustomField4,omitempty"`
	CustomField5           string                        `json:"CustomField5,omitempty"`
	CustomField6           string                        `json:"CustomField6,omitempty"`
	CustomField7           string                        `json:"CustomField7,omitempty"`
	CustomField8           string                        `json:"CustomField8,omitempty"`
	CustomField9           string                        `json:"CustomField9,omitempty"`
	CustomField10          string                        `json:"CustomField10,omitempty"`
}

type DisassemblyOrder struct {
	TaskID     string                      `json:"TaskID,omitempty"`
	Status     string                      `json:"Status,omitempty"`
	OrderLines []DisassemblyOrderLineModel `json:"OrderLines,omitempty"`
}

type Webhooks struct {
	ID                        string              `json:"ID,omitempty"`
	Type                      string              `json:"Type,omitempty"`
	Name                      string              `json:"Name,omitempty"`
	IsActive                  bool                `json:"IsActive,omitempty"`
	ExternalURL               string              `json:"ExternalURL,omitempty"`
	ExternalAuthorizationType string              `json:"ExternalAuthorizationType,omitempty"`
	ExternalUserName          string              `json:"ExternalUserName,omitempty"`
	ExternalPassword          string              `json:"ExternalPassword,omitempty"`
	ExternalBearerToken       string              `json:"ExternalBearerToken,omitempty"`
	ExternalHeaders           []map[string]string `json:"ExternalHeaders,omitempty"`
}

type FinishedGoodsList struct {
	TaskID         string  `json:"TaskID,omitempty"`
	AssemblyNumber string  `json:"AssemblyNumber,omitempty"`
	BatchSN        string  `json:"BatchSN,omitempty"`
	ExpiryDate     string  `json:"ExpiryDate,omitempty"`
	ProductID      string  `json:"ProductID,omitempty"`
	ProductCode    string  `json:"ProductCode,omitempty"`
	ProductName    string  `json:"ProductName,omitempty"`
	Quantity       float64 `json:"Quantity,omitempty"`
	LocationID     string  `json:"LocationID,omitempty"`
	Location       float64 `json:"Location,omitempty"`
	Date           string  `json:"Date,omitempty"`
	Status         string  `json:"Status,omitempty"`
	UnitCost       float64 `json:"UnitCost,omitempty"`
	Notes          string  `json:"Notes,omitempty"`
	CustomField1   string  `json:"CustomField1,omitempty"`
	CustomField2   string  `json:"CustomField2,omitempty"`
	CustomField3   string  `json:"CustomField3,omitempty"`
	CustomField4   string  `json:"CustomField4,omitempty"`
	CustomField5   string  `json:"CustomField5,omitempty"`
	CustomField6   string  `json:"CustomField6,omitempty"`
	CustomField7   string  `json:"CustomField7,omitempty"`
	CustomField8   string  `json:"CustomField8,omitempty"`
	CustomField9   string  `json:"CustomField9,omitempty"`
	CustomField10  string  `json:"CustomField10,omitempty"`
}

type InventoryWriteOffList struct {
	TaskID                  string  `json:"TaskID,omitempty"`
	InventoryWriteOffNumber string  `json:"InventoryWriteOffNumber,omitempty"`
	Status                  string  `json:"Status,omitempty"`
	LocationID              string  `json:"LocationID,omitempty"`
	Location                float64 `json:"Location,omitempty"`
	Date                    string  `json:"Date,omitempty"`
	Notes                   string  `json:"Notes,omitempty"`
}

type InventoryWriteOff struct {
	TaskID                  string                       `json:"TaskID,omitempty"`
	InventoryWriteOffNumber string                       `json:"InventoryWriteOffNumber,omitempty"`
	Status                  string                       `json:"Status,omitempty"`
	LocationID              string                       `json:"LocationID,omitempty"`
	Location                float64                      `json:"Location,omitempty"`
	Account                 string                       `json:"Account,omitempty"`
	EffectiveDate           string                       `json:"EffectiveDate,omitempty"`
	Notes                   string                       `json:"Notes,omitempty"`
	Lines                   []InventoryWriteOffLineModel `json:"Lines,omitempty"`
	Transactions            []TransactionStockLineModel  `json:"Transactions,omitempty"`
	Errors                  []ErrorModel                 `json:"Errors,omitempty"`
}

type Journal struct {
	TaskID                 string                `json:"TaskID,omitempty"`
	Status                 string                `json:"Status,omitempty"`
	JournalNumber          string                `json:"JournalNumber,omitempty"`
	Currency               string                `json:"Currency,omitempty"`
	CurrencyConversionRate float64               `json:"CurrencyConversionRate,omitempty"`
	EffectiveDate          string                `json:"EffectiveDate,omitempty"`
	Narration              string                `json:"Narration,omitempty"`
	Notes                  string                `json:"Notes,omitempty"`
	Lines                  []JournalLineModel    `json:"Lines,omitempty"`
	Attachments            []AttachmentLineModel `json:"Attachments,omitempty"`
}

type MoneyTaskList struct {
	TaskID                 string                      `json:"TaskID,omitempty"`
	TaskType               string                      `json:"TaskType,omitempty"`
	Status                 string                      `json:"Status,omitempty"`
	BankAccount            string                      `json:"BankAccount,omitempty"`
	CurrencyConversionRate float64                     `json:"CurrencyConversionRate,omitempty"`
	SupplierCustomerName   string                      `json:"SupplierCustomerName,omitempty"`
	SupplierID             string                      `json:"SupplierID,omitempty"`
	CustomerID             string                      `json:"CustomerID,omitempty"`
	Reference              string                      `json:"Reference,omitempty"`
	Date                   string                      `json:"Date,omitempty"`
	TaxInclusive           bool                        `json:"TaxInclusive,omitempty"`
	Note                   string                      `json:"Note,omitempty"`
	Lines                  []MoneyTaskLineModel        `json:"Lines,omitempty"`
	Transactions           []TransactionStockLineModel `json:"Transactions,omitempty"`
	Attachments            []AttachmentLineModel       `json:"Attachments,omitempty"`
}

type PurchaseManualJournal struct {
	TaskID string                           `json:"TaskID,omitempty"`
	Status string                           `json:"Status,omitempty"`
	Lines  []PurchaseManualJournalLineModel `json:"Lines,omitempty"`
}

type PurchaseAttachments struct {
	TaskID string                `json:"TaskID,omitempty"`
	Lines  []AttachmentLineModel `json:"Lines,omitempty"`
}

type AdvancedPurchasePartialMAnJModel struct {
	TaskID string                           `json:"TaskID,omitempty"`
	Status string                           `json:"Status,omitempty"`
	Lines  []PurchaseManualJournalLineModel `json:"Lines,omitempty"`
}

type SaleManualJournal struct {
	SaleID string                       `json:"SaleID,omitempty"`
	Status string                       `json:"Status,omitempty"`
	Lines  []SaleManualJournalLineModel `json:"Lines,omitempty"`
}

type SaleAttachments struct {
	SaleID string                `json:"SaleID,omitempty"`
	Lines  []AttachmentLineModel `json:"Lines,omitempty"`
}

type StockAdjustmentList struct {
	TaskID          string `json:"TaskID,omitempty"`
	EffectiveDate   string `json:"EffectiveDate,omitempty"`
	StocktakeNumber string `json:"StocktakeNumber,omitempty"`
	Status          string `json:"Status,omitempty"`
	Account         string `json:"Account,omitempty"`
	Reference       string `json:"Reference,omitempty"`
}

type StockTakeList struct {
	TaskID          string `json:"TaskID,omitempty"`
	EffectiveDate   string `json:"EffectiveDate,omitempty"`
	StocktakeNumber string `json:"StocktakeNumber,omitempty"`
	Status          string `json:"Status,omitempty"`
	Account         string `json:"Account,omitempty"`
	LocationID      string `json:"LocationID,omitempty"`
	Location        string `json:"Location,omitempty"`
	Tags            string `json:"Tags,omitempty"`
	PickZones       string `json:"PickZones,omitempty"`
	StockLocators   string `json:"StockLocators,omitempty"`
	Categories      string `json:"Categories,omitempty"`
	Brands          string `json:"Brands,omitempty"`
	Bins            string `json:"Bins,omitempty"`
	Reference       string `json:"Reference,omitempty"`
	CreatedDate     string `json:"CreatedDate,omitempty"`
	LastUpdatedDate string `json:"LastUpdatedDate,omitempty"`
	LastUpdatedBy   string `json:"LastUpdatedBy,omitempty"`
}

type StockTransferList struct {
	TaskID               string `json:"TaskID,omitempty"`
	From                 string `json:"From,omitempty"`
	FromLocation         string `json:"FromLocation,omitempty"`
	To                   string `json:"To,omitempty"`
	ToLocation           string `json:"ToLocation,omitempty"`
	Status               string `json:"Status,omitempty"`
	Number               string `json:"Number,omitempty"`
	CompletionDate       string `json:"CompletionDate,omitempty"`
	CostDistributionType string `json:"CostDistributionType,omitempty"`
	InTransitAccount     string `json:"InTransitAccount,omitempty"`
	DepartureDate        string `json:"DepartureDate,omitempty"`
	Reference            string `json:"Reference,omitempty"`
}

type Transactions struct {
	TaskID            string  `json:"TaskID,omitempty"`
	DebitAccountCode  string  `json:"DebitAccountCode,omitempty"`
	CreditAccountCode string  `json:"CreditAccountCode,omitempty"`
	Amount            float64 `json:"Amount,omitempty"`
	EffectiveDate     string  `json:"EffectiveDate,omitempty"`
	Reference         string  `json:"Reference,omitempty"`
	Transaction       string  `json:"Transaction,omitempty"`
	Type              string  `json:"Type,omitempty"`
}
