//
//  RestaurantDetailsScreen.swift
//  xeddy
//
//  Created by David Ledbetter on 9/10/24.
//

import SwiftUI


struct RestaurantDetailScreen: View {
    let restaurant: Restaurant
    
    var body: some View {
        ScrollView {
            VStack {
                Image(restaurant.image)
                    .resizable()
                    .aspectRatio(contentMode: .fit)
                    .frame(height: 200)
                
                Text(restaurant.name)
                    .font(.largeTitle)
                    .fontWeight(.bold)
                    .padding(.top)
                
                Text("\(restaurant.points) Points Available")
                    .font(.title2)
                    .foregroundColor(.gray)
                
                Text("Hours of Operation: 9:00 AM - 10:00 PM")
                    .padding(.top)
                
                if restaurant.hasDeal {
                    HStack {
                        Text("Exclusive Deal: 20% off with points!")
                            .padding()
                            .background(Color.yellow)
                            .cornerRadius(10)
                    }
                    .padding(.top)
                }
                
                // Interactive menu
                MenuView()
                    .padding(.top)
            }
        }
        .padding()
    }
}
