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

// checks if the SoumettreFactureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoumettreFactureRequest{}

// SoumettreFactureRequest Soumission d'une facture Chorus Pro.
type SoumettreFactureRequest struct {
	Credentials NullableChorusProCredentials `json:"credentials,omitempty"`
	// Numéro de la facture
	NumeroFacture string `json:"numero_facture"`
	// Date de facture (format ISO: YYYY-MM-DD)
	DateFacture string `json:"date_facture"`
	DateEcheancePaiement NullableString `json:"date_echeance_paiement,omitempty"`
	// ID Chorus Pro de la structure destinataire
	IdStructureCpp int32 `json:"id_structure_cpp"`
	CodeService NullableString `json:"code_service,omitempty"`
	NumeroEngagement NullableString `json:"numero_engagement,omitempty"`
	MontantHtTotal MontantHtTotal1 `json:"montant_ht_total"`
	MontantTva MontantTva1 `json:"montant_tva"`
	MontantTtcTotal MontantTtcTotal1 `json:"montant_ttc_total"`
	PieceJointePrincipaleId NullableInt32 `json:"piece_jointe_principale_id,omitempty"`
	PieceJointePrincipaleDesignation NullableString `json:"piece_jointe_principale_designation,omitempty"`
	Commentaire NullableString `json:"commentaire,omitempty"`
	NumeroBonCommande NullableString `json:"numero_bon_commande,omitempty"`
	NumeroMarche NullableString `json:"numero_marche,omitempty"`
}

type _SoumettreFactureRequest SoumettreFactureRequest

// NewSoumettreFactureRequest instantiates a new SoumettreFactureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoumettreFactureRequest(numeroFacture string, dateFacture string, idStructureCpp int32, montantHtTotal MontantHtTotal1, montantTva MontantTva1, montantTtcTotal MontantTtcTotal1) *SoumettreFactureRequest {
	this := SoumettreFactureRequest{}
	this.NumeroFacture = numeroFacture
	this.DateFacture = dateFacture
	this.IdStructureCpp = idStructureCpp
	this.MontantHtTotal = montantHtTotal
	this.MontantTva = montantTva
	this.MontantTtcTotal = montantTtcTotal
	return &this
}

// NewSoumettreFactureRequestWithDefaults instantiates a new SoumettreFactureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoumettreFactureRequestWithDefaults() *SoumettreFactureRequest {
	this := SoumettreFactureRequest{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureRequest) GetCredentials() ChorusProCredentials {
	if o == nil || IsNil(o.Credentials.Get()) {
		var ret ChorusProCredentials
		return ret
	}
	return *o.Credentials.Get()
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureRequest) GetCredentialsOk() (*ChorusProCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials.Get(), o.Credentials.IsSet()
}

// HasCredentials returns a boolean if a field has been set.
func (o *SoumettreFactureRequest) HasCredentials() bool {
	if o != nil && o.Credentials.IsSet() {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given NullableChorusProCredentials and assigns it to the Credentials field.
func (o *SoumettreFactureRequest) SetCredentials(v ChorusProCredentials) {
	o.Credentials.Set(&v)
}
// SetCredentialsNil sets the value for Credentials to be an explicit nil
func (o *SoumettreFactureRequest) SetCredentialsNil() {
	o.Credentials.Set(nil)
}

// UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
func (o *SoumettreFactureRequest) UnsetCredentials() {
	o.Credentials.Unset()
}

// GetNumeroFacture returns the NumeroFacture field value
func (o *SoumettreFactureRequest) GetNumeroFacture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NumeroFacture
}

// GetNumeroFactureOk returns a tuple with the NumeroFacture field value
// and a boolean to check if the value has been set.
func (o *SoumettreFactureRequest) GetNumeroFactureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumeroFacture, true
}

// SetNumeroFacture sets field value
func (o *SoumettreFactureRequest) SetNumeroFacture(v string) {
	o.NumeroFacture = v
}

// GetDateFacture returns the DateFacture field value
func (o *SoumettreFactureRequest) GetDateFacture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateFacture
}

