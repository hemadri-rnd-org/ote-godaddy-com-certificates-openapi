package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Certificate_renewHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		certificateIdVal, ok := args["certificateId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: certificateId"), nil
		}
		certificateId, ok := certificateIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: certificateId"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CertificateRenew
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v1/certificates/%s/renew", cfg.BaseURL, certificateId)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result map[string]interface{}
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

func CreateCertificate_renewTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_certificates_certificateId_renew",
		mcp.WithDescription("Renew active certificate"),
		mcp.WithString("certificateId", mcp.Required(), mcp.Description("Certificate id to renew")),
		mcp.WithArray("subjectAlternativeNames", mcp.Description("Input parameter: Only used for UCC products. An array of subject alternative names to include in certificate. Not including a subject alternative name that was in the previous certificate will remove it from the renewed certificate.")),
		mcp.WithString("callbackUrl", mcp.Description("Input parameter: Required if client would like to receive stateful actions via callback during certificate lifecyle")),
		mcp.WithString("commonName", mcp.Description("Input parameter: The common name of certificate to be secured")),
		mcp.WithString("csr", mcp.Description("Input parameter: Certificate Signing Request.")),
		mcp.WithNumber("period", mcp.Description("Input parameter: Number of years for certificate validity period, if different from previous certificate")),
		mcp.WithString("rootType", mcp.Description("Input parameter: Root Type. Depending on certificate expiration date, SHA_1 not be allowed. Will default to SHA_2 if expiration date exceeds sha1 allowed date")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Certificate_renewHandler(cfg),
	}
}
