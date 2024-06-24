/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   rev_alpha.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 13:23:51 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 13:51:14 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func ReverseAlphabet() {
	for i:= 'z'; 'a' <= i; i-- {
		ft.PrintRune(i)
	}
	ft.PrintRune('\n')
}