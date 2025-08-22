package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// JurisdictionOfIncorporation represents the JurisdictionOfIncorporation schema from the OpenAPI specification
type JurisdictionOfIncorporation struct {
	County string `json:"county,omitempty"`
	State string `json:"state,omitempty"`
	City string `json:"city,omitempty"`
	Country string `json:"country"`
}

// CertificateSummariesV2 represents the CertificateSummariesV2 schema from the OpenAPI specification
type CertificateSummariesV2 struct {
	Certificates []CertificateSummaryV2 `json:"certificates"` // List of certificates for a specified customer.
	Pagination Pagination `json:"pagination"`
}

// CertificateRenew represents the CertificateRenew schema from the OpenAPI specification
type CertificateRenew struct {
	Commonname string `json:"commonName,omitempty"` // The common name of certificate to be secured
	Csr string `json:"csr,omitempty"` // Certificate Signing Request.
	Period int `json:"period,omitempty"` // Number of years for certificate validity period, if different from previous certificate
	Roottype string `json:"rootType,omitempty"` // Root Type. Depending on certificate expiration date, SHA_1 not be allowed. Will default to SHA_2 if expiration date exceeds sha1 allowed date
	Subjectalternativenames []string `json:"subjectAlternativeNames,omitempty"` // Only used for UCC products. An array of subject alternative names to include in certificate. Not including a subject alternative name that was in the previous certificate will remove it from the renewed certificate.
	Callbackurl string `json:"callbackUrl,omitempty"` // Required if client would like to receive stateful actions via callback during certificate lifecyle
}

// CertificateIdentifier represents the CertificateIdentifier schema from the OpenAPI specification
type CertificateIdentifier struct {
	Certificateid string `json:"certificateId"` // The unique identifier of the certificate request. Only present if verified.
}

// ExternalAccountBinding represents the ExternalAccountBinding schema from the OpenAPI specification
type ExternalAccountBinding struct {
	Hmackey string `json:"hmacKey"` // EAB HMAC key for the ACME account
	Keyid string `json:"keyId"` // EAB key identifier for the ACME account.
	Directoryurl string `json:"directoryUrl"` // ACME directory resource URL.
}

// CertificateSummaryV2 represents the CertificateSummaryV2 schema from the OpenAPI specification
type CertificateSummaryV2 struct {
	Certificateid string `json:"certificateId"` // The unique identifier of the certificate request.
	Period int `json:"period"` // Validity period of order. Specified in years.
	Status string `json:"status"` // Certificate status (if issued or revoked): * `CANCELED` - Certificate request was canceled by customer * `DENIED` - Certificate request was denied by customer * `EXPIRED` - Issued certificate has exceeded the valid end date * `ISSUED` - Certificate has been issued and is within validity period * `PENDING_ISSUANCE` - Certificate request has completed domain verification and is in the process of being issued * `PENDING_REKEY` - Previously issued certificate was rekeyed by customer and is in the process of being reissued * `PENDING_REVOCATION` - Previously issued certificate is in the process of being revoked * `REVOKED` - Issued certificate has been revoked * `UNUSED` - Certificate in an error state
	Completedat string `json:"completedAt,omitempty"` // The date the certificate request completed processing (if issued or revoked).
	Serialnumber string `json:"serialNumber,omitempty"` // Serial number of certificate (if issued or revoked).
	Validstartat string `json:"validStartAt,omitempty"` // The start date of the certificate's validity (if issued or revoked).
	Commonname string `json:"commonName"` // Common name for the certificate request.
	Createdat string `json:"createdAt"` // Date that the certificate request was received.
	Revokedat string `json:"revokedAt,omitempty"` // The revocation date of certificate (if revoked).
	Validendat string `json:"validEndAt,omitempty"` // The end date of the certificate's validity (if issued or revoked).
	Renewalavailable bool `json:"renewalAvailable,omitempty"` // Only returned when a renewal is available.
	Slotsize string `json:"slotSize,omitempty"` // Number of subject alternative names (SAN) to be included in certificate (if UCC): * `FIVE` - Five slot UCC request * `TEN` - Ten slot UCC request * `FIFTEEN` - Fifteen slot UCC request * `TWENTY` - Twenty slot UCC request * `THIRTY` - Thirty slot UCC request * `FOURTY` - Fourty slot UCC request * `FIFTY` - Fifty slot UCC request * `ONE_HUNDRED` - One hundred slot UCC request
	Subjectalternativenames []string `json:"subjectAlternativeNames,omitempty"` // Subject Alternative names (if UCC). Collection of subjectAlternativeNames to be included in certificate.
	TypeField string `json:"type"` // Certificate type: * `DV_SSL` - (Domain Validated Secure Sockets Layer) SSL certificate validated using domain name only * `DV_WILDCARD_SSL` - SSL certificate containing subdomains which is validated using domain name only * `EV_SSL` - (Extended Validation) SSL certificate validated using organization information, domain name, business legal status, and other factors * `OV_CODE_SIGNING` - Code signing SSL certificate used by software developers to digitally sign apps. Validated using organization information * `OV_DRIVER_SIGNING` - Driver signing SSL certificate request used by software developers to digitally sign secure code for Windows hardware drivers. Validated using organization information * `OV_SSL` - SSL certificate validated using organization information and domain name * `OV_WILDCARD_SSL` - SSL certificate containing subdomains which is validated using organization information and domain name * `UCC_DV_SSL` - (Unified Communication Certificate) Multi domain SSL certificate validated using domain name only * `UCC_EV_SSL` - Multi domain SSL certificate validated using organization information, domain name, business legal status, and other factors * `UCC_OV_SSL` - Multi domain SSL certificate validated using organization information and domain name
}

