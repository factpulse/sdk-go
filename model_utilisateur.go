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

// checks if the Utilisateur type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Utilisateur{}

// Utilisateur Modèle Pydantic représentant les données de l'utilisateur authentifié.
type Utilisateur struct {
	Id int32 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	IsActive bool `json:"is_active"`
	IsSuperuser *bool `json:"is_superuser,omitempty"`
	IsStaff *bool `json:"is_staff,omitempty"`
	BypassQuota *bool `json:"bypass_quota,omitempty"`
	TeamId NullableInt32 `json:"team_id,omitempty"`
	HasQuota *bool `json:"has_quota,omitempty"`
	QuotaInfo NullableQuotaInfo `json:"quota_info,omitempty"`
	IsTrial *bool `json:"is_trial,omitempty"`
	ClientUid NullableString `json:"client_uid,omitempty"`
}

type _Utilisateur Utilisateur

// NewUtilisateur instantiates a new Utilisateur object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilisateur(id int32, username string, email string, isActive bool) *Utilisateur {
	this := Utilisateur{}
	this.Id = id
	this.Username = username
	this.Email = email
	this.IsActive = isActive
	var isSuperuser bool = false
	this.IsSuperuser = &isSuperuser
	var isStaff bool = false
	this.IsStaff = &isStaff
	var bypassQuota bool = false
	this.BypassQuota = &bypassQuota
	var hasQuota bool = true
	this.HasQuota = &hasQuota
	var isTrial bool = false
	this.IsTrial = &isTrial
	return &this
}

// NewUtilisateurWithDefaults instantiates a new Utilisateur object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilisateurWithDefaults() *Utilisateur {
	this := Utilisateur{}
	var isSuperuser bool = false
	this.IsSuperuser = &isSuperuser
	var isStaff bool = false
	this.IsStaff = &isStaff
	var bypassQuota bool = false
	this.BypassQuota = &bypassQuota
	var hasQuota bool = true
	this.HasQuota = &hasQuota
	var isTrial bool = false
	this.IsTrial = &isTrial
	return &this
}

// GetId returns the Id field value
func (o *Utilisateur) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Utilisateur) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Utilisateur) SetId(v int32) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *Utilisateur) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Utilisateur) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Utilisateur) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value
func (o *Utilisateur) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Utilisateur) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Utilisateur) SetEmail(v string) {
	o.Email = v
}

// GetIsActive returns the IsActive field value
func (o *Utilisateur) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *Utilisateur) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *Utilisateur) SetIsActive(v bool) {
	o.IsActive = v
}

// GetIsSuperuser returns the IsSuperuser field value if set, zero value otherwise.
func (o *Utilisateur) GetIsSuperuser() bool {
	if o == nil || IsNil(o.IsSuperuser) {
		var ret bool
		return ret
	}
	return *o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utilisateur) GetIsSuperuserOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuperuser) {
		return nil, false
	}
	return o.IsSuperuser, true
}

// HasIsSuperuser returns a boolean if a field has been set.
func (o *Utilisateur) HasIsSuperuser() bool {
	if o != nil && !IsNil(o.IsSuperuser) {
		return true
	}

	return false
}

// SetIsSuperuser gets a reference to the given bool and assigns it to the IsSuperuser field.
func (o *Utilisateur) SetIsSuperuser(v bool) {
	o.IsSuperuser = &v
}

// GetIsStaff returns the IsStaff field value if set, zero value otherwise.
func (o *Utilisateur) GetIsStaff() bool {
	if o == nil || IsNil(o.IsStaff) {
		var ret bool
		return ret
	}
	return *o.IsStaff
}

// GetIsStaffOk returns a tuple with the IsStaff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utilisateur) GetIsStaffOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStaff) {
		return nil, false
	}
	return o.IsStaff, true
}

// HasIsStaff returns a boolean if a field has been set.
func (o *Utilisateur) HasIsStaff() bool {
	if o != nil && !IsNil(o.IsStaff) {
		return true
	}

	return false
}

