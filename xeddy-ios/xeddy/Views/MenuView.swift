//
//  MenuView.swift
//  xeddy
//
//  Created by David Ledbetter on 9/10/24.
//

import Foundation
import SwiftUI

struct MenuView: View {
    var menuItems: [MenuItem]
    
    var body: some View {
         VStack(alignment: .leading) {
             Text("Menu")
                 .font(.title2)
                 .fontWeight(.bold)

             // Ensure this ForEach is used without any bindings or unnecessary wrappers
             ForEach(menuItems) { item in
                 HStack(alignment: .top) {
                     if let imageURL = item.imageURL {
                         AsyncImage(url: imageURL) { image in
                             image.resizable()
                                 .frame(width: 100, height: 100)
                                 .clipShape(RoundedRectangle(cornerRadius: 10))
                         } placeholder: {
                             ProgressView()
                                 .frame(width: 100, height: 100)
                         }
                     } else {
                         Rectangle()
                             .fill(Color.gray)
                             .frame(width: 100, height: 100)
                             .clipShape(RoundedRectangle(cornerRadius: 10))
                     }

                     VStack(alignment: .leading, spacing: 5) {
                         Text(item.name)
                             .fontWeight(.semibold)
                         Text(item.description)
                             .foregroundColor(.gray)
                             .lineLimit(2)
                         Text("$\(Double(item.priceMoney) / 100, specifier: "%.2f")")
                             .font(.subheadline)
                             .fontWeight(.bold)
                             .foregroundColor(.primary)
                     }
                     Spacer()
                 }
                 .padding(.vertical, 5)
                 Divider()
             }
         }
         .padding(.horizontal)
     }
}