// SubjectAlternativeNameDetails represents the SubjectAlternativeNameDetails schema from the OpenAPI specification
type SubjectAlternativeNameDetails struct {
	Subjectalternativename string `json:"subjectAlternativeName"` // Subject alternative name to be included in certificate
	Status string `json:"status"` // Total number of page results
}

// CertificateRevoke represents the CertificateRevoke schema from the OpenAPI specification
type CertificateRevoke struct {
	Reason string `json:"reason"` // Reason for revocation
}

// ErrorField represents the ErrorField schema from the OpenAPI specification
type ErrorField struct {
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
	Message string `json:"message,omitempty"` // Description of the problem with the contents of the field
	Path string `json:"path"` // JSONPath referring to the field within the submitted data containing an error
}

// CertificateContact represents the CertificateContact schema from the OpenAPI specification
type CertificateContact struct {
	Jobtitle string `json:"jobTitle,omitempty"` // Only used for EVSSL. Job title of requestor contact
	Namefirst string `json:"nameFirst"` // First name of requestor contact
	Namelast string `json:"nameLast"` // Last name of requestor contact
	Namemiddle string `json:"nameMiddle,omitempty"` // Middle initial of requestor contact
	Phone string `json:"phone"` // Phone number for requestor contact
	Suffix string `json:"suffix,omitempty"` // Suffix of requestor contact
	Email string `json:"email"` // Email address of requestor contact
}

// Certificate represents the Certificate schema from the OpenAPI specification
type Certificate struct {
	Revokedat string `json:"revokedAt,omitempty"` // The revocation date of certificate (if revoked).
	Roottype string `json:"rootType,omitempty"` // Root Type
	Createdat string `json:"createdAt"` // The date the certificate was ordered.
	Serialnumber string `json:"serialNumber,omitempty"` // Serial number of certificate (if issued or revoked)
	Deniedreason string `json:"deniedReason,omitempty"` // Only present if certificate order has been denied
	Producttype string `json:"productType"` // Certificate product type
	Subjectalternativenames []SubjectAlternativeNameDetails `json:"subjectAlternativeNames,omitempty"` // Contains subject alternative names set
	Commonname string `json:"commonName,omitempty"` // Common name of certificate
	Progress int `json:"progress,omitempty"` // Percentage of completion for certificate vetting
	Status string `json:"status"` // Status of certificate
	Certificateid string `json:"certificateId"` // The unique identifier of the certificate request. Only present if no errors returned
	Period int `json:"period"` // Validity period of order. Specified in years
	Validend string `json:"validEnd,omitempty"` // The end date of the certificate's validity (if issued or revoked).
	Validstart string `json:"validStart,omitempty"` // The start date of the certificate's validity (if issued or revoked).
	Serialnumberhex string `json:"serialNumberHex,omitempty"` // Hexadecmial format for Serial number of certificate(if issued or revoked)
	Slotsize string `json:"slotSize,omitempty"` // Number of subject alternative names(SAN) to be included in certificate
	Organization CertificateOrganization `json:"organization,omitempty"`
	Contact CertificateContact `json:"contact"`
}

