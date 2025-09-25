// Package greek provides quick access to the greek alphabet.
//
// See Lower and Upper
package greek

import (
	"core/enum/greek/lower"
	"core/enum/greek/upper"
)

// Characters provides a slice of all Greek characters considered by JanOS.
//
// See Characters, Lower, and Upper
var Characters = append(lower.Characters, upper.Characters...)

// IsGreekCharacter returns true if the provided rune is a Greek character.
//
// See Characters, Lower, and Upper
func IsGreekCharacter(c rune) bool {
	for _, l := range Characters {
		if l == string(c) {
			return true
		}
	}
	return false
}

type _lower struct {
	// Alpha represents the Greek character "α".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Alpha string
	// Beta represents the Greek character "β".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Beta string
	// Gamma represents the Greek character "γ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Gamma string
	// Delta represents the Greek character "δ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Delta string
	// Epsilon represents the Greek character "ε".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Epsilon string
	// Zeta represents the Greek character "ζ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Zeta string
	// Eta represents the Greek character "Η".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Eta string
	// Theta represents the Greek character "θ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Theta string
	// Iota represents the Greek character "ι".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Iota string
	// Kappa represents the Greek character "κ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Kappa string
	// Lambda represents the Greek character "λ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Lambda string
	// Mu represents the Greek character "μ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Mu string
	// Nu represents the Greek character "ν".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Nu string
	// Xi represents the Greek character "ξ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Xi string
	// Omicron represents the Greek character "ο".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Omicron string
	// Pi represents the Greek character "π".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Pi string
	// Rho represents the Greek character "ρ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Rho string
	// Sigma represents the Greek character "σ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	//
	// NOTE: This is the form used at the beginning or in the middle of a word, see SigmaFinal for the final form.
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Sigma string
	// SigmaFinal represents the Greek character "ς".
	//
	// NOTE: This is the final form used at the end of a word, see Sigma for the mid-word form.
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	SigmaFinal string
	// Tau represents the Greek character "τ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Tau string
	// Upsilon represents the Greek character "υ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Upsilon string
	// Phi represents the Greek character "φ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Phi string
	// Chi represents the Greek character "χ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Chi string
	// Psi represents the Greek character "ψ".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Psi string
	// Omega represents the Greek character "ω".
	//
	// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
	Omega string
}

// Lower provides access to the lowercase Greek alphabet.
//
// See Characters, Lower, and Upper
var Lower = _lower{
	Alpha:      lower.Alpha,
	Beta:       lower.Beta,
	Gamma:      lower.Gamma,
	Delta:      lower.Delta,
	Epsilon:    lower.Epsilon,
	Zeta:       lower.Zeta,
	Eta:        lower.Eta,
	Theta:      lower.Theta,
	Iota:       lower.Iota,
	Kappa:      lower.Kappa,
	Lambda:     lower.Lambda,
	Mu:         lower.Mu,
	Nu:         lower.Nu,
	Xi:         lower.Xi,
	Omicron:    lower.Omicron,
	Pi:         lower.Pi,
	Rho:        lower.Rho,
	Sigma:      lower.Sigma,
	SigmaFinal: lower.SigmaFinal,
	Tau:        lower.Tau,
	Upsilon:    lower.Upsilon,
	Phi:        lower.Phi,
	Chi:        lower.Chi,
	Psi:        lower.Psi,
	Omega:      lower.Omega,
}

type _upper struct {
	// Alpha represents the Greek character "Α".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Alpha string
	// Beta represents the Greek character "Β".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Beta string
	// Gamma represents the Greek character "Γ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Gamma string
	// Delta represents the Greek character "Δ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Delta string
	// Epsilon represents the Greek character "Ε".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Epsilon string
	// Zeta represents the Greek character "Ζ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Zeta string
	// Eta represents the Greek character "Η".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Eta string
	// Theta represents the Greek character "Θ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Theta string
	// Iota represents the Greek character "Ι".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Iota string
	// Kappa represents the Greek character "Κ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Kappa string
	// Lambda represents the Greek character "Λ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Lambda string
	// Mu represents the Greek character "Μ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Mu string
	// Nu represents the Greek character "Ν".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Nu string
	// Xi represents the Greek character "Ξ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Xi string
	// Omicron represents the Greek character "Ο".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Omicron string
	// Pi represents the Greek character "Π".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Pi string
	// Rho represents the Greek character "Ρ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Rho string
	// Sigma represents the Greek character "Σ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Sigma string
	// Tau represents the Greek character "Τ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Tau string
	// Upsilon represents the Greek character "Υ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Upsilon string
	// Phi represents the Greek character "Φ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Phi string
	// Chi represents the Greek character "Χ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Chi string
	// Psi represents the Greek character "Ψ".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Psi string
	// Omega represents the Greek character "Ω".
	//
	// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
	Omega string
}

// Upper provides access to the uppercase Greek alphabet.
//
// See Characters, Lower, and Upper
var Upper = _upper{
	Alpha:   upper.Alpha,
	Beta:    upper.Beta,
	Gamma:   upper.Gamma,
	Delta:   upper.Delta,
	Epsilon: upper.Epsilon,
	Zeta:    upper.Zeta,
	Eta:     upper.Eta,
	Theta:   upper.Theta,
	Iota:    upper.Iota,
	Kappa:   upper.Kappa,
	Lambda:  upper.Lambda,
	Mu:      upper.Mu,
	Nu:      upper.Nu,
	Xi:      upper.Xi,
	Omicron: upper.Omicron,
	Pi:      upper.Pi,
	Rho:     upper.Rho,
	Sigma:   upper.Sigma,
	Tau:     upper.Tau,
	Upsilon: upper.Upsilon,
	Phi:     upper.Phi,
	Chi:     upper.Chi,
	Psi:     upper.Psi,
	Omega:   upper.Omega,
}
