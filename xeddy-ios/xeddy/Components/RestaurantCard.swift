//
//  RestaurantCard.swift
//  xeddy
//
//  Created by David Ledbetter on 9/10/24.
//

import Foundation
import SwiftUI
import FirebaseStorage

func fetchDownloadURL(merchantID: String, filename: String, completion: @escaping (URL?) -> Void) {
    let storageRef = Storage.storage().reference().child("\(merchantID)/\(filename)")

    storageRef.downloadURL { url, error in
        if let error = error {
            print("Error fetching download URL: \(error.localizedDescription)")
            completion(nil)
            return
        }
        
        print("Fetched download URL: \(url?.absoluteString ?? "No URL")")
        completion(url)
    }
}

struct RestaurantCard: View {
    let restaurant: Restaurant
    @State private var imageURL: URL?

    var body: some View {
        VStack {
            // Display image once URL is available
            if let imageURL = imageURL {
                AsyncImage(url: imageURL) { image in
                    image.resizable()
                         .aspectRatio(contentMode: .fit)
                         .frame(height: 200)
                } placeholder: {
                    ProgressView() // Shows a loading spinner while fetching
                }
            } else {
                Rectangle()
                    .fill(Color.gray)
                    .frame(width: 100, height: 100) // Placeholder if no image is available
            }

            Text(restaurant.name)
                .font(.headline)

        }
        .onAppear {
            loadImage()
        }
    }

    private func loadImage() {
        // Fetch the download URL dynamically based on merchantID and filename
        fetchDownloadURL(merchantID: restaurant.merchantID, filename: "card_0.jpg") { url in
            self.imageURL = url
        }
    }
}

