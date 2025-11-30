package domain

func ValidateProductVariantStructure(product *Product) bool {
	if len(product.Options) == 0 { // 1, 2
		if len(product.Variants) != 1 { // 3, 4
			return false // 5
		}
		if len(product.Variants[0].OptionValues) != 0 { // 6, 7
			return false // 8
		}
		return true // 9
	}
	for _, variant := range product.Variants { // 10
		if len(variant.OptionValues) != len(product.Options) { // 11, 12
			return false // 13
		}
	}
	return true // 14
}
