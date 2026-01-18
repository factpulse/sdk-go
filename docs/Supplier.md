# Supplier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElectronicAddress** | Pointer to [**NullableElectronicAddress**](ElectronicAddress.md) |  | [optional] 
**SupplierId** | **int32** |  | 
**PrivateId** | Pointer to **NullableString** |  | [optional] 
**SupplierBankAccountCode** | Pointer to **NullableInt32** |  | [optional] 
**SupplierServiceId** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TradingBusinessName** | Pointer to **NullableString** |  | [optional] 
**LegalDescription** | Pointer to **NullableString** |  | [optional] 
**Siren** | Pointer to **NullableString** |  | [optional] 
**Siret** | Pointer to **NullableString** |  | [optional] 
**VatNumber** | Pointer to **NullableString** |  | [optional] 
**Iban** | Pointer to **NullableString** |  | [optional] 
**Bic** | Pointer to **NullableString** |  | [optional] 
**BankAccountName** | Pointer to **NullableString** |  | [optional] 
**ProprietaryId** | Pointer to **NullableString** |  | [optional] 
**PostalAddress** | Pointer to [**NullablePostalAddress**](PostalAddress.md) |  | [optional] 
**Contact** | Pointer to [**NullableContact**](Contact.md) |  | [optional] 
**GlobalIds** | Pointer to [**[]ElectronicAddress**](ElectronicAddress.md) |  | [optional] 

## Methods

### NewSupplier

`func NewSupplier(supplierId int32, ) *Supplier`

NewSupplier instantiates a new Supplier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierWithDefaults

`func NewSupplierWithDefaults() *Supplier`

NewSupplierWithDefaults instantiates a new Supplier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElectronicAddress

`func (o *Supplier) GetElectronicAddress() ElectronicAddress`

GetElectronicAddress returns the ElectronicAddress field if non-nil, zero value otherwise.

### GetElectronicAddressOk

`func (o *Supplier) GetElectronicAddressOk() (*ElectronicAddress, bool)`

GetElectronicAddressOk returns a tuple with the ElectronicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectronicAddress

`func (o *Supplier) SetElectronicAddress(v ElectronicAddress)`

SetElectronicAddress sets ElectronicAddress field to given value.

### HasElectronicAddress

`func (o *Supplier) HasElectronicAddress() bool`

HasElectronicAddress returns a boolean if a field has been set.

### SetElectronicAddressNil

`func (o *Supplier) SetElectronicAddressNil(b bool)`

 SetElectronicAddressNil sets the value for ElectronicAddress to be an explicit nil

### UnsetElectronicAddress
`func (o *Supplier) UnsetElectronicAddress()`

UnsetElectronicAddress ensures that no value is present for ElectronicAddress, not even an explicit nil
### GetSupplierId

`func (o *Supplier) GetSupplierId() int32`

GetSupplierId returns the SupplierId field if non-nil, zero value otherwise.

### GetSupplierIdOk

`func (o *Supplier) GetSupplierIdOk() (*int32, bool)`

GetSupplierIdOk returns a tuple with the SupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierId

`func (o *Supplier) SetSupplierId(v int32)`

SetSupplierId sets SupplierId field to given value.


### GetPrivateId

`func (o *Supplier) GetPrivateId() string`

GetPrivateId returns the PrivateId field if non-nil, zero value otherwise.

### GetPrivateIdOk

`func (o *Supplier) GetPrivateIdOk() (*string, bool)`

GetPrivateIdOk returns a tuple with the PrivateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateId

`func (o *Supplier) SetPrivateId(v string)`

SetPrivateId sets PrivateId field to given value.

### HasPrivateId

`func (o *Supplier) HasPrivateId() bool`

HasPrivateId returns a boolean if a field has been set.

### SetPrivateIdNil

`func (o *Supplier) SetPrivateIdNil(b bool)`

 SetPrivateIdNil sets the value for PrivateId to be an explicit nil

### UnsetPrivateId
`func (o *Supplier) UnsetPrivateId()`

UnsetPrivateId ensures that no value is present for PrivateId, not even an explicit nil
### GetSupplierBankAccountCode

`func (o *Supplier) GetSupplierBankAccountCode() int32`

GetSupplierBankAccountCode returns the SupplierBankAccountCode field if non-nil, zero value otherwise.

### GetSupplierBankAccountCodeOk

`func (o *Supplier) GetSupplierBankAccountCodeOk() (*int32, bool)`

GetSupplierBankAccountCodeOk returns a tuple with the SupplierBankAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierBankAccountCode

`func (o *Supplier) SetSupplierBankAccountCode(v int32)`

SetSupplierBankAccountCode sets SupplierBankAccountCode field to given value.

### HasSupplierBankAccountCode

`func (o *Supplier) HasSupplierBankAccountCode() bool`

HasSupplierBankAccountCode returns a boolean if a field has been set.

### SetSupplierBankAccountCodeNil

`func (o *Supplier) SetSupplierBankAccountCodeNil(b bool)`

 SetSupplierBankAccountCodeNil sets the value for SupplierBankAccountCode to be an explicit nil

### UnsetSupplierBankAccountCode
`func (o *Supplier) UnsetSupplierBankAccountCode()`

UnsetSupplierBankAccountCode ensures that no value is present for SupplierBankAccountCode, not even an explicit nil
### GetSupplierServiceId

`func (o *Supplier) GetSupplierServiceId() int32`

GetSupplierServiceId returns the SupplierServiceId field if non-nil, zero value otherwise.

### GetSupplierServiceIdOk

`func (o *Supplier) GetSupplierServiceIdOk() (*int32, bool)`

GetSupplierServiceIdOk returns a tuple with the SupplierServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierServiceId

`func (o *Supplier) SetSupplierServiceId(v int32)`

SetSupplierServiceId sets SupplierServiceId field to given value.

### HasSupplierServiceId

`func (o *Supplier) HasSupplierServiceId() bool`

HasSupplierServiceId returns a boolean if a field has been set.

### SetSupplierServiceIdNil

`func (o *Supplier) SetSupplierServiceIdNil(b bool)`

 SetSupplierServiceIdNil sets the value for SupplierServiceId to be an explicit nil

### UnsetSupplierServiceId
`func (o *Supplier) UnsetSupplierServiceId()`

UnsetSupplierServiceId ensures that no value is present for SupplierServiceId, not even an explicit nil
### GetName

`func (o *Supplier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Supplier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Supplier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Supplier) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Supplier) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Supplier) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTradingBusinessName

`func (o *Supplier) GetTradingBusinessName() string`

GetTradingBusinessName returns the TradingBusinessName field if non-nil, zero value otherwise.

### GetTradingBusinessNameOk

`func (o *Supplier) GetTradingBusinessNameOk() (*string, bool)`

GetTradingBusinessNameOk returns a tuple with the TradingBusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingBusinessName

`func (o *Supplier) SetTradingBusinessName(v string)`

SetTradingBusinessName sets TradingBusinessName field to given value.

### HasTradingBusinessName

`func (o *Supplier) HasTradingBusinessName() bool`

HasTradingBusinessName returns a boolean if a field has been set.

### SetTradingBusinessNameNil

`func (o *Supplier) SetTradingBusinessNameNil(b bool)`

 SetTradingBusinessNameNil sets the value for TradingBusinessName to be an explicit nil

### UnsetTradingBusinessName
`func (o *Supplier) UnsetTradingBusinessName()`

UnsetTradingBusinessName ensures that no value is present for TradingBusinessName, not even an explicit nil
### GetLegalDescription

`func (o *Supplier) GetLegalDescription() string`

GetLegalDescription returns the LegalDescription field if non-nil, zero value otherwise.

### GetLegalDescriptionOk

`func (o *Supplier) GetLegalDescriptionOk() (*string, bool)`

GetLegalDescriptionOk returns a tuple with the LegalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalDescription

`func (o *Supplier) SetLegalDescription(v string)`

SetLegalDescription sets LegalDescription field to given value.

### HasLegalDescription

`func (o *Supplier) HasLegalDescription() bool`

HasLegalDescription returns a boolean if a field has been set.

### SetLegalDescriptionNil

`func (o *Supplier) SetLegalDescriptionNil(b bool)`

 SetLegalDescriptionNil sets the value for LegalDescription to be an explicit nil

### UnsetLegalDescription
`func (o *Supplier) UnsetLegalDescription()`

UnsetLegalDescription ensures that no value is present for LegalDescription, not even an explicit nil
### GetSiren

`func (o *Supplier) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *Supplier) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *Supplier) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *Supplier) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### SetSirenNil

`func (o *Supplier) SetSirenNil(b bool)`

 SetSirenNil sets the value for Siren to be an explicit nil

### UnsetSiren
`func (o *Supplier) UnsetSiren()`

UnsetSiren ensures that no value is present for Siren, not even an explicit nil
### GetSiret

