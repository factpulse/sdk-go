/*
API REST FactPulse

 API REST pour la facturation électronique en France : Factur-X, AFNOR PDP/PA, signatures électroniques.  ## 🎯 Fonctionnalités principales  ### 📄 Génération de factures Factur-X - **Formats** : XML seul ou PDF/A-3 avec XML embarqué - **Profils** : MINIMUM, BASIC, EN16931, EXTENDED - **Normes** : EN 16931 (directive UE 2014/55), ISO 19005-3 (PDF/A-3), CII (UN/CEFACT) - **🆕 Format simplifié** : Génération à partir de SIRET + auto-enrichissement (API Chorus Pro + Recherche Entreprises)  ### ✅ Validation et conformité - **Validation XML** : Schematron (45 à 210+ règles selon profil) - **Validation PDF** : PDF/A-3, métadonnées XMP Factur-X, signatures électroniques - **VeraPDF** : Validation stricte PDF/A (146+ règles ISO 19005-3) - **Traitement asynchrone** : Support Celery pour validations lourdes (VeraPDF)  ### 📡 Intégration AFNOR PDP/PA (XP Z12-013) - **Soumission de flux** : Envoi de factures vers Plateformes de Dématérialisation Partenaires - **Recherche de flux** : Consultation des factures soumises - **Téléchargement** : Récupération des PDF/A-3 avec XML - **Directory Service** : Recherche d'entreprises (SIREN/SIRET) - **Multi-client** : Support de plusieurs configs PDP par utilisateur (stored credentials ou zero-storage)  ### ✍️ Signature électronique PDF - **Standards** : PAdES-B-B, PAdES-B-T (horodatage RFC 3161), PAdES-B-LT (archivage long terme) - **Niveaux eIDAS** : SES (auto-signé), AdES (CA commerciale), QES (PSCO) - **Validation** : Vérification intégrité cryptographique et certificats - **Génération de certificats** : Certificats X.509 auto-signés pour tests  ### 🔄 Traitement asynchrone - **Celery** : Génération, validation et signature asynchrones - **Polling** : Suivi d'état via `/taches/{id_tache}/statut` - **Pas de timeout** : Idéal pour gros fichiers ou validations lourdes  ## 🔒 Authentification  Toutes les requêtes nécessitent un **token JWT** dans le header Authorization : ``` Authorization: Bearer YOUR_JWT_TOKEN ```  ### Comment obtenir un token JWT ?  #### 🔑 Méthode 1 : API `/api/token/` (Recommandée)  **URL :** `https://www.factpulse.fr/api/token/`  Cette méthode est **recommandée** pour l'intégration dans vos applications et workflows CI/CD.  **Prérequis :** Avoir défini un mot de passe sur votre compte  **Pour les utilisateurs inscrits via email/password :** - Vous avez déjà un mot de passe, utilisez-le directement  **Pour les utilisateurs inscrits via OAuth (Google/GitHub) :** - Vous devez d'abord définir un mot de passe sur : https://www.factpulse.fr/accounts/password/set/ - Une fois le mot de passe créé, vous pourrez utiliser l'API  **Exemple de requête :** ```bash curl -X POST https://www.factpulse.fr/api/token/ \\   -H \"Content-Type: application/json\" \\   -d '{     \"username\": \"votre_email@example.com\",     \"password\": \"votre_mot_de_passe\"   }' ```  **Paramètre optionnel `client_uid` :**  Pour sélectionner les credentials d'un client spécifique (PA/PDP, Chorus Pro, certificats de signature), ajoutez `client_uid` :  ```bash curl -X POST https://www.factpulse.fr/api/token/ \\   -H \"Content-Type: application/json\" \\   -d '{     \"username\": \"votre_email@example.com\",     \"password\": \"votre_mot_de_passe\",     \"client_uid\": \"550e8400-e29b-41d4-a716-446655440000\"   }' ```  Le `client_uid` sera inclus dans le JWT et permettra à l'API d'utiliser automatiquement : - Les credentials AFNOR/PDP configurés pour ce client - Les credentials Chorus Pro configurés pour ce client - Les certificats de signature électronique configurés pour ce client  **Réponse :** ```json {   \"access\": \"eyJ0eXAiOiJKV1QiLCJhbGc...\",  // Token d'accès (validité: 30 min)   \"refresh\": \"eyJ0eXAiOiJKV1QiLCJhbGc...\"  // Token de rafraîchissement (validité: 7 jours) } ```  **Avantages :** - ✅ Automatisation complète (CI/CD, scripts) - ✅ Gestion programmatique des tokens - ✅ Support du refresh token pour renouveler automatiquement l'accès - ✅ Intégration facile dans n'importe quel langage/outil  #### 🖥️ Méthode 2 : Génération via Dashboard (Alternative)  **URL :** https://www.factpulse.fr/dashboard/  Cette méthode convient pour des tests rapides ou une utilisation occasionnelle via l'interface graphique.  **Fonctionnement :** - Connectez-vous au dashboard - Utilisez les boutons \"Generate Test Token\" ou \"Generate Production Token\" - Fonctionne pour **tous** les utilisateurs (OAuth et email/password), sans nécessiter de mot de passe  **Types de tokens :** - **Token Test** : Validité 24h, quota 1000 appels/jour (gratuit) - **Token Production** : Validité 7 jours, quota selon votre forfait  **Avantages :** - ✅ Rapide pour tester l'API - ✅ Aucun mot de passe requis - ✅ Interface visuelle simple  **Inconvénients :** - ❌ Nécessite une action manuelle - ❌ Pas de refresh token - ❌ Moins adapté pour l'automatisation  ### 📚 Documentation complète  Pour plus d'informations sur l'authentification et l'utilisation de l'API : https://www.factpulse.fr/documentation-api/     

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package factpulse

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL string
	Description string
	Variables map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "OpenAPI-Generator/2.0.36/go",
		Debug:            false,
		Servers:          ServerConfigurations{
			{
				URL: "",
				Description: "No description provided",
			},
		},
		OperationServers: map[string]ServerConfigurations{
		},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, reportError("Invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return 0, reportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, reportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}
