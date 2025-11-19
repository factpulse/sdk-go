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

// checks if the ServiceStructure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceStructure{}

// ServiceStructure Service d'une structure.
type ServiceStructure struct {
	// ID du service
	IdService int32 `json:"id_service"`
	// Code du service
	CodeService string `json:"code_service"`
	// Libellé du service
	LibelleService string `json:"libelle_service"`
	// Service actif
	EstActif bool `json:"est_actif"`
}

type _ServiceStructure ServiceStructure

// NewServiceStructure instantiates a new ServiceStructure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStructure(idService int32, codeService string, libelleService string, estActif bool) *ServiceStructure {
	this := ServiceStructure{}
	this.IdService = idService
	this.CodeService = codeService
	this.LibelleService = libelleService
	this.EstActif = estActif
	return &this
}

// NewServiceStructureWithDefaults instantiates a new ServiceStructure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStructureWithDefaults() *ServiceStructure {
	this := ServiceStructure{}
	return &this
}

// GetIdService returns the IdService field value
func (o *ServiceStructure) GetIdService() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IdService
}

// GetIdServiceOk returns a tuple with the IdService field value
// and a boolean to check if the value has been set.
func (o *ServiceStructure) GetIdServiceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdService, true
}

// SetIdService sets field value
func (o *ServiceStructure) SetIdService(v int32) {
	o.IdService = v
}

// GetCodeService returns the CodeService field value
func (o *ServiceStructure) GetCodeService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeService
}

// GetCodeServiceOk returns a tuple with the CodeService field value
// and a boolean to check if the value has been set.
func (o *ServiceStructure) GetCodeServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeService, true
}

// SetCodeService sets field value
func (o *ServiceStructure) SetCodeService(v string) {
	o.CodeService = v
}

// GetLibelleService returns the LibelleService field value
func (o *ServiceStructure) GetLibelleService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LibelleService
}

// GetLibelleServiceOk returns a tuple with the LibelleService field value
// and a boolean to check if the value has been set.
func (o *ServiceStructure) GetLibelleServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LibelleService, true
}

// SetLibelleService sets field value
func (o *ServiceStructure) SetLibelleService(v string) {
	o.LibelleService = v
}

// GetEstActif returns the EstActif field value
func (o *ServiceStructure) GetEstActif() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EstActif
}

// GetEstActifOk returns a tuple with the EstActif field value
// and a boolean to check if the value has been set.
func (o *ServiceStructure) GetEstActifOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstActif, true
}

// SetEstActif sets field value
func (o *ServiceStructure) SetEstActif(v bool) {
	o.EstActif = v
}

func (o ServiceStructure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceStructure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id_service"] = o.IdService
	toSerialize["code_service"] = o.CodeService
	toSerialize["libelle_service"] = o.LibelleService
	toSerialize["est_actif"] = o.EstActif
	return toSerialize, nil
}

func (o *ServiceStructure) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id_service",
		"code_service",
		"libelle_service",
		"est_actif",
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

	varServiceStructure := _ServiceStructure{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceStructure)

	if err != nil {
		return err
	}

	*o = ServiceStructure(varServiceStructure)

	return err
}

type NullableServiceStructure struct {
	value *ServiceStructure
	isSet bool
}

func (v NullableServiceStructure) Get() *ServiceStructure {
	return v.value
}

func (v *NullableServiceStructure) Set(val *ServiceStructure) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStructure) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStructure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStructure(val *ServiceStructure) *NullableServiceStructure {
	return &NullableServiceStructure{value: val, isSet: true}
}

func (v NullableServiceStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStructure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


