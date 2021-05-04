export var diff = {};

diff.htmlButton = `
    <span class="button" id="pageDiffButton" onclick="window.frontend.loadPage('pageDiff');">
        <i class="fas fa-exchange-alt"></i> <span>Diff</span>
        <span class="tooltip">Compare two texts</span>
    </span>
`;

diff.htmlPage = `
    <div id="pageDiff" class="hidden">
        <textarea id="leftDiff" class="diff-left clear" placeholder="Left text" onkeyup="window.frontend.diff.compare();"></textarea>
        <div class="diff-container">
            <div id="outputDiffContainer">
                <div id="outputDiff" class="diff clear"></div>
            </div>
            <div class="row" style="flex-shrink: 0;">
                <span class="button" onclick="window.frontend.diff.scroll();">
                    <i class="fas fa-angle-double-down"></i> Scroll to next difference
                </span>
            </div>
        </div>
        <textarea id="rightDiff" class="diff-right clear" placeholder="Right text" onkeyup="window.frontend.diff.compare();"></textarea>
    </div>
`;

diff.compare = function () {
    var left = document.getElementById("leftDiff").value;
    var right = document.getElementById("rightDiff").value;

    window.frontend.sendEventToBackend("diff", "compare", [left, right])
        .then(function (result) {
            document.getElementById("outputDiff").innerHTML = result.value;
        });
};

diff.scroll = function () {
    var diff = document.getElementById("outputDiff").innerHTML.length.toString();

    window.frontend.sendEventToBackend("diff", "scroll", [diff])
        .then(function (result) {
            if (result.value.length > 0) {
                document.getElementById(result.value).scrollIntoView({
                    behavior: "smooth"
                });
            }
        });
};