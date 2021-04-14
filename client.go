package saxo

import (
	"fmt"
	"strings"
	"time"

	"github.com/imroc/req"
	log "github.com/sirupsen/logrus"
)

type saxoClient struct {
	HttpClient *req.Req
}

const (
	APIBaseURL = "https://gateway.saxobank.com"
	DocBaseURL = "https://www.developer.saxo"
)

var (
	bearerToken = ""
	client      = &saxoClient{req.New()}
)

func init() {
	client.HttpClient.SetTimeout(60 * time.Second)
}

func SetBearerToken(token string) {
	bearerToken = token
}

func GetClient() *saxoClient {
	return client
}

func (saxoClient *saxoClient) DoRequest(method string, url string, body interface{}) (*req.Resp, error) {
	if len(bearerToken) == 0 {
		return nil, fmt.Errorf("no bearer token set")
	}
	log.Infoln("Requesting", method, url)
	authHeader := req.Header{
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", bearerToken),
	}

	resp, err := saxoClient.HttpClient.Do(method, url, authHeader, req.BodyJSON(body))
	if err != nil {
		return nil, err
	} else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}

	return resp, nil
}

func PrepareUrlRoute(url string, params ...RouteParam) string {
	for _, param := range params {
		url = strings.Replace(url, param.key, param.value, 1)
	}
	return url
}

func PrepareUrlParams(url string, params interface{}) string {

	return url
}

type RouteParam struct {
	key   string
	value string
}

func RP(key, value string) RouteParam {
	return RouteParam{key, value}
}

type AssetType string

const (
	// Bond.
	Bond AssetType = "Bond"
	// Cash. Not tradeable!
	Cash = "Cash"
	// Mirrors the price movement of the underlying only if and when the underlying price exceeds the defined barrier. If the certificate expires below the barrier, it offers partial protection/return of investment.
	CertificateBonus = "CertificateBonus"
	// Certificate Capped Bonus.
	CertificateCappedBonus = "CertificateCappedBonus"
	// Guarantees a capped percentage increase of the underlying asset's value above the issue price at expiry/maturity. Max loss is the amount invested multiplied by the CapitalProtection percentage.
	CertificateCappedCapitalProtected = "CertificateCappedCapitalProtected"
	// Capped Outperformance Certificate.
	CertificateCappedOutperformance = "CertificateCappedOutperformance"
	// Certificate Constant Leverage.
	CertificateConstantLeverage = "CertificateConstantLeverage"
	// Yields a capped return if the underlying asset's value is above the specified cap level at expiry. If the underlying's value is below the strike at expiry, the investor received the underlying or equivalent value. Offers direct exposure in underlying at a lower price (discount) with a capped potential profit and limited loss.
	CertificateDiscount = "CertificateDiscount"
	// Certificate Express kick out.
	CertificateExpress = "CertificateExpress"
	// A certificate that mirrors the price movement of the underlying instrument. Often used to trade movements in indicies. Movements can be a fixed ratio of the underlying and can be inverted for bearish/short speculation. Risk is equivalent to owning the underlying.
	CertificateTracker = "CertificateTracker"
	// Guarantees a percentage increase of the underlying asset's value above the issue price at expiry/maturity. Max loss is the amount invested multiplied by the CapitalProtection percentage.
	CertificateUncappedCapitalProtection = "CertificateUncappedCapitalProtection"
	// Provides leveraged returns when the underlying price exceeds the threshold strike price. The amount leverage is defined by the Participation %. When the underlying is below the strike price, the certificate mirrors the underlying price 1:1.
	CertificateUncappedOutperformance = "CertificateUncappedOutperformance"
	// Cfd Index Option.
	CfdIndexOption = "CfdIndexOption"
	// Cfd on unlisted warrant issued by a corporation.
	CfdOnCompanyWarrant = "CfdOnCompanyWarrant"
	// Cfd on Futures.
	CfdOnFutures = "CfdOnFutures"
	// Cfd on Stock Index.
	CfdOnIndex = "CfdOnIndex"
	// Cfd on Stock.
	CfdOnStock = "CfdOnStock"
	// Unlisted warrant issued by a corporation, often physically settled.
	CompanyWarrant = "CompanyWarrant"
	// Contract Futures.
	ContractFutures = "ContractFutures"
	// Futures Option.
	FuturesOption = "FuturesOption"
	// Futures Strategy.
	FuturesStrategy = "FuturesStrategy"
	// Forex Binary Option.
	FxBinaryOption = "FxBinaryOption"
	// Forex Forward.
	FxForwards = "FxForwards"
	// Forex Knock In Option.
	FxKnockInOption = "FxKnockInOption"
	// Forex Knock Out Option.
	FxKnockOutOption = "FxKnockOutOption"
	// Forex No Touch Option.
	FxNoTouchOption = "FxNoTouchOption"
	// Forex One Touch Option.
	FxOneTouchOption = "FxOneTouchOption"
	// Forex Spot.
	FxSpot = "FxSpot"
	// Forex Vanilla Option.
	FxVanillaOption = "FxVanillaOption"
	// Danish investment scheme (“Grantbevis”). Not online tradeable.
	GuaranteedCapital = "GuaranteedCapital"
	// IPO on Stock
	IpoOnStock = "IpoOnStock"
	// Obsolete: Managed Fund.
	ManagedFund = "ManagedFund"
	// MiniFuture.
	MiniFuture = "MiniFuture"
	// Mutual Fund.
	MutualFund = "MutualFund"
	// Danish pooled investment scheme (“Pulje”). Not online tradeable.
	PooledInvestment = "PooledInvestment"
	// SRD. (Service de Règlement Différé) on Stock.
	SrdOnStock = "SrdOnStock"
	// Stock.
	Stock = "Stock"
	// Stock Index.
	StockIndex = "StockIndex"
	// Stock Index Option.
	StockIndexOption = "StockIndexOption"
	// Stock Option.
	StockOption = "StockOption"
	// Warrant
	Warrant = "Warrant"
	// Warrant with two knock-out barriers.
	WarrantDoubleKnockOut = "WarrantDoubleKnockOut"
	// Warrant with a knock-out barrier.
	WarrantKnockOut = "WarrantKnockOut"
	// Knock-out Warrant with no expiry.
	WarrantOpenEndKnockOut = "WarrantOpenEndKnockOut"
	// Warrant with built-in spread.
	WarrantSpread = "WarrantSpread"
)
