// package api_tests

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"drivers.api/core/authentication"
// 	"drivers.api/routers"
// 	"drivers.api/settings"
// 	"github.com/codegangsta/negroni"
// 	"github.com/stretchr/testify/assert"
// 	. "gopkg.in/check.v1"
// )

// func Test(t *testing.T) {
// 	TestingT(t)
// }

// type MiddlewaresTestSuite struct{}

// var _ = Suite(&MiddlewaresTestSuite{})
// var t *testing.T
// var token string
// var server *negroni.Negroni

// func (s *MiddlewaresTestSuite) SetUpSuite(c *C) {
// 	os.Setenv("GO_ENV", "tests")
// 	settings.Init()
// }

// func (s *MiddlewaresTestSuite) SetUpTest(c *C) {
// 	authBackend := authentication.InitJWTAuthenticationBackend()
// 	assert.NotNil(t, authBackend)
// 	token, _ = authBackend.GenerateToken("1234")

// 	router := routers.InitRoutes()
// 	server = negroni.Classic()
// 	server.UseHandler(router)
// }

// func (s *MiddlewaresTestSuite) TestRequireTokenAuthentication(c *C) {
// 	resource := "/test/hello"

// 	response := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", resource, nil)
// 	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
// 	server.ServeHTTP(response, request)

// 	assert.Equal(t, response.Code, http.StatusOK)
// }

// func (s *MiddlewaresTestSuite) TestRequireTokenAuthenticationInvalidToken(c *C) {
// 	resource := "/test/hello"

// 	response := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", resource, nil)
// 	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", "token"))
// 	server.ServeHTTP(response, request)

// 	assert.Equal(t, response.Code, http.StatusUnauthorized)
// }

// func (s *MiddlewaresTestSuite) TestRequireTokenAuthenticationEmptyToken(c *C) {
// 	resource := "/test/hello"

// 	response := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", resource, nil)
// 	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", ""))
// 	server.ServeHTTP(response, request)

// 	assert.Equal(t, response.Code, http.StatusUnauthorized)
// }

// func (s *MiddlewaresTestSuite) TestRequireTokenAuthenticationWithoutToken(c *C) {
// 	resource := "/test/hello"

// 	response := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", resource, nil)
// 	server.ServeHTTP(response, request)

// 	assert.Equal(t, response.Code, http.StatusUnauthorized)
// }

// func (suite *MiddlewaresTestSuite) TestRequireTokenAuthenticationAfterLogout(c *C) {
// 	resource := "/test/hello"

// 	requestLogout, _ := http.NewRequest("GET", resource, nil)
// 	requestLogout.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
// 	services.Logout(requestLogout)

// 	response := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", resource, nil)
// 	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
// 	server.ServeHTTP(response, request)

// 	assert.Equal(t, response.Code, http.StatusUnauthorized)
// }
