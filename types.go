package gocin7corerestapi


type AttributeSet struct {
	ID  	string           
	Name  	string             
	Attribute1Name  	string             
	Attribute1Type  	string         
	Attribute1Values  	string          
	Attribute2Name  	string             
	Attribute2Type  	string         
	Attribute2Values  	string          
	Attribute3Name  	string             
	Attribute3Type  	string         
	Attribute3Values  	string          
	Attribute4Name  	string             
	Attribute4Type  	string         
	Attribute4Values  	string          
	Attribute5Name  	string             
	Attribute5Type  	string         
	Attribute5Values  	string          
	Attribute6Name  	string             
	Attribute6Type  	string         
	Attribute6Values  	string          
	Attribute7Name  	string             
	Attribute7Type  	string         
	Attribute7Values  	string          
	Attribute8Name  	string             
	Attribute8Type  	string         
	Attribute8Values  	string          
	Attribute9Name  	string             
	Attribute9Type  	string         
	Attribute9Values  	string          
	Attribute10Name  	string             
	Attribute10Type  	string         
	Attribute10Values  	string          
	Attributes  	[]AttributeSetLineModel
}

type BankAccounts struct {
	 AccountID        	 string     
	 Bank             	 string   
	 AccountName      	 string   
	 AccountNumber    	 string   
	 AccountCode      	 string
	 Currency         	 string
	 StatementBalance 	 float32  
	 BalanceInDear    	 float32 
	 InitialBalance   	 string 
}

type Brand struct {
	 ID  	 string           
	 Name  	 string             
}

type Carrier struct {
	 CarrierID   	 string      
	 Description 	 string    
}

type ChartofAccounts struct {
	 Code   	 string         
	 Name   	 string         
	 Type   	 string         
	 Status 	 string         
	 Description 	 string    
	 Class  	 string         
	 SystemAccount 	 string  
	 ForPayments 	 string    
	 DisplayName 	 string    
	 OldCode     	 string    
	 Bank        	 string    
	 BankAccountNumber 	 string 
	 BankAccountId 	 string    
	 Currency     	 string   
}

type Customer struct {
	 ID                  	 string      
	 Name                	 string    
	 Status              	 string    
	 Currency            	 string    
	 PaymentTerm         	 string    
	 AccountReceivable   	 string    
	 RevenueAccount      	 string    
	 TaxRule             	 string    
	 PriceTier           	 string       
	 Carrier             	 string    
	 SalesRepresentative 	 string    
	 Location            	 string    
	 Discount            	 int32      
	 Comments            	 string    
	 TaxNumber           	 int32      
	 CreditLimit         	 int32      
	 Tags                	 string    
	 AttributeSet        	 string    
	 AdditionalAttribute#	 string    
	 LastModifiedOn      	 string    
	 IsOnCreditHold      	 bool      
	 ProductPrices       	[]ProductPriceModel
	 Addresses           []SupplierAddressModel
	 Contacts            []SupplierContactModel
}

type FixedAssetTypes struct {
	 FixedAssetTypeID    	 string     
	 Name                	 string   
	 DepreciationMethod  	 string   
	 AveragingMethod     	 string   
	 Rate                	 float32  
	 EffectiveLife       	 float32  
	 AssetAccountCode    	 float32  
	 AccumulatedDepreciationAccountCode 	 string 
	 AssetAccountName    	 string   
	 AccumulatedDepreciationAccountName 	 string 
	 DepreciationExpenseAccountName 	 string   
}

type Location struct {
	 ID  	 string           
	 Name  	 string             
	 IsDefault  	 bool             
	 Deprecated  	 bool             
	 Bins  	 Array          
	 FixedAssetsLocation  	 bool             
	 ParentID  	 string             
	 ReferenceCount  	 int32            
	 AddressLine1  	 string             
	 AddressLine2  	 string             
	 AddressCitySuburb  	 string             
	 AddressStateProvince  	 string             
	 AddressZipPostCode  	 string             
	 AddressCountry  	 string             
	 PickZones  	 string             
	 IsShopfloor  	 bool             
	 IsCoMan  	 bool             
	 IsStaging  	 bool        
}

type ME struct {
	 Company  	 string           
	 Currency  	 string             
	 TimeZone  	 string             
	 DefaultWeightUnits  	 string             
	 DefaultDimensionsUnits  	 string             
	 LockDate  	 Date             
	 OpeningBalanceDate  	 Date             
	 TaxCalculationMethod  	 string             
	 DefaultSaleTaxRuleId  	 string             
	 DefaultSaleTaxRuleName  	 string             
	 Maximumfloat32PlacesInQuantity  	 string             
	 ApplyCustomerDiscountsAfterOtherDiscounts  	 bool             
	 DiscountRule  	 string             
	 AutomaticallyApplyDiscounts  	 bool             
	RoundingTable	[]RoundingTableModel
}

type RoundingTableModel struct {
	 RangeTo              	 float32      
	 RoundToNearest     	 float32      
	 AdjustmentRule        	 string        
	 AdjustmentValue        	 string        
}

type MeAddress struct {
	 AddressID       	 string      
	 Line1           	 string    
	 Line2           	 string    
	 CitySuburb      	 string    
	 StateProvince   	 string    
	 ZipPostCode     	 string    
	 Country         	 string    
	 Type            	 string    
	 DefaultForType  	 bool      
}

type MeContact struct {
	 ContactID       	 string      
	 Name            	 string    
	 Phone           	 string    
	 Fax             	 string    
	 Email           	 string    
	 Website         	 string    
	 Comment         	 string    
	 Type            	 string    
	 DefaultForType  	 bool      
}

type PaymentTerm struct {
	 ID          	 string      
	 Name        	 string    
	 Duration    	 int32      
	 Method      	 string    
	 IsActive    	 bool      
	 IsDefault   	 bool      
}

type PriceTier struct {
	 Code  	 int32           
	 Name  	 string             
}

type Product struct {
	 ID  	 string           
	 SKU  	 string             
	 Name  	 string             
	 Category  	 string             
	 Brand  	 string             
	 Type  	 string             
	 CostingMethod  	 string             
	 DropShipMode  	 string             
	 DefaultLocation  	 string             
	 Length  	 float32              
	 Width  	 float32              
	 Height  	 float32              
	 Weight  	 float32              
	 CartonLength  	 float32              
	 CartonWidth  	 float32              
	 CartonHeight  	 float32              
	 CartonQuantity  	 float32              
	 CartonInnerQuantity  	 float32              
	 UOM  	 string             
	 WeightUnits  	 string             
	 DimensionsUnits  	 string             
	 Barcode  	 string             
	 MinimumBeforeReorder  	 float32          
	 ReorderQuantity  	 float32             
	 PriceTier1  	 float32             
	 PriceTier2  	 float32             
	 PriceTier3  	 float32             
	 PriceTier4  	 float32             
	 PriceTier5  	 float32             
	 PriceTier6  	 float32             
	 PriceTier7  	 float32             
	 PriceTier8  	 float32             
	 PriceTier9  	 float32             
	 PriceTier10  	 float32             
	 PriceTiers  	 []PriceTierModel
	 AverageCost  	 float32             
	 ShortDescription  	 string             
	 Description  	 string             
	 InternalNote  	 string             
	 AdditionalAttribute1  	 string             
	 AdditionalAttribute2  	 string             
	 AdditionalAttribute3  	 string             
	 AdditionalAttribute4  	 string             
	 AdditionalAttribute5  	 string             
	 AdditionalAttribute6  	 string             
	 AdditionalAttribute7  	 string             
	 AdditionalAttribute8  	 string             
	 AdditionalAttribute9  	 string             
	 AdditionalAttribute10  	 string             
	 AttributeSet  	 string             
	 DiscountRule  	 string             
	 Tags  	 string             
	 Status  	 string             
	 StockLocator  	 string             
	 COGSAccount  	 string             
	 RevenueAccount  	 string             
	 ExpenseAccount  	 string             
	 InventoryAccount  	 string             
	 PurchaseTaxRule  	 string             
	 SaleTaxRule  	 string             
	 LastModifiedOn  	 string             
	 Sellable  	 bool             
	 PickZones  	 string             
	 BillOfMaterial  	 bool             
	 AutoAssembly  	 bool             
	 AutoDisassembly  	 bool             
	 QuantityToProduce  	 float32             
	 AssemblyInstructionURL  	 string             
	 AssemblyCostEstimationMethod  	 string             
	 BOMType  	 string             
	 HSCode  	 string             
	 CountryOfOrigin  	 string    
	 CountryOfOriginCode  	 string    
	 Suppliers  	 <a href="#ProductSupplierModel"><b>[] Product Supplier Model</b></a>              
	 ReorderLevels  	 <a href="#ReorderLevelModel"><b>[] Reorder Level Model</b></a>              	         	          		
	 BillOfMaterialsProducts  	 <a href="#BillOfMaterialProductModel"><b>[] Bill Of Material Product Model</b></a>              	         	          		
	 BillOfMaterialsServices  	 <a href="#BillOfMaterialServiceModel"><b>[] Bill Of Material Service Model</b></a>              	         	          		
	 Movements  	 <a href="#ProductMovementModel"><b>[] Product Movement Model</b></a>              	         	          		
	 Attachments  	 <a href="#AttachmentLineModel"><b>[] Attachment Line Model</b></a>              	         	          		
	 CustomPrices       	<a href="#ProductPriceModel"><b>[] Custom Price Model</b></a>			 Customer specific Product Prices.         	
}

### Available Fields for Product Availability
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 ID                  	 string      	           	           	 Product ID.                                                   	
	 SKU                 	 string    	 50        	           	 SKU of Product.                                               	
	 Name                	 string    	 256       	           	 Name of Product.                                              	
	 Barcode             	 string    	 256       	           	 Barcode of Product.                                           	
	 Location            	 string    	 256       	           	 Name of product location.                                     	
	 Bin                 	 string    	 50        	           	 Binof product location.                                       	
	 Batch               	 string    	           	           	 Batch.                                                        	
	 ExpiryDate          	 string  	           	           	 Expiry Date.                                                  	
	 OnHand              	 float32   	           	           	 On Hand.                                                      	
	 Allocated           	 float32   	           	           	 Allocated.                                                    	
	 Available           	 float32   	           	           	 Available.                                                    	
	 OnOrder             	 float32   	           	           	 OnOrder.                                                      	
	 StockOnHand         	 float32   	           	           	 StockOnHand.                                                  	
	 InTransit           	 float32   	           	           	 In Transit.                                                  	

### Available Fields for Product Category
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 ID  	 string           	         	          	Required for PUT, Ignored for POST operations          	
	 Name  	 string             	 50        	 Yes         	Category name for product          	

### Available Fields for Product Family
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 Products  	 <a href="#ProductFamilyProductLineModel"><b>[] Product Family Product Line Model</b></a>              	         	          	List of Products included in the Family.	
	 ID  	 string           	         	          	Unique ID. Ignored by POST action. Required for PUT action          	
	 SKU  	 string             	 45        	 Yes         	Must be unique. Product Family SKU, used to generate Product SKUs	
	 Name  	 string             	 256        	 Yes         	Product Family name, used to generate Product names	
	 Category  	 string             	 256        	 Yes         	Category name 	
	 Brand  	 string             	 256        	          	Brand name	
	 CostingMethod  	 string             	         	  Yes        	Valid values: FIFO, Special - Batch, Special - Serial Number, FIFO - Serial Number, FIFO - Batch, FEFO - Batch, FEFO - Serial Number	
	 DefaultLocation  	 string             	 256        	   Yes       	Name of the default Location for product family. (Location with this name must exist in Reference Books>Locations and Bins)	
	 UOM  	 string             	 0      	  Yes        	Unit of Measure (Unit of measure with this name must exist in Reference Books > Units of Measure)	
	 MinimumBeforeReorder  	 float32  (Nullable)          	         	          	Minimum available product quantity for this product to appear on Reorder report/forms. Defaults to 0	
	 ReorderQuantity  	 float32  (Nullable)          	         	          	Default quantity to put into purchase order when reordering this product. Not applicable to reorder backordered form. Defaults to 0	
	 PriceTier1  	 float32          	         	          	Product Sale Price used when Price Tier 1 is selected in Sale Task. Up to 4 decimal places	
	 PriceTier2  	 float32          	         	          	Product Sale Price used when Price Tier 2 is selected in Sale Task. Up to 4 decimal places	
	 PriceTier3  	 float32          	         	          	Product Sale Price used when Price Tier 3 is selected in Sale Task. Up to 4 decimal places	
	 PriceTier4  	 float32          	         	          	Product Sale Price used when Price Tier 4 is selected in Sale Task. Up to 4 decimal places	
	 PriceTier5  	 float32          	         	          	Product Sale Price used when Price Tier 5 is selected in Sale Task. Up to 4 decimal places	
	 PriceTier6  	 float32          	         	          	Product Sale Price used when Price Tier 6 is selected in Sale Task. Up to 4 decimal places	
	 PriceTier7  	 float32          	         	          	Product Sale Price used when Price Tier 7 is selected in Sale Task. Up to 4 decimal places	
	 PriceTier8  	 float32          	         	          	Product Sale Price used when Price Tier 8 is selected in Sale Task. Up to 4 decimal places	
	 PriceTier9  	 float32          	         	          	Product Sale Price used when Price Tier 9 is selected in Sale Task. Up to 4 decimal places	
	 PriceTier10  	 float32          	         	          	Product Sale Price used when Price Tier 10 is selected in Sale Task. Up to 4 decimal places	
	 ShortDescription  	 string             	 500        	          	Multiline short description of the product	
	 Description  	 string             	 Unlimited        	          	Multiline description of the product family	
	 AttributeSet  	 string             	 50        	          	Name of the product’s additional attribute set.	
	 DiscountRule  	 string             	 128        	          	Valid discount name (discount with this name must exist in Product Discounts reference book and should be active)	
	 Tags  	 string             	 256        	          	Comma delimited list of custom tags	
	 COGSAccount  	 string             	 50        	          	Account code of an active EXPENSE class account from Chart Of Accounts. 	
	 RevenueAccount  	 string             	 50        	          	Account code of an active REVENUE class account from Chart Of Accounts.	
	 InventoryAccount  	 string             	 50        	          	Account code of an active ASSET class account from Chart Of Accounts with Type not equal to FIXED or BANK. 	
	 PurchaseTaxRule  	 string             	 50        	          	Tax rule used for this product in Purchases. Tax Rule with this name must exist in Reference Books> Taxation Rules and should be Active. Purchases flag needs to be set to true in API.	
	 SaleTaxRule  	 string             	 50        	          	Tax rule used for this product in Sales. Tax Rule with this name must exist in Reference Books > Taxation Rules and should be Active. Sales flag needs to be set to true in API.	
	 DropShipMode  	 string             	         	          	One of these values: No Drop Ship, Optional Drop Ship, Always Drop Ship. Default value is No Drop Ship.	
	 Option1Name  	 string          	  50       	 Yes         	Family Option 1 name. I.e. Size, Colour, etc. Defaults to Option 1	
	 Option2Name  	 string          	  50       	          	Family Option 2 name. I.e. Size, Colour, etc. Defaults to Option 1	
	 Option3Name  	 string          	  50       	          	Family Option 3 name. I.e. Size, Colour, etc. Defaults to Option 1	
	 Option1Values  	 string          	         	          	Comma-delimited list of all possible values for Option 1. Combination of unique values of Option 1 Value for product family products. Read-only.	
	 Option2Values  	 string          	         	          	Comma-delimited list of all possible values for Option 2. Combination of unique values of Option 1 Value for product family products. Read-only.	
	 Option3Values  	 string          	         	          	Comma-delimited list of all possible values for Option 3. Combination of unique values of Option 1 Value for product family products. Read-only.	
	 LastModifiedOn  	 string  (Nullable)        	         	          	UTC Time	
	 Attachments  	 <a href="#AttachmentLineModel"><b>[] Attachment Line Model</b></a>              	         	          		
	 HSCode  	 string             	 200      	  	Specifies the HS Codes used in the process of product export. It is a numeric value of minimum 6 digits.	
	 CountryOfOrigin  	 string    	   	  	Specifies the country of manufacture, production, design, or brand origin where a product comes from.	
	 CountryOfOriginCode  	 string    	   	  	Specifies the country code of manufacture, production, design, or brand origin where a product comes from.	

### Available Fields for FactoryCalendar
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 Year  	 Short           	         	 Yes 	         	
	 WeekStart  	 string             	 50        	 Yes* 	 Required for POST. Possible values are <a href="#DayOfWeekList"><b>values</b></a>.	
	 FactoryCalendarDays  	 Array             	         	    Yes*  	 Required for POST. There should be data for each days in a week (7 items). 	
	 FactoryCalendarSpecialDays  	 Array             	         	          	 Data will be deleted if received collection is empty. 	

### Available Fields for FactoryCalendarDay
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 DayOfWeek  	 string           	         	 Yes 	    Possible values are <a href="#DayOfWeekList"><b>values</b></a>.     	
	 IsWeekend  	 bool             	         	  	 Default false 	
	 StartTime  	 Short             	         	    Yes  	 In minutes. Should be in the range 0 - 1440 (from 0 minutes to 24h*60m). 	
	 EndTime  	 Short             	         	    Yes  	 In minutes. Should be in the range 0 - 1440 (from 0 minutes to 24h*60m). 	
	 BreakStartTime  	 Short             	         	      	 In minutes. Should be in the range 0 - 1440 (from 0 minutes to 24h*60m). 	
	 BreakEndTime  	 Short             	         	      	 In minutes. Should be in the range 0 - 1440 (from 0 minutes to 24h*60m). 	
	 Capacity  	 int32            	         	          	 Readonly. 	

### Available Fields for FactoryCalendarSpecialDay
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 Date  	 string           	         	 Yes 	    Must fall within the current factory calendar year.    	
	 IsWeekend  	 bool             	         	  	 Default false 	
	 IsHoliday  	 bool             	         	  	 Default false 	
	 StartTime  	 Short             	         	      	 In minutes. Should be in the range 0 - 1440 (from 0 minutes to 24h*60m). 	
	 EndTime  	 Short             	         	      	 In minutes. Should be in the range 0 - 1440 (from 0 minutes to 24h*60m). 	
	 BreakStartTime  	 Short             	         	      	 In minutes. Should be in the range 0 - 1440 (from 0 minutes to 24h*60m). 	
	 BreakEndTime  	 Short             	         	      	 In minutes. Should be in the range 0 - 1440 (from 0 minutes to 24h*60m). 	
	 Capacity  	 int32            	         	          	 Readonly. 	
	 Comment  	 string             	  512       	          	  	

### Available Fields for Production BOM <a name="ProductionBOM" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	BOMID	string		Yes when updating	Identifier of a Production BOM.	
	OutputQuantity	float32		Yes		
	BufferPercent	float32		Yes		
	InstructionUrl	string				
	IgnoreCumulativeLeadTime	bool			Default = false	
	ComponentProductionLeadTime	int32			Required when IgnoreCumulativeLeadTime is set to 'true'. Default = 0.	
	Version	int32		Yes		
	Name	string	256	Yes		
	IsDefault	bool		Yes		
	IsTracing	bool		No		
	DeliveryToID	string			Location: Bin / Location.	
	DeliveryToName	string	256			
	IssueMethod	string			Readonly, can be Manual = 1, DerivedFromProductFamily = 3, ExternalApi = 4.	
	Operations	<a href="#ProductionBOMOperation"><b>[] Operation Model</b></a>				
	IssueMethodComponent	string			Can be Manual = 1, Backflush = 2	
	IssueMethodParameter	string			Ignored when IssueMethodComponent = Manual. Possible values: BackflushForBOM = 1, BackflushForOperation = 2, BackflushForComponent = 3 	
	MinQuantity        	 float32    	           	    No       	 Minimum Quantity to Produce Per Production Order defines the minimum quantity threshold to initiate the Production Order creation. If this quantity is not met, Production Order is not created automatically based on the Sales Order.	
	MaxQuantity        	 float32    	           	    No       	 Maximum Quantity to Produce Per Production Order defines the maximum quantity threshold to split the Production Order while creating. If this quantity is met, Production Order is split into several Production Orders based on Max Quantity To Produce Per Production Order.	
	DeviationPercent        	 float32    	           	    No       	 Order Quantity Deviation % defines the deviation from Min/Max Quantity to Produce Per Production Order. If Order Quantity Deviation % is set to 0, the Min/Max Quantity to Produce Per Production Order is a fixed quantity.	
	RunSize        	 float32    	           	    No       	 Run Size defines the multiplicity factor for Production Run.	

### Available Fields for Production BOM Operation Type <a name="ProductionBOMOperationType" />
	Operation Type	
	:-------	
	Manufacturing	
	Setup	
	QA	
	CoManufacturing	

### Available Fields for Production BOM Resource cost calculation type <a name="ResourceCostCalculationType" />
	Resource Cost Calculation Type	
	:-------	
	CostPerUnitTime	
	CostPerFinishedProduct	

### Available Fields for Production BOM Operation <a name="ProductionBOMOperation" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	OperationID	string		Yes when updating	Identifier of a Production BOM Operation.	
	Order	int32	Not zero	Yes		
	Name	string	200	Yes		
	CycleTime	int32	Not zero	Yes	Cycle Time should be a positive value. Specified in seconds.	
	UnitsPerCycle	float32		Yes		
	WorkCenterID	string		Yes		
	WorkCenterName	string	256			
	WorkCenterCoManProcurementType	string			Readonly, can be Transfer or BuySell, is returned for Co-man workcenters only.	
	OperationType	string		Yes	Available <a href="#ProductionBOMOperationType"><b>values</b></a>.	
	DeliveryToID	string				
	DeliveryToName	string	256			
	IssueMethod	string			Readonly, can be Manual = 1, Backflush = 2, DerivedFromProductFamily = 3, ExternalApi = 4.	
	IsDropShip	bool		Yes		
	Resources	<a href="#ProductionBOMResource"><b>[] Resource Model</b></a>				
	Components	<a href="#ProductionBOMComponent"><b>[] Component Model</b></a>				
	Attachments	<a href="#ProductionBOMAttachment"><b>[] Attachment Model</b></a>				
	Notes	<a href="#ProductionBOMNote"><b>[] Note Model</b></a>				
	OperationLinks	<a href="#ProductionBOMOperationLink"><b>[] Operation Link Model</b></a>				
	InputProducts	<a href="#ProductionBOMOperationProduct"><b>[] Input Product Model</b></a>				
	OutputProducts	<a href="#ProductionBOMOperationProduct"><b>[] Output Product Model</b></a>				
	FinishedProducts	<a href="#ProductionBOMOperationProduct"><b>[] Finished Product Model</b></a>				
	IsBackflush	bool		No	 Used only when IssueMethodComponent = Backflush and IssueMethodParameter = BackflushForOperation	

### Available Fields for Production BOM Resource <a name="ProductionBOMResource" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	BOMResourceID	string		Yes when updating	Identifier of a Production BOM Resource.	
	ResourceID	string		Yes		
	ResourceName	string			Readonly	
	ResourceCode	string			Readonly	
	Quantity	float32	Not zero	Yes		
	Cost	float32			Readonly	
	CycleTime	int32			Readonly	
	Position	int32		Yes		
	CostCalculationType	string		Yes	Available <a href="#ResourceCostCalculationType"><b>values</b></a>.	

### Available Fields for Production BOM Component <a name="ProductionBOMComponent" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	BOMComponentID	string		Yes when updating	Identifier of a Production BOM Resource.	
	ProductID	string		Yes when ProductSKU is not set		
	ProductSKU	string		Yes when ProductID is not set		
	ProductName	string			Readonly	
	Quantity	float32	Not zero	Yes	Component Quantity should be a positive value.	
	WastageQty	float32	Not zero	0+	Can be specified insead of WastagePercent.	
	WastagePercent	float32	0 - 100		Can be specified insead of WastageQty.	
	Cost	float32			Readonly	
	CostingMethod	string			Readonly	
	Unit	string			Readonly	
	Position	int32		Yes		
	SalePriceTier	int32	1 - 10		Used only in Buy/Sell co-man operation.	
	IsBackflush	bool		No	 Used only when IssueMethodComponent = Backflush and IssueMethodParameter = BackflushForComponent	

### Available Fields for Production BOM Attachment <a name="ProductionBOMAttachment" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	AttachmentID	string		Yes when updating	Identifier of a Production BOM Attachment.	
	FileName	string	256			
	Date	string			Readonly	
	ContentType	string		Yes		
	Content	string		Yes	Base64 string of binary file content.	
	Position	int32		Yes		

### Available Fields for Production BOM Note <a name="ProductionBOMNote" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	NoteID	string		Yes when updating	Identifier of a Production BOM Note.	
	Note	string	4000		Yes	
	Position	int32		Yes		

### Available Fields for Production BOM Operation Link <a name="ProductionBOMOperationLink" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	string		Yes when updating	Identifier of a Production BOM Operation Link.	
	RelatedOperationID	string		Yes when RelatedOperationName is not set.		
	RelatedOperationName	string	200	Yes when RelatedOperationID is not set.		
	RelationType	int32	0+	Yes	Possible value: 1	

### Available Fields for Production BOM InputProducts, OutputProducts, FinishedProducts <a name="ProductionBOMOperationProduct" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	string		Yes when updating	Identifier of a Production BOM Operation Product.	
	ProductID	string		Yes when ProductSKU is not set.		
	ProductSKU	string		Yes when ProductID is not set.		
	ProductName	string			Readonly	
	CostCalculationType	string		Yes	Can be None. If BOM contains several finished products, should contain SalesValue or Nrv.	
	PriceTier	int32	1 - 10			
	Ratio	float32	0 - 100			
	OutputQuantity	float32	0+	Yes		
	Position	int32		Yes		
	CostingMethod	string			Readonly.	
	Unit	string			Readonly	
	AverageCost	float32			Readonly	
	WastageCost	int32				
	DeliveryTo	string			DeliveryTo location identifier. 	
	DeliveryToName	string	256		DeliveryTo location name.	

### Available Fields for Production BOM
See <a href="#ProductionBOM"><b>Production BOM Model</b></a>. The only difference is that fields 'Version', 'Name', 'IsDefault' are currently readonly.

### Available Fields for Production BOM Operation
See <a href="#ProductionBOMOperation"><b>Production BOM Operation Model</b></a>
Additional fields:
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	VariationComponents	<a href="#ProductionBOMVariationComponent"><b>[] Variation Component Model</b></a>				

### Available Fields for Production BOM Variation Component <a name="ProductionBOMVariationComponent" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	BOMVariationComponentID	string		Yes when updating	Identifier of a Production BOM Variation Component.	
	ProductFamilyID	string		Yes when ProductFamilySKU is not set		
	ProductFamilySKU	string		Yes when ProductFamilyID is not set		
	ProductFamilyName	string			Readonly	
	QuantitySettingsJSON	string		Yes		
	MapVariationsJSON	string		Yes		
	QuantityOptionName	string	50	Yes		
	CostingMethod	string			Readonly	
	Unit	string			Readonly	
	Position	int32		Yes		
	SalePriceTier	int32	1 - 10			

