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

// checks if the FactureEntrante type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FactureEntrante{}

// FactureEntrante Facture reçue d'un fournisseur via PDP/PA.  Ce modèle contient les métadonnées essentielles extraites des factures entrantes, quel que soit leur format source (CII, UBL, Factur-X).  Les montants sont en Decimal en Python mais seront sérialisés en string dans le JSON pour préserver la précision monétaire.
type FactureEntrante struct {
	FlowId NullableString `json:"flow_id,omitempty"`
	// Format source de la facture
	FormatSource FormatFacture `json:"format_source"`
	// Numéro de facture émis par le fournisseur (BT-1)
	RefFournisseur string `json:"ref_fournisseur"`
	// Type de document (BT-3)
	TypeDocument *TypeDocument `json:"type_document,omitempty"`
	// Émetteur de la facture (SellerTradeParty)
	Fournisseur FournisseurEntrant `json:"fournisseur"`
	// Nom du destinataire / votre entreprise (BT-44)
	SiteFacturationNom string `json:"site_facturation_nom"`
	SiteFacturationSiret NullableString `json:"site_facturation_siret,omitempty"`
	// Date de la facture (BT-2) - YYYY-MM-DD
	DateDePiece string `json:"date_de_piece"`
	DateReglement NullableString `json:"date_reglement,omitempty"`
	// Code devise ISO (BT-5)
	Devise *string `json:"devise,omitempty"`
	// Montant HT total (BT-109)
	MontantHt string `json:"montant_ht" validate:"regexp=^(?!^[-+.]*$)[+-]?0*\\\\d*\\\\.?\\\\d*$"`
	// Montant TVA total (BT-110)
	MontantTva string `json:"montant_tva" validate:"regexp=^(?!^[-+.]*$)[+-]?0*\\\\d*\\\\.?\\\\d*$"`
	// Montant TTC total (BT-112)
	MontantTtc string `json:"montant_ttc" validate:"regexp=^(?!^[-+.]*$)[+-]?0*\\\\d*\\\\.?\\\\d*$"`
	NumeroBonCommande NullableString `json:"numero_bon_commande,omitempty"`
	ReferenceContrat NullableString `json:"reference_contrat,omitempty"`
	ObjetFacture NullableString `json:"objet_facture,omitempty"`
	DocumentBase64 NullableString `json:"document_base64,omitempty"`
	DocumentContentType NullableString `json:"document_content_type,omitempty"`
	DocumentFilename NullableString `json:"document_filename,omitempty"`
}

type _FactureEntrante FactureEntrante

// NewFactureEntrante instantiates a new FactureEntrante object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFactureEntrante(formatSource FormatFacture, refFournisseur string, fournisseur FournisseurEntrant, siteFacturationNom string, dateDePiece string, montantHt string, montantTva string, montantTtc string) *FactureEntrante {
	this := FactureEntrante{}
	this.FormatSource = formatSource
	this.RefFournisseur = refFournisseur
	var typeDocument TypeDocument = _380
	this.TypeDocument = &typeDocument
	this.Fournisseur = fournisseur
	this.SiteFacturationNom = siteFacturationNom
	this.DateDePiece = dateDePiece
	var devise string = "EUR"
	this.Devise = &devise
	this.MontantHt = montantHt
	this.MontantTva = montantTva
	this.MontantTtc = montantTtc
	return &this
}

// NewFactureEntranteWithDefaults instantiates a new FactureEntrante object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFactureEntranteWithDefaults() *FactureEntrante {
	this := FactureEntrante{}
	var typeDocument TypeDocument = _380
	this.TypeDocument = &typeDocument
	var devise string = "EUR"
	this.Devise = &devise
	return &this
}

// GetFlowId returns the FlowId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureEntrante) GetFlowId() string {
	if o == nil || IsNil(o.FlowId.Get()) {
		var ret string
		return ret
	}
	return *o.FlowId.Get()
}

// GetFlowIdOk returns a tuple with the FlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureEntrante) GetFlowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowId.Get(), o.FlowId.IsSet()
}

