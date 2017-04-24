function formatDate(date) {
	let d = new Date(date);
	let month = '' + (d.getMonth() + 1);
	let day = '' + d.getDate();
	let year = d.getFullYear();

	if (month.length < 2) month = '0' + month;
	if (day.length < 2) day = '0' + day;

	return [year, month, day].join('-');
}

$('body').ready(() => {
	// Date pickers with class .date-now will have current day as default value
	if ($('.date-now').val() == '') {
		$('.date-now').val(formatDate(new Date()));
	}
});