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

// checks if the ConsulterStructureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsulterStructureResponse{}

// ConsulterStructureResponse Détails d'une structure.
type ConsulterStructureResponse struct {
	CodeRetour int32 `json:"code_retour"`
	Libelle string `json:"libelle"`
	IdStructureCpp NullableInt32 `json:"id_structure_cpp,omitempty"`
	IdentifiantStructure NullableString `json:"identifiant_structure,omitempty"`
	LibelleStructure NullableString `json:"libelle_structure,omitempty"`
	RaisonSocialeStructure NullableString `json:"raison_sociale_structure,omitempty"`
	NumeroTva NullableString `json:"numero_tva,omitempty"`
	EmailStructure NullableString `json:"email_structure,omitempty"`
	Parametres NullableParametresStructure `json:"parametres,omitempty"`
}

type _ConsulterStructureResponse ConsulterStructureResponse

// NewConsulterStructureResponse instantiates a new ConsulterStructureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsulterStructureResponse(codeRetour int32, libelle string) *ConsulterStructureResponse {
	this := ConsulterStructureResponse{}
	this.CodeRetour = codeRetour
	this.Libelle = libelle
	return &this
}

// NewConsulterStructureResponseWithDefaults instantiates a new ConsulterStructureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulterStructureResponseWithDefaults() *ConsulterStructureResponse {
	this := ConsulterStructureResponse{}
	return &this
}

// GetCodeRetour returns the CodeRetour field value
func (o *ConsulterStructureResponse) GetCodeRetour() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CodeRetour
}

// GetCodeRetourOk returns a tuple with the CodeRetour field value
// and a boolean to check if the value has been set.
func (o *ConsulterStructureResponse) GetCodeRetourOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeRetour, true
}

// SetCodeRetour sets field value
func (o *ConsulterStructureResponse) SetCodeRetour(v int32) {
	o.CodeRetour = v
}

// GetLibelle returns the Libelle field value
func (o *ConsulterStructureResponse) GetLibelle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Libelle
}

// GetLibelleOk returns a tuple with the Libelle field value
// and a boolean to check if the value has been set.
func (o *ConsulterStructureResponse) GetLibelleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Libelle, true
}

// SetLibelle sets field value
func (o *ConsulterStructureResponse) SetLibelle(v string) {
	o.Libelle = v
}

// GetIdStructureCpp returns the IdStructureCpp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsulterStructureResponse) GetIdStructureCpp() int32 {
	if o == nil || IsNil(o.IdStructureCpp.Get()) {
		var ret int32
		return ret
	}
	return *o.IdStructureCpp.Get()
}

// GetIdStructureCppOk returns a tuple with the IdStructureCpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsulterStructureResponse) GetIdStructureCppOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdStructureCpp.Get(), o.IdStructureCpp.IsSet()
}

// HasIdStructureCpp returns a boolean if a field has been set.
func (o *ConsulterStructureResponse) HasIdStructureCpp() bool {
	if o != nil && o.IdStructureCpp.IsSet() {
		return true
	}

	return false
}

// SetIdStructureCpp gets a reference to the given NullableInt32 and assigns it to the IdStructureCpp field.
func (o *ConsulterStructureResponse) SetIdStructureCpp(v int32) {
	o.IdStructureCpp.Set(&v)
}
// SetIdStructureCppNil sets the value for IdStructureCpp to be an explicit nil
func (o *ConsulterStructureResponse) SetIdStructureCppNil() {
	o.IdStructureCpp.Set(nil)
}

// UnsetIdStructureCpp ensures that no value is present for IdStructureCpp, not even an explicit nil
func (o *ConsulterStructureResponse) UnsetIdStructureCpp() {
	o.IdStructureCpp.Unset()
}

// GetIdentifiantStructure returns the IdentifiantStructure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsulterStructureResponse) GetIdentifiantStructure() string {
	if o == nil || IsNil(o.IdentifiantStructure.Get()) {
		var ret string
		return ret
	}
	return *o.IdentifiantStructure.Get()
}

// GetIdentifiantStructureOk returns a tuple with the IdentifiantStructure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsulterStructureResponse) GetIdentifiantStructureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentifiantStructure.Get(), o.IdentifiantStructure.IsSet()
}

// HasIdentifiantStructure returns a boolean if a field has been set.
func (o *ConsulterStructureResponse) HasIdentifiantStructure() bool {
	if o != nil && o.IdentifiantStructure.IsSet() {
		return true
	}

	return false
}

