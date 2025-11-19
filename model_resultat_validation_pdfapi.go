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

// checks if the ResultatValidationPDFAPI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultatValidationPDFAPI{}

// ResultatValidationPDFAPI Résultat complet de la validation d'un PDF Factur-X.
type ResultatValidationPDFAPI struct {
	// True si le PDF est conforme à tous les critères (XML, PDF/A, XMP)
	EstConforme bool `json:"est_conforme"`
	// True si un XML Factur-X est embarqué dans le PDF
	XmlPresent bool `json:"xml_present"`
	// True si le XML Factur-X est conforme aux règles Schematron
	XmlConforme bool `json:"xml_conforme"`
	ProfilDetecte NullableString `json:"profil_detecte,omitempty"`
	// Liste des erreurs de validation XML
	ErreursXml []string `json:"erreurs_xml,omitempty"`
	// True si le PDF est conforme PDF/A
	PdfaConforme bool `json:"pdfa_conforme"`
	VersionPdfa NullableString `json:"version_pdfa,omitempty"`
	// Méthode utilisée pour la validation PDF/A (metadata ou verapdf)
	MethodeValidationPdfa *string `json:"methode_validation_pdfa,omitempty"`
	ReglesValidees NullableInt32 `json:"regles_validees,omitempty"`
	ReglesEchouees NullableInt32 `json:"regles_echouees,omitempty"`
	// Liste des erreurs de conformité PDF/A
	ErreursPdfa []string `json:"erreurs_pdfa,omitempty"`
	// Liste des avertissements PDF/A
	AvertissementsPdfa []string `json:"avertissements_pdfa,omitempty"`
	// True si des métadonnées XMP sont présentes
	XmpPresent bool `json:"xmp_present"`
	// True si les métadonnées XMP contiennent des informations Factur-X
	XmpConformeFacturx bool `json:"xmp_conforme_facturx"`
	ProfilXmp NullableString `json:"profil_xmp,omitempty"`
	VersionXmp NullableString `json:"version_xmp,omitempty"`
	// Liste des erreurs de métadonnées XMP
	ErreursXmp []string `json:"erreurs_xmp,omitempty"`
	// Métadonnées XMP extraites du PDF
	MetadonneesXmp map[string]interface{} `json:"metadonnees_xmp,omitempty"`
	// True si le PDF contient au moins une signature
	EstSigne bool `json:"est_signe"`
	// Nombre de signatures électroniques trouvées
	NombreSignatures *int32 `json:"nombre_signatures,omitempty"`
	// Liste des signatures trouvées avec leurs informations
	Signatures []InformationSignatureAPI `json:"signatures,omitempty"`
	// Liste des erreurs lors de l'analyse des signatures
	ErreursSignatures []string `json:"erreurs_signatures,omitempty"`
	// Message résumant le résultat de la validation
	MessageResume string `json:"message_resume"`
}

type _ResultatValidationPDFAPI ResultatValidationPDFAPI

// NewResultatValidationPDFAPI instantiates a new ResultatValidationPDFAPI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultatValidationPDFAPI(estConforme bool, xmlPresent bool, xmlConforme bool, pdfaConforme bool, xmpPresent bool, xmpConformeFacturx bool, estSigne bool, messageResume string) *ResultatValidationPDFAPI {
	this := ResultatValidationPDFAPI{}
	this.EstConforme = estConforme
	this.XmlPresent = xmlPresent
	this.XmlConforme = xmlConforme
	this.PdfaConforme = pdfaConforme
	var methodeValidationPdfa string = "metadata"
	this.MethodeValidationPdfa = &methodeValidationPdfa
	this.XmpPresent = xmpPresent
	this.XmpConformeFacturx = xmpConformeFacturx
	this.EstSigne = estSigne
	var nombreSignatures int32 = 0
	this.NombreSignatures = &nombreSignatures
	this.MessageResume = messageResume
	return &this
}

