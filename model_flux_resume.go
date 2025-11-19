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

// checks if the FluxResume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FluxResume{}

// FluxResume Résumé d'un flux dans les résultats de recherche
type FluxResume struct {
	FlowId string `json:"flow_id"`
	TrackingId NullableString `json:"tracking_id,omitempty"`
	Nom string `json:"nom"`
	TypeFlux NullableString `json:"type_flux,omitempty"`
	DirectionFlux NullableString `json:"direction_flux,omitempty"`
	StatutAcquittement NullableString `json:"statut_acquittement,omitempty"`
	DateCreation NullableString `json:"date_creation,omitempty"`
	DateMaj NullableString `json:"date_maj,omitempty"`
}

type _FluxResume FluxResume

// NewFluxResume instantiates a new FluxResume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFluxResume(flowId string, nom string) *FluxResume {
	this := FluxResume{}
	this.FlowId = flowId
	this.Nom = nom
	return &this
}

// NewFluxResumeWithDefaults instantiates a new FluxResume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFluxResumeWithDefaults() *FluxResume {
	this := FluxResume{}
	return &this
}

// GetFlowId returns the FlowId field value
func (o *FluxResume) GetFlowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlowId
}

// GetFlowIdOk returns a tuple with the FlowId field value
// and a boolean to check if the value has been set.
func (o *FluxResume) GetFlowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlowId, true
}

// SetFlowId sets field value
func (o *FluxResume) SetFlowId(v string) {
	o.FlowId = v
}

// GetTrackingId returns the TrackingId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FluxResume) GetTrackingId() string {
	if o == nil || IsNil(o.TrackingId.Get()) {
		var ret string
		return ret
	}
	return *o.TrackingId.Get()
}

// GetTrackingIdOk returns a tuple with the TrackingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FluxResume) GetTrackingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingId.Get(), o.TrackingId.IsSet()
}

// HasTrackingId returns a boolean if a field has been set.
func (o *FluxResume) HasTrackingId() bool {
	if o != nil && o.TrackingId.IsSet() {
		return true
	}

	return false
}

// SetTrackingId gets a reference to the given NullableString and assigns it to the TrackingId field.
func (o *FluxResume) SetTrackingId(v string) {
	o.TrackingId.Set(&v)
}
// SetTrackingIdNil sets the value for TrackingId to be an explicit nil
func (o *FluxResume) SetTrackingIdNil() {
	o.TrackingId.Set(nil)
}

// UnsetTrackingId ensures that no value is present for TrackingId, not even an explicit nil
func (o *FluxResume) UnsetTrackingId() {
	o.TrackingId.Unset()
}

// GetNom returns the Nom field value
func (o *FluxResume) GetNom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nom
}

// GetNomOk returns a tuple with the Nom field value
// and a boolean to check if the value has been set.
func (o *FluxResume) GetNomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nom, true
}

// SetNom sets field value
func (o *FluxResume) SetNom(v string) {
	o.Nom = v
}

// GetTypeFlux returns the TypeFlux field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FluxResume) GetTypeFlux() string {
	if o == nil || IsNil(o.TypeFlux.Get()) {
		var ret string
		return ret
	}
	return *o.TypeFlux.Get()
}

// GetTypeFluxOk returns a tuple with the TypeFlux field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FluxResume) GetTypeFluxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeFlux.Get(), o.TypeFlux.IsSet()
}

// HasTypeFlux returns a boolean if a field has been set.
func (o *FluxResume) HasTypeFlux() bool {
	if o != nil && o.TypeFlux.IsSet() {
		return true
	}

	return false
}

// SetTypeFlux gets a reference to the given NullableString and assigns it to the TypeFlux field.
func (o *FluxResume) SetTypeFlux(v string) {
	o.TypeFlux.Set(&v)
}
// SetTypeFluxNil sets the value for TypeFlux to be an explicit nil
func (o *FluxResume) SetTypeFluxNil() {
	o.TypeFlux.Set(nil)
}

// UnsetTypeFlux ensures that no value is present for TypeFlux, not even an explicit nil
func (o *FluxResume) UnsetTypeFlux() {
	o.TypeFlux.Unset()
}

// GetDirectionFlux returns the DirectionFlux field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FluxResume) GetDirectionFlux() string {
	if o == nil || IsNil(o.DirectionFlux.Get()) {
		var ret string
		return ret
	}
	return *o.DirectionFlux.Get()
}

// GetDirectionFluxOk returns a tuple with the DirectionFlux field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FluxResume) GetDirectionFluxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectionFlux.Get(), o.DirectionFlux.IsSet()
}

// HasDirectionFlux returns a boolean if a field has been set.
func (o *FluxResume) HasDirectionFlux() bool {
	if o != nil && o.DirectionFlux.IsSet() {
		return true
	}

	return false
}

// SetDirectionFlux gets a reference to the given NullableString and assigns it to the DirectionFlux field.
func (o *FluxResume) SetDirectionFlux(v string) {
	o.DirectionFlux.Set(&v)
}
// SetDirectionFluxNil sets the value for DirectionFlux to be an explicit nil
func (o *FluxResume) SetDirectionFluxNil() {
	o.DirectionFlux.Set(nil)
}

// UnsetDirectionFlux ensures that no value is present for DirectionFlux, not even an explicit nil
func (o *FluxResume) UnsetDirectionFlux() {
	o.DirectionFlux.Unset()
}

