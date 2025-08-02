// Package greek provides quick access to the greek alphabet.
//
// See Lower and Upper
package greek

import (
	"github.com/ignite-laboratories/core/enum/greek/lower"
	"github.com/ignite-laboratories/core/enum/greek/upper"
)

type _lower struct {
	// Alpha represents the Greek character "α".
	Alpha string
	// Beta represents the Greek character "β".
	Beta string
	// Gamma represents the Greek character "γ".
	Gamma string
	// Delta represents the Greek character "δ".
	Delta string
	// Epsilon represents the Greek character "ε".
	Epsilon string
	// Zeta represents the Greek character "ζ".
	Zeta string
	// Eta represents the Greek character "Η".
	Eta string
	// Theta represents the Greek character "θ".
	Theta string
	// Iota represents the Greek character "ι".
	Iota string
	// Kappa represents the Greek character "κ".
	Kappa string
	// Lambda represents the Greek character "λ".
	Lambda string
	// Mu represents the Greek character "μ".
	Mu string
	// Nu represents the Greek character "ν".
	Nu string
	// Xi represents the Greek character "ξ".
	Xi string
	// Omicron represents the Greek character "ο".
	Omicron string
	// Pi represents the Greek character "π".
	Pi string
	// Rho represents the Greek character "ρ".
	Rho string
	// Sigma represents the Greek character "σ".
	//
	// NOTE: This is the form used at the beginning or in the middle of a word, see SigmaFinal for the final form.
	Sigma string
	// SigmaFinal represents the Greek character "ς".
	//
	// NOTE: This is the final form used at the end of a word, see Sigma for the mid-word form.
	SigmaFinal string
	// Tau represents the Greek character "τ".
	Tau string
	// Upsilon represents the Greek character "υ".
	Upsilon string
	// Phi represents the Greek character "φ".
	Phi string
	// Chi represents the Greek character "χ".
	Chi string
	// Psi represents the Greek character "ψ".
	Psi string
	// Omega represents the Greek character "ω".
	Omega string
}

// Lower provides access to the lowercase Greek alphabet.
//
// See lower.Alpha, lower.Beta, lower.Gamma, lower.Delta, lower.Epsilon, lower.Zeta, lower.Eta, lower.Theta, lower.Iota, lower.Kappa, lower.Lambda, lower.Mu, lower.Nu, lower.Xi, lower.Omicron, lower.Pi, lower.Rho, lower.Sigma, lower.SigmaFinal, lower.Tau, lower.Upsilon, lower.Phi, lower.Chi, lower.Psi, and lower.Omega
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
	Alpha string
	// Beta represents the Greek character "Β".
	Beta string
	// Gamma represents the Greek character "Γ".
	Gamma string
	// Delta represents the Greek character "Δ".
	Delta string
	// Epsilon represents the Greek character "Ε".
	Epsilon string
	// Zeta represents the Greek character "Ζ".
	Zeta string
	// Eta represents the Greek character "Η".
	Eta string
	// Theta represents the Greek character "Θ".
	Theta string
	// Iota represents the Greek character "Ι".
	Iota string
	// Kappa represents the Greek character "Κ".
	Kappa string
	// Lambda represents the Greek character "Λ".
	Lambda string
	// Mu represents the Greek character "Μ".
	Mu string
	// Nu represents the Greek character "Ν".
	Nu string
	// Xi represents the Greek character "Ξ".
	Xi string
	// Omicron represents the Greek character "Ο".
	Omicron string
	// Pi represents the Greek character "Π".
	Pi string
	// Rho represents the Greek character "Ρ".
	Rho string
	// Sigma represents the Greek character "Σ".
	Sigma string
	// Tau represents the Greek character "Τ".
	Tau string
	// Upsilon represents the Greek character "Υ".
	Upsilon string
	// Phi represents the Greek character "Φ".
	Phi string
	// Chi represents the Greek character "Χ".
	Chi string
	// Psi represents the Greek character "Ψ".
	Psi string
	// Omega represents the Greek character "Ω".
	Omega string
}

// Upper provides access to the uppercase Greek alphabet.
//
// See upper.Alpha, upper.Beta, upper.Gamma, upper.Delta, upper.Epsilon, upper.Zeta, upper.Eta, upper.Theta, upper.Iota, upper.Kappa, upper.Lambda, upper.Mu, upper.Nu, upper.Xi, upper.Omicron, upper.Pi, upper.Rho, upper.Sigma, upper.Tau, upper.Upsilon, upper.Phi, upper.Chi, upper.Psi, and upper.Omega
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
