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

// checks if the RequeteSoumissionFlux type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequeteSoumissionFlux{}

// RequeteSoumissionFlux Requête pour soumettre une facture à une PDP/PA via AFNOR
type RequeteSoumissionFlux struct {
	// Nom du flux (ex: 'Facture 2025-001')
	NomFlux string `json:"nom_flux"`
	// Syntaxe du flux (CII pour Factur-X)
	SyntaxeFlux *SyntaxeFlux `json:"syntaxe_flux,omitempty"`
	ProfilFlux NullableProfilFlux `json:"profil_flux,omitempty"`
	TrackingId NullableString `json:"tracking_id,omitempty"`
	RequestId NullableString `json:"request_id,omitempty"`
	PdpCredentials NullablePDPCredentials `json:"pdp_credentials,omitempty"`
}

type _RequeteSoumissionFlux RequeteSoumissionFlux

// NewRequeteSoumissionFlux instantiates a new RequeteSoumissionFlux object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequeteSoumissionFlux(nomFlux string) *RequeteSoumissionFlux {
	this := RequeteSoumissionFlux{}
	this.NomFlux = nomFlux
	return &this
}

// NewRequeteSoumissionFluxWithDefaults instantiates a new RequeteSoumissionFlux object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequeteSoumissionFluxWithDefaults() *RequeteSoumissionFlux {
	this := RequeteSoumissionFlux{}
	return &this
}

// GetNomFlux returns the NomFlux field value
func (o *RequeteSoumissionFlux) GetNomFlux() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NomFlux
}

// GetNomFluxOk returns a tuple with the NomFlux field value
// and a boolean to check if the value has been set.
func (o *RequeteSoumissionFlux) GetNomFluxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NomFlux, true
}

// SetNomFlux sets field value
func (o *RequeteSoumissionFlux) SetNomFlux(v string) {
	o.NomFlux = v
}

// GetSyntaxeFlux returns the SyntaxeFlux field value if set, zero value otherwise.
func (o *RequeteSoumissionFlux) GetSyntaxeFlux() SyntaxeFlux {
	if o == nil || IsNil(o.SyntaxeFlux) {
		var ret SyntaxeFlux
		return ret
	}
	return *o.SyntaxeFlux
}

// GetSyntaxeFluxOk returns a tuple with the SyntaxeFlux field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequeteSoumissionFlux) GetSyntaxeFluxOk() (*SyntaxeFlux, bool) {
	if o == nil || IsNil(o.SyntaxeFlux) {
		return nil, false
	}
	return o.SyntaxeFlux, true
}

// HasSyntaxeFlux returns a boolean if a field has been set.
func (o *RequeteSoumissionFlux) HasSyntaxeFlux() bool {
	if o != nil && !IsNil(o.SyntaxeFlux) {
		return true
	}

	return false
}

// SetSyntaxeFlux gets a reference to the given SyntaxeFlux and assigns it to the SyntaxeFlux field.
func (o *RequeteSoumissionFlux) SetSyntaxeFlux(v SyntaxeFlux) {
	o.SyntaxeFlux = &v
}

// GetProfilFlux returns the ProfilFlux field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequeteSoumissionFlux) GetProfilFlux() ProfilFlux {
	if o == nil || IsNil(o.ProfilFlux.Get()) {
		var ret ProfilFlux
		return ret
	}
	return *o.ProfilFlux.Get()
}

// GetProfilFluxOk returns a tuple with the ProfilFlux field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequeteSoumissionFlux) GetProfilFluxOk() (*ProfilFlux, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilFlux.Get(), o.ProfilFlux.IsSet()
}

// HasProfilFlux returns a boolean if a field has been set.
func (o *RequeteSoumissionFlux) HasProfilFlux() bool {
	if o != nil && o.ProfilFlux.IsSet() {
		return true
	}

	return false
}

// SetProfilFlux gets a reference to the given NullableProfilFlux and assigns it to the ProfilFlux field.
func (o *RequeteSoumissionFlux) SetProfilFlux(v ProfilFlux) {
	o.ProfilFlux.Set(&v)
}
// SetProfilFluxNil sets the value for ProfilFlux to be an explicit nil
func (o *RequeteSoumissionFlux) SetProfilFluxNil() {
	o.ProfilFlux.Set(nil)
}

// UnsetProfilFlux ensures that no value is present for ProfilFlux, not even an explicit nil
func (o *RequeteSoumissionFlux) UnsetProfilFlux() {
	o.ProfilFlux.Unset()
}

