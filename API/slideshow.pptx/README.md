#  PDF viewer

The gekkota PPTX viewer allows you to play PPTX files in an HTML div element. in order to create the PPTX viewer, you need to load its script:

````html
<script src="resource://app/res/pptx-viewer.js"></script>
````

Then you can create the PPTX viewer in JavaScript:
````javascript
var pptxFile = "example.pptx";
var pptxDiv = document.getElementById("pptxDiv");
var pptxViewer = new PptxViewer(
  function ended() {
    // PPTX ended
  },
  function error(e) {
    // PPTX error
  },
  function pageChanged() {
    // PPTX page changed
  }
);
pptxViewer.init(pptxDiv, pptxFile);
pptxViewer.start();
````

The PPTX overwrite functions allow you to catch informations about the document state and interact accordingly. Once created, init the PPTX viewer with the div where the viewer will play the corresponding file.

Here is the PPTX viewer API to interract with your document:

* start(): starts the viewer, displays first slide and goes on if a duration has been set and autoPageProgress is activated. Otherwise you will need to call specific functions to change pages,
* autoPageProgress: a boolean attribute of the viewer that sets auto progress. If set to true and if the PPTX document has a duration, the slides will go on automatically,
* close(): closes the viewer,
* nextPage(): viewer plays next slide,
* previousPage(): viewer plays previous slide,
* goToPage(pageNumber): viewer plays slide of number pageNumber,
* pageDuration: an integer attribute of the viewer that gives the duration of the page, null of none is set,
* numPages: an integer attribute of the viewer that gives the number of pages of the document.

You can find an HTML example using this API [here.](example1.html)