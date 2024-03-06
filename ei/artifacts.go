// Based on
// https://github.com/fanaticscripter/EggContractor/blob/3ce2cdc9ee767ecc8cbdfa4ae0ac90d248dc8694/api/computed.go#L308-L1369

package ei

import (
	"fmt"
	"strings"
)

// GameName is in all caps. Use CasedName for cased version.
func (a ArtifactSpec_Name) GameName() string {
	name := strings.ReplaceAll(a.String(), "_", " ")
	switch a {
	case ArtifactSpec_VIAL_MARTIAN_DUST:
		name = "VIAL OF MARTIAN DUST"
	case ArtifactSpec_ORNATE_GUSSET:
		name = "GUSSET"
	case ArtifactSpec_MERCURYS_LENS:
		name = "MERCURY'S LENS"
	}
	return name
}

func (a ArtifactSpec_Name) CasedName() string {
	return capitalizeArtifactName(strings.ToLower(a.GameName()))
}

func (a *ArtifactSpec) GenericBenefitString() string {
	switch *a.Name {
	case ArtifactSpec_LUNAR_TOTEM:
		return "[+^b] away earnings"
	case ArtifactSpec_NEODYMIUM_MEDALLION:
		return "[+^b] drone frequency"
	case ArtifactSpec_BEAK_OF_MIDAS:
		return "[+^b] chance of gold in gifts and drones"
	case ArtifactSpec_LIGHT_OF_EGGENDIL:
		return "Enlightenment egg value increased by [^b]"
	case ArtifactSpec_TUNGSTEN_ANKH:
		fallthrough
	case ArtifactSpec_DEMETERS_NECKLACE:
		return "[+^b] egg value"
	case ArtifactSpec_VIAL_MARTIAN_DUST:
		return "Increases max running chicken bonus by [^b]"
	case ArtifactSpec_ORNATE_GUSSET:
		return "[+^b] hab capacity"
	case ArtifactSpec_THE_CHALICE:
		return "[+^b] internal hatchery rate"
	case ArtifactSpec_BOOK_OF_BASAN:
		return "[+^b] to Egg of Prophecy bonus"
	case ArtifactSpec_PHOENIX_FEATHER:
		return "[+^b] Soul Egg collection rate"
	case ArtifactSpec_AURELIAN_BROOCH:
		return "[+^b] drone rewards"
	case ArtifactSpec_CARVED_RAINSTICK:
		return "[+^b] chance of cash in gifts and drones"
	case ArtifactSpec_PUZZLE_CUBE:
		return "[-^b] research cost"
	case ArtifactSpec_QUANTUM_METRONOME:
		return "[+^b] egg laying rate"
	case ArtifactSpec_SHIP_IN_A_BOTTLE:
		return "[+^b] co-op teammate's earnings"
	case ArtifactSpec_TACHYON_DEFLECTOR:
		return "[+^b] co-op teammate's egg laying rate"
	case ArtifactSpec_INTERSTELLAR_COMPASS:
		return "[+^b] shipping rate"
	case ArtifactSpec_DILITHIUM_MONOCLE:
		return "[+^b] boost boost"
	case ArtifactSpec_TITANIUM_ACTUATOR:
		return "[+^b] hold to hatch rate"
	case ArtifactSpec_MERCURYS_LENS:
		return "[+^b] farm valuation"
	// Stones
	case ArtifactSpec_TACHYON_STONE:
		return "[+^b] egg laying rate"
	case ArtifactSpec_DILITHIUM_STONE:
		return "[+^b] boost duration"
	case ArtifactSpec_SHELL_STONE:
		return "[+^b] egg value"
	case ArtifactSpec_LUNAR_STONE:
		return "[+^b] away earnings"
	case ArtifactSpec_SOUL_STONE:
		return "[+^b] bonus per Soul Egg"
	case ArtifactSpec_PROPHECY_STONE:
		return "[+^b] to Egg of Prophecy bonus"
	case ArtifactSpec_QUANTUM_STONE:
		return "[+^b] shipping rate"
	case ArtifactSpec_TERRA_STONE:
		return "Increases max running chicken bonus by [^b]"
	case ArtifactSpec_LIFE_STONE:
		return "[+^b] internal hatchery rate"
	case ArtifactSpec_CLARITY_STONE:
		return "[^b] effect of host artifact on enlightenment egg farms"
	default:
		return ""
	}
}

