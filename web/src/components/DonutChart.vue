<template>
   <article>
        <apexchart width="600" type="donut" :options="chartOptions" :series="series"></apexchart>
   </article>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'DonutChart',
  props: {
    series: Array,
  },
  data() {
    return {
      chartOptions: {
        chart: {
          id: 'class-chart',
          toolbar: {
            show: false,
          },
        },
        labels: ['Mercenary', 'Medic', 'Hero', 'Engineer', 'Soldier', 'Ninja', 'Sniper', 'Scientist', 'Biologist', 'Looper'],
        grid: {
          xaxis: {
            lines: {
              show: true,
            },
          },
          yaxis: {
            lines: {
              show: true,
            },
          },
        },
      },
    };
  },
};
</script>

<style scoped>
article {
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
}

</style>

