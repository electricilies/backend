package http

import (
	"time"

	"backend/internal/domain"

	"github.com/google/uuid"
)

type ProductResponseDto struct {
	ID            uuid.UUID                     `json:"id"            binding:"required"`
	Name          string                        `json:"name"          binding:"required"`
	Description   string                        `json:"description"   binding:"required"`
	ViewsCount    int                           `json:"viewsCount"    binding:"required"`
	TotalPurchase int                           `json:"totalPurchase" binding:"required"`
	Price         float64                       `json:"price"         binding:"required"`
	Rating        float64                       `json:"rating"        binding:"required"`
	CreatedAt     time.Time                     `json:"createdAt"     binding:"required"`
	UpdatedAt     time.Time                     `json:"updatedAt"     binding:"required"`
	DeletedAt     *time.Time                    `json:"deletedAt"`
	Category      ProductCategoryResponseDto    `json:"category"      binding:"required"`
	Attributes    []ProductAttributeResponseDto `json:"attributes"    binding:"required"`
	Options       []ProductOptionResponseDto    `json:"options"       binding:"required"`
	Variants      []ProductVariantResponseDto   `json:"variants"      binding:"required"`
	Images        []ProductImageResponseDto     `json:"images"        binding:"required"`
}

type ProductCategoryResponseDto struct {
	ID        uuid.UUID  `json:"id"        binding:"required"`
	Name      string     `json:"name"      binding:"required"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type ProductAttributeResponseDto struct {
	ID        uuid.UUID                        `json:"id"        binding:"required"`
	Code      string                           `json:"code"      binding:"required"`
	Name      string                           `json:"name"      binding:"required"`
	Value     ProductAttributeValueResponseDto `json:"value"     binding:"required"`
	DeletedAt *time.Time                       `json:"deletedAt"`
}

type ProductAttributeValueResponseDto struct {
	ID        uuid.UUID  `json:"id"        binding:"required"`
	Value     string     `json:"value"     binding:"required"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type ProductOptionResponseDto struct {
	ID        uuid.UUID                       `json:"id"        binding:"required"`
	Name      string                          `json:"name"      binding:"required"`
	Values    []ProductOptionValueResponseDto `json:"values"    binding:"required"`
	DeletedAt *time.Time                      `json:"deletedAt"`
}

type ProductOptionValueResponseDto struct {
	ID        uuid.UUID  `json:"id"        binding:"required"`
	Value     string     `json:"value"     binding:"required"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type ProductVariantResponseDto struct {
	ID            uuid.UUID                       `json:"id"            binding:"required"`
	SKU           string                          `json:"sku"           binding:"required"`
	Price         int64                           `json:"price"         binding:"required"`
	Quantity      int                             `json:"quantity"      binding:"required"`
	PurchaseCount int                             `json:"purchaseCount" binding:"required"`
	CreatedAt     time.Time                       `json:"createdAt"     binding:"required"`
	UpdatedAt     time.Time                       `json:"updatedAt"     binding:"required"`
	DeletedAt     *time.Time                      `json:"deletedAt"`
	OptionValues  []ProductOptionValueResponseDto `json:"optionValues"  binding:"required"`
	Images        []ProductImageResponseDto       `json:"images"        binding:"required"`
}

type ProductImageResponseDto struct {
	ID        uuid.UUID  `json:"id"        binding:"required"`
	URL       string     `json:"url"       binding:"required"`
	Order     int        `json:"order"     binding:"required"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// ToProductResponseDto maps a domain.Product to ProductResponseDto
// Note: Category and Attributes need to be populated separately
func ToProductResponseDto(p *domain.Product) *ProductResponseDto {
	if p == nil {
		return nil
	}

	options := make([]ProductOptionResponseDto, 0, len(p.Options))
	for _, opt := range p.Options {
		options = append(options, *ToProductOptionResponseDto(&opt))
	}

	variants := make([]ProductVariantResponseDto, 0, len(p.Variants))
	for _, v := range p.Variants {
		variants = append(variants, *ToProductVariantResponseDto(&v))
	}

	images := make([]ProductImageResponseDto, 0, len(p.Images))
	for _, img := range p.Images {
		images = append(images, *ToProductImageResponseDto(&img))
	}

	var deletedAt *time.Time
	if !p.DeletedAt.IsZero() {
		deletedAt = &p.DeletedAt
	}
	return &ProductResponseDto{
		ID:            p.ID,
		Name:          p.Name,
		Description:   p.Description,
		ViewsCount:    p.ViewsCount,
		TotalPurchase: p.TotalPurchase,
		Price:         float64(p.Price),
		Rating:        p.Rating,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
		DeletedAt:     deletedAt,
		Category:      ProductCategoryResponseDto{},    // To be populated separately
		Attributes:    []ProductAttributeResponseDto{}, // To be populated separately
		Options:       options,
		Variants:      variants,
		Images:        images,
	}
}

// ToProductOptionResponseDto maps a domain.Option to ProductOptionResponseDto
func ToProductOptionResponseDto(o *domain.Option) *ProductOptionResponseDto {
	if o == nil {
		return nil
	}

	values := make([]ProductOptionValueResponseDto, 0, len(o.Values))
	for _, v := range o.Values {
		values = append(values, *ToProductOptionValueResponseDto(&v))
	}

	var deletedAt *time.Time
	if !o.DeletedAt.IsZero() {
		deletedAt = &o.DeletedAt
	}

	return &ProductOptionResponseDto{
		ID:        o.ID,
		Name:      o.Name,
		Values:    values,
		DeletedAt: deletedAt,
	}
}

// ToProductOptionValueResponseDto maps a domain.OptionValue to ProductOptionValueResponseDto
func ToProductOptionValueResponseDto(ov *domain.OptionValue) *ProductOptionValueResponseDto {
	if ov == nil {
		return nil
	}

	var deletedAt *time.Time
	if !ov.DeletedAt.IsZero() {
		deletedAt = &ov.DeletedAt
	}
	return &ProductOptionValueResponseDto{
		ID:        ov.ID,
		Value:     ov.Value,
		DeletedAt: deletedAt,
	}
}

// ToProductVariantResponseDto maps a domain.ProductVariant to ProductVariantResponseDto
func ToProductVariantResponseDto(v *domain.ProductVariant) *ProductVariantResponseDto {
	if v == nil {
		return nil
	}

	optionValues := make([]ProductOptionValueResponseDto, 0, len(v.OptionValues))
	for _, ov := range v.OptionValues {
		optionValues = append(optionValues, *ToProductOptionValueResponseDto(&ov))
	}

	images := make([]ProductImageResponseDto, 0, len(v.Images))
	for _, img := range v.Images {
		images = append(images, *ToProductImageResponseDto(&img))
	}

	var deletedAt *time.Time
	if !v.DeletedAt.IsZero() {
		deletedAt = &v.DeletedAt
	}
	return &ProductVariantResponseDto{
		ID:            v.ID,
		SKU:           v.SKU,
		Price:         v.Price,
		Quantity:      v.Quantity,
		PurchaseCount: v.PurchaseCount,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
		DeletedAt:     deletedAt,
		OptionValues:  optionValues,
		Images:        images,
	}
}

// ToProductImageResponseDto maps a domain.ProductImage to ProductImageResponseDto
func ToProductImageResponseDto(img *domain.ProductImage) *ProductImageResponseDto {
	if img == nil {
		return nil
	}

	var deletedAt *time.Time
	if !img.DeletedAt.IsZero() {
		deletedAt = &img.DeletedAt
	}

	return &ProductImageResponseDto{
		ID:        img.ID,
		URL:       img.URL,
		Order:     img.Order,
		CreatedAt: img.CreatedAt,
		DeletedAt: deletedAt,
	}
}

// ToProductResponseDtoList maps a slice of domain.Product to a slice of ProductResponseDto
func ToProductResponseDtoList(products []domain.Product) []ProductResponseDto {
	result := make([]ProductResponseDto, 0, len(products))
	for _, p := range products {
		dto := ToProductResponseDto(&p)
		if dto != nil {
			result = append(result, *dto)
		}
	}
	return result
}

// ToProductVariantResponseDtoList maps a slice of domain.ProductVariant to a slice of ProductVariantResponseDto
func ToProductVariantResponseDtoList(variants []domain.ProductVariant) []ProductVariantResponseDto {
	result := make([]ProductVariantResponseDto, 0, len(variants))
	for _, v := range variants {
		dto := ToProductVariantResponseDto(&v)
		if dto != nil {
			result = append(result, *dto)
		}
	}
	return result
}

// ToProductOptionResponseDtoList maps a slice of domain.Option to a slice of ProductOptionResponseDto
func ToProductOptionResponseDtoList(options []domain.Option) []ProductOptionResponseDto {
	result := make([]ProductOptionResponseDto, 0, len(options))
	for _, o := range options {
		dto := ToProductOptionResponseDto(&o)
		if dto != nil {
			result = append(result, *dto)
		}
	}
	return result
}

// ToProductOptionValueResponseDtoList maps a slice of domain.OptionValue to a slice of ProductOptionValueResponseDto
func ToProductOptionValueResponseDtoList(optionValues []domain.OptionValue) []ProductOptionValueResponseDto {
	result := make([]ProductOptionValueResponseDto, 0, len(optionValues))
	for _, ov := range optionValues {
		dto := ToProductOptionValueResponseDto(&ov)
		if dto != nil {
			result = append(result, *dto)
		}
	}
	return result
}

// ToProductImageResponseDtoList maps a slice of domain.ProductImage to a slice of ProductImageResponseDto
func ToProductImageResponseDtoList(images []domain.ProductImage) []ProductImageResponseDto {
	result := make([]ProductImageResponseDto, 0, len(images))
	for _, img := range images {
		dto := ToProductImageResponseDto(&img)
		if dto != nil {
			result = append(result, *dto)
		}
	}
	return result
}

// ToProductCategoryResponseDto maps a domain.Category to ProductCategoryResponseDto
func ToProductCategoryResponseDto(c *domain.Category) *ProductCategoryResponseDto {
	if c == nil {
		return nil
	}

	var deletedAt *time.Time
	if !c.DeletedAt.IsZero() {
		deletedAt = &c.DeletedAt
	}
	return &ProductCategoryResponseDto{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

// ToProductAttributeResponseDto maps a domain.Attribute and specific value to ProductAttributeResponseDto
func ToProductAttributeResponseDto(attr *domain.Attribute, valueID uuid.UUID) *ProductAttributeResponseDto {
	if attr == nil {
		return nil
	}

	// Find the specific attribute value
	var attrValue *domain.AttributeValue
	for _, v := range attr.Values {
		if v.ID == valueID {
			attrValue = &v
			break
		}
	}

	if attrValue == nil {
		return nil
	}

	var deletedAt *time.Time
	if !attr.DeletedAt.IsZero() {
		deletedAt = &attr.DeletedAt
	}
	var valueDeletedAt *time.Time
	if !attrValue.DeletedAt.IsZero() {
		valueDeletedAt = &attrValue.DeletedAt
	}
	return &ProductAttributeResponseDto{
		ID:   attr.ID,
		Code: attr.Code,
		Name: attr.Name,
		Value: ProductAttributeValueResponseDto{
			ID:        attrValue.ID,
			Value:     attrValue.Value,
			DeletedAt: valueDeletedAt,
		},
		DeletedAt: deletedAt,
	}
}

// WithCategory adds category information to ProductResponseDto
func (p *ProductResponseDto) WithCategory(category *domain.Category) *ProductResponseDto {
	if category != nil {
		p.Category = *ToProductCategoryResponseDto(category)
	}
	return p
}

// WithAttributes adds attribute information to ProductResponseDto
func (p *ProductResponseDto) WithAttributes(
	attributes []domain.Attribute,
	attributeValueIDs []uuid.UUID,
) *ProductResponseDto {
	// Create a map of attribute ID to attribute
	attrMap := make(map[uuid.UUID]*domain.Attribute)
	for i := range attributes {
		attrMap[attributes[i].ID] = &attributes[i]
	}

	// Create a map of value ID to attribute ID
	valueToAttrMap := make(map[uuid.UUID]uuid.UUID)
	for _, attr := range attributes {
		for _, val := range attr.Values {
			valueToAttrMap[val.ID] = attr.ID
		}
	}

	// Build the attribute response list
	attributeResponses := make([]ProductAttributeResponseDto, 0, len(attributeValueIDs))
	for _, valueID := range attributeValueIDs {
		if attrID, exists := valueToAttrMap[valueID]; exists {
			if attr, exists := attrMap[attrID]; exists {
				if attrDto := ToProductAttributeResponseDto(attr, valueID); attrDto != nil {
					attributeResponses = append(attributeResponses, *attrDto)
				}
			}
		}
	}

	p.Attributes = attributeResponses
	return p
}
