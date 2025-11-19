/*
API REST FactPulse

 API REST pour la facturation électronique en France : Factur-X, AFNOR PDP/PA, signatures électroniques.  ## 🎯 Fonctionnalités principales  ### 📄 Génération de factures Factur-X - **Formats** : XML seul ou PDF/A-3 avec XML embarqué - **Profils** : MINIMUM, BASIC, EN16931, EXTENDED - **Normes** : EN 16931 (directive UE 2014/55), ISO 19005-3 (PDF/A-3), CII (UN/CEFACT) - **🆕 Format simplifié** : Génération à partir de SIRET + auto-enrichissement (API Chorus Pro + Recherche Entreprises)  ### ✅ Validation et conformité - **Validation XML** : Schematron (45 à 210+ règles selon profil) - **Validation PDF** : PDF/A-3, métadonnées XMP Factur-X, signatures électroniques - **VeraPDF** : Validation stricte PDF/A (146+ règles ISO 19005-3) - **Traitement asynchrone** : Support Celery pour validations lourdes (VeraPDF)  ### 📡 Intégration AFNOR PDP/PA (XP Z12-013) - **Soumission de flux** : Envoi de factures vers Plateformes de Dématérialisation Partenaires - **Recherche de flux** : Consultation des factures soumises - **Téléchargement** : Récupération des PDF/A-3 avec XML - **Directory Service** : Recherche d'entreprises (SIREN/SIRET) - **Multi-client** : Support de plusieurs configs PDP par utilisateur (stored credentials ou zero-storage)  ### ✍️ Signature électronique PDF - **Standards** : PAdES-B-B, PAdES-B-T (horodatage RFC 3161), PAdES-B-LT (archivage long terme) - **Niveaux eIDAS** : SES (auto-signé), AdES (CA commerciale), QES (PSCO) - **Validation** : Vérification intégrité cryptographique et certificats - **Génération de certificats** : Certificats X.509 auto-signés pour tests  ### 🔄 Traitement asynchrone - **Celery** : Génération, validation et signature asynchrones - **Polling** : Suivi d'état via `/taches/{id_tache}/statut` - **Pas de timeout** : Idéal pour gros fichiers ou validations lourdes  ## 🔒 Authentification  Toutes les requêtes nécessitent un **token JWT** dans le header Authorization : ``` Authorization: Bearer YOUR_JWT_TOKEN ```  ### Comment obtenir un token JWT ?  #### 🔑 Méthode 1 : API `/api/token/` (Recommandée)  **URL :** `https://www.factpulse.fr/api/token/`  Cette méthode est **recommandée** pour l'intégration dans vos applications et workflows CI/CD.  **Prérequis :** Avoir défini un mot de passe sur votre compte  **Pour les utilisateurs inscrits via email/password :** - Vous avez déjà un mot de passe, utilisez-le directement  **Pour les utilisateurs inscrits via OAuth (Google/GitHub) :** - Vous devez d'abord définir un mot de passe sur : https://www.factpulse.fr/accounts/password/set/ - Une fois le mot de passe créé, vous pourrez utiliser l'API  **Exemple de requête :** ```bash curl -X POST https://www.factpulse.fr/api/token/ \\   -H \"Content-Type: application/json\" \\   -d '{     \"username\": \"votre_email@example.com\",     \"password\": \"votre_mot_de_passe\"   }' ```  **Paramètre optionnel `client_uid` :**  Pour sélectionner les credentials d'un client spécifique (PA/PDP, Chorus Pro, certificats de signature), ajoutez `client_uid` :  ```bash curl -X POST https://www.factpulse.fr/api/token/ \\   -H \"Content-Type: application/json\" \\   -d '{     \"username\": \"votre_email@example.com\",     \"password\": \"votre_mot_de_passe\",     \"client_uid\": \"550e8400-e29b-41d4-a716-446655440000\"   }' ```  Le `client_uid` sera inclus dans le JWT et permettra à l'API d'utiliser automatiquement : - Les credentials AFNOR/PDP configurés pour ce client - Les credentials Chorus Pro configurés pour ce client - Les certificats de signature électronique configurés pour ce client  **Réponse :** ```json {   \"access\": \"eyJ0eXAiOiJKV1QiLCJhbGc...\",  // Token d'accès (validité: 30 min)   \"refresh\": \"eyJ0eXAiOiJKV1QiLCJhbGc...\"  // Token de rafraîchissement (validité: 7 jours) } ```  **Avantages :** - ✅ Automatisation complète (CI/CD, scripts) - ✅ Gestion programmatique des tokens - ✅ Support du refresh token pour renouveler automatiquement l'accès - ✅ Intégration facile dans n'importe quel langage/outil  #### 🖥️ Méthode 2 : Génération via Dashboard (Alternative)  **URL :** https://www.factpulse.fr/dashboard/  Cette méthode convient pour des tests rapides ou une utilisation occasionnelle via l'interface graphique.  **Fonctionnement :** - Connectez-vous au dashboard - Utilisez les boutons \"Generate Test Token\" ou \"Generate Production Token\" - Fonctionne pour **tous** les utilisateurs (OAuth et email/password), sans nécessiter de mot de passe  **Types de tokens :** - **Token Test** : Validité 24h, quota 1000 appels/jour (gratuit) - **Token Production** : Validité 7 jours, quota selon votre forfait  **Avantages :** - ✅ Rapide pour tester l'API - ✅ Aucun mot de passe requis - ✅ Interface visuelle simple  **Inconvénients :** - ❌ Nécessite une action manuelle - ❌ Pas de refresh token - ❌ Moins adapté pour l'automatisation  ### 📚 Documentation complète  Pour plus d'informations sur l'authentification et l'utilisation de l'API : https://www.factpulse.fr/documentation-api/     

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package factpulse

import (
	"encoding/json"
)

// checks if the DestinationAFNOR type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationAFNOR{}

// DestinationAFNOR Configuration spécifique pour la destination AFNOR PDP.
type DestinationAFNOR struct {
	Type *string `json:"type,omitempty"`
	Credentials NullableCredentialsAFNOR `json:"credentials,omitempty"`
	// Syntaxe du flux à envoyer
	FlowSyntax *string `json:"flow_syntax,omitempty"`
	TrackingId NullableString `json:"tracking_id,omitempty"`
}

// NewDestinationAFNOR instantiates a new DestinationAFNOR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationAFNOR() *DestinationAFNOR {
	this := DestinationAFNOR{}
	var type_ string = "afnor"
	this.Type = &type_
	var flowSyntax string = "Factur-X"
	this.FlowSyntax = &flowSyntax
	return &this
}

// NewDestinationAFNORWithDefaults instantiates a new DestinationAFNOR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationAFNORWithDefaults() *DestinationAFNOR {
	this := DestinationAFNOR{}
	var type_ string = "afnor"
	this.Type = &type_
	var flowSyntax string = "Factur-X"
	this.FlowSyntax = &flowSyntax
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DestinationAFNOR) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationAFNOR) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DestinationAFNOR) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DestinationAFNOR) SetType(v string) {
	o.Type = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DestinationAFNOR) GetCredentials() CredentialsAFNOR {
	if o == nil || IsNil(o.Credentials.Get()) {
		var ret CredentialsAFNOR
		return ret
	}
	return *o.Credentials.Get()
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DestinationAFNOR) GetCredentialsOk() (*CredentialsAFNOR, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials.Get(), o.Credentials.IsSet()
}

// HasCredentials returns a boolean if a field has been set.
func (o *DestinationAFNOR) HasCredentials() bool {
	if o != nil && o.Credentials.IsSet() {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given NullableCredentialsAFNOR and assigns it to the Credentials field.
func (o *DestinationAFNOR) SetCredentials(v CredentialsAFNOR) {
	o.Credentials.Set(&v)
}
// SetCredentialsNil sets the value for Credentials to be an explicit nil
func (o *DestinationAFNOR) SetCredentialsNil() {
	o.Credentials.Set(nil)
}

// UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
func (o *DestinationAFNOR) UnsetCredentials() {
	o.Credentials.Unset()
}

// GetFlowSyntax returns the FlowSyntax field value if set, zero value otherwise.
func (o *DestinationAFNOR) GetFlowSyntax() string {
	if o == nil || IsNil(o.FlowSyntax) {
		var ret string
		return ret
	}
	return *o.FlowSyntax
}

// GetFlowSyntaxOk returns a tuple with the FlowSyntax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationAFNOR) GetFlowSyntaxOk() (*string, bool) {
	if o == nil || IsNil(o.FlowSyntax) {
		return nil, false
	}
	return o.FlowSyntax, true
}

// HasFlowSyntax returns a boolean if a field has been set.
func (o *DestinationAFNOR) HasFlowSyntax() bool {
	if o != nil && !IsNil(o.FlowSyntax) {
		return true
	}

	return false
}

// SetFlowSyntax gets a reference to the given string and assigns it to the FlowSyntax field.
func (o *DestinationAFNOR) SetFlowSyntax(v string) {
	o.FlowSyntax = &v
}

// GetTrackingId returns the TrackingId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DestinationAFNOR) GetTrackingId() string {
	if o == nil || IsNil(o.TrackingId.Get()) {
		var ret string
		return ret
	}
	return *o.TrackingId.Get()
}

// GetTrackingIdOk returns a tuple with the TrackingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DestinationAFNOR) GetTrackingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingId.Get(), o.TrackingId.IsSet()
}

// HasTrackingId returns a boolean if a field has been set.
func (o *DestinationAFNOR) HasTrackingId() bool {
	if o != nil && o.TrackingId.IsSet() {
		return true
	}

	return false
}

// SetTrackingId gets a reference to the given NullableString and assigns it to the TrackingId field.
func (o *DestinationAFNOR) SetTrackingId(v string) {
	o.TrackingId.Set(&v)
}
// SetTrackingIdNil sets the value for TrackingId to be an explicit nil
func (o *DestinationAFNOR) SetTrackingIdNil() {
	o.TrackingId.Set(nil)
}

// UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
func (o *DestinationAFNOR) UnsetTrackingId() {
	o.TrackingId.Unset()
}

func (o DestinationAFNOR) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationAFNOR) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Credentials.IsSet() {
		toSerialize["credentials"] = o.Credentials.Get()
	}
	if !IsNil(o.FlowSyntax) {
		toSerialize["flow_syntax"] = o.FlowSyntax
	}
	if o.TrackingId.IsSet() {
		toSerialize["tracking_id"] = o.TrackingId.Get()
	}
	return toSerialize, nil
}

type NullableDestinationAFNOR struct {
	value *DestinationAFNOR
	isSet bool
}

func (v NullableDestinationAFNOR) Get() *DestinationAFNOR {
	return v.value
}

func (v *NullableDestinationAFNOR) Set(val *DestinationAFNOR) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationAFNOR) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationAFNOR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationAFNOR(val *DestinationAFNOR) *NullableDestinationAFNOR {
	return &NullableDestinationAFNOR{value: val, isSet: true}
}

func (v NullableDestinationAFNOR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationAFNOR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


