package iso2955

import "fmt"

// Unit represents an ISO 2955 unit.
type Unit struct {
	Name                string
	InternationalSymbol string
	Representation      Representation
}

func (u Unit) String() string {
	return fmt.Sprintf("%s (%s)", u.Name, u.InternationalSymbol)
}

// Representation holds the representation of the unit in Form I (double case)
// and Form II (single case lower & single case upper).
type Representation struct {
	FormI  string
	FormII FormII
}

// FormII holds single case lower and single case upper representations of a unit.
type FormII struct {
	SingleCaseLower string
	SingleCaseUpper string
}

// Base SI Units
var (
	Metre = Unit{
		Name:                "metre",
		InternationalSymbol: "m",
		Representation: Representation{
			FormI: "m",
			FormII: FormII{
				SingleCaseLower: "m",
				SingleCaseUpper: "M",
			},
		},
	}
	Kilogram = Unit{
		Name:                "kilogram",
		InternationalSymbol: "kg",
		Representation: Representation{
			FormI: "kg",
			FormII: FormII{
				SingleCaseLower: "kg",
				SingleCaseUpper: "KG",
			},
		},
	}
	Second = Unit{
		Name:                "second",
		InternationalSymbol: "s",
		Representation: Representation{
			FormI: "s",
			FormII: FormII{
				SingleCaseLower: "s",
				SingleCaseUpper: "S",
			},
		},
	}
	Ampere = Unit{
		Name:                "ampere",
		InternationalSymbol: "A",
		Representation: Representation{
			FormI: "A",
			FormII: FormII{
				SingleCaseLower: "a",
				SingleCaseUpper: "A",
			},
		},
	}
	Kelvin = Unit{
		Name:                "kelvin",
		InternationalSymbol: "K",
		Representation: Representation{
			FormI: "K",
			FormII: FormII{
				SingleCaseLower: "k",
				SingleCaseUpper: "K",
			},
		},
	}
	Mole = Unit{
		Name:                "mole",
		InternationalSymbol: "mol",
		Representation: Representation{
			FormI: "mol",
			FormII: FormII{
				SingleCaseLower: "mol",
				SingleCaseUpper: "MOL",
			},
		},
	}
	Candela = Unit{
		Name:                "candela",
		InternationalSymbol: "cd",
		Representation: Representation{
			FormI: "cd",
			FormII: FormII{
				SingleCaseLower: "cd",
				SingleCaseUpper: "CD",
			},
		},
	}
)

// Supplementary SI Units
var (
	Radian = Unit{
		Name:                "radian",
		InternationalSymbol: "rad",
		Representation: Representation{
			FormI: "rad",
			FormII: FormII{
				SingleCaseLower: "rad",
				SingleCaseUpper: "RAD",
			},
		},
	}
	Steradian = Unit{
		Name:                "steradian",
		InternationalSymbol: "sr",
		Representation: Representation{
			FormI: "sr",
			FormII: FormII{
				SingleCaseLower: "sr",
				SingleCaseUpper: "SR",
			},
		},
	}
)