### Available Fields for Production BOM Resource
See <a href="#ProductionBOMResource"><b>Production BOM Resource Model</b></a>

### Available Fields for Production BOM Component
See <a href="#ProductionBOMComponent"><b>Production BOM Component Model</b></a>

### Available Fields for Production BOM Attachment
See <a href="#ProductionBOMAttachment"><b>Production BOM Attachment Model</b></a>

### Available Fields for Production BOM Note
See <a href="#ProductionBOMNote"><b>Production BOM Note Model</b></a>

### Available Fields for Production BOM Operation Link
See <a href="#ProductionBOMOperationLink"><b>Production BOM Operation Link Model</b></a>

### Available Fields for Production BOM InputProducts, OutputProducts, FinishedProducts
See <a href="#ProductionBOMOperationProduct"><b>Production BOM Operation Product Model</b></a>

### Available Fields for Production Order 
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 ProductionOrderID   	 string       	        	    Yes    	 Unique identifier of Production Order. Required for Update action.             	
	 ProductID            	 string    	           	  Yes   	 Unique identifier of Product.                                    	
	 ProductSKU         	 string    	           	        	 Product SKU.                                                 	
	 ProductName         	 string    	           	        	 Product Name.                          	
	 OrderNumber         	 string    	           	        	 Production Order Number.                                	
	 LocationID          	 string    	           	  Yes   	 Production Order Location ID.                                                	
	 LocationName        	 string    	           	           	 Production Order Location Name.                                                   	
	 CostingMethod        	 string    	           	           	 CostingMethod of Product (FIFO, FEFO etc.).                                                   	
	 WarehouseName        	 string    	           	           	 Readonly.                                                   	
	 Unit        	 string    	           	           	 Readonly.                                                   	
	 OrderStatus      	 string    	           	           	 Production Order execution status (DRAFT, PLANNED, RELEASED, IN PROGRESS, COMPLETED, VOIDED). Readonly.                                         	
	 Status             	 string    	           	           	 Production Order status (DRAFT, AUTHORIZED, RELEASED, VOIDED). Readonly.                                                 	
	 InstructionURL        	 string    	           	           	 Readonly.                                                   	
	 SourceName        	 string    	           	           	 Manual, Sale, SmartReorder.                                                   	
	 SourceTaskID	 string    	           	           	Obsolete. Unique identifier of Sale Task that produced the current Production Order record.                    	
	 SourceTaskNumber      	 string    	           	           	Obsolete. Number of Sale Task that produced the current Production Order record.                                    	
	 WIPAccountCode         	 string       	           	           	 Work In Progress Account Code.                                                 	
	 FinishedGoodsAccountCode        	 string    	           	           	 Finished Goods Account Code.                                            	
	 Quantity            	 float32    	           	           	 Production Order Quantity.                                                     	
	 BOMQuantity            	 float32    	           	           	 Production BOM origin Quantity. Readonly.                                                    	
	 CapacityCalculationType           	 string       	           	           	 Type of the capacity calculation method. Available values are FromStartForward and FromRequiredBackward.  	
	 StartDate            	 string       	           	           	 Production Order Planned Date. Required when CapacityCalculationType = FromStartForward.                   	
	 ReleaseDate            	 string       	           	           	 Production Order Release Date.                         	
	 RequiredByDate            	 string       	           	           	 Production Order Required By Date. Required when CapacityCalculationType = FromRequiredBackward.                         	
	 CompletionDate            	 string       	           	           	 Production Order Completion Date. Readonly.                        	
	 IsIgnoreLeadTime            	boolean			 Readonly.          	
	 WarehouseID            	string			 Readonly.          	
	 RetailID            	string			 Readonly.          	
	 Comments            	 string    	 2000       	           	                                                      	
	 ScheduleStart            	 string       	           	           	 Readonly.                        	
	 PlannedBy            	 string    	 2000       	           	 Readonly.                                                     	
	 ReleasedBy            	 string    	 2000       	           	 Readonly.                                                     	
	 OrderCycleTime            	 int32    	        	           	 In seconds.                                                    	
	 BOMVersion            	 int32    	        	           	 Readonly.                                                     	
	 BOMName            	 string    	        	           	 Readonly.                                                     	
	 Tags                	 string    	           	           	                                                  	
	CustomField1	string			Value of Production Order's additional attribute 1	
	CustomField2	string			Value of Production Order's additional attribute 2	
	CustomField3	string			Value of Production Order's additional attribute 3	
	CustomField4	string			Value of Production Order's additional attribute 4	
	CustomField5	string			Value of Production Order's additional attribute 5	
	CustomField6	string			Value of Production Order's additional attribute 6	
	CustomField7	string			Value of Production Order's additional attribute 7	
	CustomField8	string			Value of Production Order's additional attribute 8	
	CustomField9	string			Value of Production Order's additional attribute 9	
	CustomField10	string			Value of Production Order's additional attribute 10	
	 ProductionOrderOperations                	 ProductionOrderOperation[]    	           	           	                                                  	
	 ProductionRuns                	 ProductionRun[]    	           	           	                                                  	
	 ProductionOrderDeliveryTo                	 ProductionOrderDeliveryTo[]    	           	           	                                                  	
	IssueMethodComponent	string			Readonly. Can be Manual = 1, Backflush = 2.	
	IssueMethodParameter	string			Readonly. Ignored when IssueMethodComponent = Manual. Possible values: BackflushForBOM = 1, BackflushForOperation = 2, BackflushForComponent = 3 	
	SourceTasks	ProductionOrderSourceTask[] 			Readonly. Data of Sale Tasks that produced the current Production Order record. 	
	RunSize	float32		No	Readonly. Run Size specified in Production BOM 	

### Available Fields for Production Order Operation
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 OperationID   	 string       	        	    Yes*    	 Unique identifier of Production Order Operation. Required for Update action.                                             	
	 Order            	 int32    	   1+        	Yes	 Operation sequence number.	
	 Name         	 string    	  200 	  Yes   	 Product SKU.                                                 	
	 CycleTime         	 int32    	  0+         	   Yes	 In seconds.                          	
	 UnitsPerCycle         	 float32    	  0+         	  Yes   	                                 	
	 TotalCycleTime          	 int32    	   1+        	  Yes   	                                                 	
	 TotalUnitsPerCycle        	 float32    	  0+         	           	                                                 	
	 OperationType        	 string    	           	           	 Manufacturing, Setup, QA.                                          	
	 WorkCenterID        	 string    	           	           	                                                    	
	 WorkCenterName        	 string    	           	           	                                                    	
	 WorkCenterCode      	 string    	           	           	                                          	
	 WorkCenterCoManProcurementType             	 string    	           	           	 Only for CoMan workcenters. Readonly. Transfer or Buysell.	
	 ComponentLocationID        	 string    	           	           	 Readonly.                                                   	
	 IsDropShip        	 boolean    	           	           	 For CoMan operations only.                                                   	
	 Attachments	 <a href="#ProductionOrderOperationAttachment"><b>[] Attachment Model</b></a>    	           	           	                     	
	 Components	 <a href="#ProductionOrderComponent"><b>[] Component Model</b></a>    	           	           	                     	
	 Notes	 <a href="#ProductionOrderOperationNote"><b>[] Note Model</b></a>    	           	           	                     	
	 Resources	 <a href="#ProductionOrderResource"><b>[] Resource Model</b></a>    	           	           	                     	
	 OperationLinks	 <a href="#ProductionOrderOperationLink"><b>[] Link Model</b></a>    	           	           	                     	
	 InputProducts	 <a href="#ProductionOrderOperationProduct"><b>[] OperationProduct Model</b></a>    	           	           	                     	
	 OutputProducts	 <a href="#ProductionOrderOperationProduct"><b>[] OperationProduct Model</b></a>    	           	           	                     	
	 FinishedProducts	 <a href="#ProductionOrderOperationProduct"><b>[] OperationProduct Model</b></a>    	           	           	                     	
	 IsBackflush	 bool  		 	 No 	 Readonly	

### Available Fields for Production Order Operation Attachment <a name="ProductionOrderOperationAttachment" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 AttachmentID   	 string       	        	    Yes*    	 Unique identifier of Production Order Operation Attachment. Required for Update action.                                             	
	 FileName            	 string    	   256        			
	 Date         	 string    	   	     	 Readonly	
	 ContentType         	 string    	  256         	   Yes	 	
	 Position         	 int32    	           	  Yes   		
	 Content          	 string    	           	  Yes*   	Base64 string of binary file content. Required when attachment is being created.	

### Available Fields for Production Order Operation Component <a name="ProductionOrderComponent" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 OrderComponentID   	 string       	        	    Yes*    	 Unique identifier of Production Order Operation Component. Required for Update action.                                             	
	 ProductID            	 string    	           	Yes* 	Required when ProductSKU is not specified.	
	 ProductSKU         	 string    	   	Yes*	Required when ProductID is not specified.	
	 ProductName         	 string    	           	   	 Readonly.	
	 Position         	 int32    	           	  Yes   		
	 Available          	 float32    	           	     	Readonly.	
	 Quantity          	 float32    	           	 Yes 		
	 TotalQuantity          	 float32    	    0+       	  		
	 WastageQty          	 float32    	    0+       	  	Can be specified insead of WastagePercent.	
	 WastagePercent          	 float32    	    0+       	  	Can be specified insead of WastageQty.	
	 Cost          	 float32    	    0+       	  	Readonly.	
	 TotalCost          	 float32    	    0+       	  	Readonly.	
	 ProductCost          	 float32    	    0+       	  	Readonly.	
	 CostingMethod          	 string    	           	  	Taken from Product, one of FIFO, FEFO - Serial Number etc.	
	 Unit          	 string    	           	  	Readonly.	
	 ProductType          	 string    	           	  	Readonly.	
	 SalePriceTier          	 int32    	           	  	Used only in Buy/Sell co-man operation. 	
	 IsBackflush	 bool  		 	 No 	 Readonly	

### Available Fields for Production Order Operation Note <a name="ProductionOrderOperationNote" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 NoteID   	 string       	        	    Yes*    	 Unique identifier of Production Order Operation Note. Required for Update action.                                             	
	 Note            	 string    	   4000        			
	 Position         	 int32    	           	  Yes   		

### Available Fields for Production Order Operation Resource <a name="ProductionOrderResource" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 OrderResourceID   	 string       	        	    Yes*    	 Unique identifier of Production Order Operation Resource. Required for Update action.                                             	
	 ResourceID            	 string    	           	Yes* 	Required when ResourceCode is not specified	
	 ResourceCode         	 string    	   	Yes*	Required when ResourceID is not specified	
	 ResourceName         	 string    	           	   	 Readonly.	
	 Position         	 int32    	           	  Yes   		
	 Quantity          	 float32    	    0+    	 Yes 		
	 Cost          	 float32    	          	  	Readonly.	
	 TotalCost          	 float32    	        	  	Readonly.	
	 ResourceCost          	 float32    	    0+       	  	Readonly.	
	 CycleTime          	 int32    	           	  	In seconds.	
	CostCalculationType	string		Yes	Available <a href="#ResourceCostCalculationType"><b>values</b></a>.	

### Available Fields for Production Order Operation Link. This is a readonly entity. <a name="ProductionOrderOperationLink" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 ID   	 string       	        	        	 Unique identifier of Production Order Operation Link. Readonly
	 RelatedOperationID            	 string    	          		Related Production Order Operation Identifier	
	 Position         	 int32    	           	  Yes   		

### Available Fields for Production Order Operation Product. This is a readonly entity (only deletion is allowed). <a name="ProductionOrderOperationProduct" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 ID   	 string       	        	        	 Unique identifier of Production Order Operation Product.
	 ProductID            	 string    	           	 		
	 ProductSKU         	 string    	   			
	 ProductName         	 string    	           	   	 	
	 Position         	 int32    	           	     		
	 CostCalculationType          	 string    	           	     	None, SalesValue, Nrv	
	 PriceTier          	 int32    	           	  		
	 Ratio          	 float32    	           	  		
	 OutputQuantity          	 float32    	           	  		
	 TotalOutputQuantity          	 float32    	           	  		
	 Cost          	 float32    	           	  		
	 TotalCost          	 float32    	           	  		
	 AverageCost          	 float32    	           	  		
	 FixedCost          	 float32    	           	  		
	 CostingMethod          	 string    	           	  		
	 Unit          	 string    	           	  		
	 WastageCost          	 float32    	           	  		

### Available Fields for Production Order Source Task <a name="ProductionOrderSourceTask" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 SourceTaskID	 string    	           	           	 Unique identifier of Sale Task that produced the current Production Order record.                    	
	 SourceTaskNumber      	 string    	           	           	Number of Sale Task that produced the current Production Order record.                                    	

### Available Fields for Production Order List Item
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 TaskID              	 string      	           	        	 Unique identifier of returned record (Production Order or Production Order Run).                                           	
	 ProductionOrderID   	 string       	        	        	 Unique identifier of Production Order.                                             	
	 Type              	 string    	           	        	 Type of the returned record. Available values are O - Production Order and R - Production Order Run.	
	 ProductID            	 string    	           	        	 Unique identifier of Product.                                    	
	 ProductSKU         	 string    	           	        	 Product SKU.                                                 	
	 ProductName         	 string    	           	        	 Product Name.                          	
	 OrderNumber         	 string    	           	        	 Production Order Number.                                	
	 LocationID          	 string    	           	        	 Production Order Location ID.                                                	
	 LocationName        	 string    	           	           	 Production Order Location Name.                                                   	
	 Status             	 string    	           	           	 Production Order state (DRAFT, AUTHORIZED, RELEASED, VOIDED).                                                 	
	 OrderStatus      	 string    	           	           	 Production Order status (DRAFT, PLANNED, RELEASED, IN PROGRESS, COMPLETED, VOIDED).                                         	
	 Quantity            	 float32    	           	           	 Production Order Quantity.                                                     	
	 StartDate            	 string       	           	           	 Production Order Start Date.                         	
	 ReleaseDate            	 string       	           	           	 Production Order Release Date.                         	
	 RequiredByDate            	 string       	           	           	 Production Order Required By Date.                         	
	 CompletionDate            	 string       	           	           	 Production Order Completion Date.                         	
	 Comments            	 string    	 2000       	           	 Comments.                                                     	
	 CapacityCalculationType           	 string       	           	           	 Capacity calculation method. Available values are FromStartForward and FromRequiredBackward.                                                   	
	 WIPAccountCode         	 string       	           	           	 Work In Progress Account Code.                                                 	
	 Tags                	 string    	           	           	 Tags string.                                                  	
	 FinishedGoodsAccountCode        	 string    	           	           	 Finished Goods Account Code.                                            	
	 SourceTaskID	 string    	           	           	Obsolete. Unique identifier of Sale Task that produced the current Production Order record.                    	
	 SourceTaskNumber      	 string    	           	           	Obsolete. Number of Sale Task that produced the current Production Order record.                                    	
	 SourceTaskType           	int32			Obsolete. Type of Sale Task that produced the current Production Order record. Available values are 0 - Simple and 1 - Master  	
	 IsSourceTaskVoided            	boolean			 Obsolete. True if source Sale Task is voided.          	
	CustomField1	string			Value of Production Order's additional attribute 1	
	CustomField2	string			Value of Production Order's additional attribute 2	
	CustomField3	string			Value of Production Order's additional attribute 3	
	CustomField4	string			Value of Production Order's additional attribute 4	
	CustomField5	string			Value of Production Order's additional attribute 5	
	CustomField6	string			Value of Production Order's additional attribute 6	
	CustomField7	string			Value of Production Order's additional attribute 7	
	CustomField8	string			Value of Production Order's additional attribute 8	
	CustomField9	string			Value of Production Order's additional attribute 9	
	CustomField10	string			Value of Production Order's additional attribute 10	
	SourceTasks	ProductionOrderListSourceTask[] 			Readonly. Data of Sale Tasks that produced the current Production Order record. 	

### Available Fields for Production Order List Source Task <a name="ProductionOrderListSourceTask" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 SourceTaskID	 string    	           	           	 Unique identifier of Sale Task that produced the current Production Order record.                    	
	 SourceTaskNumber      	 string    	           	           	Number of Sale Task that produced the current Production Order record.                                    	
	 SourceTaskType           	int32			 Type of Sale Task that produced the current Production Order record. Available values are 0 - Simple and 1 - Master  	
	 IsSourceTaskVoided            	boolean			 True if source Sale Task is voided.          	

### Available Fields for Production Run
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 RunID              	 string      	           	        	 Unique identifier of Production Run.                                           	
	 Number   	 int32      	        	        	 Readonly. Number of Production Run.                                             	
	 Status              	 string    	           	        	 Readonly. Available <a href="#ProductionRunStatus"><b>values</b></a>. 	
	 Quantity            	 float32    	           	        	 Quantity to produce during the Run, value must be positive.                                    	
	 WIPAccount         	 string    	     50      	        	 Work in Progress account.                                                	
	 StartDate         	 string    	           	        	 Readonly, nullable.                          	
	 EndDate         	 string    	           	        	 Readonly, nullable. 	
	 DueDate          	 string    	           	        	 Readonly, nullable.                                               	
	 ReceivedDate        	 string    	           	           	 Readonly, nullable.                                                  	
	 ScheduleStart         	 string    	           	        	 Readonly, nullable.                          	
	 ScheduleDue         	 string    	           	        	 Readonly, nullable. 	
	 Operations             	 Array <a href="#ProductionRunOperation">items<a/>   	           	           	 Operation List of Production Run. Operations cannot be added or removed, some fields can be modified depending on operation's status.                                               	
	 Output                  	 Array <a href="#ProductionRunOutput">items<a/>    	           	           	 Contains all finished products. Finished products can only be added to this array when an operation with finished products completed or the final operation in the sequence is completed.                                      	
	 PendingOutput            	 Array <a href="#ProductionRunPendingOutput">items<a/>    	           	           	 Contains planned finished products. Products can be added here until all operations are finished. Products will be removed from here automatically by ProductID as they are added to the Output.                                                          	
	 ManualJournals            	 Array <a href="#ProductionRunManualJournal">items<a/>       	           	           	 Manual Journals of Production Run. Readonly for PUT request. Use special url to update Manual Journals
	 Traceability            	 Array <a href="#ProductionRunTraceability">items<a/>       	           	           	 Traceability of Production Run.
	IssueMethodParameter	string			Ignored when IssueMethodComponent = Manual. Possible values: BackflushForBOM = 1, BackflushForOperation = 2, BackflushForComponent = 3 	

### Available Fields for Production Run Operation <a name="ProductionRunOperation"/>
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 OperationID              	 string      	           	  Yes      	 Unique identifier of Production Run Operation.                                           	
	 Status   	 string       	        	        	 Readonly, Available <a href="#ProductionRunOperationStatus"><b>values</b></a>.                                            	
	 Order              	 int32   	           	        	 Readonly. 	
	 Name            	 string    	    512       	        	 Operation's name. Can be updated only when operation has PLANNED status. In other cases, it will be ignored (readonly).                                   	
	 PlannedTime         	 int32   	           	        	 Can be modified only for CoManufacturing operation with status PLANNED.                                          	
	 ActualTime         	 int32   	           	        	 Can be modified only when operation status is COMPLETED.          	
	 UnitsPerCycle         	 float32    	           	        	 Readonly. 	
	 StartDate          	 string    	           	        	 Nullable. Can be modified only when operation status is PLANNED.                                              	
	 EndDate          	 string    	           	        	 Nullable. Can be modified only when operation status is IN PROGRESS.                                               	
	 DueDate          	 string    	           	        	 Readonly, nullable.                                               	
	 OperationType        	 string    	           	           	 Readonly. Available <a href="#ProductionBOMOperationType"><b>values</b></a>.                                                	
	 WorkCenterID        	 string    	           	           	 Can be modified only when operation status is PLANNED.                                                 	
	 WorkCenterCode        	 string    	           	           	 Readonly.                                                  	
	 WorkCenterName        	 string    	           	           	 Readonly.                                                  	
	 SuspendReason        	 string    	           	           	 Readonly.                                                  	
	 CoManStatus        	 string    	           	           	 Readonly. General status for CoManufacturing operation                                                 	
	 Components             	 Array <a href="#ProductionRunOperationComponent">items<a/>    	           	           	 All components in the operation. Can be modified only when operation has not been completed.                                             	
	 Resources                  	 Array <a href="#ProductionRunOperationResource">items<a/>   	           	           	 All resources in the operation.   	
	 ResourceCosts            	 Array <a href="#ProductionRunOperationResourceCost">items<a/>   	           	           	 Resource costs for each resource. Can be updated for completed operations.                                                        	
	 FinishedProducts            	 Array <a href="#ProductionRunOperationProduct">items<a/>      	           	           	Readonly. List of Finished Products in operation. Can be updated via completing an operation. When a finished product is saved as an Output, only wastage products will be contained in this array.   	
	 OutputProducts            	 Array <a href="#ProductionRunOperationProduct">items<a/>      	           	           	 Readonly. List of Output Products in operation. Can be updated via completing an operation.    	
	 InputProducts            	 Array <a href="#ProductionRunOperationProduct">items<a/>      	           	           	 Readonly. 	
	 CoManTasks            	 Array <a href="#ProductionRunOperationCoManTask">items<a/>        	           	           	 Readonly. Contains all tasks initiated from starting the CoManufacturing operation.  	
	 CoManTaskLines            	 Array <a href="#ProductionRunOperationCoManLine">items<a/>         	           	           	 Readonly. Output for CoManufacturing operation.  	
	 Attachments            	 Array <a href="#ProductionRunOperationAttachment">items<a/>         	           	           	 List of attachments in operation. 	
	 Notes            	 Array <a href="#ProductionRunOperationNote">items<a/>         	           	           	 List of notes in operation. 	
	 IsBackflush        	 bool    	           	           	 It is used only when IssueMethodComponent = Backflush and IssueMethodParameter = BackflushForOperation	
	 ManualStartDate        	 Date	           	           	 Is used only when Backdated Transaction is enabled in General Settings. Can be updated only when operation has status = 'PLANNED'. This overrides date of written off components.	
	 ManualEndDate        	 Date	           	           	 Is used when Backdated Transaction is enabled in General Settings. Can be updated only when operation has status = 'IN PROGRESS'. This overrides the date of transaction for service products.	

### Available Fields for Production Run Operation Component <a name="ProductionRunOperationComponent"/>
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 RunComponentID              	 string      	           	        	Readonly.                                          	
	 ProductID   	 string       	        	        	   Identifier product, required if ProductCode is empty, has a higher priority over the ProductCode.                                    	
	 ProductCode              	 string    	    50       	        	 SKU product, required if ProductID is empty. 	
	 ProductName            	 string    	           	        	 Readonly.                                  	
	 Quantity         	 float32    	           	        	 Actual quantity.                                         	
	 ExpectedQuantity         	 float32    	           	        	 Readonly.         	
	 WastageQty         	 float32    	           	        	 Wastage quantity.	
	 WastagePercent          	 float32    	           	        	 Readonly.                                           	
	 BatchSN          	 string    	      50     	        	 Can be modified ony when operation status is IN PROGRESS.                                              	
	 ExpiryDate          	 string    	           	        	 Nullable.                                               	
	 UnitCost        	 float32    	           	           	 Readonly.                                               	
	 LocationID        	 string    	           	           	 Use as location identifier.                                                	
	 LocationName        	 string    	           	           	 Also can be used as location identifier. Using format: <b>LocationName:BinName</b>.                                              	
	 ProductCost        	 string    	           	           	 Readonly.                                                  	
	 Available             	 float32    	           	           	 Available quantity in stock.                                             	
	 CostingMethod                  	 string    	           	           	 Readonly. Product costing method.  	
	 Unit            	 string    	           	           	 Readonly. Product unit of measure.                                                          	
	 ReservedQuantity            	 float32       	           	           	Readonly. Quantity of reserved components. 	
	 IsReserved            	 bool       	           	           	 Set 'true' to reserve current component with Quantity > 0 and Quantity <= Available.   	
	 IsBackflush        	 bool    	           	           	 It is used only when IssueMethodComponent = Backflush and IssueMethodParameter = BackflushForComponent	

### Available Fields for Production Run Operation Resource <a name="ProductionRunOperationResource"/>
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 RunResourceID   	 string       	        	    Yes*    	 Unique identifier of Production Run Resource. Required for Update action.                                             	
	 ResourceID              	 string      	           	        	Identifier resource, required if ResourceCode is empty, has a higher priority over the ResourceCode                                          	
	 ResourceCode   	 string       	        	        	   Identifier resource, required if ResourceID is empty                                   	
	 ResourceName              	 string    	    50       	        	 Readonly.	
	 UnitCost            	 float32    	           	        	 Readonly                                  	
	 Quantity         	 float32    	           	        	 Actual quantity                                         	
	 Cost         	 float32    	           	        	 Readonly.         	
	 CycleTime         	 float32    	           	        	 Readonly	

### Available Fields for Production Run Operation Resource Cost <a name="ProductionRunOperationResourceCost"/>
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 RunCostID              	 string      	           	   Yes     		
	 ProductID   	 string       	        	   Yes     	   Identifier service product from Resource	
	 ProductName   	 string       	        	        	   Readonly.	
	 AccountName              	 string    	    50       	        	 Readonly.	
	 ExpenseAccount            	 string    	     50      	        	 Code account                                  	
	 Comments         	 string    	           	    256    	                                       	
	 Cost         	 float32    	           	   Yes     	         	

### Available Fields for Production Run Operation Product <a name="ProductionRunOperationProduct"/>
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 ProductID   	 string       	        	        	 Readonly. Identifier product	
	 ProductSKU              	 string    	    50       	        	Readonly.	
	 ProductName            	 string    	           	        	 Readonly.                                  	
	 Unit            	 string    	           	           	 Readonly. Product unit of measure.   
	 OutputQuantity         	 float32    	           	        	 Readonly.                                         	
	 ExpectedQuantity         	 float32    	           	        	 Readonly.         	
	 WastageQuantity         	 float32    	           	        	 Readonly.	
	 BatchSN          	 string    	      50     	        	 Readonly.                                             	
	 ExpiryDate          	 string    	           	        	Readonly. Nullable.                                               	
	 LocationID        	 string    	           	           	    Readonly. Nullable.                                        	
	 LocationName        	 string    	           	           	 Readonly.  	

### Available Fields for Production Run Operation CoMan Task <a name="ProductionRunOperationCoManTask"/>
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 TaskID   	 string       	        	        	 Readonly. Identifier task	
	 Type              	 string    	           	        	Readonly. Available values 'PurchaseTask', 'SaleTask', 'TransferTask'  	
	 TaskStatus            	 string    	           	        	 Readonly.                                  	
	 TaskNumber            	 string    	           	           	 Readonly.	   

