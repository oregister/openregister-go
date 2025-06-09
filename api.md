# Search

Params Types:

- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyLegalForm">CompanyLegalForm</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyRegisterType">CompanyRegisterType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyLegalForm">CompanyLegalForm</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyRegisterType">CompanyRegisterType</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchFindCompaniesResponse">SearchFindCompaniesResponse</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchLookupCompanyByURLResponse">SearchLookupCompanyByURLResponse</a>

Methods:

- <code title="get /v0/search/company">client.Search.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchService.FindCompanies">FindCompanies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchFindCompaniesParams">SearchFindCompaniesParams</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchFindCompaniesResponse">SearchFindCompaniesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/search/lookup">client.Search.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchService.LookupCompanyByURL">LookupCompanyByURL</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchLookupCompanyByURLParams">SearchLookupCompanyByURLParams</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#SearchLookupCompanyByURLResponse">SearchLookupCompanyByURLResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Company

Response Types:

- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyAddress">CompanyAddress</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyCapital">CompanyCapital</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyName">CompanyName</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyPurpose">CompanyPurpose</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyRegister">CompanyRegister</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#EntityType">EntityType</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetResponse">CompanyGetResponse</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyListShareholdersResponse">CompanyListShareholdersResponse</a>
- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetContactResponse">CompanyGetContactResponse</a>

Methods:

- <code title="get /v0/company/{company_id}">client.Company.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, companyID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetParams">CompanyGetParams</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetResponse">CompanyGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/company/{company_id}/shareholders">client.Company.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyService.ListShareholders">ListShareholders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, companyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyListShareholdersResponse">CompanyListShareholdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/company/{company_id}/contact">client.Company.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyService.GetContact">GetContact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, companyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#CompanyGetContactResponse">CompanyGetContactResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Document

Response Types:

- <a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#DocumentGetResponse">DocumentGetResponse</a>

Methods:

- <code title="get /v0/document/{document_id}">client.Document.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#DocumentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, documentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/oregister/openregister-go">openregister</a>.<a href="https://pkg.go.dev/github.com/oregister/openregister-go#DocumentGetResponse">DocumentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
