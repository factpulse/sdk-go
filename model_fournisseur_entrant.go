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

// checks if the FournisseurEntrant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FournisseurEntrant{}

// FournisseurEntrant Fournisseur extrait d'une facture entrante.  Contrairement au modèle Fournisseur de models.py, ce modèle n'a pas de champ id_fournisseur car cette information n'est pas disponible dans les XML Factur-X/CII/UBL.
type FournisseurEntrant struct {
	// Raison sociale du fournisseur (BT-27)
	Nom string `json:"nom"`
	Siren NullableString `json:"siren,omitempty"`
	Siret NullableString `json:"siret,omitempty"`
	NumeroTvaIntra NullableString `json:"numero_tva_intra,omitempty"`
	AdressePostale NullableAdressePostale `json:"adresse_postale,omitempty"`
	AdresseElectronique NullableAdresseElectronique `json:"adresse_electronique,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Telephone NullableString `json:"telephone,omitempty"`
}

type _FournisseurEntrant FournisseurEntrant

// NewFournisseurEntrant instantiates a new FournisseurEntrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFournisseurEntrant(nom string) *FournisseurEntrant {
	this := FournisseurEntrant{}
	this.Nom = nom
	return &this
}

// NewFournisseurEntrantWithDefaults instantiates a new FournisseurEntrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFournisseurEntrantWithDefaults() *FournisseurEntrant {
	this := FournisseurEntrant{}
	return &this
}

// GetNom returns the Nom field value
func (o *FournisseurEntrant) GetNom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nom
}

// GetNomOk returns a tuple with the Nom field value
// and a boolean to check if the value has been set.
func (o *FournisseurEntrant) GetNomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nom, true
}

// SetNom sets field value
func (o *FournisseurEntrant) SetNom(v string) {
	o.Nom = v
}

// GetSiren returns the Siren field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FournisseurEntrant) GetSiren() string {
	if o == nil || IsNil(o.Siren.Get()) {
		var ret string
		return ret
	}
	return *o.Siren.Get()
}

// GetSirenOk returns a tuple with the Siren field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FournisseurEntrant) GetSirenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Siren.Get(), o.Siren.IsSet()
}

// HasSiren returns a boolean if a field has been set.
func (o *FournisseurEntrant) HasSiren() bool {
	if o != nil && o.Siren.IsSet() {
		return true
	}

	return false
}

// SetSiren gets a reference to the given NullableString and assigns it to the Siren field.
func (o *FournisseurEntrant) SetSiren(v string) {
	o.Siren.Set(&v)
}
// SetSirenNil sets the value for Siren to be an explicit nil
func (o *FournisseurEntrant) SetSirenNil() {
	o.Siren.Set(nil)
}

// UnsetSiren ensures that no value is present for Siren, not even an explicit nil
func (o *FournisseurEntrant) UnsetSiren() {
	o.Siren.Unset()
}

// GetSiret returns the Siret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FournisseurEntrant) GetSiret() string {
	if o == nil || IsNil(o.Siret.Get()) {
		var ret string
		return ret
	}
	return *o.Siret.Get()
}

// GetSiretOk returns a tuple with the Siret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FournisseurEntrant) GetSiretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Siret.Get(), o.Siret.IsSet()
}

// HasSiret returns a boolean if a field has been set.
func (o *FournisseurEntrant) HasSiret() bool {
	if o != nil && o.Siret.IsSet() {
		return true
	}

	return false
}

// SetSiret gets a reference to the given NullableString and assigns it to the Siret field.
func (o *FournisseurEntrant) SetSiret(v string) {
	o.Siret.Set(&v)
}
// SetSiretNil sets the value for Siret to be an explicit nil
func (o *FournisseurEntrant) SetSiretNil() {
	o.Siret.Set(nil)
}

// UnsetSiret ensures that no value is present for Siret, not even an explicit nil
func (o *FournisseurEntrant) UnsetSiret() {
	o.Siret.Unset()
}

// GetNumeroTvaIntra returns the NumeroTvaIntra field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FournisseurEntrant) GetNumeroTvaIntra() string {
	if o == nil || IsNil(o.NumeroTvaIntra.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroTvaIntra.Get()
}

// GetNumeroTvaIntraOk returns a tuple with the NumeroTvaIntra field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FournisseurEntrant) GetNumeroTvaIntraOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroTvaIntra.Get(), o.NumeroTvaIntra.IsSet()
}

// HasNumeroTvaIntra returns a boolean if a field has been set.
func (o *FournisseurEntrant) HasNumeroTvaIntra() bool {
	if o != nil && o.NumeroTvaIntra.IsSet() {
		return true
	}

	return false
}

// SetNumeroTvaIntra gets a reference to the given NullableString and assigns it to the NumeroTvaIntra field.
func (o *FournisseurEntrant) SetNumeroTvaIntra(v string) {
	o.NumeroTvaIntra.Set(&v)
}
// SetNumeroTvaIntraNil sets the value for NumeroTvaIntra to be an explicit nil
func (o *FournisseurEntrant) SetNumeroTvaIntraNil() {
	o.NumeroTvaIntra.Set(nil)
}

// UnsetNumeroTvaIntra ensures that no value is present for NumeroTvaIntra, not even an explicit nil
func (o *FournisseurEntrant) UnsetNumeroTvaIntra() {
	o.NumeroTvaIntra.Unset()
}

// GetAdressePostale returns the AdressePostale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FournisseurEntrant) GetAdressePostale() AdressePostale {
	if o == nil || IsNil(o.AdressePostale.Get()) {
		var ret AdressePostale
		return ret
	}
	return *o.AdressePostale.Get()
}

// GetAdressePostaleOk returns a tuple with the AdressePostale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FournisseurEntrant) GetAdressePostaleOk() (*AdressePostale, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdressePostale.Get(), o.AdressePostale.IsSet()
}

// HasAdressePostale returns a boolean if a field has been set.
func (o *FournisseurEntrant) HasAdressePostale() bool {
	if o != nil && o.AdressePostale.IsSet() {
		return true
	}

	return false
}

// SetAdressePostale gets a reference to the given NullableAdressePostale and assigns it to the AdressePostale field.
func (o *FournisseurEntrant) SetAdressePostale(v AdressePostale) {
	o.AdressePostale.Set(&v)
}
// SetAdressePostaleNil sets the value for AdressePostale to be an explicit nil
func (o *FournisseurEntrant) SetAdressePostaleNil() {
	o.AdressePostale.Set(nil)
}

// UnsetAdressePostale ensures that no value is present for AdressePostale, not even an explicit nil
func (o *FournisseurEntrant) UnsetAdressePostale() {
	o.AdressePostale.Unset()
}

// GetAdresseElectronique returns the AdresseElectronique field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FournisseurEntrant) GetAdresseElectronique() AdresseElectronique {
	if o == nil || IsNil(o.AdresseElectronique.Get()) {
		var ret AdresseElectronique
		return ret
	}
	return *o.AdresseElectronique.Get()
}

// GetAdresseElectroniqueOk returns a tuple with the AdresseElectronique field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FournisseurEntrant) GetAdresseElectroniqueOk() (*AdresseElectronique, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdresseElectronique.Get(), o.AdresseElectronique.IsSet()
}

// HasAdresseElectronique returns a boolean if a field has been set.
func (o *FournisseurEntrant) HasAdresseElectronique() bool {
	if o != nil && o.AdresseElectronique.IsSet() {
		return true
	}

	return false
}

// SetAdresseElectronique gets a reference to the given NullableAdresseElectronique and assigns it to the AdresseElectronique field.
func (o *FournisseurEntrant) SetAdresseElectronique(v AdresseElectronique) {
	o.AdresseElectronique.Set(&v)
}
// SetAdresseElectroniqueNil sets the value for AdresseElectronique to be an explicit nil
func (o *FournisseurEntrant) SetAdresseElectroniqueNil() {
	o.AdresseElectronique.Set(nil)
}

// UnsetAdresseElectronique ensures that no value is present for AdresseElectronique, not even an explicit nil
func (o *FournisseurEntrant) UnsetAdresseElectronique() {
	o.AdresseElectronique.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FournisseurEntrant) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FournisseurEntrant) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *FournisseurEntrant) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *FournisseurEntrant) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *FournisseurEntrant) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *FournisseurEntrant) UnsetEmail() {
	o.Email.Unset()
}

// GetTelephone returns the Telephone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FournisseurEntrant) GetTelephone() string {
	if o == nil || IsNil(o.Telephone.Get()) {
		var ret string
		return ret
	}
	return *o.Telephone.Get()
}

// GetTelephoneOk returns a tuple with the Telephone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FournisseurEntrant) GetTelephoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Telephone.Get(), o.Telephone.IsSet()
}

// HasTelephone returns a boolean if a field has been set.
func (o *FournisseurEntrant) HasTelephone() bool {
	if o != nil && o.Telephone.IsSet() {
		return true
	}

	return false
}

// SetTelephone gets a reference to the given NullableString and assigns it to the Telephone field.
func (o *FournisseurEntrant) SetTelephone(v string) {
	o.Telephone.Set(&v)
}
// SetTelephoneNil sets the value for Telephone to be an explicit nil
func (o *FournisseurEntrant) SetTelephoneNil() {
	o.Telephone.Set(nil)
}

// UnsetTelephone ensures that no value is present for Telephone, not even an explicit nil
func (o *FournisseurEntrant) UnsetTelephone() {
	o.Telephone.Unset()
}

func (o FournisseurEntrant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FournisseurEntrant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nom"] = o.Nom
	if o.Siren.IsSet() {
		toSerialize["siren"] = o.Siren.Get()
	}
	if o.Siret.IsSet() {
		toSerialize["siret"] = o.Siret.Get()
	}
	if o.NumeroTvaIntra.IsSet() {
		toSerialize["numero_tva_intra"] = o.NumeroTvaIntra.Get()
	}
	if o.AdressePostale.IsSet() {
		toSerialize["adresse_postale"] = o.AdressePostale.Get()
	}
	if o.AdresseElectronique.IsSet() {
		toSerialize["adresse_electronique"] = o.AdresseElectronique.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Telephone.IsSet() {
		toSerialize["telephone"] = o.Telephone.Get()
	}
	return toSerialize, nil
}

func (o *FournisseurEntrant) UnmarshalJSON(data []byte) (err error) {
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

	varFournisseurEntrant := _FournisseurEntrant{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFournisseurEntrant)

	if err != nil {
		return err
	}

	*o = FournisseurEntrant(varFournisseurEntrant)

	return err
}

type NullableFournisseurEntrant struct {
	value *FournisseurEntrant
	isSet bool
}

func (v NullableFournisseurEntrant) Get() *FournisseurEntrant {
	return v.value
}

func (v *NullableFournisseurEntrant) Set(val *FournisseurEntrant) {
	v.value = val
	v.isSet = true
}

func (v NullableFournisseurEntrant) IsSet() bool {
	return v.isSet
}

func (v *NullableFournisseurEntrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFournisseurEntrant(val *FournisseurEntrant) *NullableFournisseurEntrant {
	return &NullableFournisseurEntrant{value: val, isSet: true}
}

func (v NullableFournisseurEntrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFournisseurEntrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