// Pagination represents the Pagination schema from the OpenAPI specification
type Pagination struct {
	Last string `json:"last,omitempty"` // URI to access the last page
	Next string `json:"next,omitempty"` // URI to access the next page
	Previous string `json:"previous,omitempty"` // URI to access the previous page
	Total int `json:"total,omitempty"` // Number of records available
	First string `json:"first,omitempty"` // URI to access the first page
}

// CertificateBundle represents the CertificateBundle schema from the OpenAPI specification
type CertificateBundle struct {
	Pems PEMCertificates `json:"pems"`
	Serialnumber string `json:"serialNumber"` // Serial number of certificate requested
}

// CertificateAction represents the CertificateAction schema from the OpenAPI specification
type CertificateAction struct {
	TypeField string `json:"type"`
	Createdat string `json:"createdAt"` // Date action created
}

// CertificateOrganization represents the CertificateOrganization schema from the OpenAPI specification
type CertificateOrganization struct {
	Name string `json:"name"` // Name of organization that owns common name
	Phone string `json:"phone"` // Phone number for organization
	Registrationagent string `json:"registrationAgent,omitempty"` // Only for EVSSL.
	Registrationnumber string `json:"registrationNumber,omitempty"` // Only for EVSSL.
	Address CertificateAddress `json:"address"`
	Assumedname string `json:"assumedName,omitempty"` // Only for EVSSL. The DBA(does business as) name for the organization.
	Jurisdictionofincorporation JurisdictionOfIncorporation `json:"jurisdictionOfIncorporation,omitempty"`
}

// CertificateEmailHistory represents the CertificateEmailHistory schema from the OpenAPI specification
type CertificateEmailHistory struct {
	Body string `json:"body"` // Email message
	Dateentered string `json:"dateEntered"` // Date email sent
	Fromtype string `json:"fromType"` // Email from address
	Id int `json:"id"` // Email Id
	Recipients string `json:"recipients"` // Email address email was sent
	Subject string `json:"subject"` // Email subject
	Templatetype string `json:"templateType"` // Email template type name
	Accountid int `json:"accountId"` // Shopper Id requested certificate
}

// CertificateReissue represents the CertificateReissue schema from the OpenAPI specification
type CertificateReissue struct {
	Roottype string `json:"rootType,omitempty"` // Root Type. Depending on certificate expiration date, SHA_1 not be allowed. Will default to SHA_2 if expiration date exceeds sha1 allowed date
	Subjectalternativenames []string `json:"subjectAlternativeNames,omitempty"` // Only used for UCC products. An array of subject alternative names to include in certificate.
	Callbackurl string `json:"callbackUrl,omitempty"` // Required if client would like to receive stateful action via callback during certificate lifecyle
	Commonname string `json:"commonName,omitempty"` // The common name of certificate to be secured
	Csr string `json:"csr,omitempty"` // Certificate Signing Request.
	Delayexistingrevoke int `json:"delayExistingRevoke,omitempty"` // In hours, time to delay revoking existing certificate after issuance of new certificate. If revokeExistingCertOnIssuance is enabled, this value will be ignored
	Forcedomainrevetting []string `json:"forceDomainRevetting,omitempty"` // Optional field. Domain verification will be required for each domain listed here. Specify a value of * to indicate that all domains associated with the request should have their domain information reverified.
}

// ErrorLimit represents the ErrorLimit schema from the OpenAPI specification
type ErrorLimit struct {
	Retryaftersec int `json:"retryAfterSec"` // Number of seconds to wait before attempting a similar request
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
	Fields []ErrorField `json:"fields,omitempty"` // List of the specific fields, and the errors found with their contents
	Message string `json:"message,omitempty"` // Description of the error
}

