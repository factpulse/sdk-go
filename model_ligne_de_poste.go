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

// checks if the LigneDePoste type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LigneDePoste{}

// LigneDePoste Représente une ligne de détail dans une facture.
type LigneDePoste struct {
	Numero int32 `json:"numero"`
	Reference NullableString `json:"reference,omitempty"`
	Denomination string `json:"denomination"`
	Quantite Quantite `json:"quantite"`
	Unite Unite `json:"unite"`
	MontantUnitaireHt MontantUnitaireHt `json:"montantUnitaireHt"`
	// Montant de la remise HT.
	MontantRemiseHt NullableFloat64 `json:"montantRemiseHt,omitempty" validate:"regexp=^(?!^[-+.]*$)[+-]?0*(?:\\\\d{0,8}|(?=[\\\\d.]{1,13}0*$)\\\\d{0,8}\\\\.\\\\d{0,4}0*$)"`
	// Montant total HT de la ligne (quantité × prix unitaire - remise).
	MontantTotalLigneHt *float64 `json:"montantTotalLigneHt,omitempty" validate:"regexp=^(?!^[-+.]*$)[+-]?0*(?:\\\\d{0,10}|(?=[\\\\d.]{1,13}0*$)\\\\d{0,10}\\\\.\\\\d{0,2}0*$)"`
	TauxTva NullableString `json:"tauxTva,omitempty"`
	// Taux de TVA avec valeur manuelle.
	TauxTvaManuel NullableFloat64 `json:"tauxTvaManuel,omitempty" validate:"regexp=^(?!^[-+.]*$)[+-]?0*(?:\\\\d{0,8}|(?=[\\\\d.]{1,13}0*$)\\\\d{0,8}\\\\.\\\\d{0,4}0*$)"`
	CategorieTva NullableCategorieTVA `json:"categorieTva,omitempty"`
	DateDebutPeriode NullableString `json:"dateDebutPeriode,omitempty"`
	DateFinPeriode NullableString `json:"dateFinPeriode,omitempty"`
	CodeRaisonReduction NullableCodeRaisonReduction `json:"codeRaisonReduction,omitempty"`
	RaisonReduction NullableString `json:"raisonReduction,omitempty"`
}

type _LigneDePoste LigneDePoste

// NewLigneDePoste instantiates a new LigneDePoste object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLigneDePoste(numero int32, denomination string, quantite Quantite, unite Unite, montantUnitaireHt MontantUnitaireHt) *LigneDePoste {
	this := LigneDePoste{}
	this.Numero = numero
	this.Denomination = denomination
	this.Quantite = quantite
	this.Unite = unite
	this.MontantUnitaireHt = montantUnitaireHt
	return &this
}

// NewLigneDePosteWithDefaults instantiates a new LigneDePoste object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLigneDePosteWithDefaults() *LigneDePoste {
	this := LigneDePoste{}
	return &this
}

// GetNumero returns the Numero field value
func (o *LigneDePoste) GetNumero() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Numero
}

// GetNumeroOk returns a tuple with the Numero field value
// and a boolean to check if the value has been set.
func (o *LigneDePoste) GetNumeroOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Numero, true
}

// SetNumero sets field value
func (o *LigneDePoste) SetNumero(v int32) {
	o.Numero = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LigneDePoste) GetReference() string {
	if o == nil || IsNil(o.Reference.Get()) {
		var ret string
		return ret
	}
	return *o.Reference.Get()
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LigneDePoste) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reference.Get(), o.Reference.IsSet()
}

// HasReference returns a boolean if a field has been set.
func (o *LigneDePoste) HasReference() bool {
	if o != nil && o.Reference.IsSet() {
		return true
	}

	return false
}

// SetReference gets a reference to the given NullableString and assigns it to the Reference field.
func (o *LigneDePoste) SetReference(v string) {
	o.Reference.Set(&v)
}
// SetReferenceNil sets the value for Reference to be an explicit nil
func (o *LigneDePoste) SetReferenceNil() {
	o.Reference.Set(nil)
}

// UnsetReference ensures that no value is present for Reference, not even an explicit nil
func (o *LigneDePoste) UnsetReference() {
	o.Reference.Unset()
}

// GetDenomination returns the Denomination field value
func (o *LigneDePoste) GetDenomination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Denomination
}

