//
//  AuthView.swift
//  xeddy-ios
//
//  Created by David Ledbetter on 11/1/24.
//

import Foundation

import SwiftUI
import FirebaseAuth

struct AuthView: View {
    @State private var email = ""
    @State private var password = ""
    @State private var errorMessage = ""
    @Binding var isLoggedIn: Bool
    
    var body: some View {
        VStack(spacing: 20) {
            Text("Xeddy Login")
                .font(.largeTitle)
                .padding()

            TextField("Email", text: $email)
                .padding()
                .background(Color(.secondarySystemBackground))
                .cornerRadius(8)

            SecureField("Password", text: $password)
                .padding()
                .background(Color(.secondarySystemBackground))
                .cornerRadius(8)

            if !errorMessage.isEmpty {
                Text(errorMessage)
                    .foregroundColor(.red)
            }

            Button(action: login) {
                Text("Log In")
                    .frame(maxWidth: .infinity)
                    .padding()
                    .background(Color.blue)
                    .foregroundColor(.white)
                    .cornerRadius(8)
            }

            Button(action: signUp) {
                Text("Sign Up")
                    .frame(maxWidth: .infinity)
                    .padding()
                    .background(Color.green)
                    .foregroundColor(.white)
                    .cornerRadius(8)
            }
        }
        .padding()
        .fullScreenCover(isPresented: $isLoggedIn) {
            // Destination view after login, e.g., the main app screen
            MainAppView()
        }
    }

    // Firebase Authentication - Log In
    private func login() {
        Auth.auth().signIn(withEmail: email, password: password) { authResult, error in
            if let error = error {
                errorMessage = "Error: \(error.localizedDescription)"
                print("Firebase Auth Error: \(error.localizedDescription)")
            } else {
                isLoggedIn = true
                errorMessage = "Login successful!"
            }
        }
    }

    // Firebase Authentication - Sign Up
    private func signUp() {
        Auth.auth().createUser(withEmail: email, password: password) { authResult, error in
            if let error = error {
                errorMessage = "Error: \(error.localizedDescription)"
                print("Firebase Auth Error: \(error.localizedDescription)")
            } else {
                isLoggedIn = true
                errorMessage = "Signup successful!"
            }
        }
    }
}