// NewResultatValidationPDFAPIWithDefaults instantiates a new ResultatValidationPDFAPI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultatValidationPDFAPIWithDefaults() *ResultatValidationPDFAPI {
	this := ResultatValidationPDFAPI{}
	var methodeValidationPdfa string = "metadata"
	this.MethodeValidationPdfa = &methodeValidationPdfa
	var nombreSignatures int32 = 0
	this.NombreSignatures = &nombreSignatures
	return &this
}

// GetEstConforme returns the EstConforme field value
func (o *ResultatValidationPDFAPI) GetEstConforme() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EstConforme
}

// GetEstConformeOk returns a tuple with the EstConforme field value
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetEstConformeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstConforme, true
}

// SetEstConforme sets field value
func (o *ResultatValidationPDFAPI) SetEstConforme(v bool) {
	o.EstConforme = v
}

// GetXmlPresent returns the XmlPresent field value
func (o *ResultatValidationPDFAPI) GetXmlPresent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.XmlPresent
}

// GetXmlPresentOk returns a tuple with the XmlPresent field value
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetXmlPresentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XmlPresent, true
}

// SetXmlPresent sets field value
func (o *ResultatValidationPDFAPI) SetXmlPresent(v bool) {
	o.XmlPresent = v
}

// GetXmlConforme returns the XmlConforme field value
func (o *ResultatValidationPDFAPI) GetXmlConforme() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.XmlConforme
}

// GetXmlConformeOk returns a tuple with the XmlConforme field value
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetXmlConformeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XmlConforme, true
}

// SetXmlConforme sets field value
func (o *ResultatValidationPDFAPI) SetXmlConforme(v bool) {
	o.XmlConforme = v
}

// GetProfilDetecte returns the ProfilDetecte field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultatValidationPDFAPI) GetProfilDetecte() string {
	if o == nil || IsNil(o.ProfilDetecte.Get()) {
		var ret string
		return ret
	}
	return *o.ProfilDetecte.Get()
}

// GetProfilDetecteOk returns a tuple with the ProfilDetecte field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultatValidationPDFAPI) GetProfilDetecteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilDetecte.Get(), o.ProfilDetecte.IsSet()
}

// HasProfilDetecte returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasProfilDetecte() bool {
	if o != nil && o.ProfilDetecte.IsSet() {
		return true
	}

	return false
}

// SetProfilDetecte gets a reference to the given NullableString and assigns it to the ProfilDetecte field.
func (o *ResultatValidationPDFAPI) SetProfilDetecte(v string) {
	o.ProfilDetecte.Set(&v)
}
// SetProfilDetecteNil sets the value for ProfilDetecte to be an explicit nil
func (o *ResultatValidationPDFAPI) SetProfilDetecteNil() {
	o.ProfilDetecte.Set(nil)
}

// UnsetProfilDetecte ensures that no value is present for ProfilDetecte, not even an explicit nil
func (o *ResultatValidationPDFAPI) UnsetProfilDetecte() {
	o.ProfilDetecte.Unset()
}

// GetErreursXml returns the ErreursXml field value if set, zero value otherwise.
func (o *ResultatValidationPDFAPI) GetErreursXml() []string {
	if o == nil || IsNil(o.ErreursXml) {
		var ret []string
		return ret
	}
	return o.ErreursXml
}

// GetErreursXmlOk returns a tuple with the ErreursXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetErreursXmlOk() ([]string, bool) {
	if o == nil || IsNil(o.ErreursXml) {
		return nil, false
	}
	return o.ErreursXml, true
}

// HasErreursXml returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasErreursXml() bool {
	if o != nil && !IsNil(o.ErreursXml) {
		return true
	}

	return false
}

// SetErreursXml gets a reference to the given []string and assigns it to the ErreursXml field.
func (o *ResultatValidationPDFAPI) SetErreursXml(v []string) {
	o.ErreursXml = v
}

// GetPdfaConforme returns the PdfaConforme field value
func (o *ResultatValidationPDFAPI) GetPdfaConforme() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PdfaConforme
}

// GetPdfaConformeOk returns a tuple with the PdfaConforme field value
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetPdfaConformeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PdfaConforme, true
}

