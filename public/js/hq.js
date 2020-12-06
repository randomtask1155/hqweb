getHQStatus();
window.setInterval(function(){
   getHQStatus();
}, 60000);


function getHQStatus() {
    $.ajax({
        url: "/status", // todo paramertise
        type: 'get',
        dataType: 'json',
        beforeSend: function() {
        },
        success: function (data) {
            var roku = findDevice("Roku Players", data);
            var pump = findDevice("Sump Pump", data);
            updateRoku(roku);
            updateSumpPump(pump);
            if ((roku.status != "healthy") || (pump.status != "healthy" )) {
                $('#mainPageStatus').html('Status [Down]');
                $('#mainPageStatus').addClass("text-danger");
                $('#mainPageStatus').removeClass("text-success");
            } else {
                $('#mainPageStatus').html('Status [OK]');
                $('#mainPageStatus').removeClass("text-danger");
                $('#mainPageStatus').addClass("text-success");
            }
        },
        error: function(data) {
            console.log(data);
            alert(JSON.stringify(data));
        },
        complete: function () {

        }
    });
}

function findDevice(name, devices) {
    for ( i = 0; i < devices.length; i++) {
        if (devices[i].name == name ) {
            return devices[i];
        }
    }
}

function updateRoku(d){
    for ( i = 0; i < d.events.length; i++) {
        var rd = d.events[i].split(":");
        if ( rd.length < 3) {
            continue;
        }

        $('#rokuplayer' + i).html(rd[0]+'<span class="font-s16 font-w400 text-success animated infinite pulse pull-right">'+rd[2]+'</span>');
    }        
}

function updateSumpPump(d) {
    if ( d.status == "healthy") {
        $('#sumpPumpStatus').html('SumpPump [<span class="text-success">OK</span>]');
    } else {
        $('#sumpPumpStatus').html('SumpPump [<span class="text-danger">Down</span>]');
    }
}