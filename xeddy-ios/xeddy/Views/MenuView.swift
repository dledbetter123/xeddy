//
//  MenuView.swift
//  xeddy
//
//  Created by David Ledbetter on 9/10/24.
//

import Foundation
import SwiftUI

struct MenuView: View {
    var body: some View {
        VStack(alignment: .leading) {
            Text("Menu")
                .font(.title2)
                .fontWeight(.bold)
            
            ForEach(0..<5) { index in
                HStack {
                    Text("Item \(index + 1)")
                    Spacer()
                    Text("\(index * 5) Points")
                        .foregroundColor(.gray)
                }
                .padding(.vertical, 10)
                Divider()
            }
        }
    }
}
