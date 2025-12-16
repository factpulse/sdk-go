/*
API REST FactPulse

 API REST pour la facturation électronique en France : Factur-X, AFNOR PDP/PA, signatures électroniques.  ## 🎯 Fonctionnalités principales  ### 📄 Génération de factures Factur-X - **Formats** : XML seul ou PDF/A-3 avec XML embarqué - **Profils** : MINIMUM, BASIC, EN16931, EXTENDED - **Normes** : EN 16931 (directive UE 2014/55), ISO 19005-3 (PDF/A-3), CII (UN/CEFACT) - **🆕 Format simplifié** : Génération à partir de SIRET + auto-enrichissement (API Chorus Pro + Recherche Entreprises)  ### ✅ Validation et conformité - **Validation XML** : Schematron (45 à 210+ règles selon profil) - **Validation PDF** : PDF/A-3, métadonnées XMP Factur-X, signatures électroniques - **VeraPDF** : Validation stricte PDF/A (146+ règles ISO 19005-3) - **Traitement asynchrone** : Support Celery pour validations lourdes (VeraPDF)  ### 📡 Intégration AFNOR PDP/PA (XP Z12-013) - **Soumission de flux** : Envoi de factures vers Plateformes de Dématérialisation Partenaires - **Recherche de flux** : Consultation des factures soumises - **Téléchargement** : Récupération des PDF/A-3 avec XML - **Directory Service** : Recherche d'entreprises (SIREN/SIRET) - **Multi-client** : Support de plusieurs configs PDP par utilisateur (stored credentials ou zero-storage)  ### ✍️ Signature électronique PDF - **Standards** : PAdES-B-B, PAdES-B-T (horodatage RFC 3161), PAdES-B-LT (archivage long terme) - **Niveaux eIDAS** : SES (auto-signé), AdES (CA commerciale), QES (PSCO) - **Validation** : Vérification intégrité cryptographique et certificats - **Génération de certificats** : Certificats X.509 auto-signés pour tests  ### 🔄 Traitement asynchrone - **Celery** : Génération, validation et signature asynchrones - **Polling** : Suivi d'état via `/taches/{id_tache}/statut` - **Pas de timeout** : Idéal pour gros fichiers ou validations lourdes  ## 🔒 Authentification  Toutes les requêtes nécessitent un **token JWT** dans le header Authorization : ``` Authorization: Bearer YOUR_JWT_TOKEN ```  ### Comment obtenir un token JWT ?  #### 🔑 Méthode 1 : API `/api/token/` (Recommandée)  **URL :** `https://www.factpulse.fr/api/token/`  Cette méthode est **recommandée** pour l'intégration dans vos applications et workflows CI/CD.  **Prérequis :** Avoir défini un mot de passe sur votre compte  **Pour les utilisateurs inscrits via email/password :** - Vous avez déjà un mot de passe, utilisez-le directement  **Pour les utilisateurs inscrits via OAuth (Google/GitHub) :** - Vous devez d'abord définir un mot de passe sur : https://www.factpulse.fr/accounts/password/set/ - Une fois le mot de passe créé, vous pourrez utiliser l'API  **Exemple de requête :** ```bash curl -X POST https://www.factpulse.fr/api/token/ \\   -H \"Content-Type: application/json\" \\   -d '{     \"username\": \"votre_email@example.com\",     \"password\": \"votre_mot_de_passe\"   }' ```  **Paramètre optionnel `client_uid` :**  Pour sélectionner les credentials d'un client spécifique (PA/PDP, Chorus Pro, certificats de signature), ajoutez `client_uid` :  ```bash curl -X POST https://www.factpulse.fr/api/token/ \\   -H \"Content-Type: application/json\" \\   -d '{     \"username\": \"votre_email@example.com\",     \"password\": \"votre_mot_de_passe\",     \"client_uid\": \"550e8400-e29b-41d4-a716-446655440000\"   }' ```  Le `client_uid` sera inclus dans le JWT et permettra à l'API d'utiliser automatiquement : - Les credentials AFNOR/PDP configurés pour ce client - Les credentials Chorus Pro configurés pour ce client - Les certificats de signature électronique configurés pour ce client  **Réponse :** ```json {   \"access\": \"eyJ0eXAiOiJKV1QiLCJhbGc...\",  // Token d'accès (validité: 30 min)   \"refresh\": \"eyJ0eXAiOiJKV1QiLCJhbGc...\"  // Token de rafraîchissement (validité: 7 jours) } ```  **Avantages :** - ✅ Automatisation complète (CI/CD, scripts) - ✅ Gestion programmatique des tokens - ✅ Support du refresh token pour renouveler automatiquement l'accès - ✅ Intégration facile dans n'importe quel langage/outil  #### 🖥️ Méthode 2 : Génération via Dashboard (Alternative)  **URL :** https://www.factpulse.fr/dashboard/  Cette méthode convient pour des tests rapides ou une utilisation occasionnelle via l'interface graphique.  **Fonctionnement :** - Connectez-vous au dashboard - Utilisez les boutons \"Generate Test Token\" ou \"Generate Production Token\" - Fonctionne pour **tous** les utilisateurs (OAuth et email/password), sans nécessiter de mot de passe  **Types de tokens :** - **Token Test** : Validité 24h, quota 1000 appels/jour (gratuit) - **Token Production** : Validité 7 jours, quota selon votre forfait  **Avantages :** - ✅ Rapide pour tester l'API - ✅ Aucun mot de passe requis - ✅ Interface visuelle simple  **Inconvénients :** - ❌ Nécessite une action manuelle - ❌ Pas de refresh token - ❌ Moins adapté pour l'automatisation  ### 📚 Documentation complète  Pour plus d'informations sur l'authentification et l'utilisation de l'API : https://www.factpulse.fr/documentation-api/     

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package factpulse

import (
	"encoding/json"
	"fmt"
)

// TypeFacture Type de document selon BR-FR-04 (codes UNTDID 1001).  | Code | Nom | Description | |------|-----|-------------| | 380 | FACTURE | Facture commerciale | | 389 | FACTURE_AUTOFACTUREE | Facture auto-facturée | | 393 | FACTURE_AFFACTUREE | Facture affacturée | | 501 | FACTURE_AUTOFACTUREE_AFFACTUREE | Facture auto-facturée affacturée | | 386 | FACTURE_ACOMPTE | Facture d'acompte | | 500 | FACTURE_ACOMPTE_AUTOFACTUREE | Facture d'acompte auto-facturée | | 384 | FACTURE_RECTIFICATIVE | Facture rectificative | | 471 | FACTURE_RECTIFICATIVE_AUTOFACTUREE | Facture rectificative auto-facturée | | 472 | FACTURE_RECTIFICATIVE_AFFACTUREE | Facture rectificative affacturée | | 473 | FACTURE_RECTIFICATIVE_AUTOFACTUREE_AFFACTUREE | Facture rectificative auto-facturée affacturée | | 381 | AVOIR | Avoir | | 261 | AVOIR_AUTOFACTURE | Avoir auto-facturé | | 262 | AVOIR_REMISE_GLOBALE | Avoir pour remise globale | | 396 | AVOIR_AFFACTURE | Avoir affacturé | | 502 | AVOIR_AUTOFACTURE_AFFACTURE | Avoir auto-facturé affacturé | | 503 | AVOIR_ACOMPTE | Avoir de facture d'acompte |
type TypeFacture string

// List of TypeFacture
const (
	FACTURE TypeFacture = "380"
	FACTURE_AUTOFACTUREE TypeFacture = "389"
	FACTURE_AFFACTUREE TypeFacture = "393"
	FACTURE_AUTOFACTUREE_AFFACTUREE TypeFacture = "501"
	FACTURE_ACOMPTE TypeFacture = "386"
	FACTURE_ACOMPTE_AUTOFACTUREE TypeFacture = "500"
	FACTURE_RECTIFICATIVE TypeFacture = "384"
	FACTURE_RECTIFICATIVE_AUTOFACTUREE TypeFacture = "471"
	FACTURE_RECTIFICATIVE_AFFACTUREE TypeFacture = "472"
	FACTURE_RECTIFICATIVE_AUTOFACTUREE_AFFACTUREE TypeFacture = "473"
	AVOIR TypeFacture = "381"
	AVOIR_AUTOFACTURE TypeFacture = "261"
	AVOIR_REMISE_GLOBALE TypeFacture = "262"
	AVOIR_AFFACTURE TypeFacture = "396"
	AVOIR_AUTOFACTURE_AFFACTURE TypeFacture = "502"
	AVOIR_ACOMPTE TypeFacture = "503"
)

// All allowed values of TypeFacture enum
var AllowedTypeFactureEnumValues = []TypeFacture{
	"380",
	"389",
	"393",
	"501",
	"386",
	"500",
	"384",
	"471",
	"472",
	"473",
	"381",
	"261",
	"262",
	"396",
	"502",
	"503",
}

func (v *TypeFacture) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeFacture(value)
	for _, existing := range AllowedTypeFactureEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeFacture", value)
}

// NewTypeFactureFromValue returns a pointer to a valid TypeFacture
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeFactureFromValue(v string) (*TypeFacture, error) {
	ev := TypeFacture(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypeFacture: valid values are %v", v, AllowedTypeFactureEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeFacture) IsValid() bool {
	for _, existing := range AllowedTypeFactureEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TypeFacture value
func (v TypeFacture) Ptr() *TypeFacture {
	return &v
}

type NullableTypeFacture struct {
	value *TypeFacture
	isSet bool
}

func (v NullableTypeFacture) Get() *TypeFacture {
	return v.value
}

func (v *NullableTypeFacture) Set(val *TypeFacture) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeFacture) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeFacture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeFacture(val *TypeFacture) *NullableTypeFacture {
	return &NullableTypeFacture{value: val, isSet: true}
}

func (v NullableTypeFacture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeFacture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