// GetTrackingId returns the TrackingId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequeteSoumissionFlux) GetTrackingId() string {
	if o == nil || IsNil(o.TrackingId.Get()) {
		var ret string
		return ret
	}
	return *o.TrackingId.Get()
}

// GetTrackingIdOk returns a tuple with the TrackingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequeteSoumissionFlux) GetTrackingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingId.Get(), o.TrackingId.IsSet()
}

// HasTrackingId returns a boolean if a field has been set.
func (o *RequeteSoumissionFlux) HasTrackingId() bool {
	if o != nil && o.TrackingId.IsSet() {
		return true
	}

	return false
}

// SetTrackingId gets a reference to the given NullableString and assigns it to the TrackingId field.
func (o *RequeteSoumissionFlux) SetTrackingId(v string) {
	o.TrackingId.Set(&v)
}
// SetTrackingIdNil sets the value for TrackingId to be an explicit nil
func (o *RequeteSoumissionFlux) SetTrackingIdNil() {
	o.TrackingId.Set(nil)
}

// UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
func (o *RequeteSoumissionFlux) UnsetTrackingId() {
	o.TrackingId.Unset()
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequeteSoumissionFlux) GetRequestId() string {
	if o == nil || IsNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequeteSoumissionFlux) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *RequeteSoumissionFlux) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *RequeteSoumissionFlux) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *RequeteSoumissionFlux) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *RequeteSoumissionFlux) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetPdpCredentials returns the PdpCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequeteSoumissionFlux) GetPdpCredentials() PDPCredentials {
	if o == nil || IsNil(o.PdpCredentials.Get()) {
		var ret PDPCredentials
		return ret
	}
	return *o.PdpCredentials.Get()
}

// GetPdpCredentialsOk returns a tuple with the PdpCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequeteSoumissionFlux) GetPdpCredentialsOk() (*PDPCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.PdpCredentials.Get(), o.PdpCredentials.IsSet()
}

// HasPdpCredentials returns a boolean if a field has been set.
func (o *RequeteSoumissionFlux) HasPdpCredentials() bool {
	if o != nil && o.PdpCredentials.IsSet() {
		return true
	}

	return false
}

// SetPdpCredentials gets a reference to the given NullablePDPCredentials and assigns it to the PdpCredentials field.
func (o *RequeteSoumissionFlux) SetPdpCredentials(v PDPCredentials) {
	o.PdpCredentials.Set(&v)
}
// SetPdpCredentialsNil sets the value for PdpCredentials to be an explicit nil
func (o *RequeteSoumissionFlux) SetPdpCredentialsNil() {
	o.PdpCredentials.Set(nil)
}

// UnsetPdpCredentials ensures that no value is present for PdpCredentials, not even an explicit nil
func (o *RequeteSoumissionFlux) UnsetPdpCredentials() {
	o.PdpCredentials.Unset()
}

func (o RequeteSoumissionFlux) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequeteSoumissionFlux) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nom_flux"] = o.NomFlux
	if !IsNil(o.SyntaxeFlux) {
		toSerialize["syntaxe_flux"] = o.SyntaxeFlux
	}
	if o.ProfilFlux.IsSet() {
		toSerialize["profil_flux"] = o.ProfilFlux.Get()
	}
	if o.TrackingId.IsSet() {
		toSerialize["tracking_id"] = o.TrackingId.Get()
	}
	if o.RequestId.IsSet() {
		toSerialize["request_id"] = o.RequestId.Get()
	}
	if o.PdpCredentials.IsSet() {
		toSerialize["pdp_credentials"] = o.PdpCredentials.Get()
	}
	return toSerialize, nil
}

func (o *RequeteSoumissionFlux) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nom_flux",
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

	varRequeteSoumissionFlux := _RequeteSoumissionFlux{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRequeteSoumissionFlux)

	if err != nil {
		return err
	}

	*o = RequeteSoumissionFlux(varRequeteSoumissionFlux)

	return err
}

type NullableRequeteSoumissionFlux struct {
	value *RequeteSoumissionFlux
	isSet bool
}

func (v NullableRequeteSoumissionFlux) Get() *RequeteSoumissionFlux {
	return v.value
}

func (v *NullableRequeteSoumissionFlux) Set(val *RequeteSoumissionFlux) {
	v.value = val
	v.isSet = true
}

func (v NullableRequeteSoumissionFlux) IsSet() bool {
	return v.isSet
}

func (v *NullableRequeteSoumissionFlux) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequeteSoumissionFlux(val *RequeteSoumissionFlux) *NullableRequeteSoumissionFlux {
	return &NullableRequeteSoumissionFlux{value: val, isSet: true}
}

func (v NullableRequeteSoumissionFlux) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequeteSoumissionFlux) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


