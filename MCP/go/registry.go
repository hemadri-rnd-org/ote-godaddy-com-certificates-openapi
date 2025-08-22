package main

import (
	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	tools_v2 "github.com/mcp-server/mcp-server/tools/v2"
	tools_v1 "github.com/mcp-server/mcp-server/tools/v1"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_v2.CreateGetcertificatedetailbycertidentifierTool(cfg),
		tools_v1.CreateCertificate_resend_email_addressTool(cfg),
		tools_v1.CreateCertificate_callback_getTool(cfg),
		tools_v1.CreateCertificate_callback_replaceTool(cfg),
		tools_v1.CreateCertificate_callback_deleteTool(cfg),
		tools_v1.CreateCertificate_verifydomaincontrolTool(cfg),
		tools_v1.CreateCertificate_resend_emailTool(cfg),
		tools_v1.CreateCertificate_get_entitlementTool(cfg),
		tools_v2.CreateGetdomaindetailsbydomainTool(cfg),
		tools_v1.CreateCertificate_alternate_email_addressTool(cfg),
		tools_v1.CreateCertificate_siteseal_getTool(cfg),
		tools_v1.CreateCertificate_cancelTool(cfg),
		tools_v1.CreateCertificate_downloadTool(cfg),
		tools_v1.CreateCertificate_validateTool(cfg),
		tools_v1.CreateCertificate_getTool(cfg),
		tools_v2.CreateGetcustomercertificatesbycustomeridTool(cfg),
		tools_v1.CreateCertificate_createTool(cfg),
		tools_v2.CreateGetdomaininformationbycertificateidTool(cfg),
		tools_v1.CreateCertificate_action_retrieveTool(cfg),
		tools_v1.CreateCertificate_revokeTool(cfg),
		tools_v1.CreateCertificate_reissueTool(cfg),
		tools_v1.CreateCertificate_email_historyTool(cfg),
		tools_v1.CreateCertificate_download_entitlementTool(cfg),
		tools_v1.CreateCertificate_renewTool(cfg),
		tools_v2.CreateGetacmeexternalaccountbindingTool(cfg),
	}
}
