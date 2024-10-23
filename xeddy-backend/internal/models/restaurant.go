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

func ExampleRestaurant() Restaurant {
    return Restaurant{
        ID:               1,
        Name:             "Gusto Italian Grill",
        Category:         "Italian",
        ImageURL:         "https://example.com/image.jpg",
        LogoURL:          "https://example.com/logo.jpg",
        Location:         "1234 Culinary Blvd, Taste City, TC 56789",
        DescriptionLong:  "Gusto Italian Grill offers a rich taste of Italy with a diverse menu featuring traditional dishes, handpicked wines, and a cozy ambiance perfect for any occasion.",
        DescriptionShort: "Traditional Italian cuisine in the heart of Taste City.",
        Email:            "contact@gustoitalian.com",
        Phone:            "555-1234",
        OddDates:         []time.Time{time.Now(), time.Now().AddDate(0, 0, 1)},
        Hours: map[string]TimeObject{
            "Monday":  {Open: "10:00:00", Close: "22:00:00"},
            "Tuesday": {Open: "10:00:00", Close: "22:00:00"},
        },
        Closed: false,
        Menu: []Item{
            {ID: 1, Name: "Spaghetti Carbonara", Description: "Creamy pasta with pancetta, egg, and Parmesan", Price: 15.99},
            {ID: 2, Name: "Margherita Pizza", Description: "Classic pizza with fresh tomatoes, mozzarella, and basil", Price: 12.50},
        },
        Rewards: []int64{100, 200},
        Deals: []Deal{
            {ID: 1, Description: "20% off on Wednesdays", ValidUntil: time.Now().AddDate(0, 1, 0)},
        },
        OrderSystems: []string{"Toast", "Square"},
        Reviews: []Review{
            {ID: 1, Rating: 5, Text: "Outstanding service and food quality!"},
        },
    }
}