<script setup lang="ts">
  import { ERouterName } from '@/helpers/enums/RouterName.enum';

</script>


<template>
 <nav class="main-theme p-4 h-full grid grid-cols-7 w-full">
 
 
        <router-link :to="{name:ERouterName.HOME}"  class="flex items-center space-x-4 col-start-2 col-span-1">
          <span class="title text-white">Enron Mail</span>
          <i class="icon fas fa-regular fa-envelope text-white text-2xl"></i>
        </router-link>
    
        
  
  </nav>
</template>

<style scoped> 
  .title{
    color: #fff;
     font-family: sans-serif;
     font-size: 25px; font-weight: normal; 
     text-transform: uppercase; 
     text-shadow: 2px 2px 0 #000; 
   
     text-align: center;
  }

  .icon{
    transform: rotateZ(344deg) rotateX(367deg) translate3d(-10px, -6px, 9px);
  }
  

</style>



<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