// SetIdentifiantStructure gets a reference to the given NullableString and assigns it to the IdentifiantStructure field.
func (o *ConsulterStructureResponse) SetIdentifiantStructure(v string) {
	o.IdentifiantStructure.Set(&v)
}
// SetIdentifiantStructureNil sets the value for IdentifiantStructure to be an explicit nil
func (o *ConsulterStructureResponse) SetIdentifiantStructureNil() {
	o.IdentifiantStructure.Set(nil)
}

// UnsetIdentifiantStructure ensures that no value is present for IdentifiantStructure, not even an explicit nil
func (o *ConsulterStructureResponse) UnsetIdentifiantStructure() {
	o.IdentifiantStructure.Unset()
}

// GetLibelleStructure returns the LibelleStructure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsulterStructureResponse) GetLibelleStructure() string {
	if o == nil || IsNil(o.LibelleStructure.Get()) {
		var ret string
		return ret
	}
	return *o.LibelleStructure.Get()
}

// GetLibelleStructureOk returns a tuple with the LibelleStructure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsulterStructureResponse) GetLibelleStructureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LibelleStructure.Get(), o.LibelleStructure.IsSet()
}

// HasLibelleStructure returns a boolean if a field has been set.
func (o *ConsulterStructureResponse) HasLibelleStructure() bool {
	if o != nil && o.LibelleStructure.IsSet() {
		return true
	}

	return false
}

// SetLibelleStructure gets a reference to the given NullableString and assigns it to the LibelleStructure field.
func (o *ConsulterStructureResponse) SetLibelleStructure(v string) {
	o.LibelleStructure.Set(&v)
}
// SetLibelleStructureNil sets the value for LibelleStructure to be an explicit nil
func (o *ConsulterStructureResponse) SetLibelleStructureNil() {
	o.LibelleStructure.Set(nil)
}

// UnsetLibelleStructure ensures that no value is present for LibelleStructure, not even an explicit nil
func (o *ConsulterStructureResponse) UnsetLibelleStructure() {
	o.LibelleStructure.Unset()
}

// GetRaisonSocialeStructure returns the RaisonSocialeStructure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsulterStructureResponse) GetRaisonSocialeStructure() string {
	if o == nil || IsNil(o.RaisonSocialeStructure.Get()) {
		var ret string
		return ret
	}
	return *o.RaisonSocialeStructure.Get()
}

// GetRaisonSocialeStructureOk returns a tuple with the RaisonSocialeStructure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsulterStructureResponse) GetRaisonSocialeStructureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaisonSocialeStructure.Get(), o.RaisonSocialeStructure.IsSet()
}

// HasRaisonSocialeStructure returns a boolean if a field has been set.
func (o *ConsulterStructureResponse) HasRaisonSocialeStructure() bool {
	if o != nil && o.RaisonSocialeStructure.IsSet() {
		return true
	}

	return false
}

// SetRaisonSocialeStructure gets a reference to the given NullableString and assigns it to the RaisonSocialeStructure field.
func (o *ConsulterStructureResponse) SetRaisonSocialeStructure(v string) {
	o.RaisonSocialeStructure.Set(&v)
}
// SetRaisonSocialeStructureNil sets the value for RaisonSocialeStructure to be an explicit nil
func (o *ConsulterStructureResponse) SetRaisonSocialeStructureNil() {
	o.RaisonSocialeStructure.Set(nil)
}

// UnsetRaisonSocialeStructure ensures that no value is present for RaisonSocialeStructure, not even an explicit nil
func (o *ConsulterStructureResponse) UnsetRaisonSocialeStructure() {
	o.RaisonSocialeStructure.Unset()
}

// GetNumeroTva returns the NumeroTva field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsulterStructureResponse) GetNumeroTva() string {
	if o == nil || IsNil(o.NumeroTva.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroTva.Get()
}

// GetNumeroTvaOk returns a tuple with the NumeroTva field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsulterStructureResponse) GetNumeroTvaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroTva.Get(), o.NumeroTva.IsSet()
}

// HasNumeroTva returns a boolean if a field has been set.
func (o *ConsulterStructureResponse) HasNumeroTva() bool {
	if o != nil && o.NumeroTva.IsSet() {
		return true
	}

	return false
}

// SetNumeroTva gets a reference to the given NullableString and assigns it to the NumeroTva field.
func (o *ConsulterStructureResponse) SetNumeroTva(v string) {
	o.NumeroTva.Set(&v)
}
// SetNumeroTvaNil sets the value for NumeroTva to be an explicit nil
func (o *ConsulterStructureResponse) SetNumeroTvaNil() {
	o.NumeroTva.Set(nil)
}

