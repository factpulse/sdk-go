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

// checks if the CredentialsChorusPro type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsChorusPro{}

// CredentialsChorusPro Credentials Chorus Pro optionnels.  **MODE 1 - Récupération via JWT (recommandé) :** Ne pas fournir ce champ `credentials` dans le payload. Les credentials seront récupérés automatiquement via client_uid du JWT (0-trust).  **MODE 2 - Credentials dans le payload :** Fournir tous les champs requis ci-dessous. Utile pour tests ou intégrations tierces.
type CredentialsChorusPro struct {
	PisteClientId NullableString `json:"piste_client_id,omitempty"`
	PisteClientSecret NullableString `json:"piste_client_secret,omitempty"`
	ChorusLogin NullableString `json:"chorus_login,omitempty"`
	ChorusPassword NullableString `json:"chorus_password,omitempty"`
	// [MODE 2] Utiliser le mode sandbox (défaut: True)
	ModeSandbox *bool `json:"mode_sandbox,omitempty"`
}

// NewCredentialsChorusPro instantiates a new CredentialsChorusPro object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsChorusPro() *CredentialsChorusPro {
	this := CredentialsChorusPro{}
	var modeSandbox bool = true
	this.ModeSandbox = &modeSandbox
	return &this
}

// NewCredentialsChorusProWithDefaults instantiates a new CredentialsChorusPro object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsChorusProWithDefaults() *CredentialsChorusPro {
	this := CredentialsChorusPro{}
	var modeSandbox bool = true
	this.ModeSandbox = &modeSandbox
	return &this
}

// GetPisteClientId returns the PisteClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialsChorusPro) GetPisteClientId() string {
	if o == nil || IsNil(o.PisteClientId.Get()) {
		var ret string
		return ret
	}
	return *o.PisteClientId.Get()
}

// GetPisteClientIdOk returns a tuple with the PisteClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialsChorusPro) GetPisteClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PisteClientId.Get(), o.PisteClientId.IsSet()
}

// HasPisteClientId returns a boolean if a field has been set.
func (o *CredentialsChorusPro) HasPisteClientId() bool {
	if o != nil && o.PisteClientId.IsSet() {
		return true
	}

	return false
}

// SetPisteClientId gets a reference to the given NullableString and assigns it to the PisteClientId field.
func (o *CredentialsChorusPro) SetPisteClientId(v string) {
	o.PisteClientId.Set(&v)
}
// SetPisteClientIdNil sets the value for PisteClientId to be an explicit nil
func (o *CredentialsChorusPro) SetPisteClientIdNil() {
	o.PisteClientId.Set(nil)
}

// UnsetPisteClientId ensures that no value is present for PisteClientId, not even an explicit nil
func (o *CredentialsChorusPro) UnsetPisteClientId() {
	o.PisteClientId.Unset()
}

// GetPisteClientSecret returns the PisteClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialsChorusPro) GetPisteClientSecret() string {
	if o == nil || IsNil(o.PisteClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.PisteClientSecret.Get()
}

// GetPisteClientSecretOk returns a tuple with the PisteClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialsChorusPro) GetPisteClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PisteClientSecret.Get(), o.PisteClientSecret.IsSet()
}

// HasPisteClientSecret returns a boolean if a field has been set.
func (o *CredentialsChorusPro) HasPisteClientSecret() bool {
	if o != nil && o.PisteClientSecret.IsSet() {
		return true
	}

	return false
}

// SetPisteClientSecret gets a reference to the given NullableString and assigns it to the PisteClientSecret field.
func (o *CredentialsChorusPro) SetPisteClientSecret(v string) {
	o.PisteClientSecret.Set(&v)
}
// SetPisteClientSecretNil sets the value for PisteClientSecret to be an explicit nil
func (o *CredentialsChorusPro) SetPisteClientSecretNil() {
	o.PisteClientSecret.Set(nil)
}

// UnsetPisteClientSecret ensures that no value is present for PisteClientSecret, not even an explicit nil
func (o *CredentialsChorusPro) UnsetPisteClientSecret() {
	o.PisteClientSecret.Unset()
}

// GetChorusLogin returns the ChorusLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialsChorusPro) GetChorusLogin() string {
	if o == nil || IsNil(o.ChorusLogin.Get()) {
		var ret string
		return ret
	}
	return *o.ChorusLogin.Get()
}

// GetChorusLoginOk returns a tuple with the ChorusLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialsChorusPro) GetChorusLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChorusLogin.Get(), o.ChorusLogin.IsSet()
}

