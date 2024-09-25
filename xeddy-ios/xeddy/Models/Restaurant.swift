//
//  Restaurant.swift
//  xeddy
//
//  Created by David Ledbetter on 9/10/24.
//

import Foundation

struct Restaurant: Identifiable {
    let id = UUID()
    let name: String
    let points: Int
    let image: String
    let hasDeal: Bool
}
