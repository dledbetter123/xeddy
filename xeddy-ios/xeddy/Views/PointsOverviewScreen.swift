//
//  PointsOverviewScreen.swift
//  xeddy
//
//  Created by David Ledbetter on 9/10/24.
//

import SwiftUI
import Foundation

struct PointsOverviewScreen: View {
    @State private var totalPoints = 300

    var body: some View {
        VStack {
            Text("Your Points")
                .font(.largeTitle)
                .fontWeight(.bold)
                .padding(.top)
            
            Text("\(totalPoints) Total Points")
                .font(.title)
                .foregroundColor(.umbcGold)
                .padding(.vertical)
            
            Text("Points available for next order:")
                .font(.subheadline)
                .foregroundColor(.gray)
            
            Text("\(totalPoints / 2) Points")
                .font(.title2)
                .padding(.top)
            
            Spacer()
        }
        .padding()
    }
}