### Available Fields for Production Run Operation CoMan Line <a name="ProductionRunOperationCoManLine"/>
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 TaskID   	 string       	        	        	 Readonly. Identifier task	
	 ProductID   	 string       	        	        	 Readonly. Identifier product	
	 ProductCode              	 string    	    50       	        	Readonly.	
	 ProductName            	 string    	           	        	 Readonly.                                  	
	 Unit            	 string    	           	           	 Readonly. Product unit of measure .   	
	 Quantity         	 float32    	           	        	 Readonly.                                         	
	 Cost         	 float32    	           	        	 Readonly.         	
	 BatchSN          	 string    	      50     	        	 Readonly.                                             	
	 ExpiryDate          	 string    	           	        	Readonly. Nullable.                                               	
	 LocationID        	 string    	           	           	    Readonly. Nullable.                                        	
	 LocationName        	 string    	           	           	 Readonly.  	
	 ReceivedDate          	 string    	           	        	Readonly. Nullable.                                               	

### Available Fields for Production Run Operation Attachment <a name="ProductionRunOperationAttachment" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 AttachmentID   	 string       	        	    Yes*    	 Unique identifier of Production Run Operation Attachment. Required for Update action.                                            	
	 FileName            	 string    	   256        			
	 Date         	 string    	   	     	 Readonly	
	 ContentType         	 string    	  256         	   	 Readonly. 	 
	 Position         	 int32    	           	  Yes   		
	 Content          	 string    	           	  Yes*   	Base64 string of binary file content. Required when attachment is being created.	

### Available Fields for Production Run Operation Note <a name="ProductionRunOperationNote" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 NoteID   	 string       	        	    Yes*    	 Unique identifier of Production Run Operation Note. Required for Update action.                                             	
	 Note            	 string    	   4000        			
	 Position         	 int32    	           	  Yes   		

### Available Fields for Production Run Pending Output <a name="ProductionRunPendingOutput" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 ProductID   	 string       	        	        	   Identifier product, required if ProductCode is empty, has a higher priority over the ProductCode.                                    	
	 ProductCode              	 string    	    50       	        	 SKU product, required if ProductID is empty. 	
	 ProductName            	 string    	           	        	 Readonly.                                  	
	 Unit            	 string    	           	           	 Readonly. Product unit of measure.   	
	 CostingMethod            	 string    	           	           	 Readonly.  	
	 BatchSN          	 string    	      50     	        	 Pending batch/serial number.                                            	
	 ExpiryDate          	 string    	           	        	Required for folowing product's CostingMethod: FEBATCH, FESN.                                         	

### Available Fields for Production Run Pending Output <a name="ProductionRunOutput" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 ProductID   	 string       	        	        	   Identifier product, required if ProductCode is empty, has a higher priority over the ProductCode.                                    	
	 ProductCode              	 string    	    50       	        	 SKU product, required if ProductID is empty. 	
	 ProductName            	 string    	           	        	 Readonly.                                  	
	 Unit            	 string    	           	           	 Readonly. Product unit of measure .   	
	 CostingMethod            	 string    	           	           	 Readonly.  	
	 BatchSN          	 string    	      50     	        	 Pending batch/serial number.                                            	
	 ExpiryDate          	 string    	           	        	Required for folowing product's CostingMethod: FEBATCH, FESN.                                         	
	 UnitCost          	 float32    	           	        	Readonly. Nullable.                                 	
	 Quantity         	 float32    	           	    Yes    	          	
	 WastageQuantity         	 float32    	           	    Yes    	 	
	 ReceivedDate          	 string    	           	    Yes    		
	 Received          	 boolean    	           	    Yes    		
	 LocationID        	 string    	           	           	 Location identifier.                                                	
	 LocationName        	 string    	           	           	 Also can be used as location identifier. Using format: <b>LocationName:BinName</b>.                                              	

### Available Fields for Production Run Manual Journal <a name="ProductionRunManualJournal" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 JournalID   	 string       	        	    Yes*    	 Unique identifier of Production Run Manual Journals. Required for Update action.                                             	
	 Reference   	 string       	   50     	        	          	
	 Amount              	 float32    	    50       	    Yes    	 	
	 Date            	 string    	           	   Yes     	 	
	 Debit            	 string    	      50     	     Yes      	 Account code.  	
	 Credit            	 string    	     50      	     Yes      	 Account code.  	
	 OperationID          	 string    	           	        	 Nullable, Operation identifier.                                          	
	 OperationName          	 string    	           	        	Readonly. 	

### Available Fields for Production Run Traceability <a name="ProductionRunTraceability" />
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	 TraceabilityID   	 string       	        	    Yes*    	 Unique identifier of Production Run Traceability. Required for Update action.                                             	
	 ProducedProductID   	 string       	       	  Yes      	 Identifier of Produced finished product         	
	 ProducedProductBatchSN              	 string    	    50       	    Yes    	 Batch/Serial Number of Produced finished product	
	 ProducedProductExpiryDate            	 string    	           	   No     	 Expiry Date of Produced finished product	
	 UsedProductID   	 string       	       	  Yes      	  Identifier of Used Component           	
	 UsedProductBatchSN              	 string    	    50       	    Yes    	Batch/Serial Number of Used Component  	
	 UsedProductExpiryDate            	 string    	           	   No     	 Expiry Date of Used Component 	

### Available Fields for Resource
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 ResourceID  	 string           	         	 Yes* 	   Required for PUT and DELETE, Ignored for POST operations       	
	 Code  	 string             	 50        	 Yes* 	 Resource Code, required for PUT if there is no ResourceID     	
	 Name  	 string             	 512        	 Yes* 	 Resource name, required for POST     	
	 ResourceType  	 string             	         	    Yes      	 Possible values are <a href="#ResourceTypeList"><b>values</b></a>      	
	 CycleDuration  	 int32            	         	   Yes       	   In seconds, should be between 1 and 7776000        	
	 IsActive  	 bool          	         	          	 Default false           	
	 IsInfinite  	 bool          	         	          	 Default false           	
	 IsAllowAllCapacityUsage  	 bool          	         	          	 Default false           	
	 IsAvailableOnHolidays  	 bool          	         	          	 Default false           	
	 IsAvailableOnWeekends  	 bool          	         	          	 Default false           	
	 Tags  	 string          	    2000     	          	         	Comma separated	
	 ResourceCapacities  	 Array             	         	          	 	
	 ResourceCosts  	 Array             	         	          	 	
	 ResourceRemarks  	 Array             	         	          	 	
	 ResourceAttachments  	 Array             	         	          	 	

### Available Fields for ResourceCapacity
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 LocationID  	 string           	         	 Yes 	       	
	 LocationName  	 string             	 50        	  Yes*  	 Required if there is no LocationID 	
	 ResourceQuantity  	 int32            	         	  	 Readonly    	
	 NonOperationalFrom  	 string             	         	          	 Nulluble      	
	 NonOperationalTo  	 string             	         	          	   Nulluble        	
	 CustomWorkingDays  	 Array             	         	          	 	
	 ResourceUnits  	 Array             	         	          	 	

### Available Fields for CustomWorkingDay
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 DayOfWeek  	 string           	         	 Yes 	   Possible values are <a href="#DayOfWeekList"><b>values</b></a>   	
	 StartTime  	 Short             	         	  Yes*  	 Range 0 - 1440 	
	 EndTime  	 Short             	         	  Yes*  	 Range 0 - 1440 	
	 BreakStartTime  	 Short             	         	    	 Range 0 - 1440 	
	 BreakEndTime  	 Short             	         	    	 Range 0 - 1440 	
	Capacity	int32	 	 	 Readonly  	

### Available Fields for ResourceUnit
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 Name  	 string           	    256     	 Yes 	 For resource with type = Labor, it must be Email of the registered user. Otherwise any name  	
	 NonOperationalFrom  	 string             	         	   	 Nullable	
	 NonOperationalTo  	 string             	         	    	 Nullable 	
	 OperationalFrom  	 string             	         	   	 Nullable	
	 OperationalTo  	 string             	         	    	 Nullable 	

### Available Fields for ResourceCost
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 ResourceCostID  	 string           	         	  	 Readonly  	
	 ProductID  	 string             	         	 Yes  	 Identifier of service product	
	 AccountCode  	 string             	    50     	 Yes   	 Expense Account 	
	 Cost  	 float32             	         	   	 Readonly	
	 PriceTier  	 int32            	         	  Yes  	 Range 1-10 	
	 PlannedDowntimeCost  	 float32             	         	   	 Readonly	
	 PlannedDowntimePriceTier  	 int32            	         	   	 Nullable. Range 1-10 	
	 NotPlannedDowntimeCost  	 float32             	         	   	 Readonly	
	 NotPlannedDowntimePriceTier  	 int32            	         	    	 Nullable. Range 1-10 	

### Available Fields for ResourceRemark
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 RemarkID  	 string           	         	  	 Readonly  	
	 Remark  	 string             	    1024     	 Yes   	 	
	 CreatedDate  	 string             	         	    	 Readonly 	

### Available Fields for ResourceAttachment
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 AttachmentID  	 string           	         	  	 Readonly  	
	 FileName  	 string             	    256     	 Yes  	 	
	 Date  	 string             	         	    	 Readonly 	
	 ContentType  	 string             	 256        	   	 Readonly	
	 Content  	 string             	   32 Mb      	  Yes  	 Base64 encoded string 	

### Available Fields for SuspendReason
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 SuspendReasonID  	 string           	         	 Yes* 	 If it is empty, then will add new Suspend Reason, otherwise it will be modified 	
	 Reason  	 string             	         	 Yes*  	 Only for new entity 	
	 Workcenters  	 string Array             	         	    	  	

### Available Fields for WorkCenter
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	WorkCenterID	string		Yes	Identifier of a Workcenter	
	Code	string	256	Yes		
	Name	string	256	Yes		
	IsActive	boolean		Yes		
	IsCoMan	boolean		Yes		
	IsCoManPurchase	boolean		Yes		
	SupplierID	string		Yes*	Is required when IsCoMan = true or IsCoMan = true and IsCoManPurchase = true.	
	SupplierName	string				
	CoManProcurementType	string	10	Yes*	It may be 'Transfer' or 'Buysell' or 'Purchase'. Required for Co-man workcenters only.	
	WorkCenterLocations	<a href="#WorkCenterLocation"><b>[] WorkCenter Location Model</b></a>			Empty when components are picked directly from a shopfloor location.	
	WorkCenterSuppliers	<a href="#WorkCenterSupplier"><b>[] WorkCenter Supplier Model</b></a>		Yes*	Required when IsCoMan = false and IsCoManPurchase = true	

### Available Fields for WorkCenter Location <a name="WorkCenterLocation" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	string		Yes	Identifier of a WorkCenter Location	
	LocationID	string		Yes	Can be LocationID or BinID if current location is a Bin	
	Type	string	256	Yes	Consumption, Output	
	ParentLocationID	string			Parent location ID when current location is Bin	
	LocationName	string		Yes		

### Available Fields for WorkCenter Supplier <a name="WorkCenterSupplier" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	SupplierID	string		Yes*	Is required when SupplierName is empty.	
	SupplierName	string		Yes*	Is required if SupplierID is empty	

### Available Fields for Purchase List
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 ID  	 string           	         	          	Unique Cin7 Core Purchase identifier          	
	 BlindReceipt  	 boolean             	         	          	Specifies, if Purchase is a Blind Receipt order (no purchase order)         	
	 OrderNumber  	 string             	 256        	          	Purchase Order number generated by Cin7 Core         	
	 Status  	 string          	  50       	          	Current Purchase status. Current sale status. Possible values are <a href="#PurchaseStatesList"><b>values</b></a>        	
	 OrderDate  	 Date         	         	          	 Date when Purchase Order was created        	
	 InvoiceDate  	 Date            	         	          	 Date when Invoice was issued       	
	 Supplier  	 Date            	 256        	          	 Name of the Supplier        	
	 SupplierID  	 string 	         	          	 Identifier of Supplier.       	
	 InvoiceNumber  	 string 	   256      	          	 Invoice number issued by Supplier       	
	 InvoiceAmount  	 float32 	         	          	 Total Invoice amount in Supplier currency        	
	 PaidAmount  	 float32 	         	          	 Total Paid amount minus Refunded amount in Supplier currency        	
	 InvoiceDueDate  	 Date 	         	          	 Date when invoice is due according to selected payment terms         	
	 RequiredBy  	 Date 	         	          	 Date when shipment is due     	
	 BaseCurrency  	 string 	 3        	          	 3 digit Base currency code (as configured in General Settings)        	
	 SupplierCurrency  	 string 	 3        	          	 3 digit Supplier currency code       	
	 CreditNoteNumber  	 string 	 256        	          	 Credit note number issued by Supplier. Is empty unless credit note is created       	
	 OrderStatus  	 string 	 20        	          	 Purchase Order status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE	
	 StockReceivedStatus  	 string 	 20        	          	 Purchase stock received status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE        	
	 UnstockStatus  	 string 	 20        	          	 Invoice status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE      	
	 InvoiceStatus  	 string 	 20        	          	 Purchase unstock status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE   	
	 CreditNoteStatus  	 string 	 20        	          	 Credit Note status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE       	
	 LastUpdatedDate  	 Date 	         	          	 Date and time when purchase task was created of last time modified    	
	 CombinedReceivingStatus  	 string 	 20        	 Yes         	 Combined Put Away status. Possible Values are FULLY RECEIVED, PARTIALLY RECEIVED, NOT AVAILABLE, NOT RECEIVED, VOIDED         	
	 CombinedInvoiceStatus  	 string 	 20        	 Yes         	 Combined Invoice status. Possible Values are INVOICED, INVOICED / CREDITED, NOT AVAILABLE, NOT INVOICED, PARTIALLY INVOICED, PARTIALLY INVOICED / CREDITED        	
	 CombinedPaymentStatus  	 string 	 20        	 Yes         	 Payment status. Possible Values are PREPAID, PARTIALLY PAID, UNPAID, PAID, OVERPAID / CREDITED, VOIDED       	
	 Type  	 string 	 20        	 Yes         	 Type of Purchase. Possible Values are Simple Purchase, Advanced Purchase , Service Purchase        	
	 IsServiceOnly  	 boolean             	         	          	Specifies, if Purchase is a Service Purchase         	
	 DropShipTaskID  	 string             	         	          	ID of sale which created drop ship pruchase         	

### Available Fields for Purchase Credit Note List
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 ID  	 string           	         	          	Unique Cin7 Core Purchase identifier          	
	 BlindReceipt  	 boolean             	         	          	Specifies, if Purchase is a Blind Receipt order (no purchase order)         	
	 OrderNumber  	 string             	 256        	          	Purchase Order number generated by Cin7 Core         	
	 Status  	 string          	  50       	          	Current Purchase status. Current sale status. Possible values are <a href="#PurchaseStatesList"><b>values</b></a>        	
	 OrderDate  	 Date         	         	          	 Date when Purchase Order was created        	
	 InvoiceDate  	 Date            	         	          	 Date when Invoice was issued       	
	 Supplier  	 Date            	 256        	          	 Name of the Supplier        	
	 SupplierID  	 string 	         	          	 Identifier of Supplier.       	
	 InvoiceNumber  	 string 	   256      	          	 Invoice number issued by Supplier       	
	 InvoiceAmount  	 float32 	         	          	 Total Invoice amount in Supplier currency        	
	 PaidAmount  	 float32 	         	          	 Total Paid amount minus Refunded amount in Supplier currency        	
	 InvoiceDueDate  	 Date 	         	          	 Date when invoice is due according to selected payment terms         	
	 RequiredBy  	 Date 	         	          	 Date when shipment is due     	
	 BaseCurrency  	 string 	 3        	          	 3 digit Base currency code (as configured in General Settings)        	
	 SupplierCurrency  	 string 	 3        	          	 3 digit Supplier currency code       	
	 CreditNoteNumber  	 string 	 256        	          	 Credit note number issued by Supplier. Is empty unless credit note is created       	
	 OrderStatus  	 string 	 20        	          	 Purchase Order status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE	
	 StockReceivedStatus  	 string 	 20        	          	 Purchase stock received status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE        	
	 UnstockStatus  	 string 	 20        	          	 Invoice status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE      	
	 InvoiceStatus  	 string 	 20        	          	 Purchase unstock status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE   	
	 CreditNoteStatus  	 string 	 20        	          	 Credit Note status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE       	
	 LastUpdatedDate  	 Date 	         	          	 Date and time when purchase task was created of last time modified    	
	 CombinedReceivingStatus  	 string 	 20        	 Yes         	 Combined Put Away status. Possible Values are FULLY RECEIVED, PARTIALLY RECEIVED, NOT AVAILABLE, NOT RECEIVED, VOIDED         	
	 CombinedInvoiceStatus  	 string 	 20        	 Yes         	 Combined Invoice status. Possible Values are INVOICED, INVOICED / CREDITED, NOT AVAILABLE, NOT INVOICED, PARTIALLY INVOICED, PARTIALLY INVOICED / CREDITED        	
	 CombinedPaymentStatus  	 string 	 20        	 Yes         	 Payment status. Possible Values are PREPAID, PARTIALLY PAID, UNPAID, PAID, OVERPAID / CREDITED, VOIDED       	
	 Type  	 string 	 20        	 Yes         	 Type of Purchase. Possible Values are Simple Purchase, Advanced Purchase , Service Purchase, Purchase Credit Note, Credit Note        	
	 IsServiceOnly  	 boolean             	         	          	Specifies, if Purchase is a Service Purchase         	
	 DropShipTaskID  	 string             	         	          	ID of sale which created drop ship pruchase         	

### Available Fields for Purchase
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	string			Unique Cin7 Core Purchase ID	
	SupplierID	string			Identifier of Supplier	
	Supplier	string	256		Supplier name	
	Contact	string	256		Supplier Contact name	
	Phone	string	50		Supplier Contact phone	
	InventoryAccount	string	50		Account code used by default for invoice lines when no Inventory account is defined on Product. By default Inventory Control Account is used from account mapping.	
	BlindReceipt	boolean			True if there is no order in the purchase. False is default	
	Approach	string	10		INVOICE for Invoice  First, or STOCK for Stock First	
	BillingAddress	<a href="#AddressModel"><b>Address Model</b></a>			Purchase Billing address	
	ShippingAddress	<a href="#PurchaseShippingAddressModel"><b>Shipping Address Model</b></a>			Purchase Shipping address	
	BaseCurrency	string	3		3 character currency code of Base Currency defined in General Settings on the moment when Purchase was created.	
	SupplierCurrency	string	3		3 character currency code of Supplier Currency defined in Supplier card at the moment when Supplier is selected for the Purchase.	
	TaxRule	string	50		Default Tax Rule name selected for Purchase	
	TaxCalculation	string			Inclusive or Exclusive	
	Terms	string	256		Payment terms name	
	RequiredBy	string			Date when shipment is due	
	Location	string	256		Default location to store stock to	
	Note	string	1024		Custom Purchase note	
	OrderNumber	string	256		Represents Purchase Order Number	
	Status	string	20		Current Purchase status. Current sale status. Possible values are <a href="#PurchaseStatesList"><b>values</b></a>	
	RelatedDropShipSaleTask	string			ID of SaleTask this drop-ship purchase is related to.	
	CurrencyRate	float32			Conversion Rate expressed as number of Base currency units for one Supplier currency unit	
	LastUpdatedDate	string			Date and time when purchase task was created of last time modified	
	Order	<a href="#PurchaseOrderModel"><b>Purchase Order Model</b></a>			Purchase Order	
	StockReceived	<a href="#PurchaseStockModel"><b>Purchase Stock Model</b></a>			Purchased and received items	
	Invoice	<a href="#PurchaseInvoiceModel"><b>Purchase Invoice Model</b></a>			Purchase Invoice	
	CreditNote	<a href="#PurchaseCreditNoteModel"><b>Purchase Credit Note Model</b></a>			Purchase Credit Note	
	ManualJournals	<a href="#PurchaseManualJournalModel"><b>Purchase Manual Journal Model</b></a>			Purchase Manual Journals	
	AdditionalAttributes	<a href="#AdditionalAttributeModel"><b>Additional Attribute Model</b></a>			Purchase Additional Attributes	
	Attachments	<a href="#AttachmentLineModel"><b>[] Attachment Line Model</b></a>			Purchase Attachments	
	InventoryMovements	<a href="#InventoryMovementLineModel"><b>Inventory Movement Line Model</b></a>			Purchase Inventory Movements	

### Available Fields for Purchase Order
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Purchase Task ID	
	CombineAdditionalCharges	boolean		Yes	if true then additional charges lines displayed in Lines array	
	Memo	string	1024	Yes	Additional information for Order	
	Status	string		Yes	Purchase Order Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#PurchaseOrderLineModel"><b>[] Purchase Order Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#PurchaseAdditionalChargeModel"><b>[] Purchase Additional Charge Model</b></a>				
	TotalBeforeTax	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges without taxes. Not required for POST.	
	Tax	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges taxes. Not required for POST.	
	Total	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges with taxes. Not required for POST.	

### Available Fields for Purchase Stock Received
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Purchase Task ID	
	Status	string		Yes	Purchase Stock Received Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#PurchaseStockLineModel"><b>[] Purchase Stock Line Model</b></a>		Yes		

### Available Fields for Purchase Invoice
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Purchase Task ID	
	CombineAdditionalCharges	boolean		Yes	if true then additional charges lines displayed in Lines array	
	InvoiceDate	string		Yes	Date when invoice created. Default value is current date, used if not specified.	
	InvoiceDueDate	string		Yes	Date until invoice is valid. If not specified, used default value from Terms.	
	InvoiceNumber	string			Invoice Number (auto-generated)	
	Status	string		Yes	Purchase Invoice Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED, PAID. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#PurchaseInvoiceLineModel"><b>[] Purchase Invoice Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#PurchaseInvoiceAdditionalChargeModel"><b>[] Purchase Invoice Additional Charge Model</b></a>				
	TotalBeforeTax	float32 with up to 4 decimal places			Sum of order lines and additional charges without taxes. Not required for POST.	
	Tax	float32 with up to 4 decimal places			Sum of order lines and additional charges taxes. Not required for POST.	
	Total	float32 with up to 4 decimal places			Sum of order lines and additional charges with taxes. Not required for POST.	
	InvoiceTotalAmount	float32 with up to 4 decimal places			Not required for POST.	
	InvoiceTotalTaxAmount	float32 with up to 4 decimal places			Not required for POST.	

### Available Fields for Purchase Credit Note
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Purchase Task ID	
	CombineAdditionalCharges	boolean		Yes	if true then additional charges lines displayed in Lines array	
	CreditNoteNumber	string		Yes	Credit Note Number	
	CreditNoteDate	string		Yes	Credit Note Date	
	Status	string		Yes	Purchase Credit Note Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#PurchaseInvoiceLineModel"><b>[] Purchase Invoice Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#PurchaseInvoiceAdditionalChargeModel"><b>[] Purchase Invoice Additional Charge Model</b></a>				
	Unstock	<a href="#PurchaseUnStockLineModel"><b>[] Purchase Unstock Line Model</b></a>		Yes		
	TotalBeforeTax	float32 with up to 4 decimal places			Sum of order lines and additional charges without taxes. Not required for POST.	
	Tax	float32 with up to 4 decimal places			Sum of order lines and additional charges taxes. Not required for POST.	
	Total	float32 with up to 4 decimal places			Sum of order lines and additional charges with taxes. Not required for POST.	

### Available Fields for Purchase Payments
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Purchase Task ID.	
	ID	GUID		Yes*	Identifier of payment. Available for PUT	
	Type	string		Yes	Available values are PREPAYMENT,PAYMENT,REFUND. Available for POST	
	Reference	string			Payment reference number. Available for POST/PUT	
	Amount	float32 with up to 2 decimal places		Yes	Payment amount in customer currency. Available for POST/PUT	
	DatePaid	string		Yes	Date when payment has been made. Available for POST/PUT	
	Account	string		Yes	Account Code of the bank/payment account from Chart of accounts. Available for POST/PUT	
	CurrencyRate	float32 with up to 4 decimal places		Yes	Currency Conversion rate expressed as number of Base currency units for one Customer currency unit. Available for POST/PUT	
	DateCreated	string			Date of creation payment record.	

### Available Fields for Purchase
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	string			Unique Cin7 Core Purchase ID	
	SupplierID	string			Identifier of Supplier	
	Supplier	string	256		Supplier name	
	Contact	string	256		Supplier Contact name	
	Phone	string	50		Supplier Contact phone	
	InventoryAccount	string	50		Account code used by default for invoice lines when no Inventory account is defined on Product. By default Inventory Control Account is used from account mapping.	
	BlindReceipt	boolean			True if there is no order in the purchase. False is default	
	Approach	string	10		INVOICE for Invoice  First, or STOCK for Stock First	
	BillingAddress	<a href="#AddressModel"><b>Address Model</b></a>			Purchase Billing address	
	ShippingAddress	<a href="#PurchaseShippingAddressModel"><b>Shipping Address Model</b></a>			Purchase Shipping address	
	BaseCurrency	string	3		3 character currency code of Base Currency defined in General Settings on the moment when Purchase was created.	
	SupplierCurrency	string	3		3 character currency code of Supplier Currency defined in Supplier card at the moment when Supplier is selected for the Purchase.	
	TaxRule	string	50		Default Tax Rule name selected for Purchase	
	TaxCalculation	string			Inclusive or Exclusive	
	Terms	string	256		Payment terms name	
	RequiredBy	string			Date when shipment is due	
	Location	string	256		Default location to store stock to	
	Note	string	1024		Custom Purchase note	
	OrderNumber	string	256		Represents Purchase Order Number	
	OrderDate	string			Represents Purchase Order Date	
	 CombinedReceivingStatus  	 string 	 20        	          	 Combined Put Away status. Possible Values are FULLY RECEIVED, PARTIALLY RECEIVED, NOT AVAILABLE, NOT RECEIVED, VOIDED         	
	 CombinedInvoiceStatus  	 string 	 20        	          	 Combined Invoice status. Possible Values are INVOICED, INVOICED / CREDITED, NOT AVAILABLE, NOT INVOICED, PARTIALLY INVOICED, PARTIALLY INVOICED / CREDITED        	
	 CombinedPaymentStatus  	 string 	 20        	         	 Payment status. Possible Values are PREPAID, PARTIALLY PAID, UNPAID, PAID, OVERPAID / CREDITED, VOIDED       	
	 Type  	 string 	 20        	         	 Type of Purchase. Possible Values are Simple Purchase, Advanced Purchase , Service Purchase        	
	 IsServiceOnly  	 boolean             	         	          	Specifies, if Purchase is a Service Purchase         	
	Status	string	20		Current Purchase status. Current sale status. Possible values are <a href="#PurchaseStatesList"><b>values</b></a>	
	RelatedDropShipSaleTask	string			ID of SaleTask this drop-ship purchase is related to.	
	CurrencyRate	float32			Conversion Rate expressed as number of Base currency units for one Supplier currency unit	
	LastUpdatedDate	string			Date and time when purchase task was created of last time modified	
	Order	<a href="#PurchaseOrderModel"><b>Purchase Order Model</b></a>			Purchase Order	
	StockReceived	<a href="#AdvancedPurchaseStockModel"><b>[] Advanced Purchase Stock Model</b></a>			Purchased and received items. The field is filled only for Advanced Purchases.	
	PutAway	<a href="#AdvancedPurchasePutAwayModel"><b>[] Advanced Purchase Put Away Model</b></a>			Purchased and put away items	
	Invoice	<a href="#AdvancedPurchaseInvoiceModel"><b>[] Advanced Purchase Invoice Model</b></a>			Purchase Invoice	
	CreditNote	<a href="#AdvancedPurchaseCreditNoteModel"><b>[] Advanced Purchase Credit Note Model</b></a>			Purchase Credit Note	
	ManualJournals	<a href="#AdvancedPurchaseManualJournalModel"><b>[] Advanced Purchase Manual Journal Model</b></a>			Purchase Manual Journals	
	AdditionalAttributes	<a href="#AdditionalAttributeModel"><b>Additional Attribute Model</b></a>			Purchase Additional Attributes	
	Attachments	<a href="#AttachmentLineModel"><b>[] Attachment Line Model</b></a>			Purchase Attachments	
	InventoryMovements	<a href="#InventoryMovementLineModel"><b>Inventory Movement Line Model</b></a>			Purchase Inventory Movements	

