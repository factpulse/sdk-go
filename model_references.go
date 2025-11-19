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

// checks if the References type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &References{}

// References Contient les références diverses de la facture (devise, type, etc.).
type References struct {
	DeviseFacture *string `json:"deviseFacture,omitempty"`
	ModePaiement ModePaiement `json:"modePaiement"`
	TypeFacture TypeFacture `json:"typeFacture"`
	TypeTva TypeTVA `json:"typeTva"`
	NumeroMarche NullableString `json:"numeroMarche,omitempty"`
	MotifExonerationTva NullableString `json:"motifExonerationTva,omitempty"`
	NumeroBonCommande NullableString `json:"numeroBonCommande,omitempty"`
	NumeroFactureOrigine NullableString `json:"numeroFactureOrigine,omitempty"`
}

type _References References

// NewReferences instantiates a new References object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferences(modePaiement ModePaiement, typeFacture TypeFacture, typeTva TypeTVA) *References {
	this := References{}
	var deviseFacture string = "EUR"
	this.DeviseFacture = &deviseFacture
	this.ModePaiement = modePaiement
	this.TypeFacture = typeFacture
	this.TypeTva = typeTva
	return &this
}

// NewReferencesWithDefaults instantiates a new References object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencesWithDefaults() *References {
	this := References{}
	var deviseFacture string = "EUR"
	this.DeviseFacture = &deviseFacture
	return &this
}

// GetDeviseFacture returns the DeviseFacture field value if set, zero value otherwise.
func (o *References) GetDeviseFacture() string {
	if o == nil || IsNil(o.DeviseFacture) {
		var ret string
		return ret
	}
	return *o.DeviseFacture
}

// GetDeviseFactureOk returns a tuple with the DeviseFacture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *References) GetDeviseFactureOk() (*string, bool) {
	if o == nil || IsNil(o.DeviseFacture) {
		return nil, false
	}
	return o.DeviseFacture, true
}

// HasDeviseFacture returns a boolean if a field has been set.
func (o *References) HasDeviseFacture() bool {
	if o != nil && !IsNil(o.DeviseFacture) {
		return true
	}

	return false
}

// SetDeviseFacture gets a reference to the given string and assigns it to the DeviseFacture field.
func (o *References) SetDeviseFacture(v string) {
	o.DeviseFacture = &v
}

// GetModePaiement returns the ModePaiement field value
func (o *References) GetModePaiement() ModePaiement {
	if o == nil {
		var ret ModePaiement
		return ret
	}

	return o.ModePaiement
}

// GetModePaiementOk returns a tuple with the ModePaiement field value
// and a boolean to check if the value has been set.
func (o *References) GetModePaiementOk() (*ModePaiement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModePaiement, true
}

// SetModePaiement sets field value
func (o *References) SetModePaiement(v ModePaiement) {
	o.ModePaiement = v
}

// GetTypeFacture returns the TypeFacture field value
func (o *References) GetTypeFacture() TypeFacture {
	if o == nil {
		var ret TypeFacture
		return ret
	}

	return o.TypeFacture
}

// GetTypeFactureOk returns a tuple with the TypeFacture field value
// and a boolean to check if the value has been set.
func (o *References) GetTypeFactureOk() (*TypeFacture, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeFacture, true
}

// SetTypeFacture sets field value
func (o *References) SetTypeFacture(v TypeFacture) {
	o.TypeFacture = v
}

// GetTypeTva returns the TypeTva field value
func (o *References) GetTypeTva() TypeTVA {
	if o == nil {
		var ret TypeTVA
		return ret
	}

	return o.TypeTva
}

// GetTypeTvaOk returns a tuple with the TypeTva field value
// and a boolean to check if the value has been set.
func (o *References) GetTypeTvaOk() (*TypeTVA, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeTva, true
}

// SetTypeTva sets field value
func (o *References) SetTypeTva(v TypeTVA) {
	o.TypeTva = v
}

// GetNumeroMarche returns the NumeroMarche field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *References) GetNumeroMarche() string {
	if o == nil || IsNil(o.NumeroMarche.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroMarche.Get()
}

// GetNumeroMarcheOk returns a tuple with the NumeroMarche field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *References) GetNumeroMarcheOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroMarche.Get(), o.NumeroMarche.IsSet()
}

// HasNumeroMarche returns a boolean if a field has been set.
func (o *References) HasNumeroMarche() bool {
	if o != nil && o.NumeroMarche.IsSet() {
		return true
	}

	return false
}

// SetNumeroMarche gets a reference to the given NullableString and assigns it to the NumeroMarche field.
func (o *References) SetNumeroMarche(v string) {
	o.NumeroMarche.Set(&v)
}
// SetNumeroMarcheNil sets the value for NumeroMarche to be an explicit nil
func (o *References) SetNumeroMarcheNil() {
	o.NumeroMarche.Set(nil)
}

// UnsetNumeroMarche ensures that no value is present for NumeroMarche, not even an explicit nil
func (o *References) UnsetNumeroMarche() {
	o.NumeroMarche.Unset()
}

// GetMotifExonerationTva returns the MotifExonerationTva field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *References) GetMotifExonerationTva() string {
	if o == nil || IsNil(o.MotifExonerationTva.Get()) {
		var ret string
		return ret
	}
	return *o.MotifExonerationTva.Get()
}

