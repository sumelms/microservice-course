package internal

// @title           Sumé LMS Course API
// @version         1.0
// @description     This is the Sumé LMS API for Course Microservice
// @termsOfService  https://sumelms.com/docs/terms

// @contact.name   LMS Support
// @contact.url    https://sumelms.com/support
// @contact.email  support@sumelms.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used

// @securitydefinitions.oauth2.application  OAuth2Application
// @tokenUrl                                https://sso.sumelms.com/oauth/token
// @scope.write                             Grants write access
// @scope.admin                             Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit  OAuth2Implicit
// @authorizationUrl                     https://sso.sumelms.com/oauth/authorize
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.password  OAuth2Password
// @tokenUrl                             https://sso.sumelms.com/oauth/token
// @scope.read                           Grants read access
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode  OAuth2AccessCode
// @tokenUrl                               https://sso.sumelms.com/oauth/token
// @authorizationUrl                       https://sso.sumelms.com/oauth/authorize
// @scope.admin                            Grants read and write access to administrative information
