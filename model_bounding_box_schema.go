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

// checks if the BoundingBoxSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoundingBoxSchema{}

// BoundingBoxSchema Coordonnées d'une zone rectangulaire dans le PDF.  Les coordonnées sont en points PDF (1 point = 1/72 pouce). L'origine (0,0) est en bas à gauche de la page.
type BoundingBoxSchema struct {
	// Coordonnée X gauche
	X0 float32 `json:"x0"`
	// Coordonnée Y bas
	Y0 float32 `json:"y0"`
	// Coordonnée X droite
	X1 float32 `json:"x1"`
	// Coordonnée Y haut
	Y1 float32 `json:"y1"`
	// Numéro de page (0-indexed)
	Page *int32 `json:"page,omitempty"`
	// Largeur de la zone
	Width float32 `json:"width"`
	// Hauteur de la zone
	Height float32 `json:"height"`
}

type _BoundingBoxSchema BoundingBoxSchema

// NewBoundingBoxSchema instantiates a new BoundingBoxSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoundingBoxSchema(x0 float32, y0 float32, x1 float32, y1 float32, width float32, height float32) *BoundingBoxSchema {
	this := BoundingBoxSchema{}
	this.X0 = x0
	this.Y0 = y0
	this.X1 = x1
	this.Y1 = y1
	var page int32 = 0
	this.Page = &page
	this.Width = width
	this.Height = height
	return &this
}

// NewBoundingBoxSchemaWithDefaults instantiates a new BoundingBoxSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoundingBoxSchemaWithDefaults() *BoundingBoxSchema {
	this := BoundingBoxSchema{}
	var page int32 = 0
	this.Page = &page
	return &this
}

// GetX0 returns the X0 field value
func (o *BoundingBoxSchema) GetX0() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.X0
}

// GetX0Ok returns a tuple with the X0 field value
// and a boolean to check if the value has been set.
func (o *BoundingBoxSchema) GetX0Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X0, true
}

// SetX0 sets field value
func (o *BoundingBoxSchema) SetX0(v float32) {
	o.X0 = v
}

// GetY0 returns the Y0 field value
func (o *BoundingBoxSchema) GetY0() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Y0
}

// GetY0Ok returns a tuple with the Y0 field value
// and a boolean to check if the value has been set.
func (o *BoundingBoxSchema) GetY0Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y0, true
}

// SetY0 sets field value
func (o *BoundingBoxSchema) SetY0(v float32) {
	o.Y0 = v
}

// GetX1 returns the X1 field value
func (o *BoundingBoxSchema) GetX1() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.X1
}

// GetX1Ok returns a tuple with the X1 field value
// and a boolean to check if the value has been set.
func (o *BoundingBoxSchema) GetX1Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X1, true
}

// SetX1 sets field value
func (o *BoundingBoxSchema) SetX1(v float32) {
	o.X1 = v
}

// GetY1 returns the Y1 field value
func (o *BoundingBoxSchema) GetY1() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Y1
}

// GetY1Ok returns a tuple with the Y1 field value
// and a boolean to check if the value has been set.
func (o *BoundingBoxSchema) GetY1Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y1, true
}

// SetY1 sets field value
func (o *BoundingBoxSchema) SetY1(v float32) {
	o.Y1 = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *BoundingBoxSchema) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundingBoxSchema) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *BoundingBoxSchema) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *BoundingBoxSchema) SetPage(v int32) {
	o.Page = &v
}

// GetWidth returns the Width field value
func (o *BoundingBoxSchema) GetWidth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *BoundingBoxSchema) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *BoundingBoxSchema) SetWidth(v float32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *BoundingBoxSchema) GetHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *BoundingBoxSchema) GetHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *BoundingBoxSchema) SetHeight(v float32) {
	o.Height = v
}

func (o BoundingBoxSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoundingBoxSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x0"] = o.X0
	toSerialize["y0"] = o.Y0
	toSerialize["x1"] = o.X1
	toSerialize["y1"] = o.Y1
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	return toSerialize, nil
}

func (o *BoundingBoxSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"x0",
		"y0",
		"x1",
		"y1",
		"width",
		"height",
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

	varBoundingBoxSchema := _BoundingBoxSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBoundingBoxSchema)

	if err != nil {
		return err
	}

	*o = BoundingBoxSchema(varBoundingBoxSchema)

	return err
}

type NullableBoundingBoxSchema struct {
	value *BoundingBoxSchema
	isSet bool
}

func (v NullableBoundingBoxSchema) Get() *BoundingBoxSchema {
	return v.value
}

func (v *NullableBoundingBoxSchema) Set(val *BoundingBoxSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableBoundingBoxSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableBoundingBoxSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoundingBoxSchema(val *BoundingBoxSchema) *NullableBoundingBoxSchema {
	return &NullableBoundingBoxSchema{value: val, isSet: true}
}

func (v NullableBoundingBoxSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoundingBoxSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