// SetPdfaConforme sets field value
func (o *ResultatValidationPDFAPI) SetPdfaConforme(v bool) {
	o.PdfaConforme = v
}

// GetVersionPdfa returns the VersionPdfa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultatValidationPDFAPI) GetVersionPdfa() string {
	if o == nil || IsNil(o.VersionPdfa.Get()) {
		var ret string
		return ret
	}
	return *o.VersionPdfa.Get()
}

// GetVersionPdfaOk returns a tuple with the VersionPdfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultatValidationPDFAPI) GetVersionPdfaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VersionPdfa.Get(), o.VersionPdfa.IsSet()
}

// HasVersionPdfa returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasVersionPdfa() bool {
	if o != nil && o.VersionPdfa.IsSet() {
		return true
	}

	return false
}

// SetVersionPdfa gets a reference to the given NullableString and assigns it to the VersionPdfa field.
func (o *ResultatValidationPDFAPI) SetVersionPdfa(v string) {
	o.VersionPdfa.Set(&v)
}
// SetVersionPdfaNil sets the value for VersionPdfa to be an explicit nil
func (o *ResultatValidationPDFAPI) SetVersionPdfaNil() {
	o.VersionPdfa.Set(nil)
}

// UnsetVersionPdfa ensures that no value is present for VersionPdfa, not even an explicit nil
func (o *ResultatValidationPDFAPI) UnsetVersionPdfa() {
	o.VersionPdfa.Unset()
}

// GetMethodeValidationPdfa returns the MethodeValidationPdfa field value if set, zero value otherwise.
func (o *ResultatValidationPDFAPI) GetMethodeValidationPdfa() string {
	if o == nil || IsNil(o.MethodeValidationPdfa) {
		var ret string
		return ret
	}
	return *o.MethodeValidationPdfa
}

// GetMethodeValidationPdfaOk returns a tuple with the MethodeValidationPdfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetMethodeValidationPdfaOk() (*string, bool) {
	if o == nil || IsNil(o.MethodeValidationPdfa) {
		return nil, false
	}
	return o.MethodeValidationPdfa, true
}

// HasMethodeValidationPdfa returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasMethodeValidationPdfa() bool {
	if o != nil && !IsNil(o.MethodeValidationPdfa) {
		return true
	}

	return false
}

// SetMethodeValidationPdfa gets a reference to the given string and assigns it to the MethodeValidationPdfa field.
func (o *ResultatValidationPDFAPI) SetMethodeValidationPdfa(v string) {
	o.MethodeValidationPdfa = &v
}

// GetReglesValidees returns the ReglesValidees field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultatValidationPDFAPI) GetReglesValidees() int32 {
	if o == nil || IsNil(o.ReglesValidees.Get()) {
		var ret int32
		return ret
	}
	return *o.ReglesValidees.Get()
}

// GetReglesValideesOk returns a tuple with the ReglesValidees field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultatValidationPDFAPI) GetReglesValideesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReglesValidees.Get(), o.ReglesValidees.IsSet()
}

// HasReglesValidees returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasReglesValidees() bool {
	if o != nil && o.ReglesValidees.IsSet() {
		return true
	}

	return false
}

// SetReglesValidees gets a reference to the given NullableInt32 and assigns it to the ReglesValidees field.
func (o *ResultatValidationPDFAPI) SetReglesValidees(v int32) {
	o.ReglesValidees.Set(&v)
}
// SetReglesValideesNil sets the value for ReglesValidees to be an explicit nil
func (o *ResultatValidationPDFAPI) SetReglesValideesNil() {
	o.ReglesValidees.Set(nil)
}

// UnsetReglesValidees ensures that no value is present for ReglesValidees, not even an explicit nil
func (o *ResultatValidationPDFAPI) UnsetReglesValidees() {
	o.ReglesValidees.Unset()
}

// GetReglesEchouees returns the ReglesEchouees field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultatValidationPDFAPI) GetReglesEchouees() int32 {
	if o == nil || IsNil(o.ReglesEchouees.Get()) {
		var ret int32
		return ret
	}
	return *o.ReglesEchouees.Get()
}

