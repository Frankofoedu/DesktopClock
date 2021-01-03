<template>
  <div id="app">
    <!-- <img alt="Wails logo" src="./assets/images/logo.png" class="logo zoomIn"> -->
    <Clock/>
  </div>
</template>

<script>
import Clock from "./components/Clock";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    Clock
  }
};
</script>
