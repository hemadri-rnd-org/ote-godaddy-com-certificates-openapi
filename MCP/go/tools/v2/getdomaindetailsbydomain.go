package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetdomaindetailsbydomainHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		customerIdVal, ok := args["customerId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: customerId"), nil
		}
		customerId, ok := customerIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: customerId"), nil
		}
		certificateIdVal, ok := args["certificateId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: certificateId"), nil
		}
		certificateId, ok := certificateIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: certificateId"), nil
		}
		domainVal, ok := args["domain"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: domain"), nil
		}
		domain, ok := domainVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: domain"), nil
		}
		url := fmt.Sprintf("%s/v2/customers/%s/certificates/%s/domainVerifications/%s", cfg.BaseURL, customerId, certificateId, domain)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.DomainVerificationDetail
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetdomaindetailsbydomainTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v2_customers_customerId_certificates_certificateId_domainVerifications_domain",
		mcp.WithDescription("Retrieve detailed information for supplied domain"),
		mcp.WithString("customerId", mcp.Required(), mcp.Description("An identifier for a customer")),
		mcp.WithString("certificateId", mcp.Required(), mcp.Description("Certificate id to lookup")),
		mcp.WithString("domain", mcp.Required(), mcp.Description("A valid domain name in the certificate request")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetdomaindetailsbydomainHandler(cfg),
	}
}
