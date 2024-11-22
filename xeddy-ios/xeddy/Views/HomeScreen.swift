import SwiftUI

struct HomeScreen: View {
    @StateObject private var viewModel = RestaurantViewModel() // Use the view model to fetch real data

    var body: some View {
        NavigationView {
            VStack {
                Text("Restaurants")
                    .font(.system(size: 34, weight: .bold, design: .default))
                    .foregroundColor(.yellow) // Replace with `.umbcGold` if using a custom color
                
                ScrollView {
                    ForEach(viewModel.restaurants) { restaurant in
                        NavigationLink(destination: RestaurantDetailScreen(restaurant: restaurant)) {
                            RestaurantCard(restaurant: restaurant)
                        }
                        .buttonStyle(PlainButtonStyle()) // Removes default button styling
                    }
                }
                .padding(.horizontal)
                .refreshable { // Add pull-to-refresh here
                    viewModel.fetchRestaurants()
                }
            }
            .navigationBarHidden(true)
            .background(Color.black)
        }
        .onAppear {
            if viewModel.restaurants.isEmpty {
                viewModel.fetchRestaurants() // Fetch restaurants when the view appears
            }
        }
    }
}
#Preview {
    HomeScreen()
}