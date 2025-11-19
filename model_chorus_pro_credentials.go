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

// checks if the ChorusProCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChorusProCredentials{}

// ChorusProCredentials Credentials Chorus Pro pour mode Zero-Trust.  **Mode Zero-Trust** : Les credentials sont passés dans chaque requête et ne sont JAMAIS stockés.  **Sécurité** : - Les credentials ne sont jamais persistés dans la base de données - Ils sont utilisés uniquement pour la durée de la requête - Transmission sécurisée via HTTPS  **Cas d'usage** : - Environnements à haute sécurité (banques, administrations) - Conformité RGPD stricte - Tests avec credentials temporaires - Utilisateurs ne voulant pas stocker leurs credentials
type ChorusProCredentials struct {
	// Client ID PISTE (portail API gouvernement)
	PisteClientId string `json:"piste_client_id"`
	// Client Secret PISTE
	PisteClientSecret string `json:"piste_client_secret"`
	// Login Chorus Pro
	ChorusProLogin string `json:"chorus_pro_login"`
	// Mot de passe Chorus Pro
	ChorusProPassword string `json:"chorus_pro_password"`
	// Utiliser l'environnement sandbox (true) ou production (false)
	Sandbox *bool `json:"sandbox,omitempty"`
}

type _ChorusProCredentials ChorusProCredentials

// NewChorusProCredentials instantiates a new ChorusProCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChorusProCredentials(pisteClientId string, pisteClientSecret string, chorusProLogin string, chorusProPassword string) *ChorusProCredentials {
	this := ChorusProCredentials{}
	this.PisteClientId = pisteClientId
	this.PisteClientSecret = pisteClientSecret
	this.ChorusProLogin = chorusProLogin
	this.ChorusProPassword = chorusProPassword
	var sandbox bool = true
	this.Sandbox = &sandbox
	return &this
}

// NewChorusProCredentialsWithDefaults instantiates a new ChorusProCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChorusProCredentialsWithDefaults() *ChorusProCredentials {
	this := ChorusProCredentials{}
	var sandbox bool = true
	this.Sandbox = &sandbox
	return &this
}

// GetPisteClientId returns the PisteClientId field value
func (o *ChorusProCredentials) GetPisteClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PisteClientId
}

// GetPisteClientIdOk returns a tuple with the PisteClientId field value
// and a boolean to check if the value has been set.
func (o *ChorusProCredentials) GetPisteClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PisteClientId, true
}

// SetPisteClientId sets field value
func (o *ChorusProCredentials) SetPisteClientId(v string) {
	o.PisteClientId = v
}

// GetPisteClientSecret returns the PisteClientSecret field value
func (o *ChorusProCredentials) GetPisteClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PisteClientSecret
}

// GetPisteClientSecretOk returns a tuple with the PisteClientSecret field value
// and a boolean to check if the value has been set.
func (o *ChorusProCredentials) GetPisteClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PisteClientSecret, true
}

// SetPisteClientSecret sets field value
func (o *ChorusProCredentials) SetPisteClientSecret(v string) {
	o.PisteClientSecret = v
}

// GetChorusProLogin returns the ChorusProLogin field value
func (o *ChorusProCredentials) GetChorusProLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChorusProLogin
}

// GetChorusProLoginOk returns a tuple with the ChorusProLogin field value
// and a boolean to check if the value has been set.
func (o *ChorusProCredentials) GetChorusProLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChorusProLogin, true
}

// SetChorusProLogin sets field value
func (o *ChorusProCredentials) SetChorusProLogin(v string) {
	o.ChorusProLogin = v
}

// GetChorusProPassword returns the ChorusProPassword field value
func (o *ChorusProCredentials) GetChorusProPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChorusProPassword
}

// GetChorusProPasswordOk returns a tuple with the ChorusProPassword field value
// and a boolean to check if the value has been set.
func (o *ChorusProCredentials) GetChorusProPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChorusProPassword, true
}

// SetChorusProPassword sets field value
func (o *ChorusProCredentials) SetChorusProPassword(v string) {
	o.ChorusProPassword = v
}

// GetSandbox returns the Sandbox field value if set, zero value otherwise.
func (o *ChorusProCredentials) GetSandbox() bool {
	if o == nil || IsNil(o.Sandbox) {
		var ret bool
		return ret
	}
	return *o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChorusProCredentials) GetSandboxOk() (*bool, bool) {
	if o == nil || IsNil(o.Sandbox) {
		return nil, false
	}
	return o.Sandbox, true
}

// HasSandbox returns a boolean if a field has been set.
func (o *ChorusProCredentials) HasSandbox() bool {
	if o != nil && !IsNil(o.Sandbox) {
		return true
	}

	return false
}

// SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.
func (o *ChorusProCredentials) SetSandbox(v bool) {
	o.Sandbox = &v
}

func (o ChorusProCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChorusProCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["piste_client_id"] = o.PisteClientId
	toSerialize["piste_client_secret"] = o.PisteClientSecret
	toSerialize["chorus_pro_login"] = o.ChorusProLogin
	toSerialize["chorus_pro_password"] = o.ChorusProPassword
	if !IsNil(o.Sandbox) {
		toSerialize["sandbox"] = o.Sandbox
	}
	return toSerialize, nil
}

func (o *ChorusProCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"piste_client_id",
		"piste_client_secret",
		"chorus_pro_login",
		"chorus_pro_password",
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

	varChorusProCredentials := _ChorusProCredentials{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChorusProCredentials)

	if err != nil {
		return err
	}

	*o = ChorusProCredentials(varChorusProCredentials)

	return err
}

type NullableChorusProCredentials struct {
	value *ChorusProCredentials
	isSet bool
}

func (v NullableChorusProCredentials) Get() *ChorusProCredentials {
	return v.value
}

func (v *NullableChorusProCredentials) Set(val *ChorusProCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableChorusProCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableChorusProCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChorusProCredentials(val *ChorusProCredentials) *NullableChorusProCredentials {
	return &NullableChorusProCredentials{value: val, isSet: true}
}

func (v NullableChorusProCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChorusProCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


