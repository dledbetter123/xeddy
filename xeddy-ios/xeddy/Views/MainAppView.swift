//
//  MainAppView.swift
//  xeddy-ios
//
//  Created by David Ledbetter on 11/1/24.
//

import SwiftUI
import FirebaseAuth
import Foundation

struct MainAppView: View {
    @State private var isLoggedIn = Auth.auth().currentUser != nil

    var body: some View {
        VStack {
            if isLoggedIn {
                // Main content when logged in
                NavigationView {
                    ContentView() // This can be your HomeScreen or any other main screen
                        .navigationTitle("Home")
                        .toolbar {
                            ToolbarItem(placement: .navigationBarTrailing) {
                                Button(action: signOut) {
                                    Text("Log Out")
                                        .foregroundColor(.red)
                                }
                            }
                        }
                }
            } else {
                // Show the AuthView if not logged in
                AuthView(isLoggedIn: $isLoggedIn)
            }
        }
    }

    private func signOut() {
        do {
            try Auth.auth().signOut()
            isLoggedIn = false
        } catch {
            print("Error signing out: \(error.localizedDescription)")
        }
    }
}
