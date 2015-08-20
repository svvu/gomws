package gmws

type Products struct {
	*MwsBase
}

// sellerId, authToken, region string
func NewProductsClient(config map[string]string) *MwsClient {
	base := NewMwsBase(config)
	return &Products{MwsBase: base}
}

func (p Products) paramsToAugment(action string) map[string]string {
	params := p.MwsBase.paramsToAugment()
	params["Version"] = p.Version()
	params["Action"] = action
	return params
}

func (p Products) Version() string {
	return "2011-10-01"
}

func (p Products) Name() string {
	return "Products"
}

func (p Products) Path() string {
	return "/" + p.Name() + "/" + p.Version()
}

func (p Products) Endpoint() string {
	return p.Host + "/" + p.Path()
}

func (p Products) GetMatchingProductForId(idType string, idList []string) string {
	action := "GetMatchingProductForId"
	params := Parameters{"IdType": idType, "IdList": idList}
	structedParams := params.StructureKeys("IdList", "Id").ToNormalizedParameters()

	paramsToAugment := p.paramsToAugment(action)
	httpClient := newHttpClient(p.Endpoint(), structedParams)
	httpClient.AugmentParameters(paramsToAugment)
	httpClient.SignQuery(p.getCredential())
	return httpClient.Request()
}
