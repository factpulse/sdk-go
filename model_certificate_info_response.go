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

// checks if the CertificateInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateInfoResponse{}

// CertificateInfoResponse Informations sur un certificat généré.
type CertificateInfoResponse struct {
	// Common Name
	Cn string `json:"cn"`
	// Organisation
	Organisation string `json:"organisation"`
	// Code pays
	Pays string `json:"pays"`
	// Ville
	Ville string `json:"ville"`
	// Province
	Province string `json:"province"`
	Email NullableString `json:"email,omitempty"`
	// Sujet complet (RFC4514)
	Sujet string `json:"sujet"`
	// Émetteur (auto-signé = même que sujet)
	Emetteur string `json:"emetteur"`
	// Numéro de série du certificat
	NumeroSerie int32 `json:"numero_serie"`
	// Date de début de validité (ISO 8601)
	ValideDu string `json:"valide_du"`
	// Date de fin de validité (ISO 8601)
	ValideAu string `json:"valide_au"`
	// Algorithme de signature
	Algorithme string `json:"algorithme"`
}

type _CertificateInfoResponse CertificateInfoResponse

// NewCertificateInfoResponse instantiates a new CertificateInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateInfoResponse(cn string, organisation string, pays string, ville string, province string, sujet string, emetteur string, numeroSerie int32, valideDu string, valideAu string, algorithme string) *CertificateInfoResponse {
	this := CertificateInfoResponse{}
	this.Cn = cn
	this.Organisation = organisation
	this.Pays = pays
	this.Ville = ville
	this.Province = province
	this.Sujet = sujet
	this.Emetteur = emetteur
	this.NumeroSerie = numeroSerie
	this.ValideDu = valideDu
	this.ValideAu = valideAu
	this.Algorithme = algorithme
	return &this
}

// NewCertificateInfoResponseWithDefaults instantiates a new CertificateInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateInfoResponseWithDefaults() *CertificateInfoResponse {
	this := CertificateInfoResponse{}
	return &this
}

// GetCn returns the Cn field value
func (o *CertificateInfoResponse) GetCn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cn
}

// GetCnOk returns a tuple with the Cn field value
// and a boolean to check if the value has been set.
func (o *CertificateInfoResponse) GetCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cn, true
}

// SetCn sets field value
func (o *CertificateInfoResponse) SetCn(v string) {
	o.Cn = v
}

// GetOrganisation returns the Organisation field value
func (o *CertificateInfoResponse) GetOrganisation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organisation
}

// GetOrganisationOk returns a tuple with the Organisation field value
// and a boolean to check if the value has been set.
func (o *CertificateInfoResponse) GetOrganisationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organisation, true
}

// SetOrganisation sets field value
func (o *CertificateInfoResponse) SetOrganisation(v string) {
	o.Organisation = v
}

// GetPays returns the Pays field value
func (o *CertificateInfoResponse) GetPays() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pays
}

// GetPaysOk returns a tuple with the Pays field value
// and a boolean to check if the value has been set.
func (o *CertificateInfoResponse) GetPaysOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pays, true
}

// SetPays sets field value
func (o *CertificateInfoResponse) SetPays(v string) {
	o.Pays = v
}

// GetVille returns the Ville field value
func (o *CertificateInfoResponse) GetVille() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ville
}

// GetVilleOk returns a tuple with the Ville field value
// and a boolean to check if the value has been set.
func (o *CertificateInfoResponse) GetVilleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ville, true
}

// SetVille sets field value
func (o *CertificateInfoResponse) SetVille(v string) {
	o.Ville = v
}

// GetProvince returns the Province field value
func (o *CertificateInfoResponse) GetProvince() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Province
}

// GetProvinceOk returns a tuple with the Province field value
// and a boolean to check if the value has been set.
func (o *CertificateInfoResponse) GetProvinceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Province, true
}

// SetProvince sets field value
func (o *CertificateInfoResponse) SetProvince(v string) {
	o.Province = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateInfoResponse) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateInfoResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *CertificateInfoResponse) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *CertificateInfoResponse) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *CertificateInfoResponse) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *CertificateInfoResponse) UnsetEmail() {
	o.Email.Unset()
}

