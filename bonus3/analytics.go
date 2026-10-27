package analytics

type Sale struct {
	Product          string
	Category         string
	Quantity         int
	UnitPriceKopecks int64
	Cancelled        bool
}

type Report struct {
	RevenueKopecks    int64
	UnitsSold         int
	RevenueByCategory map[string]int64
	TopProduct        string
	InvalidSales      int
}

func (s Sale) Incorrect() bool {
	condition := len(s.Category) == 0 || len(s.Product) == 0 || s.Quantity <= 0 || s.UnitPriceKopecks <= 0
	return condition
}

func AnalyzeSales(sales []Sale) Report {
	// TODO: реализуйте расчёт отчёта по правилам задания.
	var report Report = Report{RevenueByCategory: make(map[string]int64)}
	productsQuantity := make(map[string]int64)
	for _, sale := range sales {
		if sale.Cancelled {
			continue
		}
		if sale.Incorrect() {
			report.InvalidSales++
			continue
		}
		rev := int64(sale.Quantity) * sale.UnitPriceKopecks
		report.RevenueKopecks += rev
		report.UnitsSold += sale.Quantity
		report.RevenueByCategory[sale.Category] += rev
		productsQuantity[sale.Product] += int64(sale.Quantity)
	}
	var (
		maxQuentity int64
		topProduct  string
	)
	for product, quantity := range productsQuantity {
		if quantity > maxQuentity || (quantity == maxQuentity && product < topProduct) {
			maxQuentity = quantity
			topProduct = product
		}
	}
	report.TopProduct = topProduct
	return report
}
