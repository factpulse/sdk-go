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

// checks if the FactureEnrichieInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FactureEnrichieInfo{}

// FactureEnrichieInfo Informations sur la facture enrichie.
type FactureEnrichieInfo struct {
	NumeroFacture string `json:"numero_facture"`
	IdEmetteur NullableInt32 `json:"id_emetteur,omitempty"`
	IdDestinataire NullableInt32 `json:"id_destinataire,omitempty"`
	NomEmetteur string `json:"nom_emetteur"`
	NomDestinataire string `json:"nom_destinataire"`
	MontantHtTotal string `json:"montant_ht_total" validate:"regexp=^(?!^[-+.]*$)[+-]?0*\\\\d*\\\\.?\\\\d*$"`
	MontantTva string `json:"montant_tva" validate:"regexp=^(?!^[-+.]*$)[+-]?0*\\\\d*\\\\.?\\\\d*$"`
	MontantTtcTotal string `json:"montant_ttc_total" validate:"regexp=^(?!^[-+.]*$)[+-]?0*\\\\d*\\\\.?\\\\d*$"`
}

type _FactureEnrichieInfo FactureEnrichieInfo

// NewFactureEnrichieInfo instantiates a new FactureEnrichieInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFactureEnrichieInfo(numeroFacture string, nomEmetteur string, nomDestinataire string, montantHtTotal string, montantTva string, montantTtcTotal string) *FactureEnrichieInfo {
	this := FactureEnrichieInfo{}
	this.NumeroFacture = numeroFacture
	this.NomEmetteur = nomEmetteur
	this.NomDestinataire = nomDestinataire
	this.MontantHtTotal = montantHtTotal
	this.MontantTva = montantTva
	this.MontantTtcTotal = montantTtcTotal
	return &this
}

// NewFactureEnrichieInfoWithDefaults instantiates a new FactureEnrichieInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFactureEnrichieInfoWithDefaults() *FactureEnrichieInfo {
	this := FactureEnrichieInfo{}
	return &this
}

// GetNumeroFacture returns the NumeroFacture field value
func (o *FactureEnrichieInfo) GetNumeroFacture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NumeroFacture
}

// GetNumeroFactureOk returns a tuple with the NumeroFacture field value
// and a boolean to check if the value has been set.
func (o *FactureEnrichieInfo) GetNumeroFactureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumeroFacture, true
}

// SetNumeroFacture sets field value
func (o *FactureEnrichieInfo) SetNumeroFacture(v string) {
	o.NumeroFacture = v
}

// GetIdEmetteur returns the IdEmetteur field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureEnrichieInfo) GetIdEmetteur() int32 {
	if o == nil || IsNil(o.IdEmetteur.Get()) {
		var ret int32
		return ret
	}
	return *o.IdEmetteur.Get()
}

// GetIdEmetteurOk returns a tuple with the IdEmetteur field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureEnrichieInfo) GetIdEmetteurOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdEmetteur.Get(), o.IdEmetteur.IsSet()
}

// HasIdEmetteur returns a boolean if a field has been set.
func (o *FactureEnrichieInfo) HasIdEmetteur() bool {
	if o != nil && o.IdEmetteur.IsSet() {
		return true
	}

	return false
}

// SetIdEmetteur gets a reference to the given NullableInt32 and assigns it to the IdEmetteur field.
func (o *FactureEnrichieInfo) SetIdEmetteur(v int32) {
	o.IdEmetteur.Set(&v)
}
// SetIdEmetteurNil sets the value for IdEmetteur to be an explicit nil
func (o *FactureEnrichieInfo) SetIdEmetteurNil() {
	o.IdEmetteur.Set(nil)
}

// UnsetIdEmetteur ensures that no value is present for IdEmetteur, not even an explicit nil
func (o *FactureEnrichieInfo) UnsetIdEmetteur() {
	o.IdEmetteur.Unset()
}

// GetIdDestinataire returns the IdDestinataire field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FactureEnrichieInfo) GetIdDestinataire() int32 {
	if o == nil || IsNil(o.IdDestinataire.Get()) {
		var ret int32
		return ret
	}
	return *o.IdDestinataire.Get()
}

// GetIdDestinataireOk returns a tuple with the IdDestinataire field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FactureEnrichieInfo) GetIdDestinataireOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdDestinataire.Get(), o.IdDestinataire.IsSet()
}

// HasIdDestinataire returns a boolean if a field has been set.
func (o *FactureEnrichieInfo) HasIdDestinataire() bool {
	if o != nil && o.IdDestinataire.IsSet() {
		return true
	}

	return false
}

