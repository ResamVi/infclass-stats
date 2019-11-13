<template>
  <section :style="'grid-area:' + area">
    <h1>{{ title }}</h1>
    <table style="margin: 0 auto;">
      <thead>
        <tr>
          <th>Name</th>
          <th>{{ column }}</th>
        </tr>
      </thead>

      <tbody>
        <Row
          v-for="(obj, index) in limitedList"
          :key="index"
          :name="obj.Key"
          :value="format(obj.Value) + ''"
          />
      </tbody>
    </table>
    <button
      @click="limit+=10"
      v-if="list.length > limit">
      Show more
    </button>
  </section>
</template>

<script>
import Row from '@/components/Row.vue';

export default {
  name: 'Table',

  components: {
    Row,
  },

  props: {
    title: String,
    area: String,
    column: String,
    list: Array,
    format: Function,
  },

  data() {
    return {
      limit: 5,
    };
  },

  computed: {
    limitedList() {
      return this.list.slice(0, this.limit);
    },
  },
};
</script>

<style scoped>
td {
  border-left: 1px solid #cbcbcb;
  border-width: 0 0 0 1px;
  font-size: inherit;
  margin: 0;
  overflow: visible;
  padding: 0.5em 1em;
  color: black;
}

button {
  background: none;
  border: none;
  cursor: pointer;
}

table {
  width: 75%;
}

</style>
