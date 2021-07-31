export function getServerSideData(count) {

    var data = []
    var os = require("os");
    var hostname = os.hostname();

    for (var i = 0; i < count; i++) {
        var d = new Date();
        data.push(
            {
                i,
                hostname,
                timestamp: getTimestamp(d),
            });
    }

    return data
}

function getTimestamp(d){
    return d.getDate() + '/' + (d.getMonth() + 1) + '/' + d.getFullYear() + ' ' + d.getHours() + ':' + d.getMinutes() + ':' + d.getSeconds()
}