// HasFlowId returns a boolean if a field has been set.
func (o *FactureEntrante) HasFlowId() bool {
	if o != nil && o.FlowId.IsSet() {
		return true
	}

	return false
}

// SetFlowId gets a reference to the given NullableString and assigns it to the FlowId field.
func (o *FactureEntrante) SetFlowId(v string) {
	o.FlowId.Set(&v)
}
// SetFlowIdNil sets the value for FlowId to be an explicit nil
func (o *FactureEntrante) SetFlowIdNil() {
	o.FlowId.Set(nil)
}

// UnsetFlowId ensures that no value is present for FlowId, not even an explicit nil
func (o *FactureEntrante) UnsetFlowId() {
	o.FlowId.Unset()
}

// GetFormatSource returns the FormatSource field value
func (o *FactureEntrante) GetFormatSource() FormatFacture {
	if o == nil {
		var ret FormatFacture
		return ret
	}

	return o.FormatSource
}

// GetFormatSourceOk returns a tuple with the FormatSource field value
// and a boolean to check if the value has been set.
func (o *FactureEntrante) GetFormatSourceOk() (*FormatFacture, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormatSource, true
}

// SetFormatSource sets field value
func (o *FactureEntrante) SetFormatSource(v FormatFacture) {
	o.FormatSource = v
}

// GetRefFournisseur returns the RefFournisseur field value
func (o *FactureEntrante) GetRefFournisseur() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefFournisseur
}

// GetRefFournisseurOk returns a tuple with the RefFournisseur field value
// and a boolean to check if the value has been set.
func (o *FactureEntrante) GetRefFournisseurOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefFournisseur, true
}

// SetRefFournisseur sets field value
func (o *FactureEntrante) SetRefFournisseur(v string) {
	o.RefFournisseur = v
}

// GetTypeDocument returns the TypeDocument field value if set, zero value otherwise.
func (o *FactureEntrante) GetTypeDocument() TypeDocument {
	if o == nil || IsNil(o.TypeDocument) {
		var ret TypeDocument
		return ret
	}
	return *o.TypeDocument
}

// GetTypeDocumentOk returns a tuple with the TypeDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FactureEntrante) GetTypeDocumentOk() (*TypeDocument, bool) {
	if o == nil || IsNil(o.TypeDocument) {
		return nil, false
	}
	return o.TypeDocument, true
}

// HasTypeDocument returns a boolean if a field has been set.
func (o *FactureEntrante) HasTypeDocument() bool {
	if o != nil && !IsNil(o.TypeDocument) {
		return true
	}

	return false
}

// SetTypeDocument gets a reference to the given TypeDocument and assigns it to the TypeDocument field.
func (o *FactureEntrante) SetTypeDocument(v TypeDocument) {
	o.TypeDocument = &v
}

// GetFournisseur returns the Fournisseur field value
func (o *FactureEntrante) GetFournisseur() FournisseurEntrant {
	if o == nil {
		var ret FournisseurEntrant
		return ret
	}

	return o.Fournisseur
}

// GetFournisseurOk returns a tuple with the Fournisseur field value
// and a boolean to check if the value has been set.
func (o *FactureEntrante) GetFournisseurOk() (*FournisseurEntrant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fournisseur, true
}

// SetFournisseur sets field value
func (o *FactureEntrante) SetFournisseur(v FournisseurEntrant) {
	o.Fournisseur = v
}

// GetSiteFacturationNom returns the SiteFacturationNom field value
func (o *FactureEntrante) GetSiteFacturationNom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteFacturationNom
}

// GetSiteFacturationNomOk returns a tuple with the SiteFacturationNom field value
// and a boolean to check if the value has been set.
func (o *FactureEntrante) GetSiteFacturationNomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteFacturationNom, true
}

// SetSiteFacturationNom sets field value
func (o *FactureEntrante) SetSiteFacturationNom(v string) {
	o.SiteFacturationNom = v
}

