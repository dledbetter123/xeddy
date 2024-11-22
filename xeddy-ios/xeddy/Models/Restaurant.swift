//
//  Restaurant.swift
//  xeddy
//
//  Created by David Ledbetter on 9/10/24.
//

// need coding keys to ensure the (Codable, Identifiable) types remain valid for swift
// syntax checks.

import Foundation

struct MenuItem: Identifiable, Codable {
    let id: String
    let name: String
    let description: String
    let priceMoney: Int // needs to updated to support priceMoney handler.
    var imageURL: URL?

    enum CodingKeys: String, CodingKey {
        case id
        case name
        case description
        case priceMoney = "price_money"
        case imageURL = "image_url"
    }

}

struct LineItem: Codable {
    let id: String
    let quantity: String
    let base_price_money: Money
}

struct Money: Codable {
    let amount: Int // Amount in cents (e.g., $9.95 should be 995)
    let currency: String // "USD"
}

extension MenuItem {
    func toLineItem(quantity: Int) -> LineItem {
        return LineItem(
            name: self.name,
            quantity: String(quantity),
            base_price_money: Money(amount: self.price_money, currency: "USD")
        )
    }
}

struct TimeObject: Codable {
    let open: String
    let close: String

    enum CodingKeys: String, CodingKey {
        case open = "Open"
        case close = "Close"
    }
}

struct Item: Identifiable, Codable {
    let id: Int // ID as an integer
    let name: String
    let description: String
    let price: Double // Double to handle decimal prices

    enum CodingKeys: String, CodingKey {
        case id = "ID"
        case name = "Name"
        case description = "Description"
        case price = "Price"
    }
}

struct Deal: Codable, Identifiable {
    let id: Int64
    let description: String
    let validUntil: Date

    enum CodingKeys: String, CodingKey {
        case id = "ID"
        case description = "Description"
        case validUntil = "ValidUntil"
    }
}

struct Review: Codable, Identifiable {
    let id: Int64
    let rating: Int
    let text: String

    enum CodingKeys: String, CodingKey {
        case id = "ID"
        case rating = "Rating"
        case text = "Text"
    }
}

struct Restaurant: Codable, Identifiable {
    let id: Int64
    let name: String
    let category: String
    let imageURL: String
    let logoURL: String
    let location: String
    let descriptionLong: String
    let descriptionShort: String
    let email: String
    let phone: String
    let merchantID: String
    let oddDates: [Date]?
    let hours: [String: TimeObject]?
    let closed: Bool
    let menu: [Item]?
    let rewards: [Int64]?
    let deals: [Deal]?
    let orderSystems: [String]?
    let reviews: [Review]?

    enum CodingKeys: String, CodingKey {
        case id = "ID"
        case name = "Name"
        case category = "Category"
        case imageURL = "ImageURL"
        case logoURL = "LogoURL"
        case location = "Location"
        case descriptionLong = "DescriptionLong"
        case descriptionShort = "DescriptionShort"
        case email = "Email"
        case phone = "Phone"
        case merchantID = "MerchantID"
        case oddDates = "OddDates"
        case hours = "Hours"
        case closed = "Closed"
        case menu = "Menu"
        case rewards = "Rewards"
        case deals = "Deals"
        case orderSystems = "OrderSystems"
        case reviews = "Reviews"
    }
}
