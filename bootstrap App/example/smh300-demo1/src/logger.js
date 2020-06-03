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
    clearConsole: function () {
      this._getHtmlConsole().innerHTML = "";
    },
    _log: function (message, isError) {
      var elConsole = this._getHtmlConsole();
      if (isError) {
        if (this._logService) {
          this._logService.error(message, null);
        }
        message = "Error: " + message;
      } else {
        if (this._logService) {
          this._logService.debug(message, null);
        }
      }
      if (elConsole.scrollTop >= elConsole.clientHeight) {
        this.clearConsole();
      }
      elConsole.innerHTML += message + "\n";
      elConsole.scrollTop = elConsole.scrollHeight;
    },
    _getHtmlConsole: function () {
      return document.getElementById("console");
    },
  };
  window.Logger = Logger;
})();