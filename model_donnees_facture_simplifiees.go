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

// checks if the DonneesFactureSimplifiees type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DonneesFactureSimplifiees{}

// DonneesFactureSimplifiees Données simplifiées de la facture (format minimal pour auto-enrichissement).
type DonneesFactureSimplifiees struct {
	// Numéro de facture unique
	Numero string `json:"numero"`
	// Informations sur l'émetteur (siret, iban, ...)
	Emetteur map[string]interface{} `json:"emetteur"`
	// Informations sur le destinataire (siret, ...)
	Destinataire map[string]interface{} `json:"destinataire"`
	// Lignes de la facture
	Lignes []map[string]interface{} `json:"lignes"`
	Date NullableString `json:"date,omitempty"`
	// Échéance en jours (défaut: 30)
	EcheanceJours *int32 `json:"echeance_jours,omitempty"`
	Commentaire NullableString `json:"commentaire,omitempty"`
	NumeroBonCommande NullableString `json:"numero_bon_commande,omitempty"`
	NumeroMarche NullableString `json:"numero_marche,omitempty"`
}

type _DonneesFactureSimplifiees DonneesFactureSimplifiees

// NewDonneesFactureSimplifiees instantiates a new DonneesFactureSimplifiees object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDonneesFactureSimplifiees(numero string, emetteur map[string]interface{}, destinataire map[string]interface{}, lignes []map[string]interface{}) *DonneesFactureSimplifiees {
	this := DonneesFactureSimplifiees{}
	this.Numero = numero
	this.Emetteur = emetteur
	this.Destinataire = destinataire
	this.Lignes = lignes
	var echeanceJours int32 = 30
	this.EcheanceJours = &echeanceJours
	return &this
}

// NewDonneesFactureSimplifieesWithDefaults instantiates a new DonneesFactureSimplifiees object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDonneesFactureSimplifieesWithDefaults() *DonneesFactureSimplifiees {
	this := DonneesFactureSimplifiees{}
	var echeanceJours int32 = 30
	this.EcheanceJours = &echeanceJours
	return &this
}

// GetNumero returns the Numero field value
func (o *DonneesFactureSimplifiees) GetNumero() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Numero
}

// GetNumeroOk returns a tuple with the Numero field value
// and a boolean to check if the value has been set.
func (o *DonneesFactureSimplifiees) GetNumeroOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Numero, true
}

// SetNumero sets field value
func (o *DonneesFactureSimplifiees) SetNumero(v string) {
	o.Numero = v
}

// GetEmetteur returns the Emetteur field value
func (o *DonneesFactureSimplifiees) GetEmetteur() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Emetteur
}

// GetEmetteurOk returns a tuple with the Emetteur field value
// and a boolean to check if the value has been set.
func (o *DonneesFactureSimplifiees) GetEmetteurOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Emetteur, true
}

// SetEmetteur sets field value
func (o *DonneesFactureSimplifiees) SetEmetteur(v map[string]interface{}) {
	o.Emetteur = v
}

// GetDestinataire returns the Destinataire field value
func (o *DonneesFactureSimplifiees) GetDestinataire() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Destinataire
}

// GetDestinataireOk returns a tuple with the Destinataire field value
// and a boolean to check if the value has been set.
func (o *DonneesFactureSimplifiees) GetDestinataireOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Destinataire, true
}

// SetDestinataire sets field value
func (o *DonneesFactureSimplifiees) SetDestinataire(v map[string]interface{}) {
	o.Destinataire = v
}

// GetLignes returns the Lignes field value
func (o *DonneesFactureSimplifiees) GetLignes() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Lignes
}

// GetLignesOk returns a tuple with the Lignes field value
// and a boolean to check if the value has been set.
func (o *DonneesFactureSimplifiees) GetLignesOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lignes, true
}

// SetLignes sets field value
func (o *DonneesFactureSimplifiees) SetLignes(v []map[string]interface{}) {
	o.Lignes = v
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DonneesFactureSimplifiees) GetDate() string {
	if o == nil || IsNil(o.Date.Get()) {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DonneesFactureSimplifiees) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *DonneesFactureSimplifiees) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableString and assigns it to the Date field.
func (o *DonneesFactureSimplifiees) SetDate(v string) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *DonneesFactureSimplifiees) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *DonneesFactureSimplifiees) UnsetDate() {
	o.Date.Unset()
}

// GetEcheanceJours returns the EcheanceJours field value if set, zero value otherwise.
func (o *DonneesFactureSimplifiees) GetEcheanceJours() int32 {
	if o == nil || IsNil(o.EcheanceJours) {
		var ret int32
		return ret
	}
	return *o.EcheanceJours
}

// GetEcheanceJoursOk returns a tuple with the EcheanceJours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DonneesFactureSimplifiees) GetEcheanceJoursOk() (*int32, bool) {
	if o == nil || IsNil(o.EcheanceJours) {
		return nil, false
	}
	return o.EcheanceJours, true
}

// HasEcheanceJours returns a boolean if a field has been set.
func (o *DonneesFactureSimplifiees) HasEcheanceJours() bool {
	if o != nil && !IsNil(o.EcheanceJours) {
		return true
	}

	return false
}

// SetEcheanceJours gets a reference to the given int32 and assigns it to the EcheanceJours field.
func (o *DonneesFactureSimplifiees) SetEcheanceJours(v int32) {
	o.EcheanceJours = &v
}

// GetCommentaire returns the Commentaire field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DonneesFactureSimplifiees) GetCommentaire() string {
	if o == nil || IsNil(o.Commentaire.Get()) {
		var ret string
		return ret
	}
	return *o.Commentaire.Get()
}

