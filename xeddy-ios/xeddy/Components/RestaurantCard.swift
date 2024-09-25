//
//  RestaurantCard.swift
//  xeddy
//
//  Created by David Ledbetter on 9/10/24.
//

import Foundation
import SwiftUI

struct RestaurantCard: View {
    let restaurant: Restaurant
    
    var body: some View {
        VStack(alignment: .leading) {
            Image(restaurant.image)
                .resizable()
                .aspectRatio(contentMode: .fill)
                .frame(height: 150)
                .clipped()
                .cornerRadius(10)
            
            HStack {
                VStack(alignment: .leading) {
                    Text(restaurant.name)
                        .font(.headline)
                    Text("\(restaurant.points) Points Available")
                        .font(.subheadline)
                        .foregroundColor(.gray)
                }
                Spacer()
                if restaurant.hasDeal {
                    Text("Deal!")
                        .font(.caption)
                        .fontWeight(.bold)
                        .padding(5)
                        .background(Color.red)
                        .foregroundColor(.white)
                        .cornerRadius(5)
                }
            }
        }
        .padding()
        .background(Color.black.opacity(0.1))
        .cornerRadius(10)
        .shadow(radius: 5)
    }
}
