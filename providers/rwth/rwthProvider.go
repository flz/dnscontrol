package rwth

import (
	"encoding/json"
	"fmt"

	"github.com/StackExchange/dnscontrol/v3/providers"
)

type rwthProvider struct {
	apiToken string
	zones    map[string]zone
}

// features is used to let dnscontrol know which features are supported by the RWTH DNS Admin.
var features = providers.DocumentationNotes{
	providers.CanAutoDNSSEC:          providers.Unimplemented("Supported by RWTH but not implemented yet."),
	providers.CanGetZones:            providers.Can(),
	providers.CanUseAlias:            providers.Cannot(),
	providers.CanUseCAA:              providers.Can(),
	providers.CanUseDS:               providers.Unimplemented("DS records are only supported at the apex and require a different API call that hasn't been implemented yet."),
	providers.CanUseLOC:              providers.Cannot(),
	providers.CanUseNAPTR:            providers.Cannot(),
	providers.CanUsePTR:              providers.Can("PTR records with empty targets are not supported"),
	providers.CanUseSRV:              providers.Can("SRV records with empty targets are not supported."),
	providers.CanUseSSHFP:            providers.Can(),
	providers.CanUseTLSA:             providers.Cannot(),
	providers.DocCreateDomains:       providers.Cannot(),
	providers.DocDualHost:            providers.Cannot(),
	providers.DocOfficiallySupported: providers.Cannot(),
}

// init registers the registrar and the domain service provider with dnscontrol.
func init() {
	fns := providers.DspFuncs{
		Initializer:   New,
		RecordAuditor: AuditRecords,
	}
	providers.RegisterDomainServiceProviderType("RWTH", fns, features)
}

// New allocates a DNS service provider.
func New(settings map[string]string, _ json.RawMessage) (providers.DNSServiceProvider, error) {
	if settings["api_token"] == "" {
		return nil, fmt.Errorf("missing RWTH api_token")
	}

	api := &rwthProvider{apiToken: settings["api_token"]}

	return api, nil
}