func (a *ArtifactSpec) DropEffectString() string {
	var replString string
	switch *a.Name {
	case ArtifactSpec_LUNAR_TOTEM:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "100%"
		case ArtifactSpec_LESSER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "300%"
			case ArtifactSpec_RARE:
				replString = "8x"
			}
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "20x"
			case ArtifactSpec_RARE:
				replString = "40x"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "50x"
			case ArtifactSpec_RARE:
				replString = "100x"
			case ArtifactSpec_EPIC:
				replString = "150x"
			case ArtifactSpec_LEGENDARY:
				replString = "200x"
			}
		}
	case ArtifactSpec_NEODYMIUM_MEDALLION:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "10%"
		case ArtifactSpec_LESSER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "25%"
			case ArtifactSpec_RARE:
				replString = "30%"
			}
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "50%"
			case ArtifactSpec_EPIC:
				replString = "60%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "100%"
			case ArtifactSpec_RARE:
				replString = "110%"
			case ArtifactSpec_EPIC:
				replString = "120%"
			case ArtifactSpec_LEGENDARY:
				replString = "130%"
			}
		}
	case ArtifactSpec_BEAK_OF_MIDAS:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "20%"
		case ArtifactSpec_LESSER:
			replString = "50%"
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "100%"
			case ArtifactSpec_RARE:
				replString = "200%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "500%"
			case ArtifactSpec_RARE:
				replString = "1000%"
			case ArtifactSpec_LEGENDARY:
				return "!!Gold gifts and drone rewards <guaranteed>"
			}
		}
	case ArtifactSpec_LIGHT_OF_EGGENDIL:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "50%"
		case ArtifactSpec_LESSER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "100%"
			case ArtifactSpec_RARE:
				replString = "120%"
			}
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "900%"
			case ArtifactSpec_RARE:
				replString = "1400%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "100x"
			case ArtifactSpec_EPIC:
				replString = "150x"
			case ArtifactSpec_LEGENDARY:
				replString = "250x"
			}
		}
	case ArtifactSpec_DEMETERS_NECKLACE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "10%"
		case ArtifactSpec_LESSER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "25%"
			case ArtifactSpec_RARE:
				replString = "35%"
			}
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "50%"
			case ArtifactSpec_RARE:
				replString = "60%"
			case ArtifactSpec_EPIC:
				replString = "75%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "100%"
			case ArtifactSpec_RARE:
				replString = "125%"
			case ArtifactSpec_EPIC:
				replString = "150%"
			case ArtifactSpec_LEGENDARY:
				replString = "200%"
			}
		}
	case ArtifactSpec_VIAL_MARTIAN_DUST:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "10"
		case ArtifactSpec_LESSER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "50"
			case ArtifactSpec_RARE:
				replString = "60"
			}
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "100"
			case ArtifactSpec_EPIC:
				replString = "150"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "200"
			case ArtifactSpec_RARE:
				replString = "300"
			case ArtifactSpec_LEGENDARY:
				replString = "500"
			}
		}
	case ArtifactSpec_ORNATE_GUSSET:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "5%"
		case ArtifactSpec_LESSER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "10%"
			case ArtifactSpec_EPIC:
				replString = "12%"
			}
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "15%"
			case ArtifactSpec_RARE:
				replString = "16%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "20%"
			case ArtifactSpec_EPIC:
				replString = "22%"
			case ArtifactSpec_LEGENDARY:
				replString = "25%"
			}
		}
	case ArtifactSpec_THE_CHALICE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "5%"
		case ArtifactSpec_LESSER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "10%"
			case ArtifactSpec_EPIC:
				replString = "15%"
			}
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "20%"
			case ArtifactSpec_RARE:
				replString = "23%"
			case ArtifactSpec_EPIC:
				replString = "25%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "30%"
			case ArtifactSpec_EPIC:
				replString = "35%"
			case ArtifactSpec_LEGENDARY:
				replString = "40%"
			}
		}
	case ArtifactSpec_BOOK_OF_BASAN:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "0.25%"
		case ArtifactSpec_LESSER:
			replString = "0.5%"
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "0.75%"
			case ArtifactSpec_EPIC:
				replString = "0.8%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "1%"
			case ArtifactSpec_EPIC:
				replString = "1.1%"
			case ArtifactSpec_LEGENDARY:
				replString = "1.2%"
			}
		}
	case ArtifactSpec_PHOENIX_FEATHER:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "25%"
		case ArtifactSpec_LESSER:
			replString = "100%"
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "400%"
			case ArtifactSpec_RARE:
				replString = "500%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "900%"
			case ArtifactSpec_RARE:
				replString = "1100%"
			case ArtifactSpec_LEGENDARY:
				replString = "1400%"
			}
		}
	case ArtifactSpec_TUNGSTEN_ANKH:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "10%"
		case ArtifactSpec_LESSER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "25%"
			case ArtifactSpec_RARE:
				replString = "28%"
			}
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "50%"
			case ArtifactSpec_RARE:
				replString = "75%"
			case ArtifactSpec_LEGENDARY:
				replString = "100%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "100%"
			case ArtifactSpec_RARE:
				replString = "125%"
			case ArtifactSpec_LEGENDARY:
				replString = "150%"
			}
		}
	case ArtifactSpec_AURELIAN_BROOCH:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "10%"
		case ArtifactSpec_LESSER:
			replString = "25%"
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "50%"
			case ArtifactSpec_RARE:
				replString = "60%"
			case ArtifactSpec_EPIC:
				replString = "70%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "100%"
			case ArtifactSpec_RARE:
				replString = "125%"
			case ArtifactSpec_EPIC:
				replString = "150%"
			case ArtifactSpec_LEGENDARY:
				replString = "200%"
			}
		}
	case ArtifactSpec_CARVED_RAINSTICK:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "20%"
		case ArtifactSpec_LESSER:
			replString = "50%"
		case ArtifactSpec_NORMAL:
			replString = "100%"
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "400%"
			case ArtifactSpec_EPIC:
				replString = "900%"
			case ArtifactSpec_LEGENDARY:
				return "!!Cash drone rewards and gifts <guaranteed>"
			}
		}
	case ArtifactSpec_PUZZLE_CUBE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "5%"
		case ArtifactSpec_LESSER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "10%"
			case ArtifactSpec_EPIC:
				replString = "15%"
			}
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "20%"
			case ArtifactSpec_RARE:
				replString = "22%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "50%"
			case ArtifactSpec_RARE:
				replString = "53%"
			case ArtifactSpec_EPIC:
				replString = "55%"
			case ArtifactSpec_LEGENDARY:
				replString = "60%"
			}
		}
	case ArtifactSpec_QUANTUM_METRONOME:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "5%"
		case ArtifactSpec_LESSER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "10%"
			case ArtifactSpec_RARE:
				replString = "12%"
			}
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "15%"
			case ArtifactSpec_RARE:
				replString = "17%"
			case ArtifactSpec_EPIC:
				replString = "20%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "25%"
			case ArtifactSpec_RARE:
				replString = "27%"
			case ArtifactSpec_EPIC:
				replString = "30%"
			case ArtifactSpec_LEGENDARY:
				replString = "35%"
			}
		}
	case ArtifactSpec_SHIP_IN_A_BOTTLE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "20%"
		case ArtifactSpec_LESSER:
			replString = "30%"
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "50%"
			case ArtifactSpec_RARE:
				replString = "60%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "70%"
			case ArtifactSpec_RARE:
				replString = "80%"
			case ArtifactSpec_EPIC:
				replString = "90%"
			case ArtifactSpec_LEGENDARY:
				replString = "100%"
			}
		}
	case ArtifactSpec_TACHYON_DEFLECTOR:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "5%"
		case ArtifactSpec_LESSER:
			replString = "8%"
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "12%"
			case ArtifactSpec_RARE:
				replString = "13%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "15%"
			case ArtifactSpec_RARE:
				replString = "17%"
			case ArtifactSpec_EPIC:
				replString = "19%"
			case ArtifactSpec_LEGENDARY:
				replString = "20%"
			}
		}
	case ArtifactSpec_INTERSTELLAR_COMPASS:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "5%"
		case ArtifactSpec_LESSER:
			replString = "10%"
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "20%"
			case ArtifactSpec_RARE:
				replString = "22%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "30%"
			case ArtifactSpec_RARE:
				replString = "35%"
			case ArtifactSpec_EPIC:
				replString = "40%"
			case ArtifactSpec_LEGENDARY:
				replString = "50%"
			}
		}
	case ArtifactSpec_DILITHIUM_MONOCLE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "5%"
		case ArtifactSpec_LESSER:
			replString = "10%"
		case ArtifactSpec_NORMAL:
			replString = "15%"
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "20%"
			case ArtifactSpec_EPIC:
				replString = "25%"
			case ArtifactSpec_LEGENDARY:
				replString = "30%"
			}
		}
	case ArtifactSpec_TITANIUM_ACTUATOR:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "1"
		case ArtifactSpec_LESSER:
			replString = "4"
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "6"
			case ArtifactSpec_RARE:
				replString = "7"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "10"
			case ArtifactSpec_EPIC:
				replString = "12"
			case ArtifactSpec_LEGENDARY:
				replString = "15"
			}
		}
	case ArtifactSpec_MERCURYS_LENS:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "10%"
		case ArtifactSpec_LESSER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "20%"
			case ArtifactSpec_RARE:
				replString = "22%"
			}
		case ArtifactSpec_NORMAL:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "50%"
			case ArtifactSpec_RARE:
				replString = "55%"
			}
		case ArtifactSpec_GREATER:
			switch *a.Rarity {
			case ArtifactSpec_COMMON:
				replString = "100%"
			case ArtifactSpec_RARE:
				replString = "125%"
			case ArtifactSpec_EPIC:
				replString = "150%"
			case ArtifactSpec_LEGENDARY:
				replString = "200%"
			}
		}
	// Stones
	case ArtifactSpec_QUANTUM_STONE:
		fallthrough
	case ArtifactSpec_TACHYON_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "2%"
		case ArtifactSpec_LESSER:
			replString = "4%"
		case ArtifactSpec_NORMAL:
			replString = "5%"
		}
	case ArtifactSpec_DILITHIUM_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "3%"
		case ArtifactSpec_LESSER:
			replString = "6%"
		case ArtifactSpec_NORMAL:
			replString = "8%"
		}
	case ArtifactSpec_SHELL_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "5%"
		case ArtifactSpec_LESSER:
			replString = "8%"
		case ArtifactSpec_NORMAL:
			replString = "10%"
		}
	case ArtifactSpec_LUNAR_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "20%"
		case ArtifactSpec_LESSER:
			replString = "30%"
		case ArtifactSpec_NORMAL:
			replString = "40%"
		}
	case ArtifactSpec_SOUL_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "5%"
		case ArtifactSpec_LESSER:
			replString = "10%"
		case ArtifactSpec_NORMAL:
			replString = "25%"
		}
	case ArtifactSpec_PROPHECY_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "0.05%"
		case ArtifactSpec_LESSER:
			replString = "0.1%"
		case ArtifactSpec_NORMAL:
			replString = "0.15%"
		}
	case ArtifactSpec_TERRA_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "10"
		case ArtifactSpec_LESSER:
			replString = "50"
		case ArtifactSpec_NORMAL:
			replString = "100"
		}
	case ArtifactSpec_LIFE_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "2%"
		case ArtifactSpec_LESSER:
			replString = "3%"
		case ArtifactSpec_NORMAL:
			replString = "4%"
		}
	case ArtifactSpec_CLARITY_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			replString = "25%"
		case ArtifactSpec_LESSER:
			replString = "50%"
		case ArtifactSpec_NORMAL:
			replString = "100%"
		}
	default:
		replString = "???" // should never happen
	}

	return strings.ReplaceAll(a.GenericBenefitString(), "^b", replString)
}

