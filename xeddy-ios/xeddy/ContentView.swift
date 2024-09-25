//
//  ContentView.swift
//  xeddy
//
//  Created by David Ledbetter on 9/4/24.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
        TabView {
            HomeScreen()
                .tabItem {
                    Image(systemName: "house")
                    Text("Home")
                }

            PointsOverviewScreen()
                .tabItem {
                    Image(systemName: "star")
                    Text("Points")
                }
        }
        .background(Color.black)
    }
}


#Preview {
    ContentView()
}