// SetIsStaff gets a reference to the given bool and assigns it to the IsStaff field.
func (o *Utilisateur) SetIsStaff(v bool) {
	o.IsStaff = &v
}

// GetBypassQuota returns the BypassQuota field value if set, zero value otherwise.
func (o *Utilisateur) GetBypassQuota() bool {
	if o == nil || IsNil(o.BypassQuota) {
		var ret bool
		return ret
	}
	return *o.BypassQuota
}

// GetBypassQuotaOk returns a tuple with the BypassQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utilisateur) GetBypassQuotaOk() (*bool, bool) {
	if o == nil || IsNil(o.BypassQuota) {
		return nil, false
	}
	return o.BypassQuota, true
}

// HasBypassQuota returns a boolean if a field has been set.
func (o *Utilisateur) HasBypassQuota() bool {
	if o != nil && !IsNil(o.BypassQuota) {
		return true
	}

	return false
}

// SetBypassQuota gets a reference to the given bool and assigns it to the BypassQuota field.
func (o *Utilisateur) SetBypassQuota(v bool) {
	o.BypassQuota = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Utilisateur) GetTeamId() int32 {
	if o == nil || IsNil(o.TeamId.Get()) {
		var ret int32
		return ret
	}
	return *o.TeamId.Get()
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Utilisateur) GetTeamIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamId.Get(), o.TeamId.IsSet()
}

// HasTeamId returns a boolean if a field has been set.
func (o *Utilisateur) HasTeamId() bool {
	if o != nil && o.TeamId.IsSet() {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given NullableInt32 and assigns it to the TeamId field.
func (o *Utilisateur) SetTeamId(v int32) {
	o.TeamId.Set(&v)
}
// SetTeamIdNil sets the value for TeamId to be an explicit nil
func (o *Utilisateur) SetTeamIdNil() {
	o.TeamId.Set(nil)
}

// UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
func (o *Utilisateur) UnsetTeamId() {
	o.TeamId.Unset()
}

// GetHasQuota returns the HasQuota field value if set, zero value otherwise.
func (o *Utilisateur) GetHasQuota() bool {
	if o == nil || IsNil(o.HasQuota) {
		var ret bool
		return ret
	}
	return *o.HasQuota
}

// GetHasQuotaOk returns a tuple with the HasQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utilisateur) GetHasQuotaOk() (*bool, bool) {
	if o == nil || IsNil(o.HasQuota) {
		return nil, false
	}
	return o.HasQuota, true
}

// HasHasQuota returns a boolean if a field has been set.
func (o *Utilisateur) HasHasQuota() bool {
	if o != nil && !IsNil(o.HasQuota) {
		return true
	}

	return false
}

// SetHasQuota gets a reference to the given bool and assigns it to the HasQuota field.
func (o *Utilisateur) SetHasQuota(v bool) {
	o.HasQuota = &v
}

// GetQuotaInfo returns the QuotaInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Utilisateur) GetQuotaInfo() QuotaInfo {
	if o == nil || IsNil(o.QuotaInfo.Get()) {
		var ret QuotaInfo
		return ret
	}
	return *o.QuotaInfo.Get()
}

// GetQuotaInfoOk returns a tuple with the QuotaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Utilisateur) GetQuotaInfoOk() (*QuotaInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuotaInfo.Get(), o.QuotaInfo.IsSet()
}

// HasQuotaInfo returns a boolean if a field has been set.
func (o *Utilisateur) HasQuotaInfo() bool {
	if o != nil && o.QuotaInfo.IsSet() {
		return true
	}

	return false
}

// SetQuotaInfo gets a reference to the given NullableQuotaInfo and assigns it to the QuotaInfo field.
func (o *Utilisateur) SetQuotaInfo(v QuotaInfo) {
	o.QuotaInfo.Set(&v)
}
// SetQuotaInfoNil sets the value for QuotaInfo to be an explicit nil
func (o *Utilisateur) SetQuotaInfoNil() {
	o.QuotaInfo.Set(nil)
}