// GetReglesEchoueesOk returns a tuple with the ReglesEchouees field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultatValidationPDFAPI) GetReglesEchoueesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReglesEchouees.Get(), o.ReglesEchouees.IsSet()
}

// HasReglesEchouees returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasReglesEchouees() bool {
	if o != nil && o.ReglesEchouees.IsSet() {
		return true
	}

	return false
}

// SetReglesEchouees gets a reference to the given NullableInt32 and assigns it to the ReglesEchouees field.
func (o *ResultatValidationPDFAPI) SetReglesEchouees(v int32) {
	o.ReglesEchouees.Set(&v)
}
// SetReglesEchoueesNil sets the value for ReglesEchouees to be an explicit nil
func (o *ResultatValidationPDFAPI) SetReglesEchoueesNil() {
	o.ReglesEchouees.Set(nil)
}

// UnsetReglesEchouees ensures that no value is present for ReglesEchouees, not even an explicit nil
func (o *ResultatValidationPDFAPI) UnsetReglesEchouees() {
	o.ReglesEchouees.Unset()
}

// GetErreursPdfa returns the ErreursPdfa field value if set, zero value otherwise.
func (o *ResultatValidationPDFAPI) GetErreursPdfa() []string {
	if o == nil || IsNil(o.ErreursPdfa) {
		var ret []string
		return ret
	}
	return o.ErreursPdfa
}

// GetErreursPdfaOk returns a tuple with the ErreursPdfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetErreursPdfaOk() ([]string, bool) {
	if o == nil || IsNil(o.ErreursPdfa) {
		return nil, false
	}
	return o.ErreursPdfa, true
}

// HasErreursPdfa returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasErreursPdfa() bool {
	if o != nil && !IsNil(o.ErreursPdfa) {
		return true
	}

	return false
}

// SetErreursPdfa gets a reference to the given []string and assigns it to the ErreursPdfa field.
func (o *ResultatValidationPDFAPI) SetErreursPdfa(v []string) {
	o.ErreursPdfa = v
}

// GetAvertissementsPdfa returns the AvertissementsPdfa field value if set, zero value otherwise.
func (o *ResultatValidationPDFAPI) GetAvertissementsPdfa() []string {
	if o == nil || IsNil(o.AvertissementsPdfa) {
		var ret []string
		return ret
	}
	return o.AvertissementsPdfa
}

// GetAvertissementsPdfaOk returns a tuple with the AvertissementsPdfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetAvertissementsPdfaOk() ([]string, bool) {
	if o == nil || IsNil(o.AvertissementsPdfa) {
		return nil, false
	}
	return o.AvertissementsPdfa, true
}

// HasAvertissementsPdfa returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasAvertissementsPdfa() bool {
	if o != nil && !IsNil(o.AvertissementsPdfa) {
		return true
	}

	return false
}

// SetAvertissementsPdfa gets a reference to the given []string and assigns it to the AvertissementsPdfa field.
func (o *ResultatValidationPDFAPI) SetAvertissementsPdfa(v []string) {
	o.AvertissementsPdfa = v
}

// GetXmpPresent returns the XmpPresent field value
func (o *ResultatValidationPDFAPI) GetXmpPresent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.XmpPresent
}

// GetXmpPresentOk returns a tuple with the XmpPresent field value
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetXmpPresentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XmpPresent, true
}

// SetXmpPresent sets field value
func (o *ResultatValidationPDFAPI) SetXmpPresent(v bool) {
	o.XmpPresent = v
}

// GetXmpConformeFacturx returns the XmpConformeFacturx field value
func (o *ResultatValidationPDFAPI) GetXmpConformeFacturx() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.XmpConformeFacturx
}

// GetXmpConformeFacturxOk returns a tuple with the XmpConformeFacturx field value
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetXmpConformeFacturxOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XmpConformeFacturx, true
}