// GetCommentaireOk returns a tuple with the Commentaire field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DonneesFactureSimplifiees) GetCommentaireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Commentaire.Get(), o.Commentaire.IsSet()
}

// HasCommentaire returns a boolean if a field has been set.
func (o *DonneesFactureSimplifiees) HasCommentaire() bool {
	if o != nil && o.Commentaire.IsSet() {
		return true
	}

	return false
}

// SetCommentaire gets a reference to the given NullableString and assigns it to the Commentaire field.
func (o *DonneesFactureSimplifiees) SetCommentaire(v string) {
	o.Commentaire.Set(&v)
}
// SetCommentaireNil sets the value for Commentaire to be an explicit nil
func (o *DonneesFactureSimplifiees) SetCommentaireNil() {
	o.Commentaire.Set(nil)
}

// UnsetCommentaire ensures that no value is present for Commentaire, not even an explicit nil
func (o *DonneesFactureSimplifiees) UnsetCommentaire() {
	o.Commentaire.Unset()
}

// GetNumeroBonCommande returns the NumeroBonCommande field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DonneesFactureSimplifiees) GetNumeroBonCommande() string {
	if o == nil || IsNil(o.NumeroBonCommande.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroBonCommande.Get()
}

// GetNumeroBonCommandeOk returns a tuple with the NumeroBonCommande field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DonneesFactureSimplifiees) GetNumeroBonCommandeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroBonCommande.Get(), o.NumeroBonCommande.IsSet()
}

// HasNumeroBonCommande returns a boolean if a field has been set.
func (o *DonneesFactureSimplifiees) HasNumeroBonCommande() bool {
	if o != nil && o.NumeroBonCommande.IsSet() {
		return true
	}

	return false
}

// SetNumeroBonCommande gets a reference to the given NullableString and assigns it to the NumeroBonCommande field.
func (o *DonneesFactureSimplifiees) SetNumeroBonCommande(v string) {
	o.NumeroBonCommande.Set(&v)
}
// SetNumeroBonCommandeNil sets the value for NumeroBonCommande to be an explicit nil
func (o *DonneesFactureSimplifiees) SetNumeroBonCommandeNil() {
	o.NumeroBonCommande.Set(nil)
}

// UnsetNumeroBonCommande ensures that no value is present for NumeroBonCommande, not even an explicit nil
func (o *DonneesFactureSimplifiees) UnsetNumeroBonCommande() {
	o.NumeroBonCommande.Unset()
}

// GetNumeroMarche returns the NumeroMarche field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DonneesFactureSimplifiees) GetNumeroMarche() string {
	if o == nil || IsNil(o.NumeroMarche.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroMarche.Get()
}

// GetNumeroMarcheOk returns a tuple with the NumeroMarche field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DonneesFactureSimplifiees) GetNumeroMarcheOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroMarche.Get(), o.NumeroMarche.IsSet()
}

// HasNumeroMarche returns a boolean if a field has been set.
func (o *DonneesFactureSimplifiees) HasNumeroMarche() bool {
	if o != nil && o.NumeroMarche.IsSet() {
		return true
	}

	return false
}

// SetNumeroMarche gets a reference to the given NullableString and assigns it to the NumeroMarche field.
func (o *DonneesFactureSimplifiees) SetNumeroMarche(v string) {
	o.NumeroMarche.Set(&v)
}
// SetNumeroMarcheNil sets the value for NumeroMarche to be an explicit nil
func (o *DonneesFactureSimplifiees) SetNumeroMarcheNil() {
	o.NumeroMarche.Set(nil)
}

// UnsetNumeroMarche ensures that no value is present for NumeroMarche, not even an explicit nil
func (o *DonneesFactureSimplifiees) UnsetNumeroMarche() {
	o.NumeroMarche.Unset()
}

func (o DonneesFactureSimplifiees) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DonneesFactureSimplifiees) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["numero"] = o.Numero
	toSerialize["emetteur"] = o.Emetteur
	toSerialize["destinataire"] = o.Destinataire
	toSerialize["lignes"] = o.Lignes
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if !IsNil(o.EcheanceJours) {
		toSerialize["echeance_jours"] = o.EcheanceJours
	}
	if o.Commentaire.IsSet() {
		toSerialize["commentaire"] = o.Commentaire.Get()
	}
	if o.NumeroBonCommande.IsSet() {
		toSerialize["numero_bon_commande"] = o.NumeroBonCommande.Get()
	}
	if o.NumeroMarche.IsSet() {
		toSerialize["numero_marche"] = o.NumeroMarche.Get()
	}
	return toSerialize, nil
}

func (o *DonneesFactureSimplifiees) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"numero",
		"emetteur",
		"destinataire",
		"lignes",
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

	varDonneesFactureSimplifiees := _DonneesFactureSimplifiees{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDonneesFactureSimplifiees)

	if err != nil {
		return err
	}

	*o = DonneesFactureSimplifiees(varDonneesFactureSimplifiees)

	return err
}

type NullableDonneesFactureSimplifiees struct {
	value *DonneesFactureSimplifiees
	isSet bool
}

func (v NullableDonneesFactureSimplifiees) Get() *DonneesFactureSimplifiees {
	return v.value
}

func (v *NullableDonneesFactureSimplifiees) Set(val *DonneesFactureSimplifiees) {
	v.value = val
	v.isSet = true
}

func (v NullableDonneesFactureSimplifiees) IsSet() bool {
	return v.isSet
}

func (v *NullableDonneesFactureSimplifiees) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDonneesFactureSimplifiees(val *DonneesFactureSimplifiees) *NullableDonneesFactureSimplifiees {
	return &NullableDonneesFactureSimplifiees{value: val, isSet: true}
}

func (v NullableDonneesFactureSimplifiees) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDonneesFactureSimplifiees) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


