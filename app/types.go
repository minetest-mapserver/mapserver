package app

import (
	"mapserver/layer"
)

type Config struct {
	ConfigVersion             int                     `json:"configversion"`
	Port                      int                     `json:"port"`
	EnablePrometheus          bool                    `json:"enableprometheus"`
	EnableRendering           bool                    `json:"enablerendering"`
	EnableSearch              bool                    `json:"enablesearch"`
	EnableInitialRendering    bool                    `json:"enableinitialrendering"`
	EnableTransparency        bool                    `json:"enabletransparency"`
	EnableMediaRepository     bool                    `json:"enablemediarepository"`
	Webdev                    bool                    `json:"webdev"`
	WebApi                    *WebApiConfig           `json:"webapi"`
	Layers                    []*layer.Layer          `json:"layers"`
	RenderingFetchLimit       int                     `json:"renderingfetchlimit"`
	RenderingJobs             int                     `json:"renderingjobs"`
	RenderingQueue            int                     `json:"renderingqueue"`
	IncrementalRenderingTimer string                  `json:"incrementalrenderingtimer"`
	MapObjects                *MapObjectConfig        `json:"mapobjects"`
	MapBlockAccessorCfg       *MapBlockAccessorConfig `json:"mapblockaccessor"`
	DefaultOverlays           []string                `json:"defaultoverlays"`
}

type MapBlockAccessorConfig struct {
	Expiretime string `json:"expiretime"`
	Purgetime  string `json:"purgetime"`
	MaxItems   int    `json:"maxitems"`
}

type MapObjectConfig struct {
	Areas              bool `json:"areas"`
	Bones              bool `json:"bones"`
	Protector          bool `json:"protector"`
	XPProtector        bool `json:"xpprotector"`
	PrivProtector      bool `json:"privprotector"`
	TechnicQuarry      bool `json:"technic_quarry"`
	TechnicSwitch      bool `json:"technic_switch"`
	TechnicAnchor      bool `json:"technic_anchor"`
	TechnicReactor     bool `json:"technic_reactor"`
	LuaController      bool `json:"luacontroller"`
	Digiterms          bool `json:"digiterms"`
	Digilines          bool `json:"digilines"`
	Travelnet          bool `json:"travelnet"`
	MapserverPlayer    bool `json:"mapserver_player"`
	MapserverPOI       bool `json:"mapserver_poi"`
	MapserverLabel     bool `json:"mapserver_label"`
	MapserverTrainline bool `json:"mapserver_trainline"`
	MapserverBorder    bool `json:"mapserver_border"`
	TileServerLegacy   bool `json:"tileserverlegacy"`
	Mission            bool `json:"mission"`
	Jumpdrive          bool `json:"jumpdrive"`
	Smartshop          bool `json:"smartshop"`
	Fancyvend          bool `json:"fancyvend"`
	ATM                bool `json:"atm"`
	Train              bool `json:"train"`
	TrainSignal        bool `json:"trainsignal"`
	Minecart           bool `json:"minecart"`
	Locator            bool `json:"locator"`
	Signs              bool `json:"signs"`
}

type WebApiConfig struct {
	//mapblock debugging
	EnableMapblock bool `json:"enablemapblock"`

	//mod http bridge secret
	SecretKey string `json:"secretkey"`
}
