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

// checks if the SoumettreFactureCompleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoumettreFactureCompleteRequest{}

// SoumettreFactureCompleteRequest Requête pour soumettre une facture complète (génération + soumission).  Workflow : 1. Auto-enrichissement (optionnel) 2. Génération PDF Factur-X 3. Signature (optionnelle) 4. Soumission vers la destination
type SoumettreFactureCompleteRequest struct {
	// Données de la facture au format simplifié (voir exemples)
	DonneesFacture DonneesFactureSimplifiees `json:"donnees_facture"`
	// PDF source encodé en base64 (sera transformé en Factur-X)
	PdfSource string `json:"pdf_source"`
	Destination Destination `json:"destination"`
	Signature NullableParametresSignature `json:"signature,omitempty"`
	// Options de traitement
	Options *OptionsProcessing `json:"options,omitempty"`
}

type _SoumettreFactureCompleteRequest SoumettreFactureCompleteRequest

// NewSoumettreFactureCompleteRequest instantiates a new SoumettreFactureCompleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoumettreFactureCompleteRequest(donneesFacture DonneesFactureSimplifiees, pdfSource string, destination Destination) *SoumettreFactureCompleteRequest {
	this := SoumettreFactureCompleteRequest{}
	this.DonneesFacture = donneesFacture
	this.PdfSource = pdfSource
	this.Destination = destination
	return &this
}

// NewSoumettreFactureCompleteRequestWithDefaults instantiates a new SoumettreFactureCompleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoumettreFactureCompleteRequestWithDefaults() *SoumettreFactureCompleteRequest {
	this := SoumettreFactureCompleteRequest{}
	return &this
}

// GetDonneesFacture returns the DonneesFacture field value
func (o *SoumettreFactureCompleteRequest) GetDonneesFacture() DonneesFactureSimplifiees {
	if o == nil {
		var ret DonneesFactureSimplifiees
		return ret
	}

	return o.DonneesFacture
}

// GetDonneesFactureOk returns a tuple with the DonneesFacture field value
// and a boolean to check if the value has been set.
func (o *SoumettreFactureCompleteRequest) GetDonneesFactureOk() (*DonneesFactureSimplifiees, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DonneesFacture, true
}

// SetDonneesFacture sets field value
func (o *SoumettreFactureCompleteRequest) SetDonneesFacture(v DonneesFactureSimplifiees) {
	o.DonneesFacture = v
}

// GetPdfSource returns the PdfSource field value
func (o *SoumettreFactureCompleteRequest) GetPdfSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PdfSource
}

// GetPdfSourceOk returns a tuple with the PdfSource field value
// and a boolean to check if the value has been set.
func (o *SoumettreFactureCompleteRequest) GetPdfSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PdfSource, true
}

// SetPdfSource sets field value
func (o *SoumettreFactureCompleteRequest) SetPdfSource(v string) {
	o.PdfSource = v
}

// GetDestination returns the Destination field value
func (o *SoumettreFactureCompleteRequest) GetDestination() Destination {
	if o == nil {
		var ret Destination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *SoumettreFactureCompleteRequest) GetDestinationOk() (*Destination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *SoumettreFactureCompleteRequest) SetDestination(v Destination) {
	o.Destination = v
}

// GetSignature returns the Signature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureCompleteRequest) GetSignature() ParametresSignature {
	if o == nil || IsNil(o.Signature.Get()) {
		var ret ParametresSignature
		return ret
	}
	return *o.Signature.Get()
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureCompleteRequest) GetSignatureOk() (*ParametresSignature, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signature.Get(), o.Signature.IsSet()
}

// HasSignature returns a boolean if a field has been set.
func (o *SoumettreFactureCompleteRequest) HasSignature() bool {
	if o != nil && o.Signature.IsSet() {
		return true
	}

	return false
}

// SetSignature gets a reference to the given NullableParametresSignature and assigns it to the Signature field.
func (o *SoumettreFactureCompleteRequest) SetSignature(v ParametresSignature) {
	o.Signature.Set(&v)
}
// SetSignatureNil sets the value for Signature to be an explicit nil
func (o *SoumettreFactureCompleteRequest) SetSignatureNil() {
	o.Signature.Set(nil)
}

// UnsetSignature ensures that no value is present for Signature, not even an explicit nil
func (o *SoumettreFactureCompleteRequest) UnsetSignature() {
	o.Signature.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SoumettreFactureCompleteRequest) GetOptions() OptionsProcessing {
	if o == nil || IsNil(o.Options) {
		var ret OptionsProcessing
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoumettreFactureCompleteRequest) GetOptionsOk() (*OptionsProcessing, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SoumettreFactureCompleteRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given OptionsProcessing and assigns it to the Options field.
func (o *SoumettreFactureCompleteRequest) SetOptions(v OptionsProcessing) {
	o.Options = &v
}

func (o SoumettreFactureCompleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoumettreFactureCompleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["donnees_facture"] = o.DonneesFacture
	toSerialize["pdf_source"] = o.PdfSource
	toSerialize["destination"] = o.Destination
	if o.Signature.IsSet() {
		toSerialize["signature"] = o.Signature.Get()
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *SoumettreFactureCompleteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"donnees_facture",
		"pdf_source",
		"destination",
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

	varSoumettreFactureCompleteRequest := _SoumettreFactureCompleteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSoumettreFactureCompleteRequest)

	if err != nil {
		return err
	}

	*o = SoumettreFactureCompleteRequest(varSoumettreFactureCompleteRequest)

	return err
}

type NullableSoumettreFactureCompleteRequest struct {
	value *SoumettreFactureCompleteRequest
	isSet bool
}

func (v NullableSoumettreFactureCompleteRequest) Get() *SoumettreFactureCompleteRequest {
	return v.value
}

func (v *NullableSoumettreFactureCompleteRequest) Set(val *SoumettreFactureCompleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSoumettreFactureCompleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSoumettreFactureCompleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoumettreFactureCompleteRequest(val *SoumettreFactureCompleteRequest) *NullableSoumettreFactureCompleteRequest {
	return &NullableSoumettreFactureCompleteRequest{value: val, isSet: true}
}

func (v NullableSoumettreFactureCompleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoumettreFactureCompleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