// SetXmpConformeFacturx sets field value
func (o *ResultatValidationPDFAPI) SetXmpConformeFacturx(v bool) {
	o.XmpConformeFacturx = v
}

// GetProfilXmp returns the ProfilXmp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultatValidationPDFAPI) GetProfilXmp() string {
	if o == nil || IsNil(o.ProfilXmp.Get()) {
		var ret string
		return ret
	}
	return *o.ProfilXmp.Get()
}

// GetProfilXmpOk returns a tuple with the ProfilXmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultatValidationPDFAPI) GetProfilXmpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilXmp.Get(), o.ProfilXmp.IsSet()
}

// HasProfilXmp returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasProfilXmp() bool {
	if o != nil && o.ProfilXmp.IsSet() {
		return true
	}

	return false
}

// SetProfilXmp gets a reference to the given NullableString and assigns it to the ProfilXmp field.
func (o *ResultatValidationPDFAPI) SetProfilXmp(v string) {
	o.ProfilXmp.Set(&v)
}
// SetProfilXmpNil sets the value for ProfilXmp to be an explicit nil
func (o *ResultatValidationPDFAPI) SetProfilXmpNil() {
	o.ProfilXmp.Set(nil)
}

// UnsetProfilXmp ensures that no value is present for ProfilXmp, not even an explicit nil
func (o *ResultatValidationPDFAPI) UnsetProfilXmp() {
	o.ProfilXmp.Unset()
}

// GetVersionXmp returns the VersionXmp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultatValidationPDFAPI) GetVersionXmp() string {
	if o == nil || IsNil(o.VersionXmp.Get()) {
		var ret string
		return ret
	}
	return *o.VersionXmp.Get()
}

// GetVersionXmpOk returns a tuple with the VersionXmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultatValidationPDFAPI) GetVersionXmpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VersionXmp.Get(), o.VersionXmp.IsSet()
}

// HasVersionXmp returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasVersionXmp() bool {
	if o != nil && o.VersionXmp.IsSet() {
		return true
	}

	return false
}

// SetVersionXmp gets a reference to the given NullableString and assigns it to the VersionXmp field.
func (o *ResultatValidationPDFAPI) SetVersionXmp(v string) {
	o.VersionXmp.Set(&v)
}
// SetVersionXmpNil sets the value for VersionXmp to be an explicit nil
func (o *ResultatValidationPDFAPI) SetVersionXmpNil() {
	o.VersionXmp.Set(nil)
}

// UnsetVersionXmp ensures that no value is present for VersionXmp, not even an explicit nil
func (o *ResultatValidationPDFAPI) UnsetVersionXmp() {
	o.VersionXmp.Unset()
}

// GetErreursXmp returns the ErreursXmp field value if set, zero value otherwise.
func (o *ResultatValidationPDFAPI) GetErreursXmp() []string {
	if o == nil || IsNil(o.ErreursXmp) {
		var ret []string
		return ret
	}
	return o.ErreursXmp
}

// GetErreursXmpOk returns a tuple with the ErreursXmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetErreursXmpOk() ([]string, bool) {
	if o == nil || IsNil(o.ErreursXmp) {
		return nil, false
	}
	return o.ErreursXmp, true
}

// HasErreursXmp returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasErreursXmp() bool {
	if o != nil && !IsNil(o.ErreursXmp) {
		return true
	}

	return false
}

// SetErreursXmp gets a reference to the given []string and assigns it to the ErreursXmp field.
func (o *ResultatValidationPDFAPI) SetErreursXmp(v []string) {
	o.ErreursXmp = v
}

// GetMetadonneesXmp returns the MetadonneesXmp field value if set, zero value otherwise.
func (o *ResultatValidationPDFAPI) GetMetadonneesXmp() map[string]interface{} {
	if o == nil || IsNil(o.MetadonneesXmp) {
		var ret map[string]interface{}
		return ret
	}
	return o.MetadonneesXmp
}

// GetMetadonneesXmpOk returns a tuple with the MetadonneesXmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetMetadonneesXmpOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetadonneesXmp) {
		return map[string]interface{}{}, false
	}
	return o.MetadonneesXmp, true
}

