package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Certificate_siteseal_getHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["theme"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("theme=%v", val))
		}
		if val, ok := args["locale"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("locale=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/certificates/%s/siteSeal%s", cfg.BaseURL, certificateId, queryString)
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
		var result models.CertificateSiteSeal
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

func CreateCertificate_siteseal_getTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_certificates_certificateId_siteSeal",
		mcp.WithDescription("Get Site seal"),
		mcp.WithString("certificateId", mcp.Required(), mcp.Description("Certificate id")),
		mcp.WithString("theme", mcp.Description("This value represents the visual theme of the seal. If seal doesn't exist, default values are used if params not present. If seal does exist, default values will not be used to update unless params present.")),
		mcp.WithString("locale", mcp.Description("Determine locale for text displayed in seal image and verification page. If seal doesn't exist, default values are used if params not present. If seal does exist, default values will not be used to update unless params present.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Certificate_siteseal_getHandler(cfg),
	}
}