// CertificateDetailV2 represents the CertificateDetailV2 schema from the OpenAPI specification
type CertificateDetailV2 struct {
	Contact CertificateContact `json:"contact"`
	Roottype string `json:"rootType,omitempty"` // Root type: * `GODADDY_SHA_1` - GoDaddy (Secure Hash Algorithm 1) SHA-1 root type * `GODADDY_SHA_2` - GoDaddy (Secure Hash Algorithm 2) SHA-2 root type * `STARFIELD_SHA_1` - Starfield SHA-1 root type * `STARFIELD_SHA_2` - Starfield SHA-2 root type
	Slotsize string `json:"slotSize,omitempty"` // Number of subject alternative names (SAN) to be included in certificate (if UCC): * `FIVE` - Five slot UCC request * `TEN` - Ten slot UCC request * `FIFTEEN` - Fifteen slot UCC request * `TWENTY` - Twenty slot UCC request * `THIRTY` - Thirty slot UCC request * `FOURTY` - Fourty slot UCC request * `FIFTY` - Fifty slot UCC request * `ONE_HUNDRED` - One hundred slot UCC request
	Csr string `json:"csr,omitempty"` // Certificate signing request (if present) in PEM format
	Revokedat string `json:"revokedAt,omitempty"` // The revocation date of certificate (if revoked).
	Serialnumberhex string `json:"serialNumberHex,omitempty"` // Hexadecmial format for Serial number of certificate(if issued or revoked)
	Organization CertificateOrganization `json:"organization,omitempty"`
	Validstartat string `json:"validStartAt,omitempty"` // The start date of the certificate's validity (if issued or revoked).
	Commonname string `json:"commonName"` // Common name of certificate
	Status string `json:"status"` // Certificate status (if issued or revoked): * `CANCELED` - Certificate request was canceled by customer * `DENIED` - Certificate request was denied by customer\n * `EXPIRED` - Issued certificate has exceeded the valid end date * `ISSUED` - Certificate has been issued and is within validity period * `PENDING_ISSUANCE` - Certificate request has completed domain verification and is in the process of being issued * `PENDING_REKEY` - Previously issued certificate was rekeyed by customer and is in the process of being reissued * `PENDING_REVOCATION` - Previously issued certificate is in the process of being revoked * `REVOKED` - Issued certificate has been revoked\n * `UNUSED` - Certificate in an error state
	Period int `json:"period"` // Validity period of order. Specified in years
	Subjectalternativenames []string `json:"subjectAlternativeNames,omitempty"` // Subject Alternative names. Collection of subjectAlternativeNames to be included in certificate.
	TypeField string `json:"type"` // Certificate type: * `DV_SSL` - (Domain Validated Secure Sockets Layer) SSL certificate validated using domain name only * `DV_WILDCARD_SSL` - SSL certificate containing subdomains which is validated using domain name only * `EV_SSL` - (Extended Validation) SSL certificate validated using organization information, domain name, business legal status, and other factors * `OV_CODE_SIGNING` - Code signing SSL certificate used by software developers to digitally sign apps. Validated using organization information * `OV_DRIVER_SIGNING` - Driver signing SSL certificate request used by software developers to digitally sign secure code for Windows hardware drivers. Validated using organization information * `OV_SSL` - SSL certificate validated using organization information and domain name * `OV_WILDCARD_SSL` - SSL certificate containing subdomains which is validated using organization information and domain name * `UCC_DV_SSL` - (Unified Communication Certificate) Multi domain SSL certificate validated using domain name only * `UCC_EV_SSL` - Multi domain SSL certificate validated using organization information, domain name, business legal status, and other factors * `UCC_OV_SSL` - Multi domain SSL certificate validated using organization information and domain name
	Progress int `json:"progress,omitempty"` // Percentage of completion for certificate vetting
	Createdat string `json:"createdAt"` // The date the certificate was ordered.
	Completedat string `json:"completedAt,omitempty"` // The date the certificate request completed processing.
	Deniedreason string `json:"deniedReason,omitempty"` // Only present if certificate order has been denied
	Validendat string `json:"validEndAt,omitempty"` // The end date of the certificate's validity (if issued or revoked).
	Serialnumber string `json:"serialNumber,omitempty"` // Serial number of certificate (if issued or revoked)
	Certificateid string `json:"certificateId"` // The unique identifier of the certificate request. Only present if no errors returned
	Renewalavailable bool `json:"renewalAvailable,omitempty"` // Only returned when a renewal is available.
}

