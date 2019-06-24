#  PDF viewer

The gekkota PDF viewer allows you to play pdf files in an HTML iframe element. Simply create an iframe and set its src attribute with the path to your pdf file.
You can also directly add it to your HTML:
````html
<iframe id="iframepdf" src="example.pdf"></iframe>
````
We use the messages from the iframe to retrieve information about the PDF document.
````javascript
window.addEventListener("message", function (event) {
  if (event.source === self._pdfIframeWindow) {
    var type;
    if (typeof event.data === "string") {
      type = event.data;
    } else {
      type = event.data.type;
    }
    if (event.data.detail) {
      console.log(type + " " + event.data.detail);
    }
  }
});
````
Messages you can get from the iframe are:
  *	numPages: detail contains the number of pages of the PDF document,
  *	pageDuration: detail contains the number of pages of the PDF document,
  *	pageChanged : sent jsut after pageDuration, its detail contains the number of the page current page.

Moreover it supports standards HTML object events plus "beginEvent", "endEvent" and "error".

Messages you can post to the iframe to interract with the PDF:
  * restart: restarts the PDF from the beginning,
  * gotonext: change current page the next one,
  * gotoprev: change current page the previous one,
  * gotopage: detail contains the page to go to.


Here is an example to go to page 7:
````javascript
var pdfIframeWindow = document.getElementById("iframepdf");
var pageNumber = 7;
var goToPageMessage = {
  type: "gotoPage",
  detail: pageNumber
};
pdfIframeWindow.postMessage(goToPageMessage, "*");
````
You can find an HTML example using this API [here.](example1.html)

**Build note**: You need to execute the **build.cmd** file to generate the boostrap app. Otherwise there will be a mismatch between the html file name and the one the manifest tries to launch. Find more information in *SDK-G4/bootstrap App/* documentation.
