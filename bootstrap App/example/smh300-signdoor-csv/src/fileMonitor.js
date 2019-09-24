(function () {

    const LOGO_FILE = "logo.png"; // Name of the app embedded logo

    // These index correspond to the data.csv file that goes with this app
    const OFFICE_INDEX_IN_ROW = 0;
    const NAME1_INDEX_IN_ROW = 1;
    const TITLE1_INDEX_IN_ROW = 2;
    const NAME2_INDEX_IN_ROW = 3;
    const TITLE2_INDEX_IN_ROW = 4;

    function FileMonitor(logger, fileToMonitor, monitorAddedFile, monitorFileModification) {
        this._monitor = undefined;
        this._logger = logger;
        this._monitorAddedFile = monitorAddedFile;
        this._monitorFileModification = monitorFileModification;
        this._screenshot = new Screenshot();
        this._fileToMonitor = fileToMonitor;
        // Preload logo div
        this._logoDiv = document.createElement("img");
        this._logoDiv.id = "logoDiv";
        this._logoDiv.src = LOGO_FILE;
    }
    FileMonitor.prototype = {
        load: function (rootPath) {
            try {
                this._logger.log("FileMonitor load: " + rootPath);
                this._monitor = new FileSystemWatcher({ "listener": this, "rootPath": rootPath, "recursive": false });
                this._monitor.startWatching();
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        onWatcherStarted: function () { },
        onWatcherStopped: function () { },
        onWatcherError: function (aErrorType, aDescription) { },
        onFileSystemChanged: function (aFilePath) {
            try {
                this._logger.log("FileMonitor onFileSystemChanged: " + aFilePath);
                // Make it case insensitive
                if (this._monitorFileModification && aFilePath.toLowerCase().includes(this._fileToMonitor.toLowerCase())) {
                    this._requestDownload(aFilePath);
                }
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        tryLoadFileAndGeneratePicture: function (aFilePath) {
            try {
                this._logger.log("FileMonitor generatePictureIfDataFile: " + aFilePath);
                // Download will fail if file doesn't exists
                this._requestDownload(aFilePath);
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        onFileSystemRemoved: function (aFilePath) {
            try {
                this._logger.log("FileMonitor onFileSystemRemoved: " + aFilePath);
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        onFileSystemAdded: function (aFilePath) {
            try {
                this._logger.log("FileMonitor onFileSystemAdded: " + aFilePath);
                // Make it case insensitive
                if (this._monitorAddedFile && aFilePath.toLowerCase().includes(this._fileToMonitor.toLowerCase())) {
                    this._requestDownload(aFilePath);
                }
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        _parseCSV: function (str) {
            try {
                this._logger.log("FileMonitor _parseCSV length: " + str.length);
                if (str.length == 0) {
                    this._logger.error("FileMonitor _parseCSV warning empty string.");
                    return;
                }
                this._logger.log("FileMonitor _parseCSV str: " + str);
                var arr = [];
                var quote = false;  // true means we're inside a quoted field

                // iterate over each character, keep track of current row and column (of the returned array)
                for (var row = 0, col = 0, c = 0; c < str.length; c++) {
                    var cc = str[c], nc = str[c + 1];        // current character, next character
                    arr[row] = arr[row] || [];             // create a new row if necessary
                    arr[row][col] = arr[row][col] || '';   // create a new column (start with empty string) if necessary

                    // If the current character is a quotation mark, and we're inside a
                    // quoted field, and the next character is also a quotation mark,
                    // add a quotation mark to the current column and skip the next character
                    if (cc == '"' && quote && nc == '"') { arr[row][col] += cc; ++c; continue; }

                    // If it's just one quotation mark, begin/end quoted field
                    if (cc == '"') { quote = !quote; continue; }

                    // If it's a comma and we're not in a quoted field, move on to the next column
                    if (cc == ',' && !quote) { ++col; continue; }

                    // If it's a newline (CRLF) and we're not in a quoted field, skip the next character
                    // and move on to the next row and move to column 0 of that new row
                    if (cc == '\r' && nc == '\n' && !quote) { ++row; col = 0; ++c; continue; }

                    // If it's a newline (LF or CR) and we're not in a quoted field,
                    // move on to the next row and move to column 0 of that new row
                    if (cc == '\n' && !quote) { ++row; col = 0; continue; }
                    if (cc == '\r' && !quote) { ++row; col = 0; continue; }

                    // Otherwise, append the current character to the current column
                    arr[row][col] += cc;
                }
                this._logger.log("FileMonitor _parseCSV arr: " + arr);
                // Start at i=1 because first row is header
                for (var i = 1; i < arr.length; i++) {
                    this._displayContactsInfo(arr[i]);
                    var slateFolderPath = "" + i + "" + "/"
                    this._generateSlatePicture(slateFolderPath + window.SLATE_PICTURE_NAME);
                }
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        _generateSlatePicture: function (aSlatePicturePath) {
            try {
                this._logger.log("FileMonitor _generateSlatePicture");
                // "image/g4" is the image type recognize by the Slate106
                this._screenshot.captureObject(window, aSlatePicturePath, "image/g4", 0);
            } catch (error) {
                this._logger.log(error);
            }
        },
        _displayContactsInfo: function (oneRowArray) {
            try {
                this._logger.log("FileMonitor _displayRowContacts: " + oneRowArray);
                var containerDiv = document.getElementById("container");

                // Rermove old container content
                if (!containerDiv) {
                    this._logger.error("Error no container DIV.");
                    return 0;
                }
                while (containerDiv.hasChildNodes()) {
                    containerDiv.removeChild(containerDiv.firstChild);
                }

                if (oneRowArray.length < 3) {
                    this._logger.error("One of the field of the needed csv file isn't fill. This is a simple example, fill at least first 3 fields for each row.");
                } else {
                    this._displayMainTopDiv(containerDiv, oneRowArray);
                    this._displayMainBottomDiv(containerDiv, oneRowArray);
                }
            } catch (error) {
                this._logger.log(error);
            }
        },
        _displayMainTopDiv: function (containerDiv, oneRowArray) {
            try {
                var mainTopDivContainer = document.createElement("div");
                var officeDiv = document.createElement("div");

                mainTopDivContainer.id = "mainTopDivContainer";
                containerDiv.appendChild(mainTopDivContainer);

                officeDiv.id = "officeDiv";
                officeDiv.textContent = oneRowArray[OFFICE_INDEX_IN_ROW];
                mainTopDivContainer.appendChild(officeDiv);

                mainTopDivContainer.appendChild(this._logoDiv);

            } catch (error) {
                this._logger.log(error);
            }
        },
        _displayMainBottomDiv: function (containerDiv, oneRowArray) {
            try {
                var mainBottomDivContainer = document.createElement("div");
                var name1Div = document.createElement("div");
                var title1Div = document.createElement("div");
                var name2Div = document.createElement("div");
                var title2Div = document.createElement("div");

                mainBottomDivContainer.id = "mainBottomDivContainer";
                containerDiv.appendChild(mainBottomDivContainer);

                name1Div.id = "name1Div";
                name1Div.textContent = oneRowArray[NAME1_INDEX_IN_ROW];
                mainBottomDivContainer.appendChild(name1Div);

                title1Div.id = "title1Div";
                title1Div.textContent = oneRowArray[TITLE1_INDEX_IN_ROW];
                mainBottomDivContainer.appendChild(title1Div);

                // Check if row has second name
                if (oneRowArray.length > 3) {
                    // Ensure that both name and title are set
                    if (oneRowArray.length == 4) {
                        this._logger.error("One of the field of the needed csv file isn't fill. This is a simple example, fill both name an title for each row.");
                    } else {
                        name2Div.id = "name2Div";
                        name2Div.textContent = oneRowArray[NAME2_INDEX_IN_ROW];
                        mainBottomDivContainer.appendChild(name2Div);

                        title2Div.id = "title2Div";
                        title2Div.textContent = oneRowArray[TITLE2_INDEX_IN_ROW];
                        mainBottomDivContainer.appendChild(title2Div);
                    }
                }
            } catch (error) {
                this._logger.log(error);
            }
        },
        _requestDownload: function (aFilePath) {
            try {
                this._logger.log("FileMonitor _requestDownload: " + aFilePath);
                var xhrParams = {};
                xhrParams.uri = aFilePath;
                xhrParams.isRemoteSource = false;
                this._xhrLoad(xhrParams);
            } catch (error) {
                this._logger.errorEx(error);
            }
        },
        _xhrLoad: function (xhrParams) {
            var xhr;
            var self = this;
            var timeoutXhrOnSend = 500;
            xhr = new XMLHttpRequest();

            xhr.onload = function (event) {
                try {
                    targetXhr = (event) ? event.target : xhr;
                    self._logger.log("XHR onload, readyState: " + targetXhr.readyState);
                    self._logger.log("XHR status : " + targetXhr.status);
                    self._parseCSV(targetXhr.response);
                } catch (err) {
                    self._logger.error("Ex xhr onload : " + err);
                }
            };

            xhr.onerror = function (event) {
                try {
                    var targetXhr;
                    targetXhr = (event) ? event.target : xhr;
                    self._logger.error("XHR error, xhr.readyState: " + targetXhr.readyState + " " + targetXhr.status);
                } catch (err) {
                    self._logger.error("Ex xhr onerror : " + err);
                }
            };

            xhr.open('GET', xhrParams.uri, false);
            setTimeout(function (xhrParams) {
                xhr.send();
            }, timeoutXhrOnSend);
        }
    };

    window.FileMonitor = FileMonitor;
})();