// HasChorusLogin returns a boolean if a field has been set.
func (o *CredentialsChorusPro) HasChorusLogin() bool {
	if o != nil && o.ChorusLogin.IsSet() {
		return true
	}

	return false
}

// SetChorusLogin gets a reference to the given NullableString and assigns it to the ChorusLogin field.
func (o *CredentialsChorusPro) SetChorusLogin(v string) {
	o.ChorusLogin.Set(&v)
}
// SetChorusLoginNil sets the value for ChorusLogin to be an explicit nil
func (o *CredentialsChorusPro) SetChorusLoginNil() {
	o.ChorusLogin.Set(nil)
}

// UnsetChorusLogin ensures that no value is present for ChorusLogin, not even an explicit nil
func (o *CredentialsChorusPro) UnsetChorusLogin() {
	o.ChorusLogin.Unset()
}

// GetChorusPassword returns the ChorusPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialsChorusPro) GetChorusPassword() string {
	if o == nil || IsNil(o.ChorusPassword.Get()) {
		var ret string
		return ret
	}
	return *o.ChorusPassword.Get()
}

// GetChorusPasswordOk returns a tuple with the ChorusPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialsChorusPro) GetChorusPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChorusPassword.Get(), o.ChorusPassword.IsSet()
}

// HasChorusPassword returns a boolean if a field has been set.
func (o *CredentialsChorusPro) HasChorusPassword() bool {
	if o != nil && o.ChorusPassword.IsSet() {
		return true
	}

	return false
}

// SetChorusPassword gets a reference to the given NullableString and assigns it to the ChorusPassword field.
func (o *CredentialsChorusPro) SetChorusPassword(v string) {
	o.ChorusPassword.Set(&v)
}
// SetChorusPasswordNil sets the value for ChorusPassword to be an explicit nil
func (o *CredentialsChorusPro) SetChorusPasswordNil() {
	o.ChorusPassword.Set(nil)
}

// UnsetChorusPassword ensures that no value is present for ChorusPassword, not even an explicit nil
func (o *CredentialsChorusPro) UnsetChorusPassword() {
	o.ChorusPassword.Unset()
}

// GetModeSandbox returns the ModeSandbox field value if set, zero value otherwise.
func (o *CredentialsChorusPro) GetModeSandbox() bool {
	if o == nil || IsNil(o.ModeSandbox) {
		var ret bool
		return ret
	}
	return *o.ModeSandbox
}

// GetModeSandboxOk returns a tuple with the ModeSandbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsChorusPro) GetModeSandboxOk() (*bool, bool) {
	if o == nil || IsNil(o.ModeSandbox) {
		return nil, false
	}
	return o.ModeSandbox, true
}

// HasModeSandbox returns a boolean if a field has been set.
func (o *CredentialsChorusPro) HasModeSandbox() bool {
	if o != nil && !IsNil(o.ModeSandbox) {
		return true
	}

	return false
}

// SetModeSandbox gets a reference to the given bool and assigns it to the ModeSandbox field.
func (o *CredentialsChorusPro) SetModeSandbox(v bool) {
	o.ModeSandbox = &v
}

func (o CredentialsChorusPro) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialsChorusPro) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PisteClientId.IsSet() {
		toSerialize["piste_client_id"] = o.PisteClientId.Get()
	}
	if o.PisteClientSecret.IsSet() {
		toSerialize["piste_client_secret"] = o.PisteClientSecret.Get()
	}
	if o.ChorusLogin.IsSet() {
		toSerialize["chorus_login"] = o.ChorusLogin.Get()
	}
	if o.ChorusPassword.IsSet() {
		toSerialize["chorus_password"] = o.ChorusPassword.Get()
	}
	if !IsNil(o.ModeSandbox) {
		toSerialize["mode_sandbox"] = o.ModeSandbox
	}
	return toSerialize, nil
}

type NullableCredentialsChorusPro struct {
	value *CredentialsChorusPro
	isSet bool
}

func (v NullableCredentialsChorusPro) Get() *CredentialsChorusPro {
	return v.value
}

func (v *NullableCredentialsChorusPro) Set(val *CredentialsChorusPro) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsChorusPro) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsChorusPro) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsChorusPro(val *CredentialsChorusPro) *NullableCredentialsChorusPro {
	return &NullableCredentialsChorusPro{value: val, isSet: true}
}

func (v NullableCredentialsChorusPro) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsChorusPro) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