// GetDenominationOk returns a tuple with the Denomination field value
// and a boolean to check if the value has been set.
func (o *LigneDePoste) GetDenominationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Denomination, true
}

// SetDenomination sets field value
func (o *LigneDePoste) SetDenomination(v string) {
	o.Denomination = v
}

// GetQuantite returns the Quantite field value
func (o *LigneDePoste) GetQuantite() Quantite {
	if o == nil {
		var ret Quantite
		return ret
	}

	return o.Quantite
}

// GetQuantiteOk returns a tuple with the Quantite field value
// and a boolean to check if the value has been set.
func (o *LigneDePoste) GetQuantiteOk() (*Quantite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantite, true
}

// SetQuantite sets field value
func (o *LigneDePoste) SetQuantite(v Quantite) {
	o.Quantite = v
}

// GetUnite returns the Unite field value
func (o *LigneDePoste) GetUnite() Unite {
	if o == nil {
		var ret Unite
		return ret
	}

	return o.Unite
}

// GetUniteOk returns a tuple with the Unite field value
// and a boolean to check if the value has been set.
func (o *LigneDePoste) GetUniteOk() (*Unite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unite, true
}

// SetUnite sets field value
func (o *LigneDePoste) SetUnite(v Unite) {
	o.Unite = v
}

// GetMontantUnitaireHt returns the MontantUnitaireHt field value
func (o *LigneDePoste) GetMontantUnitaireHt() MontantUnitaireHt {
	if o == nil {
		var ret MontantUnitaireHt
		return ret
	}

	return o.MontantUnitaireHt
}

// GetMontantUnitaireHtOk returns a tuple with the MontantUnitaireHt field value
// and a boolean to check if the value has been set.
func (o *LigneDePoste) GetMontantUnitaireHtOk() (*MontantUnitaireHt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantUnitaireHt, true
}

// SetMontantUnitaireHt sets field value
func (o *LigneDePoste) SetMontantUnitaireHt(v MontantUnitaireHt) {
	o.MontantUnitaireHt = v
}

// GetMontantRemiseHt returns the MontantRemiseHt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LigneDePoste) GetMontantRemiseHt() float64 {
	if o == nil || IsNil(o.MontantRemiseHt.Get()) {
		var ret float64
		return ret
	}
	return *o.MontantRemiseHt.Get()
}

// GetMontantRemiseHtOk returns a tuple with the MontantRemiseHt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LigneDePoste) GetMontantRemiseHtOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MontantRemiseHt.Get(), o.MontantRemiseHt.IsSet()
}

// HasMontantRemiseHt returns a boolean if a field has been set.
func (o *LigneDePoste) HasMontantRemiseHt() bool {
	if o != nil && o.MontantRemiseHt.IsSet() {
		return true
	}

	return false
}

// SetMontantRemiseHt gets a reference to the given NullableFloat64 and assigns it to the MontantRemiseHt field.
func (o *LigneDePoste) SetMontantRemiseHt(v float64) {
	o.MontantRemiseHt.Set(&v)
}
// SetMontantRemiseHtNil sets the value for MontantRemiseHt to be an explicit nil
func (o *LigneDePoste) SetMontantRemiseHtNil() {
	o.MontantRemiseHt.Set(nil)
}

// UnsetMontantRemiseHt ensures that no value is present for MontantRemiseHt, not even an explicit nil
func (o *LigneDePoste) UnsetMontantRemiseHt() {
	o.MontantRemiseHt.Unset()
}

// GetMontantTotalLigneHt returns the MontantTotalLigneHt field value if set, zero value otherwise.
func (o *LigneDePoste) GetMontantTotalLigneHt() float64 {
	if o == nil || IsNil(o.MontantTotalLigneHt) {
		var ret float64
		return ret
	}
	return *o.MontantTotalLigneHt
}

// GetMontantTotalLigneHtOk returns a tuple with the MontantTotalLigneHt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LigneDePoste) GetMontantTotalLigneHtOk() (*float64, bool) {
	if o == nil || IsNil(o.MontantTotalLigneHt) {
		return nil, false
	}
	return o.MontantTotalLigneHt, true
}

