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

// Bot represents a trading bot with its configuration parameters.
type Bot struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`

	// Basic bot settings
	AccountID int      `json:"account_id" validate:"required"`
	Pairs     Pairs    `json:"pairs" validate:"required"`
	Strategy  Strategy `json:"strategy"`

	// Deal settings
	MaxActiveDeals      int    `json:"max_active_deals"`
	BaseOrderVolume     string `json:"base_order_volume" validate:"required"`
	BaseOrderVolumeType string `json:"base_order_volume_type"`
	TakeProfit          string `json:"take_profit" validate:"required"`

	// Safety order settings
	SafetyOrderVolume           string `json:"safety_order_volume" validate:"required"`
	SafetyOrderVolumeType       string `json:"safety_order_volume_type"`
	MartingaleVolumeCoefficient string `json:"martingale_volume_coefficient" validate:"required"`
	MartingaleStepCoefficient   string `json:"martingale_step_coefficient" validate:"required"`
	MaxSafetyOrders             int    `json:"max_safety_orders" validate:"required"`
	ActiveSafetyOrdersCount     int    `json:"active_safety_orders_count" validate:"required"`
	SafetyOrderStepPercentage   string `json:"safety_order_step_percentage" validate:"required"`

	// Stop loss settings
	StopLossPercentage       string       `json:"stop_loss_percentage"`
	StopLossTimeoutEnabled   bool         `json:"stop_loss_timeout_enabled"`
	StopLossTimeoutInSeconds int          `json:"stop_loss_timeout_in_seconds"`
	StopLossType             StopLossType `json:"stop_loss_type"`

	// Trailing settings
	TrailingEnabled   bool   `json:"trailing_enabled"`
	TrailingDeviation string `json:"trailing_deviation"`

	// Miscellaneous settings
	BTCPriceLimit string `json:"btc_price_limit"`
	Cooldown      string `json:"cooldown"`

	// base, total (base)	Percentage: base – from base order, total – from total volume
	TakeProfitType TakeProfitType `json:"take_profit_type" validate:"required"`

	// For manual signals: [{"strategy":"nonstop"}] or []
	// For non-stop(1 pair only): [{"strategy":"nonstop"}]
	// QFL: {"options"=>{"type"=>"original"}, "strategy"=>"qfl"}]
	// TradingView: [{"options"=>{"time"=>"5m", "type"=>"buy_or_strong_buy"}, "strategy"=>"trading_view"}
	StrategyList StrategyList `json:"strategy_list" validate:"required"`

	// custom, cross, not_specified (not_specified)	Used for Bitmex bots only
	LeverageType        LeverageType `json:"leverage_type"`
	LeverageCustomValue int          `json:"leverage_custom_value"`
	MinPrice            int          `json:"min_price"`
	MaxPrice            int          `json:"max_price"`
	MinVolumeBTC24H     string       `json:"min_volume_btc_24h"`

	// Bitmex only
	TSLEnabled            bool `json:"tsl_enabled"`
	DealStartDelaySeconds int  `json:"deal_start_delay_seconds"`

	// quote_currency, base_currency
	ProfitCurrency ProfitCurrency `json:"profit_currency"`

	// Limit/market
	StartOrderType StartOrderType `json:"start_order_type"`

	// Bot will be disabled after opening this number of deals
	DisableAfterDealsCount int `json:"disable_after_deals_count"`

	// Allow specific number of deals on the same pair. Multibot only.
	AllowedDealsOnSamePair int `json:"allowed_deals_on_same_pair"`
}