### Available Fields for Purchase Stock Received
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	PurchaseID	string		Yes	Unique Cin7 Core Purchase Master Task ID	
	TaskID	string		Yes	Unique Cin7 Core Purchase Stock Receiving Task ID	
	Status	string		Yes	Purchase Stock Received Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST and PUT only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#AdvancedPurchaseStockLineModel"><b>[] Advanced Purchase Stock Line Model</b></a>		Yes		

### Available Fields for Purchase Put Away
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	PurchaseID	string		Yes	Unique Cin7 Core Purchase Master Task ID	
	TaskID	string		Yes	Unique Cin7 Core Purchase Put Away Task ID	
	Status	string		Yes	Purchase Put Away Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#AdvancedPurchasePutAwayLineModel"><b>[] Advanced Purchase Put Away Line Model</b></a>		Yes		

### Available Fields for Purchase Invoice

### Available Fields for Purchase Credit Note
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	PurchaseID	string		Yes	Unique Cin7 Core Purchase Master Task ID	
	CreditNotes	<a href="#AdvancedPurchasePartialCreditNoteModel"><b>[] Advanced Purchase Credit Note Model</b></a>		Yes		

### Available Fields for Purchase Payments
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Purchase Task ID.	
	ID	GUID		Yes*	Identifier of payment. Available for PUT	
	Type	string		Yes	Available values are PREPAYMENT,PAYMENT,REFUND. Available for POST	
	Reference	string			Payment reference number. Available for POST/PUT	
	Amount	float32 with up to 2 decimal places		Yes	Payment amount in customer currency. Available for POST/PUT	
	DatePaid	string		Yes	Date when payment has been made. Available for POST/PUT	
	Account	string		Yes	Account Code of the bank/payment account from Chart of accounts. Available for POST/PUT	
	CurrencyRate	float32 with up to 4 decimal places		Yes	Currency Conversion rate expressed as number of Base currency units for one Customer currency unit. Available for POST/PUT	
	DateCreated	string			Date of creation payment record.	

### Available Fields for Product Deal <a name="ProductDeal" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID               	 string      	       	 Yes      	                       	
	 Name              	 string      	       	 Yes      	      Deal Name        	
	 DateFrom        	 Date    	        	       	                       	
	 DateTo            	 Date    	        	       	                       	
	 IsActive          	 bool    	     	 Yes      	                        	
	 AllowCoupons          	 bool    	     	 Yes      	                        	
	 SingleCouponCodeUsage          	 bool    	     	 Yes      	                        	
	 CustomersGroup     	 string    	       	           	 All, B2B, POS, Specified, CustomerTags 	
	 CouponCodes        	 string   	       	        	      	
	DealCustomers	<a href="#ProductDealCustomerModel"><b>[] Customer Model</b></a>				
	DealCustomerTags	<a href="#ProductDealTagModel"><b>[] Customer Tag Model</b></a>				
	DealDiscounts	<a href="#ProductDealDiscountsModel"><b>[] Deal Discount Model</b></a>				

### Available Fields for Product Deal Customer <a name="ProductDealCustomerModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID              	 string      	       	 Yes*      	                       	
	 CustomerName     	 string      	       	 Yes*      	                         	
	 CustomerID        	 string        	           	 Yes*     	                       	

### Available Fields for Product Deal Customer Tag <a name="ProductDealTagModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID              	 string      	       	 Yes*      	                       	
	 TagName         	 string    	       	 Yes*      	                       	

### Available Fields for Product Deal Discount <a name="ProductDealDiscountModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID              	 string      	       	 Yes*      	                       	
	 DiscountID        	 string        	           	 Yes*     	                       	
	 DiscountName     	 string    	       	 Yes*      	                         	
	 DiscountType    	 string    	       	 Yes      	  Simple, QuantityBased, FreeShipping 	
	 Sequence        	 int32       	           	              	                       	
	 IsOrderLevel     	 bool        	       	 Yes      	                         	
	 BuyType            	 string    	           	              	                       	
	 BuyValue         	 float32   	       	           	                         	
	 BuyValueType    	 string       	           	              	                       	
	 BuyMore         	 bool        	       	           	                         	
	 GetType            	 string    	           	              	                       	
	 GetValue         	 float32   	       	           	                         	
	 GetValueType    	 string       	           	              	                       	
	 BuyMore         	 bool        	       	           	                         	
	DealDiscountBrands	<a href="#ProductDealDiscountBrandModel"><b>[] Deal Discount Brand Model</b></a>				
	DealDiscountCategories	<a href="#ProductDealDiscountCategoryModel"><b>[] Deal Discount Category Model</b></a>				
	DealDiscountTags	<a href="#ProductDealDiscountTagModel"><b>[] Deal Discount Tag Model</b></a>				
	DealDiscountProducts	<a href="#ProductDealDiscountProductModel"><b>[] Deal Discount Product Model</b></a>				

### Available Fields for Product Deal Discount Brand <a name="ProductDealDiscountBrandModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID              	 string      	       	 Yes*      	                       	
	 BrandName       	 string    	       	 Yes*      	                       	
	 BrandID         	 string      	       	 Yes*      	                       	
	 Type            	 string    	       	 Yes       	 Buy or Get            	

### Available Fields for Product Deal Discount Product <a name="ProductDealDiscountProductModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID              	 string      	       	 Yes*      	                       	
	 ProductSKU      	 string    	       	 Yes*      	                       	
	 ProductID       	 string      	       	 Yes*      	                       	
	 ProductName     	 string    	       	           	      Readonly         	
	 IsFamily        	 bool      	       	 Yes       	                       	
	 Type            	 string    	       	 Yes       	 Buy or Get            	

### Available Fields for Product Deal Discount Tag <a name="ProductDealDiscountTagModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID              	 string      	       	 Yes*      	                       	
	 TagName         	 string    	       	 Yes*      	                       	
	 Type            	 string    	       	 Yes       	 Buy or Get            	

### Available Fields for Product Deal Discount Category <a name="ProductDealDiscountCategoryModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID              	 string      	       	 Yes*      	                       	
	 CategoryName    	 string    	       	 Yes*      	                       	
	 Type            	 string    	       	 Yes       	 Buy or Get            	

### Available Fields for Product Discount Rule <a name="ProductDiscountRuleModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID               	 string      	       	 Yes      	                       	
	 Name              	 string      	   128    	 Yes      	      Deal Name        	
	 IsActive          	 bool    	     	 Yes      	                        	
	 Type             	 string    	       	    Yes       	 Simple, QuantityBased, FreeShipping 	
	DiscountLines	<a href="#DiscountLineModel"><b>[] Discount Line Model</b></a>				

### Available Fields for Discount Line Model <a name="DiscountLineModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID               	 string      	       	 Yes      	                       	
	 MinValue         	 float32      	       	       	 Applicable only when Product Discount Rule Type = QuantityBase and only for the first item in the list. Subsequent items MinValue is being calculated as previous item MaxValue + 0.0001.	
	 MaxValue          	 float32    	     	       	 Applicable only when Product Discount Rule Type = QuantityBase 	
	 DiscountType     	 string    	       	           	 DiscountAmount, DiscountPercent, MarkupAmount, MarkupPercent, PriceOverride, FlatAmount, FreeShipping 	
	 Amount          	 float32    	     	 Yes*      	                        	
	 OrderExceeds          	 float32    	     	 Yes*      	 Required when DiscountType = FreeShipping      	

### Available Fields for Product Supplier
See <a href="#ProductSupplierModel"><b>Product Supplier Model</b></a>

### Available Fields for Product Supplier Options Model
See <a href="#ProductSupplierOptionsModel"><b>Product Supplier Options Model</b></a>

### Available Fields for Product Supplier Options Interval Model
See <a href="#ProductSupplierOptionsIntervalModel"><b>Product Supplier Options Interval Model</b></a>

### Available Fields for Shipping Zone <a name="ShippingZoneModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ZoneID               	 string      	       	 Yes      	                       	
	 Name              	 string      	  128     	 Yes      	              	
	 IsRestZone        	 bool    	     	 Yes      	                        	
	 PricesInclTax        	 bool    	     	 Yes      	                        	
	 Negative        	 bool    	     	 Yes      	                        	
	 DefaultShippingCost        	 float32    	     	 Yes      	                        	
	 ShortDesc        	 bool    	     	       	       Readonly                 	
	 TaxRuleID        	 string    	     	 Yes*      	                        	
	 TaxRuleName        	 string    	  255   	 Yes*      	                        	
	AppliesTo	<a href="#ShipZoneAppliesToModel"><b>[] Applies To Model</b></a>				
	Conditions	<a href="#ShipZoneConditionsModel"><b>[] Condition Model</b></a>				

### Available Fields for Applies To Model <a name="ShipZoneAppliesToModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID               	 string      	       	 Yes      	                       	
	 Country2         	 string      	       	       	  US, SO, NG etc.       	
	 StateID          	 int32   	     	       	                        	
	 StateName         	 string    	       	           	 AL, LA, CA etc. 	
	 PostCodeFrom    	 string    	  100   	       	                        	
	 PostCodeTo    	 string    	  100   	       	                        	
	 PostCodesList    	 string    	     	       	                        	
	 ShippingRate    	 float32    	     	    Yes   	                        	

### Available Fields for Condition Model <a name="ShipZoneConditionModel" />
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ZoneConditionID               	 string      	       	 Yes      	                       	
	 MinValue    	 float32    	     	       	                        	
	 MaxValue    	 float32    	     	       	                        	
	 ShippingCost    	 float32    	     	    Yes   	                        	
	 ConditionType   	 string    	       	    Yes       	 Price, Weight 	
	 ConditionDescription    	 string    	  128   	       	                        	

### Available Fields for Sale List
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 SaleID  	 string           	         	  Yes        	Unique Cin7 Core Sale identifier          	
	 OrderNumber  	 string             	 256        	 Yes         	Sale Order number generated by Cin7 Core          	
	 Status  	 string          	         	 Yes         	Current sale status. Possible values are <a href="#SaleStatesList"><b>values</b></a>        	
	 OrderDate  	 Date         	         	 Yes         	 Date when Sale Order was created         	
	 InvoiceDate  	 Date            	         	          	 Date when Invoice was issued         	
	 Customer  	 Date            	 256        	 Yes         	 Name of the customer         	
	 CustomerID  	 string 	         	          	 Customer Identifier         	
	 InvoiceNumber  	 string 	   256      	          	 Invoice number generated by Cin7 Core         	
	 CustomerReference  	 string 	 256        	          	 Optional Customer Reference (typically customer Purchase order number) supplied by the customer for the sale         	
	 InvoiceAmount  	 float32 	         	  Yes        	 Total Invoice amount minus total Credit note amount in customer currency        	
	 PaidAmount  	 float32 	         	  Yes        	 Total Paid amount minus Refunded amount in customer currency         	
	 InvoiceDueDate  	 Date 	         	          	 Date when invoice is due according to selected payment terms         	
	 ShipBy  	 Date 	         	          	 Date when shipment is due     	
	 BaseCurrency  	 string 	 3        	 Yes         	 3 digit Base currency code (as configured in General Settings)        	
	 CustomerCurrency  	 string 	 3        	 Yes         	 3 digit Customer currency code       	
	 CreditNoteNumber  	 string 	 256        	          	 Credit note number generated by Cin7 Core. Is empty unless credit note is created       	
	 Updated  	 string 	         	 Yes         	 Date when the sale was last created/updated last time        	
	 QuoteStatus  	 string 	 20        	 Yes         	 Sale Quote status. Possible Values are <a href="#QuoteStatuses"><b>values</b></a>	
	 OrderStatus  	 string 	 20        	 Yes         	 Sale Order status. Possible Values are <a href="#OrderStatuses"><b>values</b></a>        	
	 CombinedPickingStatus  	 string 	 20        	 Yes         	 Pick status. Possible Values are VOIDED, NOT AVAILABLE, PICKED, PICKING , NOT PICKED , PARTIALLY PICKED        	
	 CombinedPackingStatus  	 string 	 20        	 Yes         	 Pack status. Possible Values are VOIDED, NOT AVAILABLE, PACKED, PACKING, NOT PACKED, PARTIALLY PACKED       	
	 CombinedShippingStatus  	 string 	 20        	 Yes         	 Ship status. Possible Values are VOIDED, NOT AVAILABLE, SHIPPED, SHIPPING , NOT SHIPPED   , PARTIALLY SHIPPED         	
	 FulFilmentStatus  	 string 	 20        	 Yes         	 Fulfilment status. Possible Values are FULFILLED, PARTIALLY FULFILLED, NOT AVAILABLE, NOT FULFILLED, VOIDED         	
	 CombinedInvoiceStatus  	 string 	 20        	 Yes         	 Invoice status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE, PAID        	
	 CreditNoteStatus  	 string 	 20        	 Yes         	 Credit Note status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE       	
	 CombinedPaymentStatus  	 string 	 20        	 Yes         	 Payment status. Possible Values are NOT REFUNDED, PREPAID, PARTIALLY PAID, UNPAID, PAID, VOIDED       	
	 Type  	 string 	 50        	 Yes         	 Type of Sale. Possible Values are Simple Sale, Advanced Sale , Service Sale        	
	 CombinedTrackingNumbers  	 string 	 256        	 Yes         	 Tracking Numbers        	
	 SourceChannel  	 string 	 32        	          	 Source of the sale. read-only field       	
	 ExternalID  	 string 	 256        	          	 Custom field that is only available in API and allows to set arbitrary value for the sale for later search and any custom logic        	
	 OrderLocationID  	 string 	         	          	 Sale Order Location ID         	

### Available Fields for Sale Credit Note List
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 SaleID  	 string           	         	  Yes        	Unique Cin7 Core Sale identifier          	
	 OrderNumber  	 string             	 256        	 Yes         	Sale Order number generated by Cin7 Core          	
	 Status  	 string          	         	 Yes         	Current sale status. Possible values are <a href="#SaleStatesList"><b>values</b></a>        	
	 OrderDate  	 Date         	         	 Yes         	 Date when Sale Order was created         	
	 InvoiceDate  	 Date            	         	          	 Date when Invoice was issued         	
	 Customer  	 Date            	 256        	 Yes         	 Name of the customer         	
	 CustomerID  	 string 	         	          	 Customer Identifier         	
	 InvoiceNumber  	 string 	   256      	          	 Invoice number generated by Cin7 Core         	
	 CustomerReference  	 string 	 256        	          	 Optional Customer Reference (typically customer Purchase order number) supplied by the customer for the sale         	
	 InvoiceAmount  	 float32 	         	  Yes        	 Total Invoice amount minus total Credit note amount in customer currency        	
	 PaidAmount  	 float32 	         	  Yes        	 Total Paid amount minus Refunded amount in customer currency         	
	 InvoiceDueDate  	 Date 	         	          	 Date when invoice is due according to selected payment terms         	
	 ShipBy  	 Date 	         	          	 Date when shipment is due     	
	 BaseCurrency  	 string 	 3        	 Yes         	 3 digit Base currency code (as configured in General Settings)        	
	 CustomerCurrency  	 string 	 3        	 Yes         	 3 digit Customer currency code       	
	 CreditNoteNumber  	 string 	 256        	          	 Credit note number generated by Cin7 Core. Is empty unless credit note is created       	
	 Updated  	 string 	         	 Yes         	 Date when the sale was last created/updated last time        	
	 QuoteStatus  	 string 	 20        	 Yes         	 Sale Quote status. Possible Values are <a href="#QuoteStatuses"><b>values</b></a>	
	 OrderStatus  	 string 	 20        	 Yes         	 Sale Order status. Possible Values are <a href="#OrderStatuses"><b>values</b></a>        	
	 CombinedPickingStatus  	 string 	 20        	 Yes         	 Pick status. Possible Values are VOIDED, NOT AVAILABLE, PICKED, PICKING , NOT PICKED , PARTIALLY PICKED        	
	 CombinedPackingStatus  	 string 	 20        	 Yes         	 Pack status. Possible Values are VOIDED, NOT AVAILABLE, PACKED, PACKING, NOT PACKED, PARTIALLY PACKED       	
	 CombinedShippingStatus  	 string 	 20        	 Yes         	 Ship status. Possible Values are VOIDED, NOT AVAILABLE, SHIPPED, SHIPPING , NOT SHIPPED   , PARTIALLY SHIPPED         	
	 FulFilmentStatus  	 string 	 20        	 Yes         	 Fulfilment status. Possible Values are FULFILLED, PARTIALLY FULFILLED, NOT AVAILABLE, NOT FULFILLED, VOIDED         	
	 CombinedInvoiceStatus  	 string 	 20        	 Yes         	 Invoice status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE, PAID        	
	 CreditNoteStatus  	 string 	 20        	 Yes         	 Credit Note status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE       	
	 CombinedPaymentStatus  	 string 	 20        	 Yes         	 Payment status. Possible Values are NOT REFUNDED, PREPAID, PARTIALLY PAID, UNPAID, PAID, VOIDED       	
	 Type  	 string 	 50        	 Yes         	 Type of Sale. Possible Values are Simple Sale, Advanced Sale , Service Sale, Sale Credit Note, Credit Note        	
	 CombinedTrackingNumbers  	 string 	 256        	 Yes         	 Tracking Numbers        	
	 SourceChannel  	 string 	 32        	          	 Source of the sale. read-only field       	
	 ExternalID  	 string 	 256        	          	 Custom field that is only available in API and allows to set arbitrary value for the sale for later search and any custom logic        	
	 OrderLocationID  	 string 	         	          	 Sale Order Location ID         	

### Available Fields for Sale
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	string			Unique Cin7 Core Sale ID. Required for PUT	
	Customer	string	256		Customer name	
	CustomerID	string			Customer identifier	
	Contact	string	256		Customer Contact name	
	Phone	string	50		Customer Contact phone	
	Email	string	256		Customer Contact email	
	DefaultAccount	string	50		Account code used by default for invoice lines when no revenue account is defined on Product. By default it is equal to Customer Sale account	
	SkipQuote	boolean			True if there is no quote in the sale	
	BillingAddress	<a href="#AddressModel"><b>Address Model</b></a>			Sale Billing address	
	ShippingAddress	<a href="#SaleShippingAddressModel"><b>Shipping Address Model</b></a>			Sale Shipping address	
	ShippingNotes	string	1024		Shipping Notes	
	BaseCurrency	string	3		3 character currency code of Base Currency defined in General Settings on the moment when Sale was created.	
	CustomerCurrency	string	3		3 character currency code of customer Currency defined in Customer card at the moment when customer is selected for the Sale.	
	TaxRule	string	50		Default Tax Rule name selected for Sale	
	TaxCalculation	string			Inclusive or Exclusive	
	Terms	string	256		Payment terms name	
	PriceTier	string	50		Price Tier name selected for Sale	
	ShipBy	Date			Date when shipment is due	
	Location	string	256		Default location to pick stock from	
	SaleOrderDate	Date			Date when task was created	
	LastModifiedOn	string			UTC Time	
	Note	string	1024		Custom Sale note	
	CustomerReference	string	256		Reference number used by customer to identify this sale. Could be a purchase order number generated by customer.	
	COGSAmount	Money			COGS amount in base currency	
	Status	string	25		Sale Status, see possible <a href="#SaleStatesList"><b>values</b></a>	
	 CombinedPickingStatus  	 string 	 20        	 Yes         	 Pick status. Possible Values are VOIDED, NOT AVAILABLE, PICKED, PICKING , NOT PICKED , PARTIALLY PICKED        	
	 CombinedPackingStatus  	 string 	 20        	 Yes         	 Pack status. Possible Values are VOIDED, NOT AVAILABLE, PACKED, PACKING, NOT PACKED, PARTIALLY PACKED       	
	 CombinedShippingStatus  	 string 	 20        	 Yes         	 Ship status. Possible Values are VOIDED, NOT AVAILABLE, SHIPPED, SHIPPING , NOT SHIPPED   , PARTIALLY SHIPPED         	
	 FulFilmentStatus  	 string 	 20        	          	 Fulfilment status. Possible Values are FULFILLED, PARTIALLY FULFILLED, NOT AVAILABLE, NOT FULFILLED, VOIDED         	
	 CombinedInvoiceStatus  	 string 	 20        	          	 Invoice status. Possible Values are VOIDED, DRAFT, AUTHORISED, NOT AVAILABLE, PAID        	
	 CombinedPaymentStatus  	 string 	 20        	          	 Payment status. Possible Values are NOT REFUNDED, PREPAID, PARTIALLY PAID, UNPAID, PAID, VOIDED       	
	 CombinedTrackingNumbers  	 string 	 256        	          	 Tracking Numbers        	
	 Carrier  	 string 	 256        	  Name of Carrier/shipping method        	       	
	CurrencyRate	float32	Up to 5 decimal places		Conversion Rate expressed as number of Base currency units for one Customer currency unit	
	SalesRepresentative	string	256		Sales representative name 	
	Type	string	50		Type of sale: Simple Sale or Advanced Sale, Service Sale	
	 SourceChannel  	 string 	 32        	          	 Source of the sale. read-only field       	
	 ExternalID  	 string 	 256        	          	 Custom field that is only available in API and allows to set arbitrary value for the sale for later search and any custom logic        	
	ServiceOnly	boolean			true when it is service-only sale	
	Quote	<a href="#SaleQuoteModel"><b>Sale Quote Model</b></a>			Sale Quote	
	Order	<a href="#SaleOrderModel"><b>Sale Order Model</b></a>			Sale Order	
	Fulfilments	<a href="#SaleFulfilmentModel"><b>[] Sale Fulfilment Model</b></a>			Sale Fulfilments	
	Invoices	<a href="#SaleInvoiceModel"><b>[] Sale Invoice Model</b></a>			Sale Invoices	
	CreditNotes	<a href="#SaleCreditNoteModel"><b>[] Sale Credit Note Model</b></a>			Sale Credit Notes	
	ManualJournals	<a href="#SaleManualJournalModel"><b>Sale Manual Journal Model</b></a>			Sale Manual Journals	
	AdditionalAttributes	<a href="#AdditionalAttributeModel"><b>Additional Attribute Model</b></a>			Sale Additional Attributes	
	Attachments	<a href="#AttachmentLineModel"><b>[] Attachment Line Model</b></a>			Sale Attachments	
	InventoryMovements	<a href="#InventoryMovementLineModel"><b>[] Inventory Movement Line Model</b></a>			Sale Inventory Movements	
	Transactions	<a href="#SaleTransactionLineModel"><b>[] Sale Transaction Line Model</b></a>			Sale Transactions	

### Available Fields for Sale Quote
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	SaleID	string		Yes	Unique Cin7 Core Sale ID	
	CombineAdditionalCharges	boolean		Yes	if true then additional charges lines displayed in Lines array	
	Memo	string	1024	Yes	Additional information for Quote	
	Status	string		Yes	Quote Status. Possible Values are <a href="#QuoteStatuses"><b>values</b></a>. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#SaleQuoteLineModel"><b>[] Sale Quote Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#SaleAdditionalChargeModel"><b>[] Sale Additional Charge Model</b></a>				
	TotalBeforeTax	float32 with up to 4 decimal places		Yes	Sum of quote lines and additional charges without taxes. Not required for POST.	
	Tax	float32 with up to 4 decimal places		Yes	Sum of quote lines and additional charges taxes. Not required for POST.	
	Total	float32 with up to 4 decimal places		Yes	Sum of quote lines and additional charges with taxes. Not required for POST.	

### Available Fields for Sale Order
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	SaleID	string		Yes	Unique Cin7 Core Sale ID	
	SaleOrderNumber	string	256		Sale Order Number (auto-generated)	
	CombineAdditionalCharges	boolean		Yes	if true then additional charges lines displayed in Lines array	
	Memo	string	1024	Yes	Additional information for Order	
	Status	string		Yes	Order Status. Possible Values are <a href="#OrderStatuses"><b>values</b></a>. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#SaleOrderLineModel"><b>[] Sale Order Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#SaleAdditionalChargeModel"><b>[] Sale Additional Charge Model</b></a>				
	TotalBeforeTax	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges without taxes. Not required for POST.	
	Tax	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges taxes. Not required for POST.	
	Total	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges with taxes. Not required for POST.	

### Available Fields for Sale Fulfilment
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	SaleID	string		Yes	Unique Cin7 Core Sale ID	
	Fulfilments	<a href="#SaleFulfilmentModel"><b>[] Sale Fulfilment Model</b></a>				

### Available Fields for Sale Fulfilment Pick
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Sale Task ID	
	Status	string		Yes	Pick status. Possible Values are NOT AVAILABLE, DRAFT, AUTHORISED, VOIDED. For POST available values are DRAFT, AUTHORISED	
	Lines	<a href="#SaleFulfilmentPickPackLineModel"><b>[] Sale Fulfilment Pick Pack Line Model</b></a>				

### Available Fields for Sale Fulfilment Pack
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Sale Task ID	
	Status	string		Yes	Pack status. Possible Values are NOT AVAILABLE, DRAFT, AUTHORISED, VOIDED. For POST available values are DRAFT, AUTHORISED	
	Lines	<a href="#SaleFulfilmentPickPackLineModel"><b>[] Sale Fulfilment Pick Pack Line Model</b></a>				

### Available Fields for Sale Fulfilment Ship
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Sale Task ID	
	Status	string		Yes	Ship status. Possible Values are NOT AVAILABLE, DRAFT, PARTIALLY AUTHORISED, AUTHORISED, VOIDED. For POST/PUT available values are DRAFT, PARTIALLY AUTHORISED, AUTHORISED	
	RequireBy	Date			Require By Date. Setting this will set or override the Required by date of the sale order.	
	ShippingAddress	<a href="#SaleShippingAddressModel"><b>Shipping Address Model</b></a>			Shipping Address	
	ShippingNotes	string	1024		Shipping Notes	
	Lines	<a href="#SaleFulfilmentShipLineModel"><b>[] Sale Fulfilment Ship Line Model</b></a>				

### Invoice Available Fields

### Credit Note Available Fields

