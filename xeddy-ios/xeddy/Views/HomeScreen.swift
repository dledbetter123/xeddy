import SwiftUI

struct HomeScreen: View {
    
    // placeholder for restaurants
    let restaurants = [
        Restaurant(name: "Local Pizza", points: 100, image: "pizza", hasDeal: true),
        Restaurant(name: "Sushi Spot", points: 50, image: "sushi", hasDeal: false),
        Restaurant(name: "Burger Barn", points: 120, image: "burger", hasDeal: true)
    ]
    
    var body: some View {
        NavigationView {
            VStack {
                Text("Restaurants")
                    .font(.system(size: 34, weight: .bold, design: .default)) // Combines size and weight
                    .foregroundColor(.umbcGold)
                ScrollView {
                    ForEach(restaurants) { restaurant in
                        NavigationLink(destination: RestaurantDetailScreen(restaurant: restaurant)) {
                            RestaurantCard(restaurant: restaurant)
                        }
                        .buttonStyle(PlainButtonStyle()) // Removes default button styling
                    }
                }
                .padding(.horizontal)
            }
            .navigationBarHidden(true) // Hide default navigation bar
            .background(Color.black)
        }
    }
    
}
