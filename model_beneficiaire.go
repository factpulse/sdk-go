/*
API REST FactPulse

 API REST pour la facturation électronique en France : Factur-X, AFNOR PDP/PA, signatures électroniques.  ## 🎯 Fonctionnalités principales  ### 📄 Génération de factures Factur-X - **Formats** : XML seul ou PDF/A-3 avec XML embarqué - **Profils** : MINIMUM, BASIC, EN16931, EXTENDED - **Normes** : EN 16931 (directive UE 2014/55), ISO 19005-3 (PDF/A-3), CII (UN/CEFACT) - **🆕 Format simplifié** : Génération à partir de SIRET + auto-enrichissement (API Chorus Pro + Recherche Entreprises)  ### ✅ Validation et conformité - **Validation XML** : Schematron (45 à 210+ règles selon profil) - **Validation PDF** : PDF/A-3, métadonnées XMP Factur-X, signatures électroniques - **VeraPDF** : Validation stricte PDF/A (146+ règles ISO 19005-3) - **Traitement asynchrone** : Support Celery pour validations lourdes (VeraPDF)  ### 📡 Intégration AFNOR PDP/PA (XP Z12-013) - **Soumission de flux** : Envoi de factures vers Plateformes de Dématérialisation Partenaires - **Recherche de flux** : Consultation des factures soumises - **Téléchargement** : Récupération des PDF/A-3 avec XML - **Directory Service** : Recherche d'entreprises (SIREN/SIRET) - **Multi-client** : Support de plusieurs configs PDP par utilisateur (stored credentials ou zero-storage)  ### ✍️ Signature électronique PDF - **Standards** : PAdES-B-B, PAdES-B-T (horodatage RFC 3161), PAdES-B-LT (archivage long terme) - **Niveaux eIDAS** : SES (auto-signé), AdES (CA commerciale), QES (PSCO) - **Validation** : Vérification intégrité cryptographique et certificats - **Génération de certificats** : Certificats X.509 auto-signés pour tests  ### 🔄 Traitement asynchrone - **Celery** : Génération, validation et signature asynchrones - **Polling** : Suivi d'état via `/taches/{id_tache}/statut` - **Pas de timeout** : Idéal pour gros fichiers ou validations lourdes  ## 🔒 Authentification  Toutes les requêtes nécessitent un **token JWT** dans le header Authorization : ``` Authorization: Bearer YOUR_JWT_TOKEN ```  ### Comment obtenir un token JWT ?  #### 🔑 Méthode 1 : API `/api/token/` (Recommandée)  **URL :** `https://www.factpulse.fr/api/token/`  Cette méthode est **recommandée** pour l'intégration dans vos applications et workflows CI/CD.  **Prérequis :** Avoir défini un mot de passe sur votre compte  **Pour les utilisateurs inscrits via email/password :** - Vous avez déjà un mot de passe, utilisez-le directement  **Pour les utilisateurs inscrits via OAuth (Google/GitHub) :** - Vous devez d'abord définir un mot de passe sur : https://www.factpulse.fr/accounts/password/set/ - Une fois le mot de passe créé, vous pourrez utiliser l'API  **Exemple de requête :** ```bash curl -X POST https://www.factpulse.fr/api/token/ \\   -H \"Content-Type: application/json\" \\   -d '{     \"username\": \"votre_email@example.com\",     \"password\": \"votre_mot_de_passe\"   }' ```  **Paramètre optionnel `client_uid` :**  Pour sélectionner les credentials d'un client spécifique (PA/PDP, Chorus Pro, certificats de signature), ajoutez `client_uid` :  ```bash curl -X POST https://www.factpulse.fr/api/token/ \\   -H \"Content-Type: application/json\" \\   -d '{     \"username\": \"votre_email@example.com\",     \"password\": \"votre_mot_de_passe\",     \"client_uid\": \"550e8400-e29b-41d4-a716-446655440000\"   }' ```  Le `client_uid` sera inclus dans le JWT et permettra à l'API d'utiliser automatiquement : - Les credentials AFNOR/PDP configurés pour ce client - Les credentials Chorus Pro configurés pour ce client - Les certificats de signature électronique configurés pour ce client  **Réponse :** ```json {   \"access\": \"eyJ0eXAiOiJKV1QiLCJhbGc...\",  // Token d'accès (validité: 30 min)   \"refresh\": \"eyJ0eXAiOiJKV1QiLCJhbGc...\"  // Token de rafraîchissement (validité: 7 jours) } ```  **Avantages :** - ✅ Automatisation complète (CI/CD, scripts) - ✅ Gestion programmatique des tokens - ✅ Support du refresh token pour renouveler automatiquement l'accès - ✅ Intégration facile dans n'importe quel langage/outil  #### 🖥️ Méthode 2 : Génération via Dashboard (Alternative)  **URL :** https://www.factpulse.fr/dashboard/  Cette méthode convient pour des tests rapides ou une utilisation occasionnelle via l'interface graphique.  **Fonctionnement :** - Connectez-vous au dashboard - Utilisez les boutons \"Generate Test Token\" ou \"Generate Production Token\" - Fonctionne pour **tous** les utilisateurs (OAuth et email/password), sans nécessiter de mot de passe  **Types de tokens :** - **Token Test** : Validité 24h, quota 1000 appels/jour (gratuit) - **Token Production** : Validité 7 jours, quota selon votre forfait  **Avantages :** - ✅ Rapide pour tester l'API - ✅ Aucun mot de passe requis - ✅ Interface visuelle simple  **Inconvénients :** - ❌ Nécessite une action manuelle - ❌ Pas de refresh token - ❌ Moins adapté pour l'automatisation  ### 📚 Documentation complète  Pour plus d'informations sur l'authentification et l'utilisation de l'API : https://www.factpulse.fr/documentation-api/     

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package factpulse

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Beneficiaire type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Beneficiaire{}

// Beneficiaire Informations sur le bénéficiaire du paiement (BG-10 / PayeeTradeParty).  Le bénéficiaire est la partie qui reçoit le paiement. Ce bloc est utilisé uniquement si le bénéficiaire est différent du vendeur (fournisseur).  **Cas d'usage principal** : Affacturage (factoring) Quand une facture est affacturée, le factor (société d'affacturage) devient le bénéficiaire du paiement à la place du fournisseur.  **Business Terms (EN16931)** : - BT-59 : Nom du bénéficiaire (obligatoire) - BT-60 : Identifiant du bénéficiaire (SIRET avec schemeID 0009) - BT-61 : Identifiant légal du bénéficiaire (SIREN avec schemeID 0002)  **Référence** : docs/guide_affacturage.md
type Beneficiaire struct {
	// Nom du bénéficiaire (BT-59). Obligatoire.
	Nom string `json:"nom"`
	Siret NullableString `json:"siret,omitempty" validate:"regexp=^\\\\d{14}$"`
	Siren NullableString `json:"siren,omitempty" validate:"regexp=^\\\\d{9}$"`
	AdresseElectronique NullableAdresseElectronique `json:"adresseElectronique,omitempty"`
	Iban NullableString `json:"iban,omitempty"`
	Bic NullableString `json:"bic,omitempty"`
}

type _Beneficiaire Beneficiaire

// NewBeneficiaire instantiates a new Beneficiaire object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeneficiaire(nom string) *Beneficiaire {
	this := Beneficiaire{}
	this.Nom = nom
	return &this
}

// NewBeneficiaireWithDefaults instantiates a new Beneficiaire object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeneficiaireWithDefaults() *Beneficiaire {
	this := Beneficiaire{}
	return &this
}

// GetNom returns the Nom field value
func (o *Beneficiaire) GetNom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nom
}

// GetNomOk returns a tuple with the Nom field value
// and a boolean to check if the value has been set.
func (o *Beneficiaire) GetNomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nom, true
}

// SetNom sets field value
func (o *Beneficiaire) SetNom(v string) {
	o.Nom = v
}

// GetSiret returns the Siret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Beneficiaire) GetSiret() string {
	if o == nil || IsNil(o.Siret.Get()) {
		var ret string
		return ret
	}
	return *o.Siret.Get()
}

// GetSiretOk returns a tuple with the Siret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Beneficiaire) GetSiretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Siret.Get(), o.Siret.IsSet()
}

// HasSiret returns a boolean if a field has been set.
func (o *Beneficiaire) HasSiret() bool {
	if o != nil && o.Siret.IsSet() {
		return true
	}

	return false
}

// SetSiret gets a reference to the given NullableString and assigns it to the Siret field.
func (o *Beneficiaire) SetSiret(v string) {
	o.Siret.Set(&v)
}
// SetSiretNil sets the value for Siret to be an explicit nil
func (o *Beneficiaire) SetSiretNil() {
	o.Siret.Set(nil)
}

// UnsetSiret ensures that no value is present for Siret, not even an explicit nil
func (o *Beneficiaire) UnsetSiret() {
	o.Siret.Unset()
}

// GetSiren returns the Siren field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Beneficiaire) GetSiren() string {
	if o == nil || IsNil(o.Siren.Get()) {
		var ret string
		return ret
	}
	return *o.Siren.Get()
}

// GetSirenOk returns a tuple with the Siren field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Beneficiaire) GetSirenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Siren.Get(), o.Siren.IsSet()
}

// HasSiren returns a boolean if a field has been set.
func (o *Beneficiaire) HasSiren() bool {
	if o != nil && o.Siren.IsSet() {
		return true
	}

	return false
}

// SetSiren gets a reference to the given NullableString and assigns it to the Siren field.
func (o *Beneficiaire) SetSiren(v string) {
	o.Siren.Set(&v)
}
// SetSirenNil sets the value for Siren to be an explicit nil
func (o *Beneficiaire) SetSirenNil() {
	o.Siren.Set(nil)
}

// UnsetSiren ensures that no value is present for Siren, not even an explicit nil
func (o *Beneficiaire) UnsetSiren() {
	o.Siren.Unset()
}

// GetAdresseElectronique returns the AdresseElectronique field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Beneficiaire) GetAdresseElectronique() AdresseElectronique {
	if o == nil || IsNil(o.AdresseElectronique.Get()) {
		var ret AdresseElectronique
		return ret
	}
	return *o.AdresseElectronique.Get()
}

// GetAdresseElectroniqueOk returns a tuple with the AdresseElectronique field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Beneficiaire) GetAdresseElectroniqueOk() (*AdresseElectronique, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdresseElectronique.Get(), o.AdresseElectronique.IsSet()
}

// HasAdresseElectronique returns a boolean if a field has been set.
func (o *Beneficiaire) HasAdresseElectronique() bool {
	if o != nil && o.AdresseElectronique.IsSet() {
		return true
	}

	return false
}

// SetAdresseElectronique gets a reference to the given NullableAdresseElectronique and assigns it to the AdresseElectronique field.
func (o *Beneficiaire) SetAdresseElectronique(v AdresseElectronique) {
	o.AdresseElectronique.Set(&v)
}
// SetAdresseElectroniqueNil sets the value for AdresseElectronique to be an explicit nil
func (o *Beneficiaire) SetAdresseElectroniqueNil() {
	o.AdresseElectronique.Set(nil)
}

// UnsetAdresseElectronique ensures that no value is present for AdresseElectronique, not even an explicit nil
func (o *Beneficiaire) UnsetAdresseElectronique() {
	o.AdresseElectronique.Unset()
}

// GetIban returns the Iban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Beneficiaire) GetIban() string {
	if o == nil || IsNil(o.Iban.Get()) {
		var ret string
		return ret
	}
	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Beneficiaire) GetIbanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// HasIban returns a boolean if a field has been set.
func (o *Beneficiaire) HasIban() bool {
	if o != nil && o.Iban.IsSet() {
		return true
	}

	return false
}

// SetIban gets a reference to the given NullableString and assigns it to the Iban field.
func (o *Beneficiaire) SetIban(v string) {
	o.Iban.Set(&v)
}
// SetIbanNil sets the value for Iban to be an explicit nil
func (o *Beneficiaire) SetIbanNil() {
	o.Iban.Set(nil)
}

// UnsetIban ensures that no value is present for Iban, not even an explicit nil
func (o *Beneficiaire) UnsetIban() {
	o.Iban.Unset()
}

// GetBic returns the Bic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Beneficiaire) GetBic() string {
	if o == nil || IsNil(o.Bic.Get()) {
		var ret string
		return ret
	}
	return *o.Bic.Get()
}

// GetBicOk returns a tuple with the Bic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Beneficiaire) GetBicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bic.Get(), o.Bic.IsSet()
}

// HasBic returns a boolean if a field has been set.
func (o *Beneficiaire) HasBic() bool {
	if o != nil && o.Bic.IsSet() {
		return true
	}

	return false
}

// SetBic gets a reference to the given NullableString and assigns it to the Bic field.
func (o *Beneficiaire) SetBic(v string) {
	o.Bic.Set(&v)
}
// SetBicNil sets the value for Bic to be an explicit nil
func (o *Beneficiaire) SetBicNil() {
	o.Bic.Set(nil)
}

// UnsetBic ensures that no value is present for Bic, not even an explicit nil
func (o *Beneficiaire) UnsetBic() {
	o.Bic.Unset()
}

func (o Beneficiaire) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Beneficiaire) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nom"] = o.Nom
	if o.Siret.IsSet() {
		toSerialize["siret"] = o.Siret.Get()
	}
	if o.Siren.IsSet() {
		toSerialize["siren"] = o.Siren.Get()
	}
	if o.AdresseElectronique.IsSet() {
		toSerialize["adresseElectronique"] = o.AdresseElectronique.Get()
	}
	if o.Iban.IsSet() {
		toSerialize["iban"] = o.Iban.Get()
	}
	if o.Bic.IsSet() {
		toSerialize["bic"] = o.Bic.Get()
	}
	return toSerialize, nil
}

func (o *Beneficiaire) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nom",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBeneficiaire := _Beneficiaire{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBeneficiaire)

	if err != nil {
		return err
	}

	*o = Beneficiaire(varBeneficiaire)

	return err
}

type NullableBeneficiaire struct {
	value *Beneficiaire
	isSet bool
}

func (v NullableBeneficiaire) Get() *Beneficiaire {
	return v.value
}

func (v *NullableBeneficiaire) Set(val *Beneficiaire) {
	v.value = val
	v.isSet = true
}

func (v NullableBeneficiaire) IsSet() bool {
	return v.isSet
}

func (v *NullableBeneficiaire) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeneficiaire(val *Beneficiaire) *NullableBeneficiaire {
	return &NullableBeneficiaire{value: val, isSet: true}
}

func (v NullableBeneficiaire) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeneficiaire) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