### Available Fields for Stock Adjustment
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string			Unique Cin7 Core Stock Adjustment ID	
	EffectiveDate	string			Date of transaction	
	StocktakeNumber	string			Stock Take/Adjustment Number	
	Status	string			Status of stock adjustment. Available values are DRAFT,COMPLETED,VOIDED.	
	Account	string			Expense account for inventory adjustment	
	Reference	string			Custom reference	  
	ExistingStockLines	<a href="#ExistingStockLineModel"><b>[] Existing Stock Line Model</b></a>			Changes in non-zero stock	  
	NewStockLines	<a href="#NewStockLineModel"><b>[] New Stock Line Model</b></a>			Changes in zero stock	  
	Transactions	<a href="#TransactionStockLineModel"><b>[] Transaction Stock Line Model</b></a>			Created transactions due to stock changes.	      

### Available Fields for Post/Put
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string			Unique Cin7 Core Stock Adjustment ID. Required for PUT	
	EffectiveDate	string		Yes	Date of transaction	
	Status	string		Yes	Status of stock adjustment. Available values are DRAFT,COMPLETED.	
	Account	string			Expense account for inventory adjustment	
	Reference	string			Custom reference	  
	UpdateOnHand	boolean			If true, On hand quantity will be adjusted, otherwise available quantity will be adjusted.	  
	Lines	<a href="#NewStockLineModel"><b>[] New Stock Line Model</b></a>		Yes		  

