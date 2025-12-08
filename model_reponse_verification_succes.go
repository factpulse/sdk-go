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

// checks if the ReponseVerificationSucces type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReponseVerificationSucces{}

// ReponseVerificationSucces Réponse de vérification réussie avec données unifiées.
type ReponseVerificationSucces struct {
	// True si aucun écart critique
	EstConforme bool `json:"est_conforme"`
	// Score de conformité (0-100%)
	ScoreConformite float32 `json:"score_conformite"`
	// Nombre de champs vérifiés
	ChampsVerifies *int32 `json:"champs_verifies,omitempty"`
	// Nombre de champs conformes
	ChampsConformes *int32 `json:"champs_conformes,omitempty"`
	// True si le PDF contient du XML Factur-X
	EstFacturx *bool `json:"est_facturx,omitempty"`
	ProfilFacturx NullableString `json:"profil_facturx,omitempty"`
	// Liste des champs vérifiés avec valeurs, statuts et coordonnées PDF
	Champs []ChampVerifieSchema `json:"champs,omitempty"`
	// Notes obligatoires (PMT, PMD, AAB) avec localisation PDF
	NotesObligatoires []NoteObligatoireSchema `json:"notes_obligatoires,omitempty"`
	// Dimensions de chaque page du PDF (largeur, hauteur)
	DimensionsPages []DimensionPageSchema `json:"dimensions_pages,omitempty"`
	// Avertissements non bloquants
	Avertissements []string `json:"avertissements,omitempty"`
}

type _ReponseVerificationSucces ReponseVerificationSucces

// NewReponseVerificationSucces instantiates a new ReponseVerificationSucces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReponseVerificationSucces(estConforme bool, scoreConformite float32) *ReponseVerificationSucces {
	this := ReponseVerificationSucces{}
	this.EstConforme = estConforme
	this.ScoreConformite = scoreConformite
	var champsVerifies int32 = 0
	this.ChampsVerifies = &champsVerifies
	var champsConformes int32 = 0
	this.ChampsConformes = &champsConformes
	var estFacturx bool = false
	this.EstFacturx = &estFacturx
	return &this
}

// NewReponseVerificationSuccesWithDefaults instantiates a new ReponseVerificationSucces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReponseVerificationSuccesWithDefaults() *ReponseVerificationSucces {
	this := ReponseVerificationSucces{}
	var champsVerifies int32 = 0
	this.ChampsVerifies = &champsVerifies
	var champsConformes int32 = 0
	this.ChampsConformes = &champsConformes
	var estFacturx bool = false
	this.EstFacturx = &estFacturx
	return &this
}

// GetEstConforme returns the EstConforme field value
func (o *ReponseVerificationSucces) GetEstConforme() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EstConforme
}

// GetEstConformeOk returns a tuple with the EstConforme field value
// and a boolean to check if the value has been set.
func (o *ReponseVerificationSucces) GetEstConformeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstConforme, true
}

// SetEstConforme sets field value
func (o *ReponseVerificationSucces) SetEstConforme(v bool) {
	o.EstConforme = v
}

// GetScoreConformite returns the ScoreConformite field value
func (o *ReponseVerificationSucces) GetScoreConformite() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ScoreConformite
}

// GetScoreConformiteOk returns a tuple with the ScoreConformite field value
// and a boolean to check if the value has been set.
func (o *ReponseVerificationSucces) GetScoreConformiteOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScoreConformite, true
}

// SetScoreConformite sets field value
func (o *ReponseVerificationSucces) SetScoreConformite(v float32) {
	o.ScoreConformite = v
}

// GetChampsVerifies returns the ChampsVerifies field value if set, zero value otherwise.
func (o *ReponseVerificationSucces) GetChampsVerifies() int32 {
	if o == nil || IsNil(o.ChampsVerifies) {
		var ret int32
		return ret
	}
	return *o.ChampsVerifies
}

// GetChampsVerifiesOk returns a tuple with the ChampsVerifies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReponseVerificationSucces) GetChampsVerifiesOk() (*int32, bool) {
	if o == nil || IsNil(o.ChampsVerifies) {
		return nil, false
	}
	return o.ChampsVerifies, true
}

// HasChampsVerifies returns a boolean if a field has been set.
func (o *ReponseVerificationSucces) HasChampsVerifies() bool {
	if o != nil && !IsNil(o.ChampsVerifies) {
		return true
	}

	return false
}

// SetChampsVerifies gets a reference to the given int32 and assigns it to the ChampsVerifies field.
func (o *ReponseVerificationSucces) SetChampsVerifies(v int32) {
	o.ChampsVerifies = &v
}

// GetChampsConformes returns the ChampsConformes field value if set, zero value otherwise.
func (o *ReponseVerificationSucces) GetChampsConformes() int32 {
	if o == nil || IsNil(o.ChampsConformes) {
		var ret int32
		return ret
	}
	return *o.ChampsConformes
}

