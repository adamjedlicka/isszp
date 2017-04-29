var stopCounter = false;
var hours;
var minutes;
var seconds;

$(document).ready(function() {
  $('#startTimer').on('click', function() {

    var startDate = new Date().getTime();

    // TODO vlozit startDate do databaze prislusnemu uzivateli.

    startCounter(startDate);

  });

  $('#stopTimer').on('click', function() {

    stopCounter = true;

    alert(hours + ' hodin ' + minutes + ' minut ' + seconds + ' sekund');

    // TODO vlozit do databaze hours minutes seconds

  });

  $('#resetTimer').on('click', function() {
    $('#counter').text('00:00:00');
  });
});

// TODO vyndat startDate z databaze. var startDate = SELECT startDate FROM users
// WHERE ID =?, UUID

// Ukazka
var startDate =
    new Date('Apr 27, 2017 14:50:00').getTime();  // Pokud bude startDate
                                                  // nastaven - v databazi bude
                                                  // ulozen cas
var startDate = 0;                                // Pokud nebude nastaveny cas

startCounter(startDate);

function startCounter(startDate) {
  var timer = null;

  if (startDate > 0) {
    document.getElementById('stopTimer').disabled = false;
    document.getElementById('startTimer').disabled = true;
    document.getElementById('resetTimer').disabled = true;

    timer = setInterval(function() {

      if (stopCounter) {
        clearInterval(timer);
        startDate = 0;

        document.getElementById('stopTimer').disabled = true;
        document.getElementById('startTimer').disabled = false;
        document.getElementById('resetTimer').disabled = false;

        stopCounter = false;
        return;
        }

      var now = new Date().getTime();
      var diference = now - startDate;

      hours =
          Math.floor((diference % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
      minutes = Math.floor((diference % (1000 * 60 * 60)) / (1000 * 60));
      seconds = Math.floor((diference % (1000 * 60)) / 1000);

      if (hours < 10) {
        hours = ('0' + hours);
        }

      if (minutes < 10) {
        minutes = ('0' + minutes);
        }

      if (seconds < 10) {
        seconds = ('0' + seconds);
      }

      document.getElementById('counter').innerHTML =
          hours + ':' + minutes + ':' + seconds;

    }, 1000);
  }
}