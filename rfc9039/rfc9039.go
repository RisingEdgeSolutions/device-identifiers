// SPDX-License-Identifier: BSD-3-Clause

// Package rfc9039 provides tools for parsing RFC 9039 specified urn:dev strings.
package rfc9039

import (
	"errors"
	"regexp"
	"strings"
)

const UrnDevPrefix = "urn:dev:"

const UrnDevMaxSectionCount = 16

const SubTypeRegEx = "^[a-z][0-9a-z]*$"
const DevUrnReservedNoDashRegEx = "^[A-Za-z0-9\\.]+$"
const DevUrnReservedRegEx = "^[A-Za-z0-9\\.\\-]+$"
const HexStringRegEx = "^([0-9a-f][0-9a-f])+$"
const PosNumberRegEx = "^[1-9][0-9]*$"

// UrnDev captures Parse output for urn:dev into own fields.
type UrnDev struct {
	// FullName is the full formed urn:dev string.
	FullName string
	// Subtype defines what type of urn:dev is specified. Examples for subtypes are "mac", "ow", "org", "os", "ops" and also any valid future types.
	Subtype string
	// Organization captures value of organization for Subtype "org", "os" and "ops".
	Organization string
	// Product captures value of product for Subtype "ops".
	Product string
	// Serial captures value of serial number for Subtype "os" and "ops".
	Serial string
	// Component captures all components.
	Component []string
	// Identifier captures all identifiers.
	Identifier []string
	// Eui64Identifier captures value of mac address for Subtype "mac".
	Eui64Identifier string
	// OwIdentifier captures value of 1-wire address for Subtype "ow"
	OwIdentifier string
}

// HasUrnDevPrefix can be used to determine whether urn:dev prefix is present, and it would be suitable for parsing with Parse.
func HasUrnDevPrefix(name string) bool {
	return strings.HasPrefix(strings.ToLower(name), UrnDevPrefix)
}

func isValidEui64(name string) bool {
	if len(name) != 16 {
		return false
	}

	match, _ := regexp.MatchString(HexStringRegEx, name)

	return match
}

func isValidOwAddress(name string) bool {
	if len(name) != 16 {
		return false
	}

	match, _ := regexp.MatchString(HexStringRegEx, name)

	return match
}

func isValidPosNumber(name string) bool {
	match, _ := regexp.MatchString(PosNumberRegEx, name)

	return match
}

func isValidIdentifier(name string) bool {
	match, _ := regexp.MatchString(DevUrnReservedRegEx, name)

	return match
}

func isValidIdentifierNoDash(name string) bool {
	match, _ := regexp.MatchString(DevUrnReservedNoDashRegEx, name)

	return match
}

func isValidSubType(name string) bool {
	match, _ := regexp.MatchString(SubTypeRegEx, name)

	return match
}

