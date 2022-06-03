<template>
  <div>
    <h1>ACTIVITY OVER 7 DAYS</h1>
    <apexchart width='600' type='area' :options='chartOptions' :series='chartActivity'></apexchart>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'ActivityChart',
  data() {
    return {
      chartOptions: {
        chart: {
          id: 'activity-chart',
          toolbar: {
            show: false
          }
        },
        stroke: {
          curve: 'smooth'
        },
        dataLabels: {
          enabled: false
        },
        tooltip: {
          enabled: true,
          x: {
            show: false,
            format: 'dd MMM',
            formatter(timestamp) {
              const date = new Date(timestamp * 1000);
              const day = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'][date.getDay()];
              const hours = `0${date.getHours()}`;
              const minutes = `0${date.getMinutes()}`;
              const formattedTime = `${day}, ${hours.substr(-2)}:${minutes.substr(-2)}`;
              return formattedTime;
            }
          }
        },
        grid: {
          xaxis: {
            lines: {
              show: true
            }
          },
        },
        xaxis: {
          labels: {
            formatter(value, timestamp) {
              const date = new Date(timestamp * 1000);
              const day = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'][date.getDay()];
              const hours = `0${date.getHours()}`;
              const minutes = `0${date.getMinutes()}`;
              const formattedTime = `${day}, ${hours.substr(-2)}:${minutes.substr(-2)}`;
              return formattedTime;
            }
          }
        }
      }
    };
  },
  computed: {
    ...mapGetters(['chartActivity'])
  }
};
</script>

<style scoped>
</style>

