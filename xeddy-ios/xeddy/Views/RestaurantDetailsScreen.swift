import SwiftUI
import FirebaseStorage

struct RestaurantDetailScreen: View {
    let restaurant: Restaurant
    @StateObject private var menuViewModel = MenuViewModel()
    @State private var imageURL: URL?
    
    var body: some View {
        ScrollView {
            VStack {
                if let imageURL = imageURL {
                    AsyncImage(url: imageURL) { image in
                        image.resizable()
                            .aspectRatio(contentMode: .fit)
                            .frame(height: 200)
                    } placeholder: {
                        ProgressView()
                    }
                } else {
                    Rectangle()
                        .fill(Color.gray)
                        .frame(height: 200)
                }
                
                Text(restaurant.name)
                    .font(.largeTitle)
                    .fontWeight(.bold)
                    .padding(.top)
                
                Text("100 Points Available")
                    .font(.title2)
                    .foregroundColor(.gray)
                
                Text("Hours of Operation: 9:00 AM - 10:00 PM")
                    .padding(.top)
                
                if hasDeal() {
                    HStack {
                        Text("Exclusive Deal: 20% off with points!")
                            .padding()
                            .background(Color.yellow)
                            .cornerRadius(10)
                    }
                    .padding(.top)
                }
                
                if menuViewModel.isLoading {
                    ProgressView()
                        .padding(.top)
                } else if let errorMessage = menuViewModel.errorMessage {
                    Text("Error: \(errorMessage)")
                        .foregroundColor(.red)
                        .padding(.top)
                } else {
                    MenuView(menuItems: menuViewModel.menuItems)
                        .padding(.top)
                }
            }
        }
        .padding()
        .onAppear {
            menuViewModel.fetchMenuItems(merchantID: restaurant.merchantID)
            loadImage()
        }
    }
    
    private func loadImage() {
        fetchDownloadURL(merchantID: restaurant.merchantID, filename: "card_0.jpg") { url in
            self.imageURL = url
        }
    }
    
    private func hasDeal() -> Bool {
        restaurant.deals?.count ?? 0 > 0
    }
}