// GetDateFactureOk returns a tuple with the DateFacture field value
// and a boolean to check if the value has been set.
func (o *SoumettreFactureRequest) GetDateFactureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateFacture, true
}

// SetDateFacture sets field value
func (o *SoumettreFactureRequest) SetDateFacture(v string) {
	o.DateFacture = v
}

// GetDateEcheancePaiement returns the DateEcheancePaiement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureRequest) GetDateEcheancePaiement() string {
	if o == nil || IsNil(o.DateEcheancePaiement.Get()) {
		var ret string
		return ret
	}
	return *o.DateEcheancePaiement.Get()
}

// GetDateEcheancePaiementOk returns a tuple with the DateEcheancePaiement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureRequest) GetDateEcheancePaiementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateEcheancePaiement.Get(), o.DateEcheancePaiement.IsSet()
}

// HasDateEcheancePaiement returns a boolean if a field has been set.
func (o *SoumettreFactureRequest) HasDateEcheancePaiement() bool {
	if o != nil && o.DateEcheancePaiement.IsSet() {
		return true
	}

	return false
}

// SetDateEcheancePaiement gets a reference to the given NullableString and assigns it to the DateEcheancePaiement field.
func (o *SoumettreFactureRequest) SetDateEcheancePaiement(v string) {
	o.DateEcheancePaiement.Set(&v)
}
// SetDateEcheancePaiementNil sets the value for DateEcheancePaiement to be an explicit nil
func (o *SoumettreFactureRequest) SetDateEcheancePaiementNil() {
	o.DateEcheancePaiement.Set(nil)
}

// UnsetDateEcheancePaiement ensures that no value is present for DateEcheancePaiement, not even an explicit nil
func (o *SoumettreFactureRequest) UnsetDateEcheancePaiement() {
	o.DateEcheancePaiement.Unset()
}

// GetIdStructureCpp returns the IdStructureCpp field value
func (o *SoumettreFactureRequest) GetIdStructureCpp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IdStructureCpp
}

// GetIdStructureCppOk returns a tuple with the IdStructureCpp field value
// and a boolean to check if the value has been set.
func (o *SoumettreFactureRequest) GetIdStructureCppOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdStructureCpp, true
}

// SetIdStructureCpp sets field value
func (o *SoumettreFactureRequest) SetIdStructureCpp(v int32) {
	o.IdStructureCpp = v
}

// GetCodeService returns the CodeService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureRequest) GetCodeService() string {
	if o == nil || IsNil(o.CodeService.Get()) {
		var ret string
		return ret
	}
	return *o.CodeService.Get()
}

// GetCodeServiceOk returns a tuple with the CodeService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureRequest) GetCodeServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeService.Get(), o.CodeService.IsSet()
}

// HasCodeService returns a boolean if a field has been set.
func (o *SoumettreFactureRequest) HasCodeService() bool {
	if o != nil && o.CodeService.IsSet() {
		return true
	}

	return false
}

// SetCodeService gets a reference to the given NullableString and assigns it to the CodeService field.
func (o *SoumettreFactureRequest) SetCodeService(v string) {
	o.CodeService.Set(&v)
}
// SetCodeServiceNil sets the value for CodeService to be an explicit nil
func (o *SoumettreFactureRequest) SetCodeServiceNil() {
	o.CodeService.Set(nil)
}

// UnsetCodeService ensures that no value is present for CodeService, not even an explicit nil
func (o *SoumettreFactureRequest) UnsetCodeService() {
	o.CodeService.Unset()
}

// GetNumeroEngagement returns the NumeroEngagement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureRequest) GetNumeroEngagement() string {
	if o == nil || IsNil(o.NumeroEngagement.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroEngagement.Get()
}

// GetNumeroEngagementOk returns a tuple with the NumeroEngagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureRequest) GetNumeroEngagementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroEngagement.Get(), o.NumeroEngagement.IsSet()
}

// HasNumeroEngagement returns a boolean if a field has been set.
func (o *SoumettreFactureRequest) HasNumeroEngagement() bool {
	if o != nil && o.NumeroEngagement.IsSet() {
		return true
	}

	return false
}

