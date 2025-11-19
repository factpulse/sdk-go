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

// checks if the PDFFacturXInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PDFFacturXInfo{}

// PDFFacturXInfo Informations sur le PDF Factur-X généré.
type PDFFacturXInfo struct {
	// Taille du PDF en octets
	Taille int32 `json:"taille"`
	// Profil Factur-X utilisé
	Profil string `json:"profil"`
	// PDF signé électroniquement
	Signe *bool `json:"signe,omitempty"`
}

type _PDFFacturXInfo PDFFacturXInfo

// NewPDFFacturXInfo instantiates a new PDFFacturXInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPDFFacturXInfo(taille int32, profil string) *PDFFacturXInfo {
	this := PDFFacturXInfo{}
	this.Taille = taille
	this.Profil = profil
	var signe bool = false
	this.Signe = &signe
	return &this
}

// NewPDFFacturXInfoWithDefaults instantiates a new PDFFacturXInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPDFFacturXInfoWithDefaults() *PDFFacturXInfo {
	this := PDFFacturXInfo{}
	var signe bool = false
	this.Signe = &signe
	return &this
}

// GetTaille returns the Taille field value
func (o *PDFFacturXInfo) GetTaille() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Taille
}

// GetTailleOk returns a tuple with the Taille field value
// and a boolean to check if the value has been set.
func (o *PDFFacturXInfo) GetTailleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Taille, true
}

// SetTaille sets field value
func (o *PDFFacturXInfo) SetTaille(v int32) {
	o.Taille = v
}

// GetProfil returns the Profil field value
func (o *PDFFacturXInfo) GetProfil() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profil
}

// GetProfilOk returns a tuple with the Profil field value
// and a boolean to check if the value has been set.
func (o *PDFFacturXInfo) GetProfilOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profil, true
}

// SetProfil sets field value
func (o *PDFFacturXInfo) SetProfil(v string) {
	o.Profil = v
}

// GetSigne returns the Signe field value if set, zero value otherwise.
func (o *PDFFacturXInfo) GetSigne() bool {
	if o == nil || IsNil(o.Signe) {
		var ret bool
		return ret
	}
	return *o.Signe
}

// GetSigneOk returns a tuple with the Signe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDFFacturXInfo) GetSigneOk() (*bool, bool) {
	if o == nil || IsNil(o.Signe) {
		return nil, false
	}
	return o.Signe, true
}

// HasSigne returns a boolean if a field has been set.
func (o *PDFFacturXInfo) HasSigne() bool {
	if o != nil && !IsNil(o.Signe) {
		return true
	}

	return false
}

// SetSigne gets a reference to the given bool and assigns it to the Signe field.
func (o *PDFFacturXInfo) SetSigne(v bool) {
	o.Signe = &v
}

func (o PDFFacturXInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PDFFacturXInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["taille"] = o.Taille
	toSerialize["profil"] = o.Profil
	if !IsNil(o.Signe) {
		toSerialize["signe"] = o.Signe
	}
	return toSerialize, nil
}

func (o *PDFFacturXInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"taille",
		"profil",
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

	varPDFFacturXInfo := _PDFFacturXInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPDFFacturXInfo)

	if err != nil {
		return err
	}

	*o = PDFFacturXInfo(varPDFFacturXInfo)

	return err
}

type NullablePDFFacturXInfo struct {
	value *PDFFacturXInfo
	isSet bool
}

func (v NullablePDFFacturXInfo) Get() *PDFFacturXInfo {
	return v.value
}

func (v *NullablePDFFacturXInfo) Set(val *PDFFacturXInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePDFFacturXInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePDFFacturXInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePDFFacturXInfo(val *PDFFacturXInfo) *NullablePDFFacturXInfo {
	return &NullablePDFFacturXInfo{value: val, isSet: true}
}

func (v NullablePDFFacturXInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePDFFacturXInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


