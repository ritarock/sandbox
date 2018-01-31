<template>
  <div id="app">
    <TopContent/>
    <BodyContent/>
  </div>
</template>

<script>
import TopContent from './components/TopContent'
import BodyContent from './components/BodyContent'

export default {
  name: 'App',
  components: {
    TopContent,
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