// CertificateOrganizationCreate represents the CertificateOrganizationCreate schema from the OpenAPI specification
type CertificateOrganizationCreate struct {
	Assumedname string `json:"assumedName,omitempty"` // Only for EVSSL. The DBA(does business as) name for the organization.
	Name string `json:"name"` // Name of organization that owns common name
	Phone string `json:"phone"` // Phone number for organization
	Registrationagent string `json:"registrationAgent,omitempty"` // Only for EVSSL.
	Registrationnumber string `json:"registrationNumber,omitempty"` // Only for EVSSL.
	Address CertificateAddress `json:"address,omitempty"`
}

// CertificateCallback represents the CertificateCallback schema from the OpenAPI specification
type CertificateCallback struct {
	Callbackurl string `json:"callbackUrl"` // Callback url registered to receive stateful actions
}

// CertificateAddress represents the CertificateAddress schema from the OpenAPI specification
type CertificateAddress struct {
	Address1 string `json:"address1"` // Address line 1 of organization address
	Address2 string `json:"address2,omitempty"` // Address line 2 of organization address
	City string `json:"city,omitempty"` // City/Locality of organization address
	Country string `json:"country"` // Two character country code of organization
	Postalcode string `json:"postalCode,omitempty"` // Postal code of organization address
	State string `json:"state,omitempty"` // Full name of State/Province/Territory of organization address
}

// CertificateCreate represents the CertificateCreate schema from the OpenAPI specification
type CertificateCreate struct {
	Csr string `json:"csr"` // Certificate Signing Request
	Slotsize string `json:"slotSize,omitempty"` // Number of subject alternative names(SAN) to be included in certificate
	Subjectalternativenames []string `json:"subjectAlternativeNames,omitempty"` // Subject Alternative names. Collection of subjectAlternativeNames to be included in certificate.
	Contact CertificateContact `json:"contact"`
	Producttype string `json:"productType"` // Type of product requesting a certificate. Only required non-renewal
	Roottype string `json:"rootType,omitempty"` // Root Type. Depending on certificate expiration date, SHA_1 not be allowed. Will default to SHA_2 if expiration date exceeds sha1 allowed date
	Callbackurl string `json:"callbackUrl,omitempty"` // Required if client would like to receive stateful actions via callback during certificate lifecyle
	Commonname string `json:"commonName,omitempty"` // Name to be secured in certificate. If provided, CN field in CSR will be ignored.
	Intelvpro bool `json:"intelVPro,omitempty"` // Only used for OV
	Organization CertificateOrganizationCreate `json:"organization,omitempty"`
	Period int `json:"period"` // Number of years for certificate validity period
}

// PEMCertificates represents the PEMCertificates schema from the OpenAPI specification
type PEMCertificates struct {
	Cross string `json:"cross,omitempty"` // CA Cross Intermediate certificate in PEM format
	Intermediate string `json:"intermediate,omitempty"` // CA Signing Intermediate certificate in PEM format
	Root string `json:"root,omitempty"` // CA Root certificate in PEM format
	Certificate string `json:"certificate"` // End entity certificate in PEM format
}

