$(document).ready(() => {
    window.setInterval(() => {
        $.getJSON("/api/notify", (data) => {
            if(data.length > 0) {
                if (Notification.permission !== "granted") {
                     Notification.requestPermission();
                }

                data.forEach((v, k) => {
                    var notification = new Notification(v.Header, {
                        icon: "/static/img/notify.png",
                        body: v.Message,
                    });

                    if (v.Address != "") {
                        notification.onclick = () => {
                            window.location.replace(v.Address);
                            notification.close();
                        }
                    }
                });
            }
        });
    }, 1000);
});