// HasMontantTotalLigneHt returns a boolean if a field has been set.
func (o *LigneDePoste) HasMontantTotalLigneHt() bool {
	if o != nil && !IsNil(o.MontantTotalLigneHt) {
		return true
	}

	return false
}

// SetMontantTotalLigneHt gets a reference to the given float64 and assigns it to the MontantTotalLigneHt field.
func (o *LigneDePoste) SetMontantTotalLigneHt(v float64) {
	o.MontantTotalLigneHt = &v
}

// GetTauxTva returns the TauxTva field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LigneDePoste) GetTauxTva() string {
	if o == nil || IsNil(o.TauxTva.Get()) {
		var ret string
		return ret
	}
	return *o.TauxTva.Get()
}

// GetTauxTvaOk returns a tuple with the TauxTva field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LigneDePoste) GetTauxTvaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TauxTva.Get(), o.TauxTva.IsSet()
}

// HasTauxTva returns a boolean if a field has been set.
func (o *LigneDePoste) HasTauxTva() bool {
	if o != nil && o.TauxTva.IsSet() {
		return true
	}

	return false
}

// SetTauxTva gets a reference to the given NullableString and assigns it to the TauxTva field.
func (o *LigneDePoste) SetTauxTva(v string) {
	o.TauxTva.Set(&v)
}
// SetTauxTvaNil sets the value for TauxTva to be an explicit nil
func (o *LigneDePoste) SetTauxTvaNil() {
	o.TauxTva.Set(nil)
}

// UnsetTauxTva ensures that no value is present for TauxTva, not even an explicit nil
func (o *LigneDePoste) UnsetTauxTva() {
	o.TauxTva.Unset()
}

// GetTauxTvaManuel returns the TauxTvaManuel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LigneDePoste) GetTauxTvaManuel() float64 {
	if o == nil || IsNil(o.TauxTvaManuel.Get()) {
		var ret float64
		return ret
	}
	return *o.TauxTvaManuel.Get()
}

// GetTauxTvaManuelOk returns a tuple with the TauxTvaManuel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LigneDePoste) GetTauxTvaManuelOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TauxTvaManuel.Get(), o.TauxTvaManuel.IsSet()
}

// HasTauxTvaManuel returns a boolean if a field has been set.
func (o *LigneDePoste) HasTauxTvaManuel() bool {
	if o != nil && o.TauxTvaManuel.IsSet() {
		return true
	}

	return false
}

// SetTauxTvaManuel gets a reference to the given NullableFloat64 and assigns it to the TauxTvaManuel field.
func (o *LigneDePoste) SetTauxTvaManuel(v float64) {
	o.TauxTvaManuel.Set(&v)
}
// SetTauxTvaManuelNil sets the value for TauxTvaManuel to be an explicit nil
func (o *LigneDePoste) SetTauxTvaManuelNil() {
	o.TauxTvaManuel.Set(nil)
}

// UnsetTauxTvaManuel ensures that no value is present for TauxTvaManuel, not even an explicit nil
func (o *LigneDePoste) UnsetTauxTvaManuel() {
	o.TauxTvaManuel.Unset()
}

// GetCategorieTva returns the CategorieTva field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LigneDePoste) GetCategorieTva() CategorieTVA {
	if o == nil || IsNil(o.CategorieTva.Get()) {
		var ret CategorieTVA
		return ret
	}
	return *o.CategorieTva.Get()
}

// GetCategorieTvaOk returns a tuple with the CategorieTva field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LigneDePoste) GetCategorieTvaOk() (*CategorieTVA, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategorieTva.Get(), o.CategorieTva.IsSet()
}

// HasCategorieTva returns a boolean if a field has been set.
func (o *LigneDePoste) HasCategorieTva() bool {
	if o != nil && o.CategorieTva.IsSet() {
		return true
	}

	return false
}

// SetCategorieTva gets a reference to the given NullableCategorieTVA and assigns it to the CategorieTva field.
func (o *LigneDePoste) SetCategorieTva(v CategorieTVA) {
	o.CategorieTva.Set(&v)
}
// SetCategorieTvaNil sets the value for CategorieTva to be an explicit nil
func (o *LigneDePoste) SetCategorieTvaNil() {
	o.CategorieTva.Set(nil)
}

