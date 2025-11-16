package application

type Product struct {
	productRepo   ProductRepository
	attributeRepo AttributeRepository
	service       ProductService
}

func (p *Product) ListProducts() {
	products := p.productRepo.ListProducts()
	productVariants := p.productRepo.ListProductVariants()
	// Ghep do
	attributeValues := p.attributeRepo.ListAttributesValues(products.IDs) // Map sang array
	p.service.ListAttributesByProductIDs(products.IDs)
}
