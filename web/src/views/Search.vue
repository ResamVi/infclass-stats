<template>
  <div class="content">
    <h3>Type in player name</h3>
    <h6>(nameless tee /= Nameless Tee)</h6>
    <div>
        <input type="text" v-model="input" @keyup.enter="visitProfile" />
        <button @click="visitProfile">Search</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Player',

  data() {
    return {
      input: '',
    };
  },

  methods: {
    visitProfile() {
      this.$router.push(`player/${this.input}`);
    },
  },
};
</script>


<style scoped>
h1 {
  color: #666666;
}

h3 {
    margin-bottom: 0;
}

h6 {
    margin-top: 0;
}

.content {
    border: solid 1px #ddd;
    border-radius: 6px;
    display: flex;
    justify-content: center;
    flex-direction: column;
    background-color: #fff;
    width: 80%;
    height:30vh;
    margin: 0 auto;
}

</style>

