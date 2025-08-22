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

func Certificate_resend_emailHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		emailIdVal, ok := args["emailId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: emailId"), nil
		}
		emailId, ok := emailIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: emailId"), nil
		}
		url := fmt.Sprintf("%s/v1/certificates/%s/email/%s/resend", cfg.BaseURL, certificateId, emailId)
		req, err := http.NewRequest("POST", url, nil)
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

func CreateCertificate_resend_emailTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_certificates_certificateId_email_emailId_resend",
		mcp.WithDescription("Resend an email"),
		mcp.WithString("certificateId", mcp.Required(), mcp.Description("Certificate id to resend email")),
		mcp.WithString("emailId", mcp.Required(), mcp.Description("Email id for email to resend")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Certificate_resend_emailHandler(cfg),
	}
}