// GameName is in all caps. Use CasedName for cased version.
func (a *ArtifactSpec) GameName() string {
	switch *a.Name {
	// Artifacts
	case ArtifactSpec_LUNAR_TOTEM:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "BASIC LUNAR TOTEM"
		case ArtifactSpec_LESSER:
			return "LUNAR TOTEM"
		case ArtifactSpec_NORMAL:
			return "POWERFUL LUNAR TOTEM"
		case ArtifactSpec_GREATER:
			return "EGGCEPTIONAL LUNAR TOTEM"
		}
	case ArtifactSpec_NEODYMIUM_MEDALLION:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "WEAK NEODYMIUM MEDALLION"
		case ArtifactSpec_LESSER:
			return "NEODYMIUM MEDALLION"
		case ArtifactSpec_NORMAL:
			return "PRECISE NEODYMIUM MEDALLION"
		case ArtifactSpec_GREATER:
			return "EGGCEPTIONAL NEODYMIUM MEDALLION"
		}
	case ArtifactSpec_BEAK_OF_MIDAS:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "DULL BEAK OF MIDAS"
		case ArtifactSpec_LESSER:
			return "BEAK OF MIDAS"
		case ArtifactSpec_NORMAL:
			return "JEWELED BEAK OF MIDAS"
		case ArtifactSpec_GREATER:
			return "GLISTENING BEAK OF MIDAS"
		}
	case ArtifactSpec_LIGHT_OF_EGGENDIL:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "DIM LIGHT OF EGGENDIL"
		case ArtifactSpec_LESSER:
			return "SHIMMERING LIGHT OF EGGENDIL"
		case ArtifactSpec_NORMAL:
			return "GLOWING LIGHT OF EGGENDIL"
		case ArtifactSpec_GREATER:
			return "BRILLIANT LIGHT OF EGGENDIL"
		}
	case ArtifactSpec_DEMETERS_NECKLACE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "SIMPLE DEMETERS NECKLACE"
		case ArtifactSpec_LESSER:
			return "JEWELED DEMETERS NECKLACE"
		case ArtifactSpec_NORMAL:
			return "PRISTINE DEMETERS NECKLACE"
		case ArtifactSpec_GREATER:
			return "BEGGSPOKE DEMETERS NECKLACE"
		}
	case ArtifactSpec_VIAL_MARTIAN_DUST:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "TINY VIAL OF MARTIAN DUST"
		case ArtifactSpec_LESSER:
			return "VIAL OF MARTIAN DUST"
		case ArtifactSpec_NORMAL:
			return "HERMETIC VIAL OF MARTIAN DUST"
		case ArtifactSpec_GREATER:
			return "PRIME VIAL OF MARTIAN DUST"
		}
	case ArtifactSpec_ORNATE_GUSSET:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "PLAIN GUSSET"
		case ArtifactSpec_LESSER:
			return "ORNATE GUSSET"
		case ArtifactSpec_NORMAL:
			return "DISTEGGUISHED GUSSET"
		case ArtifactSpec_GREATER:
			return "JEWELED GUSSET"
		}
	case ArtifactSpec_THE_CHALICE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "PLAIN CHALICE"
		case ArtifactSpec_LESSER:
			return "POLISHED CHALICE"
		case ArtifactSpec_NORMAL:
			return "JEWELED CHALICE"
		case ArtifactSpec_GREATER:
			return "EGGCEPTIONAL CHALICE"
		}
	case ArtifactSpec_BOOK_OF_BASAN:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "BOOK OF BASAN"
		case ArtifactSpec_LESSER:
			return "COLLECTORS BOOK OF BASAN"
		case ArtifactSpec_NORMAL:
			return "FORTIFIED BOOK OF BASAN"
		case ArtifactSpec_GREATER:
			return "GILDED BOOK OF BASAN"
		}
	case ArtifactSpec_PHOENIX_FEATHER:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "TATTERED PHOENIX FEATHER"
		case ArtifactSpec_LESSER:
			return "PHOENIX FEATHER"
		case ArtifactSpec_NORMAL:
			return "BRILLIANT PHOENIX FEATHER"
		case ArtifactSpec_GREATER:
			return "BLAZING PHOENIX FEATHER"
		}
	case ArtifactSpec_TUNGSTEN_ANKH:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "CRUDE TUNGSTEN ANKH"
		case ArtifactSpec_LESSER:
			return "TUNGSTEN ANKH"
		case ArtifactSpec_NORMAL:
			return "POLISHED TUNGSTEN ANKH"
		case ArtifactSpec_GREATER:
			return "BRILLIANT TUNGSTEN ANKH"
		}
	case ArtifactSpec_AURELIAN_BROOCH:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "PLAIN AURELIAN BROOCH"
		case ArtifactSpec_LESSER:
			return "AURELIAN BROOCH"
		case ArtifactSpec_NORMAL:
			return "JEWELED AURELIAN BROOCH"
		case ArtifactSpec_GREATER:
			return "EGGCEPTIONAL AURELIAN BROOCH"
		}
	case ArtifactSpec_CARVED_RAINSTICK:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "SIMPLE CARVED RAINSTICK"
		case ArtifactSpec_LESSER:
			return "CARVED RAINSTICK"
		case ArtifactSpec_NORMAL:
			return "ORNATE CARVED RAINSTICK"
		case ArtifactSpec_GREATER:
			return "MEGGNIFICENT CARVED RAINSTICK"
		}
	case ArtifactSpec_PUZZLE_CUBE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "ANCIENT PUZZLE CUBE"
		case ArtifactSpec_LESSER:
			return "PUZZLE CUBE"
		case ArtifactSpec_NORMAL:
			return "MYSTICAL PUZZLE CUBE"
		case ArtifactSpec_GREATER:
			return "UNSOLVABLE PUZZLE CUBE"
		}
	case ArtifactSpec_QUANTUM_METRONOME:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "MISALIGNED QUANTUM METRONOME"
		case ArtifactSpec_LESSER:
			return "ADEQUATE QUANTUM METRONOME"
		case ArtifactSpec_NORMAL:
			return "PERFECT QUANTUM METRONOME"
		case ArtifactSpec_GREATER:
			return "REGGFERENCE QUANTUM METRONOME"
		}
	case ArtifactSpec_SHIP_IN_A_BOTTLE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "SHIP IN A BOTTLE"
		case ArtifactSpec_LESSER:
			return "DETAILED SHIP IN A BOTTLE"
		case ArtifactSpec_NORMAL:
			return "COMPLEX SHIP IN A BOTTLE"
		case ArtifactSpec_GREATER:
			return "EGGQUISITE SHIP IN A BOTTLE"
		}
	case ArtifactSpec_TACHYON_DEFLECTOR:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "WEAK TACHYON DEFLECTOR"
		case ArtifactSpec_LESSER:
			return "TACHYON DEFLECTOR"
		case ArtifactSpec_NORMAL:
			return "ROBUST TACHYON DEFLECTOR"
		case ArtifactSpec_GREATER:
			return "EGGCEPTIONAL TACHYON DEFLECTOR"
		}
	case ArtifactSpec_INTERSTELLAR_COMPASS:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "MISCALIBRATED INTERSTELLAR COMPASS"
		case ArtifactSpec_LESSER:
			return "INTERSTELLAR COMPASS"
		case ArtifactSpec_NORMAL:
			return "PRECISE INTERSTELLAR COMPASS"
		case ArtifactSpec_GREATER:
			return "CLAIRVOYANT INTERSTELLAR COMPASS"
		}
	case ArtifactSpec_DILITHIUM_MONOCLE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "DILITHIUM MONOCLE"
		case ArtifactSpec_LESSER:
			return "PRECISE DILITHIUM MONOCLE"
		case ArtifactSpec_NORMAL:
			return "EGGSACTING DILITHIUM MONOCLE"
		case ArtifactSpec_GREATER:
			return "FLAWLESS DILITHIUM MONOCLE"
		}
	case ArtifactSpec_TITANIUM_ACTUATOR:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "INCONSISTENT TITANIUM ACTUATOR"
		case ArtifactSpec_LESSER:
			return "TITANIUM ACTUATOR"
		case ArtifactSpec_NORMAL:
			return "PRECISE TITANIUM ACTUATOR"
		case ArtifactSpec_GREATER:
			return "REGGFERENCE TITANIUM ACTUATOR"
		}
	case ArtifactSpec_MERCURYS_LENS:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "MISALIGNED MERCURY'S LENS"
		case ArtifactSpec_LESSER:
			return "MERCURY'S LENS"
		case ArtifactSpec_NORMAL:
			return "PRECISE MERCURY'S LENS"
		case ArtifactSpec_GREATER:
			return "MEGGNIFICENT MERCURY'S LENS"
		}
	// Stones
	case ArtifactSpec_TACHYON_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "TACHYON STONE"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE TACHYON STONE"
		case ArtifactSpec_NORMAL:
			return "BRILLIANT TACHYON STONE"
		}
	case ArtifactSpec_DILITHIUM_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "DILITHIUM STONE"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE DILITHIUM STONE"
		case ArtifactSpec_NORMAL:
			return "BRILLIANT DILITHIUM STONE"
		}
	case ArtifactSpec_SHELL_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "SHELL STONE"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE SHELL STONE"
		case ArtifactSpec_NORMAL:
			return "FLAWLESS SHELL STONE"
		}
	case ArtifactSpec_LUNAR_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "LUNAR STONE"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE LUNAR STONE"
		case ArtifactSpec_NORMAL:
			return "MEGGNIFICENT LUNAR STONE"
		}
	case ArtifactSpec_SOUL_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "SOUL STONE"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE SOUL STONE"
		case ArtifactSpec_NORMAL:
			return "RADIANT SOUL STONE"
		}
	case ArtifactSpec_PROPHECY_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "PROPHECY STONE"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE PROPHECY STONE"
		case ArtifactSpec_NORMAL:
			return "RADIANT PROPHECY STONE"
		}
	case ArtifactSpec_QUANTUM_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "QUANTUM STONE"
		case ArtifactSpec_LESSER:
			return "PHASED QUANTUM STONE"
		case ArtifactSpec_NORMAL:
			return "MEGGNIFICENT QUANTUM STONE"
		}
	case ArtifactSpec_TERRA_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "TERRA STONE"
		case ArtifactSpec_LESSER:
			return "RICH TERRA STONE"
		case ArtifactSpec_NORMAL:
			return "EGGCEPTIONAL TERRA STONE"
		}
	case ArtifactSpec_LIFE_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "LIFE STONE"
		case ArtifactSpec_LESSER:
			return "GOOD LIFE STONE"
		case ArtifactSpec_NORMAL:
			return "EGGCEPTIONAL LIFE STONE"
		}
	case ArtifactSpec_CLARITY_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "CLARITY STONE"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE CLARITY STONE"
		case ArtifactSpec_NORMAL:
			return "EGGCEPTIONAL CLARITY STONE"
		}
	// Stone fragments
	case ArtifactSpec_TACHYON_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_DILITHIUM_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_SHELL_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_LUNAR_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_SOUL_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_PROPHECY_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_QUANTUM_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_TERRA_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_LIFE_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_CLARITY_STONE_FRAGMENT:
		return strings.ReplaceAll(a.Name.String(), "_", " ")
	// Ingredients
	case ArtifactSpec_GOLD_METEORITE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "TINY GOLD METEORITE"
		case ArtifactSpec_LESSER:
			return "ENRICHED GOLD METEORITE"
		case ArtifactSpec_NORMAL:
			return "SOLID GOLD METEORITE"
		}
	case ArtifactSpec_TAU_CETI_GEODE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "TAU CETI GEODE PIECE"
		case ArtifactSpec_LESSER:
			return "GLIMMERING TAU CETI GEODE"
		case ArtifactSpec_NORMAL:
			return "RADIANT TAU CETI GEODE"
		}
	case ArtifactSpec_SOLAR_TITANIUM:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "SOLAR TITANIUM ORE"
		case ArtifactSpec_LESSER:
			return "SOLAR TITANIUM BAR"
		case ArtifactSpec_NORMAL:
			return "SOLAR TITANIUM GEOGON"
		}
	// Unconfirmed ingredients
	case ArtifactSpec_EXTRATERRESTRIAL_ALUMINUM:
		fallthrough
	case ArtifactSpec_ANCIENT_TUNGSTEN:
		fallthrough
	case ArtifactSpec_SPACE_ROCKS:
		fallthrough
	case ArtifactSpec_ALIEN_WOOD:
		fallthrough
	case ArtifactSpec_CENTAURIAN_STEEL:
		fallthrough
	case ArtifactSpec_ERIDANI_FEATHER:
		fallthrough
	case ArtifactSpec_DRONE_PARTS:
		fallthrough
	case ArtifactSpec_CELESTIAL_BRONZE:
		fallthrough
	case ArtifactSpec_LALANDE_HIDE:
		return "? " + a.Name.String()
	}
	return a.Level.String() + " " + a.Name.GameName()
}

