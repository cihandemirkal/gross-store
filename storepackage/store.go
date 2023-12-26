package store

// Units store the Gross Store unit measurement
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    12,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill create a new bill
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem add item to customer bill
func AddItem(bill map[string]int, units map[string]int, item string, unit string) bool {
	value, ok := units[unit]
	if !ok {
		return false
	}
	bill[item] += value
	return true
}

// Remove an item from the customer bill
func RemoveItem(bill map[string]int, units map[string]int, item string, unit string) bool {
	value, ok := units[unit]
	if !ok {
		return false
	}

	current, ok := bill[item]
	if !ok {
		return false
	}

	newVal := current - value
	if newVal < 0 {
		return false
	} else if newVal == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] = newVal
	}
	return true

}

// Return the quantity of a specific item that is in the customer bill
func GetItem(bill map[string]int, item string) (int, bool) {
	value, ok := bill[item]

	if !ok {
		return 0, false
	}

	return value, ok

}