// SetNumeroEngagement gets a reference to the given NullableString and assigns it to the NumeroEngagement field.
func (o *SoumettreFactureRequest) SetNumeroEngagement(v string) {
	o.NumeroEngagement.Set(&v)
}
// SetNumeroEngagementNil sets the value for NumeroEngagement to be an explicit nil
func (o *SoumettreFactureRequest) SetNumeroEngagementNil() {
	o.NumeroEngagement.Set(nil)
}

// UnsetNumeroEngagement ensures that no value is present for NumeroEngagement, not even an explicit nil
func (o *SoumettreFactureRequest) UnsetNumeroEngagement() {
	o.NumeroEngagement.Unset()
}

// GetMontantHtTotal returns the MontantHtTotal field value
func (o *SoumettreFactureRequest) GetMontantHtTotal() MontantHtTotal1 {
	if o == nil {
		var ret MontantHtTotal1
		return ret
	}

	return o.MontantHtTotal
}

// GetMontantHtTotalOk returns a tuple with the MontantHtTotal field value
// and a boolean to check if the value has been set.
func (o *SoumettreFactureRequest) GetMontantHtTotalOk() (*MontantHtTotal1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantHtTotal, true
}

// SetMontantHtTotal sets field value
func (o *SoumettreFactureRequest) SetMontantHtTotal(v MontantHtTotal1) {
	o.MontantHtTotal = v
}

// GetMontantTva returns the MontantTva field value
func (o *SoumettreFactureRequest) GetMontantTva() MontantTva1 {
	if o == nil {
		var ret MontantTva1
		return ret
	}

	return o.MontantTva
}

// GetMontantTvaOk returns a tuple with the MontantTva field value
// and a boolean to check if the value has been set.
func (o *SoumettreFactureRequest) GetMontantTvaOk() (*MontantTva1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantTva, true
}

// SetMontantTva sets field value
func (o *SoumettreFactureRequest) SetMontantTva(v MontantTva1) {
	o.MontantTva = v
}

// GetMontantTtcTotal returns the MontantTtcTotal field value
func (o *SoumettreFactureRequest) GetMontantTtcTotal() MontantTtcTotal1 {
	if o == nil {
		var ret MontantTtcTotal1
		return ret
	}

	return o.MontantTtcTotal
}

// GetMontantTtcTotalOk returns a tuple with the MontantTtcTotal field value
// and a boolean to check if the value has been set.
func (o *SoumettreFactureRequest) GetMontantTtcTotalOk() (*MontantTtcTotal1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantTtcTotal, true
}

// SetMontantTtcTotal sets field value
func (o *SoumettreFactureRequest) SetMontantTtcTotal(v MontantTtcTotal1) {
	o.MontantTtcTotal = v
}

// GetPieceJointePrincipaleId returns the PieceJointePrincipaleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureRequest) GetPieceJointePrincipaleId() int32 {
	if o == nil || IsNil(o.PieceJointePrincipaleId.Get()) {
		var ret int32
		return ret
	}
	return *o.PieceJointePrincipaleId.Get()
}

// GetPieceJointePrincipaleIdOk returns a tuple with the PieceJointePrincipaleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureRequest) GetPieceJointePrincipaleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PieceJointePrincipaleId.Get(), o.PieceJointePrincipaleId.IsSet()
}

// HasPieceJointePrincipaleId returns a boolean if a field has been set.
func (o *SoumettreFactureRequest) HasPieceJointePrincipaleId() bool {
	if o != nil && o.PieceJointePrincipaleId.IsSet() {
		return true
	}

	return false
}

// SetPieceJointePrincipaleId gets a reference to the given NullableInt32 and assigns it to the PieceJointePrincipaleId field.
func (o *SoumettreFactureRequest) SetPieceJointePrincipaleId(v int32) {
	o.PieceJointePrincipaleId.Set(&v)
}
// SetPieceJointePrincipaleIdNil sets the value for PieceJointePrincipaleId to be an explicit nil
func (o *SoumettreFactureRequest) SetPieceJointePrincipaleIdNil() {
	o.PieceJointePrincipaleId.Set(nil)
}

