<template>
  <div id="app">
    <nav>
      <ul>
        <li><router-link to="/">TODAY DEALS</router-link></li>
        <li><router-link to="/about">SAVED</router-link></li>
      </ul>
    </nav>
    <router-view/>
  </div>
</template>

<style src="./assets/css/main.css"></style>