// Derived SI units with special names
var (
	Hertz = Unit{
		Name:                "hertz",
		InternationalSymbol: "Hz",
		Representation: Representation{
			FormI: "Hz",
			FormII: FormII{
				SingleCaseLower: "hz",
				SingleCaseUpper: "HZ",
			},
		},
	}
	Newton = Unit{
		Name:                "newton",
		InternationalSymbol: "N",
		Representation: Representation{
			FormI: "N",
			FormII: FormII{
				SingleCaseLower: "n",
				SingleCaseUpper: "N",
			},
		},
	}
	Pascal = Unit{
		Name:                "pascal",
		InternationalSymbol: "Pa",
		Representation: Representation{
			FormI: "Pa",
			FormII: FormII{
				SingleCaseLower: "pal",
				SingleCaseUpper: "PAL",
			},
		},
	}
	Joule = Unit{
		Name:                "joule",
		InternationalSymbol: "J",
		Representation: Representation{
			FormI: "J",
			FormII: FormII{
				SingleCaseLower: "j",
				SingleCaseUpper: "J",
			},
		},
	}
	Watt = Unit{
		Name:                "watt",
		InternationalSymbol: "W",
		Representation: Representation{
			FormI: "W",
			FormII: FormII{
				SingleCaseLower: "w",
				SingleCaseUpper: "W",
			},
		},
	}
	Coulomb = Unit{
		Name:                "coulomb",
		InternationalSymbol: "C",
		Representation: Representation{
			FormI: "C",
			FormII: FormII{
				SingleCaseLower: "c",
				SingleCaseUpper: "C",
			},
		},
	}
	Volt = Unit{
		Name:                "volt",
		InternationalSymbol: "V",
		Representation: Representation{
			FormI: "V",
			FormII: FormII{
				SingleCaseLower: "v",
				SingleCaseUpper: "V",
			},
		},
	}
	Farad = Unit{
		Name:                "farad",
		InternationalSymbol: "F",
		Representation: Representation{
			FormI: "F",
			FormII: FormII{
				SingleCaseLower: "f",
				SingleCaseUpper: "F",
			},
		},
	}
	Ohm = Unit{
		Name:                "ohm",
		InternationalSymbol: "\u03a9",
		Representation: Representation{
			FormI: "Ohm",
			FormII: FormII{
				SingleCaseLower: "ohm",
				SingleCaseUpper: "OHM",
			},
		},
	}
	Siemens = Unit{
		Name:                "siemens",
		InternationalSymbol: "S",
		Representation: Representation{
			FormI: "S",
			FormII: FormII{
				SingleCaseLower: "sie",
				SingleCaseUpper: "SIE",
			},
		},
	}
	Weber = Unit{
		Name:                "weber",
		InternationalSymbol: "Wb",
		Representation: Representation{
			FormI: "Wb",
			FormII: FormII{
				SingleCaseLower: "wb",
				SingleCaseUpper: "WB",
			},
		},
	}
	Tesla = Unit{
		Name:                "tesla",
		InternationalSymbol: "T",
		Representation: Representation{
			FormI: "T",
			FormII: FormII{
				SingleCaseLower: "t",
				SingleCaseUpper: "T",
			},
		},
	}
	Henry = Unit{
		Name:                "henry",
		InternationalSymbol: "H",
		Representation: Representation{
			FormI: "H",
			FormII: FormII{
				SingleCaseLower: "h",
				SingleCaseUpper: "H",
			},
		},
	}
	DegreeCelcius = Unit{
		Name:                "degree Celsius",
		InternationalSymbol: "\u2103",
		Representation: Representation{
			FormI: "Cel",
			FormII: FormII{
				SingleCaseLower: "cel",
				SingleCaseUpper: "CEL",
			},
		},
	}
	Lumen = Unit{
		Name:                "lumen",
		InternationalSymbol: "lm",
		Representation: Representation{
			FormI: "lm",
			FormII: FormII{
				SingleCaseLower: "lm",
				SingleCaseUpper: "LM",
			},
		},
	}
	Lux = Unit{
		Name:                "lux",
		InternationalSymbol: "lx",
		Representation: Representation{
			FormI: "lx",
			FormII: FormII{
				SingleCaseLower: "lx",
				SingleCaseUpper: "LX",
			},
		},
	}
	Becquerel = Unit{
		Name:                "becquerel",
		InternationalSymbol: "Bq",
		Representation: Representation{
			FormI: "Bq",
			FormII: FormII{
				SingleCaseLower: "bq",
				SingleCaseUpper: "BQ",
			},
		},
	}
	Gray = Unit{
		Name:                "gray",
		InternationalSymbol: "Gy",
		Representation: Representation{
			FormI: "Gy",
			FormII: FormII{
				SingleCaseLower: "gy",
				SingleCaseUpper: "GY",
			},
		},
	}
	Sievert = Unit{
		Name:                "sievert",
		InternationalSymbol: "Sv",
		Representation: Representation{
			FormI: "Sv",
			FormII: FormII{
				SingleCaseLower: "sv",
				SingleCaseUpper: "SV",
			},
		},
	}
)