`func (o *Supplier) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *Supplier) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *Supplier) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *Supplier) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### SetSiretNil

`func (o *Supplier) SetSiretNil(b bool)`

 SetSiretNil sets the value for Siret to be an explicit nil

### UnsetSiret
`func (o *Supplier) UnsetSiret()`

UnsetSiret ensures that no value is present for Siret, not even an explicit nil
### GetVatNumber

`func (o *Supplier) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *Supplier) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *Supplier) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *Supplier) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### SetVatNumberNil

`func (o *Supplier) SetVatNumberNil(b bool)`

 SetVatNumberNil sets the value for VatNumber to be an explicit nil

### UnsetVatNumber
`func (o *Supplier) UnsetVatNumber()`

UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
### GetIban

`func (o *Supplier) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *Supplier) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *Supplier) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *Supplier) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *Supplier) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *Supplier) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetBic

`func (o *Supplier) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *Supplier) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *Supplier) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *Supplier) HasBic() bool`

HasBic returns a boolean if a field has been set.

### SetBicNil

`func (o *Supplier) SetBicNil(b bool)`

 SetBicNil sets the value for Bic to be an explicit nil

### UnsetBic
`func (o *Supplier) UnsetBic()`

UnsetBic ensures that no value is present for Bic, not even an explicit nil
### GetBankAccountName

`func (o *Supplier) GetBankAccountName() string`

GetBankAccountName returns the BankAccountName field if non-nil, zero value otherwise.

### GetBankAccountNameOk

`func (o *Supplier) GetBankAccountNameOk() (*string, bool)`

GetBankAccountNameOk returns a tuple with the BankAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountName

`func (o *Supplier) SetBankAccountName(v string)`

SetBankAccountName sets BankAccountName field to given value.

### HasBankAccountName

`func (o *Supplier) HasBankAccountName() bool`

HasBankAccountName returns a boolean if a field has been set.

### SetBankAccountNameNil

`func (o *Supplier) SetBankAccountNameNil(b bool)`

 SetBankAccountNameNil sets the value for BankAccountName to be an explicit nil

### UnsetBankAccountName
`func (o *Supplier) UnsetBankAccountName()`

UnsetBankAccountName ensures that no value is present for BankAccountName, not even an explicit nil
### GetProprietaryId

`func (o *Supplier) GetProprietaryId() string`

GetProprietaryId returns the ProprietaryId field if non-nil, zero value otherwise.

### GetProprietaryIdOk

`func (o *Supplier) GetProprietaryIdOk() (*string, bool)`

GetProprietaryIdOk returns a tuple with the ProprietaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietaryId

`func (o *Supplier) SetProprietaryId(v string)`

SetProprietaryId sets ProprietaryId field to given value.

### HasProprietaryId

`func (o *Supplier) HasProprietaryId() bool`

HasProprietaryId returns a boolean if a field has been set.

### SetProprietaryIdNil

`func (o *Supplier) SetProprietaryIdNil(b bool)`

 SetProprietaryIdNil sets the value for ProprietaryId to be an explicit nil

### UnsetProprietaryId
`func (o *Supplier) UnsetProprietaryId()`

UnsetProprietaryId ensures that no value is present for ProprietaryId, not even an explicit nil
### GetPostalAddress

`func (o *Supplier) GetPostalAddress() PostalAddress`

GetPostalAddress returns the PostalAddress field if non-nil, zero value otherwise.

### GetPostalAddressOk

`func (o *Supplier) GetPostalAddressOk() (*PostalAddress, bool)`

GetPostalAddressOk returns a tuple with the PostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddress

`func (o *Supplier) SetPostalAddress(v PostalAddress)`

SetPostalAddress sets PostalAddress field to given value.

### HasPostalAddress

`func (o *Supplier) HasPostalAddress() bool`

HasPostalAddress returns a boolean if a field has been set.

### SetPostalAddressNil

`func (o *Supplier) SetPostalAddressNil(b bool)`

 SetPostalAddressNil sets the value for PostalAddress to be an explicit nil

### UnsetPostalAddress
`func (o *Supplier) UnsetPostalAddress()`

UnsetPostalAddress ensures that no value is present for PostalAddress, not even an explicit nil
### GetContact

`func (o *Supplier) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Supplier) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Supplier) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Supplier) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *Supplier) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *Supplier) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetGlobalIds

`func (o *Supplier) GetGlobalIds() []ElectronicAddress`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *Supplier) GetGlobalIdsOk() (*[]ElectronicAddress, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *Supplier) SetGlobalIds(v []ElectronicAddress)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *Supplier) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *Supplier) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *Supplier) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


