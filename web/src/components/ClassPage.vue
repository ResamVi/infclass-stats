<template>
  <article class="container">
    <span>{{ title }}</span>

    <Table
      :title="'Class picked in total: ' + classTotalPicked(className)"
      :area="'picks'"
      :column="'# Picked'"
      :list="classPicks(className)"
      :format="identity"
    />
    <Table
      :title="'Score/Pick Ratio'"
      :area="'ratio'"
      :column="'K/P'"
      :list="classRatio(className)"
      :format="round"
    />
    <Table
      :title="'Class kills in total: ' + classTotalKills(className)"
      :area="'kills'"
      :column="'# Killed'"
      :list="classKills(className)"
      :format="identity"
    />
    <Table
      :title="'Highest amount of score in one round'"
      :area="'record'"
      :column="'# Killed'"
      :list="classRecord(className)"
      :format="identity"
    />
  </article>
</template>

<script>
import Row from "@/components/Row.vue";
import Table from "@/components/Table.vue";
import { mapGetters } from "vuex";

export default {
  name: "ClassPage",
  components: {
    Row,
    Table
  },
  props: {
    className: String
  },
  data() {
    return {
      round: a => Math.round(a * 100) / 100,
      identity: a => a,
      limit: 10
    };
  },
  computed: {
    title() {
      return this.className.charAt(0).toUpperCase() + this.className.slice(1);
    },
    ...mapGetters([
      "classPicks",
      "classKills",
      "classRatio",
      "classRecord",
      "classTotalPicked",
      "classTotalKills"
    ])
  }
};
</script>

<style scoped>
.container {
  display: grid;
  grid-template-columns: 25% 25% 25% 25%;
  grid-template-rows: auto;
  grid-template-areas:
    "title title title title"
    "picks ratio kills record";
}
button {
  background: none;
  border: none;
  cursor: pointer;
}

table {
  margin: 0 auto;
}

.center {
  display: flex;
  justify-content: center;
  flex-direction: column;
}

span {
  grid-area: title;
  font-size: 2em;
  color: #666666;
}
</style>

