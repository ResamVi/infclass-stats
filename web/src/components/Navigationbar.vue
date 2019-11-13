<template>
    <nav>
        <router-link tag="a" to="/">Stats</router-link>
        <router-link tag="a" to="/search">Search</router-link>
        <router-link tag="a" to="/about">About</router-link>
    </nav>
</template>

<script>
export default {
  name: 'Navigationbar',
};
</script>

<style scoped>
nav {
    background-color: #333;
    overflow: hidden;
    display: flex;
    justify-content: center;
}

a {
  float: left;
  color: #f2f2f2;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 1.2em;
}

a:hover {
  background-color: #ddd;
  color: black;
}

</style>
