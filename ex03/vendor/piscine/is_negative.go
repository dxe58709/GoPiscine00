/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   is_negative.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 13:53:21 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 13:55:26 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func IsNegative(nb int) {
	if nb < 0{
		ft.PrintRune('T')
	}else{
		ft.PrintRune('F')
	}
	ft.PrintRune('\n')
}
