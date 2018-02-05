<template>
  <div id="app">
    <BodyContent/>
  </div>
</template>

<script>
import BodyContent from './components/BodyContent'

export default {
  name: 'App',
  components: {
    BodyContent
  }
}
</script>

<style>
#app {
  /*font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;*/
  color: #2c3e50;
  margin: 0px;
  padding: 0px;
}
</style>