// UnsetPieceJointePrincipaleId ensures that no value is present for PieceJointePrincipaleId, not even an explicit nil
func (o *SoumettreFactureRequest) UnsetPieceJointePrincipaleId() {
	o.PieceJointePrincipaleId.Unset()
}

// GetPieceJointePrincipaleDesignation returns the PieceJointePrincipaleDesignation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureRequest) GetPieceJointePrincipaleDesignation() string {
	if o == nil || IsNil(o.PieceJointePrincipaleDesignation.Get()) {
		var ret string
		return ret
	}
	return *o.PieceJointePrincipaleDesignation.Get()
}

// GetPieceJointePrincipaleDesignationOk returns a tuple with the PieceJointePrincipaleDesignation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureRequest) GetPieceJointePrincipaleDesignationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PieceJointePrincipaleDesignation.Get(), o.PieceJointePrincipaleDesignation.IsSet()
}

// HasPieceJointePrincipaleDesignation returns a boolean if a field has been set.
func (o *SoumettreFactureRequest) HasPieceJointePrincipaleDesignation() bool {
	if o != nil && o.PieceJointePrincipaleDesignation.IsSet() {
		return true
	}

	return false
}

// SetPieceJointePrincipaleDesignation gets a reference to the given NullableString and assigns it to the PieceJointePrincipaleDesignation field.
func (o *SoumettreFactureRequest) SetPieceJointePrincipaleDesignation(v string) {
	o.PieceJointePrincipaleDesignation.Set(&v)
}
// SetPieceJointePrincipaleDesignationNil sets the value for PieceJointePrincipaleDesignation to be an explicit nil
func (o *SoumettreFactureRequest) SetPieceJointePrincipaleDesignationNil() {
	o.PieceJointePrincipaleDesignation.Set(nil)
}

// UnsetPieceJointePrincipaleDesignation ensures that no value is present for PieceJointePrincipaleDesignation, not even an explicit nil
func (o *SoumettreFactureRequest) UnsetPieceJointePrincipaleDesignation() {
	o.PieceJointePrincipaleDesignation.Unset()
}

// GetCommentaire returns the Commentaire field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureRequest) GetCommentaire() string {
	if o == nil || IsNil(o.Commentaire.Get()) {
		var ret string
		return ret
	}
	return *o.Commentaire.Get()
}

// GetCommentaireOk returns a tuple with the Commentaire field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureRequest) GetCommentaireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Commentaire.Get(), o.Commentaire.IsSet()
}

// HasCommentaire returns a boolean if a field has been set.
func (o *SoumettreFactureRequest) HasCommentaire() bool {
	if o != nil && o.Commentaire.IsSet() {
		return true
	}

	return false
}

// SetCommentaire gets a reference to the given NullableString and assigns it to the Commentaire field.
func (o *SoumettreFactureRequest) SetCommentaire(v string) {
	o.Commentaire.Set(&v)
}
// SetCommentaireNil sets the value for Commentaire to be an explicit nil
func (o *SoumettreFactureRequest) SetCommentaireNil() {
	o.Commentaire.Set(nil)
}

// UnsetCommentaire ensures that no value is present for Commentaire, not even an explicit nil
func (o *SoumettreFactureRequest) UnsetCommentaire() {
	o.Commentaire.Unset()
}

// GetNumeroBonCommande returns the NumeroBonCommande field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureRequest) GetNumeroBonCommande() string {
	if o == nil || IsNil(o.NumeroBonCommande.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroBonCommande.Get()
}

// GetNumeroBonCommandeOk returns a tuple with the NumeroBonCommande field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureRequest) GetNumeroBonCommandeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroBonCommande.Get(), o.NumeroBonCommande.IsSet()
}

// HasNumeroBonCommande returns a boolean if a field has been set.
func (o *SoumettreFactureRequest) HasNumeroBonCommande() bool {
	if o != nil && o.NumeroBonCommande.IsSet() {
		return true
	}

	return false
}