// GetChampsConformesOk returns a tuple with the ChampsConformes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReponseVerificationSucces) GetChampsConformesOk() (*int32, bool) {
	if o == nil || IsNil(o.ChampsConformes) {
		return nil, false
	}
	return o.ChampsConformes, true
}

// HasChampsConformes returns a boolean if a field has been set.
func (o *ReponseVerificationSucces) HasChampsConformes() bool {
	if o != nil && !IsNil(o.ChampsConformes) {
		return true
	}

	return false
}

// SetChampsConformes gets a reference to the given int32 and assigns it to the ChampsConformes field.
func (o *ReponseVerificationSucces) SetChampsConformes(v int32) {
	o.ChampsConformes = &v
}

// GetEstFacturx returns the EstFacturx field value if set, zero value otherwise.
func (o *ReponseVerificationSucces) GetEstFacturx() bool {
	if o == nil || IsNil(o.EstFacturx) {
		var ret bool
		return ret
	}
	return *o.EstFacturx
}

// GetEstFacturxOk returns a tuple with the EstFacturx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReponseVerificationSucces) GetEstFacturxOk() (*bool, bool) {
	if o == nil || IsNil(o.EstFacturx) {
		return nil, false
	}
	return o.EstFacturx, true
}

// HasEstFacturx returns a boolean if a field has been set.
func (o *ReponseVerificationSucces) HasEstFacturx() bool {
	if o != nil && !IsNil(o.EstFacturx) {
		return true
	}

	return false
}

// SetEstFacturx gets a reference to the given bool and assigns it to the EstFacturx field.
func (o *ReponseVerificationSucces) SetEstFacturx(v bool) {
	o.EstFacturx = &v
}

// GetProfilFacturx returns the ProfilFacturx field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReponseVerificationSucces) GetProfilFacturx() string {
	if o == nil || IsNil(o.ProfilFacturx.Get()) {
		var ret string
		return ret
	}
	return *o.ProfilFacturx.Get()
}

// GetProfilFacturxOk returns a tuple with the ProfilFacturx field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReponseVerificationSucces) GetProfilFacturxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilFacturx.Get(), o.ProfilFacturx.IsSet()
}

// HasProfilFacturx returns a boolean if a field has been set.
func (o *ReponseVerificationSucces) HasProfilFacturx() bool {
	if o != nil && o.ProfilFacturx.IsSet() {
		return true
	}

	return false
}

// SetProfilFacturx gets a reference to the given NullableString and assigns it to the ProfilFacturx field.
func (o *ReponseVerificationSucces) SetProfilFacturx(v string) {
	o.ProfilFacturx.Set(&v)
}
// SetProfilFacturxNil sets the value for ProfilFacturx to be an explicit nil
func (o *ReponseVerificationSucces) SetProfilFacturxNil() {
	o.ProfilFacturx.Set(nil)
}

// UnsetProfilFacturx ensures that no value is present for ProfilFacturx, not even an explicit nil
func (o *ReponseVerificationSucces) UnsetProfilFacturx() {
	o.ProfilFacturx.Unset()
}

// GetChamps returns the Champs field value if set, zero value otherwise.
func (o *ReponseVerificationSucces) GetChamps() []ChampVerifieSchema {
	if o == nil || IsNil(o.Champs) {
		var ret []ChampVerifieSchema
		return ret
	}
	return o.Champs
}

// GetChampsOk returns a tuple with the Champs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReponseVerificationSucces) GetChampsOk() ([]ChampVerifieSchema, bool) {
	if o == nil || IsNil(o.Champs) {
		return nil, false
	}
	return o.Champs, true
}

// HasChamps returns a boolean if a field has been set.
func (o *ReponseVerificationSucces) HasChamps() bool {
	if o != nil && !IsNil(o.Champs) {
		return true
	}

	return false
}

// SetChamps gets a reference to the given []ChampVerifieSchema and assigns it to the Champs field.
func (o *ReponseVerificationSucces) SetChamps(v []ChampVerifieSchema) {
	o.Champs = v
}

// GetNotesObligatoires returns the NotesObligatoires field value if set, zero value otherwise.
func (o *ReponseVerificationSucces) GetNotesObligatoires() []NoteObligatoireSchema {
	if o == nil || IsNil(o.NotesObligatoires) {
		var ret []NoteObligatoireSchema
		return ret
	}
	return o.NotesObligatoires
}

// GetNotesObligatoiresOk returns a tuple with the NotesObligatoires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReponseVerificationSucces) GetNotesObligatoiresOk() ([]NoteObligatoireSchema, bool) {
	if o == nil || IsNil(o.NotesObligatoires) {
		return nil, false
	}
	return o.NotesObligatoires, true
}

