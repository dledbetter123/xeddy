import Foundation
import Combine

class RestaurantViewModel: ObservableObject {
    @Published var restaurants: [Restaurant] = []
    private var cancellables = Set<AnyCancellable>()
    
    func fetchRestaurants() {
        guard let url = URL(string: "https://xeddy-backend-452699980171.us-east1.run.app/restaurants") else {
            print("Invalid URL")
            return
        }
        
        let decoder = JSONDecoder()
        
        // Custom date formatter to handle fractional seconds
        let dateFormatter = DateFormatter()
        dateFormatter.dateFormat = "yyyy-MM-dd'T'HH:mm:ss.SSSSSSZ" // Adjusted format
        decoder.dateDecodingStrategy = .formatted(dateFormatter)
        
        URLSession.shared.dataTaskPublisher(for: url)
            .map { data, response -> Data in
                if let httpResponse = response as? HTTPURLResponse {
                    print("Status Code: \(httpResponse.statusCode)")
                }
                print("Raw Response Data: \(String(data: data, encoding: .utf8) ?? "Unable to decode data")")
                return data
            }
            .decode(type: [Restaurant].self, decoder: decoder)
            .catch { error -> Just<[Restaurant]> in
                print("Decoding error: \(error.localizedDescription)")
                return Just([]) // Replace with an empty array on error
            }
            .receive(on: DispatchQueue.main)
            .sink { [weak self] restaurants in
                self?.restaurants = restaurants
                if restaurants.isEmpty {
                    print("No restaurants decoded.")
                } else {
                    print("Decoded Restaurants: \(restaurants)")
                }
            }
            .store(in: &cancellables)
    }
}
