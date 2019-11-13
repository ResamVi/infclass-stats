<template>
  <div class="content">
    <h1>{{ $route.params.name }}</h1>
    Kills: {{ profileKills($route.params.name) }} <br/>
    Time: {{ profileTime($route.params.name) }}
    <DonutChart :series="chartPlayerClass($route.params.name)" />
  </div>
</template>

<script>
import DonutChart from '@/components/DonutChart.vue';
import { mapGetters } from 'vuex';

export default {
  name: 'Profile',

  components: {
    DonutChart,
  },
  computed: {
    ...mapGetters([
      'profileTime',
      'profileKills',
      'chartPlayerClass',
    ]),
  },
};
</script>


<style scoped>
h1 {
  color: #666666;
}

.content {
    border: solid 1px #ddd;
    border-radius: 6px;
    display: flex;
    justify-content: center;
    flex-direction: column;
    background-color: #fff;
    width: 80%;
    margin: 0 auto;
}

</style>

