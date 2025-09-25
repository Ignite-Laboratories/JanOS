package blue

import "git.ignitelabs.net/janos/core/enum/sub"

// "𝑇ℎ𝑒 𝐵𝑙𝑢𝑒 𝑁𝑜𝑡𝑒" is an improvised value from 0-7.
//
// See sub.SubByte
func Note() byte {
	n := sub.NewNote()
	return byte(n.Value())
}