func (a *ArtifactSpec) CasedName() string {
	return capitalizeArtifactName(strings.ToLower(a.GameName()))
}

func (a *ArtifactSpec) CasedSmallName() string {
	return capitalizeArtifactName((strings.ToLower(strings.ReplaceAll(a.Name.String(), "_", " "))))
}

func capitalizeArtifactName(n string) string {
	n = strings.ToUpper(n[:1]) + n[1:]
	// Captalize proper nouns.
	for s, repl := range map[string]string{
		"demeters": "Demeters",
		"midas":    "Midas",
		"eggendil": "Eggendil",
		"martian":  "Martian",
		"basan":    "Basan",
		"aurelian": "Aurelian",
		"mercury":  "Mercury",
		"tau ceti": "Tau Ceti",
		"Tau ceti": "Tau Ceti",
	} {
		n = strings.ReplaceAll(n, s, repl)
	}
	return n
}

func (a *ArtifactSpec) Type() ArtifactSpec_Type {
	return a.Name.ArtifactType()
}

func (a ArtifactSpec_Name) InventoryVisualizerOrder() int {
	switch a {
	case ArtifactSpec_LUNAR_TOTEM:
		return 15
	case ArtifactSpec_NEODYMIUM_MEDALLION:
		return 21
	case ArtifactSpec_BEAK_OF_MIDAS:
		return 23
	case ArtifactSpec_LIGHT_OF_EGGENDIL:
		return 34
	case ArtifactSpec_DEMETERS_NECKLACE:
		return 16
	case ArtifactSpec_VIAL_MARTIAN_DUST:
		return 17
	case ArtifactSpec_ORNATE_GUSSET:
		return 20
	case ArtifactSpec_THE_CHALICE:
		return 26
	case ArtifactSpec_BOOK_OF_BASAN:
		return 33
	case ArtifactSpec_PHOENIX_FEATHER:
		return 27
	case ArtifactSpec_TUNGSTEN_ANKH:
		return 19
	case ArtifactSpec_AURELIAN_BROOCH:
		return 18
	case ArtifactSpec_CARVED_RAINSTICK:
		return 24
	case ArtifactSpec_PUZZLE_CUBE:
		return 14
	case ArtifactSpec_QUANTUM_METRONOME:
		return 28
	case ArtifactSpec_SHIP_IN_A_BOTTLE:
		return 31
	case ArtifactSpec_TACHYON_DEFLECTOR:
		return 32
	case ArtifactSpec_INTERSTELLAR_COMPASS:
		return 25
	case ArtifactSpec_DILITHIUM_MONOCLE:
		return 29
	case ArtifactSpec_TITANIUM_ACTUATOR:
		return 30
	case ArtifactSpec_MERCURYS_LENS:
		return 22
	case ArtifactSpec_TACHYON_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_TACHYON_STONE:
		return 6
	case ArtifactSpec_DILITHIUM_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_DILITHIUM_STONE:
		return 11
	case ArtifactSpec_SHELL_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_SHELL_STONE:
		return 4
	case ArtifactSpec_LUNAR_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_LUNAR_STONE:
		return 5
	case ArtifactSpec_SOUL_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_SOUL_STONE:
		return 8
	case ArtifactSpec_PROPHECY_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_PROPHECY_STONE:
		return 13
	case ArtifactSpec_QUANTUM_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_QUANTUM_STONE:
		return 9
	case ArtifactSpec_TERRA_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_TERRA_STONE:
		return 7
	case ArtifactSpec_LIFE_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_LIFE_STONE:
		return 10
	case ArtifactSpec_CLARITY_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_CLARITY_STONE:
		return 12
	case ArtifactSpec_GOLD_METEORITE:
		return 1
	case ArtifactSpec_TAU_CETI_GEODE:
		return 2
	case ArtifactSpec_SOLAR_TITANIUM:
		return 3
	default:
		return 0
	}
}

