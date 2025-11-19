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

// checks if the InformationSignatureAPI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InformationSignatureAPI{}

// InformationSignatureAPI Informations sur une signature électronique dans un PDF.
type InformationSignatureAPI struct {
	// Nom du champ de signature dans le PDF
	NomChamp string `json:"nom_champ"`
	Signataire NullableString `json:"signataire,omitempty"`
	DateSignature NullableString `json:"date_signature,omitempty"`
	Raison NullableString `json:"raison,omitempty"`
	Localisation NullableString `json:"localisation,omitempty"`
	EstValide NullableBool `json:"est_valide,omitempty"`
}

type _InformationSignatureAPI InformationSignatureAPI

// NewInformationSignatureAPI instantiates a new InformationSignatureAPI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInformationSignatureAPI(nomChamp string) *InformationSignatureAPI {
	this := InformationSignatureAPI{}
	this.NomChamp = nomChamp
	return &this
}

// NewInformationSignatureAPIWithDefaults instantiates a new InformationSignatureAPI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInformationSignatureAPIWithDefaults() *InformationSignatureAPI {
	this := InformationSignatureAPI{}
	return &this
}

// GetNomChamp returns the NomChamp field value
func (o *InformationSignatureAPI) GetNomChamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NomChamp
}

// GetNomChampOk returns a tuple with the NomChamp field value
// and a boolean to check if the value has been set.
func (o *InformationSignatureAPI) GetNomChampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NomChamp, true
}

// SetNomChamp sets field value
func (o *InformationSignatureAPI) SetNomChamp(v string) {
	o.NomChamp = v
}

// GetSignataire returns the Signataire field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InformationSignatureAPI) GetSignataire() string {
	if o == nil || IsNil(o.Signataire.Get()) {
		var ret string
		return ret
	}
	return *o.Signataire.Get()
}

// GetSignataireOk returns a tuple with the Signataire field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InformationSignatureAPI) GetSignataireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signataire.Get(), o.Signataire.IsSet()
}

// HasSignataire returns a boolean if a field has been set.
func (o *InformationSignatureAPI) HasSignataire() bool {
	if o != nil && o.Signataire.IsSet() {
		return true
	}

	return false
}

// SetSignataire gets a reference to the given NullableString and assigns it to the Signataire field.
func (o *InformationSignatureAPI) SetSignataire(v string) {
	o.Signataire.Set(&v)
}
// SetSignataireNil sets the value for Signataire to be an explicit nil
func (o *InformationSignatureAPI) SetSignataireNil() {
	o.Signataire.Set(nil)
}

// UnsetSignataire ensures that no value is present for Signataire, not even an explicit nil
func (o *InformationSignatureAPI) UnsetSignataire() {
	o.Signataire.Unset()
}

// GetDateSignature returns the DateSignature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InformationSignatureAPI) GetDateSignature() string {
	if o == nil || IsNil(o.DateSignature.Get()) {
		var ret string
		return ret
	}
	return *o.DateSignature.Get()
}

// GetDateSignatureOk returns a tuple with the DateSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InformationSignatureAPI) GetDateSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateSignature.Get(), o.DateSignature.IsSet()
}

// HasDateSignature returns a boolean if a field has been set.
func (o *InformationSignatureAPI) HasDateSignature() bool {
	if o != nil && o.DateSignature.IsSet() {
		return true
	}

	return false
}

// SetDateSignature gets a reference to the given NullableString and assigns it to the DateSignature field.
func (o *InformationSignatureAPI) SetDateSignature(v string) {
	o.DateSignature.Set(&v)
}
// SetDateSignatureNil sets the value for DateSignature to be an explicit nil
func (o *InformationSignatureAPI) SetDateSignatureNil() {
	o.DateSignature.Set(nil)
}

// UnsetDateSignature ensures that no value is present for DateSignature, not even an explicit nil
func (o *InformationSignatureAPI) UnsetDateSignature() {
	o.DateSignature.Unset()
}

// GetRaison returns the Raison field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InformationSignatureAPI) GetRaison() string {
	if o == nil || IsNil(o.Raison.Get()) {
		var ret string
		return ret
	}
	return *o.Raison.Get()
}

// GetRaisonOk returns a tuple with the Raison field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InformationSignatureAPI) GetRaisonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Raison.Get(), o.Raison.IsSet()
}

// HasRaison returns a boolean if a field has been set.
func (o *InformationSignatureAPI) HasRaison() bool {
	if o != nil && o.Raison.IsSet() {
		return true
	}

	return false
}

