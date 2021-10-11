// The Bot entity or data type
package bots

// Type of bot, i.e simple or composite
type Type string

const (
	Simple    Type = "simple"
	Composite Type = "composite"
)

type Scope string

const (
	Enabled  Scope = "enabled"
	Disabled Scope = "disabled"
)

// Pairs of currencies, i.e ETH/BTC
type Pairs []string

// Strategy, i.e short/long
type Strategy string

const (
	Long  Strategy = "long"
	Short Strategy = "short"
)

// TakeProfitType -
type TakeProfitType string

const (
	Total TakeProfitType = "total"
	Base  TakeProfitType = "base"
)

// StrategyList JSON string of strategy types // @todo make this an actual type?
type StrategyList []map[string]interface{}

// LeverageType - Bitmex bots only
type LeverageType string

const (
	Custom       LeverageType = "custom"
	Cross        LeverageType = "cross"
	NotSpecified LeverageType = "not_specified"
)

type ProfitCurrency string

const (
	QuoteCurrency ProfitCurrency = "quote_currency"
	BaseCurrency  ProfitCurrency = "base_currency"
)

type StartOrderType string

const (
	Limit  StartOrderType = "limit"
	Market StartOrderType = "market"
)

// StopLossType -
type StopLossType string

const (
	StopLoss              StopLossType = "stop_loss"
	StopLossAndDisableBot StopLossType = "stop_loss_and_disable_bot"
)

// Bot -
type Bot struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`

	// Fetch from accounts.List() or /ver1/accounts
	AccountID int   `json:"account_id" validate:"required"`
	Pairs     Pairs `json:"pairs" validate:"required"`

	// Default 1
	MaxActiveDeals      int    `json:"max_active_deals"`
	BaseOrderVolume     string `json:"base_order_volume" validate:"required"`
	BaseOrderVolumeType string `json:"base_order_volume_type"`

	// In percentage
	TakeProfit            string `json:"take_profit" validate:"required"`
	SafetyOrderVolume     string `json:"safety_order_volume" validate:"required"`
	SafetyOrderVolumeType string `json:"safety_order_volume_type"`

	// Default 1
	MartingaleVolumeCoefficient string `json:"martingale_volume_coefficient" validate:"required"`

	// Default 1
	MartingaleStepCoefficient string `json:"martingale_step_coefficient" validate:"required"`
	MaxSafetyOrders           int    `json:"max_safety_orders" validate:"required"`
	ActiveSafetyOrdersCount   int    `json:"active_safety_orders_count" validate:"required"`
	StopLossPercentage        string `json:"stop_loss_percentage"`
	Cooldown                  string `json:"cooldown"`

	// Default false
	TrailingEnabled bool `json:"trailing_enabled"`

	// Required if TrailingEnabled is true @todo add validation for this
	TrailingDeviation string `json:"trailing_deviation"`
	BTCPriceLimit     string `json:"btc_price_limit"`

	// Short/Long, default is Long
	Strategy                  Strategy `json:"strategy"`
	SafetyOrderStepPercentage string   `json:"safety_order_step_percentage" validate:"required"`

	// base, total (base)	Percentage: base – from base order, total – from total volume
	TakeProfitType TakeProfitType `json:"take_profit_type" validate:"required"`

	// For manual signals: [{"strategy":"nonstop"}] or []
	// For non-stop(1 pair only): [{"strategy":"nonstop"}]
	// QFL: {"options"=>{"type"=>"original"}, "strategy"=>"qfl"}]
	// TradingView: [{"options"=>{"time"=>"5m", "type"=>"buy_or_strong_buy"}, "strategy"=>"trading_view"}
	StrategyList StrategyList `json:"strategy_list" validate:"required"`

	// custom, cross, not_specified (not_specified)	Used for Bitmex bots only
	LeverageType             LeverageType `json:"leverage_type"`
	LeverageCustomValue      int          `json:"leverage_custom_value"`
	MinPrice                 int          `json:"min_price"`
	MaxPrice                 int          `json:"max_price"`
	StopLossTimeoutEnabled   bool         `json:"stop_loss_timeout_enabled"`
	StopLossTimeoutInSeconds int          `json:"stop_loss_timeout_in_seconds"`
	MinVolumeBTC24H          string       `json:"min_volume_btc_24h"`

	// Bitmex only
	TSLEnabled            bool `json:"tsl_enabled"`
	DealStartDelaySeconds int  `json:"deal_start_delay_seconds"`

	// quote_currency, base_currency
	ProfitCurrency ProfitCurrency `json:"profit_currency"`

	// Limit/market
	StartOrderType StartOrderType `json:"start_order_type"`

	// stop_loss, stop_loss_and_disable_bot
	StopLossType StopLossType `json:"stop_loss_type"`

	// Bot will be disabled after opening this number of deals
	DisableAfterDealsCount int `json:"disable_after_deals_count"`

	// Allow specific number of deals on the same pair. Multibot only.
	AllowedDealsOnSamePair int `json:"allowed_deals_on_same_pair"`
}
