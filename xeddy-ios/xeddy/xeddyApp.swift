//
//  xeddyApp.swift
//  xeddy
//
//  Created by David Ledbetter on 9/4/24.
//

import SwiftUI
import FirebaseCore
import FirebaseFirestore
import FirebaseAuth
import UIKit

class AppDelegate: NSObject, UIApplicationDelegate {
  func application(_ application: UIApplication,
                   didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey : Any]? = nil) -> Bool {
    FirebaseApp.configure()
    FirebaseConfiguration.shared.setLoggerLevel(.debug)
    return true
  }
}

//@main
//struct YourApp: App {
//
//  var body: some Scene {
//    WindowGroup {
//      NavigationView {
//        ContentView()
//      }
//    }
//  }
//}

@main
struct xeddyApp: App {
    
    // register app delegate for Firebase setup
    @UIApplicationDelegateAdaptor(AppDelegate.self) var delegate
    var body: some Scene {
        WindowGroup {
            MainAppView()
        }
        .environment(\.colorScheme, .dark)
    }
}