// DomainVerificationDetail represents the DomainVerificationDetail schema from the OpenAPI specification
type DomainVerificationDetail struct {
	Createdat string `json:"createdAt"` // Timestamp indicating when the domain verification process was started
	Dcetoken string `json:"dceToken,omitempty"` // DCE verification type token (if DCE verification type).
	Domain string `json:"domain"` // Domain name
	Domainentityid int `json:"domainEntityId"` // A unique identifier that can be leveraged for retrieving domain verification related information. Primarily used when troubleshooting a request
	Modifiedat string `json:"modifiedAt"` // Timestamp indicating when the domain verification process was last updated
	Status string `json:"status"` // Domain verification status: * `AWAITING` - Verification pending customer input * `INVALID` - SAN connected to a cancelled request * `COMPLETED` - Verification completed * `FAILED_VERIFICATION` - Verification failed * `PENDING_POSSIBLE_FRAUD` - Flagged for a system level fraud review * `VERIFIED_POSSIBLE_FRAUD` - Fraud detection reviewed but verified * `DROPPED` - SAN dropped from request * `REVOKED_CERT` - Certificate revoked * `DROPPED_GOOGLE_SAFE_BROWSING` - SAN dropped from request due to Google Safe Browsing check * `DROPPED_CERTIFICATE_AUTHORITY_AUTHORIZATION` - SAN dropped from request due to Certificate Authorization Authority DNS record check
	TypeField string `json:"type"` // Domain verification type: * `AUTO_GENERATED_DOMAIN_ACCESS_EMAIL_ADMIN` - Domain verified using domain control verification email sent to admin@<your.domain.com> * `AUTO_GENERATED_DOMAIN_ACCESS_EMAIL_ADMINSTRATOR` - Domain verified using domain control verification email sent to administrator@<your.domain.com> * `AUTO_GENERATED_DOMAIN_ACCESS_EMAIL_HOST_MASTER` - Domain verified using domain control verification email sent to hostmaster@<your.domain.com> * `AUTO_GENERATED_DOMAIN_ACCESS_EMAIL_POST_MASTER` - Domain verified using domain control verification email sent to postmaster@<your.domain.com> * `AUTO_GENERATED_DOMAIN_ACCESS_EMAIL_WEB_MASTER` - Domain verified using domain control verification email sent to webmaster@<your.domain.com> * `DOMAIN_ACCESS_EMAIL` - Domain verified using a domain access email * `DOMAIN_ACCESS_LETTER` - Customer completed a domain access letter which was used for domain verification * `DOMAIN_CONTROL_EMAIL` - Domain verified using HTML file or DNS zone file text value * `DOMAIN_ZONE_CONTROL` - DNS zone file containing a pre-generated text value used for domain verification * `MANUAL_DOMAIN_ACCESS_EMAIL` - DAE sent to an email address manually entered by a rep * `PREVIOUS_DOMAIN_ACCESS_EMAIL` - Customers domain access email for a prior certificate request was used for domain verification * `REGISTRATION_AUTHORITY_DOMAIN_ACCESS_LETTER` - Representative reviewed a customer provided domain access letter and verified domain * `REGISTRATION_AUTHORITY_DOMAIN_ZONE_CONTROL` - Representative verified domain using a manual domain zone control check * `REGISTRATION_AUTHORITY_OVERRIDE` - Representative verified domain using alternative methods * `REGISTRATION_AUTHORITY_WEBSITE_CONTROL` - Representative verified domain using a manual website control check * `CUSTOMER_OWNED` - Validated customer account information used for domain control verification * `WEBSITE_CONTROL` - HTML file in root website directory containing pre-generated value used for domain control verification
	Usage string `json:"usage"` // Type of domain name used for domain verification
	Certificateauthorityauthorization map[string]interface{} `json:"certificateAuthorityAuthorization,omitempty"` // Contains information about the last Certificate Authority Authorization (CAA) Lookup details for the specified domain. In order for a domain to be eligible to be included in the certificate, the entire domain hierarchy must be scanned for DNS CAA records, as outlined by RFC 6844. The absence of any CAA records found in the domain hierarchy indicates that the domain may be included in the certificate. Alternatively, if CAA records are found when scanning the domain hierarchy, the domain may be included in the certificate as long as `godaddy.com` or `starfieldtech.com` is found in the DNS record value. However, if CAA records are found, yet `godaddy.com` or `starfieldtech.com` is not found in any CAA record's value, then we must drop the domain from the certificate request. In the case where there are repeated DNS errors when scanning the domain hierarchy for CAA records, thus ending in an unsuccessful scan, then the domain can still be included in the certificate provided the primary domain is not setup with DNSSEC. Conversely, if DNSSEC is found to be setup on the primary domain when scanning following repeated CAA failures, the domain must be dropped from the certificate request. Finally, if DNS errors persist to the point where a successful DNSSEC query could not be obtained, then the domain must be dropped from the certificate request.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Message string `json:"message,omitempty"` // Description of the error
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
	Fields []ErrorField `json:"fields,omitempty"` // List of the specific fields, and the errors found with their contents
}