// SetRaison gets a reference to the given NullableString and assigns it to the Raison field.
func (o *InformationSignatureAPI) SetRaison(v string) {
	o.Raison.Set(&v)
}
// SetRaisonNil sets the value for Raison to be an explicit nil
func (o *InformationSignatureAPI) SetRaisonNil() {
	o.Raison.Set(nil)
}

// UnsetRaison ensures that no value is present for Raison, not even an explicit nil
func (o *InformationSignatureAPI) UnsetRaison() {
	o.Raison.Unset()
}

// GetLocalisation returns the Localisation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InformationSignatureAPI) GetLocalisation() string {
	if o == nil || IsNil(o.Localisation.Get()) {
		var ret string
		return ret
	}
	return *o.Localisation.Get()
}

// GetLocalisationOk returns a tuple with the Localisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InformationSignatureAPI) GetLocalisationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Localisation.Get(), o.Localisation.IsSet()
}

// HasLocalisation returns a boolean if a field has been set.
func (o *InformationSignatureAPI) HasLocalisation() bool {
	if o != nil && o.Localisation.IsSet() {
		return true
	}

	return false
}

// SetLocalisation gets a reference to the given NullableString and assigns it to the Localisation field.
func (o *InformationSignatureAPI) SetLocalisation(v string) {
	o.Localisation.Set(&v)
}
// SetLocalisationNil sets the value for Localisation to be an explicit nil
func (o *InformationSignatureAPI) SetLocalisationNil() {
	o.Localisation.Set(nil)
}

// UnsetLocalisation ensures that no value is present for Localisation, not even an explicit nil
func (o *InformationSignatureAPI) UnsetLocalisation() {
	o.Localisation.Unset()
}

// GetEstValide returns the EstValide field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InformationSignatureAPI) GetEstValide() bool {
	if o == nil || IsNil(o.EstValide.Get()) {
		var ret bool
		return ret
	}
	return *o.EstValide.Get()
}

// GetEstValideOk returns a tuple with the EstValide field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InformationSignatureAPI) GetEstValideOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstValide.Get(), o.EstValide.IsSet()
}

// HasEstValide returns a boolean if a field has been set.
func (o *InformationSignatureAPI) HasEstValide() bool {
	if o != nil && o.EstValide.IsSet() {
		return true
	}

	return false
}

// SetEstValide gets a reference to the given NullableBool and assigns it to the EstValide field.
func (o *InformationSignatureAPI) SetEstValide(v bool) {
	o.EstValide.Set(&v)
}
// SetEstValideNil sets the value for EstValide to be an explicit nil
func (o *InformationSignatureAPI) SetEstValideNil() {
	o.EstValide.Set(nil)
}

// UnsetEstValide ensures that no value is present for EstValide, not even an explicit nil
func (o *InformationSignatureAPI) UnsetEstValide() {
	o.EstValide.Unset()
}

func (o InformationSignatureAPI) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InformationSignatureAPI) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nom_champ"] = o.NomChamp
	if o.Signataire.IsSet() {
		toSerialize["signataire"] = o.Signataire.Get()
	}
	if o.DateSignature.IsSet() {
		toSerialize["date_signature"] = o.DateSignature.Get()
	}
	if o.Raison.IsSet() {
		toSerialize["raison"] = o.Raison.Get()
	}
	if o.Localisation.IsSet() {
		toSerialize["localisation"] = o.Localisation.Get()
	}
	if o.EstValide.IsSet() {
		toSerialize["est_valide"] = o.EstValide.Get()
	}
	return toSerialize, nil
}

func (o *InformationSignatureAPI) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nom_champ",
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

	varInformationSignatureAPI := _InformationSignatureAPI{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInformationSignatureAPI)

	if err != nil {
		return err
	}

	*o = InformationSignatureAPI(varInformationSignatureAPI)

	return err
}

type NullableInformationSignatureAPI struct {
	value *InformationSignatureAPI
	isSet bool
}

func (v NullableInformationSignatureAPI) Get() *InformationSignatureAPI {
	return v.value
}

func (v *NullableInformationSignatureAPI) Set(val *InformationSignatureAPI) {
	v.value = val
	v.isSet = true
}

func (v NullableInformationSignatureAPI) IsSet() bool {
	return v.isSet
}

func (v *NullableInformationSignatureAPI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInformationSignatureAPI(val *InformationSignatureAPI) *NullableInformationSignatureAPI {
	return &NullableInformationSignatureAPI{value: val, isSet: true}
}

func (v NullableInformationSignatureAPI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInformationSignatureAPI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


