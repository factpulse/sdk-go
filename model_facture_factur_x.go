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

// checks if the FactureFacturX type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FactureFacturX{}

// FactureFacturX Modèle de données pour une facture destinée à être convertie en Factur-X.
type FactureFacturX struct {
	NumeroFacture string `json:"numeroFacture"`
	DateEcheancePaiement string `json:"dateEcheancePaiement"`
	DateFacture *string `json:"dateFacture,omitempty"`
	ModeDepot ModeDepot `json:"modeDepot"`
	Destinataire Destinataire `json:"destinataire"`
	Fournisseur Fournisseur `json:"fournisseur"`
	CadreDeFacturation CadreDeFacturation `json:"cadreDeFacturation"`
	References References `json:"references"`
	MontantTotal MontantTotal `json:"montantTotal"`
	LignesDePoste []LigneDePoste `json:"lignesDePoste,omitempty"`
	LignesDeTva []LigneDeTVA `json:"lignesDeTva,omitempty"`
	Notes []Note `json:"notes,omitempty"`
	Commentaire NullableString `json:"commentaire,omitempty"`
	IdUtilisateurCourant NullableInt32 `json:"idUtilisateurCourant,omitempty"`
	PiecesJointesComplementaires []PieceJointeComplementaire `json:"piecesJointesComplementaires,omitempty"`
}

type _FactureFacturX FactureFacturX

// NewFactureFacturX instantiates a new FactureFacturX object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFactureFacturX(numeroFacture string, dateEcheancePaiement string, modeDepot ModeDepot, destinataire Destinataire, fournisseur Fournisseur, cadreDeFacturation CadreDeFacturation, references References, montantTotal MontantTotal) *FactureFacturX {
	this := FactureFacturX{}
	this.NumeroFacture = numeroFacture
	this.DateEcheancePaiement = dateEcheancePaiement
	this.ModeDepot = modeDepot
	this.Destinataire = destinataire
	this.Fournisseur = fournisseur
	this.CadreDeFacturation = cadreDeFacturation
	this.References = references
	this.MontantTotal = montantTotal
	return &this
}

// NewFactureFacturXWithDefaults instantiates a new FactureFacturX object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFactureFacturXWithDefaults() *FactureFacturX {
	this := FactureFacturX{}
	return &this
}

// GetNumeroFacture returns the NumeroFacture field value
func (o *FactureFacturX) GetNumeroFacture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NumeroFacture
}

// GetNumeroFactureOk returns a tuple with the NumeroFacture field value
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetNumeroFactureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumeroFacture, true
}

// SetNumeroFacture sets field value
func (o *FactureFacturX) SetNumeroFacture(v string) {
	o.NumeroFacture = v
}

// GetDateEcheancePaiement returns the DateEcheancePaiement field value
func (o *FactureFacturX) GetDateEcheancePaiement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateEcheancePaiement
}

// GetDateEcheancePaiementOk returns a tuple with the DateEcheancePaiement field value
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetDateEcheancePaiementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateEcheancePaiement, true
}

// SetDateEcheancePaiement sets field value
func (o *FactureFacturX) SetDateEcheancePaiement(v string) {
	o.DateEcheancePaiement = v
}

// GetDateFacture returns the DateFacture field value if set, zero value otherwise.
func (o *FactureFacturX) GetDateFacture() string {
	if o == nil || IsNil(o.DateFacture) {
		var ret string
		return ret
	}
	return *o.DateFacture
}

// GetDateFactureOk returns a tuple with the DateFacture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetDateFactureOk() (*string, bool) {
	if o == nil || IsNil(o.DateFacture) {
		return nil, false
	}
	return o.DateFacture, true
}

// HasDateFacture returns a boolean if a field has been set.
func (o *FactureFacturX) HasDateFacture() bool {
	if o != nil && !IsNil(o.DateFacture) {
		return true
	}

	return false
}

// SetDateFacture gets a reference to the given string and assigns it to the DateFacture field.
func (o *FactureFacturX) SetDateFacture(v string) {
	o.DateFacture = &v
}