func (a ArtifactSpec_Name) ArtifactType() ArtifactSpec_Type {
	switch a {
	// Artifacts
	case ArtifactSpec_LUNAR_TOTEM:
		fallthrough
	case ArtifactSpec_NEODYMIUM_MEDALLION:
		fallthrough
	case ArtifactSpec_BEAK_OF_MIDAS:
		fallthrough
	case ArtifactSpec_LIGHT_OF_EGGENDIL:
		fallthrough
	case ArtifactSpec_DEMETERS_NECKLACE:
		fallthrough
	case ArtifactSpec_VIAL_MARTIAN_DUST:
		fallthrough
	case ArtifactSpec_ORNATE_GUSSET:
		fallthrough
	case ArtifactSpec_THE_CHALICE:
		fallthrough
	case ArtifactSpec_BOOK_OF_BASAN:
		fallthrough
	case ArtifactSpec_PHOENIX_FEATHER:
		fallthrough
	case ArtifactSpec_TUNGSTEN_ANKH:
		fallthrough
	case ArtifactSpec_AURELIAN_BROOCH:
		fallthrough
	case ArtifactSpec_CARVED_RAINSTICK:
		fallthrough
	case ArtifactSpec_PUZZLE_CUBE:
		fallthrough
	case ArtifactSpec_QUANTUM_METRONOME:
		fallthrough
	case ArtifactSpec_SHIP_IN_A_BOTTLE:
		fallthrough
	case ArtifactSpec_TACHYON_DEFLECTOR:
		fallthrough
	case ArtifactSpec_INTERSTELLAR_COMPASS:
		fallthrough
	case ArtifactSpec_DILITHIUM_MONOCLE:
		fallthrough
	case ArtifactSpec_TITANIUM_ACTUATOR:
		fallthrough
	case ArtifactSpec_MERCURYS_LENS:
		return ArtifactSpec_ARTIFACT
	// Stones
	case ArtifactSpec_TACHYON_STONE:
		fallthrough
	case ArtifactSpec_DILITHIUM_STONE:
		fallthrough
	case ArtifactSpec_SHELL_STONE:
		fallthrough
	case ArtifactSpec_LUNAR_STONE:
		fallthrough
	case ArtifactSpec_SOUL_STONE:
		fallthrough
	case ArtifactSpec_PROPHECY_STONE:
		fallthrough
	case ArtifactSpec_QUANTUM_STONE:
		fallthrough
	case ArtifactSpec_TERRA_STONE:
		fallthrough
	case ArtifactSpec_LIFE_STONE:
		fallthrough
	case ArtifactSpec_CLARITY_STONE:
		return ArtifactSpec_STONE
	// Stone fragments
	case ArtifactSpec_TACHYON_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_DILITHIUM_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_SHELL_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_LUNAR_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_SOUL_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_PROPHECY_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_QUANTUM_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_TERRA_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_LIFE_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_CLARITY_STONE_FRAGMENT:
		return ArtifactSpec_STONE_INGREDIENT
	// Ingredients
	case ArtifactSpec_GOLD_METEORITE:
		fallthrough
	case ArtifactSpec_TAU_CETI_GEODE:
		fallthrough
	case ArtifactSpec_SOLAR_TITANIUM:
		fallthrough
	// Unconfirmed ingredients
	case ArtifactSpec_EXTRATERRESTRIAL_ALUMINUM:
		fallthrough
	case ArtifactSpec_ANCIENT_TUNGSTEN:
		fallthrough
	case ArtifactSpec_SPACE_ROCKS:
		fallthrough
	case ArtifactSpec_ALIEN_WOOD:
		fallthrough
	case ArtifactSpec_CENTAURIAN_STEEL:
		fallthrough
	case ArtifactSpec_ERIDANI_FEATHER:
		fallthrough
	case ArtifactSpec_DRONE_PARTS:
		fallthrough
	case ArtifactSpec_CELESTIAL_BRONZE:
		fallthrough
	case ArtifactSpec_LALANDE_HIDE:
		return ArtifactSpec_INGREDIENT
	}
	return ArtifactSpec_ARTIFACT
}