// UnsetCategorieTva ensures that no value is present for CategorieTva, not even an explicit nil
func (o *LigneDePoste) UnsetCategorieTva() {
	o.CategorieTva.Unset()
}

// GetDateDebutPeriode returns the DateDebutPeriode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LigneDePoste) GetDateDebutPeriode() string {
	if o == nil || IsNil(o.DateDebutPeriode.Get()) {
		var ret string
		return ret
	}
	return *o.DateDebutPeriode.Get()
}

// GetDateDebutPeriodeOk returns a tuple with the DateDebutPeriode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LigneDePoste) GetDateDebutPeriodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateDebutPeriode.Get(), o.DateDebutPeriode.IsSet()
}

// HasDateDebutPeriode returns a boolean if a field has been set.
func (o *LigneDePoste) HasDateDebutPeriode() bool {
	if o != nil && o.DateDebutPeriode.IsSet() {
		return true
	}

	return false
}

// SetDateDebutPeriode gets a reference to the given NullableString and assigns it to the DateDebutPeriode field.
func (o *LigneDePoste) SetDateDebutPeriode(v string) {
	o.DateDebutPeriode.Set(&v)
}
// SetDateDebutPeriodeNil sets the value for DateDebutPeriode to be an explicit nil
func (o *LigneDePoste) SetDateDebutPeriodeNil() {
	o.DateDebutPeriode.Set(nil)
}

// UnsetDateDebutPeriode ensures that no value is present for DateDebutPeriode, not even an explicit nil
func (o *LigneDePoste) UnsetDateDebutPeriode() {
	o.DateDebutPeriode.Unset()
}

// GetDateFinPeriode returns the DateFinPeriode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LigneDePoste) GetDateFinPeriode() string {
	if o == nil || IsNil(o.DateFinPeriode.Get()) {
		var ret string
		return ret
	}
	return *o.DateFinPeriode.Get()
}

// GetDateFinPeriodeOk returns a tuple with the DateFinPeriode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LigneDePoste) GetDateFinPeriodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateFinPeriode.Get(), o.DateFinPeriode.IsSet()
}

// HasDateFinPeriode returns a boolean if a field has been set.
func (o *LigneDePoste) HasDateFinPeriode() bool {
	if o != nil && o.DateFinPeriode.IsSet() {
		return true
	}

	return false
}

// SetDateFinPeriode gets a reference to the given NullableString and assigns it to the DateFinPeriode field.
func (o *LigneDePoste) SetDateFinPeriode(v string) {
	o.DateFinPeriode.Set(&v)
}
// SetDateFinPeriodeNil sets the value for DateFinPeriode to be an explicit nil
func (o *LigneDePoste) SetDateFinPeriodeNil() {
	o.DateFinPeriode.Set(nil)
}

// UnsetDateFinPeriode ensures that no value is present for DateFinPeriode, not even an explicit nil
func (o *LigneDePoste) UnsetDateFinPeriode() {
	o.DateFinPeriode.Unset()
}

// GetCodeRaisonReduction returns the CodeRaisonReduction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LigneDePoste) GetCodeRaisonReduction() CodeRaisonReduction {
	if o == nil || IsNil(o.CodeRaisonReduction.Get()) {
		var ret CodeRaisonReduction
		return ret
	}
	return *o.CodeRaisonReduction.Get()
}

// GetCodeRaisonReductionOk returns a tuple with the CodeRaisonReduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LigneDePoste) GetCodeRaisonReductionOk() (*CodeRaisonReduction, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeRaisonReduction.Get(), o.CodeRaisonReduction.IsSet()
}

// HasCodeRaisonReduction returns a boolean if a field has been set.
func (o *LigneDePoste) HasCodeRaisonReduction() bool {
	if o != nil && o.CodeRaisonReduction.IsSet() {
		return true
	}

	return false
}

// SetCodeRaisonReduction gets a reference to the given NullableCodeRaisonReduction and assigns it to the CodeRaisonReduction field.
func (o *LigneDePoste) SetCodeRaisonReduction(v CodeRaisonReduction) {
	o.CodeRaisonReduction.Set(&v)
}
// SetCodeRaisonReductionNil sets the value for CodeRaisonReduction to be an explicit nil
func (o *LigneDePoste) SetCodeRaisonReductionNil() {
	o.CodeRaisonReduction.Set(nil)
}

