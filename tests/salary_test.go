package tests

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestSalary(t *testing.T) {
	//BigDecimal wages = new BigDecimal("12400").multiply(new BigDecimal("9"));
	//BigDecimal consumption = new BigDecimal("4000").multiply(new BigDecimal("9"));
	basic := decimal.NewFromFloat(197908).Add(decimal.NewFromFloat(115517))
	salary := decimal.NewFromFloat(12400).Mul(decimal.NewFromFloat(9))
	debt := decimal.NewFromFloat(1500).
		Add(decimal.NewFromFloat(6000)).
		Add(decimal.NewFromFloat(4000))

	comsume := decimal.NewFromFloat(4000).Mul(decimal.NewFromFloat(9))

	sum := basic.Add(salary).Add(debt).Sub(comsume)
	t.Logf("sum: %v", sum)
}