// GetModeDepot returns the ModeDepot field value
func (o *FactureFacturX) GetModeDepot() ModeDepot {
	if o == nil {
		var ret ModeDepot
		return ret
	}

	return o.ModeDepot
}

// GetModeDepotOk returns a tuple with the ModeDepot field value
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetModeDepotOk() (*ModeDepot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModeDepot, true
}

// SetModeDepot sets field value
func (o *FactureFacturX) SetModeDepot(v ModeDepot) {
	o.ModeDepot = v
}

// GetDestinataire returns the Destinataire field value
func (o *FactureFacturX) GetDestinataire() Destinataire {
	if o == nil {
		var ret Destinataire
		return ret
	}

	return o.Destinataire
}

// GetDestinataireOk returns a tuple with the Destinataire field value
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetDestinataireOk() (*Destinataire, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destinataire, true
}

// SetDestinataire sets field value
func (o *FactureFacturX) SetDestinataire(v Destinataire) {
	o.Destinataire = v
}

// GetFournisseur returns the Fournisseur field value
func (o *FactureFacturX) GetFournisseur() Fournisseur {
	if o == nil {
		var ret Fournisseur
		return ret
	}

	return o.Fournisseur
}

// GetFournisseurOk returns a tuple with the Fournisseur field value
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetFournisseurOk() (*Fournisseur, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fournisseur, true
}

// SetFournisseur sets field value
func (o *FactureFacturX) SetFournisseur(v Fournisseur) {
	o.Fournisseur = v
}

// GetCadreDeFacturation returns the CadreDeFacturation field value
func (o *FactureFacturX) GetCadreDeFacturation() CadreDeFacturation {
	if o == nil {
		var ret CadreDeFacturation
		return ret
	}

	return o.CadreDeFacturation
}

// GetCadreDeFacturationOk returns a tuple with the CadreDeFacturation field value
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetCadreDeFacturationOk() (*CadreDeFacturation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CadreDeFacturation, true
}

// SetCadreDeFacturation sets field value
func (o *FactureFacturX) SetCadreDeFacturation(v CadreDeFacturation) {
	o.CadreDeFacturation = v
}

// GetReferences returns the References field value
func (o *FactureFacturX) GetReferences() References {
	if o == nil {
		var ret References
		return ret
	}

	return o.References
}

// GetReferencesOk returns a tuple with the References field value
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetReferencesOk() (*References, bool) {
	if o == nil {
		return nil, false
	}
	return &o.References, true
}

// SetReferences sets field value
func (o *FactureFacturX) SetReferences(v References) {
	o.References = v
}

// GetMontantTotal returns the MontantTotal field value
func (o *FactureFacturX) GetMontantTotal() MontantTotal {
	if o == nil {
		var ret MontantTotal
		return ret
	}

	return o.MontantTotal
}

// GetMontantTotalOk returns a tuple with the MontantTotal field value
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetMontantTotalOk() (*MontantTotal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantTotal, true
}

// SetMontantTotal sets field value
func (o *FactureFacturX) SetMontantTotal(v MontantTotal) {
	o.MontantTotal = v
}

// GetLignesDePoste returns the LignesDePoste field value if set, zero value otherwise.
func (o *FactureFacturX) GetLignesDePoste() []LigneDePoste {
	if o == nil || IsNil(o.LignesDePoste) {
		var ret []LigneDePoste
		return ret
	}
	return o.LignesDePoste
}

// GetLignesDePosteOk returns a tuple with the LignesDePoste field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetLignesDePosteOk() ([]LigneDePoste, bool) {
	if o == nil || IsNil(o.LignesDePoste) {
		return nil, false
	}
	return o.LignesDePoste, true
}

// HasLignesDePoste returns a boolean if a field has been set.
func (o *FactureFacturX) HasLignesDePoste() bool {
	if o != nil && !IsNil(o.LignesDePoste) {
		return true
	}

	return false
}