### Available Fields for Stock Take
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes*	Unique Cin7 Core Stock Take ID. Required for PUT	
	EffectiveDate	string		Yes	Date of transaction	
	StocktakeNumber	string			Stock Take Number (auto-generated).	
	Status	string		Yes*	Status of stocktake. Available values are DRAFT, IN PROGRESS, COMPLETED, VOIDED. Required for PUT	
	Account	string		Yes	Expense account for inventory adjustment	
	LocationID	string		Yes*	Location ID. Required if Location is empty 	
	Location	string		Yes*	Location Name. Required if LocationID is empty 
	Tags	[] string			Tags used for additional product filtration. Can be used to filter NonZeroStockOnHandProducts	
	PickZones	[] string			Pick Zones used for additional product filtration. Can be used to filter NonZeroStockOnHandProducts	
	StockLocators	[] string			Stock Locators used for additional product filtration. Can be used to filter NonZeroStockOnHandProducts	
	Categories	<a href="#IDNameModel"><b>[] IDName Model</b></a>			Categories used for additional product filtration. Can be used to filter NonZeroStockOnHandProducts	
	Brands	<a href="#IDNameModel"><b>[] IDName Model</b></a>			Brands used for additional product filtration. Can be used to filter NonZeroStockOnHandProducts	
	Bins	<a href="#IDNameModel"><b>[] IDName Model</b></a>			Bins used for additional product filtration. Can be used to filter NonZeroStockOnHandProducts	
	Reference	string			Custom reference	  
	NonZeroStockOnHandProducts	<a href="#ExistingStockLineModel"><b>[] Existing Stock Line Model</b></a>			Lines will be automaticaly filled with filtered products which have non empty stock.(in POST method or in PUT method if status is changing from DRAFT to IN PROGRESS. If current Status is IN PROGRESS, lines can be passed to PUT method (line entry will be found by product, location, batchSN and expiryDate and passed Adjustment will be set as new Adjustment)	  
	ZeroStockOnHandProducts	<a href="#NewStockLineModel"><b>[] New Stock Line Model</b></a>				  
	Transactions	<a href="#TransactionStockLineModel"><b>[] Transaction Stock Line Model</b></a>			Created transactions due to stock changes.	      

### Available Fields for Stock Transfer
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes*	Unique Cin7 Core Stock Transfer ID.Required for PUT	
	From	string		Yes*	ID of location from which transfer was done. Required if FromLocation is null	
	FromLocation	string		Yes*	Name of location from which transfer was done. Required if From is null	
	To	string		Yes*	ID of location to which transfer was done. Required if ToLocation is null	
	ToLocation	string		Yes*	Name of location to which transfer was done. Required if To is null	
	Status	string		Yes	Status of stock transfer. Available values are DRAFT, IN TRANSIT,COMPLETED,VOIDED.	
	Number	string			Stock Transfer Number (auto-generated).	
	CostDistributionType	string			Type of additional journals capitalization cost distribution. Available values are Cost, Quantity,Weight,Volume. Default value is Cost	
	InTransitAccount	string		Yes*	Asset Account holding stock value while it is in transit. Required if Status is IN TRANSIT.	
	DepartureDate	string		Yes*	Date when stock left the From warehouse. Required if Status is IN TRANSIT.	
	CompletionDate	string		Yes	Date of transaction	
	RequiredByDate	string			Date when stock is required to be delivered to destination warehouse.	
	Reference	string			Custom reference	
	SkipOrder	boolean			Skip stock transfer order. Default value true.	
	Lines	<a href="#StockTransferLineModel"><b>[] Stock Transfer Line Model</b></a>		Yes		
	Order	<a href="#StockTransferOrderModel"><b>Stock Transfer Order Model</b></a>				

### Available Fields for Stock Transfer Order
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Stock Transfer ID	
	Status	string		Yes	Status of stock transfer order. Available values are DRAFT, AUTHORISED, NOT AVAILABLE.	
	Lines	<a href="#StockTransferOrderLineModel"><b>[] Stock Transfer Order Line Model</b></a>		Yes		

### Available Fields for Supplier
	 Property              	 Type      	 Length    	 Required  	 Notes                                                             	
	:----------------------	-----------	-----------	-----------	-------------------------------------------------------------------	
	 ID                  	 string      	           	 Yes       	 Unique Supplier ID.                                               	
	 Name                	 string    	 256       	 Yes       	 Name of Supplier.                                                 	
	 Currency            	 string    	           	 Yes       	 Currency code of Supplier.                                        	
	 PaymentTerm         	 string    	           	 Yes       	 PaymentTerm name.                                                 	
	 AccountPayable      	 string    	           	 Yes       	 AccountPayable code of Supplier.                                  	
	 TaxRule             	 string    	           	 Yes       	 TaxRule name.                                                     	
	 Status              	 string    	           	 Yes*      	 Points that Supplier is Active. Can be Active or Deprecated   	
	 Discount            	 int32      	           	           	 Discount must be between 0% and 100%. 0% used as default.         	
	 Comments            	 string    	 256       	           	 Comments.                                                         	
	 TaxNumber           	 int32      	           	           	 TaxNumber.                                                        	
	 AttributeSet        	 string    	           	           	 AttributeSet name.                                                	
	 AdditionalAttribute#	 string    	           	           	 Additional attribute value. # - int(1-10).                        	
	 LastModifiedOn      	 string    	           	           	 Date of last modification.                                        	
	 Addresses           	<a href="#SupplierAddressModel"><b>[] Supplier Address Model</b></a>			 List of addresses.             	
	 Contacts            	<a href="#SupplierContactModel"><b>[] Supplier Contact Model</b></a>			 List of contacts.              	

### Available Fields for Tax
	 Property              	 Type      	 Length    	 Required  	 Notes                 	
	:----------------------	-----------	-----------	-----------	-----------------------	
	 ID                  	 string      	 50        	           	 Unique ID             	
	 Name                	 string    	 50        	 Yes       	 Tax name              	
	 Account             	 string    	           	 Yes       	 ChartOfAccount Code with Class == LIABILITY && Status == ACTIVE 	
	 IsActive            	 bool      	           	 Yes       	 Points that tax is Active 	
	 TaxInclusive        	 bool      	           	 Yes       	 Points that tax is Inclusive 	
	 TaxPercent          	 float32   	           	           	 Tax percentage, should be between 0 and 100. Read-only 	
	 IsTaxForSale        	 bool      	           	           	 Points that tax is used for Sale 	
	 IsTaxForPurchase    	 bool      	           	           	 Points that tax is used for Purchase 	
	 Components 	<a href="#TaxComponentModel"><b>[] Tax Component Model</b></a>				

### Available Fields for Unit of Measure
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 ID  	 string           	         	          	Required for PUT, Ignored for POST operations          	
	 Name  	 string             	 50        	 Yes         	Unit of Measure name       	

### Available Fields for Lead
	Property               	Type       	Length     	Required   	Notes	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	ID	string		yes	Unique lead ID generated by Cin7 Core.	
	LeadStatus	string	256	yes	Lead status. Available values are NEW, ATTEMPTING TO CONTACT, QUALIFIED, DISQUALIFIED.	
	Name	string	256	yes	Name of the lead.	
	Currency	string	3	yes	Currency code of the lead.	
	PaymentTerm	int		yes	Payment term for the lead.	
	PriceTier	string	50	yes	Price tier.	
	SalesRepresentative	string	256	yes	Sale representative name.	
	TaxRule	string	255	yes	Tax rule name for the lead.	
	Amount	decimal			Lead value.	
	CloseChance	int		yes	 % chance of wining a lead. Available values are 0%, 10%, 20%, 30%..... etc.	
	CloseDate	string		yes	Lead deal estimated close date.	
	Comments	string	2000		Additional comments.	
	Contacts	<a href="#SupplierContactModel"><b>[] Lead Contact Model</b></a>			List of contacts.	
	Addresses	<a href="#SupplierAddressModel"><b>[] Lead Address Model</b></a>			List of addresses.	

### Available Fields for Opportunity
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	ID	string		yes	Unique Cin7 Core opportunity ID.	
	CustomerID	string			Customer identifier.	
	LeadID	string			Lead identifier.	
	CustomerName	string	256	yes	Lead or customer name.	
	Contact	string	256		Lead or customer contact name.	
	Phone	string	50		Lead or customer phone number.	
	DefaultAccount	string	50		Account code used by default for invoice lines when no revenue account is defined on Product. By default it is set to default revenue account.	
	BillingAddressLine1	string	256	yes	Opportunity billing address line 1.	
	BillingAddressLine2	string	256		Opportunity billing address line 2.	
	Currency	string	3	yes	Opportunity Currency.	
	TaxPercent	decimal			Opportunity Tax %.	
	TaxInclusive	bool			Opportunity is tax Inclusive or Exclusive.	
	TaxRule	string	255	yes	Default Tax Rule name selected for opportunity.	
	TermDays	int			Opportunity payment term number of days.	
	Terms	string	50	yes	Opportunity payment term name.	
	PriceTier	string	50	yes	Price Tier name selected for opportunity.	
	ShippingAddressLine1	string	256		Opportunity shipping address line 1.	
	ShippingAddressLine2	string	256		Opportunity shipping address line 2.	
	OpportunityLocation	string	256	yes	Default location to pick stock from once converted to sale.	
	OpportunityNumber	string	256		Number of opportunities.	
	OpportunityDate	string			Date opportunity is created.	
	OpportunityComment	string	1024		Additional comments added on opportunity.	
	OpportunityMemo	string	1024		Additional information for the opportunity.	
	OpportunityStatus	string	25		Status of the opportunity. Possible values: Draft, In Progress, Won, Lost.	
	TaxTotal	decimal			Total tax amount for the opportunity.	
	Total	decimal			Total opportunity.	
	CustomerReference	string	256		Customer reference.	
	CurrencyConversionRate	decimal			Currency conversion rate.	
	CustomerCurrency	string	3	yes	Customer Currency.	
	TermMethod	int		yes		
	SalesRepresentative	string	256	yes	Sales Representative name.	
	TermDueNextMonth	int				
	ShipToOther	bool		yes	If false and customer or leads' address matching Line 1 not found, then a new customer/ lead shipping address will be created.	
	ShipToCompany	string	128		Ship to company name when ShipToOther is true.	
	ShipToContact	string	512		Ship to contact name when ShipToOther is true.	
	ShipToCountry	string	64		Ship to country when ShipToOther is true.	
	ShipToPostCode	string	32		Ship to post code when ShipToOther is true.	
	ShipToState	string	128		Ship to state when ShipToOther is true.	
	ShipToCity	string	128		Ship to city when ShipToOther is true.	
	ShipToAddress1	string	256		Ship to address line 1.	
	ShipToAddress2	string	256		Ship to address line 2.	
	CustomField1	string	256		Value of opportunity additional attribute 1.	
	CustomField2	string	256		Value of opportunity additional attribute 2.	
	CustomField3	string	256		Value of opportunity additional attribute 3.	
	CustomField4	string	256		Value of opportunity additional attribute 4.	
	CustomField5	string	256		Value of opportunity additional attribute 5.	
	CustomField6	string	256		Value of opportunity additional attribute 6.	
	CustomField7	string	256		Value of opportunity additional attribute 7.	
	CustomField8	string	256		Value of opportunity additional attribute 8.	
	CustomField9	string	256		Value of opportunity additional attribute 9.	
	CustomField10	string	256		Value of opportunity additional attribute 10.	
	Lines	Opportunity Line[]			Opportunity line items.	
	AdditionalCharges	Opportunity Additional Charge[]			Opportunity line additional charges. If true then additional charges lines displayed in Lines array.	

### Available Fields for Opportunity Line
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string			Unique Opportunity ID.
ProductID	string		Either ProductID or ProductName or ProductSKU should be presented	Product identifier referenced by this line.
ProductName	string	256		Product Name referenced by this line.
ProductSKU	string	50		Product SKU referenced by this line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false
Quantity	decimal		yes	Product or service quantity. Minimal value is 1.
Price	decimal		yes	Price per unit in Customer/ lead currency.
Discount	decimal			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0.
Tax	decimal		yes	Tax applicable for the opportunity line item.
Total	decimal		yes	Line Total. For validation.
TaxRule	string	255		Line Tax Rule name.
TaxPercent	decimal			Tax percentage.
Comment	string	1024		Additional comments on line item.

### Available Fields for Opportunity Opportunity Additional Charge 
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string		yes	Unique opportunity ID
Description	string	256	yes	Name of Service Product referenced by this line.
Quantity	decimal		yes	Product or service quantity. Minimal value is 1.
Amount	decimal		yes	Price per unit in Customer currency.
Discount	decimal			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0.
Tax	decimal		yes	Tax  amount for the line item.
Total	decimal		yes	Line Total. For validation.
TaxRule	string	255		Line Tax Rule name.
TaxPercent	decimal			Tax percentage for line item.
Comment	string	1024		Additional comments on line item.

### Available Fields for Task model

### Available Fields for task Category model

### Available Fields for Workflow model

### Available Fields for WorkflowStep model

### Available Fields for Opportunity Opportunity Additional Charge 
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string		yes	Unique opportunity ID
Description	string	256	yes	Name of Service Product referenced by this line.
Quantity	decimal		yes	Product or service quantity. Minimal value is 1.
Amount	decimal		yes	Price per unit in Customer currency.
Discount	decimal			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0.
Tax	decimal		yes	Tax  amount for the line item.
Total	decimal		yes	Line Total. For validation.
TaxRule	string	255		Line Tax Rule name.
TaxPercent	decimal			Tax percentage for line item.
Comment	string	1024		Additional comments on line item.

### Available Fields for Task model
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string		yes	
Name	string	256	yes	
Description	string			
StartDate	string		yes	
EndDate	string		yes	
IsAllDay	bool			
CompletedDate	string			
Category	string	256		
WorkflowName	string	256		
AssignedTo	string	256		
EntityType	string	64	yes	See available types below
EntityID	string		yes	
IsImportant	bool			
TaskStatus	string	64	yes	
AssignedBy	string	256		
VoidDate	string			

### Available Fields for task Category model
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string		yes	
Name	string	256	yes	
Color	string	6		RGB HEX color string, e.g. F7CAFE
BackgroundColor	string	6		RGB HEX color string, e.g. F7CAFE

### Available Fields for Workflow model
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string		yes	
Name	string	256	yes	
IsActive	bool			
EntityType	string	64	yes	same types as for Task
DueDaysType	string	64	yes	can be “COMPLETED_DATE” or “START_DATE”
Steps	WorkflowStep[]			

### Available Fields for WorkflowStep model
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string		yes	
Name	string	256	yes	
Category	string	256		
Days	int			
StartTime	string	8		Time in ISO 8601 format hh:mm:ss
EndTime	string	8		If Start Time and End Time are the same and equal to 0:00, a task will be created for an entire day.
UserName	string	256		
SkipHoliday	string	32	yes	Available values: move_next, move_previous, leave, ph_move_next, ph_move_previous

SkipHoliday values mean:

## New Stock Line Model <a name="NewStockLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	GUID		Yes*	Product identifier referenced by this Line. Required if SKU is empty.	
	SKU	string	50	Yes*	Product SKU referenced by this Line. Required if ProductID is empty.	
	ProductName	string	1024		Product Name referenced by this Line	
	Quantity	float32		Yes	New value for QuantityOnHand	
	UnitCost	float32		Yes	Cost of 1 product unit.	
	LocationID	GUID		Yes*	Id of location on which stock changes will be done. Required if Location is empty.	
	Location	string	256	Yes*	Name of location on which stock changes will be done. Required if LocationID is empty.	
	BatchSN	string	50	Yes*	Batch Serial Number. Required if Product's Costing Method is not FIFO.	
	ExpiryDate	string			Date when Batch expires. Required if Product's Costing Method is FESN	
	ReceivedDate	string			Date when new stock is received	
	Comments	string	256		Comment	
	ProductLength	float32			Length of product. Read-only.	
	ProductWidth	float32			Width of product. Read-only.	
	ProductHeight	float32			Height of product. Read-only.	
	ProductWeight	float32			Weight of product. Read-only.	
	 WeightUnits  	 string             	 10        	          	Unit of measure for unit weight. Available Values are <a href="#DimensionUnitAvailableValues"><b>values</b></a>	
	 DimensionsUnits  	 string             	 10        	      	Unit of measure for unit length/width/height. Available Values are <a href="#DimensionUnitAvailableValues"><b>values</b></a>	
	ProductCustomField1	string			Value of Product's additional attribute 1. Read-only.	
	ProductCustomField2	string			Value of Product's additional attribute 2. Read-only.	
	ProductCustomField3	string			Value of Product's additional attribute 3. Read-only.	
	ProductCustomField4	string			Value of Product's additional attribute 4. Read-only.	
	ProductCustomField5	string			Value of Product's additional attribute 5. Read-only.	
	ProductCustomField6	string			Value of Product's additional attribute 6. Read-only.	
	ProductCustomField7	string			Value of Product's additional attribute 7. Read-only.	
	ProductCustomField8	string			Value of Product's additional attribute 8. Read-only.	
	ProductCustomField9	string			Value of Product's additional attribute 9. Read-only.	
	ProductCustomField10	string			Value of Product's additional attribute 10. Read-only.	
	Image	<a href="#AttachmentLineModel"><b>Image</b></a>			Product Image. Read-only. this field is available for stock take	
	Barcode	string			Value of Product's Barcode. Read-only.this field is available for stock take	
	StockLocator	string			Value of Product's StockLocator. Read-only.this field is available for stock take	
	Unit	string			Value of Product's Default Unit of Measure. Read-only.this field is available for stock take	
	CostingMethod	string			Value of Product's Costing method. Read-only.this field is available for stock take. Valid values: FIFO, Special - Batch, Special - Serial Number, FIFO - Serial Number, FIFO - Batch, FEFO - Batch, FEFO - Serial Number 	

### Available Fields for Unit of Measure
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 ID  	 string           	         	          	Required for PUT, Ignored for POST operations          	
	 Name  	 string             	 50        	 Yes         	Unit of Measure name       	

### Available Fields for Opportunity
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	ID	string		yes	Unique Cin7 Core opportunity ID.	
	CustomerID	string			Customer identifier.	
	LeadID	string			Lead identifier.	
	CustomerName	string	256	yes	Lead or customer name.	
	Contact	string	256		Lead or customer contact name.	
	Phone	string	50		Lead or customer phone number.	
	DefaultAccount	string	50		Account code used by default for invoice lines when no revenue account is defined on Product. By default it is set to default revenue account.	
	BillingAddressLine1	string	256	yes	Opportunity billing address line 1.	
	BillingAddressLine2	string	256		Opportunity billing address line 2.	
	Currency	string	3	yes	Opportunity Currency.	
	TaxPercent	decimal			Opportunity Tax %.	
	TaxInclusive	bool			Opportunity is tax Inclusive or Exclusive.	
	TaxRule	string	255	yes	Default Tax Rule name selected for opportunity.	
	TermDays	int			Opportunity payment term number of days.	
	Terms	string	50	yes	Opportunity payment term name.	
	PriceTier	string	50	yes	Price Tier name selected for opportunity.	
	ShippingAddressLine1	string	256		Opportunity shipping address line 1.	
	ShippingAddressLine2	string	256		Opportunity shipping address line 2.	
	OpportunityLocation	string	256	yes	Default location to pick stock from once converted to sale.	
	OpportunityNumber	string	256		Number of opportunities.	
	OpportunityDate	string			Date opportunity is created.	
	OpportunityComment	string	1024		Additional comments added on opportunity.	
	OpportunityMemo	string	1024		Additional information for the opportunity.	
	OpportunityStatus	string	25		Status of the opportunity. Possible values: Draft, In Progress, Won, Lost.	
	TaxTotal	decimal			Total tax amount for the opportunity.	
	Total	decimal			Total opportunity.	
	CustomerReference	string	256		Customer reference.	
	CurrencyConversionRate	decimal			Currency conversion rate.	
	CustomerCurrency	string	3	yes	Customer Currency.	
	TermMethod	int		yes		
	SalesRepresentative	string	256	yes	Sales Representative name.	
	TermDueNextMonth	int				
	ShipToOther	bool		yes	If false and customer or leads' address matching Line 1 not found, then a new customer/ lead shipping address will be created.	
	ShipToCompany	string	128		Ship to company name when ShipToOther is true.	
	ShipToContact	string	512		Ship to contact name when ShipToOther is true.	
	ShipToCountry	string	64		Ship to country when ShipToOther is true.	
	ShipToPostCode	string	32		Ship to post code when ShipToOther is true.	
	ShipToState	string	128		Ship to state when ShipToOther is true.	
	ShipToCity	string	128		Ship to city when ShipToOther is true.	
	ShipToAddress1	string	256		Ship to address line 1.	
	ShipToAddress2	string	256		Ship to address line 2.	
	CustomField1	string	256		Value of opportunity additional attribute 1.	
	CustomField2	string	256		Value of opportunity additional attribute 2.	
	CustomField3	string	256		Value of opportunity additional attribute 3.	
	CustomField4	string	256		Value of opportunity additional attribute 4.	
	CustomField5	string	256		Value of opportunity additional attribute 5.	
	CustomField6	string	256		Value of opportunity additional attribute 6.	
	CustomField7	string	256		Value of opportunity additional attribute 7.	
	CustomField8	string	256		Value of opportunity additional attribute 8.	
	CustomField9	string	256		Value of opportunity additional attribute 9.	
	CustomField10	string	256		Value of opportunity additional attribute 10.	
	Lines	Opportunity Line[]			Opportunity line items.	
	AdditionalCharges	Opportunity Additional Charge[]			Opportunity line additional charges. If true then additional charges lines displayed in Lines array.	

### Available Fields for Opportunity Line
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string			Unique Opportunity ID.
ProductID	string		Either ProductID or ProductName or ProductSKU should be presented	Product identifier referenced by this line.
ProductName	string	256		Product Name referenced by this line.
ProductSKU	string	50		Product SKU referenced by this line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false
Quantity	decimal		yes	Product or service quantity. Minimal value is 1.
Price	decimal		yes	Price per unit in Customer/ lead currency.
Discount	decimal			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0.
Tax	decimal		yes	Tax applicable for the opportunity line item.
Total	decimal		yes	Line Total. For validation.
TaxRule	string	255		Line Tax Rule name.
TaxPercent	decimal			Tax percentage.
Comment	string	1024		Additional comments on line item.

### Available Fields for Opportunity Opportunity Additional Charge 
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string		yes	Unique opportunity ID
Description	string	256	yes	Name of Service Product referenced by this line.
Quantity	decimal		yes	Product or service quantity. Minimal value is 1.
Amount	decimal		yes	Price per unit in Customer currency.
Discount	decimal			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0.
Tax	decimal		yes	Tax  amount for the line item.
Total	decimal		yes	Line Total. For validation.
TaxRule	string	255		Line Tax Rule name.
TaxPercent	decimal			Tax percentage for line item.
Comment	string	1024		Additional comments on line item.

### Available Fields for Task model
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string		yes	
Name	string	256	yes	
Description	string			
StartDate	string		yes	
EndDate	string		yes	
IsAllDay	bool			
CompletedDate	string			
Category	string	256		
WorkflowName	string	256		
AssignedTo	string	256		
EntityType	string	64	yes	See available types below
EntityID	string		yes	
IsImportant	bool			
TaskStatus	string	64	yes	
AssignedBy	string	256		
VoidDate	string			

### Available Fields for task Category model
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string		yes	
Name	string	256	yes	
Color	string	6		RGB HEX color string, e.g. F7CAFE
BackgroundColor	string	6		RGB HEX color string, e.g. F7CAFE

### Available Fields for Workflow model
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string		yes	
Name	string	256	yes	
IsActive	bool			
EntityType	string	64	yes	same types as for Task
DueDaysType	string	64	yes	can be “COMPLETED_DATE” or “START_DATE”
Steps	WorkflowStep[]			

### Available Fields for WorkflowStep model
	 Property              	 Type      	 Length    	 Required  	 Notes                                                         	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
ID	string		yes	
Name	string	256	yes	
Category	string	256		
Days	int			
StartTime	string	8		Time in ISO 8601 format hh:mm:ss
EndTime	string	8		If Start Time and End Time are the same and equal to 0:00, a task will be created for an entire day.
UserName	string	256		
SkipHoliday	string	32	yes	Available values: move_next, move_previous, leave, ph_move_next, ph_move_previous

## Dimension Unit Available Values <a name="DimensionUnitAvailableValues" />
	Property	Type	Abbreviation	Name	
	:-------	----	------	--------	-----	
	WeightUnits	weight	oz	ounces	
	WeightUnits	weight	mg	miligramm	
	WeightUnits	weight	kg	kilogramm	
	WeightUnits	weight	lb	pound	
	WeightUnits	weight	g	gram	
	DimensionsUnits	length	m	metre	
	DimensionsUnits	length	cm	centimetre	
	DimensionsUnits	length	mi	mile	
	DimensionsUnits	length	mm	millimetre	
	DimensionsUnits	length	in	inch	
	DimensionsUnits	length	ft	foot	
	DimensionsUnits	length	yd	yard	
	DimensionsUnits	length	km	kilometre	

## Address Model <a name="AddressModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	GUID		No	Address ID, optional, applicable for sales only, POST/PUT methods. If specified then existing address is taken by ID.	
	DisplayAddressLine1	string	256		Address Line 1 as displayed on Sale form. = Line1 + Line2	
	DisplayAddressLine2	string	256		Address Line 2 as displayed on Sale form. = City + State/Region + Zip/Postcode + Country	
	Line1	string	256	Yes	Address Line 1	
	Line2	string	256		Address Line 2	
	City	string	256		City	
	State	string	256		State	
	Postcode	string	20		Post code	
	Country	string	256	Yes	Country	

## Sale Shipping Address Model <a name="SaleShippingAddressModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	GUID		No	Address ID, optional. If specified then existing address is taken by ID. Ignored if ShipToOther = true	
	DisplayAddressLine1	string	256		Address Line 1 as displayed on Sale form. = Line1 + Line2	
	DisplayAddressLine2	string	256		Address Line 2 as displayed on Sale form. = City + State/Region + Zip/Postcode + Country	
	Line1	string	256	Yes	Address Line 1	
	Line2	string	256		Address Line 2	
	City	string	256		City	
	State	string	256		State	
	Postcode	string	20		Post code	
	Country	string	256	Yes	Country	
	Company	string	128		Company Name	
	Contact	string	512		Contact	
	ShipToOther	boolean			If false and Customer's address matching Line 1 not found, than new customer shipping address will be created.	

## Purchase Shipping Address Model <a name="PurchaseShippingAddressModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	DisplayAddressLine1	string	256		Address Line 1 as displayed on Sale form. = Line1 + Line2	
	DisplayAddressLine2	string	256		Address Line 2 as displayed on Sale form. = City + State/Region + Zip/Postcode + Country	
	Line1	string	256	Yes	Address Line 1	
	Line2	string	256		Address Line 2	
	City	string	256		City	
	State	string	256		State	
	Postcode	string	20		Post code	
	Country	string	256	Yes	Country	
	ShipToOther	boolean			Ship to different company	
	Company	string	128		Company name when ShipToOther is set to True	

## Additional Attribute Model <a name="AdditionalAttributeModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	AdditionalAttribute1	string	256		Additional Attribute 1 value	
	AdditionalAttribute2	string	256		Additional Attribute 2 value	
	AdditionalAttribute3	string	256		Additional Attribute 3 value	
	AdditionalAttribute4	string	256		Additional Attribute 4 value	
	AdditionalAttribute5	string	256		Additional Attribute 5 value	
	AdditionalAttribute6	string	256		Additional Attribute 6 value	
	AdditionalAttribute7	string	256		Additional Attribute 7 value	
	AdditionalAttribute8	string	256		Additional Attribute 8 value	
	AdditionalAttribute9	string	256		Additional Attribute 9 value	
	AdditionalAttribute10	string	256		Additional Attribute 10 value	

## Sale Quote Model <a name="SaleQuoteModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Memo	string	1024	Yes	Additional information for Quote	
	Status	string		Yes	Quote Status. Possible Values are <a href="#QuoteStatuses"><b>values</b></a>	
	Prepayments	<a href="#SalePaymentLineModel"><b>[] Sale Payment Line Model</b></a>				
	Lines	<a href="#SaleQuoteLineModel"><b>[] Sale Quote Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#SaleAdditionalChargeModel"><b>[] Sale Additional Charge Model</b></a>				
	TotalBeforeTax	float32 with up to 4 decimal places		Yes	Sum of quote lines and additional charges without taxes	
	Tax	float32 with up to 4 decimal places		Yes	Sum of quote lines and additional charges taxes	
	Total	float32 with up to 4 decimal places		Yes	Sum of quote lines and additional charges with taxes	

## Sale Quote Line Model <a name="SaleQuoteLineModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	GUID		Yes*	Product identifier referenced by this Line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false	
	SKU	string	50	Yes*	Product SKU referenced by this Line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false	
	Name	string	1024	Yes	Product Name referenced by this Line	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	Price	float32 with up to 4 decimal places		Yes	Price per unit in Customer currency	
	Discount	float32 with up to 2 decimal places			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0	
	Tax	float32 with up to 4 decimal places		Yes	Tax	
	AverageCost	float32 with up to 4 decimal places		Yes	Average product cost	
	TaxRule	string	50	Yes	Line Tax Rule name.	
	Comment	string	256	Yes	Comment for this line	
	Total	float32 with up to 4 decimal places			Line Total.For validation	

## Sale Order Model <a name="SaleOrderModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	SaleOrderNumber	string	256		Sale Order Number (auto-generated)	
	Memo	string	1024	Yes	Additional information for Quote	
	Status	string		Yes	Quote Status. Possible Values are <a href="#QuoteStatuses"><b>values</b></a>	
	Lines	<a href="#SaleOrderLineModel"><b>[] Sale Quote Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#SaleAdditionalChargeModel"><b>[] Sale Additional Charge Model</b></a>				
	TotalBeforeTax	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges without taxes	
	Tax	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges taxes	
	Total	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges with taxes	

## Sale Order Line Model <a name="SaleOrderLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	GUID		Yes*	Product identifier referenced by this Line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false	
	SKU	string	50	Yes*	Product SKU referenced by this Line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false	
	Name	string	1024	Yes	Product Name referenced by this Line	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	Price	float32 with up to 4 decimal places		Yes	Price per unit in Customer currency	
	Discount	float32 with up to 2 decimal places			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0	
	Tax	float32 with up to 4 decimal places		Yes	Tax	
	AverageCost	float32 with up to 4 decimal places			Average product cost	
	TaxRule	string	50	Yes	Line Tax Rule name.	
	Comment	string	256		Comment for this line	
	DropShip	boolean			Required if product is not a service with “Optional Drop Ship” mode. Otherwise it ignored.	
	BackorderQuantity	float32 with up to 4 decimal places			Quantity of the ordered product on backorder. Read Only for POST	
	Total	float32 with up to 4 decimal places			Line Total.For validation	

## Sale Fulfilment Model <a name="SaleFulfilmentModel"/>
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	GUID		Yes	Identifier of sale fulfilment task	
	FulfillmentNumber	int32	Yes	Fulfillment Number (auto-generated)	
	LinkedInvoiceNumber	string			Invoice number linked to this fulfilment	
	FulFilmentStatus	string		Yes	Fulfilment status. Possible Values are NOT AVAILABLE, NOT FULFILLED, FULFILLED, PARTIALLY FULFILLED, VOIDED	
	Pick	<a href="#SaleFulfilmentPickPackModel"><b>Sale Fulfilment Pick Pack Model</b></a>				
	Pack	<a href="#SaleFulfilmentPickPackModel"><b>Sale Fulfilment Pick Pack Model</b></a>				
	Ship	<a href="#SaleFulfilmentShipModel"><b>Sale Fulfilment Ship Model</b></a>				

## Sale Fulfilment Pick Pack Model <a name="SaleFulfilmentPickPackModel"/>
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Status	string		Yes	Pick/Pack status. Possible Values are NOT AVAILABLE, DRAFT, AUTHORISED, VOIDED	
	Lines	<a href="#SaleFulfilmentPickPackLineModel"><b>[] Sale Fulfilment Pick Pack Line Model</b></a>				

## Sale Fulfilment Pick Pack Line Model <a name="SaleFulfilmentPickPackLineModel"/>
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	GUID			Product identifier referenced by this Line.	
	SKU	string	50	Yes	Product SKU referenced by this Line.	
	Name	string	1024		Product Name referenced by this Line	
	Location	string	256	Yes*	Location from which pick product. Required if LocationID is empty	
	LocationID	GUID		Yes*	LocationID from which pick product. Required if Location is empty	
	Quantity	float32 with up to 4 decimal places		Yes	Product quantity. Minimal value is 1.	
	BatchSN	string	256		Batch Serial Number	
	ExpiryDate	string			Expiry Date	
	Box	string	256		Pack box name. This field available only for packing.	
	NonInventory	boolean			Indication that this line relates to non-inventory product and won't involve any stock movement.	
	WarrantyRegistrationNumber	string	256		Warranty registration number. This field available only for packing.	
	RestockLocation	string	256	No	Used only for CreditNote. Location name of return of the product.	
	RestockLocationID	GUID		No	Used only for CreditNote. Location ID of return of the product.	

## Sale Fulfilment Ship Model <a name="SaleFulfilmentShipModel"/>
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Status	string		Yes	Ship status. Possible Values are NOT AVAILABLE, DRAFT, PARTIALLY AUTHORISED, AUTHORISED, VOIDED	
	RequireBy	Date			Require By Date	
	ShippingAddress	<a href="#SaleShippingAddressModel"><b>Shipping Address Model</b></a>			Shipping Address	
	ShippingNotes	string	1024		Shipping Notes	
	Lines	<a href="#SaleFulfilmentShipLineModel"><b>[] Sale Fulfilment Ship Line Model</b></a>				

## Sale Fulfilment Ship Line Model <a name="SaleFulfilmentShipLineModel"/>
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	GUID			Shipment identifier. Not Required for POST/PUT	
	ShipmentDate	string		Yes	Shipment Date	
	Carrier	string	256		Carrier referenced by this Line.	
	Boxes	string	256	Yes	Boxes. For POST/PUT methods, this field's name is Box	
	TrackingNumber	string	256		Shipment tracking number.	
	TrackingURL	string	512		Shipment tracking URL. Read-only.	
	IsShipped	boolean			Is package shipped.	

## Sale Invoice Model <a name="SaleInvoiceModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	GUID		Yes	Identifier of sale Invoice task	
	InvoiceNumber	string			Invoice Number (auto-generated)	
	Memo	string	1024		Additional information for Invoice.	
	Status	string		Yes	Invoice status. Possible Values are <a href="#InvoiceStatuses"><b>values</b></a>. For POST available values are DRAFT, AUTHORISED	
	InvoiceDate	string		Yes	Invoice Date.	
	InvoiceDueDate	string		Yes	Invoice Due Date.	
	CurrencyConversionRate	float32 with up to 4 decimal places				
	BillingAddressLine1	string	256		Billing Address Line 1	
	BillingAddressLine2	string	256		Billing Address Line 2	
	LinkedFulfillmentNumber	int32			Number of Fulfilment linked to this invoice	
	Lines	<a href="#SaleInvoiceLineModel"><b>[] Sale Invoice Line Model</b></a>				
	AdditionalCharges	<a href="#SaleInvoiceAdditionalChargeModel"><b>[] Sale Invoice Additional Charge Model</b></a>				
	Payments	<a href="#SalePaymentLineModel"><b>[] Sale Payment Line Model</b></a>				
	TotalBeforeTax	float32 with up to 4 decimal places			Sum of Invoice lines and additional charges without taxes	
	Tax	float32 with up to 4 decimal places			Sum of Invoice lines and additional charges taxes	
	Total	float32 with up to 4 decimal places			Sum of Invoice lines and additional charges with taxes	
	Paid	float32 with up to 4 decimal places			Sum of payments.	

## Sale Additional Charge Model <a name="SaleAdditionalChargeModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Description	string	256	Yes	Name of Service Product referenced by this Line	
	Price	float32 with up to 4 decimal places		Yes	Price per unit in Customer currency	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	Discount	float32 with up to 2 decimal places			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0	
	Tax	float32 with up to 4 decimal places		Yes	Tax	
	Total	float32 with up to 4 decimal places			Line Total. For validation	
	TaxRule	string	50	Yes	Line Tax Rule name.	
	Comment	string	1024		Comment	

## Sale Invoice Additional Charge Model <a name="SaleInvoiceAdditionalChargeModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Description	string	256	Yes	Name of Service Product referenced by this Line	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	Price	float32 with up to 4 decimal places		Yes	Price per unit in Customer currency	
	Discount	float32 with up to 2 decimal places			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0	
	Tax	float32 with up to 4 decimal places		Yes	Tax	
	Total	float32 with up to 4 decimal places			Line Total.For validation	
	TaxRule	string	50	Yes	Line Tax Rule name.	
	Account	string	50		Revenue account	
	Comment	string	1024		Comment	

## Sale Invoice Line Model <a name="SaleInvoiceLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	GUID		Yes*	Product identifier referenced by this Line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false	
	SKU	string	50	Yes*	Product SKU referenced by this Line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false	
	Name	string	1024	Yes	Product Name referenced by this Line	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	Price	float32 with up to 4 decimal places		Yes	Price per unit in Customer currency	
	Discount	float32 with up to 2 decimal places			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0	
	Tax	float32 with up to 4 decimal places		Yes	Tax	
	Total	float32 with up to 4 decimal places		Yes	Line Total. For validation	
	AverageCost	float32 with up to 4 decimal places			Average product cost	
	TaxRule	string	50	Yes	Line Tax Rule name.	
	Account	string	50		Revenue account	
	Comment	string	256		Comment for this line	

## Sale Credit Note Model <a name="SaleCreditNoteModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	GUID		Yes	Identifier of sale Invoice task	
	CreditNoteInvoiceNumber	string		Yes	Number of Invoice linked to this Credit Note	
	Memo	string	1024		Additional information for Credit Note.	
	Status	string		Yes	Credit Note status. Possible Values are NOT AVAILABLE, DRAFT, AUTHORISED, VOIDED. For POST available values are DRAFT, AUTHORISED	
	CreditNoteDate	string		Yes	Credit Note Date.	
	CreditNoteNumber	string			Credit Note Number (auto-generated)	
	CreditNoteConversionRate	float32 with up to 4 decimal places				
	Lines	<a href="#SaleInvoiceLineModel"><b>[] Sale Invoice Line Model</b></a>				
	AdditionalCharges	<a href="#SaleInvoiceAdditionalChargeModel"><b>[] Sale Invoice Additional Charge Model</b></a>				
	Refunds	<a href="#SalePaymentLineModel"><b>[] Sale Payment Line Model</b></a>				
	Restock	<a href="#SaleFulfilmentPickPackLineModel"><b>[] Sale Fulfilment Pick Pack Line Model</b></a>				
	TotalBeforeTax	float32 with up to 4 decimal places			Sum of Credit Note lines and additional charges without taxes	
	Tax	float32 with up to 4 decimal places			Sum of Credit Note lines and additional charges taxes	
	Total	float32 with up to 4 decimal places			Sum of Credit Note lines and additional charges with taxes	

## Sale Payment Line Model <a name="SalePaymentLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	GUID			Identifier of payment	
	Reference	string			Payment reference number	
	Amount	float32 with up to 2 decimal places			Payment amount in customer currency	
	DatePaid	string			Date when payment has been made	
	Account	string			Account Code of the bank/payment account from Chart of accounts	
	CurrencyRate	float32 with up to 4 decimal places			Currency Conversion rate expressed as number of Base currency units for one Customer currency unit.	
	DateCreated	string			Date of creation payment record.	

## Sale Manual Journal Model <a name="SaleManualJournalModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Status	string		Yes	Manual Journal Status. Available values are NOT AVAILABLE, DRAFT, AUTHORISED. Available values for POST are DRAFT, AUTHORISED	
	Lines	<a href="#SaleManualJournalLineModel"><b>[] Sale Manual Journal Line Model</b></a>				

## Sale Manual Journal Line Model <a name="SaleManualJournalLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Reference	string			Reference Service Product	
	Amount	float32 with up to 2 decimal places		Yes	Amount in base currency for manual journal	
	Date	string		Yes	Effective date of manual journal	
	Debit	string		Yes	Debit Account	
	Credit	string		Yes	Credit Account	

## Attachment Line Model <a name="AttachmentLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	GUID			Attached file ID	
	ContentType	string			Attached file content type	
	IsDefault	boolean			Is this attachment will be used as default product image. Is visible only for product Attachment Endpoint.	
	FileName	string			Attached file name	
	DownloadUrl	string			Url to download attached File	

## Product Family Product Line Model <a name="ProductFamilyProductLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	GUID		Yes	Unique Product ID	
	SKU	string	50		Must be unique across all products. Ignored in POST and PUT	
	Name	string	256		Product variation name. Ignored in POST and PUT	
	Option1	string	256	Yes	The value that corresponds to the Option 1. Should be included in Option1Values.	
	Option2	string	256		The value that corresponds to the Option 2. Should be included in Option2Values. Should be empty if Option 2 wasn’t set.	
	Option3	string	256		The value that corresponds to the Option 2. Should be included in Option3Values. Should be empty if Option 3 wasn’t set.	

## Inventory Movement Line Model <a name="InventoryMovementLineModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	GUID			Identifier of sale Invoice task	
	ProductID	GUID			Product identifier referenced by this Line.	
	Date	string			Inventory Movement Date	
	COGS	Deciaml			Inventory Movement Amount	

## Sale Transaction Line Model <a name="SaleTransactionLineModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	GUID			Identifier of sale task	
	TransactionID	GUID			Identifier of transaction	
	Debit	string			Debit Account	
	Credit	string			Credit Account	
	Description	string			Transaction description	
	Amount	float32 with up to 2 decimal places			Amount in base currency	
	EffectiveDate	string			Transaction Date	

## Purchase Manual Journal Model <a name="PurchaseManualJournalModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Status	string		Yes	Manual Journal Status. Available values are NOT AVAILABLE, DRAFT, AUTHORISED. Available values for POST are DRAFT, AUTHORISED	
	Lines	<a href="#PurchaseManualJournalLineModel"><b>[] Purchase Manual Journal Line Model</b></a>				

## Purchase Manual Journal Line Model <a name="PurchaseManualJournalLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Reference	string			Reference Service Object	
	Amount	float32 with up to 2 decimal places		Yes	Amount in base currency for manual journal	
	Date	string		Yes	Effective date of manual journal	
	Debit	string		Yes	Debit Account	
	Credit	string		Yes	Credit Account	
	IsSystem	boolean			Read only. rows with true value cannot be deleted or modified.	

## Purchase Order Model <a name="PurchaseOrderModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Memo	string	1024	Yes	Additional information for Order	
	Status	string		Yes	Purchase Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#PurchaseOrderLineModel"><b>[] Purchase Order Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#PurchaseAdditionalChargeModel"><b>[] Purchase Additional Charge Model</b></a>				
	Prepayments	<a href="#SalePaymentLineModel"><b>[] Sale Payment Line Model</b></a>				
	TotalBeforeTax	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges without taxes. Not required for POST.	
	Tax	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges taxes. Not required for POST.	
	Total	float32 with up to 4 decimal places		Yes	Sum of order lines and additional charges with taxes. Not required for POST.	

## Purchase Order Line Model <a name="PurchaseOrderLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	GUID		Yes*	Product identifier referenced by this Line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false	
	SKU	string	50	Yes*	Product SKU referenced by this Line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false	
	Name	string	1024	Yes	Product Name referenced by this Line	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	Price	float32 with up to 4 decimal places		Yes	Price per unit in Customer currency	
	Discount	float32 with up to 2 decimal places			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0	
	Tax	float32 with up to 4 decimal places		Yes	Tax	
	TaxRule	string	50	Yes	Line Tax Rule name.	
	SupplierSKU	string	50		Supplier SKU referenced by this line.	
	Comment	string	256		Comment for this line	
	Total	float32 with up to 4 decimal places		Yes	Line Total. For validation	

## Purchase Additional Charge Model <a name="PurchaseAdditionalChargeModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Description	string	256	Yes	Name of Service Product referenced by this Line	
	Reference	string	256		Comment for this Line	Custom reference	
	Price	float32 with up to 4 decimal places		Yes	Price per unit in Customer currency	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	Discount	float32 with up to 2 decimal places			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0	
	Tax	float32 with up to 4 decimal places		Yes	Tax	
	Total	float32 with up to 4 decimal places		Yes	Line Total. For validation	
	TaxRule	string	50	Yes	Line Tax Rule name.	

## Purchase Stock Model <a name="PurchaseStockModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Status	string		Yes	Purchase Stock Received Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#PurchaseStockLineModel"><b>[] Purchase Stock Line Model</b></a>		Yes		

## Purchase Stock Line Model <a name="PurchaseStockLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Date	string		Yes	Date when items were received	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	ProductID	GUID		Yes*	Product identifier referenced by this Line.	
	SKU	string	50	Yes*	Product SKU referenced by this Line	
	Name	string	1024		Read-only. Product name referenced by this Line	
	Location	string	256	Yes*	Location where the product would be stock to/from. Required if LocationID is empty.	
	LocationID	GUID		Yes*	LocationID where the product would be stock to/from. Required if Location is empty.	
	Received	boolean			Flag, which indicates if items were received. Read-only.	
	BatchSN	string	50		Batch Serial Number	
	SupplierSKU	string	50		Supplier SKU referenced by this line.	
	ExpiryDate	string			Date when selected Batch expires	
	CardID	GUID			Stock Batch ID	

## Purchase Unstock Line Model <a name="PurchaseUnStockLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	CardID	GUID		Yes	Stock Batch ID	
	Date	string			Date when items were received	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Value should me more than zero.	
	ProductID	GUID			Read-only.Product identifier referenced by this Line.	
	SKU	string	50		Read-only.Product SKU referenced by this Line	
	Name	string	1024		Read-only. Product name referenced by this Line	
	Location	string	256		Read-only.Location where the product would be stock to/from	
	BatchSN	string	50		Read-only.Batch Serial Number	
	ExpiryDate	string			Read-only.Date when selected Batch expires	

## Purchase Invoice Model <a name="PurchaseInvoiceModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	InvoiceDate	string		Yes	Date when invoice created. Default value is current date, used if not specified.	
	InvoiceDueDate	string		Yes	Date until invoice is valid. If not specified, used default value from Terms.	
	InvoiceNumber	string			Invoice Number (auto-generated)	
	Status	string		Yes	Purchase Invoice Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED, PAID. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#PurchaseInvoiceLineModel"><b>[] Purchase Invoice Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#PurchaseInvoiceAdditionalChargeModel"><b>[] Purchase Invoice Additional Charge Model</b></a>				
	Payments	<a href="#SalePaymentLineModel"><b>[] Sale Payment Line Model</b></a>				
	TotalBeforeTax	float32 with up to 4 decimal places			Sum of order lines and additional charges without taxes. Not required for POST.	
	Tax	float32 with up to 4 decimal places			Sum of order lines and additional charges taxes. Not required for POST.	
	Total	float32 with up to 4 decimal places			Sum of order lines and additional charges with taxes. Not required for POST.	
	Paid	float32 with up to 4 decimal places			Invoice Paid Total in Supplier Currency. Not required for POST.	

## Purchase Credit Note Model <a name="PurchaseCreditNoteModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	CreditNoteNumber	string		Yes	Credit Note Number	
	CreditNoteDate	string		Yes	Credit Note Date	
	Status	string		Yes	Purchase Credit Note Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#PurchaseInvoiceLineModel"><b>[] Purchase Invoice Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#PurchaseInvoiceAdditionalChargeModel"><b>[] Purchase Invoice Additional Charge Model</b></a>				
	Refunds	<a href="#SalePaymentLineModel"><b>[] Sale Payment Line Model</b></a>				
	Unstock	<a href="#PurchaseUnStockLineModel"><b>[] Purchase Unstock Line Model</b></a>		Yes		
	TotalBeforeTax	float32 with up to 4 decimal places			Sum of order lines and additional charges without taxes. Not required for POST.	
	Tax	float32 with up to 4 decimal places			Sum of order lines and additional charges taxes. Not required for POST.	
	Total	float32 with up to 4 decimal places			Sum of order lines and additional charges with taxes. Not required for POST.	

## Purchase Invoice Line Model <a name="PurchaseInvoiceLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	GUID		Yes*	Product identifier referenced by this Line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false	
	SKU	string	50	Yes*	Product SKU referenced by this Line. Required If CombineAdditionalCharges param exist for this endpoint and it have values = false	
	Name	string	1024	Yes	Product Name referenced by this Line	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	Price	float32 with up to 4 decimal places		Yes	Price per unit in Customer currency	
	Discount	float32 with up to 2 decimal places			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0	
	Tax	float32 with up to 4 decimal places		Yes	Tax	
	TaxRule	string	50	Yes	Line Tax Rule name.	
	Account	string	50	Yes	Revenue account	
	Comment	string	256		Comment for this line	
	Total	float32 with up to 4 decimal places		Yes	Line Total. For validation	

## Purchase Invoice Additional Charge Model <a name="PurchaseInvoiceAdditionalChargeModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Description	string	256	Yes	Name of Service Product referenced by this Line	
	Reference	string	256		Comment for this Line	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	Price	float32 with up to 4 decimal places		Yes	Price per unit in Customer currency	
	Discount	float32 with up to 2 decimal places			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0	
	Tax	float32 with up to 4 decimal places		Yes	Tax	
	TaxRule	string	50	Yes	Line Tax Rule name.	
	Account	string	50	Yes	Revenue account	
	Total	float32 with up to 4 decimal places			Line Total. For validation	

## Existing Stock Line Model <a name="ExistingStockLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	GUID			Product identifier referenced by this Line	
	SKU	string	50		Product SKU referenced by this Line.	
	ProductName	string	1024		Product Name referenced by this Line	
	QuantityOnHand	float32			Quantity of Stock ‘‘On Hand’’	
	Available	float32			Available quantity of stock	
	Adjustment	float32			New value for QuantityOnHand	
	LocationID	GUID			Id of location where stock changes will be done	
	Location	string	256		Name of location where stock changes will be done	
	BatchSN	string	50		Batch Serial Number	
	ExpiryDate	string			Date when Batch expires	
	Comments	string	256		Comment	
	ProductLength	float32			Length of product. Read-only.	
	ProductWidth	float32			Width of product. Read-only.	
	ProductHeight	float32			Height of product. Read-only.	
	ProductWeight	float32			Weight of product. Read-only.	
	 WeightUnits  	 string             	 10        	          	Unit of measure for unit weight. Available Values are <a href="#DimensionUnitAvailableValues"><b>values</b></a>	
	 DimensionsUnits  	 string             	 10        	      	Unit of measure for unit length/width/height. Available Values are <a href="#DimensionUnitAvailableValues"><b>values</b></a>	
	ProductCustomField1	string			Value of Product's additional attribute 1. Read-only.	
	ProductCustomField2	string			Value of Product's additional attribute 2. Read-only.	
	ProductCustomField3	string			Value of Product's additional attribute 3. Read-only.	
	ProductCustomField4	string			Value of Product's additional attribute 4. Read-only.	
	ProductCustomField5	string			Value of Product's additional attribute 5. Read-only.	
	ProductCustomField6	string			Value of Product's additional attribute 6. Read-only.	
	ProductCustomField7	string			Value of Product's additional attribute 7. Read-only.	
	ProductCustomField8	string			Value of Product's additional attribute 8. Read-only.	
	ProductCustomField9	string			Value of Product's additional attribute 9. Read-only.	
	ProductCustomField10	string			Value of Product's additional attribute 10. Read-only.	
	Image	<a href="#AttachmentLineModel"><b>Image</b></a>			Product Image. Read-only. this field is available for stock take	
	Barcode	string			Value of Product's Barcode. Read-only.this field is available for stock take	
	StockLocator	string			Value of Product's StockLocator. Read-only.this field is available for stock take	
	Unit	string			Value of Product's Default Unit of Measure. Read-only.this field is available for stock take	
	CostingMethod	string			Value of Product's Costing method. Read-only.this field is available for stock take. Valid values: FIFO, Special - Batch, Special - Serial Number, FIFO - Serial Number, FIFO - Batch, FEFO - Batch, FEFO - Serial Number 	

## New Stock Line Model <a name="NewStockLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	GUID		Yes*	Product identifier referenced by this Line. Required if SKU is empty.	
	SKU	string	50	Yes*	Product SKU referenced by this Line. Required if ProductID is empty.	
	ProductName	string	1024		Product Name referenced by this Line	
	Quantity	float32		Yes	New value for QuantityOnHand	
	UnitCost	float32		Yes	Cost of 1 product unit.	
	LocationID	GUID		Yes*	Id of location on which stock changes will be done. Required if Location is empty.	
	Location	string	256	Yes*	Name of location on which stock changes will be done. Required if LocationID is empty.	
	BatchSN	string	50	Yes*	Batch Serial Number. Required if Product's Costing Method is not FIFO.	
	ExpiryDate	string			Date when Batch expires. Required if Product's Costing Method is FESN	
	ReceivedDate	string			Date when new stock is received	
	Comments	string	256		Comment	
	ProductLength	float32			Length of product. Read-only.	
	ProductWidth	float32			Width of product. Read-only.	
	ProductHeight	float32			Height of product. Read-only.	
	ProductWeight	float32			Weight of product. Read-only.	
	 WeightUnits  	 string             	 10        	          	Unit of measure for unit weight. Available Values are <a href="#DimensionUnitAvailableValues"><b>values</b></a>	
	 DimensionsUnits  	 string             	 10        	      	Unit of measure for unit length/width/height. Available Values are <a href="#DimensionUnitAvailableValues"><b>values</b></a>	
	ProductCustomField1	string			Value of Product's additional attribute 1. Read-only.	
	ProductCustomField2	string			Value of Product's additional attribute 2. Read-only.	
	ProductCustomField3	string			Value of Product's additional attribute 3. Read-only.	
	ProductCustomField4	string			Value of Product's additional attribute 4. Read-only.	
	ProductCustomField5	string			Value of Product's additional attribute 5. Read-only.	
	ProductCustomField6	string			Value of Product's additional attribute 6. Read-only.	
	ProductCustomField7	string			Value of Product's additional attribute 7. Read-only.	
	ProductCustomField8	string			Value of Product's additional attribute 8. Read-only.	
	ProductCustomField9	string			Value of Product's additional attribute 9. Read-only.	
	ProductCustomField10	string			Value of Product's additional attribute 10. Read-only.	
	Image	<a href="#AttachmentLineModel"><b>Image</b></a>			Product Image. Read-only. this field is available for stock take	
	Barcode	string			Value of Product's Barcode. Read-only.this field is available for stock take	
	StockLocator	string			Value of Product's StockLocator. Read-only.this field is available for stock take	
	Unit	string			Value of Product's Default Unit of Measure. Read-only.this field is available for stock take	
	CostingMethod	string			Value of Product's Costing method. Read-only.this field is available for stock take. Valid values: FIFO, Special - Batch, Special - Serial Number, FIFO - Serial Number, FIFO - Batch, FEFO - Batch, FEFO - Serial Number 	

## Transaction Stock Line Model <a name="TransactionStockLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	GUID			ID of transaction.	
	Debit	string			Debit Account.	
	Credit	string			Credit Account.	
	Amount	float32			Amount on which products quantity changed.	
	EffectiveDate	string			Transaction Date	

## Stock Transfer Line Model <a name="StockTransferLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	GUID		Yes*	Product identifier referenced by this Line. Required if SKU is empty.	
	SKU	string	50	Yes*	Product SKU referenced by this Line. Required if ProductID is empty.	
	ProductName	string	1024		Product Name referenced by this Line	
	QuantityOnHand	float32			Quantity of Stock ‘‘On Hand’’	
	QuantityAvailable	float32			Available quantity of stock	
	TransferQuantity	float32		Yes	Quantity to transfer from one location to another.	
	BatchSN	string	50	Yes*	Batch Serial Number. Required if Product's Costing Method is not FIFO.	
	ExpiryDate	string			Date when Batch expires. Required if Product's Costing Method is FESN	
	Comments	string	256		Comment	
	ProductLength	float32			Length of product. Read-only.	
	ProductWidth	float32			Width of product. Read-only.	
	ProductHeight	float32			Height of product. Read-only.	
	ProductWeight	float32			Weight of product. Read-only.	
	 WeightUnits  	 string             	 10        	          	Unit of measure for unit weight. Available Values are <a href="#DimensionUnitAvailableValues"><b>values</b></a>	
	 DimensionsUnits  	 string             	 10        	      	Unit of measure for unit length/width/height. Available Values are <a href="#DimensionUnitAvailableValues"><b>values</b></a>	
	ProductCustomField1	string			Value of Product's additional attribute 1. Read-only.	
	ProductCustomField2	string			Value of Product's additional attribute 2. Read-only.	
	ProductCustomField3	string			Value of Product's additional attribute 3. Read-only.	
	ProductCustomField4	string			Value of Product's additional attribute 4. Read-only.	
	ProductCustomField5	string			Value of Product's additional attribute 5. Read-only.	
	ProductCustomField6	string			Value of Product's additional attribute 6. Read-only.	
	ProductCustomField7	string			Value of Product's additional attribute 7. Read-only.	
	ProductCustomField8	string			Value of Product's additional attribute 8. Read-only.	
	ProductCustomField9	string			Value of Product's additional attribute 9. Read-only.	
	ProductCustomField10	string			Value of Product's additional attribute 10. Read-only.	

## Price Tier Model <a name="PriceTierModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Price Tier Name	float32			Product Sale Price used when Price Tier is selected in Sale Task. Up to 4 decimal places.	

## Product Supplier Model <a name="ProductSupplierModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	SupplierID	string		Yes*	ID of supplier. Required if SupplierName is empty	
	SupplierName	string	256	Yes*	Name of supplier. Required if SupplierID is empty	
	ProductID	string		Yes*	ID of product. Required when not nested within the Product and ProductSKU is empty	
	ProductSKU	string	256	Yes*	SKU of product. Required when not nested within the Product and ProductID is empty	
	ProductSupplierID	string		No	ID of ProductSupplier record.	
	SupplierInventoryCode	string	256		Supplier Product SKU	
	SupplierProductName	string	256		Supplier Product Name	
	Cost	float32			Latest purchase cost	
	FixedCost	float32			Fixed Cost	
	Currency	string			Supplier currency. Read-only.	
	DropShip	boolean		true if this supplier should be used for drop-shipping.	
	SupplierProductURL	string	256		Link to supplier product specification.	
	LastSupplied	string			The latest supplied date of a product.	
	ProductSupplierOptions	<a href="#ProductSupplierOptionsModel"><b>[] Product Supplier Options</b></a>				

## Product Supplier Options Model <a name="ProductSupplierOptionsModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	string		Yes*	ID of product supplier option set.	
	LocationID	string		Yes*	ID of location. Required if LocationName is empty	
	LocationName	string	256	Yes*	Name of Location. Required if LocationID is empty	
	ReorderQuantity	float32			Reorder Quantity	
	Lead	int32			Lead	
	Safety	int32			Safety.	
	MinimumToReorder	float32		Can be set for default options set only.	
	SupplyIntervals	<a href="#ProductSupplierOptionsIntervalModel"><b>[] Product Supplier Options Intervals </b></a>				

## Product Supplier Options Interval Model <a name="ProductSupplierOptionsIntervalModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	string		Yes*	ID of product supplier option Interval settings set.	
	DeliveryMethod	string	256	Yes*	Delivery method: can be 'Fixed' and 'Interval'	
	IntervalDays	int32			Interval Days, required when DeliveryMethod = Interval	
	IntervalStartDate	Date			Interval Start Date, required when DeliveryMethod = Interval	
	IsMonday	boolean			Required when DeliveryMethod = Fixed.	
	IsTuesday	boolean			Required when DeliveryMethod = Fixed.	
	IsWednesday	boolean			Required when DeliveryMethod = Fixed.	
	IsThursday	boolean			Required when DeliveryMethod = Fixed.	
	IsFriday	boolean			Required when DeliveryMethod = Fixed.	
	IsSaturday	boolean			Required when DeliveryMethod = Fixed.	
	IsSunday	boolean			Required when DeliveryMethod = Fixed.	

## Reorder Level Model <a name="ReorderLevelModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	LocationID	string		Yes*	ID of location. Required if LocationName is empty	
	LocationName	string	256	Yes*	Name of Location. Required if LocationID is empty	
	MinimumBeforeReorder	float32			Minimum available product quantity for this product to appear on Reorder report/forms. Defaults to 0	
	ReorderQuantity	float32			Default quantity to put into purchase order when reordering this product. Not applicable to reorder backordered form. Defaults to 0	
	StockLocator	string	256		Location of this product in your warehouse	
	PickZones	string	512	Yes	Comma delimited list of pick zones	

## Bill Of Material Product Model <a name="BillOfMaterialProductModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ComponentProductID	string		Yes*	ID of product. Required if ProductCode is empty	
	ProductCode	string	256	Yes*	SKU of product. Required if ComponentProductID is empty	
	Name	string	256		Name of product. Read-only.	
	Quantity	float32		Yes	Quantity	
	WastagePercent	float32			 WastagePercent and WastageQuantity are mutually exclusive.	
	WastageQuantity	float32			 WastagePercent and WastageQuantity are mutually exclusive.	
	CostPercentage	float32			Prorated cost distribution of a product that has been disassembled to current component. If left as zero equal distribution will be used between components.	

## Bill Of Material Service Model <a name="BillOfMaterialServiceModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ComponentProductID	string		Yes*	ID of service product. Required if Name is empty	
	Name	string	256	Yes*	Name of service product. Required if ComponentProductID is empty	
	Quantity	float32		Yes	Quantity	
	ExpenseAccount	string			Expense Account	
	PriceTier	int32			Price Tier Number	

## Product Movement Model <a name="ProductMovementModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string			Movement task ID	
	Type	string	256		Movement type. Possible values are <a href="#ProductMovementTypes"><b>Values</b></a> 	
	Date	string			Movement date	
	Number	string			Task Number	
	Status	int32			Task Status	
	Quantity	float32			Quantity of goods reallocated	
	Amount	float32			Cost of the reallocated goods	
	Location	float32			Location name	
	BatchSN	float32			Batch #	
	ExpiryDate	string			Expiry Date	
	FromTo	string			Supplier/CustomerName	

## Attribute Set Line Model <a name="AttributeSetLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Name	string			Attribute name	
	Type	string			Attribute type. Available values are Not used, Text, Checkbox, List. 	
	Values	string			Attribute values	

## Error Model <a name="ErrorModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ErrorCode	int32			Error code	
	Exception	string			Error text	

## Disassembly Pick Line Model <a name="DisassemblyPickLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	BinID	string			ID of Bin.	
	Bin	string	256		Name of Bin	
	Date  	 string           	         	          	Date Received         	
	BatchSN  	 string           	         	          	Batch serial number       	
	ExpiryDate  	 string           	         	          	Batch Expiry Date          	
	Available	float32			Available Quantity.	
	UnitCost 	float32			Cost.	
	Quantity	float32			Picked Quantity.	
	TotalCost 	float32			Total Line Cost.	

## Finished Goods Pick Line Model <a name="FinishedGoodsPickLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	string		Yes*	ID of product. Required if ProductCode is empty	
	ProductCode	string	256	Yes*	SKU of product. Required if ProductID is empty	
	Name	string	256		Name of product. Read-only.	
	BinID	string			ID of Bin.	
	Bin	string	256		Name of Bin.	
	BatchSN  	 string           	         	          	Batch serial number       	
	ExpiryDate  	 string           	         	          	Batch Expiry Date          	
	Quantity	float32		Yes	Quantity	
	Unit 	string			Unit of measure. Read-only.	
	Cost 	float32			Cost. Read-only.	

## Finished Goods Order Line Model <a name="FinishedGoodsOrderLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	string		Yes*	ID of product. Required if ProductCode is empty	
	ProductCode	string	256	Yes*	SKU of product. Required if ProductID is empty	
	Name	string	256		Name of product. Read-only.	
	ExpenseAccount	string		Yes*	Required if product type is Service	
	Quantity	float32		Yes	Quantity	
	Unit 	string			Unit of measure. Read-only.	
	WastagePercent	float32			 WastagePercent and WastageQuantity are mutually exclusive.	
	WastageQuantity	float32			 WastagePercent and WastageQuantity are mutually exclusive.	
	TotalQuantity	float32			Prorated cost distribution of a product that has been disassembled to current component. Read-only.	
	TotalCost	float32		Yes*	Total cost. For products with type = Stock this parameter is auto-calculated, for other type’s its required.	

## Disassembly Order Line Model <a name="DisassemblyOrderLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	string		Yes*	ID of product. Required if ProductCode is empty	
	ProductCode	string	256	Yes*	SKU of product. Required if ProductID is empty	
	Name	string	256		Name of product. Read-only.	
	BinID	string			ID of Bin.	
	Bin	string	256		Name of Bin.	
	Unit 	string			Unit of measure. Read-only.	
	BatchSN  	 string           	         	          	Batch serial number       	
	ExpiryDate  	 string           	         	          	Batch Expiry Date          	
	Account	string			Inventory Account	
	Quantity	float32		Yes	Quantity	
	Cost	float32		Yes	Cost.	

## Disassembly Order Service Line Model <a name="DisassemblyOrderServiceLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	string		Yes*	ID of product. Required if Name is empty	
	Name	string	256	Yes*	Name of service product. Required if ProductID is empty	
	Account	string		Yes	Expense Account	
	Comments	string			Comment to line.	
	Amount	float32		Yes	Amount.	

## Inventory Write-Off Line Model <a name="InventoryWriteOffLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	string		Yes*	ID of product. Required if ProductCode is empty	
	ProductCode	string	256	Yes*	SKU of product. Required if ProductID is empty	
	Name	string	256		Name of product. Read-only.	
	BinID	string			ID of Bin.	
	Bin	string	256		Name of Bin.	
	BatchSN  	 string           	         	          	Batch serial number       	
	ExpiryDate  	 string           	         	          	Batch Expiry Date          	
	ExpenseAccount	string		Yes*	Required if product type is Service	
	Quantity	float32		Yes	Quantity	
	Cost 	float32		Yes*	Cost. Required if product type is Service. Default value = 0.	
	TotalCost 	float32			Cost * Quantity. Read-only.	

## Tax Component Model <a name="TaxComponentModel" />
	Property           	 Type      	 Length 	 Required 	 Notes                         	
	:------------------	-----------	--------	----------	-------------------------------	
	Name             	 string    	 50     	 Yes      	 Name of product. Read-only.   	
	Percent          	 float32   	        	 Yes      	 Cost. Required if product type is Service. Default value = 0.	
	AccountCode      	 string    	        	 Yes      	 ChartOfAccount Code with Class == LIABILITY && Status == ACTIVE 	
	ComponentOrder   	 int32      	        	 Yes      	 The order of sequence taxes components. 	

## Money Task Line Model <a name="MoneyTaskLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Name	string	256	Yes	Name	
	Comment	string	256		Comment for this Line	
	Quantity	float32 with up to 4 decimal places		Yes	Quantity.	
	Price	float32 with up to 4 decimal places			Price per unit in Base currency. Default value is 0	
	Discount	float32 with up to 2 decimal places			Discount. Value between 0 and 100. For free items discount is 100. Default value is 0	
	Tax	float32 with up to 4 decimal places			Tax. Default value is 0	
	TaxRule	string	50	Yes	Line Tax Rule name.	
	Account	string	50	Yes	Revenue account	
	Total	float32 with up to 4 decimal places		Yes	Line Total. For validation	

## Supplier/Customer Address Model <a name="SupplierAddressModel" /> 
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ID           	 string    	    	           	 If passed in PUT method, entry will be searched by id, found entry will be updated, otherwise created.      	
	 Line1           	 string    	 256   	 Yes       	 Address Line 1.       	
	 Line2           	 string    	 256   	           	 Address Line 2.       	
	 City            	 string    	 256   	           	 City / Suburb.        	
	 State           	 string    	 256   	           	 State / Province.     	
	 Postcode        	 string    	 20    	           	 Zip / PostCode.       	
	 Country         	 string    	       	 Yes       	 Country name          	
	 Type            	 string    	       	 Yes       	 Address Type. Should be one of the following values: Billing, Business or Shipping. 	
	 DefaultForType  	 bool      	       	           	 Points that Address is used as default for chosen Type. False as default. 	

## Supplier/Customer Contact Model <a name="SupplierContactModel" /> 
	 Property          	 Type      	 Length    	 Required  	 Notes                 	
	:------------------	-----------	-----------	-----------	-----------------------	
	 ID           	 string    	    	           	 If passed in PUT method, entry will be searched by id, found entry will be updated, otherwise created.      	
	 Name            	 string    	 256       	 Yes       	 Name of Contact.      	
	 Phone           	 string    	 50        	           	 Phone                 	
	 MobilePhone           	 string    	 50        	           	 Mobile Phone                 	
	 Fax             	 string    	 50        	           	 Fax                   	
	 Email           	 string    	 256       	           	 Email                 	
	 Website         	 string    	 256       	           	 Website               	
	 Comment         	 string    	 256       	           	 Comment               	
	 Default         	 bool      	           	           	 Points that Contact is used as default. False as default. 	
	 IncludeInEmail  	 bool      	           	           	 Points that Contact is included in Email. false as default. 

## Customer specific Product Price Model <a name="ProductPriceModel" /> 
	 Property          	 Type      	 Length	 Required  	 Notes                 	
	:------------------	-----------	-------	-----------	-----------------------	
	 ProductID       	 string      	       	 Yes*      	                       	
	 CustomerID      	 string      	       	 Yes*      	                       	
	 CustomerName    	 string    	 256   	 Yes*      	                       	
	 ProductSKU      	 string    	 50    	 Yes*      	                        	
	 ProductName     	 string    	       	           	 Readonly                	
	 Price           	 float32   	       	   Yes     	      	

## Journal Line Model <a name="JournalLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Debit	string		Yes	Debit Account	
	Credit	string		Yes	Credit Account	
	Reference	string			Custom Reference	
	Amount	float32 with up to 2 decimal places		Yes	Amount in selected currency.	
	BaseAmount	float32 with up to 2 decimal places		Yes	Amount in base currency.	

## Purchase Payment Line Model <a name="PurchasePaymentLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	PurchaseID	GUID			Identifier of Master Purchase Task	
	TaskID	GUID			Identifier of Purchase Task	
	ID	GUID			Identifier of payment	
	Reference	string			Payment reference number	
	Amount	float32 with up to 2 decimal places			Payment amount in customer currency	
	DatePaid	string			Date when payment has been made	
	Account	string			Account Code of the bank/payment account from Chart of accounts	
	CurrencyRate	float32 with up to 4 decimal places			Currency Conversion rate expressed as number of Base currency units for one Customer currency unit.	
	DateCreated	string			Date of creation payment record.	

## Advanced Purchase Stock Model <a name="AdvancedPurchaseStockModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	GUID		Yes	Identifier of Purchase Stock Receiving Task	
	Status	string	50	Yes	Purchase Stock Received Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#AdvancedPurchaseStockLineModel"><b>[] Advanced Purchase Stock Line Model</b></a>		Yes		

## Advanced Purchase Stock Line Model <a name="AdvancedPurchaseStockLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Date	string		Yes	Date when items were received	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	ProductID	GUID		Yes*	Product identifier referenced by this Line.	
	SKU	string	50	Yes*	Product SKU referenced by this Line	
	Name	string	1024		Read-only. Product name referenced by this Line	
	Location	string	256	No	Location where the product would be stock to/from.	
	LocationID	GUID		No	LocationID where the product would be stock to/from.	
	Received	boolean			Flag, which indicates if items were received. Read-only.	
	BatchSN	string	50		Batch Serial Number	
	SupplierSKU	string	50		Supplier SKU referenced by this line.	
	ExpiryDate	string			Date when selected Batch expires	

## Advanced Purchase Put Away Model <a name="AdvancedPurchasePutAwayModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	GUID		Yes	Identifier of Purchase Put Away Task	
	Status	string	50	Yes	Purchase Put Away Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#AdvancedPurchasePutAwayLineModel"><b>[] Advanced Purchase Put Away Line Model</b></a>		Yes		

## Advanced Purchase Put Away Line Model <a name="AdvancedPurchasePutAwayLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Date	string		Yes	Date when items were received	
	Quantity	float32 with up to 4 decimal places		Yes	Product or service quantity. Minimal value is 1.	
	ProductID	GUID		Yes*	Product identifier referenced by this Line.	
	SKU	string	50	Yes*	Product SKU referenced by this Line	
	Name	string	1024		Read-only. Product name referenced by this Line	
	Location	string	256	Yes*	Location where the product would be stock to/from. Required if LocationID is empty.	
	LocationID	GUID		Yes*	LocationID where the product would be stock to/from. Required if Location is empty.	
	Received	boolean			Flag, which indicates if items were received. Read-only.	
	BatchSN	string	50		Batch Serial Number	
	SupplierSKU	string	50		Supplier SKU referenced by this line.	
	ExpiryDate	string			Date when selected Batch expires	
	CardID	GUID			Stock Batch ID	

## Advanced Purchase Invoice Model <a name="AdvancedPurchaseInvoiceModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	GUID		Yes	Identifier of Purchase Invoice Task	
	InvoiceDate	string		Yes	Date when invoice created. Default value is current date, used if not specified.	
	InvoiceDueDate	string		Yes	Date until invoice is valid. If not specified, used default value from Terms.	
	InvoiceNumber	string			Invoice Number (auto-generated)	
	Status	string		Yes	Purchase Invoice Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED, PAID. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#PurchaseInvoiceLineModel"><b>[] Purchase Invoice Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#PurchaseInvoiceAdditionalChargeModel"><b>[] Purchase Invoice Additional Charge Model</b></a>				
	Payments	<a href="#PurchasePaymentLineModel"><b>[] Purchase Payment Line Model</b></a>				
	TotalBeforeTax	float32 with up to 4 decimal places			Sum of order lines and additional charges without taxes. Not required for POST.	
	Tax	float32 with up to 4 decimal places			Sum of order lines and additional charges taxes. Not required for POST.	
	Total	float32 with up to 4 decimal places			Sum of order lines and additional charges with taxes. Not required for POST.	
	Paid	float32 with up to 4 decimal places			Invoice Paid Total in Supplier Currency. Not required for POST.	

## Advanced Purchase Credit Note Model <a name="AdvancedPurchaseCreditNoteModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	GUID		Yes	Identifier of Purchase Credit Note Task	
	CreditNoteNumber	string		Yes	Credit Note Number	
	CreditNoteDate	string		Yes	Credit Note Date	
	Status	string		Yes	Purchase Credit Note Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST only DRAFT and AUTHORISED values accepted	
	Lines	<a href="#PurchaseInvoiceLineModel"><b>[] Purchase Invoice Line Model</b></a>		Yes		
	AdditionalCharges	<a href="#PurchaseInvoiceAdditionalChargeModel"><b>[] Purchase Invoice Additional Charge Model</b></a>				
	Refunds	<a href="#PurchasePaymentLineModel"><b>[] Purchase Payment Line Model</b></a>				
	Unstock	<a href="#PurchaseUnStockLineModel"><b>[] Purchase Unstock Line Model</b></a>		Yes		
	TotalBeforeTax	float32 with up to 4 decimal places			Sum of order lines and additional charges without taxes. Not required for POST.	
	Tax	float32 with up to 4 decimal places			Sum of order lines and additional charges taxes. Not required for POST.	
	Total	float32 with up to 4 decimal places			Sum of order lines and additional charges with taxes. Not required for POST.	

## Advanced Purchase Manual Journal Model <a name="'AdvancedPurchaseManualJournalModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	GUID		Yes	Identifier of Purchase Invoice Task	
	Status	string		Yes	Manual Journal Status. Available values are NOT AVAILABLE, DRAFT, AUTHORISED. Available values for POST are DRAFT, AUTHORISED	
	Lines	<a href="#PurchaseManualJournalLineModel"><b>[] Purchase Manual Journal Line Model</b></a>				

## IDName Model <a name="IDNameModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ID	GUID			Identifier	
	Name	string			Name	

## Stock Transfer Order Model <a name="StockTransferOrderModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	Status	string	50	Yes	Status of stock transfer order. Passible values are NOT AVAILABLE, DRAFT, AUTHORISED	
	Lines	<a href="#StockTransferOrderLineModel"><b>[] Stock Transfer Order Line Model</b></a>				

## Stock Transfer Order Line Model <a name="StockTransferOrderLineModel" /> 
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	ProductID	GUID		Yes*	Product identifier referenced by this Line. Required if SKU is empty.	
	SKU	string	50	Yes*	Product SKU referenced by this Line. Required if ProductID is empty.	
	ProductName	string	1024		Product Name referenced by this Line	
	QuantityOnHand	float32			Quantity of Stock ‘‘On Hand’’	
	QuantityAvailable	float32			Available quantity of stock	
	TransferQuantity	float32		Yes	Quantity to transfer from one location to another.	
	Comments	string	256		Comment	
	ProductLength	float32			Length of product. Read-only.	
	ProductWidth	float32			Width of product. Read-only.	
	ProductHeight	float32			Height of product. Read-only.	
	ProductWeight	float32			Weight of product. Read-only.	
	 WeightUnits  	 string             	 10        	          	Unit of measure for unit weight. Available Values are <a href="#DimensionUnitAvailableValues"><b>values</b></a>	
	 DimensionsUnits  	 string             	 10        	      	Unit of measure for unit length/width/height. Available Values are <a href="#DimensionUnitAvailableValues"><b>values</b></a>	
	ProductCustomField1	string			Value of Product's additional attribute 1. Read-only.	
	ProductCustomField2	string			Value of Product's additional attribute 2. Read-only.	
	ProductCustomField3	string			Value of Product's additional attribute 3. Read-only.	
	ProductCustomField4	string			Value of Product's additional attribute 4. Read-only.	
	ProductCustomField5	string			Value of Product's additional attribute 5. Read-only.	
	ProductCustomField6	string			Value of Product's additional attribute 6. Read-only.	
	ProductCustomField7	string			Value of Product's additional attribute 7. Read-only.	
	ProductCustomField8	string			Value of Product's additional attribute 8. Read-only.	
	ProductCustomField9	string			Value of Product's additional attribute 9. Read-only.	
	ProductCustomField10	string			Value of Product's additional attribute 10. Read-only.	
### Available fields for Disassembly List
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 DisassemblyNumber  	 string           	         	          	Disassembly Number        	
	 ProductID  	 string           	         	          	Product ID        	
	 ProductCode  	 string           	         	          	Product Code        	
	 ProductName  	 string           	         	          	Product Name         	
	 Quantity  	 float32           	         	          	Quantity          	
	 LocationID  	 string           	         	          	Location ID    	
	 Location  	 float32           	         	          	Location Name     	
	 Date  	 string           	         	          	Date         	
	 Status  	 string           	         	          	Disassembly Task Status. Available values are DRAFT,WORK IN PROGRESS,COMPLETED,VOIDED         	

### Available fields for Disassembly
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 DisassemblyNumber  	 string           	         	          	Disassembly Number        	
	 Status  	 string           	         	          	Disassembly Task Status. Available values are DRAFT,WORK IN PROGRESS,COMPLETED,VOIDED         	
	 ProductID  	 string           	         	          	Product ID        	
	 ProductCode  	 string           	         	          	Product Code        	
	 ProductName  	 string           	         	          	Product Name         	
	 LocationID  	 string           	         	          	Location ID    	
	 Location  	 float32           	         	          	Location Name     	
	 WIPAccount  	 string           	         	          	Work in Progress Account          	
	 Quantity  	 float32           	         	          	Quantity          	
	 AssemblyInstructionURL  	 string           	         	          	Assembly Instruction URL     	
	 PickLines  	<a href="#DisassemblyPickLineModel"><b>[] Disassembly Pick Line Model</b></a>	         	          	        	
	 OrderLines  	<a href="#DisassemblyOrderLineModel"><b>[] Disassembly Order Line Model</b></a>	         	          	        	
	 OrderServiceLines  	<a href="#DisassemblyOrderServiceLineModel"><b>[] Disassembly Order Service Line Model</b></a>	         	          	        	
	 Transactions  	<a href="#TransactionStockLineModel"><b>[] Transaction Stock Line Model</b></a>	         	          	        	
	 Errors  	<a href="#ErrorModel"><b>[] Error Model</b></a>	         	          	If While processing POST method, some errors occurred, but task was created, this property will contain array of error messages.    	

### Available fields for POST method
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 Status  	 string           	         	     Yes     	Disassembly Task Status. Available values are DRAFT,AUTHORISED,IN PROGRESS,COMPLETED.	
	 ProductID  	 string           	         	Yes*	Product ID. Required if ProductCode is empty       	
	 ProductCode  	 string           	         	Yes*	Product Code. Required if ProductID is empty 	
	 LocationID  	 string           	         	 Yes*         	Location ID. Required if Location is empty    	
	 Location  	 float32           	         	 Yes*         	Location Name. Required if LocationID is empty     	
	 WIPAccount  	 string           	         	    Yes      	Work in Progress Account          	
	 Quantity  	 float32           	         	    Yes      	Quantity          	

### Available fields for Disassembly Order
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 Status  	 string           	         	          	Disassembly Task Status. Available values are DRAFT,WORK IN PROGRESS,COMPLETED,VOIDED. Available values for POST are WORK IN PROGRESS and COMPLETED.	
	 OrderLines  	<a href="#DisassemblyOrderLineModel"><b>[] Disassembly Order Line Model</b></a>	         	        	

### Available fields for Finished Goods List
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 AssemblyNumber  	 string           	         	          	Assembly Number        	
	 BatchSN  	 string           	         	          	Batch serial number       	
	 ExpiryDate  	 string           	         	          	Batch Expiry Date          	
	 ProductID  	 string           	         	          	Product ID        	
	 ProductCode  	 string           	         	          	Product Code        	
	 ProductName  	 string           	         	          	Product Name         	
	 Quantity  	 float32           	         	          	Quantity          	
	 LocationID  	 string           	         	          	Location ID    	
	 Location  	 float32           	         	          	Location Name     	
	 Date  	 string           	         	          	Date         	
	 Status  	 string           	         	          	Finished Goods Task Status. Available values are DRAFT,AUTHORISED,IN PROGRESS,COMPLETED,VOIDED         	
	 UnitCost  	 float32           	         	          	Unit Cost         	
	 Notes  	 string           	         	          	Notes        	
	CustomField1	string			Value of Finished Goods additional attribute 1	
	CustomField2	string			Value of Finished Goods additional attribute 2	
	CustomField3	string			Value of Finished Goods additional attribute 3	
	CustomField4	string			Value of Finished Goods additional attribute 4	
	CustomField5	string			Value of Finished Goods additional attribute 5	
	CustomField6	string			Value of Finished Goods additional attribute 6	
	CustomField7	string			Value of Finished Goods additional attribute 7	
	CustomField8	string			Value of Finished Goods additional attribute 8	
	CustomField9	string			Value of Finished Goods additional attribute 9	
	CustomField10	string			Value of Finished Goods additional attribute 10	

### Available fields for Finished Goods
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 AssemblyNumber  	 string           	         	          	Assembly Number        	
	 Status  	 string           	         	          	Finished Goods Task Status. Available values are DRAFT,AUTHORISED,IN PROGRESS,COMPLETED,VOIDED         	
	 ProductID  	 string           	         	          	Product ID        	
	 ProductCode  	 string           	         	          	Product Code        	
	 ProductName  	 string           	         	          	Product Name         	
	 LocationID  	 string           	         	          	Location ID    	
	 Location  	 float32           	         	          	Location Name     	
	 BinID  	 string           	         	          	Bin ID    	
	 Bin  	 float32           	         	          	Bin Name     	
	 WIPAccount  	 string           	         	          	Work in Progress Account          	
	 WIPDate  	 string           	         	          	Work in Progress Date          	
	 Account  	 string           	         	          	Finished Goods Account          	
	 Quantity  	 float32           	         	          	Quantity          	
	 AssemblyInstructionURL  	 string           	         	          	Assembly Instruction URL     	
	 CompletionDate  	 string           	         	          	Completion Date         	
	 BatchSN  	 string           	         	          	Batch serial number       	
	 ExpiryDate  	 string           	         	          	Batch Expiry Date          	
	 Notes  	 string           	         	          	Notes        	
	 OrderLines  	<a href="#FinishedGoodsOrderLineModel"><b>[] Finished Goods Order Line Model</b></a>	         	          	        	
	 PickLines  	<a href="#FinishedGoodsPickLineModel"><b>[] Finished Goods Pick Line Model</b></a>	         	          	        	
	 Transactions  	<a href="#TransactionStockLineModel"><b>[] Transaction Stock Line Model</b></a>	         	          	        	
	 Errors  	<a href="#ErrorModel"><b>[] Error Model</b></a>	         	          	If While processing POST method, some errors occurred, but task was created, this property will contain array of error messages.    	
	CustomField1	string			Value of Finished Goods additional attribute 1	
	CustomField2	string			Value of Finished Goods additional attribute 2	
	CustomField3	string			Value of Finished Goods additional attribute 3	
	CustomField4	string			Value of Finished Goods additional attribute 4	
	CustomField5	string			Value of Finished Goods additional attribute 5	
	CustomField6	string			Value of Finished Goods additional attribute 6	
	CustomField7	string			Value of Finished Goods additional attribute 7	
	CustomField8	string			Value of Finished Goods additional attribute 8	
	CustomField9	string			Value of Finished Goods additional attribute 9	
	CustomField10	string			Value of Finished Goods additional attribute 10	

### Available fields for Disassembly Order
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 Status  	 string           	         	          	Disassembly Task Status. Available values are DRAFT,WORK IN PROGRESS,COMPLETED,VOIDED. Available values for POST are WORK IN PROGRESS and COMPLETED.	
	 OrderLines  	<a href="#DisassemblyOrderLineModel"><b>[] Disassembly Order Line Model</b></a>	         	        	

### Available Fields for Lead
	Property               	Type       	Length     	Required   	Notes	
	:----------------------	-----------	-----------	-----------	---------------------------------------------------------------	
	ID	string		yes	Unique lead ID generated by Cin7 Core.	
	LeadStatus	string	256	yes	Lead status. Available values are NEW, ATTEMPTING TO CONTACT, QUALIFIED, DISQUALIFIED.	
	Name	string	256	yes	Name of the lead.	
	Currency	string	3	yes	Currency code of the lead.	
	PaymentTerm	int		yes	Payment term for the lead.	
	PriceTier	string	50	yes	Price tier.	
	SalesRepresentative	string	256	yes	Sale representative name.	
	TaxRule	string	255	yes	Tax rule name for the lead.	
	Amount	decimal			Lead value.	
	CloseChance	int		yes	 % chance of wining a lead. Available values are 0%, 10%, 20%, 30%..... etc.	
	CloseDate	string		yes	Lead deal estimated close date.	
	Comments	string	2000		Additional comments.	
	Contacts	<a href="#SupplierContactModel"><b>[] Lead Contact Model</b></a>			List of contacts.	
	Addresses	<a href="#SupplierAddressModel"><b>[] Lead Address Model</b></a>			List of addresses.	

### Available fields for Webhooks
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 ID  	 string           	         	Yes*	Unique ID. Required for PUT     	
	 Type  	 string           	         	Yes	Webhook Type. Available values are  <a href="#webhookTypeValues"><b>values</b></a>   	
	 Name  	 string           	         		Webhook Friendly Name. Read-only.    	
	 IsActive  	 boolean           	         	Yes	Is webhook active.    	
	 ExternalURL  	 string           	         	Yes	Callback url.   	
	 ExternalAuthorizationType  	 string           	         	Yes	Authorisation type. Available values are noauth,basicauth, and bearerauth 	
	 ExternalUserName  	 string           	         	Yes*	User name. Required if ExternalAuthorizationType is basicauth 	
	 ExternalPassword  	 string           	         	Yes*	Password. Required if ExternalAuthorizationType is basicauth 	
	 ExternalBearerToken  	 string           	         	Yes*	Bearer token. Required if ExternalAuthorizationType is bearerauth 	
	 ExternalHeaders  	 Array of Key and Value objects.	         		Additional headers.	

<a name="webhookTypeValues" />
### Available values for Webhook Type 
	 Type     	Description	
	:-------	----------	
	 Sale/QuoteAuthorised  	   	
	 Sale/OrderAuthorised  	   	
	 Sale/Voided  	   	
	 Sale/Backordered  	   	
	 Sale/ShipmentAuthorised  	   	
	 Sale/InvoiceAuthorised  	   	
	 Sale/PickAuthorised  	   	
	 Sale/PackAuthorised  	   	
	 Sale/CreditNoteAuthorised  	   	
	 Sale/Undo  	   	
	 Sale/PartialPaymentReceived  	   	
	 Sale/FullPaymentReceived  	   	
	 Sale/ShipmentTrackingNumberChanged  	   	
	 Purchase/OrderAuthorised  	   	
	 Purchase/InvoiceAuthorised  	   	
	 Purchase/StockReceivedAuthorised  	   	
	 Purchase/CreditNoteAuthorised  	   	
	 Customer/Updated  	   	
	 Supplier/Updated  	   	
	 Stock/AvailableStockLevelChanged  	   	

### Available fields for Finished Goods List
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 AssemblyNumber  	 string           	         	          	Assembly Number        	
	 BatchSN  	 string           	         	          	Batch serial number       	
	 ExpiryDate  	 string           	         	          	Batch Expiry Date          	
	 ProductID  	 string           	         	          	Product ID        	
	 ProductCode  	 string           	         	          	Product Code        	
	 ProductName  	 string           	         	          	Product Name         	
	 Quantity  	 float32           	         	          	Quantity          	
	 LocationID  	 string           	         	          	Location ID    	
	 Location  	 float32           	         	          	Location Name     	
	 Date  	 string           	         	          	Date         	
	 Status  	 string           	         	          	Finished Goods Task Status. Available values are DRAFT,AUTHORISED,IN PROGRESS,COMPLETED,VOIDED         	
	 UnitCost  	 float32           	         	          	Unit Cost         	
	 Notes  	 string           	         	          	Notes        	
	CustomField1	string			Value of Finished Goods additional attribute 1	
	CustomField2	string			Value of Finished Goods additional attribute 2	
	CustomField3	string			Value of Finished Goods additional attribute 3	
	CustomField4	string			Value of Finished Goods additional attribute 4	
	CustomField5	string			Value of Finished Goods additional attribute 5	
	CustomField6	string			Value of Finished Goods additional attribute 6	
	CustomField7	string			Value of Finished Goods additional attribute 7	
	CustomField8	string			Value of Finished Goods additional attribute 8	
	CustomField9	string			Value of Finished Goods additional attribute 9	
	CustomField10	string			Value of Finished Goods additional attribute 10	

### Available fields for Finished Goods
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 AssemblyNumber  	 string           	         	          	Assembly Number        	
	 Status  	 string           	         	          	Finished Goods Task Status. Available values are DRAFT,AUTHORISED,IN PROGRESS,COMPLETED,VOIDED         	
	 ProductID  	 string           	         	          	Product ID        	
	 ProductCode  	 string           	         	          	Product Code        	
	 ProductName  	 string           	         	          	Product Name         	
	 LocationID  	 string           	         	          	Location ID    	
	 Location  	 float32           	         	          	Location Name     	
	 BinID  	 string           	         	          	Bin ID    	
	 Bin  	 float32           	         	          	Bin Name     	
	 WIPAccount  	 string           	         	          	Work in Progress Account          	
	 WIPDate  	 string           	         	          	Work in Progress Date          	
	 Account  	 string           	         	          	Finished Goods Account          	
	 Quantity  	 float32           	         	          	Quantity          	
	 AssemblyInstructionURL  	 string           	         	          	Assembly Instruction URL     	
	 CompletionDate  	 string           	         	          	Completion Date         	
	 BatchSN  	 string           	         	          	Batch serial number       	
	 ExpiryDate  	 string           	         	          	Batch Expiry Date          	
	 Notes  	 string           	         	          	Notes        	
	 OrderLines  	<a href="#FinishedGoodsOrderLineModel"><b>[] Finished Goods Order Line Model</b></a>	         	          	        	
	 PickLines  	<a href="#FinishedGoodsPickLineModel"><b>[] Finished Goods Pick Line Model</b></a>	         	          	        	
	 Transactions  	<a href="#TransactionStockLineModel"><b>[] Transaction Stock Line Model</b></a>	         	          	        	
	 Errors  	<a href="#ErrorModel"><b>[] Error Model</b></a>	         	          	If While processing POST method, some errors occurred, but task was created, this property will contain array of error messages.    	
	CustomField1	string			Value of Finished Goods additional attribute 1	
	CustomField2	string			Value of Finished Goods additional attribute 2	
	CustomField3	string			Value of Finished Goods additional attribute 3	
	CustomField4	string			Value of Finished Goods additional attribute 4	
	CustomField5	string			Value of Finished Goods additional attribute 5	
	CustomField6	string			Value of Finished Goods additional attribute 6	
	CustomField7	string			Value of Finished Goods additional attribute 7	
	CustomField8	string			Value of Finished Goods additional attribute 8	
	CustomField9	string			Value of Finished Goods additional attribute 9	
	CustomField10	string			Value of Finished Goods additional attribute 10	

### Available fields for Finished Goods Order
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 Status  	 string           	         	          	Finished Goods Task Status. Available values are DRAFT,AUTHORISED,IN PROGRESS,COMPLETED,VOIDED. Available values for POST are DRAFT and AUTHORISED.	
	 OrderLines  	<a href="#FinishedGoodsOrderLineModel"><b>[] Finished Goods Order Line Model</b></a>	         	        	

### Available fields for Finished Goods Pick
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 Status  	 string           	         	          	Finished Goods Task Status. Available values are DRAFT,AUTHORISED,IN PROGRESS,COMPLETED,VOIDED. Available values for POST are AUTHORISED, IN PROGRESS and COMPLETED.	
	 WIPAccount  	 string           	         	          	Work in Progress Account. Read-only if current Status is not AUTHORISED	
	 WIPDate  	 string           	         	          	Work in Progress Date. If Status is COMPLETED but task's WIPDate is empty, current date will be set. Read-only if current Status is not AUTHORISED	
	 Account  	 string           	         	          	Finished Goods Account. Read-only if current Status is not AUTHORISED or IN PROGRESS	
	 CompletionDate  	 string           	         	   Yes       	Completion Date. If Status is COMPLETED but task's CompletionDate is empty, current date will be set. Read-only if current Status is not AUTHORISED or IN PROGRESS	
	 PickLines  	<a href="#FinishedGoodsPickLineModel"><b>[] Finished Goods Pick Line Model</b></a>	         	          	        	

### Available fields for Inventory Write-Off List
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 InventoryWriteOffNumber  	 string           	         	          	Inventory Write-Off Number        	
	 Status  	 string           	         	          	Inventory Write-Off Task Status. Available values are DRAFT, COMPLETED, VOIDED         	
	 LocationID  	 string           	         	          	Location ID    	
	 Location  	 float32           	         	          	Location Name     	
	 Date  	 string           	         	          	Date  When Task Created.  	
	 Notes  	 string           	         	          	Notes        	

### Available fields for Inventory Write-Off
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 InventoryWriteOffNumber  	 string           	         	          	Inventory Write-Off Number        	
	 Status  	 string           	         	          	Inventory Write-Off Task Status. Available values are DRAFT, COMPLETED, VOIDED         	
	 LocationID  	 string           	         	          	Location ID    	
	 Location  	 float32           	         	          	Location Name     	
	 Account  	 string           	         	          	Expense Account         	
	 EffectiveDate  	 string           	         	          	Effective Date         	
	 Notes  	 string           	         	          	Notes        	
	 Lines  	<a href="#InventoryWriteOffLineModel"><b>[] Inventory Write-Off Line Model</b></a>	         	          	        	
	 Transactions  	<a href="#TransactionStockLineModel"><b>[] Transaction Stock Line Model</b></a>	         	          	        	
	 Errors  	<a href="#ErrorModel"><b>[] Error Model</b></a>	         	          	If While processing POST or PUT method, some errors occurred, but task was created, this property will contain array of error messages.    	

### Available fields for Inventory Write-Off
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 InventoryWriteOffNumber  	 string           	         	          	Inventory Write-Off Number        	
	 Status  	 string           	         	          	Inventory Write-Off Task Status. Available values are DRAFT, COMPLETED, VOIDED         	
	 LocationID  	 string           	         	          	Location ID    	
	 Location  	 float32           	         	          	Location Name     	
	 Account  	 string           	         	          	Expense Account         	
	 EffectiveDate  	 string           	         	          	Effective Date         	
	 Notes  	 string           	         	          	Notes        	
	 Lines  	<a href="#InventoryWriteOffLineModel"><b>[] Inventory Write-Off Line Model</b></a>	         	          	        	
	 Transactions  	<a href="#TransactionStockLineModel"><b>[] Transaction Stock Line Model</b></a>	         	          	        	
	 Errors  	<a href="#ErrorModel"><b>[] Error Model</b></a>	         	          	If While processing POST or PUT method, some errors occurred, but task was created, this property will contain array of error messages.    	

### Available fields for Journal
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	Yes*	  Unique ID. Required for PUT       	
	 Status  	 string           	         	    Yes      	Jorunal Status. Available values are DRAFT, COMPLETED, VOIDED         	
	 JournalNumber  	 string           	         	          	Jorunal Number. Read-only.        	
	 Currency  	 string           	         	    Yes  	Currency code.        	
	 CurrencyConversionRate 	 float32 	     	       Yes   	 Conversion rate between tenant currency and Currency currency.	
	 EffectiveDate  	 string           	         	       Yes   	Effective Date. 	
	 Narration  	 string           	         	          	Narration        	
	 Notes  	 string           	         	          	Notes        	
	Lines	<a href="#JournalLineModel"><b>[] Journal Line Model</b></a>				
	Attachments	<a href="#AttachmentLineModel"><b>[] Attachment Line Model</b></a>				

### Available fields for Money Task List
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	          	  Unique ID.        	
	 Date  	 string           	         	          	Date  of money operation.  	
	 TaskType  	 string           	         	          	Task Type. Available values are Spend Money, Receive Money, Transfer Money         	
	 Status  	 string           	         	          	Inventory Write-Off Task Status. Available values are DRAFT, COMPLETED, VOIDED         	
	 SupplierCustomerName  	 string           	         	          	Customer or Supplier who recive/spend money.      	
	 SupplierID  	 string           	         	          	Supplier ID   	
	 CustomerID  	 string           	         	          	Customer ID   	
	 Reference  	 float32           	         	          	Custom reference    	
	 TotalAmount  	 string           	         	          	Operation total amount.     	

### Available fields for Money Task List
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	Yes*	  Unique ID. Required for PUT       	
	 TaskType  	 string           	         	     Yes     	Task Type. Available values are Spend Money, Receive Money, Transfer Money         	
	 Status  	 string           	         	    Yes      	Money Task Status. Available values are DRAFT, COMPLETED, VOIDED         	
	 BankAccount 	 string 	     	     Yes     	 Bank Account Number.	
	 CurrencyConversionRate 	 float32 	     	          	 Conversion rate between BankAccount currency and base currency. If not passed, will be calculated automatically as base currency / bank account currency.	
	 SupplierCustomerName  	 string           	         	          	Customer or Supplier who recive/spend money.      	
	 SupplierID  	 string           	         	          	Supplier ID. CustomerID and SupplierID are mutually exclusive.   	
	 CustomerID  	 string           	         	          	Customer ID. CustomerID and SupplierID are mutually exclusive.  	
	 Reference  	 string           	         	          	Custom reference    	
	 Date  	 string           	         	     Yes     	Date  of money operation.  	
	 TaxInclusive  	 boolean           	         	          	true if inclusive. Default false.  	
	 Note  	 string           	         	          	Custom note.  	
	 Lines  	<a href="#MoneyTaskLineModel"><b>[] Money Task Line Model</b></a>	         	          	        	
	 Transactions  	<a href="#TransactionStockLineModel"><b>[] Transaction Stock Line Model</b></a>	         	          	        	
	Attachments	<a href="#AttachmentLineModel"><b>[] Attachment Line Model</b></a>			Sale Attachments	

### Available fields for Money Task List
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         	Yes*	  Unique ID. Required for PUT       	
	 Status  	 string           	         	    Yes      	Bank Transfer Status. Available values are DRAFT, COMPLETED, VOIDED         	
	 FromAccount 	 string 	     	     Yes     	From Bank Account Number.	
	 ToAccount 	 string 	     	     Yes     	To Bank Account Number.	
	 FromAmount 	 float32 	     	     Yes     	From Amount. In From Bank Account currency.	
	 ToAmount 	 float32 	     	     Yes     	 To Amount. In To Bank Account currency.	
	 CurrencyConversionRate 	 float32 	     	          	 Conversion rate between BankAccount currency and base currency. Will be calculated automatically as From Bank Account currency / To Bank Account currency.	
	 Reference  	 string           	         	          	Custom reference    	
	 Date  	 string           	         	     Yes     	Date of Bank Transfer.  	
	 Note  	 string           	         	          	Custom note.  	
	 Transactions  	<a href="#TransactionStockLineModel"><b>[] Transaction Stock Line Model</b></a>	         	          	        	
	Attachments	<a href="#AttachmentLineModel"><b>[] Attachment Line Model</b></a>			Sale Attachments	

### Available field for Purchase Manual Journal
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Purchase ID	
	Status	string		Yes	Manual Journal Status. Available values are NOT AVAILABLE, DRAFT, AUTHORISED. Available values for POST are DRAFT, AUTHORISED	
	Lines	<a href="#PurchaseManualJournalLineModel"><b>[] Purchase Manual Journal Line Model</b></a>				

### Available fields for Purchase Attachments
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Purchase ID	
	Lines	<a href="#AttachmentLineModel"><b>[] Attachment Line Model</b></a>				

### Available field for Purchase Manual Journal
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	PurchaseID	string		Yes	Unique Cin7 Core Purchase Master Task ID	
	ManualJournals	<a href="#AdvancedPurchasePartialMAnJModel"><b>[] Advanced Purchase Manual Journal Model</b></a>		Yes		

### Advanced purchase manual journal partial model
<a name="AdvancedPurchasePartialMAnJModel" />
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string		Yes	Unique Cin7 Core Purchase Invoice Task ID	
	Status	string		Yes	Manual Journal Status. Available values are NOT AVAILABLE, DRAFT, AUTHORISED. Available values for POST are DRAFT, AUTHORISED	
	Lines	<a href="#PurchaseManualJournalLineModel"><b>[] Purchase Manual Journal Line Model</b></a>				

### Available field for Sale Manual Journal
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	SaleID	string		Yes	Unique Cin7 Core Sale ID	
	Status	string		Yes	Manual Journal Status. Available values are NOT AVAILABLE, DRAFT, AUTHORISED. Available values for POST are DRAFT, AUTHORISED	
	Lines	<a href="#SaleManualJournalLineModel"><b>[] Sale Manual Journal Line Model</b></a>				

### Available fields for Sale Attachments
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	SaleID	string		Yes	Unique Cin7 Core Sale ID	
	Lines	<a href="#AttachmentLineModel"><b>[] Attachment Line Model</b></a>				

### Available fields for Stock Adjustment List
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string			Unique Cin7 Core Stock Adjustment ID	
	EffectiveDate	string			Date of transaction	
	StocktakeNumber	string			Stock Number (auto-generated).	
	Status	string			Status of stock adjustment. Available values are DRAFT,COMPLETED,VOIDED.	
	Account	string			Expense account for inventory adjustment	
	Reference	string			Custom reference	

### Available fields for Stock Take List
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string			Unique Cin7 Core Stock Take ID	
	EffectiveDate	string			Date of transaction	
	StocktakeNumber	string			Stock Take Number (auto-generated).	
	Status	string			Status of stock take. Available values are DRAFT,COMPLETED,VOIDED.	
	Account	string			Expense account for inventory stocktake	
	LocationID	string			Location ID 	
	Location	string			Location Name	
	Tags	string			Comma delimited list of custom tags	
	PickZones	string			Comma delimited list of custom pick zones	
	StockLocators	string			Comma delimited list of custom stock locators	
	Categories	string			Comma delimited list of custom categories	
	Brands	string			Comma delimited list of custom brands	
	Bins	string			Comma delimited list of custom bins	
	Reference	string			Custom reference	
	CreatedDate	string			Date of creation stocktake	
	LastUpdatedDate	string			Date and time when stocktake was last time modified	
	LastUpdatedBy	string			User which made last update	

### Available fields for Stock Transfer List
	Property	Type	Length	Required	Notes	
	:-------	----	------	--------	-----	
	TaskID	string			Unique Cin7 Core Stock Transfer ID	
	From	string			ID of location from which transfer was done	
	FromLocation	string			Name of location from which transfer was done	
	To	string			ID of location to which transfer was done	
	ToLocation	string			Name of location to which transfer was done	
	Status	string			Status of stock transfer. Available values are DRAFT, IN TRANSIT,COMPLETED,VOIDED.	
	Number	string			Stock Transfer Number (auto-generated).	
	CompletionDate	string			Date of transaction	
	CostDistributionType	string			Type of additional journals capitalization cost distribution. Available values are Cost, Quantity,Weight,Volume.	
	InTransitAccount	string			Asset Account holding stock value while it is in transit.	
	DepartureDate	string			Date when stock left the From warehouse.	
	Reference	string			Custom reference	

### Available fields for Transactions
	 Property 	 Type           	 Length 	 Required 	Notes       	
	:-------	----------------------	---------	----------	----------	
	 TaskID  	 string           	         		Task ID     	
	DebitAccountCode	string			Debit Account.	
	CreditAccountCode	string			Credit Account.	
	Amount	float32			Amount.	
	EffectiveDate	string			Transaction Date	
	Reference	string			Custom Reference	
	Transaction	string			Transaction description.	
	Type	string			Transaction Type. Available values are Purchase,Sale,MoneySpend,MoneyReceive,BankTransfer,ExpenseClaimTask,FinishedGoods,InventoryWriteOff,StockTake,StockAdjustment,Journal,Disassembly,Depreciation	

