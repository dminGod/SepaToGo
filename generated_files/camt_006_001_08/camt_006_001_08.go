// Code generated by ___go_build_github_com_dminGod_sepaToGo. DO NOT EDIT.

package iso20022_camt_006_001_08

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AmtWthCcy"`
}

type Amount3Choice struct {
	AmtWthCcy  ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AmtWthCcy"`
	AmtWthtCcy float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AmtWthtCcy"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PstlAdr,omitempty"`
}

// May be one of CANI, CANS, CSUB
type CancelledStatusReason1Code string

type CashAccount39 struct {
	Id   AccountIdentification4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id"`
	Tp   CashAccountType2Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Ccy,omitempty"`
	Nm   Max70Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Nm,omitempty"`
	Prxy ProxyAccountIdentification1                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prxy,omitempty"`
	Ownr PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Ownr,omitempty"`
	Svcr BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Svcr,omitempty"`
}

type CashAccountAndEntry3 struct {
	Acct CashAccount39 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Acct"`
	Ntry CashEntry2    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Ntry,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

type CashEntry2 struct {
	Amt          ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Amt,omitempty"`
	Dt           DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Dt,omitempty"`
	Sts          EntryStatus1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Sts,omitempty"`
	Id           Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id,omitempty"`
	StmtId       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 StmtId,omitempty"`
	AcctSvcrRef  float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AcctSvcrRef,omitempty"`
	AddtlNtryInf []Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AddtlNtryInf,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 DtTm"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CtryOfBirth"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 ToDtTm"`
}

type DateTimePeriod1Choice struct {
	FrDtTm ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 FrDtTm"`
	ToDtTm ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 ToDtTm"`
	DtTmRg DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 DtTmRg"`
}

type Document struct {
	RtrTx ReturnTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 RtrTx"`
}

// May be one of BOOK, PDNG, FUTR
type EntryStatus1Code string

// Must match the pattern [BEOVW]{1,1}[0-9]{2,2}|DUM
type EntryTypeIdentifier string

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Err"`
	Desc Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Desc,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalMarketInfrastructure1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// Must be at least 1 items long
type ExternalSystemErrorHandling1Code string

// May be one of STLD, RJTD, CAND, FNLD
type FinalStatus1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Issr,omitempty"`
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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LongPaymentIdentification2 struct {
	TxId           Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 TxId,omitempty"`
	UETR           UUIDv4Identifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 UETR,omitempty"`
	IntrBkSttlmAmt float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 IntrBkSttlmAmt"`
	IntrBkSttlmDt  ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 IntrBkSttlmDt"`
	PmtMtd         PaymentOrigin1Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PmtMtd,omitempty"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 InstdAgt"`
	NtryTp         EntryTypeIdentifier                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 NtryTp,omitempty"`
	EndToEndId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 EndToEndId,omitempty"`
}

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

// Must be at least 1 items long
type Max10Text string

// Must be at least 1 items long
type Max128Text string

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max20000Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max4Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max70Text string

type MessageHeader8 struct {
	MsgId       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 MsgId"`
	CreDtTm     ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CreDtTm,omitempty"`
	MsgPgntn    Pagination1            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 MsgPgntn,omitempty"`
	OrgnlBizQry OriginalBusinessQuery1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 OrgnlBizQry,omitempty"`
	ReqTp       RequestType4Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 ReqTp,omitempty"`
	QryNm       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 QryNm,omitempty"`
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type NumberAndSumOfTransactions2 struct {
	NbOfNtries    Max15NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 NbOfNtries,omitempty"`
	Sum           float64          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Sum,omitempty"`
	TtlNetNtryAmt float64          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 TtlNetNtryAmt,omitempty"`
	CdtDbtInd     CreditDebitCode  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CdtDbtInd,omitempty"`
}

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

type OriginalBusinessQuery1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 MsgNmId,omitempty"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CreDtTm,omitempty"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id,omitempty"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 LastPgInd"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PrvtId"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Agt"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CtctDtls,omitempty"`
}

type PaymentCommon4 struct {
	PmtFr       System2                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PmtFr,omitempty"`
	PmtTo       System2                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PmtTo,omitempty"`
	CmonSts     []PaymentStatus6       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CmonSts,omitempty"`
	ReqdExctnDt DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 ReqdExctnDt,omitempty"`
	NtryDt      DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 NtryDt,omitempty"`
	CdtDbtInd   CreditDebitCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CdtDbtInd,omitempty"`
	PmtMtd      PaymentOrigin1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PmtMtd,omitempty"`
}

type PaymentIdentification6Choice struct {
	TxId      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 TxId"`
	QId       QueueTransactionIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 QId"`
	LngBizId  LongPaymentIdentification2      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 LngBizId"`
	ShrtBizId ShortPaymentIdentification2     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 ShrtBizId"`
	PrtryId   Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PrtryId"`
}

type PaymentInstruction32 struct {
	MsgId          Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 MsgId,omitempty"`
	ReqdExctnDt    DateAndDateTime2Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 ReqdExctnDt,omitempty"`
	Sts            []PaymentStatus6         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Sts,omitempty"`
	InstdAmt       Amount3Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 InstdAmt,omitempty"`
	IntrBkSttlmAmt Amount2Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 IntrBkSttlmAmt,omitempty"`
	Purp           Max10Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Purp,omitempty"`
	PmtMtd         PaymentOrigin1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PmtMtd,omitempty"`
	Prty           Priority1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prty,omitempty"`
	PrcgVldtyTm    DateTimePeriod1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PrcgVldtyTm,omitempty"`
	InstrCpy       Max20000Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 InstrCpy,omitempty"`
	Tp             PaymentType4Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Tp,omitempty"`
	GnrtdOrdr      bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 GnrtdOrdr,omitempty"`
	TxId           Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 TxId,omitempty"`
	IntrBkSttlmDt  ISODate                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 IntrBkSttlmDt,omitempty"`
	EndToEndId     Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 EndToEndId,omitempty"`
	Pties          PaymentTransactionParty3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Pties,omitempty"`
}

// May be one of BDT, BCT, CDT, CCT, CHK, BKT, DCP, CCP, RTI, CAN
type PaymentInstrument1Code string

type PaymentOrigin1Choice struct {
	FINMT    Max3NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 FINMT"`
	XMLMsgNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 XMLMsgNm"`
	Prtry    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
	Instrm   PaymentInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Instrm"`
}

type PaymentStatus6 struct {
	Cd   PaymentStatusCode6Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd,omitempty"`
	DtTm DateAndDateTime2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 DtTm,omitempty"`
	Rsn  []PaymentStatusReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Rsn,omitempty"`
}

type PaymentStatusCode6Choice struct {
	Pdg   PendingStatus4Code   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Pdg"`
	Fnl   FinalStatus1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Fnl"`
	RTGS  Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 RTGS"`
	Sttlm Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Sttlm"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

type PaymentStatusReason1Choice struct {
	Umtchd       UnmatchedStatusReason1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Umtchd"`
	Canc         CancelledStatusReason1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Canc"`
	Sspd         SuspendedStatusReason1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Sspd"`
	PdgFlngSttlm PendingFailingSettlement1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PdgFlngSttlm"`
	PdgSttlm     PendingSettlement2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PdgSttlm"`
	PrtryRjctn   ProprietaryStatusJustification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PrtryRjctn"`
	Prtry        Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

type PaymentTransactionParty3 struct {
	InstgAgt         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 InstgAgt,omitempty"`
	InstdAgt         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 InstdAgt,omitempty"`
	UltmtDbtr        Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 UltmtDbtr,omitempty"`
	Dbtr             Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Dbtr,omitempty"`
	DbtrAgt          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 DbtrAgt,omitempty"`
	InstgRmbrsmntAgt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 InstgRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 InstdRmbrsmntAgt,omitempty"`
	IntrmyAgt1       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 IntrmyAgt1,omitempty"`
	IntrmyAgt2       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 IntrmyAgt2,omitempty"`
	IntrmyAgt3       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 IntrmyAgt3,omitempty"`
	CdtrAgt          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CdtrAgt,omitempty"`
	Cdtr             Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cdtr,omitempty"`
	UltmtCdtr        Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 UltmtCdtr,omitempty"`
}

// May be one of CBS, BCK, BAL, CLS, CTR, CBH, CBP, DPG, DPN, EXP, TCH, LMT, LIQ, DPP, DPH, DPS, STF, TRP, TCS, LOA, LOR, TCP, OND, MGL
type PaymentType3Code string

type PaymentType4Choice struct {
	Cd    PaymentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

// May be one of AWMO, AWSH, LAAW, DOCY, CLAT, CERT, MINO, PHSE, SBLO, DKNY, STCD, BENO, LACK, LATE, CANR, MLAT, OBJT, DOCC, BLOC, CHAS, NEWI, CLAC, PART, CMON, COLL, DEPO, FLIM, NOFX, INCA, LINK, BYIY, CAIS, LALO, MONY, NCON, YCOL, REFS, SDUT, CYCL, BATC, GUAD, PREA, GLOB, CPEC, MUNO
type PendingFailingSettlement1Code string

// May be one of AWMO, CAIS, REFU, AWSH, PHSE, TAMM, DOCY, DOCC, BLOC, CHAS, NEWI, CLAC, MUNO, GLOB, PREA, GUAD, PART, NMAS, CMON, YCOL, COLL, DEPO, FLIM, NOFX, INCA, LINK, FUTU, LACK, LALO, MONY, NCON, REFS, SDUT, BATC, CYCL, SBLO, CPEC, MINO, PCAP
type PendingSettlement2Code string

// May be one of ACPD, VALD, MATD, AUTD, INVD, UMAC, STLE, STLM, SSPD, PCAN, PSTL, PFST, SMLR, RMLR, SRBL, AVLB, SRML
type PendingStatus4Code string

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type Priority1Choice struct {
	Cd    Priority5Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

// May be one of HIGH, LOWW, NORM, URGT
type Priority5Code string

type ProprietaryStatusJustification2 struct {
	PrtryStsRsn Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PrtryStsRsn"`
	Rsn         Max256Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Rsn"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

type QueueTransactionIdentification1 struct {
	QId    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 QId"`
	PosInQ Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PosInQ"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Enqry"`
	Prtry   GenericIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Prtry"`
}

type ReturnTransactionV08 struct {
	MsgHdr      MessageHeader8                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 MsgHdr"`
	RptOrErr    TransactionReportOrError4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 RptOrErr"`
	SplmtryData []SupplementaryData1            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 SplmtryData,omitempty"`
}

type SecuritiesTransactionReferences1 struct {
	AcctOwnrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AcctOwnrTxId,omitempty"`
	AcctSvcrTxId      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AcctSvcrTxId,omitempty"`
	MktInfrstrctrTxId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 MktInfrstrctrTxId,omitempty"`
	PrcgId            Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PrcgId,omitempty"`
}

type ShortPaymentIdentification2 struct {
	TxId          Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 TxId"`
	IntrBkSttlmDt ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 IntrBkSttlmDt"`
	InstgAgt      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 InstgAgt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of SUBY, SUBS
type SuspendedStatusReason1Code string

type System2 struct {
	SysId  MarketInfrastructureIdentification1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 SysId,omitempty"`
	MmbId  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 MmbId,omitempty"`
	Ctry   CountryCode                                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Ctry,omitempty"`
	AcctId AccountIdentification4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AcctId,omitempty"`
}

type Transaction66 struct {
	PmtTo        System2                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PmtTo,omitempty"`
	PmtFr        System2                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PmtFr,omitempty"`
	CdtDbtInd    CreditDebitCode                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 CdtDbtInd,omitempty"`
	Pmt          PaymentInstruction32             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Pmt,omitempty"`
	AcctNtry     CashAccountAndEntry3             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 AcctNtry,omitempty"`
	SctiesTxRefs SecuritiesTransactionReferences1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 SctiesTxRefs,omitempty"`
}

type TransactionOrError4Choice struct {
	Tx     Transaction66    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 Tx"`
	BizErr []ErrorHandling5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 BizErr"`
}

type TransactionReport5 struct {
	PmtId   PaymentIdentification6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PmtId"`
	TxOrErr TransactionOrError4Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 TxOrErr"`
}

type TransactionReportOrError4Choice struct {
	BizRpt  Transactions8    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 BizRpt"`
	OprlErr []ErrorHandling5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 OprlErr"`
}

type Transactions8 struct {
	PmtCmonInf PaymentCommon4              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 PmtCmonInf,omitempty"`
	TxsSummry  NumberAndSumOfTransactions2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 TxsSummry,omitempty"`
	TxRpt      []TransactionReport5        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 TxRpt"`
}

// Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

// May be one of CMIS, DDAT, DELN, DEPT, DMON, DDEA, DQUA, CADE, SETR, DSEC, VASU, DTRA, RSPR, REPO, CLAT, RERT, REPA, REPP, PHYS, IIND, FRAP, PLCE, PODU, FORF, REGD, RTGS, ICAG, CPCA, CHAR, IEXE, NCRR, NMAS, SAFE, DTRD, LATE, TERM, ICUS
type UnmatchedStatusReason1Code string

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