// SetLignesDePoste gets a reference to the given []LigneDePoste and assigns it to the LignesDePoste field.
func (o *FactureFacturX) SetLignesDePoste(v []LigneDePoste) {
	o.LignesDePoste = v
}

// GetLignesDeTva returns the LignesDeTva field value if set, zero value otherwise.
func (o *FactureFacturX) GetLignesDeTva() []LigneDeTVA {
	if o == nil || IsNil(o.LignesDeTva) {
		var ret []LigneDeTVA
		return ret
	}
	return o.LignesDeTva
}

// GetLignesDeTvaOk returns a tuple with the LignesDeTva field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetLignesDeTvaOk() ([]LigneDeTVA, bool) {
	if o == nil || IsNil(o.LignesDeTva) {
		return nil, false
	}
	return o.LignesDeTva, true
}

// HasLignesDeTva returns a boolean if a field has been set.
func (o *FactureFacturX) HasLignesDeTva() bool {
	if o != nil && !IsNil(o.LignesDeTva) {
		return true
	}

	return false
}

// SetLignesDeTva gets a reference to the given []LigneDeTVA and assigns it to the LignesDeTva field.
func (o *FactureFacturX) SetLignesDeTva(v []LigneDeTVA) {
	o.LignesDeTva = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *FactureFacturX) GetNotes() []Note {
	if o == nil || IsNil(o.Notes) {
		var ret []Note
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FactureFacturX) GetNotesOk() ([]Note, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *FactureFacturX) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []Note and assigns it to the Notes field.
func (o *FactureFacturX) SetNotes(v []Note) {
	o.Notes = v
}

// GetCommentaire returns the Commentaire field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureFacturX) GetCommentaire() string {
	if o == nil || IsNil(o.Commentaire.Get()) {
		var ret string
		return ret
	}
	return *o.Commentaire.Get()
}

// GetCommentaireOk returns a tuple with the Commentaire field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureFacturX) GetCommentaireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Commentaire.Get(), o.Commentaire.IsSet()
}

// HasCommentaire returns a boolean if a field has been set.
func (o *FactureFacturX) HasCommentaire() bool {
	if o != nil && o.Commentaire.IsSet() {
		return true
	}

	return false
}

// SetCommentaire gets a reference to the given NullableString and assigns it to the Commentaire field.
func (o *FactureFacturX) SetCommentaire(v string) {
	o.Commentaire.Set(&v)
}
// SetCommentaireNil sets the value for Commentaire to be an explicit nil
func (o *FactureFacturX) SetCommentaireNil() {
	o.Commentaire.Set(nil)
}

// UnsetCommentaire ensures that no value is present for Commentaire, not even an explicit nil
func (o *FactureFacturX) UnsetCommentaire() {
	o.Commentaire.Unset()
}

// GetIdUtilisateurCourant returns the IdUtilisateurCourant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureFacturX) GetIdUtilisateurCourant() int32 {
	if o == nil || IsNil(o.IdUtilisateurCourant.Get()) {
		var ret int32
		return ret
	}
	return *o.IdUtilisateurCourant.Get()
}

// GetIdUtilisateurCourantOk returns a tuple with the IdUtilisateurCourant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureFacturX) GetIdUtilisateurCourantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdUtilisateurCourant.Get(), o.IdUtilisateurCourant.IsSet()
}

// HasIdUtilisateurCourant returns a boolean if a field has been set.
func (o *FactureFacturX) HasIdUtilisateurCourant() bool {
	if o != nil && o.IdUtilisateurCourant.IsSet() {
		return true
	}

	return false
}

// SetIdUtilisateurCourant gets a reference to the given NullableInt32 and assigns it to the IdUtilisateurCourant field.
func (o *FactureFacturX) SetIdUtilisateurCourant(v int32) {
	o.IdUtilisateurCourant.Set(&v)
}
// SetIdUtilisateurCourantNil sets the value for IdUtilisateurCourant to be an explicit nil
func (o *FactureFacturX) SetIdUtilisateurCourantNil() {
	o.IdUtilisateurCourant.Set(nil)
}