// Family returns the family the artifact belongs to, which is the corresponding
// stone for stone fragments.
func (a *ArtifactSpec) Family() ArtifactSpec_Name {
	return a.Name.Family()
}

// Family returns the family of the artifact, which is simply itself other than
// when it is a stone fragment, in which case the corresponding stone is
// returned.
func (a ArtifactSpec_Name) Family() ArtifactSpec_Name {
	if a.ArtifactType() == ArtifactSpec_STONE_INGREDIENT {
		return a.CorrespondingStone()
	}
	return a
}

// CorrespondingStone returns the corresponding stone for a stone fragment.
// Result is undefined for non-stone fragments.
func (a ArtifactSpec_Name) CorrespondingStone() ArtifactSpec_Name {
	switch a {
	case ArtifactSpec_TACHYON_STONE_FRAGMENT:
		return ArtifactSpec_TACHYON_STONE
	case ArtifactSpec_DILITHIUM_STONE_FRAGMENT:
		return ArtifactSpec_DILITHIUM_STONE
	case ArtifactSpec_SHELL_STONE_FRAGMENT:
		return ArtifactSpec_SHELL_STONE
	case ArtifactSpec_LUNAR_STONE_FRAGMENT:
		return ArtifactSpec_LUNAR_STONE
	case ArtifactSpec_SOUL_STONE_FRAGMENT:
		return ArtifactSpec_SOUL_STONE
	case ArtifactSpec_PROPHECY_STONE_FRAGMENT:
		return ArtifactSpec_PROPHECY_STONE
	case ArtifactSpec_QUANTUM_STONE_FRAGMENT:
		return ArtifactSpec_QUANTUM_STONE
	case ArtifactSpec_TERRA_STONE_FRAGMENT:
		return ArtifactSpec_TERRA_STONE
	case ArtifactSpec_LIFE_STONE_FRAGMENT:
		return ArtifactSpec_LIFE_STONE
	case ArtifactSpec_CLARITY_STONE_FRAGMENT:
		return ArtifactSpec_CLARITY_STONE
	}
	return ArtifactSpec_UNKNOWN
}

