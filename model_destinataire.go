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

// checks if the Destinataire type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Destinataire{}

// Destinataire Informations sur le destinataire de la facture (le client).
type Destinataire struct {
	AdresseElectronique AdresseElectronique `json:"adresseElectronique"`
	CodeServiceExecutant NullableString `json:"codeServiceExecutant,omitempty"`
	Nom NullableString `json:"nom,omitempty"`
	AdressePostale NullableAdressePostale `json:"adressePostale,omitempty"`
}

type _Destinataire Destinataire

// NewDestinataire instantiates a new Destinataire object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinataire(adresseElectronique AdresseElectronique) *Destinataire {
	this := Destinataire{}
	this.AdresseElectronique = adresseElectronique
	return &this
}

// NewDestinataireWithDefaults instantiates a new Destinataire object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinataireWithDefaults() *Destinataire {
	this := Destinataire{}
	return &this
}

// GetAdresseElectronique returns the AdresseElectronique field value
func (o *Destinataire) GetAdresseElectronique() AdresseElectronique {
	if o == nil {
		var ret AdresseElectronique
		return ret
	}

	return o.AdresseElectronique
}

// GetAdresseElectroniqueOk returns a tuple with the AdresseElectronique field value
// and a boolean to check if the value has been set.
func (o *Destinataire) GetAdresseElectroniqueOk() (*AdresseElectronique, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdresseElectronique, true
}

// SetAdresseElectronique sets field value
func (o *Destinataire) SetAdresseElectronique(v AdresseElectronique) {
	o.AdresseElectronique = v
}

// GetCodeServiceExecutant returns the CodeServiceExecutant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Destinataire) GetCodeServiceExecutant() string {
	if o == nil || IsNil(o.CodeServiceExecutant.Get()) {
		var ret string
		return ret
	}
	return *o.CodeServiceExecutant.Get()
}

// GetCodeServiceExecutantOk returns a tuple with the CodeServiceExecutant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Destinataire) GetCodeServiceExecutantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeServiceExecutant.Get(), o.CodeServiceExecutant.IsSet()
}

// HasCodeServiceExecutant returns a boolean if a field has been set.
func (o *Destinataire) HasCodeServiceExecutant() bool {
	if o != nil && o.CodeServiceExecutant.IsSet() {
		return true
	}

	return false
}

// SetCodeServiceExecutant gets a reference to the given NullableString and assigns it to the CodeServiceExecutant field.
func (o *Destinataire) SetCodeServiceExecutant(v string) {
	o.CodeServiceExecutant.Set(&v)
}
// SetCodeServiceExecutantNil sets the value for CodeServiceExecutant to be an explicit nil
func (o *Destinataire) SetCodeServiceExecutantNil() {
	o.CodeServiceExecutant.Set(nil)
}

// UnsetCodeServiceExecutant ensures that no value is present for CodeServiceExecutant, not even an explicit nil
func (o *Destinataire) UnsetCodeServiceExecutant() {
	o.CodeServiceExecutant.Unset()
}

// GetNom returns the Nom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Destinataire) GetNom() string {
	if o == nil || IsNil(o.Nom.Get()) {
		var ret string
		return ret
	}
	return *o.Nom.Get()
}

// GetNomOk returns a tuple with the Nom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Destinataire) GetNomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nom.Get(), o.Nom.IsSet()
}

// HasNom returns a boolean if a field has been set.
func (o *Destinataire) HasNom() bool {
	if o != nil && o.Nom.IsSet() {
		return true
	}

	return false
}

// SetNom gets a reference to the given NullableString and assigns it to the Nom field.
func (o *Destinataire) SetNom(v string) {
	o.Nom.Set(&v)
}
// SetNomNil sets the value for Nom to be an explicit nil
func (o *Destinataire) SetNomNil() {
	o.Nom.Set(nil)
}

// UnsetNom ensures that no value is present for Nom, not even an explicit nil
func (o *Destinataire) UnsetNom() {
	o.Nom.Unset()
}

// GetAdressePostale returns the AdressePostale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Destinataire) GetAdressePostale() AdressePostale {
	if o == nil || IsNil(o.AdressePostale.Get()) {
		var ret AdressePostale
		return ret
	}
	return *o.AdressePostale.Get()
}

// GetAdressePostaleOk returns a tuple with the AdressePostale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Destinataire) GetAdressePostaleOk() (*AdressePostale, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdressePostale.Get(), o.AdressePostale.IsSet()
}

// HasAdressePostale returns a boolean if a field has been set.
func (o *Destinataire) HasAdressePostale() bool {
	if o != nil && o.AdressePostale.IsSet() {
		return true
	}

	return false
}

// SetAdressePostale gets a reference to the given NullableAdressePostale and assigns it to the AdressePostale field.
func (o *Destinataire) SetAdressePostale(v AdressePostale) {
	o.AdressePostale.Set(&v)
}
// SetAdressePostaleNil sets the value for AdressePostale to be an explicit nil
func (o *Destinataire) SetAdressePostaleNil() {
	o.AdressePostale.Set(nil)
}

// UnsetAdressePostale ensures that no value is present for AdressePostale, not even an explicit nil
func (o *Destinataire) UnsetAdressePostale() {
	o.AdressePostale.Unset()
}

func (o Destinataire) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Destinataire) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adresseElectronique"] = o.AdresseElectronique
	if o.CodeServiceExecutant.IsSet() {
		toSerialize["codeServiceExecutant"] = o.CodeServiceExecutant.Get()
	}
	if o.Nom.IsSet() {
		toSerialize["nom"] = o.Nom.Get()
	}
	if o.AdressePostale.IsSet() {
		toSerialize["adressePostale"] = o.AdressePostale.Get()
	}
	return toSerialize, nil
}

func (o *Destinataire) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"adresseElectronique",
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

	varDestinataire := _Destinataire{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDestinataire)

	if err != nil {
		return err
	}

	*o = Destinataire(varDestinataire)

	return err
}

type NullableDestinataire struct {
	value *Destinataire
	isSet bool
}

func (v NullableDestinataire) Get() *Destinataire {
	return v.value
}

func (v *NullableDestinataire) Set(val *Destinataire) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinataire) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinataire) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinataire(val *Destinataire) *NullableDestinataire {
	return &NullableDestinataire{value: val, isSet: true}
}

func (v NullableDestinataire) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinataire) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