// SetNumeroBonCommande gets a reference to the given NullableString and assigns it to the NumeroBonCommande field.
func (o *SoumettreFactureRequest) SetNumeroBonCommande(v string) {
	o.NumeroBonCommande.Set(&v)
}
// SetNumeroBonCommandeNil sets the value for NumeroBonCommande to be an explicit nil
func (o *SoumettreFactureRequest) SetNumeroBonCommandeNil() {
	o.NumeroBonCommande.Set(nil)
}

// UnsetNumeroBonCommande ensures that no value is present for NumeroBonCommande, not even an explicit nil
func (o *SoumettreFactureRequest) UnsetNumeroBonCommande() {
	o.NumeroBonCommande.Unset()
}

// GetNumeroMarche returns the NumeroMarche field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoumettreFactureRequest) GetNumeroMarche() string {
	if o == nil || IsNil(o.NumeroMarche.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroMarche.Get()
}

// GetNumeroMarcheOk returns a tuple with the NumeroMarche field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoumettreFactureRequest) GetNumeroMarcheOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroMarche.Get(), o.NumeroMarche.IsSet()
}

// HasNumeroMarche returns a boolean if a field has been set.
func (o *SoumettreFactureRequest) HasNumeroMarche() bool {
	if o != nil && o.NumeroMarche.IsSet() {
		return true
	}

	return false
}

// SetNumeroMarche gets a reference to the given NullableString and assigns it to the NumeroMarche field.
func (o *SoumettreFactureRequest) SetNumeroMarche(v string) {
	o.NumeroMarche.Set(&v)
}
// SetNumeroMarcheNil sets the value for NumeroMarche to be an explicit nil
func (o *SoumettreFactureRequest) SetNumeroMarcheNil() {
	o.NumeroMarche.Set(nil)
}

// UnsetNumeroMarche ensures that no value is present for NumeroMarche, not even an explicit nil
func (o *SoumettreFactureRequest) UnsetNumeroMarche() {
	o.NumeroMarche.Unset()
}

func (o SoumettreFactureRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoumettreFactureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Credentials.IsSet() {
		toSerialize["credentials"] = o.Credentials.Get()
	}
	toSerialize["numero_facture"] = o.NumeroFacture
	toSerialize["date_facture"] = o.DateFacture
	if o.DateEcheancePaiement.IsSet() {
		toSerialize["date_echeance_paiement"] = o.DateEcheancePaiement.Get()
	}
	toSerialize["id_structure_cpp"] = o.IdStructureCpp
	if o.CodeService.IsSet() {
		toSerialize["code_service"] = o.CodeService.Get()
	}
	if o.NumeroEngagement.IsSet() {
		toSerialize["numero_engagement"] = o.NumeroEngagement.Get()
	}
	toSerialize["montant_ht_total"] = o.MontantHtTotal
	toSerialize["montant_tva"] = o.MontantTva
	toSerialize["montant_ttc_total"] = o.MontantTtcTotal
	if o.PieceJointePrincipaleId.IsSet() {
		toSerialize["piece_jointe_principale_id"] = o.PieceJointePrincipaleId.Get()
	}
	if o.PieceJointePrincipaleDesignation.IsSet() {
		toSerialize["piece_jointe_principale_designation"] = o.PieceJointePrincipaleDesignation.Get()
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

func (o *SoumettreFactureRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"numero_facture",
		"date_facture",
		"id_structure_cpp",
		"montant_ht_total",
		"montant_tva",
		"montant_ttc_total",
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

	varSoumettreFactureRequest := _SoumettreFactureRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSoumettreFactureRequest)

	if err != nil {
		return err
	}

	*o = SoumettreFactureRequest(varSoumettreFactureRequest)

	return err
}

type NullableSoumettreFactureRequest struct {
	value *SoumettreFactureRequest
	isSet bool
}

func (v NullableSoumettreFactureRequest) Get() *SoumettreFactureRequest {
	return v.value
}

func (v *NullableSoumettreFactureRequest) Set(val *SoumettreFactureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSoumettreFactureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSoumettreFactureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoumettreFactureRequest(val *SoumettreFactureRequest) *NullableSoumettreFactureRequest {
	return &NullableSoumettreFactureRequest{value: val, isSet: true}
}

func (v NullableSoumettreFactureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoumettreFactureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