// UnsetNumeroTva ensures that no value is present for NumeroTva, not even an explicit nil
func (o *ConsulterStructureResponse) UnsetNumeroTva() {
	o.NumeroTva.Unset()
}

// GetEmailStructure returns the EmailStructure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsulterStructureResponse) GetEmailStructure() string {
	if o == nil || IsNil(o.EmailStructure.Get()) {
		var ret string
		return ret
	}
	return *o.EmailStructure.Get()
}

// GetEmailStructureOk returns a tuple with the EmailStructure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsulterStructureResponse) GetEmailStructureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailStructure.Get(), o.EmailStructure.IsSet()
}

// HasEmailStructure returns a boolean if a field has been set.
func (o *ConsulterStructureResponse) HasEmailStructure() bool {
	if o != nil && o.EmailStructure.IsSet() {
		return true
	}

	return false
}

// SetEmailStructure gets a reference to the given NullableString and assigns it to the EmailStructure field.
func (o *ConsulterStructureResponse) SetEmailStructure(v string) {
	o.EmailStructure.Set(&v)
}
// SetEmailStructureNil sets the value for EmailStructure to be an explicit nil
func (o *ConsulterStructureResponse) SetEmailStructureNil() {
	o.EmailStructure.Set(nil)
}

// UnsetEmailStructure ensures that no value is present for EmailStructure, not even an explicit nil
func (o *ConsulterStructureResponse) UnsetEmailStructure() {
	o.EmailStructure.Unset()
}

// GetParametres returns the Parametres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsulterStructureResponse) GetParametres() ParametresStructure {
	if o == nil || IsNil(o.Parametres.Get()) {
		var ret ParametresStructure
		return ret
	}
	return *o.Parametres.Get()
}

// GetParametresOk returns a tuple with the Parametres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsulterStructureResponse) GetParametresOk() (*ParametresStructure, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parametres.Get(), o.Parametres.IsSet()
}

// HasParametres returns a boolean if a field has been set.
func (o *ConsulterStructureResponse) HasParametres() bool {
	if o != nil && o.Parametres.IsSet() {
		return true
	}

	return false
}

// SetParametres gets a reference to the given NullableParametresStructure and assigns it to the Parametres field.
func (o *ConsulterStructureResponse) SetParametres(v ParametresStructure) {
	o.Parametres.Set(&v)
}
// SetParametresNil sets the value for Parametres to be an explicit nil
func (o *ConsulterStructureResponse) SetParametresNil() {
	o.Parametres.Set(nil)
}

// UnsetParametres ensures that no value is present for Parametres, not even an explicit nil
func (o *ConsulterStructureResponse) UnsetParametres() {
	o.Parametres.Unset()
}

func (o ConsulterStructureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsulterStructureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code_retour"] = o.CodeRetour
	toSerialize["libelle"] = o.Libelle
	if o.IdStructureCpp.IsSet() {
		toSerialize["id_structure_cpp"] = o.IdStructureCpp.Get()
	}
	if o.IdentifiantStructure.IsSet() {
		toSerialize["identifiant_structure"] = o.IdentifiantStructure.Get()
	}
	if o.LibelleStructure.IsSet() {
		toSerialize["libelle_structure"] = o.LibelleStructure.Get()
	}
	if o.RaisonSocialeStructure.IsSet() {
		toSerialize["raison_sociale_structure"] = o.RaisonSocialeStructure.Get()
	}
	if o.NumeroTva.IsSet() {
		toSerialize["numero_tva"] = o.NumeroTva.Get()
	}
	if o.EmailStructure.IsSet() {
		toSerialize["email_structure"] = o.EmailStructure.Get()
	}
	if o.Parametres.IsSet() {
		toSerialize["parametres"] = o.Parametres.Get()
	}
	return toSerialize, nil
}

func (o *ConsulterStructureResponse) UnmarshalJSON(data []byte) (err error) {
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

	varConsulterStructureResponse := _ConsulterStructureResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConsulterStructureResponse)

	if err != nil {
		return err
	}

	*o = ConsulterStructureResponse(varConsulterStructureResponse)

	return err
}

type NullableConsulterStructureResponse struct {
	value *ConsulterStructureResponse
	isSet bool
}

func (v NullableConsulterStructureResponse) Get() *ConsulterStructureResponse {
	return v.value
}

func (v *NullableConsulterStructureResponse) Set(val *ConsulterStructureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConsulterStructureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConsulterStructureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsulterStructureResponse(val *ConsulterStructureResponse) *NullableConsulterStructureResponse {
	return &NullableConsulterStructureResponse{value: val, isSet: true}
}

func (v NullableConsulterStructureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsulterStructureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