// GetMotifExonerationTvaOk returns a tuple with the MotifExonerationTva field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *References) GetMotifExonerationTvaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MotifExonerationTva.Get(), o.MotifExonerationTva.IsSet()
}

// HasMotifExonerationTva returns a boolean if a field has been set.
func (o *References) HasMotifExonerationTva() bool {
	if o != nil && o.MotifExonerationTva.IsSet() {
		return true
	}

	return false
}

// SetMotifExonerationTva gets a reference to the given NullableString and assigns it to the MotifExonerationTva field.
func (o *References) SetMotifExonerationTva(v string) {
	o.MotifExonerationTva.Set(&v)
}
// SetMotifExonerationTvaNil sets the value for MotifExonerationTva to be an explicit nil
func (o *References) SetMotifExonerationTvaNil() {
	o.MotifExonerationTva.Set(nil)
}

// UnsetMotifExonerationTva ensures that no value is present for MotifExonerationTva, not even an explicit nil
func (o *References) UnsetMotifExonerationTva() {
	o.MotifExonerationTva.Unset()
}

// GetNumeroBonCommande returns the NumeroBonCommande field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *References) GetNumeroBonCommande() string {
	if o == nil || IsNil(o.NumeroBonCommande.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroBonCommande.Get()
}

// GetNumeroBonCommandeOk returns a tuple with the NumeroBonCommande field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *References) GetNumeroBonCommandeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroBonCommande.Get(), o.NumeroBonCommande.IsSet()
}

// HasNumeroBonCommande returns a boolean if a field has been set.
func (o *References) HasNumeroBonCommande() bool {
	if o != nil && o.NumeroBonCommande.IsSet() {
		return true
	}

	return false
}

// SetNumeroBonCommande gets a reference to the given NullableString and assigns it to the NumeroBonCommande field.
func (o *References) SetNumeroBonCommande(v string) {
	o.NumeroBonCommande.Set(&v)
}
// SetNumeroBonCommandeNil sets the value for NumeroBonCommande to be an explicit nil
func (o *References) SetNumeroBonCommandeNil() {
	o.NumeroBonCommande.Set(nil)
}

// UnsetNumeroBonCommande ensures that no value is present for NumeroBonCommande, not even an explicit nil
func (o *References) UnsetNumeroBonCommande() {
	o.NumeroBonCommande.Unset()
}

// GetNumeroFactureOrigine returns the NumeroFactureOrigine field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *References) GetNumeroFactureOrigine() string {
	if o == nil || IsNil(o.NumeroFactureOrigine.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroFactureOrigine.Get()
}

// GetNumeroFactureOrigineOk returns a tuple with the NumeroFactureOrigine field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *References) GetNumeroFactureOrigineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroFactureOrigine.Get(), o.NumeroFactureOrigine.IsSet()
}

// HasNumeroFactureOrigine returns a boolean if a field has been set.
func (o *References) HasNumeroFactureOrigine() bool {
	if o != nil && o.NumeroFactureOrigine.IsSet() {
		return true
	}

	return false
}

// SetNumeroFactureOrigine gets a reference to the given NullableString and assigns it to the NumeroFactureOrigine field.
func (o *References) SetNumeroFactureOrigine(v string) {
	o.NumeroFactureOrigine.Set(&v)
}
// SetNumeroFactureOrigineNil sets the value for NumeroFactureOrigine to be an explicit nil
func (o *References) SetNumeroFactureOrigineNil() {
	o.NumeroFactureOrigine.Set(nil)
}

// UnsetNumeroFactureOrigine ensures that no value is present for NumeroFactureOrigine, not even an explicit nil
func (o *References) UnsetNumeroFactureOrigine() {
	o.NumeroFactureOrigine.Unset()
}

func (o References) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o References) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviseFacture) {
		toSerialize["deviseFacture"] = o.DeviseFacture
	}
	toSerialize["modePaiement"] = o.ModePaiement
	toSerialize["typeFacture"] = o.TypeFacture
	toSerialize["typeTva"] = o.TypeTva
	if o.NumeroMarche.IsSet() {
		toSerialize["numeroMarche"] = o.NumeroMarche.Get()
	}
	if o.MotifExonerationTva.IsSet() {
		toSerialize["motifExonerationTva"] = o.MotifExonerationTva.Get()
	}
	if o.NumeroBonCommande.IsSet() {
		toSerialize["numeroBonCommande"] = o.NumeroBonCommande.Get()
	}
	if o.NumeroFactureOrigine.IsSet() {
		toSerialize["numeroFactureOrigine"] = o.NumeroFactureOrigine.Get()
	}
	return toSerialize, nil
}

func (o *References) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"modePaiement",
		"typeFacture",
		"typeTva",
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

	varReferences := _References{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReferences)

	if err != nil {
		return err
	}

	*o = References(varReferences)

	return err
}

type NullableReferences struct {
	value *References
	isSet bool
}

func (v NullableReferences) Get() *References {
	return v.value
}

func (v *NullableReferences) Set(val *References) {
	v.value = val
	v.isSet = true
}

func (v NullableReferences) IsSet() bool {
	return v.isSet
}

func (v *NullableReferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferences(val *References) *NullableReferences {
	return &NullableReferences{value: val, isSet: true}
}

func (v NullableReferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


