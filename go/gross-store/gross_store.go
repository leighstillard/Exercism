package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	list := map[string]int{}

	list["quarter_of_a_dozen"] = 3
	list["half_of_a_dozen"] = 6
	list["dozen"] = 12
	list["small_gross"] = 120
	list["gross"] = 144
	list["great_gross"] = 1728

	return list
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return false
	}

	bill[item] += value
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return false
	}
	_, exists = bill[item]
	if !exists {
		return false
	}
	if bill[item] < value {
		return false
	}
	if bill[item] == value {
		delete(bill, item)
		return true
	}
	bill[item] -= value
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	if !exists {
		return 0, false
	}
	return value, true
}
