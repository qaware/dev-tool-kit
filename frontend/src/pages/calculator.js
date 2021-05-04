export var calculator = {};

calculator.htmlButton = `
    <span class="button" id="pageCalculatorButton" onclick="window.frontend.loadPage('pageCalculator');">
        <i class="fas fa-calculator"></i> <span>Calculator</span>
        <span class="tooltip">Calculate the result of an algebraic expression</span>
    </span>
`;

calculator.htmlPage = `
    <div id="pageCalculator" class="hidden">
        <input type="text" id="inputCalculator" class="clear" placeholder="Arithmetic expression"
            onkeyup="window.frontend.calculator.calculate(event);">
        <textarea class="soft-wrap" id="outputCalculator" placeholder="Results" readonly></textarea>
        <div class="row">
            <span class="button" onclick="window.frontend.calculator.example();">
                <i class="fas fa-question-circle"></i> Example
            </span>
        </div>
    </div>
`;

calculator.calculate = function (event) {
    if (event.key !== "Enter") {
        return;
    }

    window.frontend.sendEventToBackend("calculator", "calculate", [document.getElementById("inputCalculator").value.trim()])
        .then(function (result) {
            var output = document.getElementById("outputCalculator");
            if (output.value.length > 0) {
                output.value = output.value + "\n" + result.value;
            } else {
                output.value = result.value;
            }
            document.getElementById("inputCalculator").value = "";
        });
};

calculator.example = function () {
    window.frontend.sendEventToBackend("calculator", "example", [])
        .then(function (result) {
            document.getElementById("inputCalculator").value = result.value;
        });
};