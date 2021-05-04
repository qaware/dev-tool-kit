export var xml = {};

xml.htmlButton = `
    <span class="button" id="pageXmlButton" onclick="window.frontend.loadPage('pageXml');">
        <i class="fas fa-code"></i> <span>XML</span>
        <span class="tooltip">Pretty-print an XML document</span>
    </span>
`;

xml.htmlPage = `
    <div id="pageXml" class="hidden">
        <textarea id="inputXml" class="clear autodetect-input" placeholder="Raw XML"
            onfocusout="window.frontend.xml.format();"
            onpaste="window.frontend.xml.formatPasted();"></textarea>
        <textarea id="outputXml" class="clear" placeholder="Pretty-printed XML" readonly></textarea>
    </div>
`;

xml.format = function () {
    window.frontend.sendEventToBackend("xml", "", [document.getElementById("inputXml").value])
        .then(function (result) {
            document.getElementById("outputXml").value = result.value;
        });
};

xml.formatPasted = function () {
    setTimeout(function () {
        window.frontend.xml.format();
    }, 10);
};