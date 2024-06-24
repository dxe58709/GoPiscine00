/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcomb2.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 13:58:17 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 14:41:06 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func digitToRune(i int, j int) {
	ft.PrintRune(rune(i + '0'))
	ft.PrintRune(rune(j + '0'))
}

func PrintComb2() {
	for i := 10; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			if i == 10 && j == 11 {
				digitToRune(i / 10, j % 10)
				ft.PrintRune(' ')
				digitToRune(j / 10, i % 10)
			} else {
				ft.PrintRune(',')
				ft.PrintRune(' ')
				digitToRune(i / 10, i % 10)
				ft.PrintRune(' ')
				digitToRune(j / 10, j % 10)
			}
		}
	}
	ft.PrintRune('\n')
}