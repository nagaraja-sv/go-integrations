package forcego

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

// SuccessResponse is to
type SuccessResponse struct {
	Accesstoken string `json:"access_token"`
	Instanceurl string `json:"instance_url"`
	ID          string `json:"id"`
	Tokentype   string `json:"token_type"`
	Issuedat    string `json:"issued_at"`
	Signature   string `json:"signature"`
}
type Application struct {
	FirstName                      string `json:"FirstName__c,omitempty"`
	LastName                       string `json:"LastName__c,omitempty"`
	HomeAddress                    string `json:"HomeAddress1__c,omitempty"`
	HomeAddressCountry             string `json:"HomeAddressCountry__c,omitempty"`
	HomeAddressCity                string `json:"HomeAddressCity__c,omitempty"`
	HomeAddressState               string `json:"HomeAddressState__c,omitempty"`
	HomeAddressZip                 string `json:"HomeAddressZip__c,omitempty"`
	PreferredMethodofCommunication string `json:"PreferredCommunication__c,omitempty"`
	BranchofService                string `json:"BranchOfService__c,omitempty"`
	ServiceStartDate               string `json:"ServiceStartDate__c,omitempty"`
	DischargeType                  string `json:"DischargeType__c,omitempty"`
	ServiceEndDate                 string `json:"ServiceEndDate__c,omitempty"`
	DutyStartDate                  string `json:"DutyStartDate__c,omitempty"`
	DutyEndDate                    string `json:"DutyEndDate__c,omitempty"`
	DutyCampaign                   string `json:"DutyCampaign__c,omitempty"`
	DutyResponsibilities           string `json:"DutyResponsibilities__c,omitempty"`
	Supervisor1                    string `json:"Supervisor1__c,omitempty"`
}

type Response struct {
	ID      string `json:"id"`
	Success bool   `json:"success"`
}

//GetAccountHandler is to
func GetAccountHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	SuccessResponse := &SuccessResponse{}
	Application := &Application{}
	Response := &Response{}

	//resp, err := client.Post("https://login.salesforce.com/services/oauth2/token?grant_type=password&client_id=3MVG9ZL0ppGP5UrAg6Xz54_N2LAWOBA5Q.WT0h7sZgDFmrtztZUZibtu0_o9yDiFPYyq9QJQB1OeOl55p9uyo&client_secret=1889558308606956941&username=nagaraja@457.com&password=Spiderman@57valynZY8bqq5SL5asRmJyKZ2f", "", nil)
	resp, err := client.Post("https://test.salesforce.com/services/oauth2/token?grant_type=password&client_id=3MVG9zZht._ZaMumhxW.PRK5V_Du7eu9olkLjG9z3QAI7DLBS7LGRXGow1wfKJrcvvnz63tdF1Te8F5j_MlOO&client_secret=5527763041516414834&username=hari.jella@vifm.us.dev1&password=KashyakSupport@1X8ru4ps4iKsQYegrn5FowF1QI", "", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		//bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {

		}
		//bodyString := string(bodyBytes)
		if err := json.NewDecoder(resp.Body).Decode(SuccessResponse); err != nil {

			return
		}
		// Write the
		/*if err := json.NewEncoder(w).Encode(SuccessResponse); err != nil {
			panic(err)
		}*/
	}

	//Genrating Access Tocken

	//Calling Endpoint

	if SuccessResponse.Accesstoken != "" {
		//r.Header.Set("Authorization", "Bearer"+SuccessResponse.Accesstoken)
		//r.Body()
		if err := json.NewDecoder(r.Body).Decode(Application); err != nil {

			return
		}
		//url := "https://nagaraja457-dev-ed.my.salesforce.com/services/data/v32.0/sobjects/Account"
		url := "https://cs21.salesforce.com/services/data/v32.0/sobjects/Student__c"
		fmt.Println("URL:>", url)

		//var jsonStr = []byte(`{"name":"Buy cheese and bread for breakfast."}`)
		var jsonStr, err = json.Marshal(Application)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

		req.Header.Set("Authorization", "Bearer "+SuccessResponse.Accesstoken)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(Response); err != nil {

			return
		}
		if err := json.NewEncoder(w).Encode(Response); err != nil {
			panic(err)
		}

	}

}
