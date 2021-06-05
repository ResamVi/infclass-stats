<template>
  <div id="app">
    <Navigationbar />
    <img alt="logo" src="@/assets/bg.png" class="responsive">
    <router-view/>
  </div>
</template>

<script>
import Navigationbar from '@/components/Navigationbar.vue';

export default {
  name: 'App',

  components: {
    Navigationbar,
  },

  beforeMount() {
    const ws = new WebSocket(process.env.VUE_APP_API_URL);
    ws.onmessage = ({ data }) => {
      this.$store.commit('update', JSON.parse(data));
    };
  },
};
</script>

<!-- Global styles -->
<style>
#app {
  font-family: 'Open Sans', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}

body {
  margin: 0;
  background-color: #f2f2f2;
}

h1, h2, h3, h4, h5, h6 {
  color: #666666;
}

table {
  border-collapse: collapse;
  border-spacing: 0;
  empty-cells: show;
  border: 1px solid #cbcbcb;
}

thead {
  background-color: rgb(117, 117, 117);
  color: #FFF;
  text-align: left;
  vertical-align: bottom;
}

th {
  border-left: 1px solid #cbcbcb;
  border-width: 0 0 0 1px;
  font-size: inherit;
  text-align: center;
  margin: 0;
  overflow: visible;
  padding: 0.5em 1em;
}

td {
  text-align:center;
  padding: 10px;
  color: black;
}

tr:nth-child(even) {
  background-color: #d3d3d3;
}

.responsive {
  width: 100%;
  height: auto;
}

/*
  Used for <tabs>
  From https://github.com/spatie/vue-tabs-component/blob/master/docs/resources/tabs-component.css
*/
.tabs-component {
  margin: 4em 0;
}

.tabs-component-tabs {
  border: solid 1px #ddd;
  border-radius: 6px;
  margin-bottom: 5px;
  display: flex;
  justify-content: center;
}

@media (min-width: 700px) {
  .tabs-component-tabs {
    border: 0;
    align-items: stretch;
    display: flex;
    justify-content: center;
    margin-bottom: -1px;
    padding: 0;
  }
}

.tabs-component-tab {
  color: #999;
  font-size: 14px;
  font-weight: 600;
  margin-right: 0;
  list-style: none;
}

.tabs-component-tab:not(:last-child) {
  border-bottom: dotted 1px #ddd;
}

.tabs-component-tab:hover {
  color: #666;
}

.tabs-component-tab.is-active {
  color: #000;
}

.tabs-component-tab.is-disabled * {
  color: #cdcdcd;
  cursor: not-allowed !important;
}

@media (min-width: 700px) {
  .tabs-component-tab {
    background-color: #fff;
    border: solid 1px #ddd;
    border-radius: 3px 3px 0 0;
    margin-right: .5em;
    transform: translateY(2px);
    transition: transform .3s ease;
  }

  .tabs-component-tab.is-active {
    border-bottom: solid 1px #fff;
    z-index: 2;
    transform: translateY(0);
  }
}

.tabs-component-tab-a {
  align-items: center;
  color: inherit;
  display: flex;
  padding: .5em .5em;
  text-decoration: none;
}

.tabs-component-panels {
  padding: 4em 0;
}

@media (min-width: 700px) {
  .tabs-component-panels {
    border-top-left-radius: 0;
    background-color: #fff;
    border: solid 1px #ddd;
    border-radius: 0 6px 6px 6px;
    box-shadow: 0 0 10px rgba(0, 0, 0, .05);
    padding: 4em 2em;
  }
}

</style>