// UnsetQuotaInfo ensures that no value is present for QuotaInfo, not even an explicit nil
func (o *Utilisateur) UnsetQuotaInfo() {
	o.QuotaInfo.Unset()
}

// GetIsTrial returns the IsTrial field value if set, zero value otherwise.
func (o *Utilisateur) GetIsTrial() bool {
	if o == nil || IsNil(o.IsTrial) {
		var ret bool
		return ret
	}
	return *o.IsTrial
}

// GetIsTrialOk returns a tuple with the IsTrial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utilisateur) GetIsTrialOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTrial) {
		return nil, false
	}
	return o.IsTrial, true
}

// HasIsTrial returns a boolean if a field has been set.
func (o *Utilisateur) HasIsTrial() bool {
	if o != nil && !IsNil(o.IsTrial) {
		return true
	}

	return false
}

// SetIsTrial gets a reference to the given bool and assigns it to the IsTrial field.
func (o *Utilisateur) SetIsTrial(v bool) {
	o.IsTrial = &v
}

// GetClientUid returns the ClientUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Utilisateur) GetClientUid() string {
	if o == nil || IsNil(o.ClientUid.Get()) {
		var ret string
		return ret
	}
	return *o.ClientUid.Get()
}

// GetClientUidOk returns a tuple with the ClientUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Utilisateur) GetClientUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientUid.Get(), o.ClientUid.IsSet()
}

// HasClientUid returns a boolean if a field has been set.
func (o *Utilisateur) HasClientUid() bool {
	if o != nil && o.ClientUid.IsSet() {
		return true
	}

	return false
}

// SetClientUid gets a reference to the given NullableString and assigns it to the ClientUid field.
func (o *Utilisateur) SetClientUid(v string) {
	o.ClientUid.Set(&v)
}
// SetClientUidNil sets the value for ClientUid to be an explicit nil
func (o *Utilisateur) SetClientUidNil() {
	o.ClientUid.Set(nil)
}

// UnsetClientUid ensures that no value is present for ClientUid, not even an explicit nil
func (o *Utilisateur) UnsetClientUid() {
	o.ClientUid.Unset()
}

func (o Utilisateur) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Utilisateur) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["username"] = o.Username
	toSerialize["email"] = o.Email
	toSerialize["is_active"] = o.IsActive
	if !IsNil(o.IsSuperuser) {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	if !IsNil(o.IsStaff) {
		toSerialize["is_staff"] = o.IsStaff
	}
	if !IsNil(o.BypassQuota) {
		toSerialize["bypass_quota"] = o.BypassQuota
	}
	if o.TeamId.IsSet() {
		toSerialize["team_id"] = o.TeamId.Get()
	}
	if !IsNil(o.HasQuota) {
		toSerialize["has_quota"] = o.HasQuota
	}
	if o.QuotaInfo.IsSet() {
		toSerialize["quota_info"] = o.QuotaInfo.Get()
	}
	if !IsNil(o.IsTrial) {
		toSerialize["is_trial"] = o.IsTrial
	}
	if o.ClientUid.IsSet() {
		toSerialize["client_uid"] = o.ClientUid.Get()
	}
	return toSerialize, nil
}

func (o *Utilisateur) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"username",
		"email",
		"is_active",
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

	varUtilisateur := _Utilisateur{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUtilisateur)

	if err != nil {
		return err
	}

	*o = Utilisateur(varUtilisateur)

	return err
}

type NullableUtilisateur struct {
	value *Utilisateur
	isSet bool
}

func (v NullableUtilisateur) Get() *Utilisateur {
	return v.value
}

func (v *NullableUtilisateur) Set(val *Utilisateur) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilisateur) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilisateur) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilisateur(val *Utilisateur) *NullableUtilisateur {
	return &NullableUtilisateur{value: val, isSet: true}
}

func (v NullableUtilisateur) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilisateur) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


