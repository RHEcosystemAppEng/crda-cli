// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

// Defines values for IssueSeverity.
const (
	CRITICAL IssueSeverity = "CRITICAL"
	HIGH     IssueSeverity = "HIGH"
	LOW      IssueSeverity = "LOW"
	MEDIUM   IssueSeverity = "MEDIUM"
)

// Defines values for ComponentAnalysisParamsProviders.
const (
	ComponentAnalysisParamsProvidersSnyk     ComponentAnalysisParamsProviders = "snyk"
	ComponentAnalysisParamsProvidersTidelift ComponentAnalysisParamsProviders = "tidelift"
)

// Defines values for ComponentAnalysisParamsPkgManager.
const (
	ComponentAnalysisParamsPkgManagerMaven ComponentAnalysisParamsPkgManager = "maven"
)

// Defines values for DependencyAnalysisParamsProviders.
const (
	DependencyAnalysisParamsProvidersSnyk     DependencyAnalysisParamsProviders = "snyk"
	DependencyAnalysisParamsProvidersTidelift DependencyAnalysisParamsProviders = "tidelift"
)

// Defines values for DependencyAnalysisParamsPkgManager.
const (
	DependencyAnalysisParamsPkgManagerMaven DependencyAnalysisParamsPkgManager = "maven"
)

// AnalysisReport defines model for AnalysisReport.
type AnalysisReport struct {
	Dependencies *[]DependencyReport `json:"dependencies,omitempty"`
	Summary      *Summary            `json:"summary,omitempty"`
}

// ComponentRequest defines model for ComponentRequest.
type ComponentRequest = []PackageRef

// CvssVector defines model for CvssVector.
type CvssVector struct {
	AttackComplexity      *string `json:"attackComplexity,omitempty"`
	AttackVector          *string `json:"attackVector,omitempty"`
	AvailabilityImpact    *string `json:"availabilityImpact,omitempty"`
	ConfidentialityImpact *string `json:"confidentialityImpact,omitempty"`
	Cvss                  *string `json:"cvss,omitempty"`
	ExploitCodeMaturity   *string `json:"exploitCodeMaturity,omitempty"`
	IntegrityImpact       *string `json:"integrityImpact,omitempty"`
	PrivilegesRequired    *string `json:"privilegesRequired,omitempty"`
	RemediationLevel      *string `json:"remediationLevel,omitempty"`
	ReportConficende      *string `json:"reportConficende,omitempty"`
	Scope                 *string `json:"scope,omitempty"`
	UserInteraction       *string `json:"userInteraction,omitempty"`
}

// DependenciesSummary defines model for DependenciesSummary.
type DependenciesSummary struct {
	Scanned    *float32 `json:"scanned,omitempty"`
	Transitive *float32 `json:"transitive,omitempty"`
}

// DependencyReport defines model for DependencyReport.
type DependencyReport struct {
	Issues         *[]Issue    `json:"issues,omitempty"`
	Recommendation *PackageRef `json:"recommendation,omitempty"`
	Ref            *PackageRef `json:"ref,omitempty"`

	// Remediations Trusted Content remediation related to identified security vulnerabilities
	Remediations *map[string]Remediation       `json:"remediations,omitempty"`
	Transitive   *[]TransitiveDependencyReport `json:"transitive,omitempty"`
}

// Issue defines model for Issue.
type Issue struct {
	Cves     *[]string      `json:"cves,omitempty"`
	Cvss     *CvssVector    `json:"cvss,omitempty"`
	Id       *string        `json:"id,omitempty"`
	Score    *float32       `json:"score,omitempty"`
	Severity *IssueSeverity `json:"severity,omitempty"`
	Source   *string        `json:"source,omitempty"`
	Title    *string        `json:"title,omitempty"`
}

// IssueSeverity defines model for Issue.Severity.
type IssueSeverity string

// PackageRef defines model for PackageRef.
type PackageRef struct {
	// Name <groupId>:<artifactId> for Java packages
	Name *string `json:"name,omitempty"`

	// Version Package version
	Version *string `json:"version,omitempty"`
}

// Remediation defines model for Remediation.
type Remediation struct {
	IssueRef      *string     `json:"issueRef,omitempty"`
	MavenPackage  *PackageRef `json:"mavenPackage,omitempty"`
	ProductStatus *string     `json:"productStatus,omitempty"`
}

// Summary defines model for Summary.
type Summary struct {
	Dependencies    *DependenciesSummary    `json:"dependencies,omitempty"`
	Vulnerabilities *VulnerabilitiesSummary `json:"vulnerabilities,omitempty"`
}

// TransitiveDependencyReport defines model for TransitiveDependencyReport.
type TransitiveDependencyReport struct {
	Issues       *[]Issue                `json:"issues,omitempty"`
	Ref          *PackageRef             `json:"ref,omitempty"`
	Remediations *map[string]Remediation `json:"remediations,omitempty"`
}

// VulnerabilitiesSummary defines model for VulnerabilitiesSummary.
type VulnerabilitiesSummary struct {
	Critical *float32 `json:"critical,omitempty"`
	Direct   *float32 `json:"direct,omitempty"`
	High     *float32 `json:"high,omitempty"`
	Low      *float32 `json:"low,omitempty"`
	Medium   *float32 `json:"medium,omitempty"`
	Total    *float32 `json:"total,omitempty"`
}

// ComponentAnalysisParams defines parameters for ComponentAnalysis.
type ComponentAnalysisParams struct {
	// Providers List of vulnerability providers to get the report from
	Providers *ComponentAnalysisParamsProviders `form:"providers,omitempty" json:"providers,omitempty"`
}

// ComponentAnalysisParamsProviders defines parameters for ComponentAnalysis.
type ComponentAnalysisParamsProviders string

// ComponentAnalysisParamsPkgManager defines parameters for ComponentAnalysis.
type ComponentAnalysisParamsPkgManager string

// DependencyAnalysisParams defines parameters for DependencyAnalysis.
type DependencyAnalysisParams struct {
	// Providers List of vulnerability providers to get the report from
	Providers *DependencyAnalysisParamsProviders `form:"providers,omitempty" json:"providers,omitempty"`
}

// DependencyAnalysisParamsProviders defines parameters for DependencyAnalysis.
type DependencyAnalysisParamsProviders string

// DependencyAnalysisParamsPkgManager defines parameters for DependencyAnalysis.
type DependencyAnalysisParamsPkgManager string

// ComponentAnalysisJSONRequestBody defines body for ComponentAnalysis for application/json ContentType.
type ComponentAnalysisJSONRequestBody = ComponentRequest
