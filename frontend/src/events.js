export function sendEventToBackend(topic, action, values) {
    var event = {
        source: topic,
        action: action,
        values: values.map(function (value) {
            return value.replace(/\u00A0/g, "");
        })
    };

    document.body.classList.add("waiting");

    return window.backend.Bus.SendEvent(JSON.stringify(event))
        .catch(function (error) {
            document.body.classList.remove("waiting");
            document.getElementById("info").classList.add("hidden");
            document.getElementById("error").classList.remove("hidden");
            document.getElementById("footerText").innerText = error;
            throw new Error(error);
        })
        .then(function (returnedEventJson) {
            document.body.classList.remove("waiting");
            var returnedEvent = JSON.parse(returnedEventJson);

            document.getElementById("info").classList.add("hidden");
            document.getElementById("error").classList.add("hidden");
            document.getElementById("footerText").innerText = "";

            if (returnedEvent.info.length > 0) {
                document.getElementById("info").classList.remove("hidden");
                document.getElementById("footerText").innerText = returnedEvent.info;
            }
            return returnedEvent;
        });
}

export function receiveEventFromBackend(returnedEventJson) {
    var returnedEvent = JSON.parse(returnedEventJson);

    if (returnedEvent.info.length > 0) {
        document.getElementById("error").classList.add("hidden");
        document.getElementById("info").classList.remove("hidden");
        document.getElementById("footerText").innerText = returnedEvent.info;
    }

    return returnedEvent.value;
}