// GetSiteFacturationSiret returns the SiteFacturationSiret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureEntrante) GetSiteFacturationSiret() string {
	if o == nil || IsNil(o.SiteFacturationSiret.Get()) {
		var ret string
		return ret
	}
	return *o.SiteFacturationSiret.Get()
}

// GetSiteFacturationSiretOk returns a tuple with the SiteFacturationSiret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureEntrante) GetSiteFacturationSiretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SiteFacturationSiret.Get(), o.SiteFacturationSiret.IsSet()
}

// HasSiteFacturationSiret returns a boolean if a field has been set.
func (o *FactureEntrante) HasSiteFacturationSiret() bool {
	if o != nil && o.SiteFacturationSiret.IsSet() {
		return true
	}

	return false
}

// SetSiteFacturationSiret gets a reference to the given NullableString and assigns it to the SiteFacturationSiret field.
func (o *FactureEntrante) SetSiteFacturationSiret(v string) {
	o.SiteFacturationSiret.Set(&v)
}
// SetSiteFacturationSiretNil sets the value for SiteFacturationSiret to be an explicit nil
func (o *FactureEntrante) SetSiteFacturationSiretNil() {
	o.SiteFacturationSiret.Set(nil)
}

// UnsetSiteFacturationSiret ensures that no value is present for SiteFacturationSiret, not even an explicit nil
func (o *FactureEntrante) UnsetSiteFacturationSiret() {
	o.SiteFacturationSiret.Unset()
}

// GetDateDePiece returns the DateDePiece field value
func (o *FactureEntrante) GetDateDePiece() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateDePiece
}

// GetDateDePieceOk returns a tuple with the DateDePiece field value
// and a boolean to check if the value has been set.
func (o *FactureEntrante) GetDateDePieceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateDePiece, true
}

// SetDateDePiece sets field value
func (o *FactureEntrante) SetDateDePiece(v string) {
	o.DateDePiece = v
}

// GetDateReglement returns the DateReglement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureEntrante) GetDateReglement() string {
	if o == nil || IsNil(o.DateReglement.Get()) {
		var ret string
		return ret
	}
	return *o.DateReglement.Get()
}

// GetDateReglementOk returns a tuple with the DateReglement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureEntrante) GetDateReglementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateReglement.Get(), o.DateReglement.IsSet()
}

// HasDateReglement returns a boolean if a field has been set.
func (o *FactureEntrante) HasDateReglement() bool {
	if o != nil && o.DateReglement.IsSet() {
		return true
	}

	return false
}

// SetDateReglement gets a reference to the given NullableString and assigns it to the DateReglement field.
func (o *FactureEntrante) SetDateReglement(v string) {
	o.DateReglement.Set(&v)
}
// SetDateReglementNil sets the value for DateReglement to be an explicit nil
func (o *FactureEntrante) SetDateReglementNil() {
	o.DateReglement.Set(nil)
}

// UnsetDateReglement ensures that no value is present for DateReglement, not even an explicit nil
func (o *FactureEntrante) UnsetDateReglement() {
	o.DateReglement.Unset()
}

// GetDevise returns the Devise field value if set, zero value otherwise.
func (o *FactureEntrante) GetDevise() string {
	if o == nil || IsNil(o.Devise) {
		var ret string
		return ret
	}
	return *o.Devise
}

// GetDeviseOk returns a tuple with the Devise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FactureEntrante) GetDeviseOk() (*string, bool) {
	if o == nil || IsNil(o.Devise) {
		return nil, false
	}
	return o.Devise, true
}

// HasDevise returns a boolean if a field has been set.
func (o *FactureEntrante) HasDevise() bool {
	if o != nil && !IsNil(o.Devise) {
		return true
	}

	return false
}

// SetDevise gets a reference to the given string and assigns it to the Devise field.
func (o *FactureEntrante) SetDevise(v string) {
	o.Devise = &v
}

// GetMontantHt returns the MontantHt field value
func (o *FactureEntrante) GetMontantHt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MontantHt
}

