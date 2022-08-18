package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	chosenCar := ""
	if option1 > option2 {
		chosenCar = option2
	} else {
		chosenCar = option1
	}

	return chosenCar + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	resellPrice := 0.0
	if age < 3 {
		resellPrice = originalPrice * 0.8
	} else if age >= 10 {
		resellPrice = originalPrice * 0.5
	} else {
		resellPrice = originalPrice * 0.7
	}
	return resellPrice
}
