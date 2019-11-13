<template>
  <tr>
    <td>{{ name }}</td>
    <td>{{ value }}</td>
  </tr>
</template>

<script>
export default {
  name: 'Row',
  props: {
    name: String,
    value: String,
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
</style>
