(function () {

    const ID_MASK = 0XFFF; // Mask to retrieve Slate ID from event
    // Device types that can send HID events
    const EdeviceTypes = {
        slate: 1,
        wpanPushButton: 2,
        unknown: 3
    };

    function WpanKeyListener(logger, udpMulticastConfig) {
        this._logger = logger;
        this._logger.log("WpanKeyListener");
        this._udpMulticastConfig = udpMulticastConfig;
    }
    WpanKeyListener.prototype = {
        init: function () {
            try {
                this._logger.log("WpanKeyListener init");
                this._initKeyListener();
                this._initUdpSocket();
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        _initKeyListener: function () {
            try {
                this._logger.log("WpanKeyListener _initKeyListener");
                window.top.focus();
                window.top.document.addEventListener("keydown", this._onKeyDown.bind(this));
                window.top.document.addEventListener("keyup", this._onKeyUp.bind(this));
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        _initUdpSocket: function () {
            try {
                this._logger.log("WpanKeyListener _initUdpSocket");
                var jsonConfig = {
                    localPort: this._udpMulticastConfig.localPort,
                };
                if (typeof UDPSocket === "function") {
                    this._socket = new UDPSocket(jsonConfig);
                } else {
                    this._socket = new GktUDPSocket(jsonConfig);
                }
                this._socket.joinMulticast(this._udpMulticastConfig.multicastGroup);
                this._socket.onmessage = this._onSocketMessageHandler.bind(this);
                this._socket.onerror = this._onSocketErrorHandler.bind(this);
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        _sendOverUdp: function (message) {
            try {
                //this._socket.send(message);
                this._socket.send(message, this._udpMulticastConfig.multicastGroup, this._udpMulticastConfig.localPort);
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        _onKeyUp: function (event) {
            try {
                this._logger.log("WpanKeyListener _onKeyUp");
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        _onKeyDown: function (event) {
            try {
                this._logger.log("WpanKeyListener _onKeyDown");
                this._logger.log("Event keyCode " + event.keyCode);
                var id = -1;
                var domKeyLocationSlate = event.DOM_KEY_LOCATION_SLATE || 0x80000000;
                var domKeyLocationWpanPushButton = event.DOM_KEY_LOCATION_PUSHBUTTON || 0x10000000;
                var deviceType = EdeviceTypes.unknown;
                var strType;
                if (event.location & domKeyLocationSlate) {
                    deviceType = EdeviceTypes.slate;
                    strType = "slate";
                }
                else if (event.location & domKeyLocationWpanPushButton) {
                    deviceType = EdeviceTypes.wpanPushButton;
                    strType = "wpanPushButton";
                }
                if (deviceType != EdeviceTypes.unknown) {
                    id = event.location & ID_MASK;
                } else {
                    strType = "unknown";
                }
                this._logger.log("Event location is from deviceof id " + id + " and type : " + strType);
                if (deviceType == EdeviceTypes.wpanPushButton) {
                    var jsonToString = {
                        id: id.toString(),
                        keyCode: event.keyCode
                    };
                    var stringifyJsonMessage = JSON.stringify(jsonToString);
                    this._logger.log("send over UDP message: " + stringifyJsonMessage);
                    this._sendOverUdp(stringifyJsonMessage);
                }
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        _onSocketMessageHandler: function (UDPMessageEvent) {
            try {
                var data = "";
                for (var i = 0; i < UDPMessageEvent.data.length; i++) {
                    data += String.fromCharCode(UDPMessageEvent.data[i]);
                }
                this._logger.log("Receive response: '" + data + "'");
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        _onSocketErrorHandler: function (UDPErrorEvent) {
            try {
                this._logger.log("*************** _onSocketErrorHandler ***************");
                this._logger.error("UDPErrorEvent name: " + UDPErrorEvent.name + ", message: " + UDPErrorEvent.message);
            } catch (error) {
                this._logger.errorEx(error);
            }
        }
    };

    window.WpanKeyListener = WpanKeyListener;
})();