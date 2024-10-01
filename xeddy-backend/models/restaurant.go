// const (
//     first = iota  // first is 0
//     second        // second is 1
//     third         // third is 2
// )

// func main() {
//     fmt.Println(first, second, third)  // Outputs: 0 1 2
// }

// models/restaurant.go

package models

import "time"

// Use string instead of custom WeekDay type for Firestore compatibility
type TimeObject struct {
	Open  string
	Close string
}

type Item struct {
	ID          int64
	Name        string
	Description string
	Price       float64
}

type Deal struct {
	ID          int64
	Description string
	ValidUntil  time.Time
}

type Review struct {
	ID     int64
	Rating int
	Text   string
}

type Restaurant struct {
	ID               int64
	Name             string
	Category         string
	ImageURL         string
	LogoURL          string
	Location         string
	DescriptionLong  string
	DescriptionShort string
	Email            string
	Phone            string
	OddDates         []time.Time
	Hours            map[string]TimeObject // Use string instead of WeekDay
	Closed           bool
	Menu             []Item
	Rewards          []int64
	Deals            []Deal
	OrderSystems     []string
	Reviews          []Review
}
