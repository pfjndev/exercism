package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    return map[string]int {
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    if _, exists := units[unit]; !exists {
        return false
    }
    // bill[item] defaults to 0 if it doesn't exist.
    // Add the unit value directly.
    bill[item] += units[unit]
    return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    /* // 1. Return false if the given item is not in the bill
    if _, exists := bill[item]; !exists {
        return false
    }

    // 2. Return false if the given unit is not in the units map.
    if _, exists := units[unit]; !exists {
        return false
    }

    // 3. Return false if the new quantity would be less than 0.
    if bill[item] - units[unit] < 0 { return false }

    // 4. If the new quantity is 0, remove the item from the bill, then return true.
    if bill[item] - units[unit] == 0 { 
        delete(bill, item)
        return true
    }

    // 5. Otherwise, reduce the quantity of the item and return true.
    bill[item] -= units[unit]
    return true
    */
    // Look up both values once
	uQty, unitExists := units[unit]
	bQty, itemExists := bill[item]

	if !unitExists || !itemExists {
		return false
	}

	newQty := bQty - uQty

	switch {
	case newQty < 0:
		return false
	case newQty == 0:
		delete(bill, item)
	default:
		bill[item] = newQty
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    // 1. Return 0 and false if the item is not in the bill.
    /* if _, exists := bill[item]; !exists {
        return 0, false
    } */
    // 2. Otherwise, return the quantity of the item in the bill and true.
    qnt, ok := bill[item]
    return qnt, ok
}