// DomainVerificationSummary represents the DomainVerificationSummary schema from the OpenAPI specification
type DomainVerificationSummary struct {
	Domain string `json:"domain"` // Domain name
	Domainentityid int `json:"domainEntityId"` // A unique identifier that can be leveraged for retrieving domain verification related information. Primarily used when troubleshooting a request
	Modifiedat string `json:"modifiedAt"` // Timestamp indicating when the domain verification process was last updated
	Status string `json:"status"` // Domain verification status: * `AWAITING` - Verification pending customer input * `INVALID` - SAN connected to a cancelled request * `COMPLETED` - Verification completed * `FAILED_VERIFICATION` - Verification failed * `PENDING_POSSIBLE_FRAUD` - Flagged for a system level fraud review * `VERIFIED_POSSIBLE_FRAUD` - Fraud detection reviewed but verified * `DROPPED` - SAN dropped from request * `REVOKED_CERT` - Certificate revoked * `DROPPED_GOOGLE_SAFE_BROWSING` - SAN dropped from request due to Google Safe Browsing check * `DROPPED_CERTIFICATE_AUTHORITY_AUTHORIZATION` - SAN dropped from request due to Certificate Authorization Authority DNS record check
	TypeField string `json:"type"` // Domain verification type: * `AUTO_GENERATED_DOMAIN_ACCESS_EMAIL_ADMIN` - Domain verified using domain control verification email sent to admin@<your.domain.com> * `AUTO_GENERATED_DOMAIN_ACCESS_EMAIL_ADMINSTRATOR` - Domain verified using domain control verification email sent to administrator@<your.domain.com> * `AUTO_GENERATED_DOMAIN_ACCESS_EMAIL_HOST_MASTER` - Domain verified using domain control verification email sent to hostmaster@<your.domain.com> * `AUTO_GENERATED_DOMAIN_ACCESS_EMAIL_POST_MASTER` - Domain verified using domain control verification email sent to postmaster@<your.domain.com> * `AUTO_GENERATED_DOMAIN_ACCESS_EMAIL_WEB_MASTER` - Domain verified using domain control verification email sent to webmaster@<your.domain.com> * `DOMAIN_ACCESS_EMAIL` - Domain verified using a domain access email * `DOMAIN_ACCESS_LETTER` - Customer completed a domain access letter which was used for domain verification * `DOMAIN_CONTROL_EMAIL` - Domain verified using HTML file or DNS zone file text value * `DOMAIN_ZONE_CONTROL` - DNS zone file containing a pre-generated text value used for domain verification * `MANUAL_DOMAIN_ACCESS_EMAIL` - DAE sent to an email address manually entered by a rep * `PREVIOUS_DOMAIN_ACCESS_EMAIL` - Customers domain access email for a prior certificate request was used for domain verification * `REGISTRATION_AUTHORITY_DOMAIN_ACCESS_LETTER` - Representative reviewed a customer provided domain access letter and verified domain * `REGISTRATION_AUTHORITY_DOMAIN_ZONE_CONTROL` - Representative verified domain using a manual domain zone control check * `REGISTRATION_AUTHORITY_OVERRIDE` - Representative verified domain using alternative methods * `REGISTRATION_AUTHORITY_WEBSITE_CONTROL` - Representative verified domain using a manual website control check * `CUSTOMER_OWNED` - Validated customer account information used for domain control verification * `WEBSITE_CONTROL` - HTML file in root website directory containing pre-generated value used for domain control verification
	Usage string `json:"usage"` // Type of domain name used for domain verification
	Createdat string `json:"createdAt"` // Timestamp indicating when the domain verification process was started
	Dcetoken string `json:"dceToken,omitempty"` // DCE verification type token (if DCE verification type).
}

// CertificateSiteSeal represents the CertificateSiteSeal schema from the OpenAPI specification
type CertificateSiteSeal struct {
	Html string `json:"html"` // Certificate Seal HTML
}
