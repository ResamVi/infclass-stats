<template>
   <article>
        <apexchart width="1000" type="bar" :options="chartOptions" :series="series"></apexchart>
   </article>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'SurvivalChart',
  props: {
    series: Array,
  },
  data() {
    return {
      chartOptions: {
        chart: {
          id: 'radial-chart',
        },
        plotOptions: {
          bar: {
            horizontal: true,
          }
        },
        toolbar: {
          show: true,
        },
        dataLabels: {
          formatter: function (val, opts) {
            let seconds = Math.floor(val);
            const minutes = Math.floor(seconds / 60);
            seconds -= minutes * 60;
            return `${minutes}:${(`00${seconds}`).substr(-2, 2)}`;
          },
        },
        labels: ['Mercenary', 'Medic', 'Hero', 'Engineer', 'Soldier', 'Ninja', 'Sniper', 'Scientist', 'Biologist', 'Looper'],
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

