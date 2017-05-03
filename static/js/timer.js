var time = {hours: 0, minutes: 0, seconds: 0};

$(document).ready(function() {

  $('#startTimer').on('click', function() {

    var startDate = new Date().getTime();

    // TODO vlozit startDate do databaze prislusnemu uzivateli. (Kdy zapocal
    // praci)

    $('#counter').val('00:00:00');
    startCounter(startDate);

  });

  $('#stopTimer').on('click', function() {

    $('#stopTimer').prop('disabled', true);
    $('#startTimer').prop('disabled', false);
    $('#resetTimer').prop('disabled', false);

  });

  $('#resetTimer').on('click', function() {

    var task = $('#selectTasks').find(':selected').text();

    alert(
        time.hours + ' hodin ' + time.minutes + ' minut ' + time.seconds +
        ' sekund ' +
        'k ukolu: ' + task);

    // TODO vlozit do databaze hours minutes seconds task (Time recond, jak
    // dlouho a na cem se pracovalo)

    $('#counter').val('00:00:00');
  });
});

// TODO vyndat startDate z databaze. var startDate = SELECT startDate FROM users
// WHERE ID =?, UUID

// -------------------------Ukazka
var startDate =
    new Date('Apr 27, 2017 14:50:00').getTime();  // Pokud bude startDate
                                                  // nastaven - v databazi bude
                                                  // ulozen cas
var startDate = 0;                                // Pokud nebude nastaveny cas
// -------------------------Ukazka

startCounter(startDate);

function startCounter(startDate) {
  var timer = null;

  if (startDate > 0) {
    document.getElementById('stopTimer').disabled = false;
    document.getElementById('startTimer').disabled = true;
    document.getElementById('resetTimer').disabled = true;

    timer = setInterval(function() {

      if (document.getElementById('stopTimer').hasAttribute('disabled')) {
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