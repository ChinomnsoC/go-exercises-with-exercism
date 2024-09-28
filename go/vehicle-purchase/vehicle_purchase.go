package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck"{
        return true
    } else {
        return false
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    var answer string
	if option1 > option2 {
        answer = option2
    } else if option2 > option1 {
        answer = option1
    }
    return fmt.Sprintf("%s is clearly the better choice.", answer)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    var eightyPercent float64 = 0.80
    var seventyPercent float64 = 0.70
    var price float64
	if age < 3 {
        price = originalPrice * eightyPercent
    } else if age >= 10 {
        price = originalPrice  / 2
    } else if age >= 3 && age < 10 {
        price = originalPrice * seventyPercent
    }
    return price
}