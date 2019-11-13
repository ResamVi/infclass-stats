<template>
  <section :style="'grid-area:' + area">
    <h1 style="font-size: 1.5em"> {{ title }} </h1>
    <img :src="'https://stats.resamvi.io/classes/' + title + '.png'">
    
    <div v-for="(obj, index) in list" :key="index">
      <div :style="'margin: 20px;' + 'font-size: ' + (1.3 - 0.1*index) + 'em; display: flex; align-items: center; justify-content: center;'">
        <img :src="medalPicture(index)">
        <span>{{ obj.Key }} </span>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'List',

  props: {
    title: String,
    area: String,
    list: Array,
  },

  methods: {
    medalPicture(index) {
      return require(`@/assets/medal${index}.png`);
    }
  },
};
</script>

<style scoped>
section {
  box-shadow: 0 0 10px rgba(0, 0, 0, .05);
  background-color: #fff;
  border: solid 1px #ddd;
  margin: 10px;

}
</style>