// CorrespondingFragment returns the corresponding stone fragment for a stone.
// Result is undefined for non-stones.
func (a ArtifactSpec_Name) CorrespondingFragment() ArtifactSpec_Name {
	switch a {
	case ArtifactSpec_TACHYON_STONE:
		return ArtifactSpec_TACHYON_STONE_FRAGMENT
	case ArtifactSpec_DILITHIUM_STONE:
		return ArtifactSpec_DILITHIUM_STONE_FRAGMENT
	case ArtifactSpec_SHELL_STONE:
		return ArtifactSpec_SHELL_STONE_FRAGMENT
	case ArtifactSpec_LUNAR_STONE:
		return ArtifactSpec_LUNAR_STONE_FRAGMENT
	case ArtifactSpec_SOUL_STONE:
		return ArtifactSpec_SOUL_STONE_FRAGMENT
	case ArtifactSpec_PROPHECY_STONE:
		return ArtifactSpec_PROPHECY_STONE_FRAGMENT
	case ArtifactSpec_QUANTUM_STONE:
		return ArtifactSpec_QUANTUM_STONE_FRAGMENT
	case ArtifactSpec_TERRA_STONE:
		return ArtifactSpec_TERRA_STONE_FRAGMENT
	case ArtifactSpec_LIFE_STONE:
		return ArtifactSpec_LIFE_STONE_FRAGMENT
	case ArtifactSpec_CLARITY_STONE:
		return ArtifactSpec_CLARITY_STONE_FRAGMENT
	}
	return ArtifactSpec_UNKNOWN
}

func (a *ArtifactSpec) TierNumber() int {
	switch a.Type() {
	case ArtifactSpec_ARTIFACT:
		// 0, 1, 2, 3 => T1, T2, T3, T4
		return int(*a.Level) + 1
	case ArtifactSpec_STONE:
		// 0, 1, 2 => T2, T3, T4 (fragment as T1)
		return int(*a.Level) + 2
	case ArtifactSpec_STONE_INGREDIENT:
		return 1
	case ArtifactSpec_INGREDIENT:
		// 0, 1, 2 => T1, T2, T3
		return int(*a.Level) + 1
	}
	return 1
}

