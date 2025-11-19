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

// checks if the SoumettreFactureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoumettreFactureResponse{}

// SoumettreFactureResponse Réponse après soumission de facture.
type SoumettreFactureResponse struct {
	// Code retour (0 = succès)
	CodeRetour int32 `json:"code_retour"`
	// Message de retour
	Libelle string `json:"libelle"`
	IdentifiantFactureCpp NullableInt32 `json:"identifiant_facture_cpp,omitempty"`
	NumeroFluxDepot NullableString `json:"numero_flux_depot,omitempty"`
}

type _SoumettreFactureResponse SoumettreFactureResponse

// NewSoumettreFactureResponse instantiates a new SoumettreFactureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoumettreFactureResponse(codeRetour int32, libelle string) *SoumettreFactureResponse {
	this := SoumettreFactureResponse{}
	this.CodeRetour = codeRetour
	this.Libelle = libelle
	return &this
}

// NewSoumettreFactureResponseWithDefaults instantiates a new SoumettreFactureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoumettreFactureResponseWithDefaults() *SoumettreFactureResponse {
	this := SoumettreFactureResponse{}
	return &this
}

// GetCodeRetour returns the CodeRetour field value
func (o *SoumettreFactureResponse) GetCodeRetour() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CodeRetour
}

// GetCodeRetourOk returns a tuple with the CodeRetour field value
// and a boolean to check if the value has been set.
func (o *SoumettreFactureResponse) GetCodeRetourOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeRetour, true
}

// SetCodeRetour sets field value
func (o *SoumettreFactureResponse) SetCodeRetour(v int32) {
	o.CodeRetour = v
}

// GetLibelle returns the Libelle field value
func (o *SoumettreFactureResponse) GetLibelle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Libelle
}

// GetLibelleOk returns a tuple with the Libelle field value
// and a boolean to check if the value has been set.
func (o *SoumettreFactureResponse) GetLibelleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Libelle, true
}

// SetLibelle sets field value
func (o *SoumettreFactureResponse) SetLibelle(v string) {
	o.Libelle = v
}

// GetIdentifiantFactureCpp returns the IdentifiantFactureCpp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureResponse) GetIdentifiantFactureCpp() int32 {
	if o == nil || IsNil(o.IdentifiantFactureCpp.Get()) {
		var ret int32
		return ret
	}
	return *o.IdentifiantFactureCpp.Get()
}

// GetIdentifiantFactureCppOk returns a tuple with the IdentifiantFactureCpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureResponse) GetIdentifiantFactureCppOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentifiantFactureCpp.Get(), o.IdentifiantFactureCpp.IsSet()
}

// HasIdentifiantFactureCpp returns a boolean if a field has been set.
func (o *SoumettreFactureResponse) HasIdentifiantFactureCpp() bool {
	if o != nil && o.IdentifiantFactureCpp.IsSet() {
		return true
	}

	return false
}

// SetIdentifiantFactureCpp gets a reference to the given NullableInt32 and assigns it to the IdentifiantFactureCpp field.
func (o *SoumettreFactureResponse) SetIdentifiantFactureCpp(v int32) {
	o.IdentifiantFactureCpp.Set(&v)
}
// SetIdentifiantFactureCppNil sets the value for IdentifiantFactureCpp to be an explicit nil
func (o *SoumettreFactureResponse) SetIdentifiantFactureCppNil() {
	o.IdentifiantFactureCpp.Set(nil)
}

// UnsetIdentifiantFactureCpp ensures that no value is present for IdentifiantFactureCpp, not even an explicit nil
func (o *SoumettreFactureResponse) UnsetIdentifiantFactureCpp() {
	o.IdentifiantFactureCpp.Unset()
}

// GetNumeroFluxDepot returns the NumeroFluxDepot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureResponse) GetNumeroFluxDepot() string {
	if o == nil || IsNil(o.NumeroFluxDepot.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroFluxDepot.Get()
}

// GetNumeroFluxDepotOk returns a tuple with the NumeroFluxDepot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureResponse) GetNumeroFluxDepotOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroFluxDepot.Get(), o.NumeroFluxDepot.IsSet()
}

// HasNumeroFluxDepot returns a boolean if a field has been set.
func (o *SoumettreFactureResponse) HasNumeroFluxDepot() bool {
	if o != nil && o.NumeroFluxDepot.IsSet() {
		return true
	}

	return false
}

// SetNumeroFluxDepot gets a reference to the given NullableString and assigns it to the NumeroFluxDepot field.
func (o *SoumettreFactureResponse) SetNumeroFluxDepot(v string) {
	o.NumeroFluxDepot.Set(&v)
}
// SetNumeroFluxDepotNil sets the value for NumeroFluxDepot to be an explicit nil
func (o *SoumettreFactureResponse) SetNumeroFluxDepotNil() {
	o.NumeroFluxDepot.Set(nil)
}

// UnsetNumeroFluxDepot ensures that no value is present for NumeroFluxDepot, not even an explicit nil
func (o *SoumettreFactureResponse) UnsetNumeroFluxDepot() {
	o.NumeroFluxDepot.Unset()
}

func (o SoumettreFactureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoumettreFactureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code_retour"] = o.CodeRetour
	toSerialize["libelle"] = o.Libelle
	if o.IdentifiantFactureCpp.IsSet() {
		toSerialize["identifiant_facture_cpp"] = o.IdentifiantFactureCpp.Get()
	}
	if o.NumeroFluxDepot.IsSet() {
		toSerialize["numero_flux_depot"] = o.NumeroFluxDepot.Get()
	}
	return toSerialize, nil
}

func (o *SoumettreFactureResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code_retour",
		"libelle",
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

	varSoumettreFactureResponse := _SoumettreFactureResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSoumettreFactureResponse)

	if err != nil {
		return err
	}

	*o = SoumettreFactureResponse(varSoumettreFactureResponse)

	return err
}

type NullableSoumettreFactureResponse struct {
	value *SoumettreFactureResponse
	isSet bool
}

func (v NullableSoumettreFactureResponse) Get() *SoumettreFactureResponse {
	return v.value
}

func (v *NullableSoumettreFactureResponse) Set(val *SoumettreFactureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSoumettreFactureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSoumettreFactureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoumettreFactureResponse(val *SoumettreFactureResponse) *NullableSoumettreFactureResponse {
	return &NullableSoumettreFactureResponse{value: val, isSet: true}
}

func (v NullableSoumettreFactureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoumettreFactureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