// Parse parses RFC 9039 specified urn:dev into its components. If incorrectly formed urn:dev string is given as input an error is returned.
func Parse(name string) (UrnDev, error) {
	// From: RFC 9039 - Uniform Resource Names for Device Identifiers
	//
	// 3.2.  Syntax
	//
	//   The identifier is expressed in ASCII characters and has a
	//   hierarchical structure as follows:
	//
	//   devurn = "urn:dev:" body componentpart
	//   body = macbody / owbody / orgbody / osbody / opsbody / otherbody
	//   macbody = %s"mac:" hexstring
	//   owbody = %s"ow:" hexstring
	//   orgbody = %s"org:" posnumber "-" identifier *( ":" identifier )
	//   osbody = %s"os:" posnumber "-" serial *( ":" identifier )
	//   opsbody = %s"ops:" posnumber "-" product "-" serial
	//             *( ":" identifier )
	//   otherbody = subtype ":" identifier *( ":" identifier )
	//   subtype = LALPHA *(DIGIT / LALPHA)
	//   identifier = 1*devunreserved
	//   identifiernodash = 1*devunreservednodash
	//   product = identifiernodash
	//   serial = identifier
	//   componentpart = *( "_" identifier )
	//   devunreservednodash = ALPHA / DIGIT / "."
	//   devunreserved = devunreservednodash / "-"
	//   hexstring = 1*(hexdigit hexdigit)
	//   hexdigit = DIGIT / "a" / "b" / "c" / "d" / "e" / "f"
	//   posnumber = NZDIGIT *DIGIT
	//   ALPHA =  %x41-5A / %x61-7A
	//   LALPHA =  %x41-5A
	//   NZDIGIT = %x31-39
	//   DIGIT =  %x30-39

	out := UrnDev{}

	out.FullName = name

	sections := strings.Split(name, ":")

	if len(sections) < 4 || len(sections) >= UrnDevMaxSectionCount {
		return UrnDev{}, errors.New("invalid input")
	}

	// urn needs to be normalized for comparison
	if strings.ToLower(sections[0]) != "urn" {
		return UrnDev{}, errors.New("invalid input (missing urn)")
	}

	// dev needs to be normalized for comparison
	if strings.ToLower(sections[1]) != "dev" {
		return UrnDev{}, errors.New("invalid input (missing dev)")
	}

	out.Subtype = sections[2]

	var componentPart = strings.Split(sections[len(sections)-1], "_")

	if len(componentPart) > 1 {
		// Remove the processed component part
		sections[len(sections)-1] = componentPart[0]

		out.Component = componentPart[1:]
		for _, component := range out.Component {
			if !isValidIdentifier(component) {
				return UrnDev{}, errors.New("invalid input (componentpart)")
			}
		}
	} else {
		out.Component = []string{}
	}

	out.Identifier = sections[3:]
	for _, identifier := range out.Identifier {
		if !isValidIdentifier(identifier) {
			return UrnDev{}, errors.New("invalid input (identifier)")
		}
	}

	switch out.Subtype {
	case "mac":
		if len(sections) == 5 {
			return UrnDev{}, errors.New("invalid input (mac)")
		}

		out.Identifier = out.Identifier[1:]
		out.Eui64Identifier = sections[3]

		if !isValidEui64(out.Eui64Identifier) {
			return UrnDev{}, errors.New("invalid input (EUI-64)")
		}

	case "ow":
		if len(sections) == 5 {
			return UrnDev{}, errors.New("invalid input (ow)")
		}

		out.Identifier = out.Identifier[1:]
		out.OwIdentifier = sections[3]

		if !isValidOwAddress(out.OwIdentifier) {
			return UrnDev{}, errors.New("invalid input (ow)")
		}

	case "org":
		org := strings.SplitN(sections[3], "-", 2)

		if len(org) != 2 {
			return UrnDev{}, errors.New("invalid input (org)")
		}

		out.Organization = org[0]
		if !isValidPosNumber(out.Organization) {
			return UrnDev{}, errors.New("invalid input (org)")
		}

		if !isValidIdentifier(org[1]) {
			return UrnDev{}, errors.New("invalid input (org)")
		}

		// Inject organization part's identifier into identifier index 0
		out.Identifier[0] = org[1]

	case "os":
		out.Identifier = out.Identifier[1:]
		os := strings.SplitN(sections[3], "-", 2)

		if len(os) != 2 {
			return UrnDev{}, errors.New("invalid input (os)")
		}

		out.Organization = os[0]
		if !isValidPosNumber(out.Organization) {
			return UrnDev{}, errors.New("invalid input (os)")
		}

		out.Serial = os[1]
		if !isValidIdentifier(out.Serial) {
			return UrnDev{}, errors.New("invalid input (os)")
		}

	case "ops":
		out.Identifier = out.Identifier[1:]
		ops := strings.Split(sections[3], "-")

		if len(ops) != 3 {
			return UrnDev{}, errors.New("invalid input (ops)")
		}

		out.Organization = ops[0]
		if !isValidPosNumber(out.Organization) {
			return UrnDev{}, errors.New("invalid input (ops)")
		}
		out.Product = ops[1]
		if !isValidIdentifierNoDash(out.Product) {
			return UrnDev{}, errors.New("invalid input (ops)")
		}
		out.Serial = ops[2]
		if !isValidIdentifier(out.Serial) {
			return UrnDev{}, errors.New("invalid input (ops)")
		}

	default:
		// otherbody
		if !isValidSubType(out.Subtype) {
			return UrnDev{}, errors.New("invalid sub type (" + out.Subtype + ")")
		}
	}

	return out, nil
}
