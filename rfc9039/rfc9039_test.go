// SPDX-License-Identifier: BSD-3-Clause

package rfc9039

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleHasUrnDevPrefix() {
	fmt.Println(HasUrnDevPrefix("urn:dev:ops:32473-Refrigerator-5002"))
	fmt.Println(HasUrnDevPrefix("urn:not-dev:value"))
	// Output: true
	// false
}

func ExampleParse() {
	devUrn, _ := Parse("urn:dev:ops:32473-Refrigerator-5002")
	fmt.Println(devUrn.Organization)
	fmt.Println(devUrn.Product)
	fmt.Println(devUrn.Serial)
	// Output: 32473
	// Refrigerator
	// 5002
}

func assertEmptyUrnDevStruct(t *testing.T, value UrnDev) {
	assert.Equal(t, "", value.FullName)
	assert.Equal(t, "", value.Subtype)
	assert.Equal(t, "", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "", value.Serial)
	assert.Equal(t, []string(nil), value.Component)
	assert.Equal(t, []string(nil), value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestIsUrnDev(t *testing.T) {
	assert.True(t, HasUrnDevPrefix("urn:dev:mac:0024beffff804ff1"))
	assert.True(t, HasUrnDevPrefix("URN:DEV:mac:0024beffff804ff1"))
	assert.False(t, HasUrnDevPrefix("urn:uuid:f81d4fae-7dec-11d0-a765-00a0c91e6bf6"))
}

func TestUrnDevInvalidNoSubtype(t *testing.T) {
	value, err := Parse("urn:dev:")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevInvalidMissingUrn(t *testing.T) {
	value, err := Parse("foo:dev:mac:0024beffff804ff1")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevInvalidMissingDev(t *testing.T) {
	value, err := Parse("urn:foo:mac:0024beffff804ff1")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevInvalidComponent(t *testing.T) {
	value, err := Parse("urn:dev:mac:0024beffff804ff1_fa%il")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevInvalidIdentifier(t *testing.T) {
	value, err := Parse("urn:dev:mac:0024beffff804ff1:fa%il")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevMac(t *testing.T) {
	value, err := Parse("urn:dev:mac:0024beffff804ff1")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:mac:0024beffff804ff1", value.FullName)
	assert.Equal(t, "mac", value.Subtype)
	assert.Equal(t, "", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "", value.Serial)
	assert.Equal(t, []string{}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "0024beffff804ff1", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevMac2(t *testing.T) {
	value, err := Parse("urn:dev:mac:0024befffe804ff1")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:mac:0024befffe804ff1", value.FullName)
	assert.Equal(t, "mac", value.Subtype)
	assert.Equal(t, "", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "", value.Serial)
	assert.Equal(t, []string{}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "0024befffe804ff1", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevMac3(t *testing.T) {
	value, err := Parse("urn:dev:mac:acde48234567019f")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:mac:acde48234567019f", value.FullName)
	assert.Equal(t, "mac", value.Subtype)
	assert.Equal(t, "", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "", value.Serial)
	assert.Equal(t, []string{}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "acde48234567019f", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevMacInvalidIdentifier(t *testing.T) {
	value, err := Parse("urn:dev:mac:acde48234567019f:invalid")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevMacInvalidTooShort(t *testing.T) {
	value, err := Parse("urn:dev:mac:acde48234567019")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevMacInvalidTooLong(t *testing.T) {
	value, err := Parse("urn:dev:mac:acde48234567019fa")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevMacInvalidCharacters(t *testing.T) {
	value, err := Parse("urn:dev:mac:acdefail4567019f")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevMacNoValue(t *testing.T) {
	value, err := Parse("urn:dev:mac")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}
func TestUrnDevMacInvalidValue(t *testing.T) {
	value, err := Parse("urn:dev:mac:")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOw(t *testing.T) {
	value, err := Parse("urn:dev:ow:10e2073a01080063")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:ow:10e2073a01080063", value.FullName)
	assert.Equal(t, "ow", value.Subtype)
	assert.Equal(t, "", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "", value.Serial)
	assert.Equal(t, []string{}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "10e2073a01080063", value.OwIdentifier)
}

func TestUrnDevOw2(t *testing.T) {
	value, err := Parse("urn:dev:ow:264437f5000000ed_humidity")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:ow:264437f5000000ed_humidity", value.FullName)
	assert.Equal(t, "ow", value.Subtype)
	assert.Equal(t, "", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "", value.Serial)
	assert.Equal(t, []string{"humidity"}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "264437f5000000ed", value.OwIdentifier)
}

func TestUrnDevOw3(t *testing.T) {
	value, err := Parse("urn:dev:ow:264437f5000000ed_temperature")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:ow:264437f5000000ed_temperature", value.FullName)
	assert.Equal(t, "ow", value.Subtype)
	assert.Equal(t, "", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "", value.Serial)
	assert.Equal(t, []string{"temperature"}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "264437f5000000ed", value.OwIdentifier)
}

func TestUrnDevOwInvalidIdentifier(t *testing.T) {
	value, err := Parse("urn:dev:ow:264437f5000000ed:invalid_humidity")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOwInvalidIdentifier2(t *testing.T) {
	value, err := Parse("urn:dev:ow:10e2073a01080063:invalid")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOwInvalidTooShort(t *testing.T) {
	value, err := Parse("urn:dev:ow:10e2073a0108006")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOwInvalidTooLong(t *testing.T) {
	value, err := Parse("urn:dev:ow:10e2073a01080063a")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOwInvalidCharacters(t *testing.T) {
	value, err := Parse("urn:dev:ow:10e20fail1080063")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOrg(t *testing.T) {
	value, err := Parse("urn:dev:org:32473-foo")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:org:32473-foo", value.FullName)
	assert.Equal(t, "org", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "", value.Serial)
	assert.Equal(t, []string{}, value.Component)
	assert.Equal(t, []string{"foo"}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOrgWithComponent(t *testing.T) {
	value, err := Parse("urn:dev:org:32473-foo_component")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:org:32473-foo_component", value.FullName)
	assert.Equal(t, "org", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "", value.Serial)
	assert.Equal(t, []string{"component"}, value.Component)
	assert.Equal(t, []string{"foo"}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOrgWithMultipleIdentifiers(t *testing.T) {
	value, err := Parse("urn:dev:org:32473-foo:bar:zoo")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:org:32473-foo:bar:zoo", value.FullName)
	assert.Equal(t, "org", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "", value.Serial)
	assert.Equal(t, []string{}, value.Component)
	assert.Equal(t, []string{"foo", "bar", "zoo"}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}
func TestUrnDevOrgInvalidOrg(t *testing.T) {
	value, err := Parse("urn:dev:org:032473-foo")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOrgInvalidOrgCharacters(t *testing.T) {
	value, err := Parse("urn:dev:org:32473fail-foo")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOrgInvalidNoDashes(t *testing.T) {
	value, err := Parse("urn:dev:org:32473")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOrgInvalidNoIdentifier(t *testing.T) {
	value, err := Parse("urn:dev:org:32473-")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOs(t *testing.T) {
	value, err := Parse("urn:dev:os:32473-123456")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:os:32473-123456", value.FullName)
	assert.Equal(t, "os", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "123456", value.Serial)
	assert.Equal(t, []string{}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOsWithComponent(t *testing.T) {
	value, err := Parse("urn:dev:os:32473-123456_component")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:os:32473-123456_component", value.FullName)
	assert.Equal(t, "os", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "123456", value.Serial)
	assert.Equal(t, []string{"component"}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOsWithIdentifierAndComponent(t *testing.T) {
	value, err := Parse("urn:dev:os:32473-123456:identifier_component")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:os:32473-123456:identifier_component", value.FullName)
	assert.Equal(t, "os", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "123456", value.Serial)
	assert.Equal(t, []string{"component"}, value.Component)
	assert.Equal(t, []string{"identifier"}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOsSerialWithDashes(t *testing.T) {
	value, err := Parse("urn:dev:os:32473-12-34-56")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:os:32473-12-34-56", value.FullName)
	assert.Equal(t, "os", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "12-34-56", value.Serial)
	assert.Equal(t, []string{}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOsSerialWithDashesWithComponent(t *testing.T) {
	value, err := Parse("urn:dev:os:32473-12-34-56_component")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:os:32473-12-34-56_component", value.FullName)
	assert.Equal(t, "os", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "12-34-56", value.Serial)
	assert.Equal(t, []string{"component"}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOsSerialWithDashesWithIdentrifierAndComponent(t *testing.T) {
	value, err := Parse("urn:dev:os:32473-12-34-56:identifier_component")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:os:32473-12-34-56:identifier_component", value.FullName)
	assert.Equal(t, "os", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "12-34-56", value.Serial)
	assert.Equal(t, []string{"component"}, value.Component)
	assert.Equal(t, []string{"identifier"}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOsInvalidOrg(t *testing.T) {
	value, err := Parse("urn:dev:os:032473-12-34-56")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOsInvalidOrgCharacters(t *testing.T) {
	value, err := Parse("urn:dev:os:32473fail-12-34-56")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOsInvalidNoSerial(t *testing.T) {
	value, err := Parse("urn:dev:os:32473")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOsInvalidInvalidSerial(t *testing.T) {
	value, err := Parse("urn:dev:os:32473-")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOps(t *testing.T) {
	value, err := Parse("urn:dev:ops:32473-Refrigerator-5002")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:ops:32473-Refrigerator-5002", value.FullName)
	assert.Equal(t, "ops", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "Refrigerator", value.Product)
	assert.Equal(t, "5002", value.Serial)
	assert.Equal(t, []string{}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOpsWithComponent(t *testing.T) {
	value, err := Parse("urn:dev:ops:32473-Refrigerator-5002_component")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:ops:32473-Refrigerator-5002_component", value.FullName)
	assert.Equal(t, "ops", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "Refrigerator", value.Product)
	assert.Equal(t, "5002", value.Serial)
	assert.Equal(t, []string{"component"}, value.Component)
	assert.Equal(t, []string{}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOpsWithIdentifier(t *testing.T) {
	value, err := Parse("urn:dev:ops:32473-Refrigerator-5002:identifier")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:ops:32473-Refrigerator-5002:identifier", value.FullName)
	assert.Equal(t, "ops", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "Refrigerator", value.Product)
	assert.Equal(t, "5002", value.Serial)
	assert.Equal(t, []string{}, value.Component)
	assert.Equal(t, []string{"identifier"}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOpsWithIdentifierAndComponent(t *testing.T) {
	value, err := Parse("urn:dev:ops:32473-Refrigerator-5002:identifier_component")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:ops:32473-Refrigerator-5002:identifier_component", value.FullName)
	assert.Equal(t, "ops", value.Subtype)
	assert.Equal(t, "32473", value.Organization)
	assert.Equal(t, "Refrigerator", value.Product)
	assert.Equal(t, "5002", value.Serial)
	assert.Equal(t, []string{"component"}, value.Component)
	assert.Equal(t, []string{"identifier"}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevOpsInvalidOrg(t *testing.T) {
	value, err := Parse("urn:dev:ops:032473-Refrigerator-5002")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOpsInvalidOrgCharacters(t *testing.T) {
	value, err := Parse("urn:dev:ops:32473fail-Refrigerator-5002")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOpsInvalidNoProductAndSerial(t *testing.T) {
	value, err := Parse("urn:dev:ops:32473")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOpsInvalidInvalidProduct(t *testing.T) {
	value, err := Parse("urn:dev:ops:32473--5002")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOpsInvalidNoSerial(t *testing.T) {
	value, err := Parse("urn:dev:ops:32473-Refrigerator")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevOpsInvalidInvalidSerial(t *testing.T) {
	value, err := Parse("urn:dev:ops:32473-Refrigerator-")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}

func TestUrnDevExample(t *testing.T) {
	value, err := Parse("urn:dev:example:new-1-2-3_comp")
	if err != nil {
		t.Fatalf("Failed to parse")
		return
	}
	assert.Equal(t, "urn:dev:example:new-1-2-3_comp", value.FullName)
	assert.Equal(t, "example", value.Subtype)
	assert.Equal(t, "", value.Organization)
	assert.Equal(t, "", value.Product)
	assert.Equal(t, "", value.Serial)
	assert.Equal(t, []string{"comp"}, value.Component)
	assert.Equal(t, []string{"new-1-2-3"}, value.Identifier)
	assert.Equal(t, "", value.Eui64Identifier)
	assert.Equal(t, "", value.OwIdentifier)
}

func TestUrnDevInvalidSubType(t *testing.T) {
	value, err := Parse("urn:dev:INVALID:new-1-2-3_comp")
	if err == nil {
		t.Fatalf("Failed to parse")
		return
	}

	assertEmptyUrnDevStruct(t, value)
}