func (a *ArtifactSpec) TierName() string {
	switch *a.Name {
	// Artifacts
	case ArtifactSpec_LUNAR_TOTEM:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "BASIC"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "POWERFUL"
		case ArtifactSpec_GREATER:
			return "EGGCEPTIONAL"
		}
	case ArtifactSpec_NEODYMIUM_MEDALLION:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "WEAK"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "PRECISE"
		case ArtifactSpec_GREATER:
			return "EGGCEPTIONAL"
		}
	case ArtifactSpec_BEAK_OF_MIDAS:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "DULL"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "JEWELED"
		case ArtifactSpec_GREATER:
			return "GLISTENING"
		}
	case ArtifactSpec_LIGHT_OF_EGGENDIL:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "DIM"
		case ArtifactSpec_LESSER:
			return "SHIMMERING"
		case ArtifactSpec_NORMAL:
			return "GLOWING"
		case ArtifactSpec_GREATER:
			return "BRILLIANT"
		}
	case ArtifactSpec_DEMETERS_NECKLACE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "SIMPLE"
		case ArtifactSpec_LESSER:
			return "JEWELED"
		case ArtifactSpec_NORMAL:
			return "PRISTINE"
		case ArtifactSpec_GREATER:
			return "BEGGSPOKE"
		}
	case ArtifactSpec_VIAL_MARTIAN_DUST:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "TINY"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "HERMETIC"
		case ArtifactSpec_GREATER:
			return "PRIME"
		}
	case ArtifactSpec_ORNATE_GUSSET:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "PLAIN"
		case ArtifactSpec_LESSER:
			return "ORNATE"
		case ArtifactSpec_NORMAL:
			return "DISTEGGUISHED"
		case ArtifactSpec_GREATER:
			return "JEWELED"
		}
	case ArtifactSpec_THE_CHALICE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "PLAIN"
		case ArtifactSpec_LESSER:
			return "POLISHED"
		case ArtifactSpec_NORMAL:
			return "JEWELED"
		case ArtifactSpec_GREATER:
			return "EGGCEPTIONAL"
		}
	case ArtifactSpec_BOOK_OF_BASAN:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "REGULAR"
		case ArtifactSpec_LESSER:
			return "COLLECTORS"
		case ArtifactSpec_NORMAL:
			return "FORTIFIED"
		case ArtifactSpec_GREATER:
			return "GILDED"
		}
	case ArtifactSpec_PHOENIX_FEATHER:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "TATTERED"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "BRILLIANT"
		case ArtifactSpec_GREATER:
			return "BLAZING"
		}
	case ArtifactSpec_TUNGSTEN_ANKH:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "CRUDE"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "POLISHED"
		case ArtifactSpec_GREATER:
			return "BRILLIANT"
		}
	case ArtifactSpec_AURELIAN_BROOCH:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "PLAIN"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "JEWELED"
		case ArtifactSpec_GREATER:
			return "EGGCEPTIONAL"
		}
	case ArtifactSpec_CARVED_RAINSTICK:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "SIMPLE"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "ORNATE"
		case ArtifactSpec_GREATER:
			return "MEGGNIFICENT"
		}
	case ArtifactSpec_PUZZLE_CUBE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "ANCIENT"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "MYSTICAL"
		case ArtifactSpec_GREATER:
			return "UNSOLVABLE"
		}
	case ArtifactSpec_QUANTUM_METRONOME:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "MISALIGNED"
		case ArtifactSpec_LESSER:
			return "ADEQUATE"
		case ArtifactSpec_NORMAL:
			return "PERFECT"
		case ArtifactSpec_GREATER:
			return "REGGFERENCE"
		}
	case ArtifactSpec_SHIP_IN_A_BOTTLE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "REGULAR"
		case ArtifactSpec_LESSER:
			return "DETAILED"
		case ArtifactSpec_NORMAL:
			return "COMPLEX"
		case ArtifactSpec_GREATER:
			return "EGGQUISITE"
		}
	case ArtifactSpec_TACHYON_DEFLECTOR:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "WEAK"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "ROBUST"
		case ArtifactSpec_GREATER:
			return "EGGCEPTIONAL"
		}
	case ArtifactSpec_INTERSTELLAR_COMPASS:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "MISCALIBRATED"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "PRECISE"
		case ArtifactSpec_GREATER:
			return "CLAIRVOYANT"
		}
	case ArtifactSpec_DILITHIUM_MONOCLE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "REGULAR"
		case ArtifactSpec_LESSER:
			return "PRECISE"
		case ArtifactSpec_NORMAL:
			return "EGGSACTING"
		case ArtifactSpec_GREATER:
			return "FLAWLESS"
		}
	case ArtifactSpec_TITANIUM_ACTUATOR:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "INCONSISTENT"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "PRECISE"
		case ArtifactSpec_GREATER:
			return "REGGFERENCE"
		}
	case ArtifactSpec_MERCURYS_LENS:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "MISALIGNED"
		case ArtifactSpec_LESSER:
			return "REGULAR"
		case ArtifactSpec_NORMAL:
			return "PRECISE"
		case ArtifactSpec_GREATER:
			return "MEGGNIFICENT"
		}
	// Stones
	case ArtifactSpec_DILITHIUM_STONE:
	case ArtifactSpec_TACHYON_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "REGULAR"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE"
		case ArtifactSpec_NORMAL:
			return "BRILLIANT"
		}
	case ArtifactSpec_SHELL_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "REGULAR"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE"
		case ArtifactSpec_NORMAL:
			return "FLAWLESS"
		}
	case ArtifactSpec_LUNAR_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "REGULAR"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE"
		case ArtifactSpec_NORMAL:
			return "MEGGNIFICENT"
		}
	case ArtifactSpec_PROPHECY_STONE:
	case ArtifactSpec_SOUL_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "REGULAR"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE"
		case ArtifactSpec_NORMAL:
			return "RADIANT"
		}
	case ArtifactSpec_QUANTUM_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "REGULAR"
		case ArtifactSpec_LESSER:
			return "PHASED"
		case ArtifactSpec_NORMAL:
			return "MEGGNIFICENT"
		}
	case ArtifactSpec_TERRA_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "REGULAR"
		case ArtifactSpec_LESSER:
			return "RICH"
		case ArtifactSpec_NORMAL:
			return "EGGCEPTIONAL"
		}
	case ArtifactSpec_LIFE_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "REGULAR"
		case ArtifactSpec_LESSER:
			return "GOOD"
		case ArtifactSpec_NORMAL:
			return "EGGCEPTIONAL"
		}
	case ArtifactSpec_CLARITY_STONE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "REGULAR"
		case ArtifactSpec_LESSER:
			return "EGGSQUISITE"
		case ArtifactSpec_NORMAL:
			return "EGGCEPTIONAL"
		}
	// Stone fragments
	case ArtifactSpec_TACHYON_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_DILITHIUM_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_SHELL_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_LUNAR_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_SOUL_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_PROPHECY_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_QUANTUM_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_TERRA_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_LIFE_STONE_FRAGMENT:
		fallthrough
	case ArtifactSpec_CLARITY_STONE_FRAGMENT:
		return "FRAGMENT"
	// Ingredients
	case ArtifactSpec_GOLD_METEORITE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "TINY"
		case ArtifactSpec_LESSER:
			return "ENRICHED"
		case ArtifactSpec_NORMAL:
			return "SOLID"
		}
	case ArtifactSpec_TAU_CETI_GEODE:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "TAU"
		case ArtifactSpec_LESSER:
			return "GLIMMERING"
		case ArtifactSpec_NORMAL:
			return "RADIANT"
		}
	case ArtifactSpec_SOLAR_TITANIUM:
		switch *a.Level {
		case ArtifactSpec_INFERIOR:
			return "ORE"
		case ArtifactSpec_LESSER:
			return "BAR"
		case ArtifactSpec_NORMAL:
			return "GEOGON"
		}
	// Unconfirmed ingredients
	case ArtifactSpec_EXTRATERRESTRIAL_ALUMINUM:
		fallthrough
	case ArtifactSpec_ANCIENT_TUNGSTEN:
		fallthrough
	case ArtifactSpec_SPACE_ROCKS:
		fallthrough
	case ArtifactSpec_ALIEN_WOOD:
		fallthrough
	case ArtifactSpec_CENTAURIAN_STEEL:
		fallthrough
	case ArtifactSpec_ERIDANI_FEATHER:
		fallthrough
	case ArtifactSpec_DRONE_PARTS:
		fallthrough
	case ArtifactSpec_CELESTIAL_BRONZE:
		fallthrough
	case ArtifactSpec_LALANDE_HIDE:
		return "?"
	}
	return "?"
}

func (a *ArtifactSpec) CasedTierName() string {
	return strings.Title(strings.ToLower(a.TierName()))
}

func (a *ArtifactSpec) Display() string {
	s := fmt.Sprintf("%s (T%d)", a.CasedName(), a.TierNumber())
	if *a.Rarity > 0 {
		s += fmt.Sprintf(", %s", a.Rarity.Display())
	}
	return s
}

func (r ArtifactSpec_Rarity) Display() string {
	switch r {
	case ArtifactSpec_COMMON:
		return "Common"
	case ArtifactSpec_RARE:
		return "Rare"
	case ArtifactSpec_EPIC:
		return "Epic"
	case ArtifactSpec_LEGENDARY:
		return "Legendary"
	}
	return "Unknown"
}

func (t ArtifactSpec_Type) Display() string {
	switch t {
	case ArtifactSpec_ARTIFACT:
		return "Artifact"
	case ArtifactSpec_STONE:
		return "Stone"
	case ArtifactSpec_INGREDIENT:
		return "Ingredient"
	case ArtifactSpec_STONE_INGREDIENT:
		return "Stone ingredient"
	}
	return "Unknown"
}