// HasMetadonneesXmp returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasMetadonneesXmp() bool {
	if o != nil && !IsNil(o.MetadonneesXmp) {
		return true
	}

	return false
}

// SetMetadonneesXmp gets a reference to the given map[string]interface{} and assigns it to the MetadonneesXmp field.
func (o *ResultatValidationPDFAPI) SetMetadonneesXmp(v map[string]interface{}) {
	o.MetadonneesXmp = v
}

// GetEstSigne returns the EstSigne field value
func (o *ResultatValidationPDFAPI) GetEstSigne() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EstSigne
}

// GetEstSigneOk returns a tuple with the EstSigne field value
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetEstSigneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstSigne, true
}

// SetEstSigne sets field value
func (o *ResultatValidationPDFAPI) SetEstSigne(v bool) {
	o.EstSigne = v
}

// GetNombreSignatures returns the NombreSignatures field value if set, zero value otherwise.
func (o *ResultatValidationPDFAPI) GetNombreSignatures() int32 {
	if o == nil || IsNil(o.NombreSignatures) {
		var ret int32
		return ret
	}
	return *o.NombreSignatures
}

// GetNombreSignaturesOk returns a tuple with the NombreSignatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetNombreSignaturesOk() (*int32, bool) {
	if o == nil || IsNil(o.NombreSignatures) {
		return nil, false
	}
	return o.NombreSignatures, true
}

// HasNombreSignatures returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasNombreSignatures() bool {
	if o != nil && !IsNil(o.NombreSignatures) {
		return true
	}

	return false
}

// SetNombreSignatures gets a reference to the given int32 and assigns it to the NombreSignatures field.
func (o *ResultatValidationPDFAPI) SetNombreSignatures(v int32) {
	o.NombreSignatures = &v
}

// GetSignatures returns the Signatures field value if set, zero value otherwise.
func (o *ResultatValidationPDFAPI) GetSignatures() []InformationSignatureAPI {
	if o == nil || IsNil(o.Signatures) {
		var ret []InformationSignatureAPI
		return ret
	}
	return o.Signatures
}

// GetSignaturesOk returns a tuple with the Signatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetSignaturesOk() ([]InformationSignatureAPI, bool) {
	if o == nil || IsNil(o.Signatures) {
		return nil, false
	}
	return o.Signatures, true
}

// HasSignatures returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasSignatures() bool {
	if o != nil && !IsNil(o.Signatures) {
		return true
	}

	return false
}

// SetSignatures gets a reference to the given []InformationSignatureAPI and assigns it to the Signatures field.
func (o *ResultatValidationPDFAPI) SetSignatures(v []InformationSignatureAPI) {
	o.Signatures = v
}

// GetErreursSignatures returns the ErreursSignatures field value if set, zero value otherwise.
func (o *ResultatValidationPDFAPI) GetErreursSignatures() []string {
	if o == nil || IsNil(o.ErreursSignatures) {
		var ret []string
		return ret
	}
	return o.ErreursSignatures
}

// GetErreursSignaturesOk returns a tuple with the ErreursSignatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetErreursSignaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.ErreursSignatures) {
		return nil, false
	}
	return o.ErreursSignatures, true
}

// HasErreursSignatures returns a boolean if a field has been set.
func (o *ResultatValidationPDFAPI) HasErreursSignatures() bool {
	if o != nil && !IsNil(o.ErreursSignatures) {
		return true
	}

	return false
}

// SetErreursSignatures gets a reference to the given []string and assigns it to the ErreursSignatures field.
func (o *ResultatValidationPDFAPI) SetErreursSignatures(v []string) {
	o.ErreursSignatures = v
}

// GetMessageResume returns the MessageResume field value
func (o *ResultatValidationPDFAPI) GetMessageResume() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageResume
}

// GetMessageResumeOk returns a tuple with the MessageResume field value
// and a boolean to check if the value has been set.
func (o *ResultatValidationPDFAPI) GetMessageResumeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageResume, true
}