// HasNotesObligatoires returns a boolean if a field has been set.
func (o *ReponseVerificationSucces) HasNotesObligatoires() bool {
	if o != nil && !IsNil(o.NotesObligatoires) {
		return true
	}

	return false
}

// SetNotesObligatoires gets a reference to the given []NoteObligatoireSchema and assigns it to the NotesObligatoires field.
func (o *ReponseVerificationSucces) SetNotesObligatoires(v []NoteObligatoireSchema) {
	o.NotesObligatoires = v
}

// GetDimensionsPages returns the DimensionsPages field value if set, zero value otherwise.
func (o *ReponseVerificationSucces) GetDimensionsPages() []DimensionPageSchema {
	if o == nil || IsNil(o.DimensionsPages) {
		var ret []DimensionPageSchema
		return ret
	}
	return o.DimensionsPages
}

// GetDimensionsPagesOk returns a tuple with the DimensionsPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReponseVerificationSucces) GetDimensionsPagesOk() ([]DimensionPageSchema, bool) {
	if o == nil || IsNil(o.DimensionsPages) {
		return nil, false
	}
	return o.DimensionsPages, true
}

// HasDimensionsPages returns a boolean if a field has been set.
func (o *ReponseVerificationSucces) HasDimensionsPages() bool {
	if o != nil && !IsNil(o.DimensionsPages) {
		return true
	}

	return false
}

// SetDimensionsPages gets a reference to the given []DimensionPageSchema and assigns it to the DimensionsPages field.
func (o *ReponseVerificationSucces) SetDimensionsPages(v []DimensionPageSchema) {
	o.DimensionsPages = v
}

// GetAvertissements returns the Avertissements field value if set, zero value otherwise.
func (o *ReponseVerificationSucces) GetAvertissements() []string {
	if o == nil || IsNil(o.Avertissements) {
		var ret []string
		return ret
	}
	return o.Avertissements
}

// GetAvertissementsOk returns a tuple with the Avertissements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReponseVerificationSucces) GetAvertissementsOk() ([]string, bool) {
	if o == nil || IsNil(o.Avertissements) {
		return nil, false
	}
	return o.Avertissements, true
}

// HasAvertissements returns a boolean if a field has been set.
func (o *ReponseVerificationSucces) HasAvertissements() bool {
	if o != nil && !IsNil(o.Avertissements) {
		return true
	}

	return false
}

// SetAvertissements gets a reference to the given []string and assigns it to the Avertissements field.
func (o *ReponseVerificationSucces) SetAvertissements(v []string) {
	o.Avertissements = v
}

func (o ReponseVerificationSucces) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReponseVerificationSucces) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["est_conforme"] = o.EstConforme
	toSerialize["score_conformite"] = o.ScoreConformite
	if !IsNil(o.ChampsVerifies) {
		toSerialize["champs_verifies"] = o.ChampsVerifies
	}
	if !IsNil(o.ChampsConformes) {
		toSerialize["champs_conformes"] = o.ChampsConformes
	}
	if !IsNil(o.EstFacturx) {
		toSerialize["est_facturx"] = o.EstFacturx
	}
	if o.ProfilFacturx.IsSet() {
		toSerialize["profil_facturx"] = o.ProfilFacturx.Get()
	}
	if !IsNil(o.Champs) {
		toSerialize["champs"] = o.Champs
	}
	if !IsNil(o.NotesObligatoires) {
		toSerialize["notes_obligatoires"] = o.NotesObligatoires
	}
	if !IsNil(o.DimensionsPages) {
		toSerialize["dimensions_pages"] = o.DimensionsPages
	}
	if !IsNil(o.Avertissements) {
		toSerialize["avertissements"] = o.Avertissements
	}
	return toSerialize, nil
}

func (o *ReponseVerificationSucces) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"est_conforme",
		"score_conformite",
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

	varReponseVerificationSucces := _ReponseVerificationSucces{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReponseVerificationSucces)

	if err != nil {
		return err
	}

	*o = ReponseVerificationSucces(varReponseVerificationSucces)

	return err
}

type NullableReponseVerificationSucces struct {
	value *ReponseVerificationSucces
	isSet bool
}

func (v NullableReponseVerificationSucces) Get() *ReponseVerificationSucces {
	return v.value
}

func (v *NullableReponseVerificationSucces) Set(val *ReponseVerificationSucces) {
	v.value = val
	v.isSet = true
}

func (v NullableReponseVerificationSucces) IsSet() bool {
	return v.isSet
}

func (v *NullableReponseVerificationSucces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReponseVerificationSucces(val *ReponseVerificationSucces) *NullableReponseVerificationSucces {
	return &NullableReponseVerificationSucces{value: val, isSet: true}
}

func (v NullableReponseVerificationSucces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReponseVerificationSucces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


