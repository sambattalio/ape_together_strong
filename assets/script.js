// Script to Display change in weights for an exercise over time

var weight = [15, 25, 35];
var dates = ['03/01', '03/08', '03/15'];

var ctx = document.getElementById('exerciseChangeChart');

var exerciseChangeChart = new Chart(ctx, {
    type: 'line',
    data: {
	labels: dates,
	datasets: [{
	    label: 'Exercise Weights over Time',
	    data: weight,
	    backgroundColor: "rgba(255, 99, 132, 0.2)",
            borderColor: "rgba(255, 99, 132, 1)",
            borderWidth: 1,
	    fill: false,
	    lineTension: 0
	}]
    },
    
    options: {
		responsive: false
		//maintainAspectRatio: false
		//beginAtZero: true
    }
});