// SetMessageResume sets field value
func (o *ResultatValidationPDFAPI) SetMessageResume(v string) {
	o.MessageResume = v
}

func (o ResultatValidationPDFAPI) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultatValidationPDFAPI) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["est_conforme"] = o.EstConforme
	toSerialize["xml_present"] = o.XmlPresent
	toSerialize["xml_conforme"] = o.XmlConforme
	if o.ProfilDetecte.IsSet() {
		toSerialize["profil_detecte"] = o.ProfilDetecte.Get()
	}
	if !IsNil(o.ErreursXml) {
		toSerialize["erreurs_xml"] = o.ErreursXml
	}
	toSerialize["pdfa_conforme"] = o.PdfaConforme
	if o.VersionPdfa.IsSet() {
		toSerialize["version_pdfa"] = o.VersionPdfa.Get()
	}
	if !IsNil(o.MethodeValidationPdfa) {
		toSerialize["methode_validation_pdfa"] = o.MethodeValidationPdfa
	}
	if o.ReglesValidees.IsSet() {
		toSerialize["regles_validees"] = o.ReglesValidees.Get()
	}
	if o.ReglesEchouees.IsSet() {
		toSerialize["regles_echouees"] = o.ReglesEchouees.Get()
	}
	if !IsNil(o.ErreursPdfa) {
		toSerialize["erreurs_pdfa"] = o.ErreursPdfa
	}
	if !IsNil(o.AvertissementsPdfa) {
		toSerialize["avertissements_pdfa"] = o.AvertissementsPdfa
	}
	toSerialize["xmp_present"] = o.XmpPresent
	toSerialize["xmp_conforme_facturx"] = o.XmpConformeFacturx
	if o.ProfilXmp.IsSet() {
		toSerialize["profil_xmp"] = o.ProfilXmp.Get()
	}
	if o.VersionXmp.IsSet() {
		toSerialize["version_xmp"] = o.VersionXmp.Get()
	}
	if !IsNil(o.ErreursXmp) {
		toSerialize["erreurs_xmp"] = o.ErreursXmp
	}
	if !IsNil(o.MetadonneesXmp) {
		toSerialize["metadonnees_xmp"] = o.MetadonneesXmp
	}
	toSerialize["est_signe"] = o.EstSigne
	if !IsNil(o.NombreSignatures) {
		toSerialize["nombre_signatures"] = o.NombreSignatures
	}
	if !IsNil(o.Signatures) {
		toSerialize["signatures"] = o.Signatures
	}
	if !IsNil(o.ErreursSignatures) {
		toSerialize["erreurs_signatures"] = o.ErreursSignatures
	}
	toSerialize["message_resume"] = o.MessageResume
	return toSerialize, nil
}

func (o *ResultatValidationPDFAPI) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"est_conforme",
		"xml_present",
		"xml_conforme",
		"pdfa_conforme",
		"xmp_present",
		"xmp_conforme_facturx",
		"est_signe",
		"message_resume",
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

	varResultatValidationPDFAPI := _ResultatValidationPDFAPI{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultatValidationPDFAPI)

	if err != nil {
		return err
	}

	*o = ResultatValidationPDFAPI(varResultatValidationPDFAPI)

	return err
}

type NullableResultatValidationPDFAPI struct {
	value *ResultatValidationPDFAPI
	isSet bool
}

func (v NullableResultatValidationPDFAPI) Get() *ResultatValidationPDFAPI {
	return v.value
}

func (v *NullableResultatValidationPDFAPI) Set(val *ResultatValidationPDFAPI) {
	v.value = val
	v.isSet = true
}

func (v NullableResultatValidationPDFAPI) IsSet() bool {
	return v.isSet
}

func (v *NullableResultatValidationPDFAPI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultatValidationPDFAPI(val *ResultatValidationPDFAPI) *NullableResultatValidationPDFAPI {
	return &NullableResultatValidationPDFAPI{value: val, isSet: true}
}

func (v NullableResultatValidationPDFAPI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultatValidationPDFAPI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


