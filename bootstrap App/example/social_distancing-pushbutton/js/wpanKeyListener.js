(function () {

	// Serial
	const Ci = Components.interfaces;
	const nsISystemAdapterSerial = Ci.nsISystemAdapterSerial;

	const SERIAL_SYS_PATH = "/dev/ttyACM0";

	var logger;
	var processEnOceanButton;
	var onError;
	var onStart;

	function Controller(aSerialAdapterManager) {
		try {
			logger.log("Controller");
			this._serialCmd = "";
			this._bleUtils = new BLEUtils();
			this._serialAdapterManager = aSerialAdapterManager;
			this._Estatus = {
				scanning: 1,
				setObserverRole: 2,
				stopScan: 3,
				connecting: 4,
				stopConnecting: 5,
				connected: 6,
				rebooting: 7
			};
			this._EconnectedStep = {
				none: 0,
				centralIdentification: 1,
				peripheralAuthentication: 2,
				enableHIDReport: 3
			};
			this._status = this._Estatus.scanning;
			this._connectedStep = this._EconnectedStep.none;
			this._messagePartsAsByteArray = [];
			this._enableObserverRole = true;
			this._waitButtonRelease = false;
			this._buttonRepetition = false;

			var decimalPinCode;
			if (systemGeneralSettings.field2) {
				if ((0 <= systemGeneralSettings.field2) && (systemGeneralSettings.field2 <= 9999)) {
					decimalPinCode = systemGeneralSettings.field2;
					logger.log("Found authentication pincode in player variables " + decimalPinCode);
				} else {
					logger.error("Found authentication pincode in player variables but value is wrong ");
				}
			} else {
				logger.log("No data in player variable field2 : no authentication value");
			}

			var hexaPin = parseInt(decimalPinCode).toString(16);
			for (var i = hexaPin.length; i < 4; i++) {
				hexaPin = "0" + hexaPin;
			}
			logger.log("hexaPin " + hexaPin);
			this._hexadecimalAuthenticationPincode = hexaPin;
			this._cmdReturnTimer;
		}
		catch (e) {
			logger.error("Exception  " + e + " line = " + e.lineNumber);
			return null;
		}
	}
	Controller.prototype = {
		initController: function () {
			logger.log("initController");
			this._connectedStep = this._EconnectedStep.none;
			this._serialAdapterManager.sendCmd("F,0064,0064\n");
		},
		onCTSChanged: function (newValue) {
			logger.log("onCTSChanged value " + newValue);
		},
		onDSRChanged: function (newValue) {
			logger.log("onDSRChanged value " + newValue);
		},
		onRINGChanged: function (newValue) {
			logger.log("onRINGChanged value " + newValue);
		},
		onDCDChanged: function (newValue) {
			logger.log("onDCDChanged value " + newValue);
		},
		onDataAvailable: function (inputStream) {
			try {
				var bytesAvailable = 0;
				bytesAvailable = inputStream.available(bytesAvailable);
				while (bytesAvailable > 0) {
					this._serialCmd += String.fromCharCode(inputStream.read8());
					--bytesAvailable;
					if (this._serialCmd.search(("\n")) != -1) {
						//logger.log("Found carriage return cmd is : " + this._serialCmd);
						if (this._serialCmd.search("ERR") != -1) {
							logger.log("Response is ERR");
							this._onErrorReceived();
						}
						else if (this._serialCmd.search("AOK") != -1) {
							logger.log("Response is AOK");
							this._onAcknowledgeReceived();
						}
						else if (this._serialCmd.search("CMD") != -1) {
							logger.log("Response is CMD");
							this._onCMDReceived();
						}
						else if (this._serialCmd.search("Reboot") != -1) {
							logger.log("Response is Reboot");
						}
						else if (this._serialCmd.search("Connected") != -1) {
							logger.log("Response is Connected");
							this._onConnectedReceived();
						}
						else if (this._serialCmd.search("Connection End") != -1) {
							logger.log("Response is Connection End");
							this._onConnectionEndReceived();
						}
						else if (this._serialCmd.search("Notify,") != -1) {
							logger.log("Notification : Connected device wrote a value in a charac handle");
							this._onNotifyReceived(this._serialCmd);
						}
						else if ((DEVICE_TO_CONNECT_TYPE === "enocean") && this._bleUtils.isEnOceanAdvertise(this._serialCmd)) {
							logger.log("Advertise is from an EnOcean device " + this._serialCmd);
							if (this._bleUtils.isEnOceanToProcess()) {
								logger.log("Current advertise is the EnOcean device we want to process buttons from");
								this._processSwitchStatus(this._bleUtils.getEnOceanButtonHexSwitchStatus());
							}
						}
						else if ((DEVICE_TO_CONNECT_TYPE === "slate") && this._bleUtils.isSlateAdvertise(this._serialCmd)) {
							logger.log("Advertise is from a Slate");
							if (this._bleUtils.isSlateToConnect()) {
								logger.log("Current advertise is the Slate we want to connect to");
								this._stopScan();
							}
						}
						this._serialCmd = "";
					}
				}
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_startScan: function () {
			try {
				logger.log("_startScan");
				this._status = this._Estatus.scanning;
				this._serialAdapterManager.sendCmd("F,0064,0064\n");
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_setObserverRole: function () {
			try {
				logger.log("_setObserverRole");
				this._status = this._Estatus.setObserverRole;
				this._serialAdapterManager.sendCmd("J,1\n");
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_stopScan: function () {
			try {
				logger.log("_stopScan");
				this._status = this._Estatus.stopScan;
				this._serialAdapterManager.sendCmd("X\n");
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_connectToSlate: function () {
			try {
				logger.log("_connectToSlate");
				this._status = this._Estatus.connecting;
				var connectionCmd = "E,0,";
				connectionCmd += this._bleUtils.getSlateToConnectMacAddr();
				connectionCmd += "\n";
				this._serialAdapterManager.sendCmd(connectionCmd);
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_stopConnecting: function () {
			try {
				logger.log("_stopConnecting");
				this._status = this._Estatus.stopConnecting;
				this._serialAdapterManager.sendCmd("Z\n");
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_onAcknowledgeReceived: function () {
			try {
				logger.log("_onAcknowledgeReceived");
				clearTimeout(this._cmdReturnTimer);
				switch (this._status) {
					case this._Estatus.none:
						logger.log("Status none : this is start scan AOK");
						this._status = this._Estatus.scanning;
						break;
					case this._Estatus.scanning:
						logger.log("Status scanning");
						if (this._enableObserverRole) {
							this._setObserverRole();
						}
						break;
					case this._Estatus.setObserverRole:
						logger.log("Status setObserverRole");
						break;
					case this._Estatus.stopScan:
						logger.log("Status stopScan");
						this._connectToSlate();
						break;
					case this._Estatus.connecting:
						logger.log("Status connecting, waiting for connected input");
						this._cmdReturnTimer = setTimeout(this._cmdTimeoutCallback.bind(this), CONNECTED_WAIT_TIMEOUT);
						break;
					case this._Estatus.stopConnecting:
						logger.log("Status stopConnecting");
						this._startScan();
						break;
					case this._Estatus.connected:
						this._onAOKwhileConnected();
						break;
					case this._Estatus.rebooting:
						logger.error("AOK received during wrong state, should not happen");
						break;
					default:
						logger.log("Default case should not happen");
				}
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_onAOKwhileConnected: function () {
			try {
				logger.log("_onAOKwhileConnected");
				if (DEVICE_TO_CONNECT_TYPE === "slate") { //AKYR
					var cmdValue = "";
					if (this._connectedStep === this._EconnectedStep.centralIdentification) {
						logger.log("CentralIdentification OK");
						this._connectedStep = this._EconnectedStep.peripheralAuthentication;
						if (SLATE_TO_CONNECT_AUTHENTICATION_METHOD === "pincode") {
							logger.log("Authentication mode is by pin code");
							cmdValue = "0100" + this._hexadecimalAuthenticationPincode;
							cmdValue = this._serialAdapterManager.reverseStringbyPair(cmdValue);
							cmdValue = "CUWV,000000000000000040000000034E4553," + cmdValue + "\n";
							this._serialAdapterManager.sendCmd(cmdValue);
						} else {
							this._serialAdapterManager.sendCmd("CUWV,000000000000000040000000034E4553,00000000\n");
						}
					}
					else if (this._connectedStep === this._EconnectedStep.peripheralAuthentication) {
						logger.log("PeripheralAuthentication OK");
						this._connectedStep = this._EconnectedStep.enableHIDReport;
						this._serialAdapterManager.sendCmd("CUWC,2A22,1\n");
					}
					else if (this._connectedStep === this._EconnectedStep.enableHIDReport) {
						logger.log("Status connected : enableHIDReport OK");
					}
				}
				logger.log("_onAOKwhileConnected end");
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_onErrorReceived: function () {
			try {
				logger.log("_onErrorReceived");
				clearTimeout(this._cmdReturnTimer);
				switch (this._status) {
					case this._Estatus.scanning:
					case this._Estatus.setObserverRole:
					case this._Estatus.stopScan:
					case this._Estatus.connecting:
					case this._Estatus.stopConnecting:
					case this._Estatus.connected:
						logger.log("Error while generic state");
						this._rebootController();
						break;
					case this._Estatus.rebooting:
						logger.error("Error while rebooting state");
						break;
					default:
						logger.log("Default case should not happen");
				}
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_onCMDReceived: function () {
			try {
				logger.log("_onCMDReceived status " + this._status);
				clearTimeout(this._cmdReturnTimer);
				switch (this._status) {
					case this._Estatus.none:
						logger.log("CMD while status is none : this is start-up CMD");
						this._startScan();
						break;
					case this._Estatus.scanning:
					case this._Estatus.setObserverRole:
					case this._Estatus.stopScan:
					case this._Estatus.connecting:
					case this._Estatus.stopConnecting:
					case this._Estatus.connected:
						logger.error("CMD while generic state, should not happen");
						this._startScan();
						break;
					case this._Estatus.rebooting:
						logger.log("CMD while rebooting state : ready to scan");
						this._startScan();
						break;
					default:
						logger.log("Default case should not happen");
				}
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_onConnectedReceived: function () {
			try {
				logger.log("_onConnectedReceived");
				switch (this._status) {
					case this._Estatus.scanning:
					case this._Estatus.setObserverRole:
					case this._Estatus.stopScan:
					case this._Estatus.rebooting:
					case this._Estatus.connected:
					case this._Estatus.stopConnecting:
						logger.error("Connected received during wrong state, should not happen");
						return null;
					case this._Estatus.connecting:
						logger.log("Connected to Slate");
						clearTimeout(this._cmdReturnTimer);
						this._status = this._Estatus.connected;
						break;
					default:
						logger.log("Default case should not happen");
						return null;
				}
				logger.log("Central identification process");
				this._connectedStep = this._EconnectedStep.centralIdentification;
				this._serialAdapterManager.sendCmd("CUWV,000000000000000040000000024E4553,00000006\n");
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_onConnectionEndReceived: function () {
			try {
				logger.log("_onConnectionEndReceived");
				switch (this._status) {
					case this._Estatus.scanning:
					case this._Estatus.setObserverRole:
					case this._Estatus.stopScan:
					case this._Estatus.rebooting:
					case this._Estatus.connecting:
					case this._Estatus.stopConnecting:
						logger.error("Connection End received during wrong state, should not happen");
						this._startScan();
						break;
					case this._Estatus.connected:
						logger.log("Connection End received, ready to scan");
						this._startScan();
						break;
					default:
						logger.log("Default case should not happen");
				}
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_onNotifyReceived: function (aNotifyStr) {
			try {
				logger.log("_onNotifyReceived");
				var strPos = 0;
				var bleHandle = "";
				var bleValue = "";
				var keyCode = "";
				switch (this._status) {
					case this._Estatus.scanning:
					case this._Estatus.setObserverRole:
					case this._Estatus.stopScan:
					case this._Estatus.rebooting:
					case this._Estatus.connecting:
					case this._Estatus.stopConnecting:
						logger.error("Notify received during wrong state, should not happen");
						break;
					case this._Estatus.connected:
						logger.log("Notify received while connected : " + aNotifyStr);
						strPos = aNotifyStr.search(',');
						bleHandle = aNotifyStr.slice(strPos + 1, aNotifyStr.length);
						strPos = bleHandle.search(',');
						bleValue = bleHandle.slice(strPos + 1, bleHandle.length);
						bleHandle = bleHandle.slice(0, strPos);
						bleValue = this._serialAdapterManager.reverseStringbyPair(bleValue.slice(0, bleValue.length - bleHandle.length + 1));
						keyCode = bleValue.slice(6, 8);
						logger.log("bleHandle " + bleHandle + " bleValue " + bleValue + " keyCode " + keyCode);
						this._mediaManager.processSlateKey(keyCode);
						break;
					default:
						logger.log("Default case should not happen");
				}
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_cmdTimeoutCallback: function () {
			try {
				logger.log("_cmdTimeoutCallback");
				switch (this._status) {
					case this._Estatus.scanning:
					case this._Estatus.setObserverRole:
					case this._Estatus.stopScan:
					case this._Estatus.stopConnecting:
					case this._Estatus.rebooting:
					case this._Estatus.connected:
						logger.log("Command timeout during generic state");
						this._rebootController();
						break;
					case this._Estatus.connecting:
						logger.log("Command timeout during connecting");
						this._stopConnecting();
						break;
					default:
						logger.log("Default case should not happen");
						return null;
				}
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_rebootController: function () {
			try {
				logger.log("_rebootController");
				this._status = this._Estatus.rebooting;
				this._serialAdapterManager.sendCmd("R,1\n");
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_processSwitchStatus: function (aHexSwitchStatus) {
			try {
				logger.log("_processSwitchStatus status " + aHexSwitchStatus + " Repetition " + this._buttonRepetition);
				const buttonA0Mask = 0x02;
				const buttonA1Mask = 0x04;
				const buttonB0Mask = 0x08;
				const buttonB1Mask = 0x10;
				var decimalSwitchStatus = parseInt(aHexSwitchStatus.toString(), 16);
				if (decimalSwitchStatus && !this._buttonRepetition) {
					this._buttonRepetition = true;
					logger.log("Action type is press");
					if (decimalSwitchStatus & buttonA0Mask) {
						logger.log("Button is A0");
						//this._mediaManager.processEnOceanButton("A0");
						processEnOceanButton("A0");
					}
					if (decimalSwitchStatus & buttonA1Mask) {
						logger.log("Button is A1");
						processEnOceanButton("A1");
					}
					if (decimalSwitchStatus & buttonB0Mask) {
						logger.log("Button is B0");
						processEnOceanButton("B0");
					}
					if (decimalSwitchStatus & buttonB1Mask) {
						logger.log("Button is B1");
						processEnOceanButton("B1");
					}
					var self = this;
					setTimeout(function () {
						self._buttonRepetition = false;
					}, 500);
				}
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_toUTF8Array: function (str, aMaxDisplayChar) {
			try {
				// TODO(user): Use native implementations if/when available
				var out = [], p = 0;
				var nbDisplayChar = 0;
				for (var i = 0; i < str.length; i++) {
					var c = str.charCodeAt(i);
					if (c < 128) {
						out[p++] = c;
						nbDisplayChar++;
					} else if (c < 2048) {
						out[p++] = (c >> 6) | 192;
						out[p++] = (c & 63) | 128;
						nbDisplayChar++;
					} else if (
						((c & 0xFC00) == 0xD800) && (i + 1) < str.length &&
						((str.charCodeAt(i + 1) & 0xFC00) == 0xDC00)) {
						// Surrogate Pair
						c = 0x10000 + ((c & 0x03FF) << 10) + (str.charCodeAt(++i) & 0x03FF);
						out[p++] = (c >> 18) | 240;
						out[p++] = ((c >> 12) & 63) | 128;
						out[p++] = ((c >> 6) & 63) | 128;
						out[p++] = (c & 63) | 128;
						nbDisplayChar++;
					} else {
						out[p++] = (c >> 12) | 224;
						out[p++] = ((c >> 6) & 63) | 128;
						out[p++] = (c & 63) | 128;
						nbDisplayChar++;
					}
					if (nbDisplayChar >= aMaxDisplayChar) {
						logger.error("_toUTF8Array max display char length reach " + nbDisplayChar);
						break;
					}
				}
				logger.log("_toUTF8Array nbDisplayChar " + nbDisplayChar + " out length " + out.length);
				return out;
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_toHexaString: function (aInt) {
			hexString = aInt.toString(16);
			if (hexString.length % 2) {
				hexString = '0' + hexString;
			}
			return hexString;
		}
	}


	function BLEUtils() {
		try {
			logger.log("BLEUtils");
			this._currentDevice = new BLEDevice();
			this._remoteDeviceMacAddr = "";
			this._valueToConnect = "";


			DEVICE_TO_CONNECT_TYPE = "enocean"; // values : "enocean", "slate"
			if (DEVICE_TO_CONNECT_TYPE === "slate") {
				if (systemGeneralSettings.field1) {
					this._valueToConnect = systemGeneralSettings.field1;
					logger.log("Found method to connect value in player variable field 1 " + this._valueToConnect);
					if (SLATE_TO_CONNECT_SELECTION_METHOD == "psn") {
						logger.log("method to connect is by PSN : converting PSN to hexa");
						this._valueToConnect = this._convertDecInnesPsnToHex(this._valueToConnect);
						if (this._valueToConnect == null) {
							logger.log("Wrong PSN format, catch up with default");
						}
					} else {
						logger.log("method to connect is by Mac Addr");
						if (this._valueToConnect.length != 12) {
							logger.error("Wrong Mac format, catch up with default");
						}
					}
				} else {
					logger.log("No data found in player variable field1, will read from all");
				}
			} else if (DEVICE_TO_CONNECT_TYPE === "enocean") {
				if (systemGeneralSettings.field1) {
					this._valueToConnect = systemGeneralSettings.field1;
					logger.log("Found method to connect value in player variable field 1 " + this._valueToConnect);
				}
				else {
					logger.log("No data found in player variable field1, will read from all");
				}
			} else {
				logger.error("Device to connect type unknown");
			}

		}
		catch (e) {
			logger.error("Exception  " + e + " line = " + e.lineNumber);
			return null;
		}
	}

	BLEUtils.prototype =
		{
			isSlateAdvertise: function (aAdvertiseStr) {
				try {
					var macAddr;
					var hostName;
					var rxPower;
					var strPos;

					//Retrieving Mac Address and cutting the used info
					strPos = aAdvertiseStr.search(',');
					if (strPos === -1) {
						return false;
					}
					if (strPos != 12) {
						return false;
					}
					macAddr = aAdvertiseStr.substring(0, strPos);
					aAdvertiseStr = aAdvertiseStr.slice(strPos + 1, aAdvertiseStr.length);

					//Retrieving Mac Address and cutting the used info
					strPos = aAdvertiseStr.search(',');
					if (strPos === -1) {
						return false;
					}
					aAdvertiseStr = aAdvertiseStr.slice(strPos + 1, aAdvertiseStr.length);

					strPos = aAdvertiseStr.search(',');
					if (strPos === -1) {
						return false;
					}
					hostName = aAdvertiseStr.substring(0, strPos);
					aAdvertiseStr = aAdvertiseStr.slice(strPos + 1, aAdvertiseStr.length);

					strPos = aAdvertiseStr.search(',');
					if (strPos === -1) {
						return false;
					}
					rxPower = aAdvertiseStr.substring(strPos + 1, aAdvertiseStr.length);
					aAdvertiseStr = aAdvertiseStr.slice(0, strPos);

					//Now aAdvertiseStr is the advertise UUID, verify size
					if (aAdvertiseStr.length != 32) {
						return false;
					}
					//Verify PSN root
					if (aAdvertiseStr.search("05A") != 0) {
						return false;
					}
					//Verify Innes PnPId
					if (aAdvertiseStr.search("4E4553") != 26) {
						return false;
					}

					this._currentDevice.macAddr = macAddr;
					this._currentDevice.hostName = hostName;
					this._currentDevice.rxPower = rxPower;
					this._currentDevice.uuid = aAdvertiseStr;
					this._currentDevice.hexaPsn = aAdvertiseStr.substring(0, 8);
					return true;
				}
				catch (e) {
					logger.error("Exception  " + e + " line = " + e.lineNumber);
					return null;
				}
			},
			isEnOceanAdvertise: function (aAdvertiseStr) {
				try {
					var macAddr;
					var rxPower;
					var strPos;
					var uuid;
					var manufacturerId;
					const enOceanManufacturerId = "03DA";
					var switchStatus;

					strPos = aAdvertiseStr.search("Brcst:");
					if (strPos === -1) {
						return false;
					}
					uuid = aAdvertiseStr.slice(strPos + 6, aAdvertiseStr.length);
					manufacturerId = uuid.substring(4, 8);
					manufacturerId = this._reverseStringbyPair(manufacturerId);

					if (manufacturerId != enOceanManufacturerId) {
						return false;
					}

					switchStatus = uuid.substring(16, 18);

					//Retrieving Mac Address and cutting the used info
					strPos = aAdvertiseStr.search(',');
					if (strPos === -1) {
						return false;
					}
					if (strPos != 12) {
						return false;
					}
					macAddr = aAdvertiseStr.substring(0, strPos);
					aAdvertiseStr = aAdvertiseStr.slice(strPos + 1, aAdvertiseStr.length);

					//Retrieving Mac Address and cutting the used info
					strPos = aAdvertiseStr.search(',');
					if (strPos === -1) {
						return false;
					}
					aAdvertiseStr = aAdvertiseStr.slice(strPos + 1, aAdvertiseStr.length);

					strPos = aAdvertiseStr.search(',');
					if (strPos === -1) {
						return false;
					}
					rxPower = aAdvertiseStr.substring(strPos + 1, aAdvertiseStr.length);
					aAdvertiseStr = aAdvertiseStr.slice(0, strPos);


					this._currentDevice.macAddr = macAddr;
					this._currentDevice.rxPower = rxPower;
					this._currentDevice.switchStatus = parseInt(switchStatus);


					logger.log("Advertise is EnOcean, mac " + macAddr);
					return true;
				}
				catch (e) {
					logger.error("Exception  " + e + " line = " + e.lineNumber);
					return null;
				}
			},
			isSlateToConnect: function () {
				try {
					var isSlateToConnect = false;

					if (SLATE_TO_CONNECT_SELECTION_METHOD === "psn") {
						if (this._currentDevice.hexaPsn === this._valueToConnect) {
							isSlateToConnect = true;
						}
					} else if (SLATE_TO_CONNECT_SELECTION_METHOD === "mac") {
						if (this._currentDevice.macAddr === this._valueToConnect) {
							isSlateToConnect = true;
						}
					} else {
						logger.error("Slate to connect selection method unknown");
						return null;
					}

					if (isSlateToConnect) {
						this._remoteDeviceMacAddr = this._currentDevice.macAddr;
						logger.log(" It is the slate to connect  " + this._currentDevice.hostName +
							" macAddr: " + this._currentDevice.macAddr +
							" uuid: " + this._currentDevice.uuid +
							" hexaPsn: " + this._currentDevice.hexaPsn +
							" rxPower: " + this._currentDevice.rxPower);
						return true;
					}
					logger.log("It is not the slate to connect");
					return false;
				}
				catch (e) {
					logger.error("Exception  " + e + " line = " + e.lineNumber);
					return null;
				}
			},
			isEnOceanToProcess: function () {
				try {
					var isEnOceanToProcess = false;

					if ((this._valueToConnect == "") || (this._currentDevice.macAddr === this._valueToConnect)) {
						isEnOceanToProcess = true;
					}

					if (isEnOceanToProcess) {
						this._remoteDeviceMacAddr = this._currentDevice.macAddr;
						logger.log(" It is EnOcean device to process button from  " +
							" macAddr: " + this._currentDevice.macAddr +
							" switchStatus: " + this._currentDevice.switchStatus +
							" rxPower: " + this._currentDevice.rxPower);
						return true;
					}
					logger.log("It is not the enocean button to process");
					return false;
				}
				catch (e) {
					logger.error("Exception  " + e + " line = " + e.lineNumber);
					return null;
				}
			},
			getSlateToConnectMacAddr: function () {
				try {
					if (SLATE_TO_CONNECT_SELECTION_METHOD === "mac") {
						return this._valueToConnect;
					} else {
						return this._remoteDeviceMacAddr;
					}
				}
				catch (e) {
					logger.error("Exception  " + e + " line = " + e.lineNumber);
					return null;
				}
			},
			getEnOceanButtonHexSwitchStatus: function () {
				try {
					return this._currentDevice.switchStatus;
				}
				catch (e) {
					logger.error("Exception  " + e + " line = " + e.lineNumber);
					return null;
				}
			},
			_convertDecInnesPsnToHex: function (aDecPsn) {
				try {
					logger.log("_convertDecInnesPsnToHex dec PSN " + aDecPsn);
					if (aDecPsn.length === 0) {
						logger.error("Psn is empty");
						return null;
					}
					if (aDecPsn.length != 11) {
						logger.error("Wrong format for PSN to convert");
						return null;
					}
					var leftPart, rightPart, innesNb, hexPsn;
					leftPart = aDecPsn.substring(0, 4);
					rightPart = aDecPsn.substring(6, 11);
					innesNb = aDecPsn.substring(4, 5);

					leftPart = parseInt(leftPart).toString(16);
					for (var i = leftPart.length; i < 3; i++) {
						leftPart = '0' + leftPart;
					}
					hexPsn = leftPart;

					innesNb = parseInt(innesNb).toString(16);
					hexPsn += innesNb;

					rightPart = parseInt(rightPart).toString(16);
					for (var i = rightPart.length; i < 4; i++) {
						rightPart = '0' + rightPart;
					}
					hexPsn += rightPart;

					hexPsn = hexPsn.toUpperCase();
					logger.log("hexPsn " + hexPsn);
					return hexPsn;
				}
				catch (e) {
					logger.error("Exception  " + e + " line = " + e.lineNumber);
					return null;
				}
			},
			_reverseStringbyPair: function (aValue) {
				if ((aValue.length % 2) != 0) {
					logger.error("reverseStringbyPair has to be multiple of 2 in size");
					return null;
				}
				var array = Array.from(aValue);
				array.reverse();
				for (var i = 0; i < array.length; i += 2) {
					[array[i], array[i + 1]] = [array[i + 1], array[i]];
				}
				return array.join('').toString();
			}
		}

	function BLEDevice() {
		this.macAddr = "";
		this.hostname = "";
		this.uuid = "";
		this.rxPower = "";
		this.hexaPsn = "";
		this.switchStatus;
		try {
			logger.log("BLEDevice");
		}
		catch (e) {
			logger.error("Exception  " + e + " line = " + e.lineNumber);
			return null;
		}
	}

	function SerialAdapterManager(appProcessEnOceanButton, appStart, appError, appLogger) {

		logger = {
			log: function (message) {
				appLogger(message);
			},
			error: function (message) {
				appLogger("ERROR : " + message, 2);
			}
		};

		processEnOceanButton = appProcessEnOceanButton;
		onStart = appStart;
		onError = appError;

		this._serial = null;
		this._inited = false;
		this._inputStream = null;
		this._outputStream = null;
		this._slateToConnectMacAddr = "";
		try {
			this._reinit(SERIAL_SYS_PATH);
		}
		catch (e) {
			logger.error("Exception  " + e + " line = " + e.lineNumber);
			return null;
		}
	}
	
	SerialAdapterManager.prototype = {
		_getSerial: function (aSysPath) {
			try {
				var system = systemManager;
				var serials = system.getAdaptersByIId(Ci.nsISystemAdapterSerial);
				var serial = null;
				logger.log("serials.length = " + serials.length);
				for (var i = 0; i < serials.length; i++) {
					serial = serials.queryElementAt(i, Ci.nsISystemAdapterSerial);
					logger.log("serial.sysPath = " + serial.sysPath);
					if (aSysPath === "*") {
						break;
					}
					if (serial.sysPath === aSysPath) {
						break;
					}
				}
				return serial;
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_initSerial: function (aSysPath) {
			try {
				logger.log("_initSerial");
				if (!this._serial) {
					this._serial = this._getSerial(aSysPath);
				}
				if (this._serial) {
					this._serial.recieveMode = nsISystemAdapterSerial.RECIEVE_MODE_ASYNC;
					this._serial.setConfig(nsISystemAdapterSerial.DIRECTION_IN | nsISystemAdapterSerial.DIRECTION_OUT,
						nsISystemAdapterSerial.BAUD115200, 8,
						nsISystemAdapterSerial.PARITY_NONE,
						nsISystemAdapterSerial.STOPBIT_1,
						nsISystemAdapterSerial.FLOWCONTROL_HARDWARE);
					if (!this._listener) {
						logger.log("Creating serial listener");
						this._listener = new Controller(this);
					}
					this._serial.addListener(this._listener);
					this._serial.open();
					this._outputStream = this._serial.outputStream;
					this._inputStream = this._serial.inputStream;
					this._inited = true;
					onStart();
				}
			}
			catch (e) {
				logger.error("_initSerial : Exception  " + e + " line = " + e.lineNumber);
				onError();
				return null;
			}
		},
		deinitSerial: function () {
			try {
				logger.log("deinitSerial");
				this._outputStream = null;
				this._inputStream = null;
				if (this._listener) {
					logger.log("Removing serial listener");
					this._serial.removeListener(this._listener);
					this._listener = null;
				}
				if (this._serial != null) {
					this._serial.close();
					this._serial = null;
				}
				this._inited = false;
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		_reinit: function (aSysPath) {
			try {
				logger.log("_reinit serial");
				this.deinitSerial();
				this._initSerial(aSysPath);
				logger.log("_reinit OK\n");
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		sendCmd: function (aCmd) {
			if (!this._inited) {
				throw "Protocol not inited";
			}
			try {
				var count = 0;
				logger.log("sendCmd command " + aCmd);
				this._outputStream.write(aCmd, aCmd.length, count);
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		startListenProcess: function () {
			if (!this._inited) {
				throw "Protocol not inited";
			}
			try {
				logger.log("startListenProcess");
				this._listener.initController();
			}
			catch (e) {
				logger.error("Exception  " + e + " line = " + e.lineNumber);
				return null;
			}
		},
		reverseStringbyPair: function (aValue) {
			if ((aValue.length % 2) != 0) {
				logger.error("reverseStringbyPair has to be multiple of 2 in size");
				return null;
			}
			var array = Array.from(aValue);
			array.reverse();
			for (var i = 0; i < array.length; i += 2) {
				[array[i], array[i + 1]] = [array[i + 1], array[i]];
			}
			return array.join('').toString();
		}
	}


	window.SerialAdapterManager = SerialAdapterManager;

 })();