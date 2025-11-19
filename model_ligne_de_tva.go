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

// checks if the LigneDeTVA type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LigneDeTVA{}

// LigneDeTVA Représente une ligne de totalisation par taux de TVA.
type LigneDeTVA struct {
	MontantBaseHt MontantBaseHt `json:"montantBaseHt"`
	MontantTva MontantTva `json:"montantTva"`
	Taux NullableString `json:"taux,omitempty"`
	TauxManuel *Tauxmanuel `json:"tauxManuel,omitempty"`
	Categorie NullableCategorieTVA `json:"categorie,omitempty"`
}

type _LigneDeTVA LigneDeTVA

// NewLigneDeTVA instantiates a new LigneDeTVA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLigneDeTVA(montantBaseHt MontantBaseHt, montantTva MontantTva) *LigneDeTVA {
	this := LigneDeTVA{}
	this.MontantBaseHt = montantBaseHt
	this.MontantTva = montantTva
	return &this
}

// NewLigneDeTVAWithDefaults instantiates a new LigneDeTVA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLigneDeTVAWithDefaults() *LigneDeTVA {
	this := LigneDeTVA{}
	return &this
}

// GetMontantBaseHt returns the MontantBaseHt field value
func (o *LigneDeTVA) GetMontantBaseHt() MontantBaseHt {
	if o == nil {
		var ret MontantBaseHt
		return ret
	}

	return o.MontantBaseHt
}

// GetMontantBaseHtOk returns a tuple with the MontantBaseHt field value
// and a boolean to check if the value has been set.
func (o *LigneDeTVA) GetMontantBaseHtOk() (*MontantBaseHt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantBaseHt, true
}

// SetMontantBaseHt sets field value
func (o *LigneDeTVA) SetMontantBaseHt(v MontantBaseHt) {
	o.MontantBaseHt = v
}

// GetMontantTva returns the MontantTva field value
func (o *LigneDeTVA) GetMontantTva() MontantTva {
	if o == nil {
		var ret MontantTva
		return ret
	}

	return o.MontantTva
}

// GetMontantTvaOk returns a tuple with the MontantTva field value
// and a boolean to check if the value has been set.
func (o *LigneDeTVA) GetMontantTvaOk() (*MontantTva, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantTva, true
}

// SetMontantTva sets field value
func (o *LigneDeTVA) SetMontantTva(v MontantTva) {
	o.MontantTva = v
}

// GetTaux returns the Taux field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LigneDeTVA) GetTaux() string {
	if o == nil || IsNil(o.Taux.Get()) {
		var ret string
		return ret
	}
	return *o.Taux.Get()
}

// GetTauxOk returns a tuple with the Taux field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LigneDeTVA) GetTauxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Taux.Get(), o.Taux.IsSet()
}

// HasTaux returns a boolean if a field has been set.
func (o *LigneDeTVA) HasTaux() bool {
	if o != nil && o.Taux.IsSet() {
		return true
	}

	return false
}

// SetTaux gets a reference to the given NullableString and assigns it to the Taux field.
func (o *LigneDeTVA) SetTaux(v string) {
	o.Taux.Set(&v)
}
// SetTauxNil sets the value for Taux to be an explicit nil
func (o *LigneDeTVA) SetTauxNil() {
	o.Taux.Set(nil)
}

// UnsetTaux ensures that no value is present for Taux, not even an explicit nil
func (o *LigneDeTVA) UnsetTaux() {
	o.Taux.Unset()
}

// GetTauxManuel returns the TauxManuel field value if set, zero value otherwise.
func (o *LigneDeTVA) GetTauxManuel() Tauxmanuel {
	if o == nil || IsNil(o.TauxManuel) {
		var ret Tauxmanuel
		return ret
	}
	return *o.TauxManuel
}

// GetTauxManuelOk returns a tuple with the TauxManuel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LigneDeTVA) GetTauxManuelOk() (*Tauxmanuel, bool) {
	if o == nil || IsNil(o.TauxManuel) {
		return nil, false
	}
	return o.TauxManuel, true
}

// HasTauxManuel returns a boolean if a field has been set.
func (o *LigneDeTVA) HasTauxManuel() bool {
	if o != nil && !IsNil(o.TauxManuel) {
		return true
	}

	return false
}

// SetTauxManuel gets a reference to the given Tauxmanuel and assigns it to the TauxManuel field.
func (o *LigneDeTVA) SetTauxManuel(v Tauxmanuel) {
	o.TauxManuel = &v
}

// GetCategorie returns the Categorie field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LigneDeTVA) GetCategorie() CategorieTVA {
	if o == nil || IsNil(o.Categorie.Get()) {
		var ret CategorieTVA
		return ret
	}
	return *o.Categorie.Get()
}

// GetCategorieOk returns a tuple with the Categorie field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LigneDeTVA) GetCategorieOk() (*CategorieTVA, bool) {
	if o == nil {
		return nil, false
	}
	return o.Categorie.Get(), o.Categorie.IsSet()
}

// HasCategorie returns a boolean if a field has been set.
func (o *LigneDeTVA) HasCategorie() bool {
	if o != nil && o.Categorie.IsSet() {
		return true
	}

	return false
}

// SetCategorie gets a reference to the given NullableCategorieTVA and assigns it to the Categorie field.
func (o *LigneDeTVA) SetCategorie(v CategorieTVA) {
	o.Categorie.Set(&v)
}
// SetCategorieNil sets the value for Categorie to be an explicit nil
func (o *LigneDeTVA) SetCategorieNil() {
	o.Categorie.Set(nil)
}

// UnsetCategorie ensures that no value is present for Categorie, not even an explicit nil
func (o *LigneDeTVA) UnsetCategorie() {
	o.Categorie.Unset()
}

func (o LigneDeTVA) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LigneDeTVA) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["montantBaseHt"] = o.MontantBaseHt
	toSerialize["montantTva"] = o.MontantTva
	if o.Taux.IsSet() {
		toSerialize["taux"] = o.Taux.Get()
	}
	if !IsNil(o.TauxManuel) {
		toSerialize["tauxManuel"] = o.TauxManuel
	}
	if o.Categorie.IsSet() {
		toSerialize["categorie"] = o.Categorie.Get()
	}
	return toSerialize, nil
}

func (o *LigneDeTVA) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"montantBaseHt",
		"montantTva",
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

	varLigneDeTVA := _LigneDeTVA{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLigneDeTVA)

	if err != nil {
		return err
	}

	*o = LigneDeTVA(varLigneDeTVA)

	return err
}

type NullableLigneDeTVA struct {
	value *LigneDeTVA
	isSet bool
}

func (v NullableLigneDeTVA) Get() *LigneDeTVA {
	return v.value
}

func (v *NullableLigneDeTVA) Set(val *LigneDeTVA) {
	v.value = val
	v.isSet = true
}

func (v NullableLigneDeTVA) IsSet() bool {
	return v.isSet
}

func (v *NullableLigneDeTVA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLigneDeTVA(val *LigneDeTVA) *NullableLigneDeTVA {
	return &NullableLigneDeTVA{value: val, isSet: true}
}

func (v NullableLigneDeTVA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLigneDeTVA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