// SetIdDestinataire gets a reference to the given NullableInt32 and assigns it to the IdDestinataire field.
func (o *FactureEnrichieInfo) SetIdDestinataire(v int32) {
	o.IdDestinataire.Set(&v)
}
// SetIdDestinataireNil sets the value for IdDestinataire to be an explicit nil
func (o *FactureEnrichieInfo) SetIdDestinataireNil() {
	o.IdDestinataire.Set(nil)
}

// UnsetIdDestinataire ensures that no value is present for IdDestinataire, not even an explicit nil
func (o *FactureEnrichieInfo) UnsetIdDestinataire() {
	o.IdDestinataire.Unset()
}

// GetNomEmetteur returns the NomEmetteur field value
func (o *FactureEnrichieInfo) GetNomEmetteur() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NomEmetteur
}

// GetNomEmetteurOk returns a tuple with the NomEmetteur field value
// and a boolean to check if the value has been set.
func (o *FactureEnrichieInfo) GetNomEmetteurOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NomEmetteur, true
}

// SetNomEmetteur sets field value
func (o *FactureEnrichieInfo) SetNomEmetteur(v string) {
	o.NomEmetteur = v
}

// GetNomDestinataire returns the NomDestinataire field value
func (o *FactureEnrichieInfo) GetNomDestinataire() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NomDestinataire
}

// GetNomDestinataireOk returns a tuple with the NomDestinataire field value
// and a boolean to check if the value has been set.
func (o *FactureEnrichieInfo) GetNomDestinataireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NomDestinataire, true
}

// SetNomDestinataire sets field value
func (o *FactureEnrichieInfo) SetNomDestinataire(v string) {
	o.NomDestinataire = v
}

// GetMontantHtTotal returns the MontantHtTotal field value
func (o *FactureEnrichieInfo) GetMontantHtTotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MontantHtTotal
}

// GetMontantHtTotalOk returns a tuple with the MontantHtTotal field value
// and a boolean to check if the value has been set.
func (o *FactureEnrichieInfo) GetMontantHtTotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantHtTotal, true
}

// SetMontantHtTotal sets field value
func (o *FactureEnrichieInfo) SetMontantHtTotal(v string) {
	o.MontantHtTotal = v
}

// GetMontantTva returns the MontantTva field value
func (o *FactureEnrichieInfo) GetMontantTva() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MontantTva
}

// GetMontantTvaOk returns a tuple with the MontantTva field value
// and a boolean to check if the value has been set.
func (o *FactureEnrichieInfo) GetMontantTvaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantTva, true
}

// SetMontantTva sets field value
func (o *FactureEnrichieInfo) SetMontantTva(v string) {
	o.MontantTva = v
}

// GetMontantTtcTotal returns the MontantTtcTotal field value
func (o *FactureEnrichieInfo) GetMontantTtcTotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MontantTtcTotal
}

// GetMontantTtcTotalOk returns a tuple with the MontantTtcTotal field value
// and a boolean to check if the value has been set.
func (o *FactureEnrichieInfo) GetMontantTtcTotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontantTtcTotal, true
}

// SetMontantTtcTotal sets field value
func (o *FactureEnrichieInfo) SetMontantTtcTotal(v string) {
	o.MontantTtcTotal = v
}

func (o FactureEnrichieInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FactureEnrichieInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["numero_facture"] = o.NumeroFacture
	if o.IdEmetteur.IsSet() {
		toSerialize["id_emetteur"] = o.IdEmetteur.Get()
	}
	if o.IdDestinataire.IsSet() {
		toSerialize["id_destinataire"] = o.IdDestinataire.Get()
	}
	toSerialize["nom_emetteur"] = o.NomEmetteur
	toSerialize["nom_destinataire"] = o.NomDestinataire
	toSerialize["montant_ht_total"] = o.MontantHtTotal
	toSerialize["montant_tva"] = o.MontantTva
	toSerialize["montant_ttc_total"] = o.MontantTtcTotal
	return toSerialize, nil
}

func (o *FactureEnrichieInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"numero_facture",
		"nom_emetteur",
		"nom_destinataire",
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

	varFactureEnrichieInfo := _FactureEnrichieInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFactureEnrichieInfo)

	if err != nil {
		return err
	}

	*o = FactureEnrichieInfo(varFactureEnrichieInfo)

	return err
}

type NullableFactureEnrichieInfo struct {
	value *FactureEnrichieInfo
	isSet bool
}

func (v NullableFactureEnrichieInfo) Get() *FactureEnrichieInfo {
	return v.value
}

func (v *NullableFactureEnrichieInfo) Set(val *FactureEnrichieInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFactureEnrichieInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFactureEnrichieInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFactureEnrichieInfo(val *FactureEnrichieInfo) *NullableFactureEnrichieInfo {
	return &NullableFactureEnrichieInfo{value: val, isSet: true}
}

func (v NullableFactureEnrichieInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFactureEnrichieInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