// UnsetIdUtilisateurCourant ensures that no value is present for IdUtilisateurCourant, not even an explicit nil
func (o *FactureFacturX) UnsetIdUtilisateurCourant() {
	o.IdUtilisateurCourant.Unset()
}

// GetPiecesJointesComplementaires returns the PiecesJointesComplementaires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureFacturX) GetPiecesJointesComplementaires() []PieceJointeComplementaire {
	if o == nil {
		var ret []PieceJointeComplementaire
		return ret
	}
	return o.PiecesJointesComplementaires
}

// GetPiecesJointesComplementairesOk returns a tuple with the PiecesJointesComplementaires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureFacturX) GetPiecesJointesComplementairesOk() ([]PieceJointeComplementaire, bool) {
	if o == nil || IsNil(o.PiecesJointesComplementaires) {
		return nil, false
	}
	return o.PiecesJointesComplementaires, true
}

// HasPiecesJointesComplementaires returns a boolean if a field has been set.
func (o *FactureFacturX) HasPiecesJointesComplementaires() bool {
	if o != nil && !IsNil(o.PiecesJointesComplementaires) {
		return true
	}

	return false
}

// SetPiecesJointesComplementaires gets a reference to the given []PieceJointeComplementaire and assigns it to the PiecesJointesComplementaires field.
func (o *FactureFacturX) SetPiecesJointesComplementaires(v []PieceJointeComplementaire) {
	o.PiecesJointesComplementaires = v
}

func (o FactureFacturX) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FactureFacturX) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["numeroFacture"] = o.NumeroFacture
	toSerialize["dateEcheancePaiement"] = o.DateEcheancePaiement
	if !IsNil(o.DateFacture) {
		toSerialize["dateFacture"] = o.DateFacture
	}
	toSerialize["modeDepot"] = o.ModeDepot
	toSerialize["destinataire"] = o.Destinataire
	toSerialize["fournisseur"] = o.Fournisseur
	toSerialize["cadreDeFacturation"] = o.CadreDeFacturation
	toSerialize["references"] = o.References
	toSerialize["montantTotal"] = o.MontantTotal
	if !IsNil(o.LignesDePoste) {
		toSerialize["lignesDePoste"] = o.LignesDePoste
	}
	if !IsNil(o.LignesDeTva) {
		toSerialize["lignesDeTva"] = o.LignesDeTva
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if o.Commentaire.IsSet() {
		toSerialize["commentaire"] = o.Commentaire.Get()
	}
	if o.IdUtilisateurCourant.IsSet() {
		toSerialize["idUtilisateurCourant"] = o.IdUtilisateurCourant.Get()
	}
	if o.PiecesJointesComplementaires != nil {
		toSerialize["piecesJointesComplementaires"] = o.PiecesJointesComplementaires
	}
	return toSerialize, nil
}

func (o *FactureFacturX) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"numeroFacture",
		"dateEcheancePaiement",
		"modeDepot",
		"destinataire",
		"fournisseur",
		"cadreDeFacturation",
		"references",
		"montantTotal",
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

	varFactureFacturX := _FactureFacturX{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFactureFacturX)

	if err != nil {
		return err
	}

	*o = FactureFacturX(varFactureFacturX)

	return err
}

type NullableFactureFacturX struct {
	value *FactureFacturX
	isSet bool
}

func (v NullableFactureFacturX) Get() *FactureFacturX {
	return v.value
}

func (v *NullableFactureFacturX) Set(val *FactureFacturX) {
	v.value = val
	v.isSet = true
}

func (v NullableFactureFacturX) IsSet() bool {
	return v.isSet
}

func (v *NullableFactureFacturX) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFactureFacturX(val *FactureFacturX) *NullableFactureFacturX {
	return &NullableFactureFacturX{value: val, isSet: true}
}

func (v NullableFactureFacturX) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFactureFacturX) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


