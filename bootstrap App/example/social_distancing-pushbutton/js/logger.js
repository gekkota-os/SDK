(function () {
  function Logger(name) {
    this._logService = window.log4Service ? window.log4Service.getLogger(name) : (window.console ? window.console : null);
  }
  Logger.prototype = {
    log: function (message) {
      this._log(message);
    },
    error: function (message) {
      this._log(message, true);
    },
    errorEx: function (exception) {
      this._log("Exception line: " + exception.lineNumber + ", value: " + exception.message, true);
    },
    _log: function (message, isError) {
      if (isError) {
        if (this._logService) {
          this._logService.error(message, null);
        }
      } else {
        if (this._logService) {
          this._logService.debug(message, null);
        }
      }
    }
  };
  window.Logger = Logger;
})();