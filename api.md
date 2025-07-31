# Search

Params Types:

- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyLegalForm">CompanyLegalForm</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyRegisterType">CompanyRegisterType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyLegalForm">CompanyLegalForm</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyRegisterType">CompanyRegisterType</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanySearch">CompanySearch</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchLookupCompanyByURLResponse">SearchLookupCompanyByURLResponse</a>

Methods:

- <code title="get /v0/search/company">client.Search.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchService.FindCompaniesV0">FindCompaniesV0</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchFindCompaniesV0Params">SearchFindCompaniesV0Params</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanySearch">CompanySearch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/search/company">client.Search.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchService.FindCompaniesV1">FindCompaniesV1</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchFindCompaniesV1Params">SearchFindCompaniesV1Params</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanySearch">CompanySearch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/search/lookup">client.Search.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchService.LookupCompanyByURL">LookupCompanyByURL</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchLookupCompanyByURLParams">SearchLookupCompanyByURLParams</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchLookupCompanyByURLResponse">SearchLookupCompanyByURLResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Company

Response Types:

- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyAddress">CompanyAddress</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyCapital">CompanyCapital</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyName">CompanyName</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyPurpose">CompanyPurpose</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyRegister">CompanyRegister</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyRelationType">CompanyRelationType</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#EntityType">EntityType</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetResponse">CompanyGetResponse</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetHoldingsV1Response">CompanyGetHoldingsV1Response</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetOwnersV1Response">CompanyGetOwnersV1Response</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyListShareholdersResponse">CompanyListShareholdersResponse</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetContactResponse">CompanyGetContactResponse</a>

Methods:

- <code title="get /v0/company/{company_id}">client.Company.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, companyID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetParams">CompanyGetParams</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetResponse">CompanyGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/company/{company_id}/holdings">client.Company.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyService.GetHoldingsV1">GetHoldingsV1</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, companyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetHoldingsV1Response">CompanyGetHoldingsV1Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/company/{company_id}/owners">client.Company.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyService.GetOwnersV1">GetOwnersV1</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, companyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetOwnersV1Response">CompanyGetOwnersV1Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/company/{company_id}/shareholders">client.Company.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyService.ListShareholders">ListShareholders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, companyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyListShareholdersResponse">CompanyListShareholdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/company/{company_id}/contact">client.Company.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyService.GetContact">GetContact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, companyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetContactResponse">CompanyGetContactResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Document

Response Types:

- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#DocumentGetResponse">DocumentGetResponse</a>

Methods:

- <code title="get /v0/document/{document_id}">client.Document.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#DocumentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, documentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#DocumentGetResponse">DocumentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/document/{document_id}/download">client.Document.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#DocumentService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, documentID <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Jobs

## Document

Response Types:

- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#JobDocumentNewResponse">JobDocumentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#JobDocumentGetResponse">JobDocumentGetResponse</a>

Methods:

- <code title="post /v0/jobs/document">client.Jobs.Document.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#JobDocumentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#JobDocumentNewParams">JobDocumentNewParams</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#JobDocumentNewResponse">JobDocumentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/jobs/document/{id}">client.Jobs.Document.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#JobDocumentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#JobDocumentGetResponse">JobDocumentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Autocomplete

Response Types:

- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#AutocompleteAutocompleteCompaniesV1Response">AutocompleteAutocompleteCompaniesV1Response</a>

Methods:

- <code title="get /v1/autocomplete/company">client.Autocomplete.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#AutocompleteService.AutocompleteCompaniesV1">AutocompleteCompaniesV1</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#AutocompleteAutocompleteCompaniesV1Params">AutocompleteAutocompleteCompaniesV1Params</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#AutocompleteAutocompleteCompaniesV1Response">AutocompleteAutocompleteCompaniesV1Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
