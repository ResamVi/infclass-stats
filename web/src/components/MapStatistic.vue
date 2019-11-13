<template>
  <div class="container">
    <img :src="getUrl" width="30%" height="30%">
    <div id="maps">
      <h1 style="grid-area:title">{{ mapName }}</h1>
      <h3 style="grid-area:humTitle">Human Wins</h3>
      <h3 style="grid-area:zombTitle">Zombie Wins</h3>
      <h3 style="grid-area:avgTitle">Average Survival</h3>
      <h4 style="grid-area:hum">{{ mapWin(mapName, 'Human') }}</h4>
      <h4 style="grid-area:zomb">{{ mapWin(mapName, 'Zombie') }}</h4>
      <h4 style="grid-area:avg">{{ mapAverageSurvival(mapName) }}</h4>
    </div>
  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex';

export default {
  name: 'MapStatistic',
  props: {
    mapName: String,
  },
  computed: {
    ...mapState([
      'wins',
    ]),
    ...mapGetters([
      'mapWin',
      'mapAverageSurvival',
    ]),
    getUrl() {
      return require(`@/assets/maps/${this.mapName}.png`);
    },
  },
};
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  justify-content: center;
  margin-bottom: 10px;
}

#maps {
  width:50%;
  height:100%;
  display: grid;
  grid-template-columns: 33% 33% 33%;
  grid-template-rows: auto;
  grid-template-areas:
    ".        title     ."
    "humTitle zombTitle avgTitle"
    "hum zomb avg";
}

h2, h3, h4, h5, h4 {
  margin: 2px 0;
}

</style>