// GetMontantHtOk returns a tuple with the MontantHt field value
// and a boolean to check if the value has been set.
func (o *FactureEntrante) GetMontantHtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantHt, true
}

// SetMontantHt sets field value
func (o *FactureEntrante) SetMontantHt(v string) {
	o.MontantHt = v
}

// GetMontantTva returns the MontantTva field value
func (o *FactureEntrante) GetMontantTva() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MontantTva
}

// GetMontantTvaOk returns a tuple with the MontantTva field value
// and a boolean to check if the value has been set.
func (o *FactureEntrante) GetMontantTvaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantTva, true
}

// SetMontantTva sets field value
func (o *FactureEntrante) SetMontantTva(v string) {
	o.MontantTva = v
}

// GetMontantTtc returns the MontantTtc field value
func (o *FactureEntrante) GetMontantTtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MontantTtc
}

// GetMontantTtcOk returns a tuple with the MontantTtc field value
// and a boolean to check if the value has been set.
func (o *FactureEntrante) GetMontantTtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantTtc, true
}

// SetMontantTtc sets field value
func (o *FactureEntrante) SetMontantTtc(v string) {
	o.MontantTtc = v
}

// GetNumeroBonCommande returns the NumeroBonCommande field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureEntrante) GetNumeroBonCommande() string {
	if o == nil || IsNil(o.NumeroBonCommande.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroBonCommande.Get()
}

// GetNumeroBonCommandeOk returns a tuple with the NumeroBonCommande field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureEntrante) GetNumeroBonCommandeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroBonCommande.Get(), o.NumeroBonCommande.IsSet()
}

// HasNumeroBonCommande returns a boolean if a field has been set.
func (o *FactureEntrante) HasNumeroBonCommande() bool {
	if o != nil && o.NumeroBonCommande.IsSet() {
		return true
	}

	return false
}

// SetNumeroBonCommande gets a reference to the given NullableString and assigns it to the NumeroBonCommande field.
func (o *FactureEntrante) SetNumeroBonCommande(v string) {
	o.NumeroBonCommande.Set(&v)
}
// SetNumeroBonCommandeNil sets the value for NumeroBonCommande to be an explicit nil
func (o *FactureEntrante) SetNumeroBonCommandeNil() {
	o.NumeroBonCommande.Set(nil)
}

// UnsetNumeroBonCommande ensures that no value is present for NumeroBonCommande, not even an explicit nil
func (o *FactureEntrante) UnsetNumeroBonCommande() {
	o.NumeroBonCommande.Unset()
}

// GetReferenceContrat returns the ReferenceContrat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureEntrante) GetReferenceContrat() string {
	if o == nil || IsNil(o.ReferenceContrat.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceContrat.Get()
}

// GetReferenceContratOk returns a tuple with the ReferenceContrat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureEntrante) GetReferenceContratOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceContrat.Get(), o.ReferenceContrat.IsSet()
}

// HasReferenceContrat returns a boolean if a field has been set.
func (o *FactureEntrante) HasReferenceContrat() bool {
	if o != nil && o.ReferenceContrat.IsSet() {
		return true
	}

	return false
}

// SetReferenceContrat gets a reference to the given NullableString and assigns it to the ReferenceContrat field.
func (o *FactureEntrante) SetReferenceContrat(v string) {
	o.ReferenceContrat.Set(&v)
}
// SetReferenceContratNil sets the value for ReferenceContrat to be an explicit nil
func (o *FactureEntrante) SetReferenceContratNil() {
	o.ReferenceContrat.Set(nil)
}

// UnsetReferenceContrat ensures that no value is present for ReferenceContrat, not even an explicit nil
func (o *FactureEntrante) UnsetReferenceContrat() {
	o.ReferenceContrat.Unset()
}

// GetObjetFacture returns the ObjetFacture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureEntrante) GetObjetFacture() string {
	if o == nil || IsNil(o.ObjetFacture.Get()) {
		var ret string
		return ret
	}
	return *o.ObjetFacture.Get()
}

