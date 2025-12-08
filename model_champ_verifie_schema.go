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

// checks if the ChampVerifieSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChampVerifieSchema{}

// ChampVerifieSchema Un champ vérifié avec toutes ses informations (extraction + conformité + localisation).
type ChampVerifieSchema struct {
	// Business Term EN16931 (ex: BT-1)
	BusinessTerm string `json:"business_term"`
	// Libellé du champ (ex: N° Facture)
	Label string `json:"label"`
	ValeurPdf NullableString `json:"valeur_pdf,omitempty"`
	ValeurXml NullableString `json:"valeur_xml,omitempty"`
	// Statut de conformité
	Statut StatutChampAPI `json:"statut"`
	Message NullableString `json:"message,omitempty"`
	// Score de confiance (0-1)
	Confiance *float32 `json:"confiance,omitempty"`
	// Source d'extraction
	Source *string `json:"source,omitempty"`
	Bbox NullableBoundingBoxSchema `json:"bbox,omitempty"`
}

type _ChampVerifieSchema ChampVerifieSchema

// NewChampVerifieSchema instantiates a new ChampVerifieSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChampVerifieSchema(businessTerm string, label string, statut StatutChampAPI) *ChampVerifieSchema {
	this := ChampVerifieSchema{}
	this.BusinessTerm = businessTerm
	this.Label = label
	this.Statut = statut
	var confiance float32 = 1.0
	this.Confiance = &confiance
	var source string = "pdf_natif"
	this.Source = &source
	return &this
}

// NewChampVerifieSchemaWithDefaults instantiates a new ChampVerifieSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChampVerifieSchemaWithDefaults() *ChampVerifieSchema {
	this := ChampVerifieSchema{}
	var confiance float32 = 1.0
	this.Confiance = &confiance
	var source string = "pdf_natif"
	this.Source = &source
	return &this
}

// GetBusinessTerm returns the BusinessTerm field value
func (o *ChampVerifieSchema) GetBusinessTerm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessTerm
}

// GetBusinessTermOk returns a tuple with the BusinessTerm field value
// and a boolean to check if the value has been set.
func (o *ChampVerifieSchema) GetBusinessTermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessTerm, true
}

// SetBusinessTerm sets field value
func (o *ChampVerifieSchema) SetBusinessTerm(v string) {
	o.BusinessTerm = v
}

// GetLabel returns the Label field value
func (o *ChampVerifieSchema) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ChampVerifieSchema) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ChampVerifieSchema) SetLabel(v string) {
	o.Label = v
}

// GetValeurPdf returns the ValeurPdf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChampVerifieSchema) GetValeurPdf() string {
	if o == nil || IsNil(o.ValeurPdf.Get()) {
		var ret string
		return ret
	}
	return *o.ValeurPdf.Get()
}

// GetValeurPdfOk returns a tuple with the ValeurPdf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChampVerifieSchema) GetValeurPdfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValeurPdf.Get(), o.ValeurPdf.IsSet()
}

// HasValeurPdf returns a boolean if a field has been set.
func (o *ChampVerifieSchema) HasValeurPdf() bool {
	if o != nil && o.ValeurPdf.IsSet() {
		return true
	}

	return false
}

// SetValeurPdf gets a reference to the given NullableString and assigns it to the ValeurPdf field.
func (o *ChampVerifieSchema) SetValeurPdf(v string) {
	o.ValeurPdf.Set(&v)
}
// SetValeurPdfNil sets the value for ValeurPdf to be an explicit nil
func (o *ChampVerifieSchema) SetValeurPdfNil() {
	o.ValeurPdf.Set(nil)
}

// UnsetValeurPdf ensures that no value is present for ValeurPdf, not even an explicit nil
func (o *ChampVerifieSchema) UnsetValeurPdf() {
	o.ValeurPdf.Unset()
}

// GetValeurXml returns the ValeurXml field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChampVerifieSchema) GetValeurXml() string {
	if o == nil || IsNil(o.ValeurXml.Get()) {
		var ret string
		return ret
	}
	return *o.ValeurXml.Get()
}

// GetValeurXmlOk returns a tuple with the ValeurXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChampVerifieSchema) GetValeurXmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValeurXml.Get(), o.ValeurXml.IsSet()
}

// HasValeurXml returns a boolean if a field has been set.
func (o *ChampVerifieSchema) HasValeurXml() bool {
	if o != nil && o.ValeurXml.IsSet() {
		return true
	}

	return false
}

// SetValeurXml gets a reference to the given NullableString and assigns it to the ValeurXml field.
func (o *ChampVerifieSchema) SetValeurXml(v string) {
	o.ValeurXml.Set(&v)
}
// SetValeurXmlNil sets the value for ValeurXml to be an explicit nil
func (o *ChampVerifieSchema) SetValeurXmlNil() {
	o.ValeurXml.Set(nil)
}

