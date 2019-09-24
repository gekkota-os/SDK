(function () {

    const GET_WPAN_FILE_INFO_INTERVAL = 60000; // Interval between two calls to the function
    const NUMBER_OF_SLATES = 2; // If this value is N, it implies that the Id of the Slates are 1,2...N.

    function WpanFileInfo(logger) {
        this._logger = logger;
        this._logger.log("WpanFileInfo");
        this._getWpanFileInfoInterval = GET_WPAN_FILE_INFO_INTERVAL;
        this._nbSlates = NUMBER_OF_SLATES;
        this._wpanHubSrv = window.wpanHubSrv;
        this._intervalID = null;
    }
    WpanFileInfo.prototype = {
        init: function () {
            try {
                this._logger.log("WpanFileInfo init");
                this._getSlatesFileInfos();
                this._intervalID = window.setInterval(this._getSlatesFileInfos.bind(this), this._getWpanFileInfoInterval);
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        _getSlatesFileInfos: function () {
            try {
                this._logger.log("WpanFileInfo _getSlatesFileInfos");
                for (var i = 0; i < this._nbSlates; i++) {
                    if (this._wpanHubSrv.getFileInfo) {
                        var sBaseUri = this._wpanHubSrv.BASEURI_SLATE_PICTURE;
                        var sFilePath = (i + 1) + "/" + window.SLATE_PICTURE_NAME;
                        var isFileInSync = {};
                        var lastUpdateTime = {};
                        this._wpanHubSrv.getFileInfo(sBaseUri, sFilePath, isFileInSync, lastUpdateTime);
                        this._logger.log(
                            " filePath: " + sFilePath +
                            " isFileInSync: " + JSON.stringify(isFileInSync) +
                            " lastUpdateTime: " + JSON.stringify(lastUpdateTime));
                    }
                }
            } catch (error) {
                this._logger.errorEx(error);
            }
        }
    };

    window.WpanFileInfo = WpanFileInfo;
})();