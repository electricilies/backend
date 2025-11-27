package cacheredis

const (
	CacheTTLCategory       = 3600 // 1 hour
	CacheTTLAttribute      = 3600 // 1 hour
	CacheTTLAttributeValue = 3600 // 1 hour
	CacheTTLProduct        = 3600 // 1 hour
	CacheTTLCart           = 1800 // 30 minutes
)

const (
	CategoryListPrefix       = "category:list:"
	CategoryGetPrefix        = "category:get:"
	AttributeListPrefix      = "attribute:list:"
	AttributeGetPrefix       = "attribute:get:"
	AttributeValueListPrefix = "attribute_value:list:"
	ProductListPrefix        = "product:list:"
	ProductGetPrefix         = "product:get:"
	CartGetPrefix            = "cart:get:"
)
