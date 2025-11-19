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

// checks if the StructureInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StructureInfo{}

// StructureInfo Informations d'une structure.
type StructureInfo struct {
	// ID Chorus Pro de la structure
	IdStructureCpp int32 `json:"id_structure_cpp"`
	// Identifiant (SIRET, SIREN)
	IdentifiantStructure string `json:"identifiant_structure"`
	// Nom de la structure
	DesignationStructure string `json:"designation_structure"`
	// Type d'identifiant
	TypeIdentifiantStructure string `json:"type_identifiant_structure"`
	// Statut (ACTIVE, INACTIVE)
	Statut string `json:"statut"`
}

type _StructureInfo StructureInfo

// NewStructureInfo instantiates a new StructureInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructureInfo(idStructureCpp int32, identifiantStructure string, designationStructure string, typeIdentifiantStructure string, statut string) *StructureInfo {
	this := StructureInfo{}
	this.IdStructureCpp = idStructureCpp
	this.IdentifiantStructure = identifiantStructure
	this.DesignationStructure = designationStructure
	this.TypeIdentifiantStructure = typeIdentifiantStructure
	this.Statut = statut
	return &this
}

// NewStructureInfoWithDefaults instantiates a new StructureInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructureInfoWithDefaults() *StructureInfo {
	this := StructureInfo{}
	return &this
}

// GetIdStructureCpp returns the IdStructureCpp field value
func (o *StructureInfo) GetIdStructureCpp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IdStructureCpp
}

// GetIdStructureCppOk returns a tuple with the IdStructureCpp field value
// and a boolean to check if the value has been set.
func (o *StructureInfo) GetIdStructureCppOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdStructureCpp, true
}

// SetIdStructureCpp sets field value
func (o *StructureInfo) SetIdStructureCpp(v int32) {
	o.IdStructureCpp = v
}

// GetIdentifiantStructure returns the IdentifiantStructure field value
func (o *StructureInfo) GetIdentifiantStructure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentifiantStructure
}

// GetIdentifiantStructureOk returns a tuple with the IdentifiantStructure field value
// and a boolean to check if the value has been set.
func (o *StructureInfo) GetIdentifiantStructureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentifiantStructure, true
}

// SetIdentifiantStructure sets field value
func (o *StructureInfo) SetIdentifiantStructure(v string) {
	o.IdentifiantStructure = v
}

// GetDesignationStructure returns the DesignationStructure field value
func (o *StructureInfo) GetDesignationStructure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DesignationStructure
}

// GetDesignationStructureOk returns a tuple with the DesignationStructure field value
// and a boolean to check if the value has been set.
func (o *StructureInfo) GetDesignationStructureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesignationStructure, true
}

// SetDesignationStructure sets field value
func (o *StructureInfo) SetDesignationStructure(v string) {
	o.DesignationStructure = v
}

// GetTypeIdentifiantStructure returns the TypeIdentifiantStructure field value
func (o *StructureInfo) GetTypeIdentifiantStructure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeIdentifiantStructure
}

// GetTypeIdentifiantStructureOk returns a tuple with the TypeIdentifiantStructure field value
// and a boolean to check if the value has been set.
func (o *StructureInfo) GetTypeIdentifiantStructureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeIdentifiantStructure, true
}

// SetTypeIdentifiantStructure sets field value
func (o *StructureInfo) SetTypeIdentifiantStructure(v string) {
	o.TypeIdentifiantStructure = v
}

// GetStatut returns the Statut field value
func (o *StructureInfo) GetStatut() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Statut
}

// GetStatutOk returns a tuple with the Statut field value
// and a boolean to check if the value has been set.
func (o *StructureInfo) GetStatutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statut, true
}

// SetStatut sets field value
func (o *StructureInfo) SetStatut(v string) {
	o.Statut = v
}

func (o StructureInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StructureInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id_structure_cpp"] = o.IdStructureCpp
	toSerialize["identifiant_structure"] = o.IdentifiantStructure
	toSerialize["designation_structure"] = o.DesignationStructure
	toSerialize["type_identifiant_structure"] = o.TypeIdentifiantStructure
	toSerialize["statut"] = o.Statut
	return toSerialize, nil
}

func (o *StructureInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id_structure_cpp",
		"identifiant_structure",
		"designation_structure",
		"type_identifiant_structure",
		"statut",
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

	varStructureInfo := _StructureInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStructureInfo)

	if err != nil {
		return err
	}

	*o = StructureInfo(varStructureInfo)

	return err
}

type NullableStructureInfo struct {
	value *StructureInfo
	isSet bool
}

func (v NullableStructureInfo) Get() *StructureInfo {
	return v.value
}

func (v *NullableStructureInfo) Set(val *StructureInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStructureInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStructureInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructureInfo(val *StructureInfo) *NullableStructureInfo {
	return &NullableStructureInfo{value: val, isSet: true}
}

func (v NullableStructureInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructureInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


