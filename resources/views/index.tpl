=====
Pagetitle: Dashboard
BodyClasses: layout-dashboard
=====
<h1 class="pt-3">Welcome!</h1>

<p>Server Statistics</p>

<div class="row">
    <div class="col-12 col-md-4">
        <canvas id="loadChart" width="400" height="400"></canvas>
    </div>
    <div class="col-12 col-md-4">
        <canvas id="ramChart" width="400" height="400"></canvas>
    </div>
    <div class="col-12 col-md-4">
        <canvas id="discChart" width="400" height="400"></canvas>
    </div>
</div>
<script>
var intervalSeconds = 5;

function getStats() {
    axios.get("stats")
        .then(function(resp) {
            console.log(resp);

            // if we don't get a 200 then our session is likey expired
            if (resp.headers['content-type'] !== "application/json") {
                location.reload();
            }

            // for easy peeking in the console
            window.statsData = resp.data;

            // update the bar chart
            loadChart.chart.data.datasets[0].data = [
                window.statsData.loadAverages['one-minute'],
                window.statsData.loadAverages['five-minutes'],
                window.statsData.loadAverages['fifteen-minutes']
            ];
            if (window.statsData.loadAverages['one-minute'] > 0.5) {
                loadChart.chart.data.datasets[0].backgroundColor[0] = "orange";
            }
            if (window.statsData.loadAverages['five-minutes'] > 0.5) {
                loadChart.chart.data.datasets[0].backgroundColor[1] = "orange";
            }
            if (window.statsData.loadAverages['fifteen-minutes'] > 0.5) {
                loadChart.chart.data.datasets[0].backgroundColor[2] = "orange";
            }
            if (window.statsData.loadAverages['one-minute'] > 0.8) {
                loadChart.chart.data.datasets[0].backgroundColor[0] = "red";
            }
            if (window.statsData.loadAverages['five-minutes'] > 0.8) {
                loadChart.chart.data.datasets[0].backgroundColor[1] = "red";
            }
            if (window.statsData.loadAverages['fifteen-minutes'] > 0.8) {
                loadChart.chart.data.datasets[0].backgroundColor[2] = "red";
            }
            loadChart.update();

            ramChart.chart.data.datasets[0].data = [
                window.statsData.ram.mem.total - window.statsData.ram.mem.used,
                window.statsData.ram.mem.used
            ];
            ramChart.update();
        });
}
getStats();

setInterval(function() {
    getStats();
}, intervalSeconds * 1000);

var loadCtx = document.getElementById('loadChart').getContext('2d');
var loadChart = new Chart(loadCtx, {
    type: 'bar',
    data: {
        labels: ['1 Minute', '5 Minutes', '15 Minutes'],
        datasets: [{
            label: 'Load Averages',
            data: [0, 0, 0],
            borderWidth: 1,
            backgroundColor: ["grey", "grey", "grey"]
        }]
    },
    options: {
        scales: {
            yAxes: [{
                ticks: {
                    beginAtZero: true,
                    max: 1
                }
            }]
        }
    }
});

var ramCtx = document.getElementById('ramChart').getContext('2d');
var ramChart = new Chart(ramCtx, {
    type: 'doughnut',
    data: {
        labels: ['RAM Used/Buffered', 'RAM Available'],
        datasets: [{
            label: 'RAM Usage',
            data: [0, 100],
            borderWidth: 1,
            backgroundColor: ["red", "lightGrey"]
        }]
    }
});
</script>