// GetStatutAcquittement returns the StatutAcquittement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FluxResume) GetStatutAcquittement() string {
	if o == nil || IsNil(o.StatutAcquittement.Get()) {
		var ret string
		return ret
	}
	return *o.StatutAcquittement.Get()
}

// GetStatutAcquittementOk returns a tuple with the StatutAcquittement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FluxResume) GetStatutAcquittementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatutAcquittement.Get(), o.StatutAcquittement.IsSet()
}

// HasStatutAcquittement returns a boolean if a field has been set.
func (o *FluxResume) HasStatutAcquittement() bool {
	if o != nil && o.StatutAcquittement.IsSet() {
		return true
	}

	return false
}

// SetStatutAcquittement gets a reference to the given NullableString and assigns it to the StatutAcquittement field.
func (o *FluxResume) SetStatutAcquittement(v string) {
	o.StatutAcquittement.Set(&v)
}
// SetStatutAcquittementNil sets the value for StatutAcquittement to be an explicit nil
func (o *FluxResume) SetStatutAcquittementNil() {
	o.StatutAcquittement.Set(nil)
}

// UnsetStatutAcquittement ensures that no value is present for StatutAcquittement, not even an explicit nil
func (o *FluxResume) UnsetStatutAcquittement() {
	o.StatutAcquittement.Unset()
}

// GetDateCreation returns the DateCreation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FluxResume) GetDateCreation() string {
	if o == nil || IsNil(o.DateCreation.Get()) {
		var ret string
		return ret
	}
	return *o.DateCreation.Get()
}

// GetDateCreationOk returns a tuple with the DateCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FluxResume) GetDateCreationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCreation.Get(), o.DateCreation.IsSet()
}

// HasDateCreation returns a boolean if a field has been set.
func (o *FluxResume) HasDateCreation() bool {
	if o != nil && o.DateCreation.IsSet() {
		return true
	}

	return false
}

// SetDateCreation gets a reference to the given NullableString and assigns it to the DateCreation field.
func (o *FluxResume) SetDateCreation(v string) {
	o.DateCreation.Set(&v)
}
// SetDateCreationNil sets the value for DateCreation to be an explicit nil
func (o *FluxResume) SetDateCreationNil() {
	o.DateCreation.Set(nil)
}

// UnsetDateCreation ensures that no value is present for DateCreation, not even an explicit nil
func (o *FluxResume) UnsetDateCreation() {
	o.DateCreation.Unset()
}

// GetDateMaj returns the DateMaj field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FluxResume) GetDateMaj() string {
	if o == nil || IsNil(o.DateMaj.Get()) {
		var ret string
		return ret
	}
	return *o.DateMaj.Get()
}

// GetDateMajOk returns a tuple with the DateMaj field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FluxResume) GetDateMajOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateMaj.Get(), o.DateMaj.IsSet()
}

// HasDateMaj returns a boolean if a field has been set.
func (o *FluxResume) HasDateMaj() bool {
	if o != nil && o.DateMaj.IsSet() {
		return true
	}

	return false
}

// SetDateMaj gets a reference to the given NullableString and assigns it to the DateMaj field.
func (o *FluxResume) SetDateMaj(v string) {
	o.DateMaj.Set(&v)
}
// SetDateMajNil sets the value for DateMaj to be an explicit nil
func (o *FluxResume) SetDateMajNil() {
	o.DateMaj.Set(nil)
}

// UnsetDateMaj ensures that no value is present for DateMaj, not even an explicit nil
func (o *FluxResume) UnsetDateMaj() {
	o.DateMaj.Unset()
}

func (o FluxResume) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FluxResume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["flow_id"] = o.FlowId
	if o.TrackingId.IsSet() {
		toSerialize["tracking_id"] = o.TrackingId.Get()
	}
	toSerialize["nom"] = o.Nom
	if o.TypeFlux.IsSet() {
		toSerialize["type_flux"] = o.TypeFlux.Get()
	}
	if o.DirectionFlux.IsSet() {
		toSerialize["direction_flux"] = o.DirectionFlux.Get()
	}
	if o.StatutAcquittement.IsSet() {
		toSerialize["statut_acquittement"] = o.StatutAcquittement.Get()
	}
	if o.DateCreation.IsSet() {
		toSerialize["date_creation"] = o.DateCreation.Get()
	}
	if o.DateMaj.IsSet() {
		toSerialize["date_maj"] = o.DateMaj.Get()
	}
	return toSerialize, nil
}

func (o *FluxResume) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"flow_id",
		"nom",
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

	varFluxResume := _FluxResume{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFluxResume)

	if err != nil {
		return err
	}

	*o = FluxResume(varFluxResume)

	return err
}

type NullableFluxResume struct {
	value *FluxResume
	isSet bool
}

func (v NullableFluxResume) Get() *FluxResume {
	return v.value
}

func (v *NullableFluxResume) Set(val *FluxResume) {
	v.value = val
	v.isSet = true
}

func (v NullableFluxResume) IsSet() bool {
	return v.isSet
}

func (v *NullableFluxResume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFluxResume(val *FluxResume) *NullableFluxResume {
	return &NullableFluxResume{value: val, isSet: true}
}

func (v NullableFluxResume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFluxResume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