// GetObjetFactureOk returns a tuple with the ObjetFacture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureEntrante) GetObjetFactureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjetFacture.Get(), o.ObjetFacture.IsSet()
}

// HasObjetFacture returns a boolean if a field has been set.
func (o *FactureEntrante) HasObjetFacture() bool {
	if o != nil && o.ObjetFacture.IsSet() {
		return true
	}

	return false
}

// SetObjetFacture gets a reference to the given NullableString and assigns it to the ObjetFacture field.
func (o *FactureEntrante) SetObjetFacture(v string) {
	o.ObjetFacture.Set(&v)
}
// SetObjetFactureNil sets the value for ObjetFacture to be an explicit nil
func (o *FactureEntrante) SetObjetFactureNil() {
	o.ObjetFacture.Set(nil)
}

// UnsetObjetFacture ensures that no value is present for ObjetFacture, not even an explicit nil
func (o *FactureEntrante) UnsetObjetFacture() {
	o.ObjetFacture.Unset()
}

// GetDocumentBase64 returns the DocumentBase64 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureEntrante) GetDocumentBase64() string {
	if o == nil || IsNil(o.DocumentBase64.Get()) {
		var ret string
		return ret
	}
	return *o.DocumentBase64.Get()
}

// GetDocumentBase64Ok returns a tuple with the DocumentBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureEntrante) GetDocumentBase64Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentBase64.Get(), o.DocumentBase64.IsSet()
}

// HasDocumentBase64 returns a boolean if a field has been set.
func (o *FactureEntrante) HasDocumentBase64() bool {
	if o != nil && o.DocumentBase64.IsSet() {
		return true
	}

	return false
}

// SetDocumentBase64 gets a reference to the given NullableString and assigns it to the DocumentBase64 field.
func (o *FactureEntrante) SetDocumentBase64(v string) {
	o.DocumentBase64.Set(&v)
}
// SetDocumentBase64Nil sets the value for DocumentBase64 to be an explicit nil
func (o *FactureEntrante) SetDocumentBase64Nil() {
	o.DocumentBase64.Set(nil)
}

// UnsetDocumentBase64 ensures that no value is present for DocumentBase64, not even an explicit nil
func (o *FactureEntrante) UnsetDocumentBase64() {
	o.DocumentBase64.Unset()
}

// GetDocumentContentType returns the DocumentContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureEntrante) GetDocumentContentType() string {
	if o == nil || IsNil(o.DocumentContentType.Get()) {
		var ret string
		return ret
	}
	return *o.DocumentContentType.Get()
}

// GetDocumentContentTypeOk returns a tuple with the DocumentContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureEntrante) GetDocumentContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentContentType.Get(), o.DocumentContentType.IsSet()
}

// HasDocumentContentType returns a boolean if a field has been set.
func (o *FactureEntrante) HasDocumentContentType() bool {
	if o != nil && o.DocumentContentType.IsSet() {
		return true
	}

	return false
}

// SetDocumentContentType gets a reference to the given NullableString and assigns it to the DocumentContentType field.
func (o *FactureEntrante) SetDocumentContentType(v string) {
	o.DocumentContentType.Set(&v)
}
// SetDocumentContentTypeNil sets the value for DocumentContentType to be an explicit nil
func (o *FactureEntrante) SetDocumentContentTypeNil() {
	o.DocumentContentType.Set(nil)
}

// UnsetDocumentContentType ensures that no value is present for DocumentContentType, not even an explicit nil
func (o *FactureEntrante) UnsetDocumentContentType() {
	o.DocumentContentType.Unset()
}

// GetDocumentFilename returns the DocumentFilename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureEntrante) GetDocumentFilename() string {
	if o == nil || IsNil(o.DocumentFilename.Get()) {
		var ret string
		return ret
	}
	return *o.DocumentFilename.Get()
}