// Other units from ISO 1000
var (
	AngleGrade = Unit{
		Name:                "grade (angle)",
		InternationalSymbol: "\u1da2",
		Representation: Representation{
			FormI: "gon",
			FormII: FormII{
				SingleCaseLower: "gon",
				SingleCaseUpper: "GON",
			},
		},
	}
	AngleDegree = Unit{
		Name:                "degree (angle)",
		InternationalSymbol: "\u00b0",
		Representation: Representation{
			FormI: "deg",
			FormII: FormII{
				SingleCaseLower: "deg",
				SingleCaseUpper: "DEG",
			},
		},
	}
	AngleMinute = Unit{
		Name:                "minute (angle)",
		InternationalSymbol: "\u2032",
		Representation: Representation{
			FormI: "\u2032",
			FormII: FormII{
				SingleCaseLower: "mnt",
				SingleCaseUpper: "MNT",
			},
		},
	}
	AngleSecond = Unit{
		Name:                "second (angle)",
		InternationalSymbol: "\u2033",
		Representation: Representation{
			FormI: "\u2033",
			FormII: FormII{
				SingleCaseLower: "sec",
				SingleCaseUpper: "SEC",
			},
		},
	}
	Litre = Unit{
		Name:                "litre",
		InternationalSymbol: "l",
		Representation: Representation{
			FormI: "l",
			FormII: FormII{
				SingleCaseLower: "l",
				SingleCaseUpper: "L",
			},
		},
	}
	Are = Unit{
		Name:                "are",
		InternationalSymbol: "a",
		Representation: Representation{
			FormI: "a",
			FormII: FormII{
				SingleCaseLower: "are",
				SingleCaseUpper: "ARE",
			},
		},
	}
	Hectare = Unit{
		Name:                "hectare",
		InternationalSymbol: "ha",
		Representation: Representation{
			FormI: "ha",
			FormII: FormII{
				SingleCaseLower: "har",
				SingleCaseUpper: "HAR",
			},
		},
	}
	Minute = Unit{
		Name:                "minute (time)",
		InternationalSymbol: "min",
		Representation: Representation{
			FormI: "min",
			FormII: FormII{
				SingleCaseLower: "min",
				SingleCaseUpper: "MIN",
			},
		},
	}
	Hour = Unit{
		Name:                "hour",
		InternationalSymbol: "h",
		Representation: Representation{
			FormI: "h",
			FormII: FormII{
				SingleCaseLower: "hr",
				SingleCaseUpper: "HR",
			},
		},
	}
	Day = Unit{
		Name:                "day",
		InternationalSymbol: "d",
		Representation: Representation{
			FormI: "d",
			FormII: FormII{
				SingleCaseLower: "d",
				SingleCaseUpper: "D",
			},
		},
	}
	Year = Unit{
		Name:                "year",
		InternationalSymbol: "a",
		Representation: Representation{
			FormI: "a",
			FormII: FormII{
				SingleCaseLower: "ann",
				SingleCaseUpper: "ANN",
			},
		},
	}
	Gram = Unit{
		Name:                "gram",
		InternationalSymbol: "g",
		Representation: Representation{
			FormI: "g",
			FormII: FormII{
				SingleCaseLower: "g",
				SingleCaseUpper: "G",
			},
		},
	}
	Tonne = Unit{
		Name:                "tonne",
		InternationalSymbol: "t",
		Representation: Representation{
			FormI: "t",
			FormII: FormII{
				SingleCaseLower: "tne",
				SingleCaseUpper: "TNE",
			},
		},
	}
	Bar = Unit{
		Name:                "bar",
		InternationalSymbol: "bar",
		Representation: Representation{
			FormI: "bar",
			FormII: FormII{
				SingleCaseLower: "bar",
				SingleCaseUpper: "BAR",
			},
		},
	}
	Poise = Unit{
		Name:                "poise",
		InternationalSymbol: "P",
		Representation: Representation{
			FormI: "P",
			FormII: FormII{
				SingleCaseLower: "p",
				SingleCaseUpper: "P",
			},
		},
	}
	Stokes = Unit{
		Name:                "stokes",
		InternationalSymbol: "St",
		Representation: Representation{
			FormI: "St",
			FormII: FormII{
				SingleCaseLower: "st",
				SingleCaseUpper: "ST",
			},
		},
	}
	ElectronVolt = Unit{
		Name:                "electronvolt",
		InternationalSymbol: "eV",
		Representation: Representation{
			FormI: "eV",
			FormII: FormII{
				SingleCaseLower: "ev",
				SingleCaseUpper: "EV",
			},
		},
	}
	AtomicMassUnit = Unit{
		Name:                "atomic mass unit",
		InternationalSymbol: "u",
		Representation: Representation{
			FormI: "u",
			FormII: FormII{
				SingleCaseLower: "u",
				SingleCaseUpper: "U",
			},
		},
	}
	AstronomicUnit = Unit{
		Name:                "astronomic unit",
		InternationalSymbol: "AU",
		Representation: Representation{
			FormI: "AU",
			FormII: FormII{
				SingleCaseLower: "asu",
				SingleCaseUpper: "ASU",
			},
		},
	}
	Parsec = Unit{
		Name:                "parsec",
		InternationalSymbol: "pc",
		Representation: Representation{
			FormI: "pc",
			FormII: FormII{
				SingleCaseLower: "prs",
				SingleCaseUpper: "PRS",
			},
		},
	}
)

// PrefixRepresentation holds data for an ISO 2955 prefix.
type PrefixRepresentation struct {
	Prefix               string
	MultiplicationFactor float64
	InternationalSymbol  string
	FormI                string
	FormII               FormII
}

func (p PrefixRepresentation) String() string {
	return fmt.Sprintf("%s (%s)", p.Prefix, p.InternationalSymbol)
}

