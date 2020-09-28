package stats

import (
	"reflect"
	"testing"

	"github.com/JovidYnwa/bank/v2/pkg/types"
)

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
	}
}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "found"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}

	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
	}
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 2, Category: "food"},
	}

	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}

	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "Category1", Amount: 1_000_00},
		{ID: 2, Category: "Category2", Amount: 2_000_00},
		{ID: 3, Category: "Category1", Amount: 3_000_00},
		{ID: 4, Category: "Category1", Amount: 4_000_00},
		{ID: 5, Category: "Category3", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"Category1": 8_000_00,
		"Category2": 2_000_00,
		"Category3": 5_000_00,
	}
	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic(t *testing.T) {
	firstPaymet := map[types.Category]types.Money{
		"Category1": 10,
		"Category2": 20,
	}

	secondPayment := map[types.Category]types.Money{
		"Category1": 10,
		"Category2": 15,
		"Category3": 5,
	}

	expected := map[types.Category]types.Money{
		"Category1": 0,
		"Category2": -5,
		"Category3": 5,
	}

	result := PeriodsDynamic(firstPaymet, secondPayment)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result: %v, actual: %v", expected, result)
	}
}
