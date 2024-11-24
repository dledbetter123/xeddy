//
//  MenuViewModel.swift
//  xeddy-ios
//
//  Created by David Ledbetter on 11/8/24.
//

import Foundation
import SwiftUI

class MenuViewModel: ObservableObject {
    @Published var menuItems: [MenuItem] = []
    @Published var isLoading: Bool = false
    @Published var errorMessage: String? = nil
    
    func fetchMenuItems(merchantID: String) {
        guard let url = URL(string: "https://xeddy-backend-452699980171.us-east1.run.app/menu?merchant_id=\(merchantID)") else {
            self.errorMessage = "Invalid URL"
            return
        }
        
        isLoading = true
        URLSession.shared.dataTask(with: url) { [weak self] data, response, error in
            DispatchQueue.main.async {
                self?.isLoading = false
                if let error = error {
                    self?.errorMessage = "Error fetching menu items: \(error.localizedDescription)"
                    return
                }
                
                guard let data = data else {
                    self?.errorMessage = "No data received"
                    return
                }
                
                do {
                    let decodedItems = try JSONDecoder().decode([MenuItem].self, from: data)
                    self?.menuItems = decodedItems
                } catch {
                    self?.errorMessage = "Failed to decode menu items: \(error.localizedDescription)"
                }
            }
        }.resume()
    }
}
