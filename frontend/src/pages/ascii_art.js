export var ascii = {};

ascii.htmlButton = `
    <span class="button" id="pageAsciiButton" onclick="window.frontend.loadPage('pageAscii');">
        <i class="fas fa-pen-nib"></i> <span>ASCII art</span>
        <span class="tooltip">Create ASCII art from an input text</span>
    </span>
`;

ascii.htmlPage = `
    <div id="pageAscii" class="hidden">
        <input type="text" id="inputAscii" class="clear" placeholder="Plain text" onkeyup="window.frontend.ascii.perform();">
        <textarea id="outputAscii" class="clear" placeholder="ASCII art" readonly></textarea>
    </div>
`;

ascii.perform = function () {
    window.frontend.sendEventToBackend("ascii", "", [document.getElementById("inputAscii").value.trim()])
        .then(function (result) {
            document.getElementById("outputAscii").value = result.value;
        });
};