// GetSujet returns the Sujet field value
func (o *CertificateInfoResponse) GetSujet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sujet
}

// GetSujetOk returns a tuple with the Sujet field value
// and a boolean to check if the value has been set.
func (o *CertificateInfoResponse) GetSujetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sujet, true
}

// SetSujet sets field value
func (o *CertificateInfoResponse) SetSujet(v string) {
	o.Sujet = v
}

// GetEmetteur returns the Emetteur field value
func (o *CertificateInfoResponse) GetEmetteur() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Emetteur
}

// GetEmetteurOk returns a tuple with the Emetteur field value
// and a boolean to check if the value has been set.
func (o *CertificateInfoResponse) GetEmetteurOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Emetteur, true
}

// SetEmetteur sets field value
func (o *CertificateInfoResponse) SetEmetteur(v string) {
	o.Emetteur = v
}

// GetNumeroSerie returns the NumeroSerie field value
func (o *CertificateInfoResponse) GetNumeroSerie() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumeroSerie
}

// GetNumeroSerieOk returns a tuple with the NumeroSerie field value
// and a boolean to check if the value has been set.
func (o *CertificateInfoResponse) GetNumeroSerieOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumeroSerie, true
}

// SetNumeroSerie sets field value
func (o *CertificateInfoResponse) SetNumeroSerie(v int32) {
	o.NumeroSerie = v
}

// GetValideDu returns the ValideDu field value
func (o *CertificateInfoResponse) GetValideDu() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValideDu
}

// GetValideDuOk returns a tuple with the ValideDu field value
// and a boolean to check if the value has been set.
func (o *CertificateInfoResponse) GetValideDuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValideDu, true
}

// SetValideDu sets field value
func (o *CertificateInfoResponse) SetValideDu(v string) {
	o.ValideDu = v
}

// GetValideAu returns the ValideAu field value
func (o *CertificateInfoResponse) GetValideAu() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValideAu
}

// GetValideAuOk returns a tuple with the ValideAu field value
// and a boolean to check if the value has been set.
func (o *CertificateInfoResponse) GetValideAuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValideAu, true
}

// SetValideAu sets field value
func (o *CertificateInfoResponse) SetValideAu(v string) {
	o.ValideAu = v
}

// GetAlgorithme returns the Algorithme field value
func (o *CertificateInfoResponse) GetAlgorithme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithme
}

// GetAlgorithmeOk returns a tuple with the Algorithme field value
// and a boolean to check if the value has been set.
func (o *CertificateInfoResponse) GetAlgorithmeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithme, true
}

// SetAlgorithme sets field value
func (o *CertificateInfoResponse) SetAlgorithme(v string) {
	o.Algorithme = v
}

func (o CertificateInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cn"] = o.Cn
	toSerialize["organisation"] = o.Organisation
	toSerialize["pays"] = o.Pays
	toSerialize["ville"] = o.Ville
	toSerialize["province"] = o.Province
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	toSerialize["sujet"] = o.Sujet
	toSerialize["emetteur"] = o.Emetteur
	toSerialize["numero_serie"] = o.NumeroSerie
	toSerialize["valide_du"] = o.ValideDu
	toSerialize["valide_au"] = o.ValideAu
	toSerialize["algorithme"] = o.Algorithme
	return toSerialize, nil
}

func (o *CertificateInfoResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cn",
		"organisation",
		"pays",
		"ville",
		"province",
		"sujet",
		"emetteur",
		"numero_serie",
		"valide_du",
		"valide_au",
		"algorithme",
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

	varCertificateInfoResponse := _CertificateInfoResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCertificateInfoResponse)

	if err != nil {
		return err
	}

	*o = CertificateInfoResponse(varCertificateInfoResponse)

	return err
}

type NullableCertificateInfoResponse struct {
	value *CertificateInfoResponse
	isSet bool
}

func (v NullableCertificateInfoResponse) Get() *CertificateInfoResponse {
	return v.value
}

func (v *NullableCertificateInfoResponse) Set(val *CertificateInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateInfoResponse(val *CertificateInfoResponse) *NullableCertificateInfoResponse {
	return &NullableCertificateInfoResponse{value: val, isSet: true}
}

func (v NullableCertificateInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


