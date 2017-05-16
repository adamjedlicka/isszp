var time = {hours: 0, minutes: 0, seconds: 0};

$(document).ready(function() {

  $('#startTimer').on('click', function() {

    var startDate = new Date().getTime();
    var taskID = $('#selectTasks').find(':selected').attr('id');

    // Send data to backend - manipulation with database
    $.ajax({
      url: '/api/startTimer',
      type: 'POST',
      data: {
        'taskID': taskID,
        'startDate': startDate,
        'description': $('#descriptionTR').val()
      },

      success: function() {
        $('#counter').val('00:00:00');
        startCounter(startDate);
      }
    });

  });

  $('#stopTimer').on('click', function() {

    $('#stopTimer').prop('disabled', true);
    $('#stopTimer').hide();
    $('#startTimer').show();
    $('#selectTasks').prop('disabled', false);

    var task = $('#selectTasks').find(':selected').text();

    // Give a signal to backend, so timer was stopped
    $.ajax({
      url: '/api/stopTimer',
      type: 'POST',
      data: {'description': $('#descriptionTR').val()},

      success: function() {
        $('#counter').val('00:00:00');
      }
    });

  });

});

if (document.getElementById('counter') != null) {
  startDate = document.getElementById('counter').getAttribute(
      'startTime');  // If there is an open record in a database, startTime will
                     // be set
} else {
  startDate = '';
}

startCounter(startDate);

function startCounter(startDate) {
  var timer = null;

  if (startDate > 0) {
    document.getElementById('stopTimer').disabled = false;
    document.getElementById('stopTimer').style.display = 'initial';
    document.getElementById('startTimer').style.display = 'none';
    document.getElementById('selectTasks').disabled = true;

    timer = setInterval(function() {

      if (document.getElementById('stopTimer')
              .hasAttribute('disabled')) {  // If button stop is disabled -
                                            // timer is not running, overpass
                                            // calculations
        clearInterval(timer);
        startDate = 0;

        return;
        }

      var now = new Date().getTime();
      var diference = now - startDate;

      time.hours =
          Math.floor((diference % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
      time.minutes = Math.floor((diference % (1000 * 60 * 60)) / (1000 * 60));
      time.seconds = Math.floor((diference % (1000 * 60)) / 1000);

      if (time.hours < 10) {
        time.hours = ('0' + time.hours);
        }

      if (time.minutes < 10) {
        time.minutes = ('0' + time.minutes);
        }

      if (time.seconds < 10) {
        time.seconds = ('0' + time.seconds);
      }

      document.getElementById('counter').value =
          time.hours + ':' + time.minutes + ':' + time.seconds;

    }, 1000);
  }
}