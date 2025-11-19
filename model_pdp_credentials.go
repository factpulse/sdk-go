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

// checks if the PDPCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PDPCredentials{}

// PDPCredentials Credentials PDP pour la stratégie zero-storage (Strategy B).  Permet de fournir directement les credentials PDP dans la requête au lieu de les stocker dans Django.  Utile pour : - Tests ponctuels sans persister les credentials - Intégrations temporaires - Environnements de développement
type PDPCredentials struct {
	// URL de base du Flow Service AFNOR
	FlowServiceUrl string `json:"flow_service_url"`
	DirectoryServiceUrl NullableString `json:"directory_service_url,omitempty"`
	// URL du serveur OAuth2
	TokenUrl string `json:"token_url"`
	// Client ID OAuth2
	ClientId string `json:"client_id"`
	// Client Secret OAuth2 (sensible)
	ClientSecret string `json:"client_secret"`
}

type _PDPCredentials PDPCredentials

// NewPDPCredentials instantiates a new PDPCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPDPCredentials(flowServiceUrl string, tokenUrl string, clientId string, clientSecret string) *PDPCredentials {
	this := PDPCredentials{}
	this.FlowServiceUrl = flowServiceUrl
	this.TokenUrl = tokenUrl
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewPDPCredentialsWithDefaults instantiates a new PDPCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPDPCredentialsWithDefaults() *PDPCredentials {
	this := PDPCredentials{}
	return &this
}

// GetFlowServiceUrl returns the FlowServiceUrl field value
func (o *PDPCredentials) GetFlowServiceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlowServiceUrl
}

// GetFlowServiceUrlOk returns a tuple with the FlowServiceUrl field value
// and a boolean to check if the value has been set.
func (o *PDPCredentials) GetFlowServiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlowServiceUrl, true
}

// SetFlowServiceUrl sets field value
func (o *PDPCredentials) SetFlowServiceUrl(v string) {
	o.FlowServiceUrl = v
}

// GetDirectoryServiceUrl returns the DirectoryServiceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PDPCredentials) GetDirectoryServiceUrl() string {
	if o == nil || IsNil(o.DirectoryServiceUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DirectoryServiceUrl.Get()
}

// GetDirectoryServiceUrlOk returns a tuple with the DirectoryServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PDPCredentials) GetDirectoryServiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectoryServiceUrl.Get(), o.DirectoryServiceUrl.IsSet()
}

// HasDirectoryServiceUrl returns a boolean if a field has been set.
func (o *PDPCredentials) HasDirectoryServiceUrl() bool {
	if o != nil && o.DirectoryServiceUrl.IsSet() {
		return true
	}

	return false
}

// SetDirectoryServiceUrl gets a reference to the given NullableString and assigns it to the DirectoryServiceUrl field.
func (o *PDPCredentials) SetDirectoryServiceUrl(v string) {
	o.DirectoryServiceUrl.Set(&v)
}
// SetDirectoryServiceUrlNil sets the value for DirectoryServiceUrl to be an explicit nil
func (o *PDPCredentials) SetDirectoryServiceUrlNil() {
	o.DirectoryServiceUrl.Set(nil)
}

// UnsetDirectoryServiceUrl ensures that no value is present for DirectoryServiceUrl, not even an explicit nil
func (o *PDPCredentials) UnsetDirectoryServiceUrl() {
	o.DirectoryServiceUrl.Unset()
}

// GetTokenUrl returns the TokenUrl field value
func (o *PDPCredentials) GetTokenUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value
// and a boolean to check if the value has been set.
func (o *PDPCredentials) GetTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenUrl, true
}

// SetTokenUrl sets field value
func (o *PDPCredentials) SetTokenUrl(v string) {
	o.TokenUrl = v
}

// GetClientId returns the ClientId field value
func (o *PDPCredentials) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *PDPCredentials) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *PDPCredentials) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *PDPCredentials) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *PDPCredentials) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *PDPCredentials) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o PDPCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PDPCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["flow_service_url"] = o.FlowServiceUrl
	if o.DirectoryServiceUrl.IsSet() {
		toSerialize["directory_service_url"] = o.DirectoryServiceUrl.Get()
	}
	toSerialize["token_url"] = o.TokenUrl
	toSerialize["client_id"] = o.ClientId
	toSerialize["client_secret"] = o.ClientSecret
	return toSerialize, nil
}

func (o *PDPCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"flow_service_url",
		"token_url",
		"client_id",
		"client_secret",
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

	varPDPCredentials := _PDPCredentials{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPDPCredentials)

	if err != nil {
		return err
	}

	*o = PDPCredentials(varPDPCredentials)

	return err
}

type NullablePDPCredentials struct {
	value *PDPCredentials
	isSet bool
}

func (v NullablePDPCredentials) Get() *PDPCredentials {
	return v.value
}

func (v *NullablePDPCredentials) Set(val *PDPCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullablePDPCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullablePDPCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePDPCredentials(val *PDPCredentials) *NullablePDPCredentials {
	return &NullablePDPCredentials{value: val, isSet: true}
}

func (v NullablePDPCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePDPCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


