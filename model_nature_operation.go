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

// NatureOperation Nature de l'opération (BT-23) pour Factur-X - Réforme française.  BR-FR-08: Le cadre de facturation doit être l'un des codes suivants. La première lettre indique : B = Biens, S = Services, M = Mixte.  Ref: XP Z12-012, article_conformite_pdf_facturx.md  Exemple d'utilisation:     >>> cadre = CadreDeFacturation(     ...     code_cadre_facturation=CodeCadreFacturation.A1_FACTURE_FOURNISSEUR,     ...     nature_operation=NatureOperation.BIENS     ... )
type NatureOperation string

// List of NatureOperation
const (
	B1 NatureOperation = "B1"
	B12 NatureOperation = "B1"
	S1 NatureOperation = "S1"
	S12 NatureOperation = "S1"
	M1 NatureOperation = "M1"
	M12 NatureOperation = "M1"
	B2 NatureOperation = "B2"
	B22 NatureOperation = "B2"
	S2 NatureOperation = "S2"
	S22 NatureOperation = "S2"
	M2 NatureOperation = "M2"
	M22 NatureOperation = "M2"
	B4 NatureOperation = "B4"
	B42 NatureOperation = "B4"
	S4 NatureOperation = "S4"
	S42 NatureOperation = "S4"
	M4 NatureOperation = "M4"
	M42 NatureOperation = "M4"
	S3 NatureOperation = "S3"
	S32 NatureOperation = "S3"
	S5 NatureOperation = "S5"
	S52 NatureOperation = "S5"
	S6 NatureOperation = "S6"
	S62 NatureOperation = "S6"
	B7 NatureOperation = "B7"
	B72 NatureOperation = "B7"
	S7 NatureOperation = "S7"
	S72 NatureOperation = "S7"
)

// All allowed values of NatureOperation enum
var AllowedNatureOperationEnumValues = []NatureOperation{
	"B1",
	"B1",
	"S1",
	"S1",
	"M1",
	"M1",
	"B2",
	"B2",
	"S2",
	"S2",
	"M2",
	"M2",
	"B4",
	"B4",
	"S4",
	"S4",
	"M4",
	"M4",
	"S3",
	"S3",
	"S5",
	"S5",
	"S6",
	"S6",
	"B7",
	"B7",
	"S7",
	"S7",
}

func (v *NatureOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NatureOperation(value)
	for _, existing := range AllowedNatureOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NatureOperation", value)
}

// NewNatureOperationFromValue returns a pointer to a valid NatureOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNatureOperationFromValue(v string) (*NatureOperation, error) {
	ev := NatureOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NatureOperation: valid values are %v", v, AllowedNatureOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NatureOperation) IsValid() bool {
	for _, existing := range AllowedNatureOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NatureOperation value
func (v NatureOperation) Ptr() *NatureOperation {
	return &v
}

type NullableNatureOperation struct {
	value *NatureOperation
	isSet bool
}

func (v NullableNatureOperation) Get() *NatureOperation {
	return v.value
}

func (v *NullableNatureOperation) Set(val *NatureOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableNatureOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableNatureOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatureOperation(val *NatureOperation) *NullableNatureOperation {
	return &NullableNatureOperation{value: val, isSet: true}
}

func (v NullableNatureOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatureOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