// UnsetValeurXml ensures that no value is present for ValeurXml, not even an explicit nil
func (o *ChampVerifieSchema) UnsetValeurXml() {
	o.ValeurXml.Unset()
}

// GetStatut returns the Statut field value
func (o *ChampVerifieSchema) GetStatut() StatutChampAPI {
	if o == nil {
		var ret StatutChampAPI
		return ret
	}

	return o.Statut
}

// GetStatutOk returns a tuple with the Statut field value
// and a boolean to check if the value has been set.
func (o *ChampVerifieSchema) GetStatutOk() (*StatutChampAPI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statut, true
}

// SetStatut sets field value
func (o *ChampVerifieSchema) SetStatut(v StatutChampAPI) {
	o.Statut = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChampVerifieSchema) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChampVerifieSchema) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ChampVerifieSchema) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ChampVerifieSchema) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *ChampVerifieSchema) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ChampVerifieSchema) UnsetMessage() {
	o.Message.Unset()
}

// GetConfiance returns the Confiance field value if set, zero value otherwise.
func (o *ChampVerifieSchema) GetConfiance() float32 {
	if o == nil || IsNil(o.Confiance) {
		var ret float32
		return ret
	}
	return *o.Confiance
}

// GetConfianceOk returns a tuple with the Confiance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChampVerifieSchema) GetConfianceOk() (*float32, bool) {
	if o == nil || IsNil(o.Confiance) {
		return nil, false
	}
	return o.Confiance, true
}

// HasConfiance returns a boolean if a field has been set.
func (o *ChampVerifieSchema) HasConfiance() bool {
	if o != nil && !IsNil(o.Confiance) {
		return true
	}

	return false
}

// SetConfiance gets a reference to the given float32 and assigns it to the Confiance field.
func (o *ChampVerifieSchema) SetConfiance(v float32) {
	o.Confiance = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ChampVerifieSchema) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChampVerifieSchema) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ChampVerifieSchema) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ChampVerifieSchema) SetSource(v string) {
	o.Source = &v
}

// GetBbox returns the Bbox field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChampVerifieSchema) GetBbox() BoundingBoxSchema {
	if o == nil || IsNil(o.Bbox.Get()) {
		var ret BoundingBoxSchema
		return ret
	}
	return *o.Bbox.Get()
}

// GetBboxOk returns a tuple with the Bbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChampVerifieSchema) GetBboxOk() (*BoundingBoxSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bbox.Get(), o.Bbox.IsSet()
}

// HasBbox returns a boolean if a field has been set.
func (o *ChampVerifieSchema) HasBbox() bool {
	if o != nil && o.Bbox.IsSet() {
		return true
	}

	return false
}

// SetBbox gets a reference to the given NullableBoundingBoxSchema and assigns it to the Bbox field.
func (o *ChampVerifieSchema) SetBbox(v BoundingBoxSchema) {
	o.Bbox.Set(&v)
}
// SetBboxNil sets the value for Bbox to be an explicit nil
func (o *ChampVerifieSchema) SetBboxNil() {
	o.Bbox.Set(nil)
}

// UnsetBbox ensures that no value is present for Bbox, not even an explicit nil
func (o *ChampVerifieSchema) UnsetBbox() {
	o.Bbox.Unset()
}

func (o ChampVerifieSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChampVerifieSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_term"] = o.BusinessTerm
	toSerialize["label"] = o.Label
	if o.ValeurPdf.IsSet() {
		toSerialize["valeur_pdf"] = o.ValeurPdf.Get()
	}
	if o.ValeurXml.IsSet() {
		toSerialize["valeur_xml"] = o.ValeurXml.Get()
	}
	toSerialize["statut"] = o.Statut
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if !IsNil(o.Confiance) {
		toSerialize["confiance"] = o.Confiance
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if o.Bbox.IsSet() {
		toSerialize["bbox"] = o.Bbox.Get()
	}
	return toSerialize, nil
}

func (o *ChampVerifieSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_term",
		"label",
		"statut",
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

	varChampVerifieSchema := _ChampVerifieSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChampVerifieSchema)

	if err != nil {
		return err
	}

	*o = ChampVerifieSchema(varChampVerifieSchema)

	return err
}

type NullableChampVerifieSchema struct {
	value *ChampVerifieSchema
	isSet bool
}

func (v NullableChampVerifieSchema) Get() *ChampVerifieSchema {
	return v.value
}

func (v *NullableChampVerifieSchema) Set(val *ChampVerifieSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableChampVerifieSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableChampVerifieSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChampVerifieSchema(val *ChampVerifieSchema) *NullableChampVerifieSchema {
	return &NullableChampVerifieSchema{value: val, isSet: true}
}

func (v NullableChampVerifieSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChampVerifieSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


