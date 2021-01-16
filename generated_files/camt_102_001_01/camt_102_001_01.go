// Code generated by ___go_build_github_com_dminGod_sepaToGo. DO NOT EDIT.

package iso20022_camt_102_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Prtry"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 AmtWthCcy"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 MmbId"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CreateStandingOrderV01 struct {
	MsgHdr      MessageHeader1               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 MsgHdr"`
	StgOrdrId   StandingOrderIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 StgOrdrId"`
	ValSet      StandingOrder7               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 ValSet"`
	SplmtryData []SupplementaryData1         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 SplmtryData,omitempty"`
}

type DatePeriod2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 ToDt"`
}

type DatePeriod2Choice struct {
	FrDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 FrDt"`
	ToDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 ToDt"`
	FrToDt DatePeriod2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 FrToDt"`
}

type Document struct {
	CretStgOrdr CreateStandingOrderV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 CretStgOrdr"`
}

type EventType1Choice struct {
	Cd    ExternalSystemEventType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Prtry"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type ExecutionType1Choice struct {
	Tm  ISOTime          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Tm"`
	Evt EventType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Evt"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// Must be at least 1 items long
type ExternalSystemEventType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Othr,omitempty"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, OVNG
type Frequency2Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type MessageHeader1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 MsgId"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 CreDtTm,omitempty"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 AdrLine,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Prtry"`
}

type StandingOrder7 struct {
	Amt          Amount2Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Amt,omitempty"`
	Cdtr         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Cdtr,omitempty"`
	CdtrAcct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 CdtrAcct,omitempty"`
	Dbtr         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Dbtr,omitempty"`
	DbtrAcct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 DbtrAcct,omitempty"`
	ExctnTp      ExecutionType1Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 ExctnTp,omitempty"`
	Frqcy        Frequency2Code                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Frqcy,omitempty"`
	VldtyPrd     DatePeriod2Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 VldtyPrd,omitempty"`
	ZeroSweepInd bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 ZeroSweepInd,omitempty"`
}

type StandingOrderIdentification4 struct {
	Id       Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Id,omitempty"`
	Acct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Acct"`
	AcctOwnr BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 AcctOwnr,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}