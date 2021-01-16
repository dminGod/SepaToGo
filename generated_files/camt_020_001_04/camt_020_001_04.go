// Code generated by ___go_build_github_com_dminGod_sepaToGo. DO NOT EDIT.

package iso20022_camt_020_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type BusinessInformationCriteria1 struct {
	NewQryNm Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 NewQryNm,omitempty"`
	SchCrit  []GeneralBusinessInformationSearchCriteria1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 SchCrit,omitempty"`
	RtrCrit  GeneralBusinessInformationReturnCriteria1   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 RtrCrit,omitempty"`
}

type BusinessInformationQueryDefinition3 struct {
	QryTp         QueryType2Code                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 QryTp,omitempty"`
	GnlBizInfCrit GeneralBusinessInformationCriteriaDefinition1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 GnlBizInfCrit,omitempty"`
}

type CharacterSearch1Choice struct {
	EQ  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 EQ"`
	NEQ Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 NEQ"`
	CT  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 CT"`
	NCT Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 NCT"`
}

type Document struct {
	GetGnlBizInf GetGeneralBusinessInformationV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 GetGnlBizInf"`
}

type GeneralBusinessInformationCriteriaDefinition1Choice struct {
	QryNm   Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 QryNm"`
	NewCrit BusinessInformationCriteria1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 NewCrit"`
}

type GeneralBusinessInformationReturnCriteria1 struct {
	QlfrInd     bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 QlfrInd,omitempty"`
	SbjtInd     bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 SbjtInd,omitempty"`
	SbjtDtlsInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 SbjtDtlsInd,omitempty"`
}

type GeneralBusinessInformationSearchCriteria1 struct {
	Ref  []Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 Ref,omitempty"`
	Sbjt []CharacterSearch1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 Sbjt,omitempty"`
	Qlfr []InformationQualifierType1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 Qlfr,omitempty"`
}

type GetGeneralBusinessInformationV04 struct {
	MsgHdr          MessageHeader1                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 MsgHdr"`
	GnlBizInfQryDef BusinessInformationQueryDefinition3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 GnlBizInfQryDef,omitempty"`
	SplmtryData     []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 SplmtryData,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type InformationQualifierType1 struct {
	IsFrmtd bool          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 IsFrmtd,omitempty"`
	Prty    Priority1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 Prty,omitempty"`
}

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

type MessageHeader1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 MsgId"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 CreDtTm,omitempty"`
}

// May be one of HIGH, NORM, LOWW
type Priority1Code string

// May be one of ALLL, CHNG, MODF, DELD
type QueryType2Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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