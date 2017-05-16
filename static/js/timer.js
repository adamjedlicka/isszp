$.getJSON("/api/timer", (data) => {
    if (data.length == 0) {
        $("#TimerStart").show();
    } else {
        $("#TimerStop").show();

        $("#TimerDescription").prop("disabled", true);
        $("#TimerTask").prop("disabled", true);

        $("#TimerDescription").val(data[0].Description);
        $("#TimerTask").val(data[0].TaskID);

        let start = new Date(data[0].Date + " " + data[0].Start);

        SetCounter(start);
        window.setInterval(() => {
            SetCounter(start);
        }, 1000);
    }
});

$("#TimerStart").click(() => {
    $.post("/api/timer/start", {
        TaskID: $("#TimerTask").val(),
        Description: $("#TimerDescription").val(),
    }).done(() => {
        window.location.reload();
    });
});

$("#TimerStop").click(() => {
    $.post("/api/timer/stop", () => {
        window.location.reload();
    })
});

function SetCounter(from) {
    let now = new Date();
    let diff = new Date(now - from);

    $("#TimerCounter").val(diff.toISOString().substr(11, 8));
}