// UnsetCodeRaisonReduction ensures that no value is present for CodeRaisonReduction, not even an explicit nil
func (o *LigneDePoste) UnsetCodeRaisonReduction() {
	o.CodeRaisonReduction.Unset()
}

// GetRaisonReduction returns the RaisonReduction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LigneDePoste) GetRaisonReduction() string {
	if o == nil || IsNil(o.RaisonReduction.Get()) {
		var ret string
		return ret
	}
	return *o.RaisonReduction.Get()
}

// GetRaisonReductionOk returns a tuple with the RaisonReduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LigneDePoste) GetRaisonReductionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaisonReduction.Get(), o.RaisonReduction.IsSet()
}

// HasRaisonReduction returns a boolean if a field has been set.
func (o *LigneDePoste) HasRaisonReduction() bool {
	if o != nil && o.RaisonReduction.IsSet() {
		return true
	}

	return false
}

// SetRaisonReduction gets a reference to the given NullableString and assigns it to the RaisonReduction field.
func (o *LigneDePoste) SetRaisonReduction(v string) {
	o.RaisonReduction.Set(&v)
}
// SetRaisonReductionNil sets the value for RaisonReduction to be an explicit nil
func (o *LigneDePoste) SetRaisonReductionNil() {
	o.RaisonReduction.Set(nil)
}

// UnsetRaisonReduction ensures that no value is present for RaisonReduction, not even an explicit nil
func (o *LigneDePoste) UnsetRaisonReduction() {
	o.RaisonReduction.Unset()
}

func (o LigneDePoste) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LigneDePoste) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["numero"] = o.Numero
	if o.Reference.IsSet() {
		toSerialize["reference"] = o.Reference.Get()
	}
	toSerialize["denomination"] = o.Denomination
	toSerialize["quantite"] = o.Quantite
	toSerialize["unite"] = o.Unite
	toSerialize["montantUnitaireHt"] = o.MontantUnitaireHt
	if o.MontantRemiseHt.IsSet() {
		toSerialize["montantRemiseHt"] = o.MontantRemiseHt.Get()
	}
	if !IsNil(o.MontantTotalLigneHt) {
		toSerialize["montantTotalLigneHt"] = o.MontantTotalLigneHt
	}
	if o.TauxTva.IsSet() {
		toSerialize["tauxTva"] = o.TauxTva.Get()
	}
	if o.TauxTvaManuel.IsSet() {
		toSerialize["tauxTvaManuel"] = o.TauxTvaManuel.Get()
	}
	if o.CategorieTva.IsSet() {
		toSerialize["categorieTva"] = o.CategorieTva.Get()
	}
	if o.DateDebutPeriode.IsSet() {
		toSerialize["dateDebutPeriode"] = o.DateDebutPeriode.Get()
	}
	if o.DateFinPeriode.IsSet() {
		toSerialize["dateFinPeriode"] = o.DateFinPeriode.Get()
	}
	if o.CodeRaisonReduction.IsSet() {
		toSerialize["codeRaisonReduction"] = o.CodeRaisonReduction.Get()
	}
	if o.RaisonReduction.IsSet() {
		toSerialize["raisonReduction"] = o.RaisonReduction.Get()
	}
	return toSerialize, nil
}

func (o *LigneDePoste) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"numero",
		"denomination",
		"quantite",
		"unite",
		"montantUnitaireHt",
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

	varLigneDePoste := _LigneDePoste{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLigneDePoste)

	if err != nil {
		return err
	}

	*o = LigneDePoste(varLigneDePoste)

	return err
}

type NullableLigneDePoste struct {
	value *LigneDePoste
	isSet bool
}

func (v NullableLigneDePoste) Get() *LigneDePoste {
	return v.value
}

func (v *NullableLigneDePoste) Set(val *LigneDePoste) {
	v.value = val
	v.isSet = true
}

func (v NullableLigneDePoste) IsSet() bool {
	return v.isSet
}

func (v *NullableLigneDePoste) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLigneDePoste(val *LigneDePoste) *NullableLigneDePoste {
	return &NullableLigneDePoste{value: val, isSet: true}
}

func (v NullableLigneDePoste) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLigneDePoste) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


