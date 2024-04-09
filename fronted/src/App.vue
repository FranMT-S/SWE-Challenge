<template>
 
  <router-view/>
</template>

<style>

.main-theme{
    background: #490075;
    background-image: linear-gradient(0.25turn, #490075, #41086a);
  }

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}

/*  Scrollbar */

/* width */
::-webkit-scrollbar {
  width: 10px;
}

/* Track */
::-webkit-scrollbar-track {
  background: #a089ec5c;
  border-radius: 10px;

}

/* Handle */
::-webkit-scrollbar-thumb {
  background: #9557b4c4;
  border-radius: 10px;
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
  background: #671192c4;
}

</style>