// Representations of prefixes.
var (
	Exa = PrefixRepresentation{
		Prefix:               "exa",
		MultiplicationFactor: 10e18,
		InternationalSymbol:  "E",
		FormI:                "E",
		FormII: FormII{
			SingleCaseLower: "ex",
			SingleCaseUpper: "EX",
		},
	}
	Peta = PrefixRepresentation{
		Prefix:               "peta",
		MultiplicationFactor: 10e15,
		InternationalSymbol:  "P",
		FormI:                "P",
		FormII: FormII{
			SingleCaseLower: "pe",
			SingleCaseUpper: "PE",
		},
	}
	Tera = PrefixRepresentation{
		Prefix:               "tera",
		MultiplicationFactor: 10e12,
		InternationalSymbol:  "T",
		FormI:                "T",
		FormII: FormII{
			SingleCaseLower: "t",
			SingleCaseUpper: "T",
		},
	}
	Giga = PrefixRepresentation{
		Prefix:               "giga",
		MultiplicationFactor: 10e9,
		InternationalSymbol:  "G",
		FormI:                "G",
		FormII: FormII{
			SingleCaseLower: "g",
			SingleCaseUpper: "G",
		},
	}
	Mega = PrefixRepresentation{
		Prefix:               "mega",
		MultiplicationFactor: 10e6,
		InternationalSymbol:  "M",
		FormI:                "M",
		FormII: FormII{
			SingleCaseLower: "ma",
			SingleCaseUpper: "MA",
		},
	}
	Kilo = PrefixRepresentation{
		Prefix:               "kilo",
		MultiplicationFactor: 10e3,
		InternationalSymbol:  "k",
		FormI:                "k",
		FormII: FormII{
			SingleCaseLower: "k",
			SingleCaseUpper: "K",
		},
	}
	Hecto = PrefixRepresentation{
		Prefix:               "hecto",
		MultiplicationFactor: 10e2,
		InternationalSymbol:  "h",
		FormI:                "h",
		FormII: FormII{
			SingleCaseLower: "h",
			SingleCaseUpper: "H",
		},
	}
	Deca = PrefixRepresentation{
		Prefix:               "deca",
		MultiplicationFactor: 10e1,
		InternationalSymbol:  "da",
		FormI:                "da",
		FormII: FormII{
			SingleCaseLower: "da",
			SingleCaseUpper: "DA",
		},
	}
	Deci = PrefixRepresentation{
		Prefix:               "deci",
		MultiplicationFactor: 10e-1,
		InternationalSymbol:  "d",
		FormI:                "d",
		FormII: FormII{
			SingleCaseLower: "d",
			SingleCaseUpper: "D",
		},
	}
	Centi = PrefixRepresentation{
		Prefix:               "centi",
		MultiplicationFactor: 10e-2,
		InternationalSymbol:  "c",
		FormI:                "c",
		FormII: FormII{
			SingleCaseLower: "c",
			SingleCaseUpper: "C",
		},
	}
	Milli = PrefixRepresentation{
		Prefix:               "milli",
		MultiplicationFactor: 10e-3,
		InternationalSymbol:  "m",
		FormI:                "m",
		FormII: FormII{
			SingleCaseLower: "m",
			SingleCaseUpper: "M",
		},
	}
	Micro = PrefixRepresentation{
		Prefix:               "micro",
		MultiplicationFactor: 10e-6,
		InternationalSymbol:  "\u00b5",
		FormI:                "u",
		FormII: FormII{
			SingleCaseLower: "u",
			SingleCaseUpper: "U",
		},
	}
	Nano = PrefixRepresentation{
		Prefix:               "nano",
		MultiplicationFactor: 10e-9,
		InternationalSymbol:  "n",
		FormI:                "n",
		FormII: FormII{
			SingleCaseLower: "n",
			SingleCaseUpper: "N",
		},
	}
	Pico = PrefixRepresentation{
		Prefix:               "pico",
		MultiplicationFactor: 10e-12,
		InternationalSymbol:  "p",
		FormI:                "p",
		FormII: FormII{
			SingleCaseLower: "p",
			SingleCaseUpper: "P",
		},
	}
	Femto = PrefixRepresentation{
		Prefix:               "femto",
		MultiplicationFactor: 10e-15,
		InternationalSymbol:  "f",
		FormI:                "f",
		FormII: FormII{
			SingleCaseLower: "f",
			SingleCaseUpper: "F",
		},
	}
	Atto = PrefixRepresentation{
		Prefix:               "atto",
		MultiplicationFactor: 10e-18,
		InternationalSymbol:  "a",
		FormI:                "a",
		FormII: FormII{
			SingleCaseLower: "a",
			SingleCaseUpper: "A",
		},
	}
)
