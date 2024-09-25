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

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type TimeObject struct {
	Open  time.Time
	Close time.Time
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

// Restaurant represents a restaurant structure
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
	Hours            map[WeekDay]TimeObject
	Closed           bool
	Menu             []Item
	Rewards          []int64
	Deals            []Deal
	OrderSystems     []string // 'Toast', 'Square', whatever
	Reviews          []Review
}
