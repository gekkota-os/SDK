(function(global) {
	"use strict";
	function getLocationInfoFromException(e, level) {
		if(!e.stack) {
			return null;
		}
		var lines = e.stack.split("\n");
		var reg = new RegExp("^(.*)@(.*):(.*)", "g");

		var line = lines[level];
		var file = line.replace(reg, "$2");
		var func = line.replace(reg, "$1");
		var line = line.replace(reg, "$3");
		return log4Service.createLocationInfo(file, func, line);
	}

	function getLocationInfo() {
		try {
			throw new Error();
		} catch(e) {
			return getLocationInfoFromException(e, 2);
		}
	}

	function getMessage() {
		let msg = "";
		for(let i = 0; i < arguments.length; ++i) {
			msg += arguments[i];
		}
		return msg;
	}
	
	function setProperty(obj, name, prop) {
		Object.defineProperty(obj, name, {
			configurable: false,
			enumerable: false,
			writable: false,
			value: prop
		});
	}

	function XpfLogger(name) {
		setProperty(this, "_logger", log4Service.getLogger(name));
		setProperty(this, "_debugEnabled", this._logger.isDebugEnabled());
		setProperty(this, "_warnEnabled", this._logger.isWarnEnabled());
		setProperty(this, "_errorEnabled", this._logger.isErrorEnabled());
		setProperty(this, "_fatalEnabled", this._logger.isFatalEnabled());
	}

	XpfLogger.prototype = {
		constructor: XpfLogger,

		//debug
		isDebugEnabled: function isDebugEnabled() {
			return this._debugEnabled;
		},
		get debugEnabled() {
			return this._debugEnabled;
		},
		debug: function debug() {
			if(this._debugEnabled) {
				this._logger.debugAS(getMessage.apply(null, arguments), getLocationInfo());
			}
		},
		debugEx: function debugEx(e){
			if(this._debugEnabled) {
				this._logger.debugAS(e, getLocationInfoFromException(e, 0));
			}
		},

		//warn
		isWarnEnabled: function isWarnEnabled() {
			return this._warnEnabled;
		},
		get warnEnabled() {
			return this._warnEnabled;
		},
		warn: function warn() {
			if(this._warnEnabled) {
				this._logger.warnAS(getMessage.apply(null, arguments), getLocationInfo());
			}
		},
		warnEx: function warnEx(e){
			if(this._warnEnabled) {
				this._logger.warnAS(e, getLocationInfoFromException(e, 0));
			}
		},

		//error
		isErrorEnabled: function isErrorEnabled() {
			return this._errorEnabled;
		},
		get errorEnabled() {
			return this._errorEnabled;
		},
		error: function error() {
			if(this._errorEnabled) {
				this._logger.errorAS(getMessage.apply(null, arguments), getLocationInfo());
			}
		},
		errorEx: function errorEx(e){
			if(this._errorEnabled) {
				this._logger.errorAS(e, getLocationInfoFromException(e, 0));
			}
		},

		//fatal
		isFatalEnabled: function isFatalEnabled() {
			return this._fatalEnabled;
		},
		get fatalEnabled() {
			return this._fatalEnabled;
		},
		fatal: function fatal() {
			if(this._fatalEnabled) {
				this._logger.fatalAS(getMessage.apply(null, arguments), getLocationInfo());
			}
		},
		fatalEx: function fatalEx(e){
			if(this._fatalEnabled) {
				this._logger.fatalAS(e, getLocationInfoFromException(e, 0));
			}
		}
	};

	Object.freeze(XpfLogger.prototype);
	Object.freeze(XpfLogger);

	Object.defineProperty(global, "XpfLogger", {
		configurable: false,
		writable: false,
		enumerable: false,
		value: XpfLogger
	});
})(window)