// GetDocumentFilenameOk returns a tuple with the DocumentFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureEntrante) GetDocumentFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentFilename.Get(), o.DocumentFilename.IsSet()
}

// HasDocumentFilename returns a boolean if a field has been set.
func (o *FactureEntrante) HasDocumentFilename() bool {
	if o != nil && o.DocumentFilename.IsSet() {
		return true
	}

	return false
}

// SetDocumentFilename gets a reference to the given NullableString and assigns it to the DocumentFilename field.
func (o *FactureEntrante) SetDocumentFilename(v string) {
	o.DocumentFilename.Set(&v)
}
// SetDocumentFilenameNil sets the value for DocumentFilename to be an explicit nil
func (o *FactureEntrante) SetDocumentFilenameNil() {
	o.DocumentFilename.Set(nil)
}

// UnsetDocumentFilename ensures that no value is present for DocumentFilename, not even an explicit nil
func (o *FactureEntrante) UnsetDocumentFilename() {
	o.DocumentFilename.Unset()
}

func (o FactureEntrante) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FactureEntrante) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FlowId.IsSet() {
		toSerialize["flow_id"] = o.FlowId.Get()
	}
	toSerialize["format_source"] = o.FormatSource
	toSerialize["ref_fournisseur"] = o.RefFournisseur
	if !IsNil(o.TypeDocument) {
		toSerialize["type_document"] = o.TypeDocument
	}
	toSerialize["fournisseur"] = o.Fournisseur
	toSerialize["site_facturation_nom"] = o.SiteFacturationNom
	if o.SiteFacturationSiret.IsSet() {
		toSerialize["site_facturation_siret"] = o.SiteFacturationSiret.Get()
	}
	toSerialize["date_de_piece"] = o.DateDePiece
	if o.DateReglement.IsSet() {
		toSerialize["date_reglement"] = o.DateReglement.Get()
	}
	if !IsNil(o.Devise) {
		toSerialize["devise"] = o.Devise
	}
	toSerialize["montant_ht"] = o.MontantHt
	toSerialize["montant_tva"] = o.MontantTva
	toSerialize["montant_ttc"] = o.MontantTtc
	if o.NumeroBonCommande.IsSet() {
		toSerialize["numero_bon_commande"] = o.NumeroBonCommande.Get()
	}
	if o.ReferenceContrat.IsSet() {
		toSerialize["reference_contrat"] = o.ReferenceContrat.Get()
	}
	if o.ObjetFacture.IsSet() {
		toSerialize["objet_facture"] = o.ObjetFacture.Get()
	}
	if o.DocumentBase64.IsSet() {
		toSerialize["document_base64"] = o.DocumentBase64.Get()
	}
	if o.DocumentContentType.IsSet() {
		toSerialize["document_content_type"] = o.DocumentContentType.Get()
	}
	if o.DocumentFilename.IsSet() {
		toSerialize["document_filename"] = o.DocumentFilename.Get()
	}
	return toSerialize, nil
}

func (o *FactureEntrante) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"format_source",
		"ref_fournisseur",
		"fournisseur",
		"site_facturation_nom",
		"date_de_piece",
		"montant_ht",
		"montant_tva",
		"montant_ttc",
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

	varFactureEntrante := _FactureEntrante{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFactureEntrante)

	if err != nil {
		return err
	}

	*o = FactureEntrante(varFactureEntrante)

	return err
}

type NullableFactureEntrante struct {
	value *FactureEntrante
	isSet bool
}

func (v NullableFactureEntrante) Get() *FactureEntrante {
	return v.value
}

func (v *NullableFactureEntrante) Set(val *FactureEntrante) {
	v.value = val
	v.isSet = true
}

func (v NullableFactureEntrante) IsSet() bool {
	return v.isSet
}

func (v *NullableFactureEntrante) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFactureEntrante(val *FactureEntrante) *NullableFactureEntrante {
	return &NullableFactureEntrante{value: val, isSet: true}
}

func (v NullableFactureEntrante) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFactureEntrante) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


