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

// checks if the GenerateCertificateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateCertificateResponse{}

// GenerateCertificateResponse Réponse après génération d'un certificat de test.  Contient le certificat PEM, la clé privée PEM, et optionnellement le PKCS#12.
type GenerateCertificateResponse struct {
	// Statut de l'opération
	Status *string `json:"status,omitempty"`
	// Certificat X.509 au format PEM
	CertificatPem string `json:"certificat_pem"`
	// Clé privée RSA au format PEM
	ClePriveePem string `json:"cle_privee_pem"`
	Pkcs12Base64 NullableString `json:"pkcs12_base64,omitempty"`
	// Informations sur le certificat généré
	Info CertificateInfoResponse `json:"info"`
	// Avertissement sur l'utilisation du certificat
	Avertissement *string `json:"avertissement,omitempty"`
}

type _GenerateCertificateResponse GenerateCertificateResponse

// NewGenerateCertificateResponse instantiates a new GenerateCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateCertificateResponse(certificatPem string, clePriveePem string, info CertificateInfoResponse) *GenerateCertificateResponse {
	this := GenerateCertificateResponse{}
	var status string = "success"
	this.Status = &status
	this.CertificatPem = certificatPem
	this.ClePriveePem = clePriveePem
	this.Info = info
	var avertissement string = "⚠️ Ce certificat est AUTO-SIGNÉ et destiné uniquement aux TESTS. Ne PAS utiliser en production. Niveau eIDAS : SES (Simple Electronic Signature)"
	this.Avertissement = &avertissement
	return &this
}

// NewGenerateCertificateResponseWithDefaults instantiates a new GenerateCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateCertificateResponseWithDefaults() *GenerateCertificateResponse {
	this := GenerateCertificateResponse{}
	var status string = "success"
	this.Status = &status
	var avertissement string = "⚠️ Ce certificat est AUTO-SIGNÉ et destiné uniquement aux TESTS. Ne PAS utiliser en production. Niveau eIDAS : SES (Simple Electronic Signature)"
	this.Avertissement = &avertissement
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GenerateCertificateResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateCertificateResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GenerateCertificateResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GenerateCertificateResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCertificatPem returns the CertificatPem field value
func (o *GenerateCertificateResponse) GetCertificatPem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificatPem
}

// GetCertificatPemOk returns a tuple with the CertificatPem field value
// and a boolean to check if the value has been set.
func (o *GenerateCertificateResponse) GetCertificatPemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificatPem, true
}

// SetCertificatPem sets field value
func (o *GenerateCertificateResponse) SetCertificatPem(v string) {
	o.CertificatPem = v
}

// GetClePriveePem returns the ClePriveePem field value
func (o *GenerateCertificateResponse) GetClePriveePem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClePriveePem
}

// GetClePriveePemOk returns a tuple with the ClePriveePem field value
// and a boolean to check if the value has been set.
func (o *GenerateCertificateResponse) GetClePriveePemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClePriveePem, true
}

// SetClePriveePem sets field value
func (o *GenerateCertificateResponse) SetClePriveePem(v string) {
	o.ClePriveePem = v
}

// GetPkcs12Base64 returns the Pkcs12Base64 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateCertificateResponse) GetPkcs12Base64() string {
	if o == nil || IsNil(o.Pkcs12Base64.Get()) {
		var ret string
		return ret
	}
	return *o.Pkcs12Base64.Get()
}

// GetPkcs12Base64Ok returns a tuple with the Pkcs12Base64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateCertificateResponse) GetPkcs12Base64Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pkcs12Base64.Get(), o.Pkcs12Base64.IsSet()
}

// HasPkcs12Base64 returns a boolean if a field has been set.
func (o *GenerateCertificateResponse) HasPkcs12Base64() bool {
	if o != nil && o.Pkcs12Base64.IsSet() {
		return true
	}

	return false
}

// SetPkcs12Base64 gets a reference to the given NullableString and assigns it to the Pkcs12Base64 field.
func (o *GenerateCertificateResponse) SetPkcs12Base64(v string) {
	o.Pkcs12Base64.Set(&v)
}
// SetPkcs12Base64Nil sets the value for Pkcs12Base64 to be an explicit nil
func (o *GenerateCertificateResponse) SetPkcs12Base64Nil() {
	o.Pkcs12Base64.Set(nil)
}

// UnsetPkcs12Base64 ensures that no value is present for Pkcs12Base64, not even an explicit nil
func (o *GenerateCertificateResponse) UnsetPkcs12Base64() {
	o.Pkcs12Base64.Unset()
}

// GetInfo returns the Info field value
func (o *GenerateCertificateResponse) GetInfo() CertificateInfoResponse {
	if o == nil {
		var ret CertificateInfoResponse
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *GenerateCertificateResponse) GetInfoOk() (*CertificateInfoResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value
func (o *GenerateCertificateResponse) SetInfo(v CertificateInfoResponse) {
	o.Info = v
}

// GetAvertissement returns the Avertissement field value if set, zero value otherwise.
func (o *GenerateCertificateResponse) GetAvertissement() string {
	if o == nil || IsNil(o.Avertissement) {
		var ret string
		return ret
	}
	return *o.Avertissement
}

// GetAvertissementOk returns a tuple with the Avertissement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateCertificateResponse) GetAvertissementOk() (*string, bool) {
	if o == nil || IsNil(o.Avertissement) {
		return nil, false
	}
	return o.Avertissement, true
}

// HasAvertissement returns a boolean if a field has been set.
func (o *GenerateCertificateResponse) HasAvertissement() bool {
	if o != nil && !IsNil(o.Avertissement) {
		return true
	}

	return false
}

// SetAvertissement gets a reference to the given string and assigns it to the Avertissement field.
func (o *GenerateCertificateResponse) SetAvertissement(v string) {
	o.Avertissement = &v
}

func (o GenerateCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateCertificateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["certificat_pem"] = o.CertificatPem
	toSerialize["cle_privee_pem"] = o.ClePriveePem
	if o.Pkcs12Base64.IsSet() {
		toSerialize["pkcs12_base64"] = o.Pkcs12Base64.Get()
	}
	toSerialize["info"] = o.Info
	if !IsNil(o.Avertissement) {
		toSerialize["avertissement"] = o.Avertissement
	}
	return toSerialize, nil
}

func (o *GenerateCertificateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"certificat_pem",
		"cle_privee_pem",
		"info",
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

	varGenerateCertificateResponse := _GenerateCertificateResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenerateCertificateResponse)

	if err != nil {
		return err
	}

	*o = GenerateCertificateResponse(varGenerateCertificateResponse)

	return err
}

type NullableGenerateCertificateResponse struct {
	value *GenerateCertificateResponse
	isSet bool
}

func (v NullableGenerateCertificateResponse) Get() *GenerateCertificateResponse {
	return v.value
}

func (v *NullableGenerateCertificateResponse) Set(val *GenerateCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateCertificateResponse(val *GenerateCertificateResponse) *NullableGenerateCertificateResponse {
	return &NullableGenerateCertificateResponse{value: val, isSet: true}
}

func (v NullableGenerateCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


