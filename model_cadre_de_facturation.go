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

// checks if the CadreDeFacturation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CadreDeFacturation{}

// CadreDeFacturation Définit le cadre de facturation (ex: A1 pour une facture fournisseur).
type CadreDeFacturation struct {
	CodeCadreFacturation CodeCadreFacturation `json:"codeCadreFacturation"`
	CodeServiceValideur NullableString `json:"codeServiceValideur,omitempty"`
	CodeStructureValideur NullableString `json:"codeStructureValideur,omitempty"`
}

type _CadreDeFacturation CadreDeFacturation

// NewCadreDeFacturation instantiates a new CadreDeFacturation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCadreDeFacturation(codeCadreFacturation CodeCadreFacturation) *CadreDeFacturation {
	this := CadreDeFacturation{}
	this.CodeCadreFacturation = codeCadreFacturation
	return &this
}

// NewCadreDeFacturationWithDefaults instantiates a new CadreDeFacturation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCadreDeFacturationWithDefaults() *CadreDeFacturation {
	this := CadreDeFacturation{}
	return &this
}

// GetCodeCadreFacturation returns the CodeCadreFacturation field value
func (o *CadreDeFacturation) GetCodeCadreFacturation() CodeCadreFacturation {
	if o == nil {
		var ret CodeCadreFacturation
		return ret
	}

	return o.CodeCadreFacturation
}

// GetCodeCadreFacturationOk returns a tuple with the CodeCadreFacturation field value
// and a boolean to check if the value has been set.
func (o *CadreDeFacturation) GetCodeCadreFacturationOk() (*CodeCadreFacturation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeCadreFacturation, true
}

// SetCodeCadreFacturation sets field value
func (o *CadreDeFacturation) SetCodeCadreFacturation(v CodeCadreFacturation) {
	o.CodeCadreFacturation = v
}

// GetCodeServiceValideur returns the CodeServiceValideur field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CadreDeFacturation) GetCodeServiceValideur() string {
	if o == nil || IsNil(o.CodeServiceValideur.Get()) {
		var ret string
		return ret
	}
	return *o.CodeServiceValideur.Get()
}

// GetCodeServiceValideurOk returns a tuple with the CodeServiceValideur field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CadreDeFacturation) GetCodeServiceValideurOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeServiceValideur.Get(), o.CodeServiceValideur.IsSet()
}

// HasCodeServiceValideur returns a boolean if a field has been set.
func (o *CadreDeFacturation) HasCodeServiceValideur() bool {
	if o != nil && o.CodeServiceValideur.IsSet() {
		return true
	}

	return false
}

// SetCodeServiceValideur gets a reference to the given NullableString and assigns it to the CodeServiceValideur field.
func (o *CadreDeFacturation) SetCodeServiceValideur(v string) {
	o.CodeServiceValideur.Set(&v)
}
// SetCodeServiceValideurNil sets the value for CodeServiceValideur to be an explicit nil
func (o *CadreDeFacturation) SetCodeServiceValideurNil() {
	o.CodeServiceValideur.Set(nil)
}

// UnsetCodeServiceValideur ensures that no value is present for CodeServiceValideur, not even an explicit nil
func (o *CadreDeFacturation) UnsetCodeServiceValideur() {
	o.CodeServiceValideur.Unset()
}

// GetCodeStructureValideur returns the CodeStructureValideur field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CadreDeFacturation) GetCodeStructureValideur() string {
	if o == nil || IsNil(o.CodeStructureValideur.Get()) {
		var ret string
		return ret
	}
	return *o.CodeStructureValideur.Get()
}

// GetCodeStructureValideurOk returns a tuple with the CodeStructureValideur field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CadreDeFacturation) GetCodeStructureValideurOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeStructureValideur.Get(), o.CodeStructureValideur.IsSet()
}

// HasCodeStructureValideur returns a boolean if a field has been set.
func (o *CadreDeFacturation) HasCodeStructureValideur() bool {
	if o != nil && o.CodeStructureValideur.IsSet() {
		return true
	}

	return false
}

// SetCodeStructureValideur gets a reference to the given NullableString and assigns it to the CodeStructureValideur field.
func (o *CadreDeFacturation) SetCodeStructureValideur(v string) {
	o.CodeStructureValideur.Set(&v)
}
// SetCodeStructureValideurNil sets the value for CodeStructureValideur to be an explicit nil
func (o *CadreDeFacturation) SetCodeStructureValideurNil() {
	o.CodeStructureValideur.Set(nil)
}

// UnsetCodeStructureValideur ensures that no value is present for CodeStructureValideur, not even an explicit nil
func (o *CadreDeFacturation) UnsetCodeStructureValideur() {
	o.CodeStructureValideur.Unset()
}

func (o CadreDeFacturation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CadreDeFacturation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["codeCadreFacturation"] = o.CodeCadreFacturation
	if o.CodeServiceValideur.IsSet() {
		toSerialize["codeServiceValideur"] = o.CodeServiceValideur.Get()
	}
	if o.CodeStructureValideur.IsSet() {
		toSerialize["codeStructureValideur"] = o.CodeStructureValideur.Get()
	}
	return toSerialize, nil
}

func (o *CadreDeFacturation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"codeCadreFacturation",
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

	varCadreDeFacturation := _CadreDeFacturation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCadreDeFacturation)

	if err != nil {
		return err
	}

	*o = CadreDeFacturation(varCadreDeFacturation)

	return err
}

type NullableCadreDeFacturation struct {
	value *CadreDeFacturation
	isSet bool
}

func (v NullableCadreDeFacturation) Get() *CadreDeFacturation {
	return v.value
}

func (v *NullableCadreDeFacturation) Set(val *CadreDeFacturation) {
	v.value = val
	v.isSet = true
}

func (v NullableCadreDeFacturation) IsSet() bool {
	return v.isSet
}

func (v *NullableCadreDeFacturation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCadreDeFacturation(val *CadreDeFacturation) *NullableCadreDeFacturation {
	return &NullableCadreDeFacturation{value: val, isSet: true}
}

func (v NullableCadreDeFacturation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCadreDeFacturation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


