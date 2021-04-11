package clientmanagement

type GetClientRenewalDataResponse struct {
	RenewalBy   string `json:"RenewalBy"`
	RenewalData struct {
		AustraliaData struct {
			AnnualIncome string `json:"AnnualIncome"`
			NetWorth     string `json:"NetWorth"`
		} `json:"AustraliaData"`
		FinlandData struct {
			EuroClearSectorCode string `json:"EuroClearSectorCode"`
		} `json:"FinlandData"`
		HongkongData struct {
			NetWorthInUsd                   string `json:"NetWorthInUsd"`
			TotalEstimatedAnnualIncomeInUsd string `json:"TotalEstimatedAnnualIncomeInUsd"`
		} `json:"HongkongData"`
		JapanData struct {
			AnnualIncomeJpy string   `json:"AnnualIncomeJpy"`
			NetWorthJpy     string   `json:"NetWorthJpy"`
			TickerCodes     []string `json:"TickerCodes"`
		} `json:"JapanData"`
		PersonalInformation struct {
			AdditionalNationalities []struct {
				CountryCode    string  `json:"CountryCode"`
				NationalID     string  `json:"NationalId"`
				NationalIDType float64 `json:"NationalIdType"`
			} `json:"AdditionalNationalities"`
			AdditionalTaxableCountries []struct {
				CountryCode     string `json:"CountryCode"`
				TaxID           string `json:"TaxId"`
				TinNotAvailable bool   `json:"TinNotAvailable"`
			} `json:"AdditionalTaxableCountries"`
			AliasFirstName     string `json:"AliasFirstName"`
			AliasLastName      string `json:"AliasLastName"`
			CityOfBirth        string `json:"CityOfBirth"`
			ContactInformation struct {
				EmailAddress string `json:"EmailAddress"`
				MobileNumber struct {
					CountryCode string `json:"CountryCode"`
					Number      string `json:"Number"`
				} `json:"MobileNumber"`
				PrimaryPhoneNumber struct {
					CountryCode string `json:"CountryCode"`
					Number      string `json:"Number"`
				} `json:"PrimaryPhoneNumber"`
				SecondaryPhoneNumber struct {
					CountryCode string `json:"CountryCode"`
					Number      string `json:"Number"`
				} `json:"SecondaryPhoneNumber"`
			} `json:"ContactInformation"`
			CountryOfBirth        string `json:"CountryOfBirth"`
			DateOfBirth           string `json:"DateOfBirth"`
			EmploymentInformation struct {
				EmployerName     string   `json:"EmployerName"`
				NatureOfBusiness string   `json:"NatureOfBusiness"`
				OccupationTypes  []string `json:"OccupationTypes"`
				Position         string   `json:"Position"`
			} `json:"EmploymentInformation"`
			FirstName               string `json:"FirstName"`
			LastName                string `json:"LastName"`
			NationalityCode         string `json:"NationalityCode"`
			OriginalScriptFirstName string `json:"OriginalScriptFirstName"`
			OriginalScriptLastName  string `json:"OriginalScriptLastName"`
			OtherAddress            struct {
				BuildingName           string `json:"BuildingName"`
				BuildingNumber         string `json:"BuildingNumber"`
				City                   string `json:"City"`
				CoName                 string `json:"CoName"`
				CountryOfResidenceCode string `json:"CountryOfResidenceCode"`
				Floor                  string `json:"Floor"`
				LocalArea              string `json:"LocalArea"`
				PostBoxOffice          string `json:"PostBoxOffice"`
				PostalCode             string `json:"PostalCode"`
				SideDoor               string `json:"SideDoor"`
				State                  string `json:"State"`
				StreetName             string `json:"StreetName"`
				Unit                   string `json:"Unit"`
			} `json:"OtherAddress"`
			PersonalID               string `json:"PersonalId"`
			PersonalIDType           string `json:"PersonalIdType"`
			PoliticallyExposedPerson bool   `json:"PoliticallyExposedPerson"`
			ResidentialAddress       struct {
				BuildingName           string `json:"BuildingName"`
				BuildingNumber         string `json:"BuildingNumber"`
				City                   string `json:"City"`
				CoName                 string `json:"CoName"`
				CountryOfResidenceCode string `json:"CountryOfResidenceCode"`
				Floor                  string `json:"Floor"`
				LocalArea              string `json:"LocalArea"`
				PostBoxOffice          string `json:"PostBoxOffice"`
				PostalCode             string `json:"PostalCode"`
				SideDoor               string `json:"SideDoor"`
				State                  string `json:"State"`
				StreetName             string `json:"StreetName"`
				Unit                   string `json:"Unit"`
			} `json:"ResidentialAddress"`
			TaxID            string `json:"TaxId"`
			TinMissingReason string `json:"TinMissingReason"`
			TinNotAvailable  bool   `json:"TinNotAvailable"`
		} `json:"PersonalInformation"`
		ProfileInformation struct {
			AnnualIncomeInformation struct {
				AnnualSalaryAfterTax          string   `json:"AnnualSalaryAfterTax"`
				SecondaryIncomeOther          string   `json:"SecondaryIncomeOther"`
				SecondarySourcesOfIncome      []string `json:"SecondarySourcesOfIncome"`
				SecondarySourcesOfIncomeTotal string   `json:"SecondarySourcesOfIncomeTotal"`
			} `json:"AnnualIncomeInformation"`
			InvestableAssets struct {
				IntendToInvest           string   `json:"IntendToInvest"`
				PrimarySourcesOfWealth   []string `json:"PrimarySourcesOfWealth"`
				SourceOfWealthOther      string   `json:"SourceOfWealthOther"`
				ValueOfCashAndSecurities string   `json:"ValueOfCashAndSecurities"`
			} `json:"InvestableAssets"`
			TradingProfile struct {
				SoleBoOfAssets bool `json:"SoleBoOfAssets"`
			} `json:"TradingProfile"`
		} `json:"ProfileInformation"`
		RegulatoryInformation struct {
			FatcaDeclaration struct {
				UnitedStatesCitizen   bool   `json:"UnitedStatesCitizen"`
				UnitedStatesTaxID     string `json:"UnitedStatesTaxId"`
				UnitedStatesTaxLiable bool   `json:"UnitedStatesTaxLiable"`
			} `json:"FatcaDeclaration"`
		} `json:"RegulatoryInformation"`
		SingaporeData struct {
			AnnualIncomeSgd string `json:"AnnualIncomeSgd"`
			NetWorthSgd     string `json:"NetWorthSgd"`
		} `json:"SingaporeData"`
		SwitzerlandData struct {
			AnnualIncomeChf string `json:"AnnualIncomeChf"`
			NetWorthChf     string `json:"NetWorthChf"`
		} `json:"SwitzerlandData"`
		UkData struct {
			EstimatedSavingAndInvestmentGbp string `json:"EstimatedSavingAndInvestmentGbp"`
			MonthlyIncomeAfterTaxGbp        string `json:"MonthlyIncomeAfterTaxGbp"`
		} `json:"UkData"`
	} `json:"RenewalData"`
	RenewalEntityID string `json:"RenewalEntityId"`
	RenewalStatus   string `json:"RenewalStatus"`
}
